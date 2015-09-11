package main

import (
	"fmt"
	"log"
	"os"
)

type Domain struct {
	Major    string // major api version
	Minor    string // minor api version
	Filename string
	Domain   string
	Hidden   bool
	Types    []*Type
	Events   []*Event
	Commands []*Command
	// basicTypes holds a map of type.RefName and type.Underlying type so we can replace $ref
	// with the underlying type (provided it's not another object or array)

	// However array types that simply point to another type are also added here.
	basicTypes map[string]string
}

func NewDomain(major, minor, domain string) *Domain {
	d := &Domain{Major: major, Minor: minor, Domain: domain}
	d.Types = make([]*Type, 0)
	d.Events = make([]*Event, 0)
	d.Commands = make([]*Command, 0)
	d.basicTypes = make(map[string]string)
	return d
}

// Do two passes over types, first is to set all the types and sub types
// Second is to resolve references. Store base types (int/string) in a map
// so we don't have to have single typed values like type ChromeSomething int everywhere.
func (d *Domain) PopulateTypes(types []*ProtoType) error {
	// do first pass to get all underlying type information
	for _, protoType := range types {
		fmt.Printf("Populating type: %s\n", protoType.Id)
		newType := NewType(d.Domain, protoType)
		newType.SetGoType()

		// we don't want to create single typed values if we can avoid it.
		if newType.UnderlyingType != "object" && newType.UnderlyingType != "array" {
			d.basicTypes[newType.RefName] = newType.GoType
		}
		// see if this new type has sub objects or arrays we need to build
		d.processProperties(newType, protoType.Properties, protoType.Items)
		d.Types = append(d.Types, newType)

		//d.Types = append(d.Types, newType)

	}
	// second pass, over our converted ProtoType -> Type fill out arrays and objects
	for _, t := range d.Types {
		for _, props := range t.Properties {
			props.ResolveRefs(d.basicTypes)
		}
		/*
			if t.UnderlyingType != "object" && t.UnderlyingType != "array" {
				//fmt.Printf("Type: %s RefName: %s UnderlyingType: %s\n", t.Name, t.RefName, t.UnderlyingType)
				continue
			}

			if t.UnderlyingType == "object" {
				t.createObjectType(d.basicTypes)
			}*/
		//d.writeTypes()
		//os.Exit(1)
		//if t.UnderlyingType == "array" {
		//	t.createArrayType(d.basicTypes)
		//}

	}
	return nil
}

// process the new types properties, we may have sub-objects and arrays we need to create
func (d *Domain) processProperties(newType *Type, protoProps []*ProtoProperty, protoItems *ProtoItem) {
	// add the objects properties to a new TypeProperties container.
	if newType.UnderlyingType == "object" {
		d.processObjectProperties(newType, protoProps)
	}
	// If a top level type is an array it doesn't have properties, just Items (which may have properties)
	if newType.UnderlyingType == "array" {
		d.processArrayProperties(newType, protoItems)
	}
}

// We have an object type that has properties, this could be nested
// Create a new TypeProperties to add to this newType.
func (d *Domain) processObjectProperties(newType *Type, protoProps []*ProtoProperty) {
	for _, protoProp := range protoProps {
		newProps := NewTypeProperties(newType, protoProp)
		newProps.SetGoType()
		// HANDLE ARRAYS HERE???

		newType.Properties = append(newType.Properties, newProps)
		// this object has a sub object with it's own properties
		// create a new sub type
		if protoProp.Properties != nil && len(protoProp.Properties) > 0 {
			subType := NewSubType(d.Domain, newProps, protoProp)
			d.Types = append(d.Types, subType)
			d.processObjectProperties(subType, protoProp.Properties)
			newProps.Ref = subType.Name // update the reference to this new subtype
		}

	}

}

// arrays can be one of three things
// An item with a single ref
// An item with a single type (integer/string)
// An item with an embedded object
func (d *Domain) processArrayProperties(newType *Type, protoItems *ProtoItem) {
	if protoItems.Ref != "" {
		d.basicTypes[d.Domain+"."+newType.Name] = "[]*Chrome" + d.Domain + protoItems.Ref
		return
	}

	if protoItems.Type != "" {
		d.basicTypes[d.Domain+"."+newType.Name] = "[]" + d.Domain + getGoType(protoItems.Type)
		return
	}

	if protoItems.Properties != nil && len(protoItems.Properties) > 0 {

	}

}

func (d *Domain) writeTypes() {
	wr := os.Stdout

	err := templates.ExecuteTemplate(wr, "type_template.txt", d)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}
}

func recurseSubTypes(parent *TypeProperties, items []*ProtoItem) {

}
