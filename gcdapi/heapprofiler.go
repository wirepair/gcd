// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeapProfiler functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

//
type HeapProfilerAddHeapSnapshotChunkEvent struct {
	Chunk string `json:"chunk"` //
}

//
type HeapProfilerReportHeapSnapshotProgressEvent struct {
	Done     int  `json:"done"`               //
	Total    int  `json:"total"`              //
	Finished bool `json:"finished,omitempty"` //
}

// If heap objects tracking has been started then backend regulary sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type HeapProfilerLastSeenObjectIdEvent struct {
	LastSeenObjectId int     `json:"lastSeenObjectId"` //
	Timestamp        float64 `json:"timestamp"`        //
}

// If heap objects tracking has been started then backend may send update for one or more fragments
type HeapProfilerHeapStatsUpdateEvent struct {
	StatsUpdate int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
}

type HeapProfiler struct {
	target gcdmessage.ChromeTargeter
}

func NewHeapProfiler(target gcdmessage.ChromeTargeter) *HeapProfiler {
	c := &HeapProfiler{target: target}
	return c
}

//
func (c *HeapProfiler) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.enable"})
}

//
func (c *HeapProfiler) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.disable"})
}

// StartTrackingHeapObjects -
// trackAllocations -
func (c *HeapProfiler) StartTrackingHeapObjects(trackAllocations bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["trackAllocations"] = trackAllocations
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: paramRequest})
}

// StopTrackingHeapObjects -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
func (c *HeapProfiler) StopTrackingHeapObjects(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: paramRequest})
}

// TakeHeapSnapshot -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (c *HeapProfiler) TakeHeapSnapshot(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["reportProgress"] = reportProgress
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.takeHeapSnapshot", Params: paramRequest})
}

//
func (c *HeapProfiler) CollectGarbage() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.collectGarbage"})
}

// GetObjectByHeapObjectId -
// objectId -
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectId(objectId string, objectGroup string) (*RuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["objectId"] = objectId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getObjectByHeapObjectId"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result *RuntimeRemoteObject
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Result, nil
}

// AddInspectedHeapObject - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// heapObjectId - Heap snapshot object id to be accessible by means of $x command line API.
func (c *HeapProfiler) AddInspectedHeapObject(heapObjectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["heapObjectId"] = heapObjectId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.addInspectedHeapObject", Params: paramRequest})
}

// GetHeapObjectId -
// objectId - Identifier of the object to get heap object id for.
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectId(objectId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getHeapObjectId"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			HeapSnapshotObjectId string
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.HeapSnapshotObjectId, nil
}
