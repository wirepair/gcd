// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Debugger functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Location in the source code.
type DebuggerLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the `Debugger.scriptParsed`.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}

// Location in the source code.
type DebuggerScriptPosition struct {
	LineNumber   int `json:"lineNumber"`   //
	ColumnNumber int `json:"columnNumber"` //
}

// Location range within one script.
type DebuggerLocationRange struct {
	ScriptId string                  `json:"scriptId"` //
	Start    *DebuggerScriptPosition `json:"start"`    //
	End      *DebuggerScriptPosition `json:"end"`      //
}

// JavaScript call frame. Array of call frames form the call stack.
type DebuggerCallFrame struct {
	CallFrameId      string               `json:"callFrameId"`                // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName     string               `json:"functionName"`               // Name of the JavaScript function called on this call frame.
	FunctionLocation *DebuggerLocation    `json:"functionLocation,omitempty"` // Location in the source code.
	Location         *DebuggerLocation    `json:"location"`                   // Location in the source code.
	Url              string               `json:"url"`                        // JavaScript script name or url. Deprecated in favor of using the `location.scriptId` to resolve the URL via a previously sent `Debugger.scriptParsed` event.
	ScopeChain       []*DebuggerScope     `json:"scopeChain"`                 // Scope chain for this call frame.
	This             *RuntimeRemoteObject `json:"this"`                       // `this` object for this call frame.
	ReturnValue      *RuntimeRemoteObject `json:"returnValue,omitempty"`      // The value being returned, if the function is at return point.
	CanBeRestarted   bool                 `json:"canBeRestarted,omitempty"`   // Valid only while the VM is paused and indicates whether this frame can be restarted or not. Note that a `true` value here does not guarantee that Debugger#restartFrame with this CallFrameId will be successful, but it is very likely.
}

// Scope description.
type DebuggerScope struct {
	Type          string               `json:"type"`                    // Scope type.
	Object        *RuntimeRemoteObject `json:"object"`                  // Object representing the scope. For `global` and `with` scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
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
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the `Debugger.scriptParsed`.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
	Type         string `json:"type,omitempty"`         //
}

// No Description.
type DebuggerWasmDisassemblyChunk struct {
	Lines           []string `json:"lines"`           // The next chunk of disassembled lines.
	BytecodeOffsets []int    `json:"bytecodeOffsets"` // The bytecode offsets describing the start of each line.
}

// Debug symbols available for a wasm script.
type DebuggerDebugSymbols struct {
	Type        string `json:"type"`                  // Type of the debug symbols.
	ExternalURL string `json:"externalURL,omitempty"` // URL of the external symbol source.
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
		CallFrames            []*DebuggerCallFrame   `json:"callFrames"`                      // Call stack the virtual machine stopped on.
		Reason                string                 `json:"reason"`                          // Pause reason.
		Data                  map[string]interface{} `json:"data,omitempty"`                  // Object containing break-specific auxiliary properties.
		HitBreakpoints        []string               `json:"hitBreakpoints,omitempty"`        // Hit breakpoints IDs
		AsyncStackTrace       *RuntimeStackTrace     `json:"asyncStackTrace,omitempty"`       // Async stack trace, if any.
		AsyncStackTraceId     *RuntimeStackTraceId   `json:"asyncStackTraceId,omitempty"`     // Async stack trace, if any.
		AsyncCallStackTraceId *RuntimeStackTraceId   `json:"asyncCallStackTraceId,omitempty"` // Never present, will be removed.
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
		Hash                    string                 `json:"hash"`                              // Content hash of the script, SHA-256.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data likely matching {isDefault: boolean, type: 'default'|'isolated'|'worker', frameId: string}
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
		CodeOffset              int                    `json:"codeOffset,omitempty"`              // If the scriptLanguage is WebAssembly, the code section offset in the module.
		ScriptLanguage          string                 `json:"scriptLanguage,omitempty"`          // The language of the script. enum values: JavaScript, WebAssembly
		EmbedderName            string                 `json:"embedderName,omitempty"`            // The name the embedder supplied for this script.
	} `json:"Params,omitempty"`
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
		Hash                    string                 `json:"hash"`                              // Content hash of the script, SHA-256.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data likely matching {isDefault: boolean, type: 'default'|'isolated'|'worker', frameId: string}
		IsLiveEdit              bool                   `json:"isLiveEdit,omitempty"`              // True, if this script is generated as a result of the live edit operation.
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
		CodeOffset              int                    `json:"codeOffset,omitempty"`              // If the scriptLanguage is WebAssembly, the code section offset in the module.
		ScriptLanguage          string                 `json:"scriptLanguage,omitempty"`          // The language of the script. enum values: JavaScript, WebAssembly
		DebugSymbols            *DebuggerDebugSymbols  `json:"debugSymbols,omitempty"`            // If the scriptLanguage is WebASsembly, the source of debug symbols for the module.
		EmbedderName            string                 `json:"embedderName,omitempty"`            // The name the embedder supplied for this script.
	} `json:"Params,omitempty"`
}

