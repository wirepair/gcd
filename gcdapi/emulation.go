// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Emulation functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Screen orientation.
type EmulationScreenOrientation struct {
	Type  string `json:"type"`  // Orientation type.
	Angle int    `json:"angle"` // Orientation angle.
}

type Emulation struct {
	target gcdmessage.ChromeTargeter
}

func NewEmulation(target gcdmessage.ChromeTargeter) *Emulation {
	c := &Emulation{target: target}
	return c
}

// SetDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
// scale - Scale to apply to resulting view image. Ignored in |fitWindow| mode.
// offsetX - Not used.
// offsetY - Not used.
// screenWidth - Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenHeight - Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionX - Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionY - Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenOrientation - Screen orientation override.
func (c *Emulation) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int, screenOrientation *EmulationScreenOrientation) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 13)
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["deviceScaleFactor"] = deviceScaleFactor
	paramRequest["mobile"] = mobile
	paramRequest["fitWindow"] = fitWindow
	paramRequest["scale"] = scale
	paramRequest["offsetX"] = offsetX
	paramRequest["offsetY"] = offsetY
	paramRequest["screenWidth"] = screenWidth
	paramRequest["screenHeight"] = screenHeight
	paramRequest["positionX"] = positionX
	paramRequest["positionY"] = positionY
	paramRequest["screenOrientation"] = screenOrientation
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDeviceMetricsOverride", Params: paramRequest})
}

// Clears the overriden device metrics.
func (c *Emulation) ClearDeviceMetricsOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearDeviceMetricsOverride"})
}

// ForceViewport - Overrides the visible area of the page. The change is hidden from the page, i.e. the observable scroll position and page scale does not change. In effect, the command moves the specified area of the page into the top-left corner of the frame.
// x - X coordinate of top-left corner of the area (CSS pixels).
// y - Y coordinate of top-left corner of the area (CSS pixels).
// scale - Scale to apply to the area (relative to a page scale of 1.0).
func (c *Emulation) ForceViewport(x float64, y float64, scale float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["scale"] = scale
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.forceViewport", Params: paramRequest})
}

// Resets the visible area of the page to the original viewport, undoing any effects of the <code>forceViewport</code> command.
func (c *Emulation) ResetViewport() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.resetViewport"})
}

// Requests that page scale factor is reset to initial values.
func (c *Emulation) ResetPageScaleFactor() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.resetPageScaleFactor"})
}

// SetPageScaleFactor - Sets a specified page scale factor.
// pageScaleFactor - Page scale factor.
func (c *Emulation) SetPageScaleFactor(pageScaleFactor float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["pageScaleFactor"] = pageScaleFactor
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setPageScaleFactor", Params: paramRequest})
}

// SetVisibleSize - Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
// width - Frame width (DIP).
// height - Frame height (DIP).
func (c *Emulation) SetVisibleSize(width int, height int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["width"] = width
	paramRequest["height"] = height
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVisibleSize", Params: paramRequest})
}

// SetScriptExecutionDisabled - Switches script execution in the page.
// value - Whether script execution should be disabled in the page.
func (c *Emulation) SetScriptExecutionDisabled(value bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setScriptExecutionDisabled", Params: paramRequest})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Emulation) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setGeolocationOverride", Params: paramRequest})
}

// Clears the overriden Geolocation Position and Error.
func (c *Emulation) ClearGeolocationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearGeolocationOverride"})
}

// SetTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Emulation) SetTouchEmulationEnabled(enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["enabled"] = enabled
	paramRequest["configuration"] = configuration
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setTouchEmulationEnabled", Params: paramRequest})
}

// SetEmulatedMedia - Emulates the given media for CSS media queries.
// media - Media type to emulate. Empty string disables the override.
func (c *Emulation) SetEmulatedMedia(media string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["media"] = media
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setEmulatedMedia", Params: paramRequest})
}

// SetCPUThrottlingRate - Enables CPU throttling to emulate slow CPUs.
// rate - Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
func (c *Emulation) SetCPUThrottlingRate(rate float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["rate"] = rate
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setCPUThrottlingRate", Params: paramRequest})
}

// CanEmulate - Tells whether emulation is supported.
// Returns -  result - True if emulation is supported.
func (c *Emulation) CanEmulate() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.canEmulate"})
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

// SetVirtualTimePolicy - Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
// policy -  enum values: advance, pause, pauseIfNetworkFetchesPending
// budget - If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
func (c *Emulation) SetVirtualTimePolicy(policy string, budget int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["policy"] = policy
	paramRequest["budget"] = budget
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVirtualTimePolicy", Params: paramRequest})
}

// SetDefaultBackgroundColorOverride - Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
// color - RGBA of the default background color. If not specified, any existing override will be cleared.
func (c *Emulation) SetDefaultBackgroundColorOverride(color *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["color"] = color
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDefaultBackgroundColorOverride", Params: paramRequest})
}
