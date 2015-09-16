// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Rendering functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

type Rendering struct {
	target gcdmessage.ChromeTargeter
}

func NewRendering(target gcdmessage.ChromeTargeter) *Rendering {
	c := &Rendering{target: target}
	return c
}

// SetShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *Rendering) SetShowPaintRects(result bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["result"] = result
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Rendering.setShowPaintRects", Params: paramRequest})
}

// SetShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *Rendering) SetShowDebugBorders(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Rendering.setShowDebugBorders", Params: paramRequest})
}

// SetShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *Rendering) SetShowFPSCounter(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Rendering.setShowFPSCounter", Params: paramRequest})
}

// SetShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *Rendering) SetShowScrollBottleneckRects(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Rendering.setShowScrollBottleneckRects", Params: paramRequest})
}
