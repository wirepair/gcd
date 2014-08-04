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
		if v.Types != nil && len(v.Types) > 0 {
			err := createTypes(output, v.Types)
			if err != nil {
				log.Fatal(err)
			}
			writeTypes(output)
		}
		if v.Events != nil && len(v.Events) > 0 {
			err := createEvents(output, v.Events)
		}

		if v.Commands != nil && len(v.Commands) > 0 {
			err := createCommands(output, v.Commands)
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

// if a . exists, we need to remove
// if a . doesn't exist that mean this type references a type with in its own domain.
func filterRef(currDomain, ref string) string {
	if !strings.Contains(ref, ".") {
		return currDomain + ref
	}
	return strings.Replace(ref, ".", "", -1)
}
