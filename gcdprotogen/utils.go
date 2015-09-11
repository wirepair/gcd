package main

func getGoType(protocolType string) string {
	switch protocolType {
	case "any", "string", "enum":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "object":
		return "map[string]interface{}" // default
	case "array":
		return "array"
	}
	return ""
}
