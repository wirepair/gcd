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
func (c *ChromeCanvas) DropTraceLog(traceLogId *types.ChromeCanvasTraceLogId, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["traceLogId"] = traceLogId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.dropTraceLog", Params: paramRequest})
}

// stopCapturing - 
// traceLogId - 
func (c *ChromeCanvas) StopCapturing(traceLogId *types.ChromeCanvasTraceLogId, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["traceLogId"] = traceLogId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.stopCapturing", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// hasUninstrumentedCanvases - Checks if there is any uninstrumented canvas in the inspected page.
// Returns - 
func (c *ChromeCanvas) HasUninstrumentedCanvases() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.hasUninstrumentedCanvases"})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Result bool 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}


// end commands with no parameters but special return types


// start commands with parameters and special return types

// captureFrame - Starts (or continues) a canvas frame capturing which will be stopped automatically after the next frame is prepared.
// Returns - 
// Identifier of the trace log containing captured canvas calls.
func (c *ChromeCanvas) CaptureFrame(frameId *types.ChromePageFrameId, ) (*types.ChromeCanvasTraceLogId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.captureFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			TraceLogId *types.ChromeCanvasTraceLogId 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.TraceLogId, nil
}

// startCapturing - Starts (or continues) consecutive canvas frames capturing. The capturing is stopped by the corresponding stopCapturing command.
// Returns - 
// Identifier of the trace log containing captured canvas calls.
func (c *ChromeCanvas) StartCapturing(frameId *types.ChromePageFrameId, ) (*types.ChromeCanvasTraceLogId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.startCapturing", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			TraceLogId *types.ChromeCanvasTraceLogId 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.TraceLogId, nil
}

// getTraceLog - 
// Returns - 
func (c *ChromeCanvas) GetTraceLog(traceLogId *types.ChromeCanvasTraceLogId, startOffset int, maxLength int, ) (*types.ChromeCanvasTraceLog, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["traceLogId"] = traceLogId
	paramRequest["startOffset"] = startOffset
	paramRequest["maxLength"] = maxLength
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.getTraceLog", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			TraceLog *types.ChromeCanvasTraceLog 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.TraceLog, nil
}

// replayTraceLog - 
// Returns - 
// Replay time (in milliseconds).
func (c *ChromeCanvas) ReplayTraceLog(traceLogId *types.ChromeCanvasTraceLogId, stepNo int, ) (*types.ChromeCanvasResourceState, float64, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["traceLogId"] = traceLogId
	paramRequest["stepNo"] = stepNo
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.replayTraceLog", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			ResourceState *types.ChromeCanvasResourceState 
			ReplayTime float64 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, 0, &ChromeRequestErr{Resp: cerr}
		}
		return nil, 0, err
	}

	return chromeData.Result.ResourceState, chromeData.Result.ReplayTime, nil
}

// getResourceState - 
// Returns - 
func (c *ChromeCanvas) GetResourceState(traceLogId *types.ChromeCanvasTraceLogId, resourceId *types.ChromeCanvasResourceId, ) (*types.ChromeCanvasResourceState, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["traceLogId"] = traceLogId
	paramRequest["resourceId"] = resourceId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.getResourceState", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			ResourceState *types.ChromeCanvasResourceState 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.ResourceState, nil
}

// evaluateTraceLogCallArgument - Evaluates a given trace call argument or its result.
// Returns - 
// Object wrapper for the evaluation result.
// State of the <code>Resource</code> object.
func (c *ChromeCanvas) EvaluateTraceLogCallArgument(traceLogId *types.ChromeCanvasTraceLogId, callIndex int, argumentIndex int, objectGroup string, ) (*types.ChromeRuntimeRemoteObject, *types.ChromeCanvasResourceState, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["traceLogId"] = traceLogId
	paramRequest["callIndex"] = callIndex
	paramRequest["argumentIndex"] = argumentIndex
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Canvas.evaluateTraceLogCallArgument", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Result *types.ChromeRuntimeRemoteObject 
			ResourceState *types.ChromeCanvasResourceState 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ResourceState, nil
}


// end commands with parameters and special return types

