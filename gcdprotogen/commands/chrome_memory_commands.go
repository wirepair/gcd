// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Memory commands.
// API Version: 1.1

package gcd


import (
	
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Memory() *ChromeMemory {
	if c.memory == nil {
		c.memory = newChromeMemory(c)
	}
	return c.memory
}


type ChromeMemory struct {
	target *ChromeTarget
}

func newChromeMemory(target *ChromeTarget) *ChromeMemory {
	c := &ChromeMemory{target: target}
	return c
}

// start non parameterized commands
// end non parameterized commands

// start parameterized commands with no special return types


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// getDOMCounters - 
// Returns - 
func (c *ChromeMemory) GetDOMCounters() (float64, float64, float64, error) {	
	var documents float64 
	var nodes float64 
	var jsEventListeners float64 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Memory.getDOMCounters"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return documents, nodes, jsEventListeners, &ChromeRequestErr{Resp: cerr}
		}
		return documents, nodes, jsEventListeners, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		documents = results["documents"].(float64)
		nodes = results["nodes"].(float64)
		jsEventListeners = results["jsEventListeners"].(float64)
	}
	return documents, nodes, jsEventListeners, nil
}


// end commands with no parameters but special return types

