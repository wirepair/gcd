// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Profiler types.
// API Version: 1.1
package types

 
// CPU Profile node. Holds callsite information, execution statistics and child nodes.
type ChromeProfilerCPUProfileNode struct {
	FunctionName string `json:"functionName"` // Function name.
	ScriptId *ChromeDebuggerScriptId `json:"scriptId"` // Script identifier.
	Url string `json:"url"` // URL.
	LineNumber int `json:"lineNumber"` // 1-based line number of the function start position.
	ColumnNumber int `json:"columnNumber"` // 1-based column number of the function start position.
	HitCount int `json:"hitCount"` // Number of samples where this node was on top of the call stack.
	CallUID float64 `json:"callUID"` // Call UID.
	Children []*ChromeProfilerCPUProfileNode `json:"children"` // Child nodes.
	DeoptReason string `json:"deoptReason"` // The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	Id int `json:"id"` // Unique id of the node.
}
 
 
// Profile.
type ChromeProfilerCPUProfile struct {
	Head *ChromeProfilerCPUProfileNode `json:"head"` 
	StartTime float64 `json:"startTime"` // Profiling start time in seconds.
	EndTime float64 `json:"endTime"` // Profiling end time in seconds.
	Samples []int `json:"samples,omitempty"` // Ids of samples top nodes.
	Timestamps []float64 `json:"timestamps,omitempty"` // Timestamps of the samples in microseconds.
}
 
