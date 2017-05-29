// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Debugger functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Location in the source code.
type DebuggerLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}

// Location in the source code.
type DebuggerScriptPosition struct {
	LineNumber   int `json:"lineNumber"`   //
	ColumnNumber int `json:"columnNumber"` //
}

// JavaScript call frame. Array of call frames form the call stack.
type DebuggerCallFrame struct {
	CallFrameId      string               `json:"callFrameId"`                // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName     string               `json:"functionName"`               // Name of the JavaScript function called on this call frame.
	FunctionLocation *DebuggerLocation    `json:"functionLocation,omitempty"` // Location in the source code.
	Location         *DebuggerLocation    `json:"location"`                   // Location in the source code.
	ScopeChain       []*DebuggerScope     `json:"scopeChain"`                 // Scope chain for this call frame.
	This             *RuntimeRemoteObject `json:"this"`                       // <code>this</code> object for this call frame.
	ReturnValue      *RuntimeRemoteObject `json:"returnValue,omitempty"`      // The value being returned, if the function is at return point.
}

// Scope description.
type DebuggerScope struct {
	Type          string               `json:"type"`                    // Scope type.
	Object        *RuntimeRemoteObject `json:"object"`                  // Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Name          string               `json:"name,omitempty"`          //
	StartLocation *DebuggerLocation    `json:"startLocation,omitempty"` // Location in the source code where scope starts
	EndLocation   *DebuggerLocation    `json:"endLocation,omitempty"`   // Location in the source code where scope ends
}

// Search match for resource.
type DebuggerSearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`  // Line number in resource content.
	LineContent string  `json:"lineContent"` // Line with match content.
}

// No Description.
type DebuggerBreakLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
	Type         string `json:"type,omitempty"`         //
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
type DebuggerScriptParsedEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                string                 `json:"scriptId"`                          // Identifier of the script parsed.
		Url                     string                 `json:"url"`                               // URL or name of the script parsed (if any).
		StartLine               int                    `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
		StartColumn             int                    `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
		EndLine                 int                    `json:"endLine"`                           // Last line of the script.
		EndColumn               int                    `json:"endColumn"`                         // Length of the last line of the script.
		ExecutionContextId      int                    `json:"executionContextId"`                // Specifies script creation context.
		Hash                    string                 `json:"hash"`                              // Content hash of the script.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
		IsLiveEdit              bool                   `json:"isLiveEdit,omitempty"`              // True, if this script is generated as a result of the live edit operation.
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
	} `json:"Params,omitempty"`
}

// Fired when virtual machine fails to parse the script.
type DebuggerScriptFailedToParseEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                string                 `json:"scriptId"`                          // Identifier of the script parsed.
		Url                     string                 `json:"url"`                               // URL or name of the script parsed (if any).
		StartLine               int                    `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
		StartColumn             int                    `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
		EndLine                 int                    `json:"endLine"`                           // Last line of the script.
		EndColumn               int                    `json:"endColumn"`                         // Length of the last line of the script.
		ExecutionContextId      int                    `json:"executionContextId"`                // Specifies script creation context.
		Hash                    string                 `json:"hash"`                              // Content hash of the script.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
	} `json:"Params,omitempty"`
}

// Fired when breakpoint is resolved to an actual script and location.
type DebuggerBreakpointResolvedEvent struct {
	Method string `json:"method"`
	Params struct {
		BreakpointId string            `json:"breakpointId"` // Breakpoint unique identifier.
		Location     *DebuggerLocation `json:"location"`     // Actual breakpoint location.
	} `json:"Params,omitempty"`
}

// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
type DebuggerPausedEvent struct {
	Method string `json:"method"`
	Params struct {
		CallFrames      []*DebuggerCallFrame   `json:"callFrames"`                // Call stack the virtual machine stopped on.
		Reason          string                 `json:"reason"`                    // Pause reason.
		Data            map[string]interface{} `json:"data,omitempty"`            // Object containing break-specific auxiliary properties.
		HitBreakpoints  []string               `json:"hitBreakpoints,omitempty"`  // Hit breakpoints IDs
		AsyncStackTrace *RuntimeStackTrace     `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	} `json:"Params,omitempty"`
}

