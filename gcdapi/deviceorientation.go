// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DeviceOrientation functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

type DeviceOrientation struct {
	target gcdmessage.ChromeTargeter
}

func NewDeviceOrientation(target gcdmessage.ChromeTargeter) *DeviceOrientation {
	c := &DeviceOrientation{target: target}
	return c
}

// SetDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *DeviceOrientation) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["alpha"] = alpha
	paramRequest["beta"] = beta
	paramRequest["gamma"] = gamma
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceOrientation.setDeviceOrientationOverride", Params: paramRequest})
}

// Clears the overridden Device Orientation.
func (c *DeviceOrientation) ClearDeviceOrientationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceOrientation.clearDeviceOrientationOverride"})
}