type Debugger struct {
	target gcdmessage.ChromeTargeter
}

func NewDebugger(target gcdmessage.ChromeTargeter) *Debugger {
	c := &Debugger{target: target}
	return c
}

type DebuggerContinueToLocationParams struct {
	// Location to continue to.
	Location *DebuggerLocation `json:"location"`
	//
	TargetCallFrames string `json:"targetCallFrames,omitempty"`
}

// ContinueToLocationWithParams - Continues execution until specific location is reached.
func (c *Debugger) ContinueToLocationWithParams(ctx context.Context, v *DebuggerContinueToLocationParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.continueToLocation", Params: v})
}

// ContinueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// targetCallFrames -
func (c *Debugger) ContinueToLocation(ctx context.Context, location *DebuggerLocation, targetCallFrames string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerContinueToLocationParams
	v.Location = location
	v.TargetCallFrames = targetCallFrames
	return c.ContinueToLocationWithParams(ctx, &v)
}

// Disables debugger for given page.
func (c *Debugger) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disable"})
}

type DebuggerEnableParams struct {
	// The maximum size in bytes of collected scripts (not referenced by other heap objects) the debugger can hold. Puts no limit if parameter is omitted.
	MaxScriptsCacheSize float64 `json:"maxScriptsCacheSize,omitempty"`
}

// EnableWithParams - Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
// Returns -  debuggerId - Unique identifier of the debugger.
func (c *Debugger) EnableWithParams(ctx context.Context, v *DebuggerEnableParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.enable", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			DebuggerId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.DebuggerId, nil
}

// Enable - Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
// maxScriptsCacheSize - The maximum size in bytes of collected scripts (not referenced by other heap objects) the debugger can hold. Puts no limit if parameter is omitted.
// Returns -  debuggerId - Unique identifier of the debugger.
func (c *Debugger) Enable(ctx context.Context, maxScriptsCacheSize float64) (string, error) {
	var v DebuggerEnableParams
	v.MaxScriptsCacheSize = maxScriptsCacheSize
	return c.EnableWithParams(ctx, &v)
}

type DebuggerEvaluateOnCallFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameId string `json:"callFrameId"`
	// Expression to evaluate.
	Expression string `json:"expression"`
	// String object group name to put result into (allows rapid releasing resulting object handles using `releaseObjectGroup`).
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Specifies whether command line API should be available to the evaluated expression, defaults to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
	// Terminate execution after timing out (number of milliseconds).
	Timeout float64 `json:"timeout,omitempty"`
}

// EvaluateOnCallFrameWithParams - Evaluates expression on a given call frame.
// Returns -  result - Object wrapper for the evaluation result. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrameWithParams(ctx context.Context, v *DebuggerEvaluateOnCallFrameParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.evaluateOnCallFrame", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	if chromeData.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// EvaluateOnCallFrame - Evaluates expression on a given call frame.
// callFrameId - Call frame identifier to evaluate on.
// expression - Expression to evaluate.
// objectGroup - String object group name to put result into (allows rapid releasing resulting object handles using `releaseObjectGroup`).
// includeCommandLineAPI - Specifies whether command line API should be available to the evaluated expression, defaults to false.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// throwOnSideEffect - Whether to throw an exception if side effect cannot be ruled out during evaluation.
// timeout - Terminate execution after timing out (number of milliseconds).
// Returns -  result - Object wrapper for the evaluation result. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrame(ctx context.Context, callFrameId string, expression string, objectGroup string, includeCommandLineAPI bool, silent bool, returnByValue bool, generatePreview bool, throwOnSideEffect bool, timeout float64) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v DebuggerEvaluateOnCallFrameParams
	v.CallFrameId = callFrameId
	v.Expression = expression
	v.ObjectGroup = objectGroup
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.Silent = silent
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.ThrowOnSideEffect = throwOnSideEffect
	v.Timeout = timeout
	return c.EvaluateOnCallFrameWithParams(ctx, &v)
}

type DebuggerGetPossibleBreakpointsParams struct {
	// Start of range to search possible breakpoint locations in.
	Start *DebuggerLocation `json:"start"`
	// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	End *DebuggerLocation `json:"end,omitempty"`
	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction bool `json:"restrictToFunction,omitempty"`
}

// GetPossibleBreakpointsWithParams - Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
// Returns -  locations - List of the possible breakpoint locations.
func (c *Debugger) GetPossibleBreakpointsWithParams(ctx context.Context, v *DebuggerGetPossibleBreakpointsParams) ([]*DebuggerBreakLocation, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getPossibleBreakpoints", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Locations []*DebuggerBreakLocation
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Locations, nil
}

