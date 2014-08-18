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
func (c *ChromeTracing) GetCategories() ([]types.string, error) {	
	var categories []types.string 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Tracing.getCategories"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return categories, &ChromeRequestErr{Resp: cerr}
		}
		return categories, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		categories = results["categories"].([]types.string)
	}
	return categories, nil
}


// end commands with no parameters but special return types

