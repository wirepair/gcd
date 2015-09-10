package main

func getGoType(protocolType string) string {
	switch protocolType {
	case "enum":
		return "string"
	case "integer":
		return "int"
	case "string":
		return "string"
	case "number":
		return "float64"
	case "object":
		return "object"
	case "array":
		return "array"
	}
	return ""
}
