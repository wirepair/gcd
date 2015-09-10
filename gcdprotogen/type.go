package main

import (
	"fmt"
	"strings"
)

type Type struct {
	protoType      *ProtoType
	domain         string
	Name           string // the name of the API call, Event or Type
	RefName        string // name to this type as a reference (domain.type)
	Description    string // the description/comments for the struct
	UnderlyingType string
	EnumVals       string // if it's an enum string list out the possible values as a comment
	Properties     []*TypeProperties
}

type TypeProperties struct {
	Name           string // property name
	Description    string // property description
	UnderlyingType string

	Optional bool   // is this property optional?
	Ref      string // reference to other type

}

func NewTypeProperties(name, description, ref string, optional bool) *TypeProperties {
	return &TypeProperties{Name: name, Description: description, Ref: ref, Optional: optional}
}

func NewType(domain string, protoType *ProtoType) *Type {
	t := &Type{}
	t.protoType = protoType
	t.RefName = domain + "." + protoType.Id
	t.Name = protoType.Id
	t.Description = protoType.Description
	t.UnderlyingType = protoType.Type
	t.Properties = make([]*TypeProperties, 0)
	return t
}

func (t *Type) SetUnderlyingType() {
	t.UnderlyingType = getGoType(t.protoType.Type)
	// if type is enum, set enumvals for comment output
	if t.protoType.Type == "enum" {
		t.EnumVals = strings.Join(t.protoType.Enum, ", ")
	}
}

func (t *Type) parseObject() error {
	if len(t.protoType.Properties) == 0 {
		fmt.Printf("%s is an object with no properties.", t.Name)
	}
	return nil
}

func (t *Type) createObjectType(baseTypes map[string]string) error {
	for _, p := range t.protoType.Properties {
		newProp := NewTypeProperties(p.Name, p.Description, ref, optional)
	}
	return nil
}

func (t *Type) createArrayType(baseTypes map[string]string) error {
	return nil
}
