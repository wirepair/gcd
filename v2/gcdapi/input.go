// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Input functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type InputTouchPoint struct {
	X                  float64 `json:"x"`                            // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y                  float64 `json:"y"`                            // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX            float64 `json:"radiusX,omitempty"`            // X radius of the touch area (default: 1.0).
	RadiusY            float64 `json:"radiusY,omitempty"`            // Y radius of the touch area (default: 1.0).
	RotationAngle      float64 `json:"rotationAngle,omitempty"`      // Rotation angle (default: 0.0).
	Force              float64 `json:"force,omitempty"`              // Force (default: 1.0).
	TangentialPressure float64 `json:"tangentialPressure,omitempty"` // The normalized tangential pressure, which has a range of [-1,1] (default: 0).
	TiltX              int     `json:"tiltX,omitempty"`              // The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0)
	TiltY              int     `json:"tiltY,omitempty"`              // The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	Twist              int     `json:"twist,omitempty"`              // The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
	Id                 float64 `json:"id,omitempty"`                 // Identifier used to track touch sources between events, must be unique within an event.
}

// No Description.
type InputDragDataItem struct {
	MimeType string `json:"mimeType"`          // Mime type of the dragged data.
	Data     string `json:"data"`              // Depending of the value of `mimeType`, it contains the dragged link, text, HTML markup or any other data.
	Title    string `json:"title,omitempty"`   // Title associated with a link. Only valid when `mimeType` == "text/uri-list".
	BaseURL  string `json:"baseURL,omitempty"` // Stores the base URL for the contained markup. Only valid when `mimeType` == "text/html".
}

// No Description.
type InputDragData struct {
	Items              []*InputDragDataItem `json:"items"`              //
	Files              []string             `json:"files,omitempty"`    // List of filenames that should be included when dropping
	DragOperationsMask int                  `json:"dragOperationsMask"` // Bit field representing allowed drag operations. Copy = 1, Link = 2, Move = 16
}

// Emitted only when `Input.setInterceptDrags` is enabled. Use this data with `Input.dispatchDragEvent` to restore normal drag and drop behavior.
type InputDragInterceptedEvent struct {
	Method string `json:"method"`
	Params struct {
		Data *InputDragData `json:"data"` //
	} `json:"Params,omitempty"`
}

type Input struct {
	target gcdmessage.ChromeTargeter
}

func NewInput(target gcdmessage.ChromeTargeter) *Input {
	c := &Input{target: target}
	return c
}

type InputDispatchDragEventParams struct {
	// Type of the drag event.
	TheType string `json:"type"`
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float64 `json:"y"`
	//
	Data *InputDragData `json:"data"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
}

// DispatchDragEventWithParams - Dispatches a drag event into the page.
func (c *Input) DispatchDragEventWithParams(ctx context.Context, v *InputDispatchDragEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchDragEvent", Params: v})
}

// DispatchDragEvent - Dispatches a drag event into the page.
// type - Type of the drag event.
// x - X coordinate of the event relative to the main frame's viewport in CSS pixels.
// y - Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
// data -
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
func (c *Input) DispatchDragEvent(ctx context.Context, theType string, x float64, y float64, data *InputDragData, modifiers int) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchDragEventParams
	v.TheType = theType
	v.X = x
	v.Y = y
	v.Data = data
	v.Modifiers = modifiers
	return c.DispatchDragEventWithParams(ctx, &v)
}

type InputDispatchKeyEventParams struct {
	// Type of the key event.
	TheType string `json:"type"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
	// Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
	Text string `json:"text,omitempty"`
	// Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	UnmodifiedText string `json:"unmodifiedText,omitempty"`
	// Unique key identifier (e.g., 'U+0041') (default: "").
	KeyIdentifier string `json:"keyIdentifier,omitempty"`
	// Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Code string `json:"code,omitempty"`
	// Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	Key string `json:"key,omitempty"`
	// Windows virtual key code (default: 0).
	WindowsVirtualKeyCode int `json:"windowsVirtualKeyCode,omitempty"`
	// Native virtual key code (default: 0).
	NativeVirtualKeyCode int `json:"nativeVirtualKeyCode,omitempty"`
	// Whether the event was generated from auto repeat (default: false).
	AutoRepeat bool `json:"autoRepeat,omitempty"`
	// Whether the event was generated from the keypad (default: false).
	IsKeypad bool `json:"isKeypad,omitempty"`
	// Whether the event was a system key event (default: false).
	IsSystemKey bool `json:"isSystemKey,omitempty"`
	// Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
	Location int `json:"location,omitempty"`
	// Editing commands to send with the key event (e.g., 'selectAll') (default: []). These are related to but not equal the command names used in `document.execCommand` and NSStandardKeyBindingResponding. See https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h for valid command names.
	Commands []string `json:"commands,omitempty"`
}

