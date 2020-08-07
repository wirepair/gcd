// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Emulation functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Screen orientation.
type EmulationScreenOrientation struct {
	Type  string `json:"type"`  // Orientation type.
	Angle int    `json:"angle"` // Orientation angle.
}

// No Description.
type EmulationDisplayFeature struct {
	Orientation string `json:"orientation"` // Orientation of a display feature in relation to screen
	Offset      int    `json:"offset"`      // The offset from the screen origin in either the x (for vertical orientation) or y (for horizontal orientation) direction.
	MaskLength  int    `json:"maskLength"`  // A display feature may mask content such that it is not physically displayed - this length along with the offset describes this area. A display feature that only splits content will have a 0 mask_length.
}

// No Description.
type EmulationMediaFeature struct {
	Name  string `json:"name"`  //
	Value string `json:"value"` //
}

// Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
type EmulationUserAgentBrandVersion struct {
	Brand   string `json:"brand"`   //
	Version string `json:"version"` //
}

// Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
type EmulationUserAgentMetadata struct {
	Brands          []*EmulationUserAgentBrandVersion `json:"brands"`          //
	FullVersion     string                            `json:"fullVersion"`     //
	Platform        string                            `json:"platform"`        //
	PlatformVersion string                            `json:"platformVersion"` //
	Architecture    string                            `json:"architecture"`    //
	Model           string                            `json:"model"`           //
	Mobile          bool                              `json:"mobile"`          //
}

type Emulation struct {
	target gcdmessage.ChromeTargeter
}

func NewEmulation(target gcdmessage.ChromeTargeter) *Emulation {
	c := &Emulation{target: target}
	return c
}

// CanEmulate - Tells whether emulation is supported.
// Returns -  result - True if emulation is supported.
func (c *Emulation) CanEmulate(ctx context.Context) (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.canEmulate"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears the overriden device metrics.
func (c *Emulation) ClearDeviceMetricsOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearDeviceMetricsOverride"})
}

// Clears the overriden Geolocation Position and Error.
func (c *Emulation) ClearGeolocationOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearGeolocationOverride"})
}

// Requests that page scale factor is reset to initial values.
func (c *Emulation) ResetPageScaleFactor(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.resetPageScaleFactor"})
}

type EmulationSetFocusEmulationEnabledParams struct {
	// Whether to enable to disable focus emulation.
	Enabled bool `json:"enabled"`
}

// SetFocusEmulationEnabledWithParams - Enables or disables simulating a focused and active page.
func (c *Emulation) SetFocusEmulationEnabledWithParams(ctx context.Context, v *EmulationSetFocusEmulationEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setFocusEmulationEnabled", Params: v})
}

// SetFocusEmulationEnabled - Enables or disables simulating a focused and active page.
// enabled - Whether to enable to disable focus emulation.
func (c *Emulation) SetFocusEmulationEnabled(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetFocusEmulationEnabledParams
	v.Enabled = enabled
	return c.SetFocusEmulationEnabledWithParams(ctx, &v)
}

type EmulationSetCPUThrottlingRateParams struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float64 `json:"rate"`
}

// SetCPUThrottlingRateWithParams - Enables CPU throttling to emulate slow CPUs.
func (c *Emulation) SetCPUThrottlingRateWithParams(ctx context.Context, v *EmulationSetCPUThrottlingRateParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setCPUThrottlingRate", Params: v})
}

// SetCPUThrottlingRate - Enables CPU throttling to emulate slow CPUs.
// rate - Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
func (c *Emulation) SetCPUThrottlingRate(ctx context.Context, rate float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetCPUThrottlingRateParams
	v.Rate = rate
	return c.SetCPUThrottlingRateWithParams(ctx, &v)
}

type EmulationSetDefaultBackgroundColorOverrideParams struct {
	// RGBA of the default background color. If not specified, any existing override will be cleared.
	Color *DOMRGBA `json:"color,omitempty"`
}

// SetDefaultBackgroundColorOverrideWithParams - Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
func (c *Emulation) SetDefaultBackgroundColorOverrideWithParams(ctx context.Context, v *EmulationSetDefaultBackgroundColorOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDefaultBackgroundColorOverride", Params: v})
}

