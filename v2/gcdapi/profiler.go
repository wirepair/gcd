// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Profiler functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Profile node. Holds callsite information, execution statistics and child nodes.
type ProfilerProfileNode struct {
	Id            int                         `json:"id"`                      // Unique id of the node.
	CallFrame     *RuntimeCallFrame           `json:"callFrame"`               // Function location.
	HitCount      int                         `json:"hitCount,omitempty"`      // Number of samples where this node was on top of the call stack.
	Children      []int                       `json:"children,omitempty"`      // Child node ids.
	DeoptReason   string                      `json:"deoptReason,omitempty"`   // The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	PositionTicks []*ProfilerPositionTickInfo `json:"positionTicks,omitempty"` // An array of source position ticks.
}

// Profile.
type ProfilerProfile struct {
	Nodes      []*ProfilerProfileNode `json:"nodes"`                // The list of profile nodes. First item is the root node.
	StartTime  float64                `json:"startTime"`            // Profiling start timestamp in microseconds.
	EndTime    float64                `json:"endTime"`              // Profiling end timestamp in microseconds.
	Samples    []int                  `json:"samples,omitempty"`    // Ids of samples top nodes.
	TimeDeltas []int                  `json:"timeDeltas,omitempty"` // Time intervals between adjacent samples in microseconds. The first delta is relative to the profile startTime.
}

// Specifies a number of samples attributed to a certain source position.
type ProfilerPositionTickInfo struct {
	Line  int `json:"line"`  // Source line number (1-based).
	Ticks int `json:"ticks"` // Number of samples attributed to the source line.
}

// Coverage data for a source range.
type ProfilerCoverageRange struct {
	StartOffset int `json:"startOffset"` // JavaScript script source offset for the range start.
	EndOffset   int `json:"endOffset"`   // JavaScript script source offset for the range end.
	Count       int `json:"count"`       // Collected execution count of the source range.
}

// Coverage data for a JavaScript function.
type ProfilerFunctionCoverage struct {
	FunctionName    string                   `json:"functionName"`    // JavaScript function name.
	Ranges          []*ProfilerCoverageRange `json:"ranges"`          // Source ranges inside the function with coverage data.
	IsBlockCoverage bool                     `json:"isBlockCoverage"` // Whether coverage data for this function has block granularity.
}

// Coverage data for a JavaScript script.
type ProfilerScriptCoverage struct {
	ScriptId  string                      `json:"scriptId"`  // JavaScript script id.
	Url       string                      `json:"url"`       // JavaScript script name or url.
	Functions []*ProfilerFunctionCoverage `json:"functions"` // Functions contained in the script that has coverage data.
}

type ProfilerConsoleProfileFinishedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id       string            `json:"id"`              //
		Location *DebuggerLocation `json:"location"`        // Location of console.profileEnd().
		Profile  *ProfilerProfile  `json:"profile"`         //
		Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
	} `json:"Params,omitempty"`
}

// Sent when new profile recording is started using console.profile() call.
type ProfilerConsoleProfileStartedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id       string            `json:"id"`              //
		Location *DebuggerLocation `json:"location"`        // Location of console.profile().
		Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
	} `json:"Params,omitempty"`
}

// Reports coverage delta since the last poll (either from an event like this, or from `takePreciseCoverage` for the current isolate. May only be sent if precise code coverage has been started. This event can be trigged by the embedder to, for example, trigger collection of coverage data immediately at a certain point in time.
type ProfilerPreciseCoverageDeltaUpdateEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64                   `json:"timestamp"` // Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
		Occasion  string                    `json:"occasion"`  // Identifier for distinguishing coverage events.
		Result    []*ProfilerScriptCoverage `json:"result"`    // Coverage data for the current isolate.
	} `json:"Params,omitempty"`
}

type Profiler struct {
	target gcdmessage.ChromeTargeter
}

func NewProfiler(target gcdmessage.ChromeTargeter) *Profiler {
	c := &Profiler{target: target}
	return c
}

func (c *Profiler) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.disable"})
}

func (c *Profiler) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.enable"})
}

// GetBestEffortCoverage - Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
// Returns -  result - Coverage data for the current isolate.
func (c *Profiler) GetBestEffortCoverage(ctx context.Context) ([]*ProfilerScriptCoverage, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.getBestEffortCoverage"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result []*ProfilerScriptCoverage
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

type ProfilerSetSamplingIntervalParams struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// SetSamplingIntervalWithParams - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (c *Profiler) SetSamplingIntervalWithParams(ctx context.Context, v *ProfilerSetSamplingIntervalParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.setSamplingInterval", Params: v})
}

// SetSamplingInterval - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
// interval - New sampling interval in microseconds.
func (c *Profiler) SetSamplingInterval(ctx context.Context, interval int) (*gcdmessage.ChromeResponse, error) {
	var v ProfilerSetSamplingIntervalParams
	v.Interval = interval
	return c.SetSamplingIntervalWithParams(ctx, &v)
}

func (c *Profiler) Start(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.start"})
}

type ProfilerStartPreciseCoverageParams struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount bool `json:"callCount,omitempty"`
	// Collect block-based coverage.
	Detailed bool `json:"detailed,omitempty"`
	// Allow the backend to send updates on its own initiative
	AllowTriggeredUpdates bool `json:"allowTriggeredUpdates,omitempty"`
}

// StartPreciseCoverageWithParams - Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
// Returns -  timestamp - Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
func (c *Profiler) StartPreciseCoverageWithParams(ctx context.Context, v *ProfilerStartPreciseCoverageParams) (float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.startPreciseCoverage", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Timestamp float64
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	if chromeData.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Timestamp, nil
}

// StartPreciseCoverage - Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
// callCount - Collect accurate call counts beyond simple 'covered' or 'not covered'.
// detailed - Collect block-based coverage.
// allowTriggeredUpdates - Allow the backend to send updates on its own initiative
// Returns -  timestamp - Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
func (c *Profiler) StartPreciseCoverage(ctx context.Context, callCount bool, detailed bool, allowTriggeredUpdates bool) (float64, error) {
	var v ProfilerStartPreciseCoverageParams
	v.CallCount = callCount
	v.Detailed = detailed
	v.AllowTriggeredUpdates = allowTriggeredUpdates
	return c.StartPreciseCoverageWithParams(ctx, &v)
}

// Stop -
// Returns -  profile - Recorded profile.
func (c *Profiler) Stop(ctx context.Context) (*ProfilerProfile, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.stop"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Profile *ProfilerProfile
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Profile, nil
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (c *Profiler) StopPreciseCoverage(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.stopPreciseCoverage"})
}

// TakePreciseCoverage - Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
// Returns -  result - Coverage data for the current isolate. timestamp - Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
func (c *Profiler) TakePreciseCoverage(ctx context.Context) ([]*ProfilerScriptCoverage, float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.takePreciseCoverage"})
	if err != nil {
		return nil, 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result    []*ProfilerScriptCoverage
			Timestamp float64
		}
	}

	if resp == nil {
		return nil, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, 0, err
	}

	if chromeData.Error != nil {
		return nil, 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, chromeData.Result.Timestamp, nil
}
