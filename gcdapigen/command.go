package main

type Command struct {
	Name               string
	Description        string
	Parameters         []*TypeProperties
	Returns            []*Return
	HasParams          bool
	HasReturn          bool
	NoParamReturnCalls bool
	ParamCalls         bool
	ReturnCalls        bool
	ParamReturnCalls   bool
}

func NewCommand(protoCommand *ProtoCommand) *Command {
	c := &Command{}
	c.Name = protoCommand.Name
	c.Description = protoCommand.Description
	if protoCommand.Parameters != nil && len(protoCommand.Parameters) > 0 {
		c.HasParams = true
	}

	if protoCommand.Returns != nil && len(protoCommand.Returns) > 0 {
		c.HasReturn = true
	}
	// Determine type of call for template output
	if c.HasParams == false && c.HasReturn == false {
		c.NoParamReturnCalls = true
	}

	if c.HasParams == true && c.HasReturn == false {
		c.ParamCalls = true
	}

	if c.HasParams == false && c.HasReturn == true {
		c.ReturnCalls = true
	}

	if c.HasParams == true && c.HasReturn == true {
		c.ParamReturnCalls = true
	}

	c.Returns = make([]*Return, 0)
	c.Parameters = make([]*TypeProperties, 0)
	return c
}
