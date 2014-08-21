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





// getDOMCounters - 
// Returns - 
func (c *ChromeMemory) GetDOMCounters() (float64, float64, float64, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Memory.getDOMCounters"})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Documents float64 
			Nodes float64 
			JsEventListeners float64 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return 0, 0, 0, &ChromeRequestErr{Resp: cerr}
		}
		return 0, 0, 0, err
	}

	return chromeData.Result.Documents, chromeData.Result.Nodes, chromeData.Result.JsEventListeners, nil
}




