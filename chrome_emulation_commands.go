// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Emulation commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Emulation() *ChromeEmulation {
	if c.emulation == nil {
		c.emulation = newChromeEmulation(c)
	}
	return c.emulation
}

type ChromeEmulation struct {
	target *ChromeTarget
}

func newChromeEmulation(target *ChromeTarget) *ChromeEmulation {
	c := &ChromeEmulation{target: target}
	return c
}

// Clears the overriden device metrics.
func (c *ChromeEmulation) ClearDeviceMetricsOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.clearDeviceMetricsOverride"})
}

// Requests that scroll offsets and page scale factor are reset to initial values.
func (c *ChromeEmulation) ResetScrollAndPageScaleFactor() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.resetScrollAndPageScaleFactor"})
}

// Clears the overriden Geolocation Position and Error.
func (c *ChromeEmulation) ClearGeolocationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.clearGeolocationOverride"})
}

// setDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
// scale - Scale to apply to resulting view image. Ignored in |fitWindow| mode.
// offsetX - X offset to shift resulting view image by. Ignored in |fitWindow| mode.
// offsetY - Y offset to shift resulting view image by. Ignored in |fitWindow| mode.
// screenWidth - Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenHeight - Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionX - Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionY - Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
func (c *ChromeEmulation) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 12)
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
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setDeviceMetricsOverride", Params: paramRequest})
}

// setPageScaleFactor - Sets a specified page scale factor.
// pageScaleFactor - Page scale factor.
func (c *ChromeEmulation) SetPageScaleFactor(pageScaleFactor float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["pageScaleFactor"] = pageScaleFactor
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setPageScaleFactor", Params: paramRequest})
}

// setScriptExecutionDisabled - Switches script execution in the page.
// value - Whether script execution should be disabled in the page.
func (c *ChromeEmulation) SetScriptExecutionDisabled(value bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setScriptExecutionDisabled", Params: paramRequest})
}

// setGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *ChromeEmulation) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setGeolocationOverride", Params: paramRequest})
}

// setTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *ChromeEmulation) SetTouchEmulationEnabled(enabled bool, configuration string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["enabled"] = enabled
	paramRequest["configuration"] = configuration
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setTouchEmulationEnabled", Params: paramRequest})
}

// setEmulatedMedia - Emulates the given media for CSS media queries.
// media - Media type to emulate. Empty string disables the override.
func (c *ChromeEmulation) SetEmulatedMedia(media string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["media"] = media
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.setEmulatedMedia", Params: paramRequest})
}

// canEmulate - Tells whether emulation is supported.
// Returns -
// True if emulation is supported.
func (c *ChromeEmulation) CanEmulate() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Emulation.canEmulate"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}