type Debugger struct {
	target gcdmessage.ChromeTargeter
}

func NewDebugger(target gcdmessage.ChromeTargeter) *Debugger {
	c := &Debugger{target: target}
	return c
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (c *Debugger) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.enable"})
}

// Disables debugger for given page.
func (c *Debugger) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disable"})
}

// SetBreakpointsActive - Activates / deactivates all breakpoints on the page.
// active - New value for breakpoints active state.
func (c *Debugger) SetBreakpointsActive(active bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["active"] = active
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointsActive", Params: paramRequest})
}

// SetSkipAllPauses - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
// skip - New value for skip pauses state.
func (c *Debugger) SetSkipAllPauses(skip bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["skip"] = skip
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setSkipAllPauses", Params: paramRequest})
}

// SetBreakpointByUrl - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
// lineNumber - Line number to set breakpoint at.
// url - URL of the resources to set breakpoint on.
// urlRegex - Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
// columnNumber - Offset in the line to set breakpoint at.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. locations - List of the locations this breakpoint resolved into upon addition.
func (c *Debugger) SetBreakpointByUrl(lineNumber int, url string, urlRegex string, columnNumber int, condition string) (string, []*DebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["lineNumber"] = lineNumber
	paramRequest["url"] = url
	paramRequest["urlRegex"] = urlRegex
	paramRequest["columnNumber"] = columnNumber
	paramRequest["condition"] = condition
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointByUrl", Params: paramRequest})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			BreakpointId string
			Locations    []*DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.BreakpointId, chromeData.Result.Locations, nil
}

// SetBreakpoint - Sets JavaScript breakpoint at a given location.
// location - Location to set breakpoint in.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. actualLocation - Location this breakpoint resolved into.
func (c *Debugger) SetBreakpoint(location *DebuggerLocation, condition string) (string, *DebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["condition"] = condition
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpoint", Params: paramRequest})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			BreakpointId   string
			ActualLocation *DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.BreakpointId, chromeData.Result.ActualLocation, nil
}

// RemoveBreakpoint - Removes JavaScript breakpoint.
// breakpointId -
func (c *Debugger) RemoveBreakpoint(breakpointId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["breakpointId"] = breakpointId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.removeBreakpoint", Params: paramRequest})
}

// GetPossibleBreakpoints - Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
// start - Start of range to search possible breakpoint locations in.
// end - End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
// restrictToFunction - Only consider locations which are in the same (non-nested) function as start.
// Returns -  locations - List of the possible breakpoint locations.
func (c *Debugger) GetPossibleBreakpoints(start *DebuggerLocation, end *DebuggerLocation, restrictToFunction bool) ([]*DebuggerBreakLocation, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["start"] = start
	paramRequest["end"] = end
	paramRequest["restrictToFunction"] = restrictToFunction
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getPossibleBreakpoints", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Locations []*DebuggerBreakLocation
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Locations, nil
}

// ContinueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// targetCallFrames -
func (c *Debugger) ContinueToLocation(location *DebuggerLocation, targetCallFrames string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["targetCallFrames"] = targetCallFrames
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.continueToLocation", Params: paramRequest})
}

// Steps over the statement.
func (c *Debugger) StepOver() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOver"})
}

// Steps into the function call.
func (c *Debugger) StepInto() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepInto"})
}

// Steps out of the function call.
func (c *Debugger) StepOut() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOut"})
}

// Stops on the next JavaScript statement.
func (c *Debugger) Pause() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.pause"})
}

// Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called.
func (c *Debugger) ScheduleStepIntoAsync() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.scheduleStepIntoAsync"})
}

// Resumes JavaScript execution.
func (c *Debugger) Resume() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.resume"})
}

