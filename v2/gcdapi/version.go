package gcdapi

import "github.com/goccy/go-json"

// define json here for all Domains
// useful so that each generated file doesn't need to know whether to import the JSON library
var jsonUnmarshal = json.Unmarshal

// Chrome Channel information
const CHROME_CHANNEL = "unknown" 
// Chrome Version information
const CHROME_VERSION = "unknown"
// Protocol Major version
const DEVTOOLS_MAJOR_VERSION = "1"
// Protocol Minor version
const DEVTOOLS_MINOR_VERSION = "3"