// SetDefaultBackgroundColorOverride - Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
// color - RGBA of the default background color. If not specified, any existing override will be cleared.
func (c *Emulation) SetDefaultBackgroundColorOverride(ctx context.Context, color *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetDefaultBackgroundColorOverrideParams
	v.Color = color
	return c.SetDefaultBackgroundColorOverrideWithParams(ctx, &v)
}

type EmulationSetDeviceMetricsOverrideParams struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`
	// Scale to apply to resulting view image.
	Scale float64 `json:"scale,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth int `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight int `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX int `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY int `json:"positionY,omitempty"`
	// Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize bool `json:"dontSetVisibleSize,omitempty"`
	// Screen orientation override.
	ScreenOrientation *EmulationScreenOrientation `json:"screenOrientation,omitempty"`
	// If set, the visible area of the page will be overridden to this viewport. This viewport change is not observed by the page, e.g. viewport-relative elements do not change positions.
	Viewport *PageViewport `json:"viewport,omitempty"`
	// If set, the display feature of a multi-segment screen. If not set, multi-segment support is turned-off.
	DisplayFeature *EmulationDisplayFeature `json:"displayFeature,omitempty"`
}

// SetDeviceMetricsOverrideWithParams - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (c *Emulation) SetDeviceMetricsOverrideWithParams(ctx context.Context, v *EmulationSetDeviceMetricsOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDeviceMetricsOverride", Params: v})
}

// SetDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// scale - Scale to apply to resulting view image.
// screenWidth - Overriding screen width value in pixels (minimum 0, maximum 10000000).
// screenHeight - Overriding screen height value in pixels (minimum 0, maximum 10000000).
// positionX - Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
// positionY - Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
// dontSetVisibleSize - Do not set visible view size, rely upon explicit setVisibleSize call.
// screenOrientation - Screen orientation override.
// viewport - If set, the visible area of the page will be overridden to this viewport. This viewport change is not observed by the page, e.g. viewport-relative elements do not change positions.
// displayFeature - If set, the display feature of a multi-segment screen. If not set, multi-segment support is turned-off.
func (c *Emulation) SetDeviceMetricsOverride(ctx context.Context, width int, height int, deviceScaleFactor float64, mobile bool, scale float64, screenWidth int, screenHeight int, positionX int, positionY int, dontSetVisibleSize bool, screenOrientation *EmulationScreenOrientation, viewport *PageViewport, displayFeature *EmulationDisplayFeature) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetDeviceMetricsOverrideParams
	v.Width = width
	v.Height = height
	v.DeviceScaleFactor = deviceScaleFactor
	v.Mobile = mobile
	v.Scale = scale
	v.ScreenWidth = screenWidth
	v.ScreenHeight = screenHeight
	v.PositionX = positionX
	v.PositionY = positionY
	v.DontSetVisibleSize = dontSetVisibleSize
	v.ScreenOrientation = screenOrientation
	v.Viewport = viewport
	v.DisplayFeature = displayFeature
	return c.SetDeviceMetricsOverrideWithParams(ctx, &v)
}

type EmulationSetScrollbarsHiddenParams struct {
	// Whether scrollbars should be always hidden.
	Hidden bool `json:"hidden"`
}

// SetScrollbarsHiddenWithParams -
func (c *Emulation) SetScrollbarsHiddenWithParams(ctx context.Context, v *EmulationSetScrollbarsHiddenParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setScrollbarsHidden", Params: v})
}

// SetScrollbarsHidden -
// hidden - Whether scrollbars should be always hidden.
func (c *Emulation) SetScrollbarsHidden(ctx context.Context, hidden bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetScrollbarsHiddenParams
	v.Hidden = hidden
	return c.SetScrollbarsHiddenWithParams(ctx, &v)
}

type EmulationSetDocumentCookieDisabledParams struct {
	// Whether document.coookie API should be disabled.
	Disabled bool `json:"disabled"`
}

// SetDocumentCookieDisabledWithParams -
func (c *Emulation) SetDocumentCookieDisabledWithParams(ctx context.Context, v *EmulationSetDocumentCookieDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDocumentCookieDisabled", Params: v})
}

