// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Debugger functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

//
type DebuggerSubCompileError struct {
	Message      string `json:"message"`      // Compiler error message
	LineNumber   int    `json:"lineNumber"`   // Compile error line number (1-based)
	ColumnNumber int    `json:"columnNumber"` // Compile error column number (1-based)
}

// Location in the source code.
type DebuggerLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
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

// JavaScript call stack, including async stack traces.
type DebuggerStackTrace struct {
	CallFrames      []*DebuggerCallFrame `json:"callFrames"`                // Call frames of the stack trace.
	Description     string               `json:"description,omitempty"`     // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	AsyncStackTrace *DebuggerStackTrace  `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
}

// Scope description.
type DebuggerScope struct {
	Type   string               `json:"type"`   // Scope type.
	Object *RuntimeRemoteObject `json:"object"` // Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
}

// Detailed information on exception (or error) that was thrown during script compilation or execution.
type DebuggerExceptionDetails struct {
	Text       string              `json:"text"`                 // Exception text.
	Url        string              `json:"url,omitempty"`        // URL of the message origin.
	ScriptId   string              `json:"scriptId,omitempty"`   // Script ID of the message origin.
	Line       int                 `json:"line,omitempty"`       // Line number in the resource that generated this message.
	Column     int                 `json:"column,omitempty"`     // Column number in the resource that generated this message.
	StackTrace []*ConsoleCallFrame `json:"stackTrace,omitempty"` // JavaScript stack trace for assertions and error messages.
}

// Error data for setScriptSource command. compileError is a case type for uncompilable script source error.
type DebuggerSetScriptSourceError struct {
	CompileError DebuggerSubCompileError `json:"compileError,omitempty"` //
}

// Information about the promise. All fields but id are optional and if present they reflect the new state of the property on the promise with given id.
type DebuggerPromiseDetails struct {
	Id                   int                     `json:"id"`                             // Unique id of the promise.
	Status               string                  `json:"status,omitempty"`               // Status of the promise.
	ParentId             int                     `json:"parentId,omitempty"`             // Id of the parent promise.
	CallFrame            *ConsoleCallFrame       `json:"callFrame,omitempty"`            // Top call frame on promise creation.
	CreationTime         float64                 `json:"creationTime,omitempty"`         // Creation time of the promise.
	SettlementTime       float64                 `json:"settlementTime,omitempty"`       // Settlement time of the promise.
	CreationStack        []*ConsoleCallFrame     `json:"creationStack,omitempty"`        // JavaScript stack trace on promise creation.
	AsyncCreationStack   *ConsoleAsyncStackTrace `json:"asyncCreationStack,omitempty"`   // JavaScript asynchronous stack trace on promise creation, if available.
	SettlementStack      []*ConsoleCallFrame     `json:"settlementStack,omitempty"`      // JavaScript stack trace on promise settlement.
	AsyncSettlementStack *ConsoleAsyncStackTrace `json:"asyncSettlementStack,omitempty"` // JavaScript asynchronous stack trace on promise settlement, if available.
}

// Information about the async operation.
type DebuggerAsyncOperation struct {
	Id              int                     `json:"id"`                        // Unique id of the async operation.
	Description     string                  `json:"description"`               // String description of the async operation.
	StackTrace      []*ConsoleCallFrame     `json:"stackTrace,omitempty"`      // Stack trace where async operation was scheduled.
	AsyncStackTrace *ConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Asynchronous stack trace where async operation was scheduled, if available.
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
		ScriptId         string `json:"scriptId"`                   // Identifier of the script parsed.
		Url              string `json:"url"`                        // URL or name of the script parsed (if any).
		StartLine        int    `json:"startLine"`                  // Line offset of the script within the resource with given URL (for script tags).
		StartColumn      int    `json:"startColumn"`                // Column offset of the script within the resource with given URL.
		EndLine          int    `json:"endLine"`                    // Last line of the script.
		EndColumn        int    `json:"endColumn"`                  // Length of the last line of the script.
		IsContentScript  bool   `json:"isContentScript,omitempty"`  // Determines whether this script is a user extension script.
		IsInternalScript bool   `json:"isInternalScript,omitempty"` // Determines whether this script is an internal script.
		SourceMapURL     string `json:"sourceMapURL,omitempty"`     // URL of source map associated with script (if any).
		HasSourceURL     bool   `json:"hasSourceURL,omitempty"`     // True, if this script has sourceURL.
	} `json:"Params,omitempty"`
}

