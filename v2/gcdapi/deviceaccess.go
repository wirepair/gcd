// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DeviceAccess functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Device information displayed in a user prompt to select a device.
type DeviceAccessPromptDevice struct {
	Id   string `json:"id"`   //
	Name string `json:"name"` // Display name as it appears in a device request user prompt.
}

// A device request opened a user prompt to select a device. Respond with the selectPrompt or cancelPrompt command.
type DeviceAccessDeviceRequestPromptedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id      string                      `json:"id"`      //
		Devices []*DeviceAccessPromptDevice `json:"devices"` //
	} `json:"Params,omitempty"`
}

type DeviceAccess struct {
	target gcdmessage.ChromeTargeter
}

func NewDeviceAccess(target gcdmessage.ChromeTargeter) *DeviceAccess {
	c := &DeviceAccess{target: target}
	return c
}

// Enable events in this domain.
func (c *DeviceAccess) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceAccess.enable"})
}

// Disable events in this domain.
func (c *DeviceAccess) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceAccess.disable"})
}

type DeviceAccessSelectPromptParams struct {
	//
	Id string `json:"id"`
	//
	DeviceId string `json:"deviceId"`
}

// SelectPromptWithParams - Select a device in response to a DeviceAccess.deviceRequestPrompted event.
func (c *DeviceAccess) SelectPromptWithParams(ctx context.Context, v *DeviceAccessSelectPromptParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceAccess.selectPrompt", Params: v})
}

// SelectPrompt - Select a device in response to a DeviceAccess.deviceRequestPrompted event.
// id -
// deviceId -
func (c *DeviceAccess) SelectPrompt(ctx context.Context, id string, deviceId string) (*gcdmessage.ChromeResponse, error) {
	var v DeviceAccessSelectPromptParams
	v.Id = id
	v.DeviceId = deviceId
	return c.SelectPromptWithParams(ctx, &v)
}

type DeviceAccessCancelPromptParams struct {
	//
	Id string `json:"id"`
}

// CancelPromptWithParams - Cancel a prompt in response to a DeviceAccess.deviceRequestPrompted event.
func (c *DeviceAccess) CancelPromptWithParams(ctx context.Context, v *DeviceAccessCancelPromptParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceAccess.cancelPrompt", Params: v})
}

// CancelPrompt - Cancel a prompt in response to a DeviceAccess.deviceRequestPrompted event.
// id -
func (c *DeviceAccess) CancelPrompt(ctx context.Context, id string) (*gcdmessage.ChromeResponse, error) {
	var v DeviceAccessCancelPromptParams
	v.Id = id
	return c.CancelPromptWithParams(ctx, &v)
}
