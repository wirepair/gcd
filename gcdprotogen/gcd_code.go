package main

import (
	"fmt"
	"log"
	"os"
)

type GeneratedCommand struct {
	Out                *Output // reference to the output object
	Domain             string
	Name               string   // the name of the API call, Event or Type
	Description        string   // the description/comments for the command
	ParamDescriptions  []string // the description/comments for parameters
	ReturnDescriptions []string // the description/comments for return types

}

func NewGeneratedCommand(output *Output, apiMethod, description string) *GeneratedStruct {
	structName := output.Domain + apiMethod
	s := &GeneratedStruct{Out: output, Name: structName, Description: description}
	s.Lines = make([]string, 0)
	s.StructEnd = structEnd
	return s
}

func writeCommands(output *Output) {
	wr, err := os.Create("types" + string(os.PathSeparator) + output.Filename + "_commands.go")
	if err != nil {
		log.Fatalf("Error creating output file, %s: %s\n", output.Filename, err)
	}

	err = templates.ExecuteTemplate(wr, "code_template.txt", output)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}
}

// Iterates over the Domain -> commands section of protocol.json
func createCommands(output *Output, commands []*Command) error {
	for _, c := range commands {
		//struc := NewGeneratedStruct(output, c.Name, c.Description)
		fmt.Println("Name: " + c.Name + " Type: " + c.Description)
		/*
			for _, p := range c.Parameters {
				fmt.Println("Parameters: " + p.Name + " " + p.Type)
			}
			for _, r := range c.Returns {
				fmt.Println("Returns: " + r.Name + " " + r.Type)
			}
		*/
	}
	return nil
}
