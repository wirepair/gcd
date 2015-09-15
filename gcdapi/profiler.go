// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Profiler functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// CPU Profile node. Holds callsite information, execution statistics and child nodes.
type ProfilerCPUProfileNode struct {
	FunctionName  string                      `json:"functionName"`  // Function name.
	ScriptId      string                      `json:"scriptId"`      // Script identifier.
	Url           string                      `json:"url"`           // URL.
	LineNumber    int                         `json:"lineNumber"`    // 1-based line number of the function start position.
	ColumnNumber  int                         `json:"columnNumber"`  // 1-based column number of the function start position.
	HitCount      int                         `json:"hitCount"`      // Number of samples where this node was on top of the call stack.
	CallUID       float64                     `json:"callUID"`       // Call UID.
	Children      []*ProfilerCPUProfileNode   `json:"children"`      // Child nodes.
	DeoptReason   string                      `json:"deoptReason"`   // The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	Id            int                         `json:"id"`            // Unique id of the node.
	PositionTicks []*ProfilerPositionTickInfo `json:"positionTicks"` // An array of source position ticks.
}

// Profile.
type ProfilerCPUProfile struct {
	Head       *ProfilerCPUProfileNode `json:"head"`                 //
	StartTime  float64                 `json:"startTime"`            // Profiling start time in seconds.
	EndTime    float64                 `json:"endTime"`              // Profiling end time in seconds.
	Samples    []int                   `json:"samples,omitempty"`    // Ids of samples top nodes.
	Timestamps []float64               `json:"timestamps,omitempty"` // Timestamps of the samples in microseconds.
}

// Specifies a number of samples attributed to a certain source position.
type ProfilerPositionTickInfo struct {
	Line  int `json:"line"`  // Source line number (1-based).
	Ticks int `json:"ticks"` // Number of samples attributed to the source line.
}

// Sent when new profile recodring is started using console.profile() call.
type ProfilerConsoleProfileStartedEvent struct {
	Id       string            `json:"id"`              //
	Location *DebuggerLocation `json:"location"`        // Location of console.profile().
	Title    string            `json:"title,omitempty"` // Profile title passed as argument to console.profile().
}

//
type ProfilerConsoleProfileFinishedEvent struct {
	Id       string              `json:"id"`              //
	Location *DebuggerLocation   `json:"location"`        // Location of console.profileEnd().
	Profile  *ProfilerCPUProfile `json:"profile"`         //
	Title    string              `json:"title,omitempty"` // Profile title passed as argunet to console.profile().
}

type Profiler struct {
	target gcdmessage.ChromeTargeter
}

func NewProfiler(target gcdmessage.ChromeTargeter) *Profiler {
	c := &Profiler{target: target}
	return c
}

//
func (c *Profiler) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.enable"})
}

//
func (c *Profiler) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.disable"})
}

// SetSamplingInterval - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
// interval - New sampling interval in microseconds.
func (c *Profiler) SetSamplingInterval(interval int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["interval"] = interval
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.setSamplingInterval", Params: paramRequest})
}

//
func (c *Profiler) Start() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.start"})
}

// Stop -
// Returns -  profile - Recorded profile.
func (c *Profiler) Stop() (*ProfilerCPUProfile, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.stop"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Profile *ProfilerCPUProfile
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

	return chromeData.Result.Profile, nil
}