// SetDocumentCookieDisabled -
// disabled - Whether document.coookie API should be disabled.
func (c *Emulation) SetDocumentCookieDisabled(ctx context.Context, disabled bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetDocumentCookieDisabledParams
	v.Disabled = disabled
	return c.SetDocumentCookieDisabledWithParams(ctx, &v)
}

type EmulationSetEmitTouchEventsForMouseParams struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration string `json:"configuration,omitempty"`
}

// SetEmitTouchEventsForMouseWithParams -
func (c *Emulation) SetEmitTouchEventsForMouseWithParams(ctx context.Context, v *EmulationSetEmitTouchEventsForMouseParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setEmitTouchEventsForMouse", Params: v})
}

// SetEmitTouchEventsForMouse -
// enabled - Whether touch emulation based on mouse input should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Emulation) SetEmitTouchEventsForMouse(ctx context.Context, enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetEmitTouchEventsForMouseParams
	v.Enabled = enabled
	v.Configuration = configuration
	return c.SetEmitTouchEventsForMouseWithParams(ctx, &v)
}

type EmulationSetEmulatedMediaParams struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media,omitempty"`
	// Media features to emulate.
	Features []*EmulationMediaFeature `json:"features,omitempty"`
}

// SetEmulatedMediaWithParams - Emulates the given media type or media feature for CSS media queries.
func (c *Emulation) SetEmulatedMediaWithParams(ctx context.Context, v *EmulationSetEmulatedMediaParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setEmulatedMedia", Params: v})
}

// SetEmulatedMedia - Emulates the given media type or media feature for CSS media queries.
// media - Media type to emulate. Empty string disables the override.
// features - Media features to emulate.
func (c *Emulation) SetEmulatedMedia(ctx context.Context, media string, features []*EmulationMediaFeature) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetEmulatedMediaParams
	v.Media = media
	v.Features = features
	return c.SetEmulatedMediaWithParams(ctx, &v)
}

type EmulationSetEmulatedVisionDeficiencyParams struct {
	// Vision deficiency to emulate.
	TheType string `json:"type"`
}

// SetEmulatedVisionDeficiencyWithParams - Emulates the given vision deficiency.
func (c *Emulation) SetEmulatedVisionDeficiencyWithParams(ctx context.Context, v *EmulationSetEmulatedVisionDeficiencyParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setEmulatedVisionDeficiency", Params: v})
}

// SetEmulatedVisionDeficiency - Emulates the given vision deficiency.
// type - Vision deficiency to emulate.
func (c *Emulation) SetEmulatedVisionDeficiency(ctx context.Context, theType string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetEmulatedVisionDeficiencyParams
	v.TheType = theType
	return c.SetEmulatedVisionDeficiencyWithParams(ctx, &v)
}

type EmulationSetGeolocationOverrideParams struct {
	// Mock latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy float64 `json:"accuracy,omitempty"`
}

// SetGeolocationOverrideWithParams - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (c *Emulation) SetGeolocationOverrideWithParams(ctx context.Context, v *EmulationSetGeolocationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setGeolocationOverride", Params: v})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Emulation) SetGeolocationOverride(ctx context.Context, latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetGeolocationOverrideParams
	v.Latitude = latitude
	v.Longitude = longitude
	v.Accuracy = accuracy
	return c.SetGeolocationOverrideWithParams(ctx, &v)
}

type EmulationSetIdleOverrideParams struct {
	// Mock isUserActive
	IsUserActive bool `json:"isUserActive"`
	// Mock isScreenUnlocked
	IsScreenUnlocked bool `json:"isScreenUnlocked"`
}

// SetIdleOverrideWithParams - Overrides the Idle state.
func (c *Emulation) SetIdleOverrideWithParams(ctx context.Context, v *EmulationSetIdleOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setIdleOverride", Params: v})
}

// SetIdleOverride - Overrides the Idle state.
// isUserActive - Mock isUserActive
// isScreenUnlocked - Mock isScreenUnlocked
func (c *Emulation) SetIdleOverride(ctx context.Context, isUserActive bool, isScreenUnlocked bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetIdleOverrideParams
	v.IsUserActive = isUserActive
	v.IsScreenUnlocked = isScreenUnlocked
	return c.SetIdleOverrideWithParams(ctx, &v)
}

