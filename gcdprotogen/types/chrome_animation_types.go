// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Animation types.
// API Version: 1.1
package types

// Animation instance.
type ChromeAnimationAnimation struct {
	Id           string                          `json:"id"`           // <code>Animation</code>'s id.
	PausedState  bool                            `json:"pausedState"`  // <code>Animation</code>'s internal paused state.
	PlayState    string                          `json:"playState"`    // <code>Animation</code>'s play state.
	PlaybackRate float64                         `json:"playbackRate"` // <code>Animation</code>'s playback rate.
	StartTime    float64                         `json:"startTime"`    // <code>Animation</code>'s start time.
	CurrentTime  float64                         `json:"currentTime"`  // <code>Animation</code>'s current time.
	Source       *ChromeAnimationAnimationEffect `json:"source"`       // <code>Animation</code>'s source animation node.
	Type         string                          `json:"type"`         // Animation type of <code>Animation</code>.
}

// AnimationEffect instance
type ChromeAnimationAnimationEffect struct {
	Delay          float64                       `json:"delay"`                   // <code>AnimationEffect</code>'s delay.
	EndDelay       float64                       `json:"endDelay"`                // <code>AnimationEffect</code>'s end delay.
	PlaybackRate   float64                       `json:"playbackRate"`            // <code>AnimationEffect</code>'s playbackRate.
	IterationStart float64                       `json:"iterationStart"`          // <code>AnimationEffect</code>'s iteration start.
	Iterations     float64                       `json:"iterations"`              // <code>AnimationEffect</code>'s iterations.
	Duration       float64                       `json:"duration"`                // <code>AnimationEffect</code>'s iteration duration.
	Direction      string                        `json:"direction"`               // <code>AnimationEffect</code>'s playback direction.
	Fill           string                        `json:"fill"`                    // <code>AnimationEffect</code>'s fill mode.
	Name           string                        `json:"name"`                    // <code>AnimationEffect</code>'s name.
	BackendNodeId  *ChromeDOMBackendNodeId       `json:"backendNodeId"`           // <code>AnimationEffect</code>'s target node.
	KeyframesRule  *ChromeAnimationKeyframesRule `json:"keyframesRule,omitempty"` // <code>AnimationEffect</code>'s keyframes.
	Easing         string                        `json:"easing"`                  // <code>AnimationEffect</code>'s timing function.
}

// Keyframes Rule
type ChromeAnimationKeyframesRule struct {
	Name      string                          `json:"name,omitempty"` // CSS keyframed animation's name.
	Keyframes []*ChromeAnimationKeyframeStyle `json:"keyframes"`      // List of animation keyframes.
}

// Keyframe Style
type ChromeAnimationKeyframeStyle struct {
	Offset string `json:"offset"` // Keyframe's time offset.
	Easing string `json:"easing"` // <code>AnimationEffect</code>'s timing function.
}
