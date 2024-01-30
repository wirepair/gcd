// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeapProfiler functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
type HeapProfilerSamplingHeapProfileNode struct {
	CallFrame *RuntimeCallFrame                      `json:"callFrame"` // Function location.
	SelfSize  float64                                `json:"selfSize"`  // Allocations size in bytes for the node excluding children.
	Id        int                                    `json:"id"`        // Node id. Ids are unique across all profiles collected between startSampling and stopSampling.
	Children  []*HeapProfilerSamplingHeapProfileNode `json:"children"`  // Child nodes.
}

// A single sample from a sampling profile.
type HeapProfilerSamplingHeapProfileSample struct {
	Size    float64 `json:"size"`    // Allocation size in bytes attributed to the sample.
	NodeId  int     `json:"nodeId"`  // Id of the corresponding profile tree node.
	Ordinal float64 `json:"ordinal"` // Time-ordered sample ordinal number. It is unique across all profiles retrieved between startSampling and stopSampling.
}

// Sampling profile.
type HeapProfilerSamplingHeapProfile struct {
	Head    *HeapProfilerSamplingHeapProfileNode     `json:"head"`    //
	Samples []*HeapProfilerSamplingHeapProfileSample `json:"samples"` //
}

type HeapProfilerAddHeapSnapshotChunkEvent struct {
	Method string `json:"method"`
	Params struct {
		Chunk string `json:"chunk"` //
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend may send update for one or more fragments
type HeapProfilerHeapStatsUpdateEvent struct {
	Method string `json:"method"`
	Params struct {
		StatsUpdate []int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend regularly sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type HeapProfilerLastSeenObjectIdEvent struct {
	Method string `json:"method"`
	Params struct {
		LastSeenObjectId int     `json:"lastSeenObjectId"` //
		Timestamp        float64 `json:"timestamp"`        //
	} `json:"Params,omitempty"`
}

type HeapProfilerReportHeapSnapshotProgressEvent struct {
	Method string `json:"method"`
	Params struct {
		Done     int  `json:"done"`               //
		Total    int  `json:"total"`              //
		Finished bool `json:"finished,omitempty"` //
	} `json:"Params,omitempty"`
}

type HeapProfiler struct {
	target gcdmessage.ChromeTargeter
}

func NewHeapProfiler(target gcdmessage.ChromeTargeter) *HeapProfiler {
	c := &HeapProfiler{target: target}
	return c
}

type HeapProfilerAddInspectedHeapObjectParams struct {
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId string `json:"heapObjectId"`
}

// AddInspectedHeapObjectWithParams - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (c *HeapProfiler) AddInspectedHeapObjectWithParams(ctx context.Context, v *HeapProfilerAddInspectedHeapObjectParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.addInspectedHeapObject", Params: v})
}

// AddInspectedHeapObject - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// heapObjectId - Heap snapshot object id to be accessible by means of $x command line API.
func (c *HeapProfiler) AddInspectedHeapObject(ctx context.Context, heapObjectId string) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerAddInspectedHeapObjectParams
	v.HeapObjectId = heapObjectId
	return c.AddInspectedHeapObjectWithParams(ctx, &v)
}

func (c *HeapProfiler) CollectGarbage(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.collectGarbage"})
}

func (c *HeapProfiler) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.disable"})
}

func (c *HeapProfiler) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.enable"})
}

type HeapProfilerGetHeapObjectIdParams struct {
	// Identifier of the object to get heap object id for.
	ObjectId string `json:"objectId"`
}

// GetHeapObjectIdWithParams -
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectIdWithParams(ctx context.Context, v *HeapProfilerGetHeapObjectIdParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getHeapObjectId", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			HeapSnapshotObjectId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.HeapSnapshotObjectId, nil
}

// GetHeapObjectId -
// objectId - Identifier of the object to get heap object id for.
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectId(ctx context.Context, objectId string) (string, error) {
	var v HeapProfilerGetHeapObjectIdParams
	v.ObjectId = objectId
	return c.GetHeapObjectIdWithParams(ctx, &v)
}

