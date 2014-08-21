// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Debugger commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Debugger() *ChromeDebugger {
	if c.debugger == nil {
		c.debugger = newChromeDebugger(c)
	}
	return c.debugger
}

type ChromeDebugger struct {
	target *ChromeTarget
}

func newChromeDebugger(target *ChromeTarget) *ChromeDebugger {
	c := &ChromeDebugger{target: target}
	return c
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (c *ChromeDebugger) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.enable"})
}

// Disables debugger for given page.
func (c *ChromeDebugger) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.disable"})
}

// Steps over the statement.
func (c *ChromeDebugger) StepOver() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.stepOver"})
}

// Steps into the function call.
func (c *ChromeDebugger) StepInto() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.stepInto"})
}

// Steps out of the function call.
func (c *ChromeDebugger) StepOut() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.stepOut"})
}

// Stops on the next JavaScript statement.
func (c *ChromeDebugger) Pause() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.pause"})
}

// Resumes JavaScript execution.
func (c *ChromeDebugger) Resume() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.resume"})
}

// setBreakpointsActive - Activates / deactivates all breakpoints on the page.
// active - New value for breakpoints active state.
func (c *ChromeDebugger) SetBreakpointsActive(active bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["active"] = active
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setBreakpointsActive", Params: paramRequest})
}

// setSkipAllPauses - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
// skipped - New value for skip pauses state.
// untilReload - Whether page reload should set skipped to false.
func (c *ChromeDebugger) SetSkipAllPauses(skipped bool, untilReload bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["skipped"] = skipped
	paramRequest["untilReload"] = untilReload
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setSkipAllPauses", Params: paramRequest})
}

// removeBreakpoint - Removes JavaScript breakpoint.
// breakpointId -
func (c *ChromeDebugger) RemoveBreakpoint(breakpointId *types.ChromeDebuggerBreakpointId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["breakpointId"] = breakpointId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.removeBreakpoint", Params: paramRequest})
}

// continueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// interstatementLocation - Allows breakpoints at the intemediate positions inside statements.
func (c *ChromeDebugger) ContinueToLocation(location *types.ChromeDebuggerLocation, interstatementLocation bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["interstatementLocation"] = interstatementLocation
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.continueToLocation", Params: paramRequest})
}

// setPauseOnExceptions - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
// state - Pause on exceptions mode.
func (c *ChromeDebugger) SetPauseOnExceptions(state string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["state"] = state
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setPauseOnExceptions", Params: paramRequest})
}

// setOverlayMessage - Sets overlay message.
// message - Overlay message to display when paused in debugger.
func (c *ChromeDebugger) SetOverlayMessage(message string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["message"] = message
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setOverlayMessage", Params: paramRequest})
}

// setVariableValue - Changes value of variable in a callframe or a closure. Either callframe or function must be specified. Object-based scopes are not supported and must be mutated manually.
// scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
// variableName - Variable name.
// newValue - New variable value.
// callFrameId - Id of callframe that holds variable.
// functionObjectId - Object id of closure (function) that holds variable.
func (c *ChromeDebugger) SetVariableValue(scopeNumber int, variableName string, newValue *types.ChromeRuntimeCallArgument, callFrameId *types.ChromeDebuggerCallFrameId, functionObjectId *types.ChromeRuntimeRemoteObjectId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["scopeNumber"] = scopeNumber
	paramRequest["variableName"] = variableName
	paramRequest["newValue"] = newValue
	paramRequest["callFrameId"] = callFrameId
	paramRequest["functionObjectId"] = functionObjectId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setVariableValue", Params: paramRequest})
}

// skipStackFrames - Makes backend skip steps in the sources with names matching given pattern. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// script - Regular expression defining the scripts to ignore while stepping.
func (c *ChromeDebugger) SkipStackFrames(script string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["script"] = script
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.skipStackFrames", Params: paramRequest})
}

// setAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (c *ChromeDebugger) SetAsyncCallStackDepth(maxDepth int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["maxDepth"] = maxDepth
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setAsyncCallStackDepth", Params: paramRequest})
}

// canSetScriptSource - Always returns true.
// Returns -
// True if <code>setScriptSource</code> is supported.
func (c *ChromeDebugger) CanSetScriptSource() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.canSetScriptSource"})
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

// getBacktrace - Returns call stack including variables changed since VM was paused. VM must be paused.
// Returns -
// Call stack the virtual machine stopped on.
// Async stack trace, if any.
func (c *ChromeDebugger) GetBacktrace() ([]*types.ChromeDebuggerCallFrame, *types.ChromeDebuggerStackTrace, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.getBacktrace"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*types.ChromeDebuggerCallFrame
			AsyncStackTrace *types.ChromeDebuggerStackTrace
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

	return chromeData.Result.CallFrames, chromeData.Result.AsyncStackTrace, nil
}

// setBreakpointByUrl - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
// Returns -
// Id of the created breakpoint for further reference.
// List of the locations this breakpoint resolved into upon addition.
func (c *ChromeDebugger) SetBreakpointByUrl(lineNumber int, url string, urlRegex string, columnNumber int, condition string, isAntibreakpoint bool) (*types.ChromeDebuggerBreakpointId, []*types.ChromeDebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 6)
	paramRequest["lineNumber"] = lineNumber
	paramRequest["url"] = url
	paramRequest["urlRegex"] = urlRegex
	paramRequest["columnNumber"] = columnNumber
	paramRequest["condition"] = condition
	paramRequest["isAntibreakpoint"] = isAntibreakpoint
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setBreakpointByUrl", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			BreakpointId *types.ChromeDebuggerBreakpointId
			Locations    []*types.ChromeDebuggerLocation
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

	return chromeData.Result.BreakpointId, chromeData.Result.Locations, nil
}

