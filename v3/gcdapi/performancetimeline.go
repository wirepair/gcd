// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains PerformanceTimeline functionality.
// API Version: 1.3

package gcdapi

import (
	"context"

	"github.com/wirepair/gcd/v3/gcdmessage"
)

// See https://github.com/WICG/LargestContentfulPaint and largest_contentful_paint.idl
type PerformanceTimelineLargestContentfulPaint struct {
	RenderTime float64 `json:"renderTime"`          //
	LoadTime   float64 `json:"loadTime"`            //
	Size       float64 `json:"size"`                // The number of pixels being painted.
	ElementId  string  `json:"elementId,omitempty"` // The id attribute of the element, if available.
	Url        string  `json:"url,omitempty"`       // The URL of the image (may be trimmed).
	NodeId     int     `json:"nodeId,omitempty"`    //
}

// No Description.
type PerformanceTimelineLayoutShiftAttribution struct {
	PreviousRect *DOMRect `json:"previousRect"`     //
	CurrentRect  *DOMRect `json:"currentRect"`      //
	NodeId       int      `json:"nodeId,omitempty"` //
}

// See https://wicg.github.io/layout-instability/#sec-layout-shift and layout_shift.idl
type PerformanceTimelineLayoutShift struct {
	Value          float64                                      `json:"value"`          // Score increment produced by this event.
	HadRecentInput bool                                         `json:"hadRecentInput"` //
	LastInputTime  float64                                      `json:"lastInputTime"`  //
	Sources        []*PerformanceTimelineLayoutShiftAttribution `json:"sources"`        //
}

// No Description.
type PerformanceTimelineTimelineEvent struct {
	FrameId            string                                     `json:"frameId"`                      // Identifies the frame that this event is related to. Empty for non-frame targets.
	Type               string                                     `json:"type"`                         // The event type, as specified in https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype This determines which of the optional "details" fiedls is present.
	Name               string                                     `json:"name"`                         // Name may be empty depending on the type.
	Time               float64                                    `json:"time"`                         // Time in seconds since Epoch, monotonically increasing within document lifetime.
	Duration           float64                                    `json:"duration,omitempty"`           // Event duration, if applicable.
	LcpDetails         *PerformanceTimelineLargestContentfulPaint `json:"lcpDetails,omitempty"`         //
	LayoutShiftDetails *PerformanceTimelineLayoutShift            `json:"layoutShiftDetails,omitempty"` //
}

// Sent when a performance timeline event is added. See reportPerformanceTimeline method.
type PerformanceTimelineTimelineEventAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Event *PerformanceTimelineTimelineEvent `json:"event"` //
	} `json:"Params,omitempty"`
}

type PerformanceTimeline struct {
	target gcdmessage.ChromeTargeter
}

func NewPerformanceTimeline(target gcdmessage.ChromeTargeter) *PerformanceTimeline {
	c := &PerformanceTimeline{target: target}
	return c
}

type PerformanceTimelineEnableParams struct {
	// The types of event to report, as specified in https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype The specified filter overrides any previous filters, passing empty filter disables recording. Note that not all types exposed to the web platform are currently supported.
	EventTypes []string `json:"eventTypes"`
}

// EnableWithParams - Previously buffered events would be reported before method returns. See also: timelineEventAdded
func (c *PerformanceTimeline) EnableWithParams(ctx context.Context, v *PerformanceTimelineEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PerformanceTimeline.enable", Params: v})
}

// Enable - Previously buffered events would be reported before method returns. See also: timelineEventAdded
// eventTypes - The types of event to report, as specified in https://w3c.github.io/performance-timeline/#dom-performanceentry-entrytype The specified filter overrides any previous filters, passing empty filter disables recording. Note that not all types exposed to the web platform are currently supported.
func (c *PerformanceTimeline) Enable(ctx context.Context, eventTypes []string) (*gcdmessage.ChromeResponse, error) {
	var v PerformanceTimelineEnableParams
	v.EventTypes = eventTypes
	return c.EnableWithParams(ctx, &v)
}
