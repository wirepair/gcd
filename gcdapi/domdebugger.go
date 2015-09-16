// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMDebugger functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Object event listener.
type DOMDebuggerEventListener struct {
	Type       string               `json:"type"`              // <code>EventListener</code>'s type.
	UseCapture bool                 `json:"useCapture"`        // <code>EventListener</code>'s useCapture.
	Location   *DebuggerLocation    `json:"location"`          // Handler code location.
	Handler    *RuntimeRemoteObject `json:"handler,omitempty"` // Event handler function value.
}

type DOMDebugger struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMDebugger(target gcdmessage.ChromeTargeter) *DOMDebugger {
	c := &DOMDebugger{target: target}
	return c
}

// SetDOMBreakpoint - Sets breakpoint on particular operation with DOM.
// nodeId - Identifier of the node to set breakpoint on.
// type - Type of the operation to stop upon. enum values: subtree-modified, attribute-modified, node-removed
func (c *DOMDebugger) SetDOMBreakpoint(nodeId int, theType string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["type"] = theType
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setDOMBreakpoint", Params: paramRequest})
}

// RemoveDOMBreakpoint - Removes DOM breakpoint that was set using <code>setDOMBreakpoint</code>.
// nodeId - Identifier of the node to remove breakpoint from.
// type - Type of the breakpoint to remove. enum values: subtree-modified, attribute-modified, node-removed
func (c *DOMDebugger) RemoveDOMBreakpoint(nodeId int, theType string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["type"] = theType
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeDOMBreakpoint", Params: paramRequest})
}

// SetEventListenerBreakpoint - Sets breakpoint on particular DOM event.
// eventName - DOM Event name to stop on (any DOM event will do).
// targetName - EventTarget interface name to stop on. If equal to <code>"*"</code> or not provided, will stop on any EventTarget.
func (c *DOMDebugger) SetEventListenerBreakpoint(eventName string, targetName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["eventName"] = eventName
	paramRequest["targetName"] = targetName
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setEventListenerBreakpoint", Params: paramRequest})
}

// RemoveEventListenerBreakpoint - Removes breakpoint on particular DOM event.
// eventName - Event name.
// targetName - EventTarget interface name.
func (c *DOMDebugger) RemoveEventListenerBreakpoint(eventName string, targetName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["eventName"] = eventName
	paramRequest["targetName"] = targetName
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeEventListenerBreakpoint", Params: paramRequest})
}

// SetInstrumentationBreakpoint - Sets breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *DOMDebugger) SetInstrumentationBreakpoint(eventName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["eventName"] = eventName
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setInstrumentationBreakpoint", Params: paramRequest})
}

// RemoveInstrumentationBreakpoint - Removes breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *DOMDebugger) RemoveInstrumentationBreakpoint(eventName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["eventName"] = eventName
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeInstrumentationBreakpoint", Params: paramRequest})
}

// SetXHRBreakpoint - Sets breakpoint on XMLHttpRequest.
// url - Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
func (c *DOMDebugger) SetXHRBreakpoint(url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setXHRBreakpoint", Params: paramRequest})
}

// RemoveXHRBreakpoint - Removes breakpoint from XMLHttpRequest.
// url - Resource URL substring.
func (c *DOMDebugger) RemoveXHRBreakpoint(url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeXHRBreakpoint", Params: paramRequest})
}

// GetEventListeners - Returns event listeners of the given object.
// objectId - Identifier of the object to return listeners for.
// Returns -  listeners - Array of relevant listeners.
func (c *DOMDebugger) GetEventListeners(objectId string) ([]*DOMDebuggerEventListener, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.getEventListeners", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Listeners []*DOMDebuggerEventListener
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Listeners, nil
}
