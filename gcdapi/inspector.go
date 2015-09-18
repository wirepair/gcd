// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Inspector functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

//
type InspectorEvaluateForTestInFrontendEvent struct {
	Method string `json:"method"`
	Params struct {
		TestCallId int    `json:"testCallId"` //
		Script     string `json:"script"`     //
	} `json:"Params,omitempty"`
}

//
type InspectorInspectEvent struct {
	Method string `json:"method"`
	Params struct {
		Object *RuntimeRemoteObject   `json:"object"` //
		Hints  map[string]interface{} `json:"hints"`  //
	} `json:"Params,omitempty"`
}

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
type InspectorDetachedEvent struct {
	Method string `json:"method"`
	Params struct {
		Reason string `json:"reason"` // The reason why connection has been terminated.
	} `json:"Params,omitempty"`
}

type Inspector struct {
	target gcdmessage.ChromeTargeter
}

func NewInspector(target gcdmessage.ChromeTargeter) *Inspector {
	c := &Inspector{target: target}
	return c
}

// Enables inspector domain notifications.
func (c *Inspector) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Inspector.enable"})
}

// Disables inspector domain notifications.
func (c *Inspector) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Inspector.disable"})
}