// GetPossibleBreakpoints - Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
// start - Start of range to search possible breakpoint locations in.
// end - End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
// restrictToFunction - Only consider locations which are in the same (non-nested) function as start.
// Returns -  locations - List of the possible breakpoint locations.
func (c *Debugger) GetPossibleBreakpoints(ctx context.Context, start *DebuggerLocation, end *DebuggerLocation, restrictToFunction bool) ([]*DebuggerBreakLocation, error) {
	var v DebuggerGetPossibleBreakpointsParams
	v.Start = start
	v.End = end
	v.RestrictToFunction = restrictToFunction
	return c.GetPossibleBreakpointsWithParams(ctx, &v)
}

type DebuggerGetScriptSourceParams struct {
	// Id of the script to get source for.
	ScriptId string `json:"scriptId"`
}

// GetScriptSourceWithParams - Returns source for the script with given id.
// Returns -  scriptSource - Script source (empty in case of Wasm bytecode).
func (c *Debugger) GetScriptSourceWithParams(ctx context.Context, v *DebuggerGetScriptSourceParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getScriptSource", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			ScriptSource string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.ScriptSource, nil
}

// GetScriptSource - Returns source for the script with given id.
// scriptId - Id of the script to get source for.
// Returns -  scriptSource - Script source (empty in case of Wasm bytecode).
func (c *Debugger) GetScriptSource(ctx context.Context, scriptId string) (string, error) {
	var v DebuggerGetScriptSourceParams
	v.ScriptId = scriptId
	return c.GetScriptSourceWithParams(ctx, &v)
}

type DebuggerDisassembleWasmModuleParams struct {
	// Id of the script to disassemble
	ScriptId string `json:"scriptId"`
}

// DisassembleWasmModuleWithParams -
// Returns -  streamId - For large modules, return a stream from which additional chunks of disassembly can be read successively. totalNumberOfLines - The total number of lines in the disassembly text. functionBodyOffsets - The offsets of all function bodies, in the format [start1, end1, start2, end2, ...] where all ends are exclusive. chunk - The first chunk of disassembly.
func (c *Debugger) DisassembleWasmModuleWithParams(ctx context.Context, v *DebuggerDisassembleWasmModuleParams) (string, int, []int, *DebuggerWasmDisassemblyChunk, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disassembleWasmModule", Params: v})
	if err != nil {
		return "", 0, nil, nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			StreamId            string
			TotalNumberOfLines  int
			FunctionBodyOffsets []int
			Chunk               *DebuggerWasmDisassemblyChunk
		}
	}

	if resp == nil {
		return "", 0, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", 0, nil, nil, err
	}

	if chromeData.Error != nil {
		return "", 0, nil, nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.StreamId, chromeData.Result.TotalNumberOfLines, chromeData.Result.FunctionBodyOffsets, chromeData.Result.Chunk, nil
}

// DisassembleWasmModule -
// scriptId - Id of the script to disassemble
// Returns -  streamId - For large modules, return a stream from which additional chunks of disassembly can be read successively. totalNumberOfLines - The total number of lines in the disassembly text. functionBodyOffsets - The offsets of all function bodies, in the format [start1, end1, start2, end2, ...] where all ends are exclusive. chunk - The first chunk of disassembly.
func (c *Debugger) DisassembleWasmModule(ctx context.Context, scriptId string) (string, int, []int, *DebuggerWasmDisassemblyChunk, error) {
	var v DebuggerDisassembleWasmModuleParams
	v.ScriptId = scriptId
	return c.DisassembleWasmModuleWithParams(ctx, &v)
}

type DebuggerNextWasmDisassemblyChunkParams struct {
	//
	StreamId string `json:"streamId"`
}

// NextWasmDisassemblyChunkWithParams - Disassemble the next chunk of lines for the module corresponding to the stream. If disassembly is complete, this API will invalidate the streamId and return an empty chunk. Any subsequent calls for the now invalid stream will return errors.
// Returns -  chunk - The next chunk of disassembly.
func (c *Debugger) NextWasmDisassemblyChunkWithParams(ctx context.Context, v *DebuggerNextWasmDisassemblyChunkParams) (*DebuggerWasmDisassemblyChunk, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.nextWasmDisassemblyChunk", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Chunk *DebuggerWasmDisassemblyChunk
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Chunk, nil
}

// NextWasmDisassemblyChunk - Disassemble the next chunk of lines for the module corresponding to the stream. If disassembly is complete, this API will invalidate the streamId and return an empty chunk. Any subsequent calls for the now invalid stream will return errors.
// streamId -
// Returns -  chunk - The next chunk of disassembly.
func (c *Debugger) NextWasmDisassemblyChunk(ctx context.Context, streamId string) (*DebuggerWasmDisassemblyChunk, error) {
	var v DebuggerNextWasmDisassemblyChunkParams
	v.StreamId = streamId
	return c.NextWasmDisassemblyChunkWithParams(ctx, &v)
}

