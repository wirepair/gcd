package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var file string

type Output struct {
	Major    string // major api version
	Minor    string // minor api version
	Filename string
	Domain   string
	Types    []*GeneratedStruct
	Events   []*GeneratedStruct
	Commands []*GeneratedStruct
}

func NewOutput(major, minor, domain string) *Output {
	filename := prefix + strings.ToLower(domain)
	output := &Output{Major: major, Minor: minor, Domain: domain, Filename: filename}
	output.Types = make([]*GeneratedStruct, 0)
	output.Commands = make([]*GeneratedStruct, 0)
	output.Events = make([]*GeneratedStruct, 0)
	return output
}

type GeneratedStruct struct {
	Out            *Output  // reference to the output object
	Name           string   // the name of the API call, Event or Type
	Description    string   // the description/comments for the struct
	StructStart    string   // the struct start line
	Lines          []string // the individual struct lines
	NonObjectLines string   // non object type
	StructEnd      string   // the end struct line
	IsObject       bool     // is the type an object?
	EnumVals       string   // if it's an enum string list out the possible values as a comment
}

func NewGeneratedStruct(output *Output, apiMethod, description string) *GeneratedStruct {
	structName := output.Domain + apiMethod
	s := &GeneratedStruct{Out: output, Name: structName, Description: description}
	s.Lines = make([]string, 0)
	s.StructEnd = structEnd
	return s
}

func (g *GeneratedStruct) String() string {
	var text string
	if g.IsObject {
		text = fmt.Sprintf(structStart, g.Description, g.Name)
		for _, line := range g.Lines {
			text += line + "\n"
		}
		text += structEnd + "\n"
	} else {
		text = g.NonObjectLines
	}
	return text
}

const (
	prefix       = "chrome_"
	nonStruct    = "// %s\ntype Chrome%s %s %s\n"    // description, Name, type, possible enums
	structStart  = "// %s\ntype Chrome%s struct {\n" // description, Name
	structEnd    = "}"
	structMember = "\t%s %s `json:\"%s\"` %s" // property name, type, json name{omitempty}, description
)

var templates *template.Template

func init() {
	var err error
	flag.StringVar(&file, "file", "protocol.json", "open remote debugger protocol file.")
	templates, err = template.ParseFiles("type_template.txt")
	if err != nil {
		log.Fatalf("error opening templates: %s\n", err)
	}
}

func main() {
	flag.Parse()
	fmt.Printf("Opening %s for reading...\n", file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	err = makeDirectories()

	api := &DebuggerApi{}
	err = json.Unmarshal(b, api)
	if err != nil {
		switch u := err.(type) {
		case *json.SyntaxError:
			log.Fatalf("syntax error at offset %d: %s\n", u.Offset, u)
		case *json.UnmarshalTypeError:
			log.Fatal("unmarshal type err: " + err.Error() + " " + err.(*json.UnmarshalTypeError).Value)
		case *json.InvalidUnmarshalError:
			log.Fatal("InvalidUnmarshalError: " + err.Error())
		default:
			log.Fatal("UnmarshalError: " + err.Error())
		}
	}

	for _, v := range api.Domains {
		output := NewOutput(api.Version.Major, api.Version.Minor, v.Domain)
		fmt.Printf("%s\n", v.Domain)
		if v.Types != nil {
			err := createTypes(output, v.Types)
			if err != nil {
				log.Fatal(err)
			}
			writeTypes(output)
		}

	}
}

func makeDirectories() error {
	dirs := []string{"types", "events", "commands"}
	for _, d := range dirs {
		err := os.Mkdir(d, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

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
			// we are an array of basic types, recall ourselves with the *items* type.
			return "[]" + jsonType(domain, items, "", items.Type)
		}
	default:
		if items != nil && items.Ref != "" {
			return "*Chrome" + filterRef(domain, items.Ref)
		}
	}
	return valType
}

// if a . exists, we need to remove
// if a . doesn't exist that mean this type references a type with in our own domain.
func filterRef(currDomain, ref string) string {
	if !strings.Contains(ref, ".") {
		return currDomain + ref
	}
	return strings.Replace(ref, ".", "", -1)
}