type HeapProfilerGetObjectByHeapObjectIdParams struct {
	//
	ObjectId string `json:"objectId"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// GetObjectByHeapObjectIdWithParams -
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectIdWithParams(ctx context.Context, v *HeapProfilerGetObjectByHeapObjectIdParams) (*RuntimeRemoteObject, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getObjectByHeapObjectId", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result *RuntimeRemoteObject
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// GetObjectByHeapObjectId -
// objectId -
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectId(ctx context.Context, objectId string, objectGroup string) (*RuntimeRemoteObject, error) {
	var v HeapProfilerGetObjectByHeapObjectIdParams
	v.ObjectId = objectId
	v.ObjectGroup = objectGroup
	return c.GetObjectByHeapObjectIdWithParams(ctx, &v)
}

// GetSamplingProfile -
// Returns -  profile - Return the sampling profile being collected.
func (c *HeapProfiler) GetSamplingProfile(ctx context.Context) (*HeapProfilerSamplingHeapProfile, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getSamplingProfile"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Profile *HeapProfilerSamplingHeapProfile
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Profile, nil
}

type HeapProfilerStartSamplingParams struct {
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
	SamplingInterval float64 `json:"samplingInterval,omitempty"`
	// By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by major GC, which will show which functions cause large temporary memory usage or long GC pauses.
	IncludeObjectsCollectedByMajorGC bool `json:"includeObjectsCollectedByMajorGC,omitempty"`
	// By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by minor GC, which is useful when tuning a latency-sensitive application for minimal GC activity.
	IncludeObjectsCollectedByMinorGC bool `json:"includeObjectsCollectedByMinorGC,omitempty"`
}

// StartSamplingWithParams -
func (c *HeapProfiler) StartSamplingWithParams(ctx context.Context, v *HeapProfilerStartSamplingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startSampling", Params: v})
}

// StartSampling -
// samplingInterval - Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
// includeObjectsCollectedByMajorGC - By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by major GC, which will show which functions cause large temporary memory usage or long GC pauses.
// includeObjectsCollectedByMinorGC - By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by minor GC, which is useful when tuning a latency-sensitive application for minimal GC activity.
func (c *HeapProfiler) StartSampling(ctx context.Context, samplingInterval float64, includeObjectsCollectedByMajorGC bool, includeObjectsCollectedByMinorGC bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStartSamplingParams
	v.SamplingInterval = samplingInterval
	v.IncludeObjectsCollectedByMajorGC = includeObjectsCollectedByMajorGC
	v.IncludeObjectsCollectedByMinorGC = includeObjectsCollectedByMinorGC
	return c.StartSamplingWithParams(ctx, &v)
}

type HeapProfilerStartTrackingHeapObjectsParams struct {
	//
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

// StartTrackingHeapObjectsWithParams -
func (c *HeapProfiler) StartTrackingHeapObjectsWithParams(ctx context.Context, v *HeapProfilerStartTrackingHeapObjectsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: v})
}

// StartTrackingHeapObjects -
// trackAllocations -
func (c *HeapProfiler) StartTrackingHeapObjects(ctx context.Context, trackAllocations bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStartTrackingHeapObjectsParams
	v.TrackAllocations = trackAllocations
	return c.StartTrackingHeapObjectsWithParams(ctx, &v)
}

// StopSampling -
// Returns -  profile - Recorded sampling heap profile.
func (c *HeapProfiler) StopSampling(ctx context.Context) (*HeapProfilerSamplingHeapProfile, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopSampling"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Profile *HeapProfilerSamplingHeapProfile
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Profile, nil
}

type HeapProfilerStopTrackingHeapObjectsParams struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
	ReportProgress bool `json:"reportProgress,omitempty"`
	// Deprecated in favor of `exposeInternals`.
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	// If true, numerical values are included in the snapshot
	CaptureNumericValue bool `json:"captureNumericValue,omitempty"`
	// If true, exposes internals of the snapshot.
	ExposeInternals bool `json:"exposeInternals,omitempty"`
}

// StopTrackingHeapObjectsWithParams -
func (c *HeapProfiler) StopTrackingHeapObjectsWithParams(ctx context.Context, v *HeapProfilerStopTrackingHeapObjectsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: v})
}

// StopTrackingHeapObjects -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
// treatGlobalObjectsAsRoots - Deprecated in favor of `exposeInternals`.
// captureNumericValue - If true, numerical values are included in the snapshot
// exposeInternals - If true, exposes internals of the snapshot.
func (c *HeapProfiler) StopTrackingHeapObjects(ctx context.Context, reportProgress bool, treatGlobalObjectsAsRoots bool, captureNumericValue bool, exposeInternals bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStopTrackingHeapObjectsParams
	v.ReportProgress = reportProgress
	v.TreatGlobalObjectsAsRoots = treatGlobalObjectsAsRoots
	v.CaptureNumericValue = captureNumericValue
	v.ExposeInternals = exposeInternals
	return c.StopTrackingHeapObjectsWithParams(ctx, &v)
}

type HeapProfilerTakeHeapSnapshotParams struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
	// If true, a raw snapshot without artificial roots will be generated. Deprecated in favor of `exposeInternals`.
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	// If true, numerical values are included in the snapshot
	CaptureNumericValue bool `json:"captureNumericValue,omitempty"`
	// If true, exposes internals of the snapshot.
	ExposeInternals bool `json:"exposeInternals,omitempty"`
}

// TakeHeapSnapshotWithParams -
func (c *HeapProfiler) TakeHeapSnapshotWithParams(ctx context.Context, v *HeapProfilerTakeHeapSnapshotParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.takeHeapSnapshot", Params: v})
}

// TakeHeapSnapshot -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
// treatGlobalObjectsAsRoots - If true, a raw snapshot without artificial roots will be generated. Deprecated in favor of `exposeInternals`.
// captureNumericValue - If true, numerical values are included in the snapshot
// exposeInternals - If true, exposes internals of the snapshot.
func (c *HeapProfiler) TakeHeapSnapshot(ctx context.Context, reportProgress bool, treatGlobalObjectsAsRoots bool, captureNumericValue bool, exposeInternals bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerTakeHeapSnapshotParams
	v.ReportProgress = reportProgress
	v.TreatGlobalObjectsAsRoots = treatGlobalObjectsAsRoots
	v.CaptureNumericValue = captureNumericValue
	v.ExposeInternals = exposeInternals
	return c.TakeHeapSnapshotWithParams(ctx, &v)
}