type DebuggerGetWasmBytecodeParams struct {
	// Id of the Wasm script to get source for.
	ScriptId string `json:"scriptId"`
}

// GetWasmBytecodeWithParams - This command is deprecated. Use getScriptSource instead.
// Returns -
func (c *Debugger) GetWasmBytecodeWithParams(ctx context.Context, v *DebuggerGetWasmBytecodeParams) error {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getWasmBytecode", Params: v})
	if err != nil {
		return err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
		}
	}

	if resp == nil {
		return &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return err
	}

	if chromeData.Error != nil {
		return &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return nil
}

// GetWasmBytecode - This command is deprecated. Use getScriptSource instead.
// scriptId - Id of the Wasm script to get source for.
// Returns -
func (c *Debugger) GetWasmBytecode(ctx context.Context, scriptId string) error {
	var v DebuggerGetWasmBytecodeParams
	v.ScriptId = scriptId
	return c.GetWasmBytecodeWithParams(ctx, &v)
}

type DebuggerGetStackTraceParams struct {
	//
	StackTraceId *RuntimeStackTraceId `json:"stackTraceId"`
}

// GetStackTraceWithParams - Returns stack trace with given `stackTraceId`.
// Returns -  stackTrace -
func (c *Debugger) GetStackTraceWithParams(ctx context.Context, v *DebuggerGetStackTraceParams) (*RuntimeStackTrace, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getStackTrace", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			StackTrace *RuntimeStackTrace
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.StackTrace, nil
}

// GetStackTrace - Returns stack trace with given `stackTraceId`.
// stackTraceId -
// Returns -  stackTrace -
func (c *Debugger) GetStackTrace(ctx context.Context, stackTraceId *RuntimeStackTraceId) (*RuntimeStackTrace, error) {
	var v DebuggerGetStackTraceParams
	v.StackTraceId = stackTraceId
	return c.GetStackTraceWithParams(ctx, &v)
}

// Stops on the next JavaScript statement.
func (c *Debugger) Pause(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.pause"})
}

type DebuggerPauseOnAsyncCallParams struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceId *RuntimeStackTraceId `json:"parentStackTraceId"`
}

// PauseOnAsyncCallWithParams -
func (c *Debugger) PauseOnAsyncCallWithParams(ctx context.Context, v *DebuggerPauseOnAsyncCallParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.pauseOnAsyncCall", Params: v})
}

// PauseOnAsyncCall -
// parentStackTraceId - Debugger will pause when async call with given stack trace is started.
func (c *Debugger) PauseOnAsyncCall(ctx context.Context, parentStackTraceId *RuntimeStackTraceId) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerPauseOnAsyncCallParams
	v.ParentStackTraceId = parentStackTraceId
	return c.PauseOnAsyncCallWithParams(ctx, &v)
}

type DebuggerRemoveBreakpointParams struct {
	//
	BreakpointId string `json:"breakpointId"`
}

// RemoveBreakpointWithParams - Removes JavaScript breakpoint.
func (c *Debugger) RemoveBreakpointWithParams(ctx context.Context, v *DebuggerRemoveBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.removeBreakpoint", Params: v})
}

// RemoveBreakpoint - Removes JavaScript breakpoint.
// breakpointId -
func (c *Debugger) RemoveBreakpoint(ctx context.Context, breakpointId string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerRemoveBreakpointParams
	v.BreakpointId = breakpointId
	return c.RemoveBreakpointWithParams(ctx, &v)
}

type DebuggerRestartFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameId string `json:"callFrameId"`
	// The `mode` parameter must be present and set to 'StepInto', otherwise `restartFrame` will error out.
	Mode string `json:"mode,omitempty"`
}

// RestartFrameWithParams - Restarts particular call frame from the beginning. The old, deprecated behavior of `restartFrame` is to stay paused and allow further CDP commands after a restart was scheduled. This can cause problems with restarting, so we now continue execution immediatly after it has been scheduled until we reach the beginning of the restarted frame.  To stay back-wards compatible, `restartFrame` now expects a `mode` parameter to be present. If the `mode` parameter is missing, `restartFrame` errors out.  The various return values are deprecated and `callFrames` is always empty. Use the call frames from the `Debugger#paused` events instead, that fires once V8 pauses at the beginning of the restarted function.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any. asyncStackTraceId - Async stack trace, if any.
func (c *Debugger) RestartFrameWithParams(ctx context.Context, v *DebuggerRestartFrameParams) ([]*DebuggerCallFrame, *RuntimeStackTrace, *RuntimeStackTraceId, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.restartFrame", Params: v})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			CallFrames        []*DebuggerCallFrame
			AsyncStackTrace   *RuntimeStackTrace
			AsyncStackTraceId *RuntimeStackTraceId
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	if chromeData.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.CallFrames, chromeData.Result.AsyncStackTrace, chromeData.Result.AsyncStackTraceId, nil
}

