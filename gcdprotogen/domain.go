package main

import (
	"fmt"
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

func (d *Domain) PopulateTypes(types []*ProtoType) error {
	// do first pass to get all underlying type information
	for _, t := range types {
		fmt.Printf("Populating type: %s\n", t.Id)
		newType := NewType(d.Domain, t)
		newType.SetUnderlyingType()

		// we don't want to create single typed values if we can avoid it.
		if newType.UnderlyingType != "object" && newType.UnderlyingType != "array" {
			d.basicTypes[newType.RefName] = newType.UnderlyingType
		}

		d.Types = append(d.Types, newType)

	}
	// second pass, fill out arrays and objects
	for _, t := range d.Types {
		if t.UnderlyingType != "object" && t.UnderlyingType != "array" {
			//fmt.Printf("Type: %s RefName: %s UnderlyingType: %s\n", t.Name, t.RefName, t.UnderlyingType)
			continue
		}

		if t.UnderlyingType == "object" {
			t.createObjectType(d.basicTypes)
		}

		if t.UnderlyingType == "array" {
			t.createArrayType(d.basicTypes)
		}

	}
	return nil
}
