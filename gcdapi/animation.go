// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Animation functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Animation instance.
type AnimationAnimation struct {
	Id           string                    `json:"id"`           // <code>Animation</code>'s id.
	PausedState  bool                      `json:"pausedState"`  // <code>Animation</code>'s internal paused state.
	PlayState    string                    `json:"playState"`    // <code>Animation</code>'s play state.
	PlaybackRate float64                   `json:"playbackRate"` // <code>Animation</code>'s playback rate.
	StartTime    float64                   `json:"startTime"`    // <code>Animation</code>'s start time.
	CurrentTime  float64                   `json:"currentTime"`  // <code>Animation</code>'s current time.
	Source       *AnimationAnimationEffect `json:"source"`       // <code>Animation</code>'s source animation node.
	Type         string                    `json:"type"`         // Animation type of <code>Animation</code>.
}

// AnimationEffect instance
type AnimationAnimationEffect struct {
	Delay          float64                 `json:"delay"`                   // <code>AnimationEffect</code>'s delay.
	EndDelay       float64                 `json:"endDelay"`                // <code>AnimationEffect</code>'s end delay.
	PlaybackRate   float64                 `json:"playbackRate"`            // <code>AnimationEffect</code>'s playbackRate.
	IterationStart float64                 `json:"iterationStart"`          // <code>AnimationEffect</code>'s iteration start.
	Iterations     float64                 `json:"iterations"`              // <code>AnimationEffect</code>'s iterations.
	Duration       float64                 `json:"duration"`                // <code>AnimationEffect</code>'s iteration duration.
	Direction      string                  `json:"direction"`               // <code>AnimationEffect</code>'s playback direction.
	Fill           string                  `json:"fill"`                    // <code>AnimationEffect</code>'s fill mode.
	Name           string                  `json:"name"`                    // <code>AnimationEffect</code>'s name.
	BackendNodeId  int                     `json:"backendNodeId"`           // <code>AnimationEffect</code>'s target node.
	KeyframesRule  *AnimationKeyframesRule `json:"keyframesRule,omitempty"` // <code>AnimationEffect</code>'s keyframes.
	Easing         string                  `json:"easing"`                  // <code>AnimationEffect</code>'s timing function.
}

// Keyframes Rule
type AnimationKeyframesRule struct {
	Name      string                    `json:"name,omitempty"` // CSS keyframed animation's name.
	Keyframes []*AnimationKeyframeStyle `json:"keyframes"`      // List of animation keyframes.
}

// Keyframe Style
type AnimationKeyframeStyle struct {
	Offset string `json:"offset"` // Keyframe's time offset.
	Easing string `json:"easing"` // <code>AnimationEffect</code>'s timing function.
}

// Event for each animation player that has been created.
type AnimationAnimationCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Player        *AnimationAnimation `json:"player"`        // Animation that was created.
		ResetTimeline bool                `json:"resetTimeline"` // Whether the timeline should be reset.
	} `json:"Params,omitempty"`
}

// Event for Animations in the frontend that have been cancelled.
type AnimationAnimationCanceledEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` // Id of the Animation that was cancelled.
	} `json:"Params,omitempty"`
}

type Animation struct {
	target gcdmessage.ChromeTargeter
}

func NewAnimation(target gcdmessage.ChromeTargeter) *Animation {
	c := &Animation{target: target}
	return c
}

// Enables animation domain notifications.
func (c *Animation) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.enable"})
}

// Disables animation domain notifications.
func (c *Animation) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.disable"})
}

// GetPlaybackRate - Gets the playback rate of the document timeline.
// Returns -  playbackRate - Playback rate for animations on page.
func (c *Animation) GetPlaybackRate() (float64, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.getPlaybackRate"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			PlaybackRate float64
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.PlaybackRate, nil
}

// SetPlaybackRate - Sets the playback rate of the document timeline.
// playbackRate - Playback rate for animations on page
func (c *Animation) SetPlaybackRate(playbackRate float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["playbackRate"] = playbackRate
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setPlaybackRate", Params: paramRequest})
}

// SetCurrentTime - Sets the current time of the document timeline.
// currentTime - Current time for the page animation timeline
func (c *Animation) SetCurrentTime(currentTime float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["currentTime"] = currentTime
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setCurrentTime", Params: paramRequest})
}

// SetTiming - Sets the timing of an animation node.
// playerId - AnimationPlayer id.
// duration - Duration of the animation.
// delay - Delay of the animation.
func (c *Animation) SetTiming(playerId string, duration float64, delay float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["playerId"] = playerId
	paramRequest["duration"] = duration
	paramRequest["delay"] = delay
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setTiming", Params: paramRequest})
}
