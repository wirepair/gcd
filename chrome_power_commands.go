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

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
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

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.Result, nil
}
