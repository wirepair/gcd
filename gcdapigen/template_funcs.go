package main

import (
	"strings"
)

func modifyReserved(input string) string {
	switch input {
	case "type":
		return "theType"
	case "range":
		return "theRange"
	case "interface":
		return "theInterface"
	case "for":
		return "theFor"
	}
	return input
}

func nullType(input string) string {
	if strings.Contains(input, "[]") {
		return "nil"
	}
	//fmt.Printf("INPUT: %s\n", input)
	switch input {
	case "int":
		return "0"
	case "float64":
		return "0"
	case "string":
		return "\"\""
	case "bool":
		return "false"
	case "interface{}":
		return "nil"
	}
	return "nil"
}
