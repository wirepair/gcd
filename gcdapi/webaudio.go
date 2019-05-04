// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains WebAudio functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Fields in AudioContext that change in real-time. These are not updated on OfflineAudioContext.
type WebAudioContextRealtimeData struct {
	CurrentTime    float64 `json:"currentTime,omitempty"`    // The current context time in second in BaseAudioContext.
	RenderCapacity float64 `json:"renderCapacity,omitempty"` // The time spent on rendering graph divided by render qunatum duration, and multiplied by 100. 100 means the audio renderer reached the full capacity and glitch may occur.
}

// Protocol object for BaseAudioContext
type WebAudioBaseAudioContext struct {
	ContextId             string                       `json:"contextId"`              //
	ContextType           string                       `json:"contextType"`            //  enum values: realtime, offline
	ContextState          string                       `json:"contextState"`           //  enum values: suspended, running, closed
	RealtimeData          *WebAudioContextRealtimeData `json:"realtimeData,omitempty"` //
	CallbackBufferSize    float64                      `json:"callbackBufferSize"`     // Platform-dependent callback buffer size.
	MaxOutputChannelCount float64                      `json:"maxOutputChannelCount"`  // Number of output channels supported by audio hardware in use.
	SampleRate            float64                      `json:"sampleRate"`             // Context sample rate.
}

// Notifies that a new BaseAudioContext has been created.
type WebAudioContextCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *WebAudioBaseAudioContext `json:"context"` //
	} `json:"Params,omitempty"`
}

// Notifies that existing BaseAudioContext has been destroyed.
type WebAudioContextDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId string `json:"contextId"` //
	} `json:"Params,omitempty"`
}

// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
type WebAudioContextChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *WebAudioBaseAudioContext `json:"context"` //
	} `json:"Params,omitempty"`
}

type WebAudio struct {
	target gcdmessage.ChromeTargeter
}

func NewWebAudio(target gcdmessage.ChromeTargeter) *WebAudio {
	c := &WebAudio{target: target}
	return c
}

// Enables the WebAudio domain and starts sending context lifetime events.
func (c *WebAudio) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.enable"})
}

// Disables the WebAudio domain.
func (c *WebAudio) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.disable"})
}

type WebAudioGetRealtimeDataParams struct {
	//
	ContextId string `json:"contextId"`
}

// GetRealtimeDataWithParams - Fetch the realtime data from the registered contexts.
// Returns -  realtimeData -
func (c *WebAudio) GetRealtimeDataWithParams(v *WebAudioGetRealtimeDataParams) (*WebAudioContextRealtimeData, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.getRealtimeData", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RealtimeData *WebAudioContextRealtimeData
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.RealtimeData, nil
}

// GetRealtimeData - Fetch the realtime data from the registered contexts.
// contextId -
// Returns -  realtimeData -
func (c *WebAudio) GetRealtimeData(contextId string) (*WebAudioContextRealtimeData, error) {
	var v WebAudioGetRealtimeDataParams
	v.ContextId = contextId
	return c.GetRealtimeDataWithParams(&v)
}