// Clears Idle state overrides.
func (c *Emulation) ClearIdleOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearIdleOverride"})
}

type EmulationSetNavigatorOverridesParams struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

// SetNavigatorOverridesWithParams - Overrides value returned by the javascript navigator object.
func (c *Emulation) SetNavigatorOverridesWithParams(ctx context.Context, v *EmulationSetNavigatorOverridesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setNavigatorOverrides", Params: v})
}

// SetNavigatorOverrides - Overrides value returned by the javascript navigator object.
// platform - The platform navigator.platform should return.
func (c *Emulation) SetNavigatorOverrides(ctx context.Context, platform string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetNavigatorOverridesParams
	v.Platform = platform
	return c.SetNavigatorOverridesWithParams(ctx, &v)
}

type EmulationSetPageScaleFactorParams struct {
	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

// SetPageScaleFactorWithParams - Sets a specified page scale factor.
func (c *Emulation) SetPageScaleFactorWithParams(ctx context.Context, v *EmulationSetPageScaleFactorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setPageScaleFactor", Params: v})
}

// SetPageScaleFactor - Sets a specified page scale factor.
// pageScaleFactor - Page scale factor.
func (c *Emulation) SetPageScaleFactor(ctx context.Context, pageScaleFactor float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetPageScaleFactorParams
	v.PageScaleFactor = pageScaleFactor
	return c.SetPageScaleFactorWithParams(ctx, &v)
}

type EmulationSetScriptExecutionDisabledParams struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// SetScriptExecutionDisabledWithParams - Switches script execution in the page.
func (c *Emulation) SetScriptExecutionDisabledWithParams(ctx context.Context, v *EmulationSetScriptExecutionDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setScriptExecutionDisabled", Params: v})
}

// SetScriptExecutionDisabled - Switches script execution in the page.
// value - Whether script execution should be disabled in the page.
func (c *Emulation) SetScriptExecutionDisabled(ctx context.Context, value bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetScriptExecutionDisabledParams
	v.Value = value
	return c.SetScriptExecutionDisabledWithParams(ctx, &v)
}

type EmulationSetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints int `json:"maxTouchPoints,omitempty"`
}

// SetTouchEmulationEnabledWithParams - Enables touch on platforms which do not support them.
func (c *Emulation) SetTouchEmulationEnabledWithParams(ctx context.Context, v *EmulationSetTouchEmulationEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setTouchEmulationEnabled", Params: v})
}

// SetTouchEmulationEnabled - Enables touch on platforms which do not support them.
// enabled - Whether the touch event emulation should be enabled.
// maxTouchPoints - Maximum touch points supported. Defaults to one.
func (c *Emulation) SetTouchEmulationEnabled(ctx context.Context, enabled bool, maxTouchPoints int) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetTouchEmulationEnabledParams
	v.Enabled = enabled
	v.MaxTouchPoints = maxTouchPoints
	return c.SetTouchEmulationEnabledWithParams(ctx, &v)
}

type EmulationSetVirtualTimePolicyParams struct {
	//  enum values: advance, pause, pauseIfNetworkFetchesPending
	Policy string `json:"policy"`
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	Budget float64 `json:"budget,omitempty"`
	// If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
	// If set the virtual time policy change should be deferred until any frame starts navigating. Note any previous deferred policy change is superseded.
	WaitForNavigation bool `json:"waitForNavigation,omitempty"`
	// If set, base::Time::Now will be overriden to initially return this value.
	InitialVirtualTime float64 `json:"initialVirtualTime,omitempty"`
}

// SetVirtualTimePolicyWithParams - Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
// Returns -  virtualTimeTicksBase - Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
func (c *Emulation) SetVirtualTimePolicyWithParams(ctx context.Context, v *EmulationSetVirtualTimePolicyParams) (float64, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVirtualTimePolicy", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			VirtualTimeTicksBase float64
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

	return chromeData.Result.VirtualTimeTicksBase, nil
}

