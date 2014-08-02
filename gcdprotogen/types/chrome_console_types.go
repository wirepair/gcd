// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Console types.
// API Version: 1.1
package types

 
// Number of seconds since epoch.
type ChromeConsoleTimestamp float 
 
 
// Console message.
type ChromeConsoleConsoleMessage struct {
	Source string `json:"source"` // Message source.
	Level string `json:"level"` // Message severity.
	Text string `json:"text"` // Message text.
	Type string `json:"type,omitempty"` // Console message type.
	Url string `json:"url,omitempty"` // URL of the message origin.
	Line int `json:"line,omitempty"` // Line number in the resource that generated this message.
	Column int `json:"column,omitempty"` // Column number in the resource that generated this message.
	RepeatCount int `json:"repeatCount,omitempty"` // Repeat count for repeated messages.
	Parameters []*ChromeRuntimeRemoteObject `json:"parameters,omitempty"` // Message parameters in case of the formatted message.
	StackTrace *ChromeConsoleStackTrace `json:"stackTrace,omitempty"` // JavaScript stack trace for assertions and error messages.
	AsyncStackTrace *ChromeConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Asynchronous JavaScript stack trace that preceded this message, if available.
	NetworkRequestId *ChromeNetworkRequestId `json:"networkRequestId,omitempty"` // Identifier of the network request associated with this message.
	Timestamp *ChromeConsoleTimestamp `json:"timestamp"` // Timestamp, when this message was fired.
	ExecutionContextId *ChromeRuntimeExecutionContextId `json:"executionContextId,omitempty"` // Identifier of the context where this message was created
}
 
 
// Stack entry for console errors and assertions.
type ChromeConsoleCallFrame struct {
	FunctionName string `json:"functionName"` // JavaScript function name.
	ScriptId string `json:"scriptId"` // JavaScript script id.
	Url string `json:"url"` // JavaScript script name or url.
	LineNumber int `json:"lineNumber"` // JavaScript script line number.
	ColumnNumber int `json:"columnNumber"` // JavaScript script column number.
}
 
 
// Call frames for assertions or error messages.
type ChromeConsoleStackTrace []*ChromeConsoleCallFrame 
 
 
// Asynchronous JavaScript call stack.
type ChromeConsoleAsyncStackTrace struct {
	CallFrames []*ChromeConsoleCallFrame `json:"callFrames"` // Call frames of the stack trace.
	Description string `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	AsyncStackTrace *ChromeConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Next asynchronous stack trace, if any.
}
 
