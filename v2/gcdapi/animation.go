// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Animation functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Animation instance.
type AnimationAnimation struct {
	Id           string                    `json:"id"`               // `Animation`'s id.
	Name         string                    `json:"name"`             // `Animation`'s name.
	PausedState  bool                      `json:"pausedState"`      // `Animation`'s internal paused state.
	PlayState    string                    `json:"playState"`        // `Animation`'s play state.
	PlaybackRate float64                   `json:"playbackRate"`     // `Animation`'s playback rate.
	StartTime    float64                   `json:"startTime"`        // `Animation`'s start time.
	CurrentTime  float64                   `json:"currentTime"`      // `Animation`'s current time.
	Type         string                    `json:"type"`             // Animation type of `Animation`.
	Source       *AnimationAnimationEffect `json:"source,omitempty"` // `Animation`'s source animation node.
	CssId        string                    `json:"cssId,omitempty"`  // A unique ID for `Animation` representing the sources that triggered this CSS animation/transition.
}

// AnimationEffect instance
type AnimationAnimationEffect struct {
	Delay          float64                 `json:"delay"`                   // `AnimationEffect`'s delay.
	EndDelay       float64                 `json:"endDelay"`                // `AnimationEffect`'s end delay.
	IterationStart float64                 `json:"iterationStart"`          // `AnimationEffect`'s iteration start.
	Iterations     float64                 `json:"iterations"`              // `AnimationEffect`'s iterations.
	Duration       float64                 `json:"duration"`                // `AnimationEffect`'s iteration duration.
	Direction      string                  `json:"direction"`               // `AnimationEffect`'s playback direction.
	Fill           string                  `json:"fill"`                    // `AnimationEffect`'s fill mode.
	BackendNodeId  int                     `json:"backendNodeId,omitempty"` // `AnimationEffect`'s target node.
	KeyframesRule  *AnimationKeyframesRule `json:"keyframesRule,omitempty"` // `AnimationEffect`'s keyframes.
	Easing         string                  `json:"easing"`                  // `AnimationEffect`'s timing function.
}

// Keyframes Rule
type AnimationKeyframesRule struct {
	Name      string                    `json:"name,omitempty"` // CSS keyframed animation's name.
	Keyframes []*AnimationKeyframeStyle `json:"keyframes"`      // List of animation keyframes.
}

// Keyframe Style
type AnimationKeyframeStyle struct {
	Offset string `json:"offset"` // Keyframe's time offset.
	Easing string `json:"easing"` // `AnimationEffect`'s timing function.
}

