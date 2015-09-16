// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains ScreenOrientation functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

type ScreenOrientation struct {
	target gcdmessage.ChromeTargeter
}

func NewScreenOrientation(target gcdmessage.ChromeTargeter) *ScreenOrientation {
	c := &ScreenOrientation{target: target}
	return c
}

// SetScreenOrientationOverride - Overrides the Screen Orientation.
// angle - Orientation angle
// type - Orientation type enum values: portraitPrimary, portraitSecondary, landscapePrimary, landscapeSecondary
func (c *ScreenOrientation) SetScreenOrientationOverride(angle int, theType string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["angle"] = angle
	paramRequest["type"] = theType
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ScreenOrientation.setScreenOrientationOverride", Params: paramRequest})
}

// Clears the overridden Screen Orientation.
func (c *ScreenOrientation) ClearScreenOrientationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ScreenOrientation.clearScreenOrientationOverride"})
}
