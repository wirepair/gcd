package main

type Parameter struct {
	protoParam     *ProtoCommandParameters
	Name           string // property name
	Description    string // property description
	UnderlyingType string
	GoType         string
	Optional       bool   // is this property optional?
	EnumVals       string // possible enum values as a string
	IsRef          bool   // is a reference to another type
}

func NewParameter(protoParam *ProtoCommandParameters) *Parameter {
	p := &Parameter{}
	p.protoParam = protoParam
	p.Name = protoParam.Name
	p.Description = protoParam.Description
	p.UnderlyingType = protoParam.Type
	if protoParam.Ref != "" {
		p.IsRef = true
	}
	return p
}

func (p *Parameter) IsNonPropertiesObject() bool {
	return (p.UnderlyingType == "object" && len(p.protoParam.Properties) == 0)
}

func (p *TypeProperties) GetUnderlyingType() string {
	return p.UnderlyingType
}

func (p *TypeProperties) IsArray() bool {
	return p.UnderlyingType == "array"
}

func (p *TypeProperties) GetArrayType() string {
	if p.protoProperty.Items.Type != "" {
		return p.protoProperty.Items.Type
	}

	if p.protoProperty.Items.Ref != "" {
		return p.protoProperty.Items.Ref
	}
	return "object"
}