// SetVirtualTimePolicy - Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
// policy -  enum values: advance, pause, pauseIfNetworkFetchesPending
// budget - If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
// maxVirtualTimeTaskStarvationCount - If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
// waitForNavigation - If set the virtual time policy change should be deferred until any frame starts navigating. Note any previous deferred policy change is superseded.
// initialVirtualTime - If set, base::Time::Now will be overriden to initially return this value.
// Returns -  virtualTimeTicksBase - Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
func (c *Emulation) SetVirtualTimePolicy(ctx context.Context, policy string, budget float64, maxVirtualTimeTaskStarvationCount int, waitForNavigation bool, initialVirtualTime float64) (float64, error) {
	var v EmulationSetVirtualTimePolicyParams
	v.Policy = policy
	v.Budget = budget
	v.MaxVirtualTimeTaskStarvationCount = maxVirtualTimeTaskStarvationCount
	v.WaitForNavigation = waitForNavigation
	v.InitialVirtualTime = initialVirtualTime
	return c.SetVirtualTimePolicyWithParams(ctx, &v)
}

type EmulationSetLocaleOverrideParams struct {
	// ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and restores default host system locale.
	Locale string `json:"locale,omitempty"`
}

// SetLocaleOverrideWithParams - Overrides default host system locale with the specified one.
func (c *Emulation) SetLocaleOverrideWithParams(ctx context.Context, v *EmulationSetLocaleOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setLocaleOverride", Params: v})
}

// SetLocaleOverride - Overrides default host system locale with the specified one.
// locale - ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and restores default host system locale.
func (c *Emulation) SetLocaleOverride(ctx context.Context, locale string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetLocaleOverrideParams
	v.Locale = locale
	return c.SetLocaleOverrideWithParams(ctx, &v)
}

type EmulationSetTimezoneOverrideParams struct {
	// The timezone identifier. If empty, disables the override and restores default host system timezone.
	TimezoneId string `json:"timezoneId"`
}

// SetTimezoneOverrideWithParams - Overrides default host system timezone with the specified one.
func (c *Emulation) SetTimezoneOverrideWithParams(ctx context.Context, v *EmulationSetTimezoneOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setTimezoneOverride", Params: v})
}

// SetTimezoneOverride - Overrides default host system timezone with the specified one.
// timezoneId - The timezone identifier. If empty, disables the override and restores default host system timezone.
func (c *Emulation) SetTimezoneOverride(ctx context.Context, timezoneId string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetTimezoneOverrideParams
	v.TimezoneId = timezoneId
	return c.SetTimezoneOverrideWithParams(ctx, &v)
}

type EmulationSetVisibleSizeParams struct {
	// Frame width (DIP).
	Width int `json:"width"`
	// Frame height (DIP).
	Height int `json:"height"`
}

// SetVisibleSizeWithParams - Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
func (c *Emulation) SetVisibleSizeWithParams(ctx context.Context, v *EmulationSetVisibleSizeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVisibleSize", Params: v})
}

// SetVisibleSize - Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
// width - Frame width (DIP).
// height - Frame height (DIP).
func (c *Emulation) SetVisibleSize(ctx context.Context, width int, height int) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetVisibleSizeParams
	v.Width = width
	v.Height = height
	return c.SetVisibleSizeWithParams(ctx, &v)
}

type EmulationSetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
	// Browser langugage to emulate.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	// The platform navigator.platform should return.
	Platform string `json:"platform,omitempty"`
	// To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
	UserAgentMetadata *EmulationUserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

// SetUserAgentOverrideWithParams - Allows overriding user agent with the given string.
func (c *Emulation) SetUserAgentOverrideWithParams(ctx context.Context, v *EmulationSetUserAgentOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setUserAgentOverride", Params: v})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
// acceptLanguage - Browser langugage to emulate.
// platform - The platform navigator.platform should return.
// userAgentMetadata - To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
func (c *Emulation) SetUserAgentOverride(ctx context.Context, userAgent string, acceptLanguage string, platform string, userAgentMetadata *EmulationUserAgentMetadata) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetUserAgentOverrideParams
	v.UserAgent = userAgent
	v.AcceptLanguage = acceptLanguage
	v.Platform = platform
	v.UserAgentMetadata = userAgentMetadata
	return c.SetUserAgentOverrideWithParams(ctx, &v)
}
