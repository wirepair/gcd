// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Timeline functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Timeline record contains information about the recorded activity.
type TimelineTimelineEvent struct {
	Type       string                   `json:"type"`                 // Event type.
	Data       map[string]interface{}   `json:"data"`                 // Event data.
	StartTime  float64                  `json:"startTime"`            // Start time.
	EndTime    float64                  `json:"endTime,omitempty"`    // End time.
	Children   []*TimelineTimelineEvent `json:"children,omitempty"`   // Nested records.
	Thread     string                   `json:"thread,omitempty"`     // If present, identifies the thread that produced the event.
	StackTrace []*ConsoleCallFrame      `json:"stackTrace,omitempty"` // Stack trace.
	FrameId    string                   `json:"frameId,omitempty"`    // Unique identifier of the frame within the page that the event relates to.
}

// Deprecated.
type TimelineEventRecordedEvent struct {
	Record *TimelineTimelineEvent `json:"record"` // Timeline event record data.
}

type Timeline struct {
	target gcdmessage.ChromeTargeter
}

func NewTimeline(target gcdmessage.ChromeTargeter) *Timeline {
	c := &Timeline{target: target}
	return c
}

// Deprecated.
func (c *Timeline) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Timeline.enable"})
}

// Deprecated.
func (c *Timeline) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Timeline.disable"})
}

// Start - Deprecated.
// maxCallStackDepth - Samples JavaScript stack traces up to <code>maxCallStackDepth</code>, defaults to 5.
// bufferEvents - Whether instrumentation events should be buffered and returned upon <code>stop</code> call.
// liveEvents - Coma separated event types to issue although bufferEvents is set.
// includeCounters - Whether counters data should be included into timeline events.
// includeGPUEvents - Whether events from GPU process should be collected.
func (c *Timeline) Start(maxCallStackDepth int, bufferEvents bool, liveEvents string, includeCounters bool, includeGPUEvents bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["maxCallStackDepth"] = maxCallStackDepth
	paramRequest["bufferEvents"] = bufferEvents
	paramRequest["liveEvents"] = liveEvents
	paramRequest["includeCounters"] = includeCounters
	paramRequest["includeGPUEvents"] = includeGPUEvents
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Timeline.start", Params: paramRequest})
}

// Deprecated.
func (c *Timeline) Stop() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Timeline.stop"})
}
