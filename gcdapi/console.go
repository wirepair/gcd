// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Console functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Console message.
type ConsoleConsoleMessage struct {
	Source             string                 `json:"source"`                       // Message source.
	Level              string                 `json:"level"`                        // Message severity.
	Text               string                 `json:"text"`                         // Message text.
	Type               string                 `json:"type,omitempty"`               // Console message type.
	ScriptId           string                 `json:"scriptId,omitempty"`           // Script ID of the message origin.
	Url                string                 `json:"url,omitempty"`                // URL of the message origin.
	Line               int                    `json:"line,omitempty"`               // Line number in the resource that generated this message.
	Column             int                    `json:"column,omitempty"`             // Column number in the resource that generated this message.
	RepeatCount        int                    `json:"repeatCount,omitempty"`        // Repeat count for repeated messages.
	Parameters         []*RuntimeRemoteObject `json:"parameters,omitempty"`         // Message parameters in case of the formatted message.
	Stack              *RuntimeStackTrace     `json:"stack,omitempty"`              // JavaScript stack trace for assertions and error messages.
	NetworkRequestId   string                 `json:"networkRequestId,omitempty"`   // Identifier of the network request associated with this message.
	Timestamp          float64                `json:"timestamp"`                    // Timestamp, when this message was fired.
	ExecutionContextId int                    `json:"executionContextId,omitempty"` // Identifier of the context where this message was created
	MessageId          int                    `json:"messageId,omitempty"`          // Message id.
	RelatedMessageId   int                    `json:"relatedMessageId,omitempty"`   // Related message id.
}

// Issued when new console message is added.
type ConsoleMessageAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Message *ConsoleConsoleMessage `json:"message"` // Console message that has been added.
	} `json:"Params,omitempty"`
}

// Is not issued. Will be gone in the future versions of the protocol.
type ConsoleMessageRepeatCountUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Count     int     `json:"count"`     // New repeat count value.
		Timestamp float64 `json:"timestamp"` // Timestamp of most recent message in batch.
	} `json:"Params,omitempty"`
}

type Console struct {
	target gcdmessage.ChromeTargeter
}

func NewConsole(target gcdmessage.ChromeTargeter) *Console {
	c := &Console{target: target}
	return c
}

// Enables console domain, sends the messages collected so far to the client by means of the <code>messageAdded</code> notification.
func (c *Console) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.enable"})
}

// Disables console domain, prevents further console messages from being reported to the client.
func (c *Console) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.disable"})
}

// Clears console messages collected in the browser.
func (c *Console) ClearMessages() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.clearMessages"})
}