// RestartFrame - Restarts particular call frame from the beginning. The old, deprecated behavior of `restartFrame` is to stay paused and allow further CDP commands after a restart was scheduled. This can cause problems with restarting, so we now continue execution immediatly after it has been scheduled until we reach the beginning of the restarted frame.  To stay back-wards compatible, `restartFrame` now expects a `mode` parameter to be present. If the `mode` parameter is missing, `restartFrame` errors out.  The various return values are deprecated and `callFrames` is always empty. Use the call frames from the `Debugger#paused` events instead, that fires once V8 pauses at the beginning of the restarted function.
// callFrameId - Call frame identifier to evaluate on.
// mode - The `mode` parameter must be present and set to 'StepInto', otherwise `restartFrame` will error out.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any. asyncStackTraceId - Async stack trace, if any.
func (c *Debugger) RestartFrame(ctx context.Context, callFrameId string, mode string) ([]*DebuggerCallFrame, *RuntimeStackTrace, *RuntimeStackTraceId, error) {
	var v DebuggerRestartFrameParams
	v.CallFrameId = callFrameId
	v.Mode = mode
	return c.RestartFrameWithParams(ctx, &v)
}

type DebuggerResumeParams struct {
	// Set to true to terminate execution upon resuming execution. In contrast to Runtime.terminateExecution, this will allows to execute further JavaScript (i.e. via evaluation) until execution of the paused code is actually resumed, at which point termination is triggered. If execution is currently not paused, this parameter has no effect.
	TerminateOnResume bool `json:"terminateOnResume,omitempty"`
}

// ResumeWithParams - Resumes JavaScript execution.
func (c *Debugger) ResumeWithParams(ctx context.Context, v *DebuggerResumeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.resume", Params: v})
}

// Resume - Resumes JavaScript execution.
// terminateOnResume - Set to true to terminate execution upon resuming execution. In contrast to Runtime.terminateExecution, this will allows to execute further JavaScript (i.e. via evaluation) until execution of the paused code is actually resumed, at which point termination is triggered. If execution is currently not paused, this parameter has no effect.
func (c *Debugger) Resume(ctx context.Context, terminateOnResume bool) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerResumeParams
	v.TerminateOnResume = terminateOnResume
	return c.ResumeWithParams(ctx, &v)
}

