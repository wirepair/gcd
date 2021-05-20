// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Media functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Have one type per entry in MediaLogRecord::Type Corresponds to kMessage
type MediaPlayerMessage struct {
	Level   string `json:"level"`   // Keep in sync with MediaLogMessageLevel We are currently keeping the message level 'error' separate from the PlayerError type because right now they represent different things, this one being a DVLOG(ERROR) style log message that gets printed based on what log level is selected in the UI, and the other is a representation of a media::PipelineStatus object. Soon however we're going to be moving away from using PipelineStatus for errors and introducing a new error type which should hopefully let us integrate the error log level into the PlayerError type.
	Message string `json:"message"` //
}

// Corresponds to kMediaPropertyChange
type MediaPlayerProperty struct {
	Name  string `json:"name"`  //
	Value string `json:"value"` //
}

// Corresponds to kMediaEventTriggered
type MediaPlayerEvent struct {
	Timestamp float64 `json:"timestamp"` //
	Value     string  `json:"value"`     //
}

// Corresponds to kMediaError
type MediaPlayerError struct {
	Type      string `json:"type"`      //
	ErrorCode string `json:"errorCode"` // When this switches to using media::Status instead of PipelineStatus we can remove "errorCode" and replace it with the fields from a Status instance. This also seems like a duplicate of the error level enum - there is a todo bug to have that level removed and use this instead. (crbug.com/1068454)
}

// This can be called multiple times, and can be used to set / override / remove player properties. A null propValue indicates removal.
type MediaPlayerPropertiesChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		PlayerId   string                 `json:"playerId"`   //
		Properties []*MediaPlayerProperty `json:"properties"` //
	} `json:"Params,omitempty"`
}

// Send events as a list, allowing them to be batched on the browser for less congestion. If batched, events must ALWAYS be in chronological order.
type MediaPlayerEventsAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		PlayerId string              `json:"playerId"` //
		Events   []*MediaPlayerEvent `json:"events"`   //
	} `json:"Params,omitempty"`
}

// Send a list of any messages that need to be delivered.
type MediaPlayerMessagesLoggedEvent struct {
	Method string `json:"method"`
	Params struct {
		PlayerId string                `json:"playerId"` //
		Messages []*MediaPlayerMessage `json:"messages"` //
	} `json:"Params,omitempty"`
}

// Send a list of any errors that need to be delivered.
type MediaPlayerErrorsRaisedEvent struct {
	Method string `json:"method"`
	Params struct {
		PlayerId string              `json:"playerId"` //
		Errors   []*MediaPlayerError `json:"errors"`   //
	} `json:"Params,omitempty"`
}

// Called whenever a player is created, or when a new agent joins and receives a list of active players. If an agent is restored, it will receive the full list of player ids and all events again.
type MediaPlayersCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Players []string `json:"players"` //
	} `json:"Params,omitempty"`
}

type Media struct {
	target gcdmessage.ChromeTargeter
}

func NewMedia(target gcdmessage.ChromeTargeter) *Media {
	c := &Media{target: target}
	return c
}

// Enables the Media domain
func (c *Media) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Media.enable"})
}

// Disables the Media domain.
func (c *Media) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Media.disable"})
}
