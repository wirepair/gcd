// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeadlessExperimental functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Encoding options for a screenshot.
type HeadlessExperimentalScreenshotParams struct {
	Format           string `json:"format,omitempty"`           // Image compression format (defaults to png).
	Quality          int    `json:"quality,omitempty"`          // Compression quality from range [0..100] (jpeg and webp only).
	OptimizeForSpeed bool   `json:"optimizeForSpeed,omitempty"` // Optimize image encoding for speed, not for resulting size (defaults to false)
}

type HeadlessExperimental struct {
	target gcdmessage.ChromeTargeter
}

func NewHeadlessExperimental(target gcdmessage.ChromeTargeter) *HeadlessExperimental {
	c := &HeadlessExperimental{target: target}
	return c
}

type HeadlessExperimentalBeginFrameParams struct {
	// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used.
	FrameTimeTicks float64 `json:"frameTimeTicks,omitempty"`
	// The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval float64 `json:"interval,omitempty"`
	// Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	NoDisplayUpdates bool `json:"noDisplayUpdates,omitempty"`
	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
	Screenshot *HeadlessExperimentalScreenshotParams `json:"screenshot,omitempty"`
}

// BeginFrameWithParams - Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also https://goo.gle/chrome-headless-rendering for more background.
// Returns -  hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
func (c *HeadlessExperimental) BeginFrameWithParams(ctx context.Context, v *HeadlessExperimentalBeginFrameParams) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.beginFrame", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			HasDamage bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.HasDamage, nil
}

// BeginFrame - Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also https://goo.gle/chrome-headless-rendering for more background.
// frameTimeTicks - Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used.
// interval - The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
// noDisplayUpdates - Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
// screenshot - If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
// Returns -  hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
func (c *HeadlessExperimental) BeginFrame(ctx context.Context, frameTimeTicks float64, interval float64, noDisplayUpdates bool, screenshot *HeadlessExperimentalScreenshotParams) (bool, error) {
	var v HeadlessExperimentalBeginFrameParams
	v.FrameTimeTicks = frameTimeTicks
	v.Interval = interval
	v.NoDisplayUpdates = noDisplayUpdates
	v.Screenshot = screenshot
	return c.BeginFrameWithParams(ctx, &v)
}

// Disables headless events for the target.
func (c *HeadlessExperimental) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.disable"})
}

// Enables headless events for the target.
func (c *HeadlessExperimental) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeadlessExperimental.enable"})
}
