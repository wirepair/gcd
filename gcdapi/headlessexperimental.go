// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeadlessExperimental functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Encoding options for a screenshot.
type HeadlessExperimentalScreenshotParams struct {
	Format  string `json:"format,omitempty"`  // Image compression format (defaults to png).
	Quality int    `json:"quality,omitempty"` // Compression quality from range [0..100] (jpeg only).
}

// Issued when the target starts or stops needing BeginFrames.
type HeadlessExperimentalNeedsBeginFramesChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		NeedsBeginFrames bool `json:"needsBeginFrames"` // True if BeginFrames are needed, false otherwise.
	} `json:"Params,omitempty"`
}

type HeadlessExperimental struct {
	target gcdmessage.ChromeTargeter
}

func NewHeadlessExperimental(target gcdmessage.ChromeTargeter) *HeadlessExperimental {
	c := &HeadlessExperimental{target: target}
	return c
}

type HeadlessExperimentalBeginFrameParams struct {
	// Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used.
	FrameTime float64 `json:"frameTime,omitempty"`
	// Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval.
	Deadline float64 `json:"deadline,omitempty"`
	// The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval float64 `json:"interval,omitempty"`
	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured.
	Screenshot *HeadlessExperimentalScreenshotParams `json:"screenshot,omitempty"`
}

// BeginFrameWithParams - Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl.
// Returns -  hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. mainFrameContentUpdated - Whether the main frame submitted a new display frame in response to this BeginFrame. screenshotData - Base64-encoded image data of the screenshot, if one was requested and successfully taken.
func (c *HeadlessExperimental) BeginFrameWithParams(v *HeadlessExperimentalBeginFrameParams) (bool, bool, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.beginFrame", Params: v})
	if err != nil {
		return false, false, "", err
	}

	var chromeData struct {
		Result struct {
			HasDamage               bool
			MainFrameContentUpdated bool
			ScreenshotData          string
		}
	}

	if resp == nil {
		return false, false, "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, false, "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, false, "", err
	}

	return chromeData.Result.HasDamage, chromeData.Result.MainFrameContentUpdated, chromeData.Result.ScreenshotData, nil
}

// BeginFrame - Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl.
// frameTime - Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used.
// deadline - Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval.
// interval - The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
// screenshot - If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured.
// Returns -  hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. mainFrameContentUpdated - Whether the main frame submitted a new display frame in response to this BeginFrame. screenshotData - Base64-encoded image data of the screenshot, if one was requested and successfully taken.
func (c *HeadlessExperimental) BeginFrame(frameTime float64, deadline float64, interval float64, screenshot *HeadlessExperimentalScreenshotParams) (bool, bool, string, error) {
	var v HeadlessExperimentalBeginFrameParams
	v.FrameTime = frameTime
	v.Deadline = deadline
	v.Interval = interval
	v.Screenshot = screenshot
	return c.BeginFrameWithParams(&v)
}

// Disables headless events for the target.
func (c *HeadlessExperimental) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.disable"})
}

// Enables headless events for the target.
func (c *HeadlessExperimental) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.enable"})
}