type DebuggerSearchInContentParams struct {
	// Id of the script to search in.
	ScriptId string `json:"scriptId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInContentWithParams - Searches for given string in script content.
// Returns -  result - List of search matches.
func (c *Debugger) SearchInContentWithParams(ctx context.Context, v *DebuggerSearchInContentParams) ([]*DebuggerSearchMatch, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.searchInContent", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result []*DebuggerSearchMatch
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// SearchInContent - Searches for given string in script content.
// scriptId - Id of the script to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Debugger) SearchInContent(ctx context.Context, scriptId string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v DebuggerSearchInContentParams
	v.ScriptId = scriptId
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInContentWithParams(ctx, &v)
}

type DebuggerSetAsyncCallStackDepthParams struct {
	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// SetAsyncCallStackDepthWithParams - Enables or disables async call stacks tracking.
func (c *Debugger) SetAsyncCallStackDepthWithParams(ctx context.Context, v *DebuggerSetAsyncCallStackDepthParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncCallStackDepth", Params: v})
}

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
func (c *Debugger) SetAsyncCallStackDepth(ctx context.Context, maxDepth int) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetAsyncCallStackDepthParams
	v.MaxDepth = maxDepth
	return c.SetAsyncCallStackDepthWithParams(ctx, &v)
}

type DebuggerSetBlackboxPatternsParams struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

// SetBlackboxPatternsWithParams - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
func (c *Debugger) SetBlackboxPatternsWithParams(ctx context.Context, v *DebuggerSetBlackboxPatternsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxPatterns", Params: v})
}

// SetBlackboxPatterns - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// patterns - Array of regexps that will be used to check script url for blackbox state.
func (c *Debugger) SetBlackboxPatterns(ctx context.Context, patterns []string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBlackboxPatternsParams
	v.Patterns = patterns
	return c.SetBlackboxPatternsWithParams(ctx, &v)
}

type DebuggerSetBlackboxedRangesParams struct {
	// Id of the script.
	ScriptId string `json:"scriptId"`
	//
	Positions []*DebuggerScriptPosition `json:"positions"`
}

// SetBlackboxedRangesWithParams - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
func (c *Debugger) SetBlackboxedRangesWithParams(ctx context.Context, v *DebuggerSetBlackboxedRangesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxedRanges", Params: v})
}

// SetBlackboxedRanges - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
// scriptId - Id of the script.
// positions -
func (c *Debugger) SetBlackboxedRanges(ctx context.Context, scriptId string, positions []*DebuggerScriptPosition) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBlackboxedRangesParams
	v.ScriptId = scriptId
	v.Positions = positions
	return c.SetBlackboxedRangesWithParams(ctx, &v)
}

type DebuggerSetBreakpointParams struct {
	// Location to set breakpoint in.
	Location *DebuggerLocation `json:"location"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// SetBreakpointWithParams - Sets JavaScript breakpoint at a given location.
// Returns -  breakpointId - Id of the created breakpoint for further reference. actualLocation - Location this breakpoint resolved into.
func (c *Debugger) SetBreakpointWithParams(ctx context.Context, v *DebuggerSetBreakpointParams) (string, *DebuggerLocation, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpoint", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			BreakpointId   string
			ActualLocation *DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	if chromeData.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.BreakpointId, chromeData.Result.ActualLocation, nil
}

// SetBreakpoint - Sets JavaScript breakpoint at a given location.
// location - Location to set breakpoint in.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. actualLocation - Location this breakpoint resolved into.
func (c *Debugger) SetBreakpoint(ctx context.Context, location *DebuggerLocation, condition string) (string, *DebuggerLocation, error) {
	var v DebuggerSetBreakpointParams
	v.Location = location
	v.Condition = condition
	return c.SetBreakpointWithParams(ctx, &v)
}

type DebuggerSetInstrumentationBreakpointParams struct {
	// Instrumentation name.
	Instrumentation string `json:"instrumentation"`
}

// SetInstrumentationBreakpointWithParams - Sets instrumentation breakpoint.
// Returns -  breakpointId - Id of the created breakpoint for further reference.
func (c *Debugger) SetInstrumentationBreakpointWithParams(ctx context.Context, v *DebuggerSetInstrumentationBreakpointParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setInstrumentationBreakpoint", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			BreakpointId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.BreakpointId, nil
}

// SetInstrumentationBreakpoint - Sets instrumentation breakpoint.
// instrumentation - Instrumentation name.
// Returns -  breakpointId - Id of the created breakpoint for further reference.
func (c *Debugger) SetInstrumentationBreakpoint(ctx context.Context, instrumentation string) (string, error) {
	var v DebuggerSetInstrumentationBreakpointParams
	v.Instrumentation = instrumentation
	return c.SetInstrumentationBreakpointWithParams(ctx, &v)
}

type DebuggerSetBreakpointByUrlParams struct {
	// Line number to set breakpoint at.
	LineNumber int `json:"lineNumber"`
	// URL of the resources to set breakpoint on.
	Url string `json:"url,omitempty"`
	// Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or `urlRegex` must be specified.
	UrlRegex string `json:"urlRegex,omitempty"`
	// Script hash of the resources to set breakpoint on.
	ScriptHash string `json:"scriptHash,omitempty"`
	// Offset in the line to set breakpoint at.
	ColumnNumber int `json:"columnNumber,omitempty"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// SetBreakpointByUrlWithParams - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in `locations` property. Further matching script parsing will result in subsequent `breakpointResolved` events issued. This logical breakpoint will survive page reloads.
// Returns -  breakpointId - Id of the created breakpoint for further reference. locations - List of the locations this breakpoint resolved into upon addition.
func (c *Debugger) SetBreakpointByUrlWithParams(ctx context.Context, v *DebuggerSetBreakpointByUrlParams) (string, []*DebuggerLocation, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointByUrl", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			BreakpointId string
			Locations    []*DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	if chromeData.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.BreakpointId, chromeData.Result.Locations, nil
}

// SetBreakpointByUrl - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in `locations` property. Further matching script parsing will result in subsequent `breakpointResolved` events issued. This logical breakpoint will survive page reloads.
// lineNumber - Line number to set breakpoint at.
// url - URL of the resources to set breakpoint on.
// urlRegex - Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or `urlRegex` must be specified.
// scriptHash - Script hash of the resources to set breakpoint on.
// columnNumber - Offset in the line to set breakpoint at.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. locations - List of the locations this breakpoint resolved into upon addition.
func (c *Debugger) SetBreakpointByUrl(ctx context.Context, lineNumber int, url string, urlRegex string, scriptHash string, columnNumber int, condition string) (string, []*DebuggerLocation, error) {
	var v DebuggerSetBreakpointByUrlParams
	v.LineNumber = lineNumber
	v.Url = url
	v.UrlRegex = urlRegex
	v.ScriptHash = scriptHash
	v.ColumnNumber = columnNumber
	v.Condition = condition
	return c.SetBreakpointByUrlWithParams(ctx, &v)
}

type DebuggerSetBreakpointOnFunctionCallParams struct {
	// Function object id.
	ObjectId string `json:"objectId"`
	// Expression to use as a breakpoint condition. When specified, debugger will stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// SetBreakpointOnFunctionCallWithParams - Sets JavaScript breakpoint before each call to the given function. If another function was created from the same source as a given one, calling it will also trigger the breakpoint.
// Returns -  breakpointId - Id of the created breakpoint for further reference.
func (c *Debugger) SetBreakpointOnFunctionCallWithParams(ctx context.Context, v *DebuggerSetBreakpointOnFunctionCallParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointOnFunctionCall", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			BreakpointId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.BreakpointId, nil
}

// SetBreakpointOnFunctionCall - Sets JavaScript breakpoint before each call to the given function. If another function was created from the same source as a given one, calling it will also trigger the breakpoint.
// objectId - Function object id.
// condition - Expression to use as a breakpoint condition. When specified, debugger will stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference.
func (c *Debugger) SetBreakpointOnFunctionCall(ctx context.Context, objectId string, condition string) (string, error) {
	var v DebuggerSetBreakpointOnFunctionCallParams
	v.ObjectId = objectId
	v.Condition = condition
	return c.SetBreakpointOnFunctionCallWithParams(ctx, &v)
}

type DebuggerSetBreakpointsActiveParams struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

// SetBreakpointsActiveWithParams - Activates / deactivates all breakpoints on the page.
func (c *Debugger) SetBreakpointsActiveWithParams(ctx context.Context, v *DebuggerSetBreakpointsActiveParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointsActive", Params: v})
}

// SetBreakpointsActive - Activates / deactivates all breakpoints on the page.
// active - New value for breakpoints active state.
func (c *Debugger) SetBreakpointsActive(ctx context.Context, active bool) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBreakpointsActiveParams
	v.Active = active
	return c.SetBreakpointsActiveWithParams(ctx, &v)
}

type DebuggerSetPauseOnExceptionsParams struct {
	// Pause on exceptions mode.
	State string `json:"state"`
}

// SetPauseOnExceptionsWithParams - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions, or caught exceptions, no exceptions. Initial pause on exceptions state is `none`.
func (c *Debugger) SetPauseOnExceptionsWithParams(ctx context.Context, v *DebuggerSetPauseOnExceptionsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setPauseOnExceptions", Params: v})
}

// SetPauseOnExceptions - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions, or caught exceptions, no exceptions. Initial pause on exceptions state is `none`.
// state - Pause on exceptions mode.
func (c *Debugger) SetPauseOnExceptions(ctx context.Context, state string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetPauseOnExceptionsParams
	v.State = state
	return c.SetPauseOnExceptionsWithParams(ctx, &v)
}

type DebuggerSetReturnValueParams struct {
	// New return value.
	NewValue *RuntimeCallArgument `json:"newValue"`
}

// SetReturnValueWithParams - Changes return value in top frame. Available only at return break position.
func (c *Debugger) SetReturnValueWithParams(ctx context.Context, v *DebuggerSetReturnValueParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setReturnValue", Params: v})
}

// SetReturnValue - Changes return value in top frame. Available only at return break position.
// newValue - New return value.
func (c *Debugger) SetReturnValue(ctx context.Context, newValue *RuntimeCallArgument) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetReturnValueParams
	v.NewValue = newValue
	return c.SetReturnValueWithParams(ctx, &v)
}

type DebuggerSetScriptSourceParams struct {
	// Id of the script to edit.
	ScriptId string `json:"scriptId"`
	// New content of the script.
	ScriptSource string `json:"scriptSource"`
	// If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
	DryRun bool `json:"dryRun,omitempty"`
	// If true, then `scriptSource` is allowed to change the function on top of the stack as long as the top-most stack frame is the only activation of that function.
	AllowTopFrameEditing bool `json:"allowTopFrameEditing,omitempty"`
}

// SetScriptSourceWithParams - Edits JavaScript source live.  In general, functions that are currently on the stack can not be edited with a single exception: If the edited function is the top-most stack frame and that is the only activation of that function on the stack. In this case the live edit will be successful and a `Debugger.restartFrame` for the top-most function is automatically triggered.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. asyncStackTraceId - Async stack trace, if any. status - Whether the operation was successful or not. Only `Ok` denotes a successful live edit while the other enum variants denote why the live edit failed. exceptionDetails - Exception details if any. Only present when `status` is `CompileError`.
func (c *Debugger) SetScriptSourceWithParams(ctx context.Context, v *DebuggerSetScriptSourceParams) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *RuntimeStackTraceId, string, *RuntimeExceptionDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setScriptSource", Params: v})
	if err != nil {
		return nil, false, nil, nil, "", nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			CallFrames        []*DebuggerCallFrame
			StackChanged      bool
			AsyncStackTrace   *RuntimeStackTrace
			AsyncStackTraceId *RuntimeStackTraceId
			Status            string
			ExceptionDetails  *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, false, nil, nil, "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, nil, nil, "", nil, err
	}

	if chromeData.Error != nil {
		return nil, false, nil, nil, "", nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.CallFrames, chromeData.Result.StackChanged, chromeData.Result.AsyncStackTrace, chromeData.Result.AsyncStackTraceId, chromeData.Result.Status, chromeData.Result.ExceptionDetails, nil
}

// SetScriptSource - Edits JavaScript source live.  In general, functions that are currently on the stack can not be edited with a single exception: If the edited function is the top-most stack frame and that is the only activation of that function on the stack. In this case the live edit will be successful and a `Debugger.restartFrame` for the top-most function is automatically triggered.
// scriptId - Id of the script to edit.
// scriptSource - New content of the script.
// dryRun - If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
// allowTopFrameEditing - If true, then `scriptSource` is allowed to change the function on top of the stack as long as the top-most stack frame is the only activation of that function.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. asyncStackTraceId - Async stack trace, if any. status - Whether the operation was successful or not. Only `Ok` denotes a successful live edit while the other enum variants denote why the live edit failed. exceptionDetails - Exception details if any. Only present when `status` is `CompileError`.
func (c *Debugger) SetScriptSource(ctx context.Context, scriptId string, scriptSource string, dryRun bool, allowTopFrameEditing bool) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *RuntimeStackTraceId, string, *RuntimeExceptionDetails, error) {
	var v DebuggerSetScriptSourceParams
	v.ScriptId = scriptId
	v.ScriptSource = scriptSource
	v.DryRun = dryRun
	v.AllowTopFrameEditing = allowTopFrameEditing
	return c.SetScriptSourceWithParams(ctx, &v)
}

type DebuggerSetSkipAllPausesParams struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

// SetSkipAllPausesWithParams - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (c *Debugger) SetSkipAllPausesWithParams(ctx context.Context, v *DebuggerSetSkipAllPausesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setSkipAllPauses", Params: v})
}

// SetSkipAllPauses - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
// skip - New value for skip pauses state.
func (c *Debugger) SetSkipAllPauses(ctx context.Context, skip bool) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetSkipAllPausesParams
	v.Skip = skip
	return c.SetSkipAllPausesWithParams(ctx, &v)
}

type DebuggerSetVariableValueParams struct {
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber int `json:"scopeNumber"`
	// Variable name.
	VariableName string `json:"variableName"`
	// New variable value.
	NewValue *RuntimeCallArgument `json:"newValue"`
	// Id of callframe that holds variable.
	CallFrameId string `json:"callFrameId"`
}

// SetVariableValueWithParams - Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
func (c *Debugger) SetVariableValueWithParams(ctx context.Context, v *DebuggerSetVariableValueParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setVariableValue", Params: v})
}

// SetVariableValue - Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
// scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
// variableName - Variable name.
// newValue - New variable value.
// callFrameId - Id of callframe that holds variable.
func (c *Debugger) SetVariableValue(ctx context.Context, scopeNumber int, variableName string, newValue *RuntimeCallArgument, callFrameId string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetVariableValueParams
	v.ScopeNumber = scopeNumber
	v.VariableName = variableName
	v.NewValue = newValue
	v.CallFrameId = callFrameId
	return c.SetVariableValueWithParams(ctx, &v)
}

type DebuggerStepIntoParams struct {
	// Debugger will pause on the execution of the first async task which was scheduled before next pause.
	BreakOnAsyncCall bool `json:"breakOnAsyncCall,omitempty"`
	// The skipList specifies location ranges that should be skipped on step into.
	SkipList []*DebuggerLocationRange `json:"skipList,omitempty"`
}

// StepIntoWithParams - Steps into the function call.
func (c *Debugger) StepIntoWithParams(ctx context.Context, v *DebuggerStepIntoParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepInto", Params: v})
}

// StepInto - Steps into the function call.
// breakOnAsyncCall - Debugger will pause on the execution of the first async task which was scheduled before next pause.
// skipList - The skipList specifies location ranges that should be skipped on step into.
func (c *Debugger) StepInto(ctx context.Context, breakOnAsyncCall bool, skipList []*DebuggerLocationRange) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerStepIntoParams
	v.BreakOnAsyncCall = breakOnAsyncCall
	v.SkipList = skipList
	return c.StepIntoWithParams(ctx, &v)
}

// Steps out of the function call.
func (c *Debugger) StepOut(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOut"})
}

type DebuggerStepOverParams struct {
	// The skipList specifies location ranges that should be skipped on step over.
	SkipList []*DebuggerLocationRange `json:"skipList,omitempty"`
}

// StepOverWithParams - Steps over the statement.
func (c *Debugger) StepOverWithParams(ctx context.Context, v *DebuggerStepOverParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOver", Params: v})
}

// StepOver - Steps over the statement.
// skipList - The skipList specifies location ranges that should be skipped on step over.
func (c *Debugger) StepOver(ctx context.Context, skipList []*DebuggerLocationRange) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerStepOverParams
	v.SkipList = skipList
	return c.StepOverWithParams(ctx, &v)
}
