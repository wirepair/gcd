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

// Start power events collection.
func (c *ChromePower) Start() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.start"})
}

// Stop power events collection.
func (c *ChromePower) End() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.end"})
}

// canProfilePower - Tells whether power profiling is supported.
// Returns -
// True if power profiling is supported.
func (c *ChromePower) CanProfilePower() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.canProfilePower"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}

// getAccuracyLevel - Describes the accuracy level of the data provider.
// Returns -
func (c *ChromePower) GetAccuracyLevel() (string, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Power.getAccuracyLevel"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result string
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

	return chromeData.Result.Result, nil
}
