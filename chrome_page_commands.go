// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Page commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Page() *ChromePage {
	if c.page == nil {
		c.page = newChromePage(c)
	}
	return c.page
}

type ChromePage struct {
	target *ChromeTarget
}

func newChromePage(target *ChromeTarget) *ChromePage {
	c := &ChromePage{target: target}
	return c
}

// Enables page domain notifications.
func (c *ChromePage) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.enable"})
}

// Disables page domain notifications.
func (c *ChromePage) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.disable"})
}

// Clears the overriden device metrics.
func (c *ChromePage) ClearDeviceMetricsOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearDeviceMetricsOverride"})
}

// Clears the overriden Geolocation Position and Error.
func (c *ChromePage) ClearGeolocationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearGeolocationOverride"})
}

// Clears the overridden Device Orientation.
func (c *ChromePage) ClearDeviceOrientationOverride() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.clearDeviceOrientationOverride"})
}

// Stops sending each frame in the <code>screencastFrame</code>.
func (c *ChromePage) StopScreencast() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.stopScreencast"})
}

// removeScriptToEvaluateOnLoad -
// identifier -
func (c *ChromePage) RemoveScriptToEvaluateOnLoad(identifier *types.ChromePageScriptIdentifier) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["identifier"] = identifier
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.removeScriptToEvaluateOnLoad", Params: paramRequest})
}

// reload - Reloads given page optionally ignoring the cache.
// ignoreCache - If true, browser cache is ignored (as if the user pressed Shift+refresh).
// scriptToEvaluateOnLoad - If set, the script will be injected into all frames of the inspected page after reload.
func (c *ChromePage) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["ignoreCache"] = ignoreCache
	paramRequest["scriptToEvaluateOnLoad"] = scriptToEvaluateOnLoad
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.reload", Params: paramRequest})
}

// navigateToHistoryEntry - Navigates current page to the given history entry.
// entryId - Unique id of the entry to navigate to.
func (c *ChromePage) NavigateToHistoryEntry(entryId int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["entryId"] = entryId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.navigateToHistoryEntry", Params: paramRequest})
}

// deleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *ChromePage) DeleteCookie(cookieName string, url string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.deleteCookie", Params: paramRequest})
}

// setDocumentContent - Sets given markup as the document's HTML.
// frameId - Frame id to set HTML for.
// html - HTML content to set.
func (c *ChromePage) SetDocumentContent(frameId *types.ChromePageFrameId, html string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["html"] = html
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDocumentContent", Params: paramRequest})
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
func (c *ChromePage) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int) (*ChromeResponse, error) {
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
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDeviceMetricsOverride", Params: paramRequest})
}

// setGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *ChromePage) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setGeolocationOverride", Params: paramRequest})
}

// setDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *ChromePage) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["alpha"] = alpha
	paramRequest["beta"] = beta
	paramRequest["gamma"] = gamma
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setDeviceOrientationOverride", Params: paramRequest})
}

// setTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *ChromePage) SetTouchEmulationEnabled(enabled bool, configuration string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["enabled"] = enabled
	paramRequest["configuration"] = configuration
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setTouchEmulationEnabled", Params: paramRequest})
}

// startScreencast - Starts sending each frame using the <code>screencastFrame</code> event.
// format - Image compression format.
// quality - Compression quality from range [0..100].
// maxWidth - Maximum screenshot width.
// maxHeight - Maximum screenshot height.
func (c *ChromePage) StartScreencast(format string, quality int, maxWidth int, maxHeight int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["format"] = format
	paramRequest["quality"] = quality
	paramRequest["maxWidth"] = maxWidth
	paramRequest["maxHeight"] = maxHeight
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.startScreencast", Params: paramRequest})
}

// screencastFrameAck - Acknowledges that a screencast frame has been received by the frontend.
// frameNumber - Frame number.
func (c *ChromePage) ScreencastFrameAck(frameNumber int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameNumber"] = frameNumber
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.screencastFrameAck", Params: paramRequest})
}

