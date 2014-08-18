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
var funcMap template.FuncMap

func modifyReserved(input string) string {
	switch input {
	case "type":
		return "theType"
	case "range":
		return "theRange"
	case "interface":
		return "theInterface"
	case "for":
		return "theFor"
	}
	return input
}

func nullType(input string) string {
	if strings.Contains(input, "types.") {
		return "nil"
	}
	if strings.Contains(input, "[]") {
		return "nil"
	}
	fmt.Printf("INPUT: %s\n", input)
	switch input {
	case "int":
		return "0"
	case "float64":
		return "0"
	case "string":
		return "\"\""
	case "array":
		return "[]"
	case "bool":
		return "false"
	case "interface{}":
		return "nil"
	}
	return ""
}

func init() {
	flag.StringVar(&file, "file", "protocol.json", "open remote debugger protocol file.")
	funcMap := template.FuncMap{
		"Title":    strings.Title,
		"ToLower":  strings.ToLower,
		"Reserved": modifyReserved,
		"NullType": nullType,
	}
	// kinda dumb. but need to map functions first, but need a named template first and it must match the first template in ParseFiles.
	templates = template.Must(template.New("type_template.txt").Funcs(funcMap).ParseFiles("type_template.txt", "code_template.txt"))
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

		if v.Types != nil && len(v.Types) > 0 {
			err := createTypes(output, v.Types)
			if err != nil {
				log.Fatal(err)
			}
			writeTypes(output)
		}

		if v.Commands != nil && len(v.Commands) > 0 {
			err := createCommands(output, v.Commands)
			if err != nil {
				log.Fatal(err)
			}
		}
		/*
			if v.Events != nil && len(v.Events) > 0 {
					err := createEvents(output, v.Events)
				}
		*/
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

// if a . exists, we need to remove
// if a . doesn't exist that mean this type references a type with in its own domain.
func filterRef(currDomain, ref string) string {
	if !strings.Contains(ref, ".") {
		return currDomain + ref
	}
	return strings.Replace(ref, ".", "", -1)
}

func jsonType(domain string, items *Item, ref, valType string, typeRequired bool) string {
	ret := ""
	if typeRequired {
		ret += "types."
	}
	// if we are a reference type, use that.
	if ref != "" {
		return "*" + ret + "Chrome" + filterRef(domain, ref)
	}

	switch valType {
	// No idea what any type means, see DOM.ShapeOutsideInfo (treat as string???)
	case "any":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		// we are an array, see if *that* is a ref type.
		fmt.Printf("type is array: %#v\n", items)
		if items != nil && items.Ref != "" {
			return "[]*" + ret + "Chrome" + filterRef(domain, items.Ref)
		} else {
			if items == nil {
				return "[]" + jsonType(domain, items, "", "", typeRequired)
			}
			// we are an array of basic types, recursively call with the *item's* type.
			return "[]" + jsonType(domain, items, "", items.Type, typeRequired)
		}
	default:
		if items != nil && items.Ref != "" {
			return "*" + ret + "Chrome" + filterRef(domain, items.Ref)
		}
	}
	return valType
}
