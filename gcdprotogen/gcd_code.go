package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type GeneratedApi struct {
	Domain             string
	Filename           string
	Major              string
	Minor              string
	TypesRequired      bool // does this api require types to be imported?
	JsonRequired       bool // does this api require json to be imported?
	NoParamReturnCalls []*GeneratedNoParamReturnCommand
	ParamCalls         []*GeneratedParamCommand
	ReturnCalls        []*GeneratedReturnCommand
	ParamReturnCalls   []*GeneratedReturnParamCommand
}

// For commands which have no parameters and no return values.
type GeneratedNoParamReturnCommand struct {
	Description string // the description/comments for the command
	Method      string // the name of the API call, Event or Type
}

// For commands which have parameters
type GeneratedParamCommand struct {
	Description       string // the description/comments for the command
	Method            string // the name of the API call, Event or Type
	ParamDefinition   string
	ParamLen          int
	ParamTypes        []string // The parameter types
	ParamNames        []string
	ParamDescriptions []string // the description/comments for parameters
}

// For commands which just return types
type GeneratedReturnCommand struct {
	Description        string // the description/comments for the command
	Method             string // the name of the API call, Event or Type
	Returns            []*Return
	ReturnDescriptions []string // the description/comments for return types
}

type Return struct {
	Name string // the name of the return value
	Type string // the type of return value
}

// For commands which have parameters and return types
type GeneratedReturnParamCommand struct {
	Description        string // the description/comments for the command
	Method             string // the name of the API call, Event or Type
	Params             []string
	Returns            []string
	ParamDescriptions  []string // the description/comments for parameters
	ReturnDescriptions []string // the description/comments for return types
}

type ParamInterface interface {
	AddParam(paramName, paramType, paramDescription string)
	FormalizeDefinition()
}

type ReturnInterface interface {
	AddReturn(returnName, returnType, returnDescription string)
}

func (grc *GeneratedReturnCommand) AddReturn(returnName, returnType, returnDescription string) {
	grc.Returns = append(grc.Returns, &Return{Name: returnName, Type: returnType})
	if returnDescription != "" {
		grc.ReturnDescriptions = append(grc.ReturnDescriptions, returnDescription)
	}
}

func (gpc *GeneratedParamCommand) AddParam(paramName, paramType, description string) {
	gpc.ParamNames = append(gpc.ParamNames, paramName)
	gpc.ParamTypes = append(gpc.ParamTypes, paramType)
	gpc.ParamDescriptions = append(gpc.ParamDescriptions, paramName+" - "+description)
}

func (gpc *GeneratedParamCommand) FormalizeDefinition() {
	gpc.ParamDefinition = ""
	gpc.ParamLen = len(gpc.ParamNames)
	for k, v := range gpc.ParamNames {
		v = modifyReserved(v)
		gpc.ParamDefinition += v + " " + gpc.ParamTypes[k] + ", "
	}
	gpc.ParamDefinition = gpc.ParamDefinition[:len(gpc.ParamDefinition)-2] // remove ,
}

func NewGeneratedApi(output *Output) *GeneratedApi {
	return &GeneratedApi{Domain: output.Domain, Filename: output.Filename, Major: output.Major, Minor: output.Minor}
}

func writeCommands(api *GeneratedApi) {
	wr, err := os.Create("commands" + string(os.PathSeparator) + api.Filename + "_commands.go")
	if err != nil {
		log.Fatalf("Error creating output file, %s: %s\n", api.Filename, err)
	}

	err = templates.ExecuteTemplate(wr, "code_template.txt", api)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}
}

// Iterates over the Domain -> commands section of protocol.json
func createCommands(output *Output, commands []*Command) error {
	api := NewGeneratedApi(output)
	for _, c := range commands {
		//cmd := NewGeneratedCommand(output, c.Name, c.Description)
		if len(c.Parameters) == 0 && len(c.Returns) == 0 {
			//fmt.Println("Name: " + c.Name + " Description: " + c.Description)
			cmd := &GeneratedNoParamReturnCommand{Description: c.Description, Method: c.Name}
			api.NoParamReturnCalls = append(api.NoParamReturnCalls, cmd)
		}
		if len(c.Parameters) != 0 && len(c.Returns) == 0 {
			cmd := &GeneratedParamCommand{Description: c.Description, Method: c.Name}
			for _, p := range c.Parameters {
				getParameters(api, output.Domain, cmd, p)
			}
			cmd.FormalizeDefinition()
			api.ParamCalls = append(api.ParamCalls, cmd)
		} else if len(c.Parameters) == 0 && len(c.Returns) != 0 {
			cmd := &GeneratedReturnCommand{Description: c.Description, Method: c.Name}
			for _, r := range c.Returns {
				fmt.Printf("METHOD: %s\n", c.Name)
				getReturnTypes(api, output.Domain, cmd, r)
			}
			api.ReturnCalls = append(api.ReturnCalls, cmd)
		}
		/*
			if len(c.Parameters) != 0 {
				param := make([]string, len(c.Parameters))
				paramDescription := make([]string, len(c.Parameters))

			for _, p := range c.Parameters {
				fmt.Println("Parameters: " + p.Name + " " + p.Type)
			}
			for _, r := range c.Returns {
				fmt.Println("Returns: " + r.Name + " " + r.Type)
			}
		*/
	}
	writeCommands(api)
	return nil
}

func getParameters(api *GeneratedApi, domain string, v ParamInterface, p *Property) {
	json := jsonType(domain, p.Items, p.Ref, p.Type, true)
	if api.TypesRequired == false && strings.Contains(json, "types.") {
		api.TypesRequired = true
	}
	v.AddParam(p.Name, json, p.Description)
}

func getReturnTypes(api *GeneratedApi, domain string, v ReturnInterface, r *CommandReturns) {
	api.JsonRequired = true
	json := jsonType(domain, r.Items, r.Ref, r.Type, true)
	if r.Type == "array" {
		fmt.Printf("r.Items: %#v\n", r.Items)
	}
	fmt.Printf("jsonType: %s Ref: %s Type: %s\n", json, r.Ref, r.Type)
	// having problems with int in map[string]interface{} json Unmarshalling...
	if json == "int" {
		json = "float64"
	}
	v.AddReturn(r.Name, json, r.Description)
}
