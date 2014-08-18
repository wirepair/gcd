// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DeviceOrientation commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) DeviceOrientation() *ChromeDeviceOrientation {
	if c.deviceorientation == nil {
		c.deviceorientation = newChromeDeviceOrientation(c)
	}
	return c.deviceorientation
}


type ChromeDeviceOrientation struct {
	target *ChromeTarget
}

func newChromeDeviceOrientation(target *ChromeTarget) *ChromeDeviceOrientation {
	c := &ChromeDeviceOrientation{target: target}
	return c
}

// start non parameterized commands 
// Clears the overridden Device Orientation.
func (c *ChromeDeviceOrientation) ClearDeviceOrientationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DeviceOrientation.clearDeviceOrientationOverride"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *ChromeDeviceOrientation) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["alpha"] = alpha
	paramRequest["beta"] = beta
	paramRequest["gamma"] = gamma
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DeviceOrientation.setDeviceOrientationOverride", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

