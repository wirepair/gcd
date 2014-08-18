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

// Iterates over the Domain -> Types section of protocol.json
// Creates a new GeneratedStruct per type. If the type is an object type (has properties)
// calls createObjectType with the list of properties for the object.
// if it's an enum, just converts to string and sets the possible values as code comments
// If the struct is not an object, we need to print out special lines, such as:
// type ChromeSomeSillyThing string
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
			struc.NonObjectLines = fmt.Sprintf(nonStruct, struc.Description, struc.Name, jsonType(output.Domain, t.Items, "", t.Type, false), struc.EnumVals)
		}
	}
	return nil
}

// Iterates over the properties of the object type and attempts to identify
// if any properties have nested object types
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
		realType := jsonType(struc.Out.Domain, prop.Items, prop.Ref, prop.Type, false)
		if realType == "object" {
			fmt.Printf("realType was object for %s in %s setting to interface{}", prop.Name, struc.Out.Domain, prop, newStruc)
			realType = "interface{}"
		}
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
