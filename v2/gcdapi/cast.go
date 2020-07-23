// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Cast functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type CastSink struct {
	Name    string `json:"name"`              //
	Id      string `json:"id"`                //
	Session string `json:"session,omitempty"` // Text describing the current session. Present only if there is an active session on the sink.
}

// This is fired whenever the list of available sinks changes. A sink is a device or a software surface that you can cast to.
type CastSinksUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Sinks []*CastSink `json:"sinks"` //
	} `json:"Params,omitempty"`
}

// This is fired whenever the outstanding issue/error message changes. |issueMessage| is empty if there is no issue.
type CastIssueUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		IssueMessage string `json:"issueMessage"` //
	} `json:"Params,omitempty"`
}

type Cast struct {
	target gcdmessage.ChromeTargeter
}

func NewCast(target gcdmessage.ChromeTargeter) *Cast {
	c := &Cast{target: target}
	return c
}

type CastEnableParams struct {
	//
	PresentationUrl string `json:"presentationUrl,omitempty"`
}

// EnableWithParams - Starts observing for sinks that can be used for tab mirroring, and if set, sinks compatible with |presentationUrl| as well. When sinks are found, a |sinksUpdated| event is fired. Also starts observing for issue messages. When an issue is added or removed, an |issueUpdated| event is fired.
func (c *Cast) EnableWithParams(ctx context.Context, v *CastEnableParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Cast.enable", Params: v})
}

// Enable - Starts observing for sinks that can be used for tab mirroring, and if set, sinks compatible with |presentationUrl| as well. When sinks are found, a |sinksUpdated| event is fired. Also starts observing for issue messages. When an issue is added or removed, an |issueUpdated| event is fired.
// presentationUrl -
func (c *Cast) Enable(ctx context.Context, presentationUrl string) (*gcdmessage.ChromeResponse, error) {
	var v CastEnableParams
	v.PresentationUrl = presentationUrl
	return c.EnableWithParams(ctx, &v)
}

// Stops observing for sinks and issues.
func (c *Cast) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Cast.disable"})
}

type CastSetSinkToUseParams struct {
	//
	SinkName string `json:"sinkName"`
}

// SetSinkToUseWithParams - Sets a sink to be used when the web page requests the browser to choose a sink via Presentation API, Remote Playback API, or Cast SDK.
func (c *Cast) SetSinkToUseWithParams(ctx context.Context, v *CastSetSinkToUseParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Cast.setSinkToUse", Params: v})
}

// SetSinkToUse - Sets a sink to be used when the web page requests the browser to choose a sink via Presentation API, Remote Playback API, or Cast SDK.
// sinkName -
func (c *Cast) SetSinkToUse(ctx context.Context, sinkName string) (*gcdmessage.ChromeResponse, error) {
	var v CastSetSinkToUseParams
	v.SinkName = sinkName
	return c.SetSinkToUseWithParams(ctx, &v)
}

type CastStartTabMirroringParams struct {
	//
	SinkName string `json:"sinkName"`
}

// StartTabMirroringWithParams - Starts mirroring the tab to the sink.
func (c *Cast) StartTabMirroringWithParams(ctx context.Context, v *CastStartTabMirroringParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Cast.startTabMirroring", Params: v})
}

// StartTabMirroring - Starts mirroring the tab to the sink.
// sinkName -
func (c *Cast) StartTabMirroring(ctx context.Context, sinkName string) (*gcdmessage.ChromeResponse, error) {
	var v CastStartTabMirroringParams
	v.SinkName = sinkName
	return c.StartTabMirroringWithParams(ctx, &v)
}

type CastStopCastingParams struct {
	//
	SinkName string `json:"sinkName"`
}

// StopCastingWithParams - Stops the active Cast session on the sink.
func (c *Cast) StopCastingWithParams(ctx context.Context, v *CastStopCastingParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Cast.stopCasting", Params: v})
}

// StopCasting - Stops the active Cast session on the sink.
// sinkName -
func (c *Cast) StopCasting(ctx context.Context, sinkName string) (*gcdmessage.ChromeResponse, error) {
	var v CastStopCastingParams
	v.SinkName = sinkName
	return c.StopCastingWithParams(ctx, &v)
}
