// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Debugger functionality.
// API Version: 1.1

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
	Line   int `json:"line"`   //
	Column int `json:"column"` //
}

// Information about the function.
type DebuggerFunctionDetails struct {
	Location     *DebuggerLocation `json:"location,omitempty"`   // Location of the function, none for native functions.
	FunctionName string            `json:"functionName"`         // Name of the function.
	IsGenerator  bool              `json:"isGenerator"`          // Whether this is a generator function.
	ScopeChain   []*DebuggerScope  `json:"scopeChain,omitempty"` // Scope chain for this closure.
}

// Information about the generator object.
type DebuggerGeneratorObjectDetails struct {
	Function     *RuntimeRemoteObject `json:"function"`           // Generator function.
	FunctionName string               `json:"functionName"`       // Name of the generator function.
	Status       string               `json:"status"`             // Current generator object status.
	Location     *DebuggerLocation    `json:"location,omitempty"` // If suspended, location where generator function was suspended (e.g. location of the last 'yield'). Otherwise, location of the generator function.
}

// Collection entry.
type DebuggerCollectionEntry struct {
	Key   *RuntimeRemoteObject `json:"key,omitempty"` // Entry key of a map-like collection, otherwise not provided.
	Value *RuntimeRemoteObject `json:"value"`         // Entry value.
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

// Error data for setScriptSource command. Contains uncompilable script source error.
type DebuggerSetScriptSourceError struct {
	Message      string `json:"message"`      // Compiler error message
	LineNumber   int    `json:"lineNumber"`   // Compile error line number (1-based)
	ColumnNumber int    `json:"columnNumber"` // Compile error column number (1-based)
}

// Search match for resource.
type DebuggerSearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`  // Line number in resource content.
	LineContent string  `json:"lineContent"` // Line with match content.
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
type DebuggerScriptParsedEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                 string `json:"scriptId"`                           // Identifier of the script parsed.
		Url                      string `json:"url"`                                // URL or name of the script parsed (if any).
		StartLine                int    `json:"startLine"`                          // Line offset of the script within the resource with given URL (for script tags).
		StartColumn              int    `json:"startColumn"`                        // Column offset of the script within the resource with given URL.
		EndLine                  int    `json:"endLine"`                            // Last line of the script.
		EndColumn                int    `json:"endColumn"`                          // Length of the last line of the script.
		ExecutionContextId       int    `json:"executionContextId"`                 // Specifies script creation context.
		Hash                     string `json:"hash"`                               // Content hash of the script.
		IsContentScript          bool   `json:"isContentScript,omitempty"`          // Determines whether this script is a user extension script.
		IsInternalScript         bool   `json:"isInternalScript,omitempty"`         // Determines whether this script is an internal script.
		IsLiveEdit               bool   `json:"isLiveEdit,omitempty"`               // True, if this script is generated as a result of the live edit operation.
		SourceMapURL             string `json:"sourceMapURL,omitempty"`             // URL of source map associated with script (if any).
		HasSourceURL             bool   `json:"hasSourceURL,omitempty"`             // True, if this script has sourceURL.
		DeprecatedCommentWasUsed bool   `json:"deprecatedCommentWasUsed,omitempty"` // True, if '//@ sourceURL' or '//@ sourceMappingURL' was used.
	} `json:"Params,omitempty"`
}

// Fired when virtual machine fails to parse the script.
type DebuggerScriptFailedToParseEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                 string `json:"scriptId"`                           // Identifier of the script parsed.
		Url                      string `json:"url"`                                // URL or name of the script parsed (if any).
		StartLine                int    `json:"startLine"`                          // Line offset of the script within the resource with given URL (for script tags).
		StartColumn              int    `json:"startColumn"`                        // Column offset of the script within the resource with given URL.
		EndLine                  int    `json:"endLine"`                            // Last line of the script.
		EndColumn                int    `json:"endColumn"`                          // Length of the last line of the script.
		ExecutionContextId       int    `json:"executionContextId"`                 // Specifies script creation context.
		Hash                     string `json:"hash"`                               // Content hash of the script.
		IsContentScript          bool   `json:"isContentScript,omitempty"`          // Determines whether this script is a user extension script.
		IsInternalScript         bool   `json:"isInternalScript,omitempty"`         // Determines whether this script is an internal script.
		SourceMapURL             string `json:"sourceMapURL,omitempty"`             // URL of source map associated with script (if any).
		HasSourceURL             bool   `json:"hasSourceURL,omitempty"`             // True, if this script has sourceURL.
		DeprecatedCommentWasUsed bool   `json:"deprecatedCommentWasUsed,omitempty"` // True, if '//@ sourceURL' or '//@ sourceMappingURL' was used.
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
// skipped - New value for skip pauses state.
func (c *Debugger) SetSkipAllPauses(skipped bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["skipped"] = skipped
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

// ContinueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// interstatementLocation - Allows breakpoints at the intemediate positions inside statements.
func (c *Debugger) ContinueToLocation(location *DebuggerLocation, interstatementLocation bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["interstatementLocation"] = interstatementLocation
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

// CanSetScriptSource - Always returns true.
// Returns -  result - True if <code>setScriptSource</code> is supported.
func (c *Debugger) CanSetScriptSource() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.canSetScriptSource"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// SetScriptSource - Edits JavaScript source live.
// scriptId - Id of the script to edit.
// scriptSource - New content of the script.
// preview -  If true the change will not actually be applied. Preview mode may be used to get result description without actually modifying the code.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. compileError - Error data if any.
func (c *Debugger) SetScriptSource(scriptId string, scriptSource string, preview bool) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *DebuggerSetScriptSourceError, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["scriptId"] = scriptId
	paramRequest["scriptSource"] = scriptSource
	paramRequest["preview"] = preview
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setScriptSource", Params: paramRequest})
	if err != nil {
		return nil, false, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			StackChanged    bool
			AsyncStackTrace *RuntimeStackTrace
			CompileError    *DebuggerSetScriptSourceError
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

	return chromeData.Result.CallFrames, chromeData.Result.StackChanged, chromeData.Result.AsyncStackTrace, chromeData.Result.CompileError, nil
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

// GetFunctionDetails - Returns detailed information on given function.
// functionId - Id of the function to get details for.
// Returns -  details - Information about the function.
func (c *Debugger) GetFunctionDetails(functionId string) (*DebuggerFunctionDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["functionId"] = functionId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getFunctionDetails", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Details *DebuggerFunctionDetails
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

	return chromeData.Result.Details, nil
}

// GetGeneratorObjectDetails - Returns detailed information on given generator object.
// objectId - Id of the generator object to get details for.
// Returns -  details - Information about the generator object.
func (c *Debugger) GetGeneratorObjectDetails(objectId string) (*DebuggerGeneratorObjectDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getGeneratorObjectDetails", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Details *DebuggerGeneratorObjectDetails
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

	return chromeData.Result.Details, nil
}

// GetCollectionEntries - Returns entries of given collection.
// objectId - Id of the collection to get entries for.
// Returns -  entries - Array of collection entries.
func (c *Debugger) GetCollectionEntries(objectId string) ([]*DebuggerCollectionEntry, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getCollectionEntries", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Entries []*DebuggerCollectionEntry
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

	return chromeData.Result.Entries, nil
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
// doNotPauseOnExceptionsAndMuteConsole - Specifies whether evaluation should stop on exceptions and mute console. Overrides setPauseOnException state.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// Returns -  result - Object wrapper for the evaluation result. wasThrown - True if the result was thrown during the evaluation. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrame(callFrameId string, expression string, objectGroup string, includeCommandLineAPI bool, doNotPauseOnExceptionsAndMuteConsole bool, returnByValue bool, generatePreview bool) (*RuntimeRemoteObject, bool, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["callFrameId"] = callFrameId
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.evaluateOnCallFrame", Params: paramRequest})
	if err != nil {
		return nil, false, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			WasThrown        bool
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, false, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, chromeData.Result.ExceptionDetails, nil
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

// GetBacktrace - Returns call stack including variables changed since VM was paused. VM must be paused.
// Returns -  callFrames - Call stack the virtual machine stopped on. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) GetBacktrace() ([]*DebuggerCallFrame, *RuntimeStackTrace, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getBacktrace"})
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

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (c *Debugger) SetAsyncCallStackDepth(maxDepth int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["maxDepth"] = maxDepth
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncCallStackDepth", Params: paramRequest})
}

// SetBlackboxPatterns - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// patterns - Array of regexps that will be used to check script url for blackbox state.
func (c *Debugger) SetBlackboxPatterns(patterns string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["patterns"] = patterns
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxPatterns", Params: paramRequest})
}

// SetBlackboxedRanges - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
// scriptId - Id of the script.
// positions -
func (c *Debugger) SetBlackboxedRanges(scriptId string, positions *DebuggerScriptPosition) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["scriptId"] = scriptId
	paramRequest["positions"] = positions
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxedRanges", Params: paramRequest})
}