// handleJavaScriptDialog - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
// accept - Whether to accept or dismiss the dialog.
// promptText - The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (c *ChromePage) HandleJavaScriptDialog(accept bool, promptText string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["accept"] = accept
	paramRequest["promptText"] = promptText
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.handleJavaScriptDialog", Params: paramRequest})
}

// setShowViewportSizeOnResize - Paints viewport size upon main frame resize.
// show - Whether to paint size or not.
// showGrid - Whether to paint grid as well.
func (c *ChromePage) SetShowViewportSizeOnResize(show bool, showGrid bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["show"] = show
	paramRequest["showGrid"] = showGrid
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setShowViewportSizeOnResize", Params: paramRequest})
}

// setColorPickerEnabled - Shows / hides color picker
// enabled - Shows / hides color picker
func (c *ChromePage) SetColorPickerEnabled(enabled bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setColorPickerEnabled", Params: paramRequest})
}

// setOverlayMessage - Sets overlay message.
// message - Overlay message to display when paused in debugger.
func (c *ChromePage) SetOverlayMessage(message string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["message"] = message
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.setOverlayMessage", Params: paramRequest})
}

// getNavigationHistory - Returns navigation history for the current page.
// Returns -
// Index of the current navigation history entry.
// Array of navigation history entries.
func (c *ChromePage) GetNavigationHistory() (float64, []*types.ChromePageNavigationEntry, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getNavigationHistory"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CurrentIndex float64
			Entries      []*types.ChromePageNavigationEntry
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, nil, err
	}

	return chromeData.Result.CurrentIndex, chromeData.Result.Entries, nil
}

// getCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -
// Array of cookie objects.
func (c *ChromePage) GetCookies() ([]*types.ChromeNetworkCookie, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Cookies []*types.ChromeNetworkCookie
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Cookies, nil
}

// getResourceTree - Returns present frame / resource tree structure.
// Returns -
// Present frame / resource tree structure.
func (c *ChromePage) GetResourceTree() (*types.ChromePageFrameResourceTree, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getResourceTree"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameTree *types.ChromePageFrameResourceTree
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.FrameTree, nil
}

// captureScreenshot - Capture page screenshot.
// Returns -
// Base64-encoded image data (PNG).
func (c *ChromePage) CaptureScreenshot() (string, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.captureScreenshot"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.Data, nil
}

// canScreencast - Tells whether screencast is supported.
// Returns -
// True if screencast is supported.
func (c *ChromePage) CanScreencast() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.canScreencast"})
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

// addScriptToEvaluateOnLoad -
// Returns -
// Identifier of the added script.
func (c *ChromePage) AddScriptToEvaluateOnLoad(scriptSource string) (*types.ChromePageScriptIdentifier, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scriptSource"] = scriptSource
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.addScriptToEvaluateOnLoad", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Identifier *types.ChromePageScriptIdentifier
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Identifier, nil
}

// navigate - Navigates current page to the given URL.
// Returns -
// Frame id that will be navigated.
func (c *ChromePage) Navigate(url string) (*types.ChromePageFrameId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.navigate", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameId *types.ChromePageFrameId
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.FrameId, nil
}

// getResourceContent - Returns content of the given resource.
// Returns -
// Resource content.
// True, if content was served as base64.
func (c *ChromePage) GetResourceContent(frameId *types.ChromePageFrameId, url string) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.getResourceContent", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Content       string
			Base64Encoded bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.Content, chromeData.Result.Base64Encoded, nil
}

// searchInResource - Searches for given string in resource content.
// Returns -
// List of search matches.
func (c *ChromePage) SearchInResource(frameId *types.ChromePageFrameId, url string, query string, caseSensitive bool, isRegex bool) ([]*types.ChromeDebuggerSearchMatch, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	paramRequest["query"] = query
	paramRequest["caseSensitive"] = caseSensitive
	paramRequest["isRegex"] = isRegex
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Page.searchInResource", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result []*types.ChromeDebuggerSearchMatch
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Result, nil
}
