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

// start non parameterized commands 
// Stop trace events collection.
func (c *ChromeTracing) End() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.end"})
}

// end non parameterized commands

// start parameterized commands with no special return types


// end parameterized commands with no special return types


// start commands with no parameters but special return types

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
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Categories, nil
}


// end commands with no parameters but special return types


// start commands with parameters and special return types

// start - Start trace events collection.
// Returns - 
// A system-unique identifier of the tracing session that allows associating of some of the trace events with the inspected page
func (c *ChromeTracing) Start(categories string, options string, bufferUsageReportingInterval float64, ) (string, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["categories"] = categories
	paramRequest["options"] = options
	paramRequest["bufferUsageReportingInterval"] = bufferUsageReportingInterval
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.start", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			SessionId string 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.SessionId, nil
}


// end commands with parameters and special return types

