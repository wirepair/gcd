// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOMDebugger commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) DOMDebugger() *ChromeDOMDebugger {
	if c.domdebugger == nil {
		c.domdebugger = newChromeDOMDebugger(c)
	}
	return c.domdebugger
}


type ChromeDOMDebugger struct {
	target *ChromeTarget
}

func newChromeDOMDebugger(target *ChromeTarget) *ChromeDOMDebugger {
	c := &ChromeDOMDebugger{target: target}
	return c
}



// setDOMBreakpoint - Sets breakpoint on particular operation with DOM.
// nodeId - Identifier of the node to set breakpoint on.
// type - Type of the operation to stop upon.
func (c *ChromeDOMDebugger) SetDOMBreakpoint(nodeId *types.ChromeDOMNodeId, theType *types.ChromeDOMDebuggerDOMBreakpointType, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["type"] = theType
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.setDOMBreakpoint", Params: paramRequest})
}

// removeDOMBreakpoint - Removes DOM breakpoint that was set using <code>setDOMBreakpoint</code>.
// nodeId - Identifier of the node to remove breakpoint from.
// type - Type of the breakpoint to remove.
func (c *ChromeDOMDebugger) RemoveDOMBreakpoint(nodeId *types.ChromeDOMNodeId, theType *types.ChromeDOMDebuggerDOMBreakpointType, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["type"] = theType
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.removeDOMBreakpoint", Params: paramRequest})
}

// setEventListenerBreakpoint - Sets breakpoint on particular DOM event.
// eventName - DOM Event name to stop on (any DOM event will do).
// targetName - EventTarget interface name to stop on. If equal to <code>"*"</code> or not provided, will stop on any EventTarget.
func (c *ChromeDOMDebugger) SetEventListenerBreakpoint(eventName string, targetName string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["eventName"] = eventName
	paramRequest["targetName"] = targetName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.setEventListenerBreakpoint", Params: paramRequest})
}

// removeEventListenerBreakpoint - Removes breakpoint on particular DOM event.
// eventName - Event name.
// targetName - EventTarget interface name.
func (c *ChromeDOMDebugger) RemoveEventListenerBreakpoint(eventName string, targetName string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["eventName"] = eventName
	paramRequest["targetName"] = targetName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.removeEventListenerBreakpoint", Params: paramRequest})
}

// setInstrumentationBreakpoint - Sets breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *ChromeDOMDebugger) SetInstrumentationBreakpoint(eventName string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["eventName"] = eventName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.setInstrumentationBreakpoint", Params: paramRequest})
}

// removeInstrumentationBreakpoint - Removes breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *ChromeDOMDebugger) RemoveInstrumentationBreakpoint(eventName string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["eventName"] = eventName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.removeInstrumentationBreakpoint", Params: paramRequest})
}

// setXHRBreakpoint - Sets breakpoint on XMLHttpRequest.
// url - Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
func (c *ChromeDOMDebugger) SetXHRBreakpoint(url string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.setXHRBreakpoint", Params: paramRequest})
}

// removeXHRBreakpoint - Removes breakpoint from XMLHttpRequest.
// url - Resource URL substring.
func (c *ChromeDOMDebugger) RemoveXHRBreakpoint(url string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMDebugger.removeXHRBreakpoint", Params: paramRequest})
}






