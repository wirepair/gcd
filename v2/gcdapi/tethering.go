// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Tethering functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Informs that port was successfully bound and got a specified connection id.
type TetheringAcceptedEvent struct {
	Method string `json:"method"`
	Params struct {
		Port         int    `json:"port"`         // Port number that was successfully bound.
		ConnectionId string `json:"connectionId"` // Connection id to be used.
	} `json:"Params,omitempty"`
}

type Tethering struct {
	target gcdmessage.ChromeTargeter
}

func NewTethering(target gcdmessage.ChromeTargeter) *Tethering {
	c := &Tethering{target: target}
	return c
}

type TetheringBindParams struct {
	// Port number to bind.
	Port int `json:"port"`
}

// BindWithParams - Request browser port binding.
func (c *Tethering) BindWithParams(ctx context.Context, v *TetheringBindParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tethering.bind", Params: v})
}

// Bind - Request browser port binding.
// port - Port number to bind.
func (c *Tethering) Bind(ctx context.Context, port int) (*gcdmessage.ChromeResponse, error) {
	var v TetheringBindParams
	v.Port = port
	return c.BindWithParams(ctx, &v)
}

type TetheringUnbindParams struct {
	// Port number to unbind.
	Port int `json:"port"`
}

// UnbindWithParams - Request browser port unbinding.
func (c *Tethering) UnbindWithParams(ctx context.Context, v *TetheringUnbindParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tethering.unbind", Params: v})
}

// Unbind - Request browser port unbinding.
// port - Port number to unbind.
func (c *Tethering) Unbind(ctx context.Context, port int) (*gcdmessage.ChromeResponse, error) {
	var v TetheringUnbindParams
	v.Port = port
	return c.UnbindWithParams(ctx, &v)
}
