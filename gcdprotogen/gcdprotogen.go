package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//"os"
	"io/ioutil"
	"log"
	"strings"
)

// Top level API protocol file
type DebuggerApi struct {
	Version *ApiVersion `json:"version"`
	Domains []*Domain   `json:"domains"`
}

// Version information
type ApiVersion struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

// The Domain (contains all objects, their type/commands/events)
type Domain struct {
	Domain      string     `json:"domain"`
	Description string     `json:"description,omitempty"`
	Types       []*Type    `json:"types,omitempty"`
	Commands    []*Command `json:"commands,omitempty"`
	Events      []*Event   `json:"events,omitempty"`
	Hidden      bool       `json:"hidden,omitempty"`
	Items       *Item      `json:"items,omitempty"`
}

// A Type which represents objects specific to the API method
type Type struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
	Properties  []*Property `json:"properties,omitempty"`
	Hidden      bool        `json:"hidden,omitempty"`
	Items       *Item       `json:"items,omitempty"`
	MinItems    int64       `json:"minItems,omitempty"`
	MaxItems    int64       `json:"maxItems,omitempty"`
}

// A property & Parameter type used by both commands & types
type Property struct {
	Name        string   `json:"name"`
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	Ref         string   `json:"$ref,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	Hidden      bool     `json:"hidden,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Items       *Item    `json:"items,omitempty"`
}

// An item used by types, properties and events.
type Item struct {
	Type        string      `json:"type,omitempty"`
	Ref         string      `json:"$ref,omitempty"`
	Properties  []*Property `json:"properties,omitempty"`
	Description string      `json:"description,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
}

// The API Command call.
type Command struct {
	Name        string            `json:"name"`
	Type        string            `json:"type,omitempty"`
	Description string            `json:"description,omitempty"`
	Handlers    []string          `json:"handlers,omitempty"`
	Parameters  []*Property       `json:"parameters,omitempty"`
	Returns     []*CommandReturns `json:"returns,omitempty"`
	Hidden      bool              `json:"hidden,omitempty"`
	Async       bool              `json:"async,omitempty"`
	Redirect    string            `json:"redirect,omitempty"`
}

// Parameters to the API Command call
type CommandParameters struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Ref  string `json:"$ref,omitempty"`
}

// The return parameters for an API call
type CommandReturns struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Ref         string `json:"$ref,omitempty"`
	Description string `json:"description,omitempty"`
}

// An event, asynchornous events that can come in once
// enabled.
type Event struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  []*Property `json:"parameters,omitempty"`
	Deprecated  bool        `json:"deprecated,omitempty"`
	Hidden      bool        `json:"hidden,omitempty"`
}

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
type GeneratedStruct struct {
	Name        string   // the name of the API call, Event or Type
	Description string   // the description/comments for the struct
	StructStart string   // the struct start line
	Lines       []string // the individual struct lines
	StructEnd   string   // the end struct line
}

const (
	prefix       = "chrome_"
	structStart  = "type Chrome%s struct {"
	structEnd    = "}"
	structMember = "\t%s %s `json:\"%s\"` %s"
	//eventSuffix = "_events.go" // file used for saving event data structures
	//typeSuffix = "_types.go" // file used for saving event data structures
	//commandSuffix = "_commands.go" // file used for saving event data structures
)

func init() {
	flag.StringVar(&file, "file", "protocol.json", "open remote debugger protocol file.")
}

func main() {
	flag.Parse()
	fmt.Printf("Opening %s for reading...\n", file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
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
		}
	}
}

func NewOutput(major, minor, domain string) {
	filename := prefix + strings.ToLower(domain) + ".go"
	output := &Output{Major: api.Version.Major, Minor: api.Version.Minor, Domain: domain, Filename: filename}
	output.Types = make([]*GeneratedStruct, 0)
	output.Commands = make([]*GeneratedStruct, 0)
	output.Events = make([]*GeneratedStruct, 0)
	return output
}

func NewGeneratedStruct(apiName string, description string) *GeneratedStruct {
	s := &GeneratedStruct{ApiName: apiName, Description: "//" + description}
	s.StructEnd = structEnd
	return s
}

func createTypes(output *Output, types []*Type) error {
	for _, v := range types {
		t := NewGeneratedStruct(output.Domain, v.Description)
		t.Name = v.Id // Type Name
		t.Lines = make([]string)
		if v.Type == "object" {
			for _, l := range v.Properties {
				optional := ""
				if l.Optional == true {
					optional = ",omitempty"
				}
				line := fmt.Sprintf(structMember, l.Name, l.Type, l.Name, l.Optional)
			}
		}

		fmt.Printf("type: %v\n", v)
	}
	return nil
}

func createMethods(api *DebuggerApi) error {
	return nil
}

func jsonType(type string) {
	
}