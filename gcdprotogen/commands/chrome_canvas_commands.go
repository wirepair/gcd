// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Canvas commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Canvas() *ChromeCanvas {
	if c.canvas == nil {
		c.canvas = newChromeCanvas(c)
	}
	return c.canvas
}


type ChromeCanvas struct {
	target *ChromeTarget
}

func newChromeCanvas(target *ChromeTarget) *ChromeCanvas {
	c := &ChromeCanvas{target: target}
	return c
}

// start non parameterized commands 
// Enables Canvas inspection.
func (c *ChromeCanvas) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.enable"})
}
 
// Disables Canvas inspection.
func (c *ChromeCanvas) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// dropTraceLog - 
// traceLogId - 
func (c *ChromeCanvas) DropTraceLog(traceLogId *types.ChromeCanvasTraceLogId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["traceLogId"] = traceLogId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.dropTraceLog", Params: paramRequest})
}

// stopCapturing - 
// traceLogId - 
func (c *ChromeCanvas) StopCapturing(traceLogId *types.ChromeCanvasTraceLogId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["traceLogId"] = traceLogId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.stopCapturing", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// hasUninstrumentedCanvases - Checks if there is any uninstrumented canvas in the inspected page.
// Returns - 
func (c *ChromeCanvas) HasUninstrumentedCanvases() (bool, error) {	
	var result bool 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.hasUninstrumentedCanvases"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return result, &ChromeRequestErr{Resp: cerr}
		}
		return result, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		result = results["result"].(bool)
	}
	return result, nil
}


// end commands with no parameters but special return types

