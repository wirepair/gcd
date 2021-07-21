// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Console functionality.
// API Version: 1.3

package gcdapi

import (
	"context"

	"github.com/wirepair/gcd/v3/gcdmessage"
)

// Console message.
type ConsoleConsoleMessage struct {
	Source string `json:"source"`           // Message source.
	Level  string `json:"level"`            // Message severity.
	Text   string `json:"text"`             // Message text.
	Url    string `json:"url,omitempty"`    // URL of the message origin.
	Line   int    `json:"line,omitempty"`   // Line number in the resource that generated this message (1-based).
	Column int    `json:"column,omitempty"` // Column number in the resource that generated this message (1-based).
}

// Issued when new console message is added.
type ConsoleMessageAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Message *ConsoleConsoleMessage `json:"message"` // Console message that has been added.
	} `json:"Params,omitempty"`
}

type Console struct {
	target gcdmessage.ChromeTargeter
}

func NewConsole(target gcdmessage.ChromeTargeter) *Console {
	c := &Console{target: target}
	return c
}

// Does nothing.
func (c *Console) ClearMessages(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.clearMessages"})
}

// Disables console domain, prevents further console messages from being reported to the client.
func (c *Console) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.disable"})
}

// Enables console domain, sends the messages collected so far to the client by means of the `messageAdded` notification.
func (c *Console) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.enable"})
}
