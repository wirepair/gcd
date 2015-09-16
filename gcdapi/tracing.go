// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Tracing functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
type TracingTracingCompleteEvent struct {
	Stream string `json:"stream,omitempty"` // A handle of the stream that holds resulting trace data.
}

//
type TracingBufferUsageEvent struct {
	PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
	Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
}

type Tracing struct {
	target gcdmessage.ChromeTargeter
}

func NewTracing(target gcdmessage.ChromeTargeter) *Tracing {
	c := &Tracing{target: target}
	return c
}

// Start - Start trace events collection.
// categories - Category/tag filter
// options - Tracing options
// bufferUsageReportingInterval - If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
// transferMode - Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>).
func (c *Tracing) Start(categories string, options string, bufferUsageReportingInterval float64, transferMode string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["categories"] = categories
	paramRequest["options"] = options
	paramRequest["bufferUsageReportingInterval"] = bufferUsageReportingInterval
	paramRequest["transferMode"] = transferMode
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.start", Params: paramRequest})
}

// Stop trace events collection.
func (c *Tracing) End() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.end"})
}

// GetCategories - Gets supported tracing categories.
// Returns -  categories - A list of supported tracing categories.
func (c *Tracing) GetCategories() ([]string, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.getCategories"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Categories []string
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

	return chromeData.Result.Categories, nil
}

// RequestMemoryDump - Request a global memory dump.
// Returns -  dumpGuid - GUID of the resulting global memory dump. success - True iff the global memory dump succeeded.
func (c *Tracing) RequestMemoryDump() (string, bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.requestMemoryDump"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			DumpGuid string
			Success  bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.DumpGuid, chromeData.Result.Success, nil
}
