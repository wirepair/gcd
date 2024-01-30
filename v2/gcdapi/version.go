package gcdapi

import "github.com/json-iterator/go"

// define json here for all Domains
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Chrome Channel information
const CHROME_CHANNEL = "unknown" 
// Chrome Version information
const CHROME_VERSION = "unknown"
// Protocol Major version
const DEVTOOLS_MAJOR_VERSION = "1"
// Protocol Minor version
const DEVTOOLS_MINOR_VERSION = "3"