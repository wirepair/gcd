// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains EventBreakpoints functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

type EventBreakpoints struct {
	target gcdmessage.ChromeTargeter
}

func NewEventBreakpoints(target gcdmessage.ChromeTargeter) *EventBreakpoints {
	c := &EventBreakpoints{target: target}
	return c
}

type EventBreakpointsSetInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// SetInstrumentationBreakpointWithParams - Sets breakpoint on particular native event.
func (c *EventBreakpoints) SetInstrumentationBreakpointWithParams(ctx context.Context, v *EventBreakpointsSetInstrumentationBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "EventBreakpoints.setInstrumentationBreakpoint", Params: v})
}

// SetInstrumentationBreakpoint - Sets breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *EventBreakpoints) SetInstrumentationBreakpoint(ctx context.Context, eventName string) (*gcdmessage.ChromeResponse, error) {
	var v EventBreakpointsSetInstrumentationBreakpointParams
	v.EventName = eventName
	return c.SetInstrumentationBreakpointWithParams(ctx, &v)
}

type EventBreakpointsRemoveInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// RemoveInstrumentationBreakpointWithParams - Removes breakpoint on particular native event.
func (c *EventBreakpoints) RemoveInstrumentationBreakpointWithParams(ctx context.Context, v *EventBreakpointsRemoveInstrumentationBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "EventBreakpoints.removeInstrumentationBreakpoint", Params: v})
}

// RemoveInstrumentationBreakpoint - Removes breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *EventBreakpoints) RemoveInstrumentationBreakpoint(ctx context.Context, eventName string) (*gcdmessage.ChromeResponse, error) {
	var v EventBreakpointsRemoveInstrumentationBreakpointParams
	v.EventName = eventName
	return c.RemoveInstrumentationBreakpointWithParams(ctx, &v)
}

// Removes all breakpoints
func (c *EventBreakpoints) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "EventBreakpoints.disable"})
}
