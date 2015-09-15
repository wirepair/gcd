// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Console functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Console message.
type ConsoleConsoleMessage struct {
	Source             string                  `json:"source"`                       // Message source.
	Level              string                  `json:"level"`                        // Message severity.
	Text               string                  `json:"text"`                         // Message text.
	Type               string                  `json:"type,omitempty"`               // Console message type.
	ScriptId           string                  `json:"scriptId,omitempty"`           // Script ID of the message origin.
	Url                string                  `json:"url,omitempty"`                // URL of the message origin.
	Line               int                     `json:"line,omitempty"`               // Line number in the resource that generated this message.
	Column             int                     `json:"column,omitempty"`             // Column number in the resource that generated this message.
	RepeatCount        int                     `json:"repeatCount,omitempty"`        // Repeat count for repeated messages.
	Parameters         []*RuntimeRemoteObject  `json:"parameters,omitempty"`         // Message parameters in case of the formatted message.
	StackTrace         []*ConsoleCallFrame     `json:"stackTrace,omitempty"`         // JavaScript stack trace for assertions and error messages.
	AsyncStackTrace    *ConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"`    // Asynchronous JavaScript stack trace that preceded this message, if available.
	NetworkRequestId   string                  `json:"networkRequestId,omitempty"`   // Identifier of the network request associated with this message.
	Timestamp          float64                 `json:"timestamp"`                    // Timestamp, when this message was fired.
	ExecutionContextId int                     `json:"executionContextId,omitempty"` // Identifier of the context where this message was created
	MessageId          int                     `json:"messageId,omitempty"`          // Message id.
	RelatedMessageId   int                     `json:"relatedMessageId,omitempty"`   // Related message id.
}

// Stack entry for console errors and assertions.
type ConsoleCallFrame struct {
	FunctionName string `json:"functionName"` // JavaScript function name.
	ScriptId     string `json:"scriptId"`     // JavaScript script id.
	Url          string `json:"url"`          // JavaScript script name or url.
	LineNumber   int    `json:"lineNumber"`   // JavaScript script line number.
	ColumnNumber int    `json:"columnNumber"` // JavaScript script column number.
}

// Asynchronous JavaScript call stack.
type ConsoleAsyncStackTrace struct {
	CallFrames      []*ConsoleCallFrame     `json:"callFrames"`                // Call frames of the stack trace.
	Description     string                  `json:"description,omitempty"`     // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	AsyncStackTrace *ConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Next asynchronous stack trace, if any.
}

// Issued when new console message is added.
type ConsoleMessageAddedEvent struct {
	Message *ConsoleConsoleMessage `json:"message"` // Console message that has been added.
}

// Is not issued. Will be gone in the future versions of the protocol.
type ConsoleMessageRepeatCountUpdatedEvent struct {
	Count     int     `json:"count"`     // New repeat count value.
	Timestamp float64 `json:"timestamp"` // Timestamp of most recent message in batch.
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
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.enable"})
}

// Disables console domain, prevents further console messages from being reported to the client.
func (c *Console) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.disable"})
}

// Clears console messages collected in the browser.
func (c *Console) ClearMessages() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Console.clearMessages"})
}
