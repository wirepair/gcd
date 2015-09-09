// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Rendering commands.
// API Version: 1.1

package gcd

import ()

// add this API domain to ChromeTarget
func (c *ChromeTarget) Rendering() *ChromeRendering {
	if c.rendering == nil {
		c.rendering = newChromeRendering(c)
	}
	return c.rendering
}

type ChromeRendering struct {
	target *ChromeTarget
}

func newChromeRendering(target *ChromeTarget) *ChromeRendering {
	c := &ChromeRendering{target: target}
	return c
}

// setShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *ChromeRendering) SetShowPaintRects(result bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["result"] = result
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Rendering.setShowPaintRects", Params: paramRequest})
}

// setShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *ChromeRendering) SetShowDebugBorders(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Rendering.setShowDebugBorders", Params: paramRequest})
}

// setShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *ChromeRendering) SetShowFPSCounter(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Rendering.setShowFPSCounter", Params: paramRequest})
}

// setShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *ChromeRendering) SetShowScrollBottleneckRects(show bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Rendering.setShowScrollBottleneckRects", Params: paramRequest})
}
