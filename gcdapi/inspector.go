// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Inspector functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/gcdmessage"
)

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

// Disables inspector domain notifications.
func (c *Inspector) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Inspector.disable"})
}

// Enables inspector domain notifications.
func (c *Inspector) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Inspector.enable"})
}
