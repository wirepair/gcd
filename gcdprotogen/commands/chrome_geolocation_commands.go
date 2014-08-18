// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Geolocation commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Geolocation() *ChromeGeolocation {
	if c.geolocation == nil {
		c.geolocation = newChromeGeolocation(c)
	}
	return c.geolocation
}


type ChromeGeolocation struct {
	target *ChromeTarget
}

func newChromeGeolocation(target *ChromeTarget) *ChromeGeolocation {
	c := &ChromeGeolocation{target: target}
	return c
}

// start non parameterized commands 
// Clears the overriden Geolocation Position and Error.
func (c *ChromeGeolocation) ClearGeolocationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Geolocation.clearGeolocationOverride"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setGeolocationOverride - Overrides the Geolocation Position or Error.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *ChromeGeolocation) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Geolocation.setGeolocationOverride", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

