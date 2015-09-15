package main

import (
	"strings"
)

type Type struct {
	protoType      *ProtoType
	Name           string // the name of the API call, Event or Type
	Description    string // the description/comments for the struct
	GoType         string // the type this would be in Go
	UnderlyingType string // the type defined in protocol.json
	EnumVals       string // if it's an enum string list out the possible values as a comment
	IsSubType      bool   // is this a sub type? (Should be prefixed with Sub in template)
	Properties     []*TypeProperties
}

func NewType(protoType *ProtoType) *Type {
	t := &Type{}
	t.protoType = protoType
	t.Name = protoType.Id
	t.Description = protoType.Description
	if t.Description == "" {
		t.Description = "No Description."
	}
	t.UnderlyingType = protoType.Type
	t.Properties = make([]*TypeProperties, 0)
	return t
}

func NewSubType(parentProps *TypeProperties, protoProps *ProtoProperty) *Type {
	st := &Type{}
	st.IsSubType = true
	// only convert to type if it's an array
	if protoProps.IsArray() {
		st.protoType = typeFromProperties(protoProps)
	}

	st.Name = "Sub" + strings.Title(parentProps.Name)
	st.Properties = make([]*TypeProperties, 0)
	st.UnderlyingType = parentProps.UnderlyingType
	return st
}

func (t *Type) IsNonPropertiesObject() bool {
	return (t.UnderlyingType == "object" && len(t.protoType.Properties) == 0)
}

func (t *Type) GetUnderlyingType() string {
	return t.UnderlyingType
}

func (t *Type) IsArray() bool {
	return t.UnderlyingType == "array"
}

func (t *Type) GetArrayType() string {
	if t.protoType.Items.Type != "" {
		return t.protoType.Items.Type
	}

	if t.protoType.Items.Ref != "" {
		return t.protoType.Items.Ref
	}
	return "object"
}
