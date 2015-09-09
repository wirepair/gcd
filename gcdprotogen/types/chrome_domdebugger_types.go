// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOMDebugger types.
// API Version: 1.1
package types

// DOM breakpoint type.
type ChromeDOMDebuggerDOMBreakpointType string // possible values: subtree-modified, attribute-modified, node-removed

// Object event listener.
type ChromeDOMDebuggerEventListener struct {
	Type       string                     `json:"type"`              // <code>EventListener</code>'s type.
	UseCapture bool                       `json:"useCapture"`        // <code>EventListener</code>'s useCapture.
	Location   *ChromeDebuggerLocation    `json:"location"`          // Handler code location.
	Handler    *ChromeRuntimeRemoteObject `json:"handler,omitempty"` // Event handler function value.
}
