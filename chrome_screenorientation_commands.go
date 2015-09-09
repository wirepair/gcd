// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ScreenOrientation commands.
// API Version: 1.1

package gcd

import (
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) ScreenOrientation() *ChromeScreenOrientation {
	if c.screenorientation == nil {
		c.screenorientation = newChromeScreenOrientation(c)
	}
	return c.screenorientation
}

type ChromeScreenOrientation struct {
	target *ChromeTarget
}

func newChromeScreenOrientation(target *ChromeTarget) *ChromeScreenOrientation {
	c := &ChromeScreenOrientation{target: target}
	return c
}

// Clears the overridden Screen Orientation.
func (c *ChromeScreenOrientation) ClearScreenOrientationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ScreenOrientation.clearScreenOrientationOverride"})
}

// setScreenOrientationOverride - Overrides the Screen Orientation.
// angle - Orientation angle
// type - Orientation type
func (c *ChromeScreenOrientation) SetScreenOrientationOverride(angle int, theType *types.ChromeScreenOrientationOrientationType) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["angle"] = angle
	paramRequest["type"] = theType
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ScreenOrientation.setScreenOrientationOverride", Params: paramRequest})
}
