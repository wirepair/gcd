package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func writeTypes(output *Output) {
	wr, err := os.Create("types" + string(os.PathSeparator) + output.Filename + "_types.go")
	if err != nil {
		log.Fatalf("Error creating output file, %s: %s\n", output.Filename, err)
	}

	err = templates.ExecuteTemplate(wr, "type_template.txt", output)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}
}

func createTypes(output *Output, types []*Type) error {
	for _, t := range types {
		struc := NewGeneratedStruct(output, t.Id, t.Description)
		if t.Type == "object" {
			struc.IsObject = true
			createObjectType(struc, t.Properties)
		} else {
			if t.Enum != nil && len(t.Enum) > 0 {
				struc.EnumVals = "// possible values: " + strings.Join(t.Enum, ", ")
			}
		}
		output.Types = append(output.Types, struc)
		if !struc.IsObject {
			if t.Type == "array" {

			}
			struc.NonObjectLines = fmt.Sprintf(nonStruct, struc.Description, struc.Name, jsonType(output.Domain, t.Items, "", t.Type), struc.EnumVals)
		}
	}
	return nil
}

func createObjectType(struc *GeneratedStruct, props []*Property) {
	for _, prop := range props {
		optional := ""
		description := ""
		if prop.Optional == true {
			optional = ",omitempty"
		}

		if prop.Description != "" {
			description = "// " + prop.Description
		}
		newStruc := hasNestedType(struc, prop)
		realType := jsonType(struc.Out.Domain, prop.Items, prop.Ref, prop.Type)
		if newStruc != nil {
			struc.Out.Types = append(struc.Out.Types, newStruc)
			realType = "*Chrome" + filterRef("", newStruc.Name) // don't pass domain as we already have it from the parent type
			if prop.Type == "array" {
				realType = "[]" + realType
			}
		}
		// structMember = "\t%s %s `json:\"%s\"` %s" // property name, type, json name{omitempty}, description
		jsonField := prop.Name + optional
		line := fmt.Sprintf(structMember, strings.Title(prop.Name), realType, jsonField, description)
		struc.Lines = append(struc.Lines, line)

	}
}

// checks if the type is nested, which is an un-named object (we create a sub-resource name for it)
func hasNestedType(struc *GeneratedStruct, prop *Property) *GeneratedStruct {
	if prop.Items == nil || prop.Items.Type != "object" {
		return nil
	}
	newName := strings.Title(prop.Name) // take the parent object name, and add our name to it to generate a sub-resource type.
	newStruc := NewGeneratedStruct(struc.Out, newName, prop.Description)
	newStruc.IsObject = true
	createObjectType(newStruc, prop.Items.Properties)
	return newStruc
}

func createMethods(api *DebuggerApi) error {
	return nil
}

func jsonType(domain string, items *Item, ref, valType string) string {
	// if we are a reference type, use that.
	if ref != "" {
		return "*Chrome" + filterRef(domain, ref)
	}

	switch valType {
	// No idea what any type means, see DOM.ShapeOutsideInfo (treat as string???)
	case "any":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float"
	case "boolean":
		return "bool"
	case "array":
		// we are an array, see if *that* is a ref type.
		if items != nil && items.Ref != "" {
			return "[]*Chrome" + filterRef(domain, items.Ref)
		} else {
			// we are an array of basic types, recursively call with the *item's* type.
			return "[]" + jsonType(domain, items, "", items.Type)
		}
	default:
		if items != nil && items.Ref != "" {
			return "*Chrome" + filterRef(domain, items.Ref)
		}
	}
	return valType
}
