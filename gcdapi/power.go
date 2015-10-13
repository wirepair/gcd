// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Power functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// PowerEvent item
type PowerPowerEvent struct {
	Type      string  `json:"type"`      // Power Event Type.
	Timestamp float64 `json:"timestamp"` // Power Event Time, in milliseconds.
	Value     float64 `json:"value"`     // Power Event Value.
}

//
type PowerDataAvailableEvent struct {
	Method string `json:"method"`
	Params struct {
		Value []*PowerPowerEvent `json:"value"` // List of power events.
	} `json:"Params,omitempty"`
}

type Power struct {
	target gcdmessage.ChromeTargeter
}

func NewPower(target gcdmessage.ChromeTargeter) *Power {
	c := &Power{target: target}
	return c
}

// Start power events collection.
func (c *Power) Start() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Power.start"})
}

// Stop power events collection.
func (c *Power) End() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Power.end"})
}

// CanProfilePower - Tells whether power profiling is supported.
// Returns -  result - True if power profiling is supported.
func (c *Power) CanProfilePower() (bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Power.canProfilePower"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// GetAccuracyLevel - Describes the accuracy level of the data provider.
// Returns -  result -
func (c *Power) GetAccuracyLevel() (string, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Power.getAccuracyLevel"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.Result, nil
}
