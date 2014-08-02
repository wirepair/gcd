// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Timeline types.
// API Version: 1.1
package main

 
// Current values of counters.
type ChromeTimelineCounters struct {
	Documents int `json:"documents,omitempty"` 
	Nodes int `json:"nodes,omitempty"` 
	JsEventListeners int `json:"jsEventListeners,omitempty"` 
	JsHeapSizeUsed int `json:"jsHeapSizeUsed,omitempty"` // Currently used size of JS heap.
	GpuMemoryUsedKB int `json:"gpuMemoryUsedKB,omitempty"` // Current GPU memory usage in kilobytes.
	GpuMemoryLimitKB int `json:"gpuMemoryLimitKB,omitempty"` // Current GPU memory limit in kilobytes.
}
 
 
// Timeline record contains information about the recorded activity.
type ChromeTimelineTimelineEvent struct {
	Type string `json:"type"` // Event type.
	Data object `json:"data"` // Event data.
	StartTime int `json:"startTime"` // Start time.
	EndTime int `json:"endTime,omitempty"` // End time.
	Children []*ChromeTimelineTimelineEvent `json:"children,omitempty"` // Nested records.
	Thread string `json:"thread,omitempty"` // If present, identifies the thread that produced the event.
	StackTrace *ChromeConsoleStackTrace `json:"stackTrace,omitempty"` // Stack trace.
	FrameId string `json:"frameId,omitempty"` // Unique identifier of the frame within the page that the event relates to.
}
 