// SearchInContent - Searches for given string in script content.
// scriptId - Id of the script to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Debugger) SearchInContent(scriptId string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["scriptId"] = scriptId
	paramRequest["query"] = query
	paramRequest["caseSensitive"] = caseSensitive
	paramRequest["isRegex"] = isRegex
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.searchInContent", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Result, nil
}

// SetScriptSource - Edits JavaScript source live.
// scriptId - Id of the script to edit.
// scriptSource - New content of the script.
// dryRun -  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. exceptionDetails - Exception details if any.
func (c *Debugger) SetScriptSource(scriptId string, scriptSource string, dryRun bool) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["scriptId"] = scriptId
	paramRequest["scriptSource"] = scriptSource
	paramRequest["dryRun"] = dryRun
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setScriptSource", Params: paramRequest})
	if err != nil {
		return nil, false, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CallFrames       []*DebuggerCallFrame
			StackChanged     bool
			AsyncStackTrace  *RuntimeStackTrace
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, false, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.StackChanged, chromeData.Result.AsyncStackTrace, chromeData.Result.ExceptionDetails, nil
}

// RestartFrame - Restarts particular call frame from the beginning.
// callFrameId - Call frame identifier to evaluate on.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) RestartFrame(callFrameId string) ([]*DebuggerCallFrame, *RuntimeStackTrace, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["callFrameId"] = callFrameId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.restartFrame", Params: paramRequest})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			AsyncStackTrace *RuntimeStackTrace
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.AsyncStackTrace, nil
}

// GetScriptSource - Returns source for the script with given id.
// scriptId - Id of the script to get source for.
// Returns -  scriptSource - Script source.
func (c *Debugger) GetScriptSource(scriptId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scriptId"] = scriptId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getScriptSource", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			ScriptSource string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.ScriptSource, nil
}

// SetPauseOnExceptions - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
// state - Pause on exceptions mode.
func (c *Debugger) SetPauseOnExceptions(state string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["state"] = state
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setPauseOnExceptions", Params: paramRequest})
}

// EvaluateOnCallFrame - Evaluates expression on a given call frame.
// callFrameId - Call frame identifier to evaluate on.
// expression - Expression to evaluate.
// objectGroup - String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
// includeCommandLineAPI - Specifies whether command line API should be available to the evaluated expression, defaults to false.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// throwOnSideEffect - Whether to throw an exception if side effect cannot be ruled out during evaluation.
// Returns -  result - Object wrapper for the evaluation result. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrame(callFrameId string, expression string, objectGroup string, includeCommandLineAPI bool, silent bool, returnByValue bool, generatePreview bool, throwOnSideEffect bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 8)
	paramRequest["callFrameId"] = callFrameId
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["silent"] = silent
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	paramRequest["throwOnSideEffect"] = throwOnSideEffect
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.evaluateOnCallFrame", Params: paramRequest})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// SetVariableValue - Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
// scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
// variableName - Variable name.
// newValue - New variable value.
// callFrameId - Id of callframe that holds variable.
func (c *Debugger) SetVariableValue(scopeNumber int, variableName string, newValue *RuntimeCallArgument, callFrameId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["scopeNumber"] = scopeNumber
	paramRequest["variableName"] = variableName
	paramRequest["newValue"] = newValue
	paramRequest["callFrameId"] = callFrameId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setVariableValue", Params: paramRequest})
}

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (c *Debugger) SetAsyncCallStackDepth(maxDepth int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["maxDepth"] = maxDepth
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncCallStackDepth", Params: paramRequest})
}

// SetBlackboxPatterns - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// patterns - Array of regexps that will be used to check script url for blackbox state.
func (c *Debugger) SetBlackboxPatterns(patterns []string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["patterns"] = patterns
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxPatterns", Params: paramRequest})
}

// SetBlackboxedRanges - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
// scriptId - Id of the script.
// positions -
func (c *Debugger) SetBlackboxedRanges(scriptId string, positions []*DebuggerScriptPosition) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["scriptId"] = scriptId
	paramRequest["positions"] = positions
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxedRanges", Params: paramRequest})
}
