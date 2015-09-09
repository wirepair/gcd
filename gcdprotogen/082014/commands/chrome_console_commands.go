// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Console commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Console() *ChromeConsole {
	if c.console == nil {
		c.console = newChromeConsole(c)
	}
	return c.console
}


type ChromeConsole struct {
	target *ChromeTarget
}

func newChromeConsole(target *ChromeTarget) *ChromeConsole {
	c := &ChromeConsole{target: target}
	return c
}

 
// Enables console domain, sends the messages collected so far to the client by means of the <code>messageAdded</code> notification.
func (c *ChromeConsole) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.enable"})
}
 
// Disables console domain, prevents further console messages from being reported to the client.
func (c *ChromeConsole) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.disable"})
}
 
// Clears console messages collected in the browser.
func (c *ChromeConsole) ClearMessages() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.clearMessages"})
}


// setMonitoringXHREnabled - Toggles monitoring of XMLHttpRequest. If <code>true</code>, console will receive messages upon each XHR issued.
// enabled - Monitoring enabled state.
func (c *ChromeConsole) SetMonitoringXHREnabled(enabled bool, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.setMonitoringXHREnabled", Params: paramRequest})
}

// addInspectedNode - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// nodeId - DOM node id to be accessible by means of $x command line API.
func (c *ChromeConsole) AddInspectedNode(nodeId *types.ChromeDOMNodeId, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.addInspectedNode", Params: paramRequest})
}

// addInspectedHeapObject - 
// heapObjectId - 
func (c *ChromeConsole) AddInspectedHeapObject(heapObjectId int, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["heapObjectId"] = heapObjectId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.addInspectedHeapObject", Params: paramRequest})
}






