// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains WebAudio functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Fields in AudioContext that change in real-time.
type WebAudioContextRealtimeData struct {
	CurrentTime              float64 `json:"currentTime"`              // The current context time in second in BaseAudioContext.
	RenderCapacity           float64 `json:"renderCapacity"`           // The time spent on rendering graph divided by render quantum duration, and multiplied by 100. 100 means the audio renderer reached the full capacity and glitch may occur.
	CallbackIntervalMean     float64 `json:"callbackIntervalMean"`     // A running mean of callback interval.
	CallbackIntervalVariance float64 `json:"callbackIntervalVariance"` // A running variance of callback interval.
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

// Protocol object for AudioListener
type WebAudioAudioListener struct {
	ListenerId string `json:"listenerId"` //
	ContextId  string `json:"contextId"`  //
}

// Protocol object for AudioNode
type WebAudioAudioNode struct {
	NodeId                string  `json:"nodeId"`                //
	ContextId             string  `json:"contextId"`             //
	NodeType              string  `json:"nodeType"`              //
	NumberOfInputs        float64 `json:"numberOfInputs"`        //
	NumberOfOutputs       float64 `json:"numberOfOutputs"`       //
	ChannelCount          float64 `json:"channelCount"`          //
	ChannelCountMode      string  `json:"channelCountMode"`      //  enum values: clamped-max, explicit, max
	ChannelInterpretation string  `json:"channelInterpretation"` //  enum values: discrete, speakers
}

// Protocol object for AudioParam
type WebAudioAudioParam struct {
	ParamId      string  `json:"paramId"`      //
	NodeId       string  `json:"nodeId"`       //
	ContextId    string  `json:"contextId"`    //
	ParamType    string  `json:"paramType"`    //
	Rate         string  `json:"rate"`         //  enum values: a-rate, k-rate
	DefaultValue float64 `json:"defaultValue"` //
	MinValue     float64 `json:"minValue"`     //
	MaxValue     float64 `json:"maxValue"`     //
}

// Notifies that a new BaseAudioContext has been created.
type WebAudioContextCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *WebAudioBaseAudioContext `json:"context"` //
	} `json:"Params,omitempty"`
}

// Notifies that an existing BaseAudioContext will be destroyed.
type WebAudioContextWillBeDestroyedEvent struct {
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

// Notifies that the construction of an AudioListener has finished.
type WebAudioAudioListenerCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Listener *WebAudioAudioListener `json:"listener"` //
	} `json:"Params,omitempty"`
}

// Notifies that a new AudioListener has been created.
type WebAudioAudioListenerWillBeDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId  string `json:"contextId"`  //
		ListenerId string `json:"listenerId"` //
	} `json:"Params,omitempty"`
}

// Notifies that a new AudioNode has been created.
type WebAudioAudioNodeCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Node *WebAudioAudioNode `json:"node"` //
	} `json:"Params,omitempty"`
}

// Notifies that an existing AudioNode has been destroyed.
type WebAudioAudioNodeWillBeDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId string `json:"contextId"` //
		NodeId    string `json:"nodeId"`    //
	} `json:"Params,omitempty"`
}

// Notifies that a new AudioParam has been created.
type WebAudioAudioParamCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Param *WebAudioAudioParam `json:"param"` //
	} `json:"Params,omitempty"`
}

// Notifies that an existing AudioParam has been destroyed.
type WebAudioAudioParamWillBeDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId string `json:"contextId"` //
		NodeId    string `json:"nodeId"`    //
		ParamId   string `json:"paramId"`   //
	} `json:"Params,omitempty"`
}

// Notifies that two AudioNodes are connected.
type WebAudioNodesConnectedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId             string  `json:"contextId"`                       //
		SourceId              string  `json:"sourceId"`                        //
		DestinationId         string  `json:"destinationId"`                   //
		SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`     //
		DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"` //
	} `json:"Params,omitempty"`
}

// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
type WebAudioNodesDisconnectedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId             string  `json:"contextId"`                       //
		SourceId              string  `json:"sourceId"`                        //
		DestinationId         string  `json:"destinationId"`                   //
		SourceOutputIndex     float64 `json:"sourceOutputIndex,omitempty"`     //
		DestinationInputIndex float64 `json:"destinationInputIndex,omitempty"` //
	} `json:"Params,omitempty"`
}

// Notifies that an AudioNode is connected to an AudioParam.
type WebAudioNodeParamConnectedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId         string  `json:"contextId"`                   //
		SourceId          string  `json:"sourceId"`                    //
		DestinationId     string  `json:"destinationId"`               //
		SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"` //
	} `json:"Params,omitempty"`
}

// Notifies that an AudioNode is disconnected to an AudioParam.
type WebAudioNodeParamDisconnectedEvent struct {
	Method string `json:"method"`
	Params struct {
		ContextId         string  `json:"contextId"`                   //
		SourceId          string  `json:"sourceId"`                    //
		DestinationId     string  `json:"destinationId"`               //
		SourceOutputIndex float64 `json:"sourceOutputIndex,omitempty"` //
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
func (c *WebAudio) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.enable"})
}

// Disables the WebAudio domain.
func (c *WebAudio) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.disable"})
}

type WebAudioGetRealtimeDataParams struct {
	//
	ContextId string `json:"contextId"`
}

// GetRealtimeDataWithParams - Fetch the realtime data from the registered contexts.
// Returns -  realtimeData -
func (c *WebAudio) GetRealtimeDataWithParams(ctx context.Context, v *WebAudioGetRealtimeDataParams) (*WebAudioContextRealtimeData, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAudio.getRealtimeData", Params: v})
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
func (c *WebAudio) GetRealtimeData(ctx context.Context, contextId string) (*WebAudioContextRealtimeData, error) {
	var v WebAudioGetRealtimeDataParams
	v.ContextId = contextId
	return c.GetRealtimeDataWithParams(ctx, &v)
}
