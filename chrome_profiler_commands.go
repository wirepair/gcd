// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Profiler commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Profiler() *ChromeProfiler {
	if c.profiler == nil {
		c.profiler = newChromeProfiler(c)
	}
	return c.profiler
}

type ChromeProfiler struct {
	target *ChromeTarget
}

func newChromeProfiler(target *ChromeTarget) *ChromeProfiler {
	c := &ChromeProfiler{target: target}
	return c
}

//
func (c *ChromeProfiler) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.enable"})
}

//
func (c *ChromeProfiler) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.disable"})
}

//
func (c *ChromeProfiler) Start() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.start"})
}

// setSamplingInterval - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
// interval - New sampling interval in microseconds.
func (c *ChromeProfiler) SetSamplingInterval(interval int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["interval"] = interval
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.setSamplingInterval", Params: paramRequest})
}

// stop -
// Returns -
// Recorded profile.
func (c *ChromeProfiler) Stop() (*types.ChromeProfilerCPUProfile, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Profiler.stop"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Profile *types.ChromeProfilerCPUProfile
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Profile, nil
}
