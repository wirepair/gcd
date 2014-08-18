// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Runtime commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Runtime() *ChromeRuntime {
	if c.runtime == nil {
		c.runtime = newChromeRuntime(c)
	}
	return c.runtime
}


type ChromeRuntime struct {
	target *ChromeTarget
}

func newChromeRuntime(target *ChromeTarget) *ChromeRuntime {
	c := &ChromeRuntime{target: target}
	return c
}

// start non parameterized commands 
// Tells inspected instance(worker or page) that it can run in case it was started paused.
func (c *ChromeRuntime) Run() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.run"})
}
 
// Enables reporting of execution contexts creation by means of <code>executionContextCreated</code> event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (c *ChromeRuntime) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.enable"})
}
 
// Disables reporting of execution contexts creation.
func (c *ChromeRuntime) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// releaseObject - Releases remote object with given id.
// objectId - Identifier of the object to release.
func (c *ChromeRuntime) ReleaseObject(objectId *types.ChromeRuntimeRemoteObjectId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.releaseObject", Params: paramRequest})
}

// releaseObjectGroup - Releases all remote objects that belong to a given group.
// objectGroup - Symbolic object group name.
func (c *ChromeRuntime) ReleaseObjectGroup(objectGroup string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectGroup"] = objectGroup
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.releaseObjectGroup", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// isRunRequired - 
// Returns - 
// True if the Runtime is in paused on start state.
func (c *ChromeRuntime) IsRunRequired() (bool, error) {	
	var result bool 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.isRunRequired"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return result, &ChromeRequestErr{Resp: cerr}
		}
		return result, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		result = results["result"].(bool)
	}
	return result, nil
}


// end commands with no parameters but special return types

