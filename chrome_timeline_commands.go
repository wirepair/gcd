// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Timeline commands.
// API Version: 1.1

package gcd

import ()

// add this API domain to ChromeTarget
func (c *ChromeTarget) Timeline() *ChromeTimeline {
	if c.timeline == nil {
		c.timeline = newChromeTimeline(c)
	}
	return c.timeline
}

type ChromeTimeline struct {
	target *ChromeTarget
}

func newChromeTimeline(target *ChromeTarget) *ChromeTimeline {
	c := &ChromeTimeline{target: target}
	return c
}

// Deprecated.
func (c *ChromeTimeline) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Timeline.enable"})
}

// Deprecated.
func (c *ChromeTimeline) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Timeline.disable"})
}

// Deprecated.
func (c *ChromeTimeline) Stop() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Timeline.stop"})
}

// start - Deprecated.
// maxCallStackDepth - Samples JavaScript stack traces up to <code>maxCallStackDepth</code>, defaults to 5.
// bufferEvents - Whether instrumentation events should be buffered and returned upon <code>stop</code> call.
// liveEvents - Coma separated event types to issue although bufferEvents is set.
// includeCounters - Whether counters data should be included into timeline events.
// includeGPUEvents - Whether events from GPU process should be collected.
func (c *ChromeTimeline) Start(maxCallStackDepth int, bufferEvents bool, liveEvents string, includeCounters bool, includeGPUEvents bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["maxCallStackDepth"] = maxCallStackDepth
	paramRequest["bufferEvents"] = bufferEvents
	paramRequest["liveEvents"] = liveEvents
	paramRequest["includeCounters"] = includeCounters
	paramRequest["includeGPUEvents"] = includeGPUEvents
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Timeline.start", Params: paramRequest})
}