// DispatchKeyEventWithParams - Dispatches a key event to the page.
func (c *Input) DispatchKeyEventWithParams(ctx context.Context, v *InputDispatchKeyEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchKeyEvent", Params: v})
}

// DispatchKeyEvent - Dispatches a key event to the page.
// type - Type of the key event.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
// text - Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
// unmodifiedText - Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
// keyIdentifier - Unique key identifier (e.g., 'U+0041') (default: "").
// code - Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
// key - Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
// windowsVirtualKeyCode - Windows virtual key code (default: 0).
// nativeVirtualKeyCode - Native virtual key code (default: 0).
// autoRepeat - Whether the event was generated from auto repeat (default: false).
// isKeypad - Whether the event was generated from the keypad (default: false).
// isSystemKey - Whether the event was a system key event (default: false).
// location - Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
// commands - Editing commands to send with the key event (e.g., 'selectAll') (default: []). These are related to but not equal the command names used in `document.execCommand` and NSStandardKeyBindingResponding. See https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h for valid command names.
func (c *Input) DispatchKeyEvent(ctx context.Context, theType string, modifiers int, timestamp float64, text string, unmodifiedText string, keyIdentifier string, code string, key string, windowsVirtualKeyCode int, nativeVirtualKeyCode int, autoRepeat bool, isKeypad bool, isSystemKey bool, location int, commands []string) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchKeyEventParams
	v.TheType = theType
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	v.Text = text
	v.UnmodifiedText = unmodifiedText
	v.KeyIdentifier = keyIdentifier
	v.Code = code
	v.Key = key
	v.WindowsVirtualKeyCode = windowsVirtualKeyCode
	v.NativeVirtualKeyCode = nativeVirtualKeyCode
	v.AutoRepeat = autoRepeat
	v.IsKeypad = isKeypad
	v.IsSystemKey = isSystemKey
	v.Location = location
	v.Commands = commands
	return c.DispatchKeyEventWithParams(ctx, &v)
}

type InputInsertTextParams struct {
	// The text to insert.
	Text string `json:"text"`
}

// InsertTextWithParams - This method emulates inserting text that doesn't come from a key press, for example an emoji keyboard or an IME.
func (c *Input) InsertTextWithParams(ctx context.Context, v *InputInsertTextParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.insertText", Params: v})
}

// InsertText - This method emulates inserting text that doesn't come from a key press, for example an emoji keyboard or an IME.
// text - The text to insert.
func (c *Input) InsertText(ctx context.Context, text string) (*gcdmessage.ChromeResponse, error) {
	var v InputInsertTextParams
	v.Text = text
	return c.InsertTextWithParams(ctx, &v)
}

type InputImeSetCompositionParams struct {
	// The text to insert
	Text string `json:"text"`
	// selection start
	SelectionStart int `json:"selectionStart"`
	// selection end
	SelectionEnd int `json:"selectionEnd"`
	// replacement start
	ReplacementStart int `json:"replacementStart,omitempty"`
	// replacement end
	ReplacementEnd int `json:"replacementEnd,omitempty"`
}

// ImeSetCompositionWithParams - This method sets the current candidate text for ime. Use imeCommitComposition to commit the final text. Use imeSetComposition with empty string as text to cancel composition.
func (c *Input) ImeSetCompositionWithParams(ctx context.Context, v *InputImeSetCompositionParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.imeSetComposition", Params: v})
}

