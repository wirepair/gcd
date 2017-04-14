// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Animation functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Animation instance.
type AnimationAnimation struct {
	Id           string                    `json:"id"`              // <code>Animation</code>'s id.
	Name         string                    `json:"name"`            // <code>Animation</code>'s name.
	PausedState  bool                      `json:"pausedState"`     // <code>Animation</code>'s internal paused state.
	PlayState    string                    `json:"playState"`       // <code>Animation</code>'s play state.
	PlaybackRate float64                   `json:"playbackRate"`    // <code>Animation</code>'s playback rate.
	StartTime    float64                   `json:"startTime"`       // <code>Animation</code>'s start time.
	CurrentTime  float64                   `json:"currentTime"`     // <code>Animation</code>'s current time.
	Source       *AnimationAnimationEffect `json:"source"`          // <code>Animation</code>'s source animation node.
	Type         string                    `json:"type"`            // Animation type of <code>Animation</code>.
	CssId        string                    `json:"cssId,omitempty"` // A unique ID for <code>Animation</code> representing the sources that triggered this CSS animation/transition.
}

// AnimationEffect instance
type AnimationAnimationEffect struct {
	Delay          float64                 `json:"delay"`                   // <code>AnimationEffect</code>'s delay.
	EndDelay       float64                 `json:"endDelay"`                // <code>AnimationEffect</code>'s end delay.
	IterationStart float64                 `json:"iterationStart"`          // <code>AnimationEffect</code>'s iteration start.
	Iterations     float64                 `json:"iterations"`              // <code>AnimationEffect</code>'s iterations.
	Duration       float64                 `json:"duration"`                // <code>AnimationEffect</code>'s iteration duration.
	Direction      string                  `json:"direction"`               // <code>AnimationEffect</code>'s playback direction.
	Fill           string                  `json:"fill"`                    // <code>AnimationEffect</code>'s fill mode.
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

// Event for each animation that has been created.
type AnimationAnimationCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` // Id of the animation that was created.
	} `json:"Params,omitempty"`
}

// Event for animation that has been started.
type AnimationAnimationStartedEvent struct {
	Method string `json:"method"`
	Params struct {
		Animation *AnimationAnimation `json:"animation"` // Animation that was started.
	} `json:"Params,omitempty"`
}

// Event for when an animation has been cancelled.
type AnimationAnimationCanceledEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` // Id of the animation that was cancelled.
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
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.enable"})
}

// Disables animation domain notifications.
func (c *Animation) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.disable"})
}

// GetPlaybackRate - Gets the playback rate of the document timeline.
// Returns -  playbackRate - Playback rate for animations on page.
func (c *Animation) GetPlaybackRate() (float64, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.getPlaybackRate"})
	if err != nil {
		return 0, err
	}

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

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	return chromeData.Result.PlaybackRate, nil
}

// SetPlaybackRate - Sets the playback rate of the document timeline.
// playbackRate - Playback rate for animations on page
func (c *Animation) SetPlaybackRate(playbackRate float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["playbackRate"] = playbackRate
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setPlaybackRate", Params: paramRequest})
}

// GetCurrentTime - Returns the current time of the an animation.
// id - Id of animation.
// Returns -  currentTime - Current time of the page.
func (c *Animation) GetCurrentTime(id string) (float64, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["id"] = id
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.getCurrentTime", Params: paramRequest})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			CurrentTime float64
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

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	return chromeData.Result.CurrentTime, nil
}

// SetPaused - Sets the paused state of a set of animations.
// animations - Animations to set the pause state of.
// paused - Paused state to set to.
func (c *Animation) SetPaused(animations string, paused bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["animations"] = animations
	paramRequest["paused"] = paused
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setPaused", Params: paramRequest})
}

// SetTiming - Sets the timing of an animation node.
// animationId - Animation id.
// duration - Duration of the animation.
// delay - Delay of the animation.
func (c *Animation) SetTiming(animationId string, duration float64, delay float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["animationId"] = animationId
	paramRequest["duration"] = duration
	paramRequest["delay"] = delay
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setTiming", Params: paramRequest})
}

// SeekAnimations - Seek a set of animations to a particular time within each animation.
// animations - List of animation ids to seek.
// currentTime - Set the current time of each animation.
func (c *Animation) SeekAnimations(animations string, currentTime float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["animations"] = animations
	paramRequest["currentTime"] = currentTime
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.seekAnimations", Params: paramRequest})
}

// ReleaseAnimations - Releases a set of animations to no longer be manipulated.
// animations - List of animation ids to seek.
func (c *Animation) ReleaseAnimations(animations string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["animations"] = animations
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.releaseAnimations", Params: paramRequest})
}

// ResolveAnimation - Gets the remote object of the Animation.
// animationId - Animation id.
// Returns -  remoteObject - Corresponding remote object.
func (c *Animation) ResolveAnimation(animationId string) (*RuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["animationId"] = animationId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.resolveAnimation", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RemoteObject *RuntimeRemoteObject
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.RemoteObject, nil
}
