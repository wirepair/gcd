// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Runtime commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
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
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.isRunRequired"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}

// end commands with no parameters but special return types

// start commands with parameters and special return types

// evaluate - Evaluates expression on global object.
// Returns -
// Evaluation result.
// True if the result was thrown during the evaluation.
// Exception details.
func (c *ChromeRuntime) Evaluate(expression string, objectGroup string, includeCommandLineAPI bool, doNotPauseOnExceptionsAndMuteConsole bool, contextId *types.ChromeRuntimeExecutionContextId, returnByValue bool, generatePreview bool) (*types.ChromeRuntimeRemoteObject, bool, *types.ChromeDebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["contextId"] = contextId
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.evaluate", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result           *types.ChromeRuntimeRemoteObject
			WasThrown        bool
			ExceptionDetails *types.ChromeDebuggerExceptionDetails
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, false, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, false, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, chromeData.Result.ExceptionDetails, nil
}

// callFunctionOn - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// Returns -
// Call result.
// True if the result was thrown during the evaluation.
func (c *ChromeRuntime) CallFunctionOn(objectId *types.ChromeRuntimeRemoteObjectId, functionDeclaration string, arguments []*types.ChromeRuntimeCallArgument, doNotPauseOnExceptionsAndMuteConsole bool, returnByValue bool, generatePreview bool) (*types.ChromeRuntimeRemoteObject, bool, error) {
	paramRequest := make(map[string]interface{}, 6)
	paramRequest["objectId"] = objectId
	paramRequest["functionDeclaration"] = functionDeclaration
	paramRequest["arguments"] = arguments
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.callFunctionOn", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result    *types.ChromeRuntimeRemoteObject
			WasThrown bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, false, &ChromeRequestErr{Resp: cerr}
		}
		return nil, false, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, nil
}

// getProperties - Returns properties of a given object. Object group of the result is inherited from the target object.
// Returns -
// Object properties.
// Internal object properties (only of the element itself).
func (c *ChromeRuntime) GetProperties(objectId *types.ChromeRuntimeRemoteObjectId, ownProperties bool, accessorPropertiesOnly bool) ([]*types.ChromeRuntimePropertyDescriptor, []*types.ChromeRuntimeInternalPropertyDescriptor, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["objectId"] = objectId
	paramRequest["ownProperties"] = ownProperties
	paramRequest["accessorPropertiesOnly"] = accessorPropertiesOnly
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Runtime.getProperties", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result             []*types.ChromeRuntimePropertyDescriptor
			InternalProperties []*types.ChromeRuntimeInternalPropertyDescriptor
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.InternalProperties, nil
}

// end commands with parameters and special return types