// ImeSetComposition - This method sets the current candidate text for ime. Use imeCommitComposition to commit the final text. Use imeSetComposition with empty string as text to cancel composition.
// text - The text to insert
// selectionStart - selection start
// selectionEnd - selection end
// replacementStart - replacement start
// replacementEnd - replacement end
func (c *Input) ImeSetComposition(ctx context.Context, text string, selectionStart int, selectionEnd int, replacementStart int, replacementEnd int) (*gcdmessage.ChromeResponse, error) {
	var v InputImeSetCompositionParams
	v.Text = text
	v.SelectionStart = selectionStart
	v.SelectionEnd = selectionEnd
	v.ReplacementStart = replacementStart
	v.ReplacementEnd = replacementEnd
	return c.ImeSetCompositionWithParams(ctx, &v)
}

type InputDispatchMouseEventParams struct {
	// Type of the mouse event.
	TheType string `json:"type"`
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float64 `json:"y"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
	// Mouse button (default: "none"). enum values: none, left, middle, right, back, forward
	Button string `json:"button,omitempty"`
	// A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
	Buttons int `json:"buttons,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`
	// The normalized pressure, which has a range of [0,1] (default: 0).
	Force float64 `json:"force,omitempty"`
	// The normalized tangential pressure, which has a range of [-1,1] (default: 0).
	TangentialPressure float64 `json:"tangentialPressure,omitempty"`
	// The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0).
	TiltX int `json:"tiltX,omitempty"`
	// The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	TiltY int `json:"tiltY,omitempty"`
	// The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
	Twist int `json:"twist,omitempty"`
	// X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaX float64 `json:"deltaX,omitempty"`
	// Y delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY float64 `json:"deltaY,omitempty"`
	// Pointer type (default: "mouse").
	PointerType string `json:"pointerType,omitempty"`
}

// DispatchMouseEventWithParams - Dispatches a mouse event to the page.
func (c *Input) DispatchMouseEventWithParams(ctx context.Context, v *InputDispatchMouseEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchMouseEvent", Params: v})
}

// DispatchMouseEvent - Dispatches a mouse event to the page.
// type - Type of the mouse event.
// x - X coordinate of the event relative to the main frame's viewport in CSS pixels.
// y - Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
// button - Mouse button (default: "none"). enum values: none, left, middle, right, back, forward
// buttons - A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
// clickCount - Number of times the mouse button was clicked (default: 0).
// force - The normalized pressure, which has a range of [0,1] (default: 0).
// tangentialPressure - The normalized tangential pressure, which has a range of [-1,1] (default: 0).
// tiltX - The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0).
// tiltY - The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
// twist - The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
// deltaX - X delta in CSS pixels for mouse wheel event (default: 0).
// deltaY - Y delta in CSS pixels for mouse wheel event (default: 0).
// pointerType - Pointer type (default: "mouse").
func (c *Input) DispatchMouseEvent(ctx context.Context, theType string, x float64, y float64, modifiers int, timestamp float64, button string, buttons int, clickCount int, force float64, tangentialPressure float64, tiltX int, tiltY int, twist int, deltaX float64, deltaY float64, pointerType string) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchMouseEventParams
	v.TheType = theType
	v.X = x
	v.Y = y
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	v.Button = button
	v.Buttons = buttons
	v.ClickCount = clickCount
	v.Force = force
	v.TangentialPressure = tangentialPressure
	v.TiltX = tiltX
	v.TiltY = tiltY
	v.Twist = twist
	v.DeltaX = deltaX
	v.DeltaY = deltaY
	v.PointerType = pointerType
	return c.DispatchMouseEventWithParams(ctx, &v)
}

