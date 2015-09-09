// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Animation commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Animation() *ChromeAnimation {
	if c.animation == nil {
		c.animation = newChromeAnimation(c)
	}
	return c.animation
}

type ChromeAnimation struct {
	target *ChromeTarget
}

func newChromeAnimation(target *ChromeTarget) *ChromeAnimation {
	c := &ChromeAnimation{target: target}
	return c
}

// Enables animation domain notifications.
func (c *ChromeAnimation) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.enable"})
}

// Disables animation domain notifications.
func (c *ChromeAnimation) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.disable"})
}

// setPlaybackRate - Sets the playback rate of the document timeline.
// playbackRate - Playback rate for animations on page
func (c *ChromeAnimation) SetPlaybackRate(playbackRate float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["playbackRate"] = playbackRate
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.setPlaybackRate", Params: paramRequest})
}

// setCurrentTime - Sets the current time of the document timeline.
// currentTime - Current time for the page animation timeline
func (c *ChromeAnimation) SetCurrentTime(currentTime float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["currentTime"] = currentTime
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.setCurrentTime", Params: paramRequest})
}

// setTiming - Sets the timing of an animation node.
// playerId - AnimationPlayer id.
// duration - Duration of the animation.
// delay - Delay of the animation.
func (c *ChromeAnimation) SetTiming(playerId string, duration float64, delay float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["playerId"] = playerId
	paramRequest["duration"] = duration
	paramRequest["delay"] = delay
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.setTiming", Params: paramRequest})
}

// getPlaybackRate - Gets the playback rate of the document timeline.
// Returns -
// Playback rate for animations on page.
func (c *ChromeAnimation) GetPlaybackRate() (float64, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Animation.getPlaybackRate"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			PlaybackRate float64
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.PlaybackRate, nil
}
