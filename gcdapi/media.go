// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Media functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Player Property type
type MediaPlayerProperty struct {
	Name  string `json:"name"`            //
	Value string `json:"value,omitempty"` //
}

// No Description.
type MediaPlayerEvent struct {
	Type      string  `json:"type"`      //  enum values: errorEvent, triggeredEvent, messageEvent
	Timestamp float64 `json:"timestamp"` // Events are timestamped relative to the start of the player creation not relative to the start of playback.
	Name      string  `json:"name"`      //
	Value     string  `json:"value"`     //
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

// Called whenever a player is created, or when a new agent joins and recieves a list of active players. If an agent is restored, it will recieve the full list of player ids and all events again.
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
func (c *Media) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Media.enable"})
}

// Disables the Media domain.
func (c *Media) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Media.disable"})
}
