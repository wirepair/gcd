package main

type Command struct {
	Name        string
	Description string
	Parameters  []*Parameter
	Returns     []*Return
	HasParams   bool
	HasReturn   bool
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
	return c
}