type InputDispatchTouchEventParams struct {
	// Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
	TheType string `json:"type"`
	// Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
	TouchPoints []*InputTouchPoint `json:"touchPoints"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
}

// DispatchTouchEventWithParams - Dispatches a touch event to the page.
func (c *Input) DispatchTouchEventWithParams(ctx context.Context, v *InputDispatchTouchEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchTouchEvent", Params: v})
}

// DispatchTouchEvent - Dispatches a touch event to the page.
// type - Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
// touchPoints - Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
func (c *Input) DispatchTouchEvent(ctx context.Context, theType string, touchPoints []*InputTouchPoint, modifiers int, timestamp float64) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchTouchEventParams
	v.TheType = theType
	v.TouchPoints = touchPoints
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	return c.DispatchTouchEventWithParams(ctx, &v)
}

type InputEmulateTouchFromMouseEventParams struct {
	// Type of the mouse event.
	TheType string `json:"type"`
	// X coordinate of the mouse pointer in DIP.
	X int `json:"x"`
	// Y coordinate of the mouse pointer in DIP.
	Y int `json:"y"`
	// Mouse button. Only "none", "left", "right" are supported. enum values: none, left, middle, right, back, forward
	Button string `json:"button"`
	// Time at which the event occurred (default: current time).
	Timestamp float64 `json:"timestamp,omitempty"`
	// X delta in DIP for mouse wheel event (default: 0).
	DeltaX float64 `json:"deltaX,omitempty"`
	// Y delta in DIP for mouse wheel event (default: 0).
	DeltaY float64 `json:"deltaY,omitempty"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`
}

// EmulateTouchFromMouseEventWithParams - Emulates touch event from the mouse event parameters.
func (c *Input) EmulateTouchFromMouseEventWithParams(ctx context.Context, v *InputEmulateTouchFromMouseEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.emulateTouchFromMouseEvent", Params: v})
}

// EmulateTouchFromMouseEvent - Emulates touch event from the mouse event parameters.
// type - Type of the mouse event.
// x - X coordinate of the mouse pointer in DIP.
// y - Y coordinate of the mouse pointer in DIP.
// button - Mouse button. Only "none", "left", "right" are supported. enum values: none, left, middle, right, back, forward
// timestamp - Time at which the event occurred (default: current time).
// deltaX - X delta in DIP for mouse wheel event (default: 0).
// deltaY - Y delta in DIP for mouse wheel event (default: 0).
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// clickCount - Number of times the mouse button was clicked (default: 0).
func (c *Input) EmulateTouchFromMouseEvent(ctx context.Context, theType string, x int, y int, button string, timestamp float64, deltaX float64, deltaY float64, modifiers int, clickCount int) (*gcdmessage.ChromeResponse, error) {
	var v InputEmulateTouchFromMouseEventParams
	v.TheType = theType
	v.X = x
	v.Y = y
	v.Button = button
	v.Timestamp = timestamp
	v.DeltaX = deltaX
	v.DeltaY = deltaY
	v.Modifiers = modifiers
	v.ClickCount = clickCount
	return c.EmulateTouchFromMouseEventWithParams(ctx, &v)
}

type InputSetIgnoreInputEventsParams struct {
	// Ignores input events processing when set to true.
	Ignore bool `json:"ignore"`
}

// SetIgnoreInputEventsWithParams - Ignores input events (useful while auditing page).
func (c *Input) SetIgnoreInputEventsWithParams(ctx context.Context, v *InputSetIgnoreInputEventsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.setIgnoreInputEvents", Params: v})
}

// SetIgnoreInputEvents - Ignores input events (useful while auditing page).
// ignore - Ignores input events processing when set to true.
func (c *Input) SetIgnoreInputEvents(ctx context.Context, ignore bool) (*gcdmessage.ChromeResponse, error) {
	var v InputSetIgnoreInputEventsParams
	v.Ignore = ignore
	return c.SetIgnoreInputEventsWithParams(ctx, &v)
}

type InputSetInterceptDragsParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetInterceptDragsWithParams - Prevents default drag and drop behavior and instead emits `Input.dragIntercepted` events. Drag and drop behavior can be directly controlled via `Input.dispatchDragEvent`.
func (c *Input) SetInterceptDragsWithParams(ctx context.Context, v *InputSetInterceptDragsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.setInterceptDrags", Params: v})
}

// SetInterceptDrags - Prevents default drag and drop behavior and instead emits `Input.dragIntercepted` events. Drag and drop behavior can be directly controlled via `Input.dispatchDragEvent`.
// enabled -
func (c *Input) SetInterceptDrags(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v InputSetInterceptDragsParams
	v.Enabled = enabled
	return c.SetInterceptDragsWithParams(ctx, &v)
}

type InputSynthesizePinchGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor float64 `json:"scaleFactor"`
	// Relative pointer speed in pixels per second (default: 800).
	RelativeSpeed int `json:"relativeSpeed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
}

// SynthesizePinchGestureWithParams - Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizePinchGestureWithParams(ctx context.Context, v *InputSynthesizePinchGestureParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizePinchGesture", Params: v})
}