// Fired when virtual machine fails to parse the script.
type DebuggerScriptFailedToParseEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId         string `json:"scriptId"`                   // Identifier of the script parsed.
		Url              string `json:"url"`                        // URL or name of the script parsed (if any).
		StartLine        int    `json:"startLine"`                  // Line offset of the script within the resource with given URL (for script tags).
		StartColumn      int    `json:"startColumn"`                // Column offset of the script within the resource with given URL.
		EndLine          int    `json:"endLine"`                    // Last line of the script.
		EndColumn        int    `json:"endColumn"`                  // Length of the last line of the script.
		IsContentScript  bool   `json:"isContentScript,omitempty"`  // Determines whether this script is a user extension script.
		IsInternalScript bool   `json:"isInternalScript,omitempty"` // Determines whether this script is an internal script.
		SourceMapURL     string `json:"sourceMapURL,omitempty"`     // URL of source map associated with script (if any).
		HasSourceURL     bool   `json:"hasSourceURL,omitempty"`     // True, if this script has sourceURL.
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
		AsyncStackTrace *DebuggerStackTrace    `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	} `json:"Params,omitempty"`
}

// Fired when a <code>Promise</code> is created, updated or garbage collected.
type DebuggerPromiseUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		EventType string                  `json:"eventType"` // Type of the event.
		Promise   *DebuggerPromiseDetails `json:"promise"`   // Information about the updated <code>Promise</code>.
	} `json:"Params,omitempty"`
}

// Fired when an async operation is scheduled (while in a debugger stepping session).
type DebuggerAsyncOperationStartedEvent struct {
	Method string `json:"method"`
	Params struct {
		Operation *DebuggerAsyncOperation `json:"operation"` // Information about the async operation.
	} `json:"Params,omitempty"`
}

// Fired when an async operation is completed (while in a debugger stepping session).
type DebuggerAsyncOperationCompletedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id int `json:"id"` // ID of the async operation that was completed.
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
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.enable"})
}

// Disables debugger for given page.
func (c *Debugger) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disable"})
}

// SetBreakpointsActive - Activates / deactivates all breakpoints on the page.
// active - New value for breakpoints active state.
func (c *Debugger) SetBreakpointsActive(active bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["active"] = active
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointsActive", Params: paramRequest})
}

// SetSkipAllPauses - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
// skipped - New value for skip pauses state.
func (c *Debugger) SetSkipAllPauses(skipped bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["skipped"] = skipped
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setSkipAllPauses", Params: paramRequest})
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
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointByUrl", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			BreakpointId string
			Locations    []*DebuggerLocation
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
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
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpoint", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			BreakpointId   string
			ActualLocation *DebuggerLocation
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", nil, err
	}

	return chromeData.Result.BreakpointId, chromeData.Result.ActualLocation, nil
}

// RemoveBreakpoint - Removes JavaScript breakpoint.
// breakpointId -
func (c *Debugger) RemoveBreakpoint(breakpointId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["breakpointId"] = breakpointId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.removeBreakpoint", Params: paramRequest})
}

// ContinueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// interstatementLocation - Allows breakpoints at the intemediate positions inside statements.
func (c *Debugger) ContinueToLocation(location *DebuggerLocation, interstatementLocation bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["location"] = location
	paramRequest["interstatementLocation"] = interstatementLocation
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.continueToLocation", Params: paramRequest})
}

// Steps over the statement.
func (c *Debugger) StepOver() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOver"})
}

// Steps into the function call.
func (c *Debugger) StepInto() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepInto"})
}

// Steps out of the function call.
func (c *Debugger) StepOut() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOut"})
}

// Stops on the next JavaScript statement.
func (c *Debugger) Pause() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.pause"})
}

// Resumes JavaScript execution.
func (c *Debugger) Resume() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.resume"})
}

// Steps into the first async operation handler that was scheduled by or after the current statement.
func (c *Debugger) StepIntoAsync() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepIntoAsync"})
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
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.searchInContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
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

	return chromeData.Result.Result, nil
}

// CanSetScriptSource - Always returns true.
// Returns -  result - True if <code>setScriptSource</code> is supported.
func (c *Debugger) CanSetScriptSource() (bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.canSetScriptSource"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// SetScriptSource - Edits JavaScript source live.
// scriptId - Id of the script to edit.
// scriptSource - New content of the script.
// preview -  If true the change will not actually be applied. Preview mode may be used to get result description without actually modifying the code.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) SetScriptSource(scriptId string, scriptSource string, preview bool) ([]*DebuggerCallFrame, bool, *DebuggerStackTrace, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["scriptId"] = scriptId
	paramRequest["scriptSource"] = scriptSource
	paramRequest["preview"] = preview
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setScriptSource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			StackChanged    bool
			AsyncStackTrace *DebuggerStackTrace
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, false, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.StackChanged, chromeData.Result.AsyncStackTrace, nil
}

// RestartFrame - Restarts particular call frame from the beginning.
// callFrameId - Call frame identifier to evaluate on.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) RestartFrame(callFrameId string) ([]*DebuggerCallFrame, *DebuggerStackTrace, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["callFrameId"] = callFrameId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.restartFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			AsyncStackTrace *DebuggerStackTrace
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
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
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getScriptSource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ScriptSource string
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
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
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getFunctionDetails", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Details *DebuggerFunctionDetails
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

	return chromeData.Result.Details, nil
}

// GetGeneratorObjectDetails - Returns detailed information on given generator object.
// objectId - Id of the generator object to get details for.
// Returns -  details - Information about the generator object.
func (c *Debugger) GetGeneratorObjectDetails(objectId string) (*DebuggerGeneratorObjectDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getGeneratorObjectDetails", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Details *DebuggerGeneratorObjectDetails
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

	return chromeData.Result.Details, nil
}

// GetCollectionEntries - Returns entries of given collection.
// objectId - Id of the collection to get entries for.
// Returns -  entries - Array of collection entries.
func (c *Debugger) GetCollectionEntries(objectId string) ([]*DebuggerCollectionEntry, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getCollectionEntries", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Entries []*DebuggerCollectionEntry
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

	return chromeData.Result.Entries, nil
}

// SetPauseOnExceptions - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
// state - Pause on exceptions mode.
func (c *Debugger) SetPauseOnExceptions(state string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["state"] = state
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setPauseOnExceptions", Params: paramRequest})
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
func (c *Debugger) EvaluateOnCallFrame(callFrameId string, expression string, objectGroup string, includeCommandLineAPI bool, doNotPauseOnExceptionsAndMuteConsole bool, returnByValue bool, generatePreview bool) (*RuntimeRemoteObject, bool, *DebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["callFrameId"] = callFrameId
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.evaluateOnCallFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			WasThrown        bool
			ExceptionDetails *DebuggerExceptionDetails
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, false, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, chromeData.Result.ExceptionDetails, nil
}

// CompileScript - Compiles expression.
// expression - Expression to compile.
// sourceURL - Source url to be set for the script.
// persistScript - Specifies whether the compiled script should be persisted.
// executionContextId - Specifies in which isolated context to perform script run. Each content script lives in an isolated context and this parameter may be used to specify one of those contexts. If the parameter is omitted or 0 the evaluation will be performed in the context of the inspected page.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Debugger) CompileScript(expression string, sourceURL string, persistScript bool, executionContextId int) (string, *DebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["expression"] = expression
	paramRequest["sourceURL"] = sourceURL
	paramRequest["persistScript"] = persistScript
	paramRequest["executionContextId"] = executionContextId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.compileScript", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ScriptId         string
			ExceptionDetails *DebuggerExceptionDetails
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", nil, err
	}

	return chromeData.Result.ScriptId, chromeData.Result.ExceptionDetails, nil
}

// RunScript - Runs script with given id in a given context.
// scriptId - Id of the script to run.
// executionContextId - Specifies in which isolated context to perform script run. Each content script lives in an isolated context and this parameter may be used to specify one of those contexts. If the parameter is omitted or 0 the evaluation will be performed in the context of the inspected page.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// doNotPauseOnExceptionsAndMuteConsole - Specifies whether script run should stop on exceptions and mute console. Overrides setPauseOnException state.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Debugger) RunScript(scriptId string, executionContextId int, objectGroup string, doNotPauseOnExceptionsAndMuteConsole bool) (*RuntimeRemoteObject, *DebuggerExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["scriptId"] = scriptId
	paramRequest["executionContextId"] = executionContextId
	paramRequest["objectGroup"] = objectGroup
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.runScript", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *DebuggerExceptionDetails
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// SetVariableValue - Changes value of variable in a callframe or a closure. Either callframe or function must be specified. Object-based scopes are not supported and must be mutated manually.
// scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
// variableName - Variable name.
// newValue - New variable value.
// callFrameId - Id of callframe that holds variable.
// functionObjectId - Object id of closure (function) that holds variable.
func (c *Debugger) SetVariableValue(scopeNumber int, variableName string, newValue *RuntimeCallArgument, callFrameId string, functionObjectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["scopeNumber"] = scopeNumber
	paramRequest["variableName"] = variableName
	paramRequest["newValue"] = newValue
	paramRequest["callFrameId"] = callFrameId
	paramRequest["functionObjectId"] = functionObjectId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setVariableValue", Params: paramRequest})
}

// GetStepInPositions - Lists all positions where step-in is possible for a current statement in a specified call frame
// callFrameId - Id of a call frame where the current statement should be analized
// Returns -  stepInPositions - experimental
func (c *Debugger) GetStepInPositions(callFrameId string) ([]*DebuggerLocation, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["callFrameId"] = callFrameId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getStepInPositions", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			StepInPositions []*DebuggerLocation
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

	return chromeData.Result.StepInPositions, nil
}

// GetBacktrace - Returns call stack including variables changed since VM was paused. VM must be paused.
// Returns -  callFrames - Call stack the virtual machine stopped on. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) GetBacktrace() ([]*DebuggerCallFrame, *DebuggerStackTrace, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getBacktrace"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			AsyncStackTrace *DebuggerStackTrace
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.AsyncStackTrace, nil
}

// SkipStackFrames - Makes backend skip steps in the sources with names matching given pattern. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// script - Regular expression defining the scripts to ignore while stepping.
// skipContentScripts - True, if all content scripts should be ignored.
func (c *Debugger) SkipStackFrames(script string, skipContentScripts bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["script"] = script
	paramRequest["skipContentScripts"] = skipContentScripts
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.skipStackFrames", Params: paramRequest})
}

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (c *Debugger) SetAsyncCallStackDepth(maxDepth int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["maxDepth"] = maxDepth
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncCallStackDepth", Params: paramRequest})
}

// EnablePromiseTracker - Enables promise tracking, information about <code>Promise</code>s created or updated will now be stored on the backend.
// captureStacks - Whether to capture stack traces for promise creation and settlement events (default: false).
func (c *Debugger) EnablePromiseTracker(captureStacks bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["captureStacks"] = captureStacks
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.enablePromiseTracker", Params: paramRequest})
}

// Disables promise tracking.
func (c *Debugger) DisablePromiseTracker() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disablePromiseTracker"})
}

// GetPromiseById - Returns <code>Promise</code> with specified ID.
// promiseId -
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  promise - Object wrapper for <code>Promise</code> with specified ID, if any.
func (c *Debugger) GetPromiseById(promiseId int, objectGroup string) (*RuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["promiseId"] = promiseId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getPromiseById", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Promise *RuntimeRemoteObject
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

	return chromeData.Result.Promise, nil
}

// Fires pending <code>asyncOperationStarted</code> events (if any), as if a debugger stepping session has just been started.
func (c *Debugger) FlushAsyncOperationEvents() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.flushAsyncOperationEvents"})
}

// SetAsyncOperationBreakpoint - Sets breakpoint on AsyncOperation callback handler.
// operationId - ID of the async operation to set breakpoint for.
func (c *Debugger) SetAsyncOperationBreakpoint(operationId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["operationId"] = operationId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncOperationBreakpoint", Params: paramRequest})
}

// RemoveAsyncOperationBreakpoint - Removes AsyncOperation breakpoint.
// operationId - ID of the async operation to remove breakpoint for.
func (c *Debugger) RemoveAsyncOperationBreakpoint(operationId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["operationId"] = operationId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.removeAsyncOperationBreakpoint", Params: paramRequest})
}
