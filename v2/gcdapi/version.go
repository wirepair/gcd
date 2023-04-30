package gcdapi

import "github.com/json-iterator/go"

// define json here for all Domains
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Chrome Channel information
const CHROME_CHANNEL = "stable" 
// Chrome Version information
const CHROME_VERSION = "112.0.5615.138"
// Protocol Major version
const DEVTOOLS_MAJOR_VERSION = "1"
// Protocol Minor version
const DEVTOOLS_MINOR_VERSION = "3"