// SynthesizePinchGesture - Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// scaleFactor - Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
// relativeSpeed - Relative pointer speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
func (c *Input) SynthesizePinchGesture(ctx context.Context, x float64, y float64, scaleFactor float64, relativeSpeed int, gestureSourceType string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizePinchGestureParams
	v.X = x
	v.Y = y
	v.ScaleFactor = scaleFactor
	v.RelativeSpeed = relativeSpeed
	v.GestureSourceType = gestureSourceType
	return c.SynthesizePinchGestureWithParams(ctx, &v)
}

type InputSynthesizeScrollGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// The distance to scroll along the X axis (positive to scroll left).
	XDistance float64 `json:"xDistance,omitempty"`
	// The distance to scroll along the Y axis (positive to scroll up).
	YDistance float64 `json:"yDistance,omitempty"`
	// The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	XOverscroll float64 `json:"xOverscroll,omitempty"`
	// The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	YOverscroll float64 `json:"yOverscroll,omitempty"`
	// Prevent fling (default: true).
	PreventFling bool `json:"preventFling,omitempty"`
	// Swipe speed in pixels per second (default: 800).
	Speed int `json:"speed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
	// The number of times to repeat the gesture (default: 0).
	RepeatCount int `json:"repeatCount,omitempty"`
	// The number of milliseconds delay between each repeat. (default: 250).
	RepeatDelayMs int `json:"repeatDelayMs,omitempty"`
	// The name of the interaction markers to generate, if not empty (default: "").
	InteractionMarkerName string `json:"interactionMarkerName,omitempty"`
}

// SynthesizeScrollGestureWithParams - Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizeScrollGestureWithParams(ctx context.Context, v *InputSynthesizeScrollGestureParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizeScrollGesture", Params: v})
}

// SynthesizeScrollGesture - Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// xDistance - The distance to scroll along the X axis (positive to scroll left).
// yDistance - The distance to scroll along the Y axis (positive to scroll up).
// xOverscroll - The number of additional pixels to scroll back along the X axis, in addition to the given distance.
// yOverscroll - The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
// preventFling - Prevent fling (default: true).
// speed - Swipe speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
// repeatCount - The number of times to repeat the gesture (default: 0).
// repeatDelayMs - The number of milliseconds delay between each repeat. (default: 250).
// interactionMarkerName - The name of the interaction markers to generate, if not empty (default: "").
func (c *Input) SynthesizeScrollGesture(ctx context.Context, x float64, y float64, xDistance float64, yDistance float64, xOverscroll float64, yOverscroll float64, preventFling bool, speed int, gestureSourceType string, repeatCount int, repeatDelayMs int, interactionMarkerName string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizeScrollGestureParams
	v.X = x
	v.Y = y
	v.XDistance = xDistance
	v.YDistance = yDistance
	v.XOverscroll = xOverscroll
	v.YOverscroll = yOverscroll
	v.PreventFling = preventFling
	v.Speed = speed
	v.GestureSourceType = gestureSourceType
	v.RepeatCount = repeatCount
	v.RepeatDelayMs = repeatDelayMs
	v.InteractionMarkerName = interactionMarkerName
	return c.SynthesizeScrollGestureWithParams(ctx, &v)
}

type InputSynthesizeTapGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// Duration between touchdown and touchup events in ms (default: 50).
	Duration int `json:"duration,omitempty"`
	// Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	TapCount int `json:"tapCount,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
}

// SynthesizeTapGestureWithParams - Synthesizes a tap gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizeTapGestureWithParams(ctx context.Context, v *InputSynthesizeTapGestureParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizeTapGesture", Params: v})
}

// SynthesizeTapGesture - Synthesizes a tap gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// duration - Duration between touchdown and touchup events in ms (default: 50).
// tapCount - Number of times to perform the tap (e.g. 2 for double tap, default: 1).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
func (c *Input) SynthesizeTapGesture(ctx context.Context, x float64, y float64, duration int, tapCount int, gestureSourceType string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizeTapGestureParams
	v.X = x
	v.Y = y
	v.Duration = duration
	v.TapCount = tapCount
	v.GestureSourceType = gestureSourceType
	return c.SynthesizeTapGestureWithParams(ctx, &v)
}
