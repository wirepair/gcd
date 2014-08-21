// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the HeapProfiler commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
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


// startTrackingHeapObjects - 
// trackAllocations - 
func (c *ChromeHeapProfiler) StartTrackingHeapObjects(trackAllocations bool, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["trackAllocations"] = trackAllocations
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: paramRequest})
}

// stopTrackingHeapObjects - 
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
func (c *ChromeHeapProfiler) StopTrackingHeapObjects(reportProgress bool, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: paramRequest})
}

// takeHeapSnapshot - 
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (c *ChromeHeapProfiler) TakeHeapSnapshot(reportProgress bool, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.takeHeapSnapshot", Params: paramRequest})
}





// getObjectByHeapObjectId - 
// Returns - 
// Evaluation result.
func (c *ChromeHeapProfiler) GetObjectByHeapObjectId(objectId *types.ChromeHeapProfilerHeapSnapshotObjectId, objectGroup string, ) (*types.ChromeRuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["objectId"] = objectId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.getObjectByHeapObjectId", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Result *types.ChromeRuntimeRemoteObject 
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

	return chromeData.Result.Result, nil
}

// getHeapObjectId - 
// Returns - 
// Id of the heap snapshot object corresponding to the passed remote object id.
func (c *ChromeHeapProfiler) GetHeapObjectId(objectId *types.ChromeRuntimeRemoteObjectId, ) (*types.ChromeHeapProfilerHeapSnapshotObjectId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "HeapProfiler.getHeapObjectId", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			HeapSnapshotObjectId *types.ChromeHeapProfilerHeapSnapshotObjectId 
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

	return chromeData.Result.HeapSnapshotObjectId, nil
}


