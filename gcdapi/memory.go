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

// SetPressureNotificationsSuppressed - Enable/disable suppressing memory pressure notifications in all processes.
// suppressed - If true, memory pressure notifications will be suppressed.
func (c *Memory) SetPressureNotificationsSuppressed(suppressed bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["suppressed"] = suppressed
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.setPressureNotificationsSuppressed", Params: paramRequest})
}

// SimulatePressureNotification - Simulate a memory pressure notification in all processes.
// level - Memory pressure level of the notification. enum values: moderate, critical
func (c *Memory) SimulatePressureNotification(level string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["level"] = level
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.simulatePressureNotification", Params: paramRequest})
}
