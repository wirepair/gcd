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

type GlobalReference struct {
	LocalRefName    string
	ExternalGoName  string
	IsBaseType      bool
	BaseType        string
	EnumDescription string
}

// Stores a list of all references and if they are base types
var globalRefs map[string]*GlobalReference

const (
	prefix = "chrome_"
)

var file string
var templates *template.Template // for code generation
var funcMap template.FuncMap     // helper funcs

func init() {
	flag.StringVar(&file, "file", "protocol.json", "open remote debugger protocol file.")
	funcMap := template.FuncMap{
		"Title":    strings.Title,
		"ToLower":  strings.ToLower,
		"Reserved": modifyReserved,
		"NullType": nullType,
	}
	templates = template.Must(template.New("type_template.txt").Funcs(funcMap).ParseFiles("type_template.txt"))
}

func main() {
	var domains []*Domain
	globalRefs = make(map[string]*GlobalReference)

	flag.Parse()
	protocolData := openFile()
	//createOutputDirectory()

	protocolApi := unmarshalProtocol(protocolData)
	major := protocolApi.Version.Major
	minor := protocolApi.Version.Minor

	// iterate over the protocol once to resolve references
	for _, proto := range protocolApi.Domains {
		PopulateReferences(proto.Domain, proto.Types)
	}

	for _, protoDomain := range protocolApi.Domains {
		if protoDomain.Domain != "Debugger" {
			continue
		}
		domain := NewDomain(major, minor, protoDomain.Domain)
		fmt.Printf("Creating api for domain: %s\n", protoDomain.Domain)
		// Do types first
		if protoDomain.Types != nil && len(protoDomain.Types) > 0 {
			domain.PopulateTypes(protoDomain.Types)
		}
		// Then Commands
		if protoDomain.Commands != nil && len(protoDomain.Commands) > 0 {
			domain.PopulateCommands(protoDomain.Commands)
		}
		domains = append(domains, domain)
		domain.writeTypes()
	}

	//for k, v := range globalRefs {
	//	fmt.Printf("ref: %s value: %#v\n", k, v)
	//}
}

func openFile() []byte {
	log.Printf("Opening %s for reading...\n", file)
	protocolData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	return protocolData
}

func createOutputDirectory() {
	dir := "commands"
	err := os.Mkdir(dir, 0666)
	if err != nil {
		log.Fatalf("error creating command output directory")
	}
}

func unmarshalProtocol(protocolData []byte) *ProtoDebuggerApi {
	api := &ProtoDebuggerApi{}
	err := json.Unmarshal(protocolData, api)
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
	return api
}
