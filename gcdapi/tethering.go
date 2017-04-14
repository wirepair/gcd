// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Tethering functionality.
// API Version: 1.2

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
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

// Bind - Request browser port binding.
// port - Port number to bind.
func (c *Tethering) Bind(port int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["port"] = port
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tethering.bind", Params: paramRequest})
}

// Unbind - Request browser port unbinding.
// port - Port number to unbind.
func (c *Tethering) Unbind(port int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["port"] = port
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tethering.unbind", Params: paramRequest})
}
