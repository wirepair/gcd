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

	flag.Parse()
	protocolData := openFile()
	//createOutputDirectory()

	protocolApi := unmarshalProtocol(protocolData)
	major := protocolApi.Version.Major
	minor := protocolApi.Version.Minor

	for _, proto := range protocolApi.Domains {
		domain := NewDomain(major, minor, proto.Domain)
		fmt.Printf("Creating api for domain: %s\n", proto.Domain)
		if proto.Types != nil && len(proto.Types) > 0 {
			domain.PopulateTypes(proto.Types)
		}
		domains = append(domains, domain)
	}

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