// setBreakpoint - Sets JavaScript breakpoint at a given location.
// Returns -
// Id of the created breakpoint for further reference.
// Location this breakpoint resolved into.
func (c *ChromeDebugger) SetBreakpoint(location *types.ChromeDebuggerLocation, condition string) (*types.ChromeDebuggerBreakpointId, *types.ChromeDebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["condition"] = condition
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setBreakpoint", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			BreakpointId   *types.ChromeDebuggerBreakpointId
			ActualLocation *types.ChromeDebuggerLocation
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

	return chromeData.Result.BreakpointId, chromeData.Result.ActualLocation, nil
}

// searchInContent - Searches for given string in script content.
// Returns -
// List of search matches.
func (c *ChromeDebugger) SearchInContent(scriptId *types.ChromeDebuggerScriptId, query string, caseSensitive bool, isRegex bool) ([]*types.ChromePageSearchMatch, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["scriptId"] = scriptId
	paramRequest["query"] = query
	paramRequest["caseSensitive"] = caseSensitive
	paramRequest["isRegex"] = isRegex
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.searchInContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result []*types.ChromePageSearchMatch
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Result, nil
}

// setScriptSource - Edits JavaScript source live.
// Returns -
// New stack trace in case editing has happened while VM was stopped.
// VM-specific description of the changes applied.
// Async stack trace, if any.
func (c *ChromeDebugger) SetScriptSource(scriptId *types.ChromeDebuggerScriptId, scriptSource string, preview bool) ([]*types.ChromeDebuggerCallFrame, interface{}, *types.ChromeDebuggerStackTrace, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["scriptId"] = scriptId
	paramRequest["scriptSource"] = scriptSource
	paramRequest["preview"] = preview
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.setScriptSource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*types.ChromeDebuggerCallFrame
			Result          interface{}
			AsyncStackTrace *types.ChromeDebuggerStackTrace
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.Result, chromeData.Result.AsyncStackTrace, nil
}

// restartFrame - Restarts particular call frame from the beginning.
// Returns -
// New stack trace.
// VM-specific description.
// Async stack trace, if any.
func (c *ChromeDebugger) RestartFrame(callFrameId *types.ChromeDebuggerCallFrameId) ([]*types.ChromeDebuggerCallFrame, interface{}, *types.ChromeDebuggerStackTrace, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["callFrameId"] = callFrameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.restartFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*types.ChromeDebuggerCallFrame
			Result          interface{}
			AsyncStackTrace *types.ChromeDebuggerStackTrace
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.Result, chromeData.Result.AsyncStackTrace, nil
}

// getScriptSource - Returns source for the script with given id.
// Returns -
// Script source.
func (c *ChromeDebugger) GetScriptSource(scriptId *types.ChromeDebuggerScriptId) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scriptId"] = scriptId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.getScriptSource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ScriptSource string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.ScriptSource, nil
}

// getFunctionDetails - Returns detailed informtation on given function.
// Returns -
// Information about the function.
func (c *ChromeDebugger) GetFunctionDetails(functionId *types.ChromeRuntimeRemoteObjectId) (*types.ChromeDebuggerFunctionDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["functionId"] = functionId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.getFunctionDetails", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Details *types.ChromeDebuggerFunctionDetails
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Details, nil
}

// evaluateOnCallFrame - Evaluates expression on a given call frame.
// Returns -
// Object wrapper for the evaluation result.
// True if the result was thrown during the evaluation.
// Exception details.
func (c *ChromeDebugger) EvaluateOnCallFrame(callFrameId *types.ChromeDebuggerCallFrameId, expression string, objectGroup string, includeCommandLineAPI bool, doNotPauseOnExceptionsAndMuteConsole bool, returnByValue bool, generatePreview bool) (*types.ChromeRuntimeRemoteObject, bool, *types.ChromeDebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["callFrameId"] = callFrameId
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.evaluateOnCallFrame", Params: paramRequest})
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

// compileScript - Compiles expression.
// Returns -
// Id of the script.
// Exception details.
func (c *ChromeDebugger) CompileScript(expression string, sourceURL string, executionContextId *types.ChromeRuntimeExecutionContextId) (*types.ChromeDebuggerScriptId, *types.ChromeDebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["expression"] = expression
	paramRequest["sourceURL"] = sourceURL
	paramRequest["executionContextId"] = executionContextId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.compileScript", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ScriptId         *types.ChromeDebuggerScriptId
			ExceptionDetails *types.ChromeDebuggerExceptionDetails
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

	return chromeData.Result.ScriptId, chromeData.Result.ExceptionDetails, nil
}

// runScript - Runs script with given id in a given context.
// Returns -
// Run result.
// Exception details.
func (c *ChromeDebugger) RunScript(scriptId *types.ChromeDebuggerScriptId, executionContextId *types.ChromeRuntimeExecutionContextId, objectGroup string, doNotPauseOnExceptionsAndMuteConsole bool) (*types.ChromeRuntimeRemoteObject, *types.ChromeDebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["scriptId"] = scriptId
	paramRequest["executionContextId"] = executionContextId
	paramRequest["objectGroup"] = objectGroup
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.runScript", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result           *types.ChromeRuntimeRemoteObject
			ExceptionDetails *types.ChromeDebuggerExceptionDetails
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

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// getStepInPositions - Lists all positions where step-in is possible for a current statement in a specified call frame
// Returns -
// experimental
func (c *ChromeDebugger) GetStepInPositions(callFrameId *types.ChromeDebuggerCallFrameId) ([]*types.ChromeDebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["callFrameId"] = callFrameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.getStepInPositions", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			StepInPositions []*types.ChromeDebuggerLocation
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.StepInPositions, nil
}
