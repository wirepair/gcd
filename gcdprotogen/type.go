package main

import (
	"strings"
)

type Type struct {
	protoType      *ProtoType
	Domain         string
	Name           string // the name of the API call, Event or Type
	RefName        string // name to this type as a reference (domain.type)
	Description    string // the description/comments for the struct
	GoType         string // the type this would be in Go
	UnderlyingType string // the type defined in protocol.json
	EnumVals       string // if it's an enum string list out the possible values as a comment
	Properties     []*TypeProperties
}

func NewType(domain string, protoType *ProtoType) *Type {
	t := &Type{}
	t.protoType = protoType
	t.RefName = domain + "." + protoType.Id
	t.Name = protoType.Id
	t.Description = protoType.Description
	t.Domain = domain
	t.UnderlyingType = protoType.Type
	t.Properties = make([]*TypeProperties, 0)
	return t
}

func NewSubType(domain string, parentProps *TypeProperties, protoProps *ProtoProperty) *Type {
	st := &Type{}
	st.Domain = domain
	st.Name = "Sub" + parentProps.Name
	st.Properties = make([]*TypeProperties, 0)
	st.UnderlyingType = parentProps.UnderlyingType
	//st.Properties = protoProps
	return st
}

func (t *Type) SetGoType() {
	t.GoType = getGoType(t.protoType.Type)
	// if type is enum, set enumvals for comment output
	if t.UnderlyingType == "enum" {
		t.EnumVals = strings.Join(t.protoType.Enum, ", ")
	}
}

func (t *Type) createObjectType(baseTypes map[string]string) error {
	for _, p := range t.protoType.Properties {
		newProp := NewTypeProperties(t, p)
		newProp.SetGoType()
		newProp.ResolveRefs(baseTypes)
		t.Properties = append(t.Properties, newProp)
		//if newProp.GoType == "array" {
		//	NewType(t.Domain, protoType)
		//}
	}

	return nil
}

func (t *Type) createArrayType(baseTypes map[string]string) error {
	return nil
}
