// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the HeapProfiler commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) HeapProfiler() *ChromeHeapProfiler {
	if c.heapprofiler == nil {
		c.heapprofiler = newChromeHeapProfiler(c)
	}
	return c.heapprofiler
}


type ChromeHeapProfiler struct {
	target *ChromeTarget
}

func newChromeHeapProfiler(target *ChromeTarget) *ChromeHeapProfiler {
	c := &ChromeHeapProfiler{target: target}
	return c
}

// start non parameterized commands 
// 
func (c *ChromeHeapProfiler) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.enable"})
}
 
// 
func (c *ChromeHeapProfiler) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.disable"})
}
 
// 
func (c *ChromeHeapProfiler) CollectGarbage() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.collectGarbage"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// startTrackingHeapObjects - 
// trackAllocations - 
func (c *ChromeHeapProfiler) StartTrackingHeapObjects(trackAllocations bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["trackAllocations"] = trackAllocations
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: paramRequest})
}

// stopTrackingHeapObjects - 
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
func (c *ChromeHeapProfiler) StopTrackingHeapObjects(reportProgress bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: paramRequest})
}

// takeHeapSnapshot - 
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (c *ChromeHeapProfiler) TakeHeapSnapshot(reportProgress bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.takeHeapSnapshot", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