// Event for when an animation has been cancelled.
type AnimationAnimationCanceledEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` // Id of the animation that was cancelled.
	} `json:"Params,omitempty"`
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

type Animation struct {
	target gcdmessage.ChromeTargeter
}

func NewAnimation(target gcdmessage.ChromeTargeter) *Animation {
	c := &Animation{target: target}
	return c
}

// Disables animation domain notifications.
func (c *Animation) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.disable"})
}

// Enables animation domain notifications.
func (c *Animation) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.enable"})
}

type AnimationGetCurrentTimeParams struct {
	// Id of animation.
	Id string `json:"id"`
}

// GetCurrentTimeWithParams - Returns the current time of the an animation.
// Returns -  currentTime - Current time of the page.
func (c *Animation) GetCurrentTimeWithParams(ctx context.Context, v *AnimationGetCurrentTimeParams) (float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.getCurrentTime", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			CurrentTime float64
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	if chromeData.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.CurrentTime, nil
}

// GetCurrentTime - Returns the current time of the an animation.
// id - Id of animation.
// Returns -  currentTime - Current time of the page.
func (c *Animation) GetCurrentTime(ctx context.Context, id string) (float64, error) {
	var v AnimationGetCurrentTimeParams
	v.Id = id
	return c.GetCurrentTimeWithParams(ctx, &v)
}

// GetPlaybackRate - Gets the playback rate of the document timeline.
// Returns -  playbackRate - Playback rate for animations on page.
func (c *Animation) GetPlaybackRate(ctx context.Context) (float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.getPlaybackRate"})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			PlaybackRate float64
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	if chromeData.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.PlaybackRate, nil
}

type AnimationReleaseAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

// ReleaseAnimationsWithParams - Releases a set of animations to no longer be manipulated.
func (c *Animation) ReleaseAnimationsWithParams(ctx context.Context, v *AnimationReleaseAnimationsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.releaseAnimations", Params: v})
}

// ReleaseAnimations - Releases a set of animations to no longer be manipulated.
// animations - List of animation ids to seek.
func (c *Animation) ReleaseAnimations(ctx context.Context, animations []string) (*gcdmessage.ChromeResponse, error) {
	var v AnimationReleaseAnimationsParams
	v.Animations = animations
	return c.ReleaseAnimationsWithParams(ctx, &v)
}

type AnimationResolveAnimationParams struct {
	// Animation id.
	AnimationId string `json:"animationId"`
}

// ResolveAnimationWithParams - Gets the remote object of the Animation.
// Returns -  remoteObject - Corresponding remote object.
func (c *Animation) ResolveAnimationWithParams(ctx context.Context, v *AnimationResolveAnimationParams) (*RuntimeRemoteObject, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.resolveAnimation", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			RemoteObject *RuntimeRemoteObject
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.RemoteObject, nil
}

// ResolveAnimation - Gets the remote object of the Animation.
// animationId - Animation id.
// Returns -  remoteObject - Corresponding remote object.
func (c *Animation) ResolveAnimation(ctx context.Context, animationId string) (*RuntimeRemoteObject, error) {
	var v AnimationResolveAnimationParams
	v.AnimationId = animationId
	return c.ResolveAnimationWithParams(ctx, &v)
}

type AnimationSeekAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
	// Set the current time of each animation.
	CurrentTime float64 `json:"currentTime"`
}

// SeekAnimationsWithParams - Seek a set of animations to a particular time within each animation.
func (c *Animation) SeekAnimationsWithParams(ctx context.Context, v *AnimationSeekAnimationsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.seekAnimations", Params: v})
}

// SeekAnimations - Seek a set of animations to a particular time within each animation.
// animations - List of animation ids to seek.
// currentTime - Set the current time of each animation.
func (c *Animation) SeekAnimations(ctx context.Context, animations []string, currentTime float64) (*gcdmessage.ChromeResponse, error) {
	var v AnimationSeekAnimationsParams
	v.Animations = animations
	v.CurrentTime = currentTime
	return c.SeekAnimationsWithParams(ctx, &v)
}

type AnimationSetPausedParams struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`
	// Paused state to set to.
	Paused bool `json:"paused"`
}

// SetPausedWithParams - Sets the paused state of a set of animations.
func (c *Animation) SetPausedWithParams(ctx context.Context, v *AnimationSetPausedParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setPaused", Params: v})
}

// SetPaused - Sets the paused state of a set of animations.
// animations - Animations to set the pause state of.
// paused - Paused state to set to.
func (c *Animation) SetPaused(ctx context.Context, animations []string, paused bool) (*gcdmessage.ChromeResponse, error) {
	var v AnimationSetPausedParams
	v.Animations = animations
	v.Paused = paused
	return c.SetPausedWithParams(ctx, &v)
}

type AnimationSetPlaybackRateParams struct {
	// Playback rate for animations on page
	PlaybackRate float64 `json:"playbackRate"`
}

// SetPlaybackRateWithParams - Sets the playback rate of the document timeline.
func (c *Animation) SetPlaybackRateWithParams(ctx context.Context, v *AnimationSetPlaybackRateParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setPlaybackRate", Params: v})
}

// SetPlaybackRate - Sets the playback rate of the document timeline.
// playbackRate - Playback rate for animations on page
func (c *Animation) SetPlaybackRate(ctx context.Context, playbackRate float64) (*gcdmessage.ChromeResponse, error) {
	var v AnimationSetPlaybackRateParams
	v.PlaybackRate = playbackRate
	return c.SetPlaybackRateWithParams(ctx, &v)
}

type AnimationSetTimingParams struct {
	// Animation id.
	AnimationId string `json:"animationId"`
	// Duration of the animation.
	Duration float64 `json:"duration"`
	// Delay of the animation.
	Delay float64 `json:"delay"`
}

// SetTimingWithParams - Sets the timing of an animation node.
func (c *Animation) SetTimingWithParams(ctx context.Context, v *AnimationSetTimingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Animation.setTiming", Params: v})
}

// SetTiming - Sets the timing of an animation node.
// animationId - Animation id.
// duration - Duration of the animation.
// delay - Delay of the animation.
func (c *Animation) SetTiming(ctx context.Context, animationId string, duration float64, delay float64) (*gcdmessage.ChromeResponse, error) {
	var v AnimationSetTimingParams
	v.AnimationId = animationId
	v.Duration = duration
	v.Delay = delay
	return c.SetTimingWithParams(ctx, &v)
}
