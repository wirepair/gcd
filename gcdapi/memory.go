// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Memory functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

type Memory struct {
	target gcdmessage.ChromeTargeter
}

func NewMemory(target gcdmessage.ChromeTargeter) *Memory {
	c := &Memory{target: target}
	return c
}

// GetDOMCounters -
// Returns -  documents -  nodes -  jsEventListeners -
func (c *Memory) GetDOMCounters() (int, int, int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.getDOMCounters"})
	if err != nil {
		return 0, 0, 0, err
	}

	var chromeData struct {
		Result struct {
			Documents        int
			Nodes            int
			JsEventListeners int
		}
	}

	if resp == nil {
		return 0, 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, 0, err
	}

	return chromeData.Result.Documents, chromeData.Result.Nodes, chromeData.Result.JsEventListeners, nil
}
