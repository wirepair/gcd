// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Debugger commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
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

// start non parameterized commands 
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

// end non parameterized commands

// start parameterized commands with no special return types

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


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// canSetScriptSource - Always returns true.
// Returns - 
// True if <code>setScriptSource</code> is supported.
func (c *ChromeDebugger) CanSetScriptSource() (bool, error) {	
	var result bool 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.canSetScriptSource"})
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

// getBacktrace - Returns call stack including variables changed since VM was paused. VM must be paused.
// Returns - 
// Call stack the virtual machine stopped on.
// Async stack trace, if any.
func (c *ChromeDebugger) GetBacktrace() ([]*types.ChromeDebuggerCallFrame, *types.ChromeDebuggerStackTrace, error) {	
	var callFrames []*types.ChromeDebuggerCallFrame 
	var asyncStackTrace *types.ChromeDebuggerStackTrace 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Debugger.getBacktrace"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return callFrames, asyncStackTrace, &ChromeRequestErr{Resp: cerr}
		}
		return callFrames, asyncStackTrace, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		callFrames = results["callFrames"].([]*types.ChromeDebuggerCallFrame)
		asyncStackTrace = results["asyncStackTrace"].(*types.ChromeDebuggerStackTrace)
	}
	return callFrames, asyncStackTrace, nil
}


// end commands with no parameters but special return types

