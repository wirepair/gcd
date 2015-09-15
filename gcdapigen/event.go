package main

type Event struct {
	protoEvent  *ProtoEvent
	Name        string
	Description string
	Parameters  []*TypeProperties
	HasParams   bool
}

func NewEvent(protoEvent *ProtoEvent) *Event {
	e := &Event{}
	e.Name = protoEvent.Name
	e.Description = protoEvent.Description
	if protoEvent.Parameters != nil && len(protoEvent.Parameters) > 0 {
		e.HasParams = true
	}
	e.Parameters = make([]*TypeProperties, 0)
	return e
}
