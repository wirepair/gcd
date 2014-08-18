// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Power commands.
// API Version: 1.1

package gcd


import (
	
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Power() *ChromePower {
	if c.power == nil {
		c.power = newChromePower(c)
	}
	return c.power
}


type ChromePower struct {
	target *ChromeTarget
}

func newChromePower(target *ChromeTarget) *ChromePower {
	c := &ChromePower{target: target}
	return c
}

// start non parameterized commands 
// Start power events collection.
func (c *ChromePower) Start() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.start"})
}
 
// Stop power events collection.
func (c *ChromePower) End() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.end"})
}

// end non parameterized commands

// start parameterized commands with no special return types


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// canProfilePower - Tells whether power profiling is supported.
// Returns - 
// True if power profiling is supported.
func (c *ChromePower) CanProfilePower() (bool, error) {	
	var result bool 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.canProfilePower"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return result, &ChromeRequestErr{Resp: cerr}
		}
		return result, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		result = results["result"].(bool)
	}
	return result, nil
}

// getAccuracyLevel - Describes the accuracy level of the data provider.
// Returns - 
func (c *ChromePower) GetAccuracyLevel() (string, error) {	
	var result string 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.getAccuracyLevel"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return result, &ChromeRequestErr{Resp: cerr}
		}
		return result, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		result = results["result"].(string)
	}
	return result, nil
}


// end commands with no parameters but special return types

