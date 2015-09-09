// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Debugger types.
// API Version: 1.1
package types

 
// Breakpoint identifier.
type ChromeDebuggerBreakpointId string 
 
 
// Unique script identifier.
type ChromeDebuggerScriptId string 
 
 
// Call frame identifier.
type ChromeDebuggerCallFrameId string 
 
 
// Location in the source code.
type ChromeDebuggerLocation struct {
	ScriptId *ChromeDebuggerScriptId `json:"scriptId"` // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber int `json:"lineNumber"` // Line number in the script (0-based).
	ColumnNumber int `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}
 
 
// Information about the function.
type ChromeDebuggerFunctionDetails struct {
	Location *ChromeDebuggerLocation `json:"location"` // Location of the function.
	FunctionName string `json:"functionName"` // Name of the function.
	ScopeChain []*ChromeDebuggerScope `json:"scopeChain,omitempty"` // Scope chain for this closure.
}
 
 
// JavaScript call frame. Array of call frames form the call stack.
type ChromeDebuggerCallFrame struct {
	CallFrameId *ChromeDebuggerCallFrameId `json:"callFrameId"` // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName string `json:"functionName"` // Name of the JavaScript function called on this call frame.
	Location *ChromeDebuggerLocation `json:"location"` // Location in the source code.
	ScopeChain []*ChromeDebuggerScope `json:"scopeChain"` // Scope chain for this call frame.
	This *ChromeRuntimeRemoteObject `json:"this"` // <code>this</code> object for this call frame.
	ReturnValue *ChromeRuntimeRemoteObject `json:"returnValue,omitempty"` // The value being returned, if the function is at return point.
}
 
 
// JavaScript call stack, including async stack traces.
type ChromeDebuggerStackTrace struct {
	CallFrames []*ChromeDebuggerCallFrame `json:"callFrames"` // Call frames of the stack trace.
	Description string `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	AsyncStackTrace *ChromeDebuggerStackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
}
 
 
// Scope description.
type ChromeDebuggerScope struct {
	Type string `json:"type"` // Scope type.
	Object *ChromeRuntimeRemoteObject `json:"object"` // Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
}
 
 
// Detailed information on exception (or error) that was thrown during script compilation or execution.
type ChromeDebuggerExceptionDetails struct {
	Text string `json:"text"` // Exception text.
	Url string `json:"url,omitempty"` // URL of the message origin.
	Line int `json:"line,omitempty"` // Line number in the resource that generated this message.
	Column int `json:"column,omitempty"` // Column number in the resource that generated this message.
	StackTrace *ChromeConsoleStackTrace `json:"stackTrace,omitempty"` // JavaScript stack trace for assertions and error messages.
}
 
 
// Error data for setScriptSource command. compileError is a case type for uncompilable script source error.
type ChromeDebuggerSetScriptSourceError struct {
	CompileError map[string]interface{} `json:"compileError,omitempty"` 
}
 
