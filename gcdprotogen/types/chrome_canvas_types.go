// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Canvas types.
// API Version: 1.1
package main

 
// Unique resource identifier.
type ChromeCanvasResourceId string 
 
 
// Resource state descriptor.
type ChromeCanvasResourceStateDescriptor struct {
	Name string `json:"name"` // State name.
	EnumValueForName string `json:"enumValueForName,omitempty"` // String representation of the enum value, if <code>name</code> stands for an enum.
	Value *ChromeCanvasCallArgument `json:"value,omitempty"` // The value associated with the particular state.
	Values []*ChromeCanvasResourceStateDescriptor `json:"values,omitempty"` // Array of values associated with the particular state. Either <code>value</code> or <code>values</code> will be specified.
	IsArray bool `json:"isArray,omitempty"` // True iff the given <code>values</code> items stand for an array rather than a list of grouped states.
}
 
 
// Resource state.
type ChromeCanvasResourceState struct {
	Id *ChromeCanvasResourceId `json:"id"` 
	TraceLogId *ChromeCanvasTraceLogId `json:"traceLogId"` 
	Descriptors []*ChromeCanvasResourceStateDescriptor `json:"descriptors,omitempty"` // Describes current <code>Resource</code> state.
	ImageURL string `json:"imageURL,omitempty"` // Screenshot image data URL.
}
 
 
// 
type ChromeCanvasCallArgument struct {
	Description string `json:"description"` // String representation of the object.
	EnumName string `json:"enumName,omitempty"` // Enum name, if any, that stands for the value (for example, a WebGL enum name).
	ResourceId *ChromeCanvasResourceId `json:"resourceId,omitempty"` // Resource identifier. Specified for <code>Resource</code> objects only.
	Type string `json:"type,omitempty"` // Object type. Specified for non <code>Resource</code> objects only.
	Subtype string `json:"subtype,omitempty"` // Object subtype hint. Specified for <code>object</code> type values only.
	RemoteObject *ChromeRuntimeRemoteObject `json:"remoteObject,omitempty"` // The <code>RemoteObject</code>, if requested.
}
 
 
// 
type ChromeCanvasCall struct {
	ContextId *ChromeCanvasResourceId `json:"contextId"` 
	FunctionName string `json:"functionName,omitempty"` 
	Arguments []*ChromeCanvasCallArgument `json:"arguments,omitempty"` 
	Result *ChromeCanvasCallArgument `json:"result,omitempty"` 
	IsDrawingCall bool `json:"isDrawingCall,omitempty"` 
	IsFrameEndCall bool `json:"isFrameEndCall,omitempty"` 
	Property string `json:"property,omitempty"` 
	Value *ChromeCanvasCallArgument `json:"value,omitempty"` 
	SourceURL string `json:"sourceURL,omitempty"` 
	LineNumber int `json:"lineNumber,omitempty"` 
	ColumnNumber int `json:"columnNumber,omitempty"` 
}
 
 
// Unique trace log identifier.
type ChromeCanvasTraceLogId string 
 
 
// 
type ChromeCanvasTraceLog struct {
	Id *ChromeCanvasTraceLogId `json:"id"` 
	Calls []*ChromeCanvasCall `json:"calls"` 
	Contexts []*ChromeCanvasCallArgument `json:"contexts"` 
	StartOffset int `json:"startOffset"` 
	Alive bool `json:"alive"` 
	TotalAvailableCalls int `json:"totalAvailableCalls"` 
}
 
