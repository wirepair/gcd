// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Tracing commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Tracing() *ChromeTracing {
	if c.tracing == nil {
		c.tracing = newChromeTracing(c)
	}
	return c.tracing
}

type ChromeTracing struct {
	target *ChromeTarget
}

func newChromeTracing(target *ChromeTarget) *ChromeTracing {
	c := &ChromeTracing{target: target}
	return c
}

// Stop trace events collection.
func (c *ChromeTracing) End() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.end"})
}

// start - Start trace events collection.
// categories - Category/tag filter
// options - Tracing options
// bufferUsageReportingInterval - If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
// transferMode - Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>).
func (c *ChromeTracing) Start(categories string, options string, bufferUsageReportingInterval float64, transferMode string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["categories"] = categories
	paramRequest["options"] = options
	paramRequest["bufferUsageReportingInterval"] = bufferUsageReportingInterval
	paramRequest["transferMode"] = transferMode
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.start", Params: paramRequest})
}

// getCategories - Gets supported tracing categories.
// Returns -
// A list of supported tracing categories.
func (c *ChromeTracing) GetCategories() ([]string, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.getCategories"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Categories []string
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

	return chromeData.Result.Categories, nil
}

// requestMemoryDump - Request a global memory dump.
// Returns -
// GUID of the resulting global memory dump.
// True iff the global memory dump succeeded.
func (c *ChromeTracing) RequestMemoryDump() (string, bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.requestMemoryDump"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			DumpGuid string
			Success  bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.DumpGuid, chromeData.Result.Success, nil
}
