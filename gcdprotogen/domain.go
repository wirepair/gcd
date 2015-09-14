package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SharedProperties interface {
	GetUnderlyingType() string
	GetArrayType() string
	IsArray() bool
	IsNonPropertiesObject() bool
}

type Domain struct {
	Major    string // major api version
	Minor    string // minor api version
	Filename string
	Domain   string
	Hidden   bool
	SubTypes []*Type
	Types    []*Type
	Events   []*Event
	Commands []*Command
	// basicTypes holds a map of type.RefName and type.Underlying type so we can replace $ref
	// with the underlying type (provided it's not another object or array)

	//typeMap map[string]*BaseType
}

func NewDomain(major, minor, domain string) *Domain {
	d := &Domain{Major: major, Minor: minor, Domain: domain}
	d.Types = make([]*Type, 0)
	d.SubTypes = make([]*Type, 0)
	d.Events = make([]*Event, 0)
	d.Commands = make([]*Command, 0)
	//d.typeMap = make(map[string]*BaseType)
	return d
}

// Extract each type and call handleType, add the result to our Types slice.
func (d *Domain) PopulateTypes(types []*ProtoType) {
	// do first pass to get all underlying type information
	for _, protoType := range types {
		fmt.Printf("Populating type: %s\n", protoType.Id)
		newType := NewType(protoType)
		// igore empty property types as we turn those into Refs
		if len(protoType.Properties) > 0 {
			d.handleType(newType, protoType.Properties)
			d.Types = append(d.Types, newType)
		}

	}
}

func (d *Domain) PopulateCommand(commands []*ProtoCommand) {
	for _, protoCommand := range commands {
		newCmd := NewCommand(protoCommand)
		if newCmd.HasParams {
			d.handleParameters(newCmd, protoCommand.Parameters)
		}
		if newCmd.HasReturn {
			d.handleReturns(newCmd, protoCommand.Returns)
		}
	}
}

func (d *Domain) handleParameters(newCmd *Command, protoParameters []*ProtoCommandParameters) {
	for _, protoParam := range protoParameters
}

func (d *Domain) handleReturns(newCmd *Command, protoReturn []*ProtoCommandReturns) {

}

// Takes in a new type, checks if it is a base type, or an object or an array.
func (d *Domain) handleType(newType *Type, typeProperties []*ProtoProperty) {
	// loop over properties of this new type
	for _, protoProp := range typeProperties {

		// base type, add it and fix up description
		if isBaseType(protoProp) {
			d.createBaseType(newType, protoProp)
			continue
		}

		// It's a reference, see if it points to a base type or not
		if newProp := d.resolveReference(newType, protoProp); newProp != nil {
			newType.Properties = append(newType.Properties, newProp)
			continue
		}

		// is this a subType?
		if isSubType(protoProp) {
			d.createSubType(newType, protoProp)
		}
	}
}

// Creates a base type based property (string, int etc) Also fixes up the description
// to include Enum values if available. Adds the property to newType.
func (d *Domain) createBaseType(newType *Type, protoProp *ProtoProperty) {
	newProp := NewTypeProperties(newType, protoProp)
	newProp.GoType = getGoType(newProp)
	if len(newProp.EnumVals) > 0 {
		newProp.Description = newProp.Description + "Enum values: " + newProp.EnumVals
	}
	newType.Properties = append(newType.Properties, newProp)
	return
}

// Creates a new SubType *Type object. This is for nested structs that are better
// defined outside of the original Type. It will call handleType to iterate over
// the nested properties to create it in much of the same way as a normal type
// except we prefix it with the Sub keyword.
func (d *Domain) createSubType(newType *Type, protoProp *ProtoProperty) {
	// create the property
	newProp := NewTypeProperties(newType, protoProp)
	// create the new sub type
	subType := NewSubType(newProp)
	// recursive to add props to this type
	d.handleType(subType, protoProp.Properties)
	d.SubTypes = append(d.SubTypes, subType)
	// update ref from property to new sub type
	refName := d.Domain + "Sub" + strings.Title(protoProp.Name)
	newProp.GoType = refName
	newProp.IsRef = true
	newType.Properties = append(newType.Properties, newProp)
}

// Since we don't want useless single underlying type struct definitions everywhere we resolve
// any references to single underlying type objects.
func (d *Domain) resolveReference(newType *Type, protoProp *ProtoProperty) *TypeProperties {
	if protoProp.Ref == "" {
		return nil
	}
	refName := ""
	// Local reference, add . between domain/ref type so we can look it up in our globalRefs
	if !strings.Contains(protoProp.Ref, ".") {
		refName = d.Domain + "." + protoProp.Ref
	} else {
		refName = protoProp.Ref
	}

	newProp := NewTypeProperties(newType, protoProp)
	ref := globalRefs[refName]
	// base type
	if ref.IsBaseType {
		newProp.GoType = ref.BaseType
	} else {
		newProp.GoType = ref.ExternalGoName
		newProp.IsRef = true
	}
	// add enum possible values to description
	if ref.EnumDescription != "" {
		newProp.Description = newProp.Description + ref.EnumDescription
	}

	return newProp
}

func (d *Domain) writeTypes() {
	wr := os.Stdout

	err := templates.ExecuteTemplate(wr, "type_template.txt", d)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}
}
