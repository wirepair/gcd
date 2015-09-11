package main

import (
	"fmt"
	"strings"
)

type TypeProperties struct {
	ParentType     *Type // the type ref this property belongs to
	protoProperty  *ProtoProperty
	Name           string // property name
	Description    string // property description
	UnderlyingType string
	GoType         string
	IsArray        bool
	Optional       bool   // is this property optional?
	EnumVals       string // possible enum values as a string
	Ref            string // reference to other type
	ChildProps     []*TypeProperties
}

func NewTypeProperties(parentType *Type, props *ProtoProperty) *TypeProperties {
	tp := &TypeProperties{}
	tp.ParentType = parentType
	tp.protoProperty = props
	tp.Name = props.Name
	tp.Description = props.Description
	tp.Ref = props.Ref
	tp.Optional = props.Optional
	tp.UnderlyingType = props.Type
	tp.ChildProps = make([]*TypeProperties, 0)
	return tp
}

func (p *TypeProperties) SetGoType() {
	p.GoType = getGoType(p.UnderlyingType)
	// if type is enum, set enumvals for comment output
	if p.UnderlyingType == "enum" {
		p.EnumVals = strings.Join(p.protoProperty.Enum, ", ")
	}
}

func (p *TypeProperties) ResolveRefs(baseTypes map[string]string) {
	// No references
	if p.Ref == "" {
		return
	}

	// check if this reference is to a type that is of a base type (string/int/float etc)
	// we do this so we don't have a bunch of type ChromeSomething string which is annoying
	// to work with.
	key := p.ParentType.Domain + "." + p.Ref
	if goType, ok := baseTypes[key]; ok {
		p.GoType = goType
		p.Ref = "" // unset it as being a reference
		return
	}

	// See if external ref, if so remove .
	if strings.Contains(p.Ref, ".") {
		p.GoType = fmt.Sprintf("*Chrome%s", strings.Replace(p.Ref, ".", "", 1))
	} else {
		// set it to DomainRef (So, PageFrame)
		p.GoType = fmt.Sprintf("*Chrome%s", p.ParentType.Domain+p.Ref)
	}

}
