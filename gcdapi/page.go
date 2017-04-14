// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Page functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Information about the Frame on the page.
type PageFrame struct {
	Id             string `json:"id"`                 // Frame unique identifier.
	ParentId       string `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderId       string `json:"loaderId"`           // Identifier of the loader associated with this frame.
	Name           string `json:"name,omitempty"`     // Frame's name as specified in the tag.
	Url            string `json:"url"`                // Frame document's URL.
	SecurityOrigin string `json:"securityOrigin"`     // Frame document's security origin.
	MimeType       string `json:"mimeType"`           // Frame document's mimeType as determined by the browser.
}

// Information about the Resource on the page.
type PageFrameResource struct {
	Url          string  `json:"url"`                    // Resource URL.
	Type         string  `json:"type"`                   // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
	MimeType     string  `json:"mimeType"`               // Resource mimeType as determined by the browser.
	LastModified float64 `json:"lastModified,omitempty"` // last-modified timestamp as reported by server.
	ContentSize  float64 `json:"contentSize,omitempty"`  // Resource content size.
	Failed       bool    `json:"failed,omitempty"`       // True if the resource failed to load.
	Canceled     bool    `json:"canceled,omitempty"`     // True if the resource was canceled during loading.
}

// Information about the Frame hierarchy along with their cached resources.
type PageFrameResourceTree struct {
	Frame       *PageFrame               `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*PageFrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []*PageFrameResource     `json:"resources"`             // Information about frame resources.
}

// Navigation history entry.
type PageNavigationEntry struct {
	Id    int    `json:"id"`    // Unique id of the navigation history entry.
	Url   string `json:"url"`   // URL of the navigation history entry.
	Title string `json:"title"` // Title of the navigation history entry.
}

// Screencast frame metadata.
type PageScreencastFrameMetadata struct {
	OffsetTop       float64 `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64 `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64 `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64 `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64 `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64 `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       float64 `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// Error while paring app manifest.
type PageAppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int    `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int    `json:"line"`     // Error line.
	Column   int    `json:"column"`   // Error column.
}

// Layout viewport position and dimensions.
type PageLayoutViewport struct {
	PageX        int `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// Visual viewport position, dimensions, and scale.
type PageVisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

//
type PageDomContentEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

//
type PageLoadEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

// Fired when frame has been attached to its parent.
type PageFrameAttachedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId       string             `json:"frameId"`         // Id of the frame that has been attached.
		ParentFrameId string             `json:"parentFrameId"`   // Parent frame identifier.
		Stack         *RuntimeStackTrace `json:"stack,omitempty"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
	} `json:"Params,omitempty"`
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
type PageFrameNavigatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Frame *PageFrame `json:"frame"` // Frame object.
	} `json:"Params,omitempty"`
}

// Fired when frame has been detached from its parent.
type PageFrameDetachedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has been detached.
	} `json:"Params,omitempty"`
}

// Fired when frame has started loading.
type PageFrameStartedLoadingEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has started loading.
	} `json:"Params,omitempty"`
}

// Fired when frame has stopped loading.
type PageFrameStoppedLoadingEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has stopped loading.
	} `json:"Params,omitempty"`
}

// Fired when frame schedules a potential navigation.
type PageFrameScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string  `json:"frameId"` // Id of the frame that has scheduled a navigation.
		Delay   float64 `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	} `json:"Params,omitempty"`
}

// Fired when frame no longer has a scheduled navigation.
type PageFrameClearedScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to open.
type PageJavascriptDialogOpeningEvent struct {
	Method string `json:"method"`
	Params struct {
		Message string `json:"message"` // Message that will be displayed by the dialog.
		Type    string `json:"type"`    // Dialog type. enum values: alert, confirm, prompt, beforeunload
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been closed.
type PageJavascriptDialogClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		Result bool `json:"result"` // Whether dialog was confirmed.
	} `json:"Params,omitempty"`
}

// Compressed image data requested by the <code>startScreencast</code>.
type PageScreencastFrameEvent struct {
	Method string `json:"method"`
	Params struct {
		Data      string                       `json:"data"`      // Base64-encoded compressed image.
		Metadata  *PageScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
		SessionId int                          `json:"sessionId"` // Frame number.
	} `json:"Params,omitempty"`
}

// Fired when the page with currently enabled screencast was shown or hidden </code>.
type PageScreencastVisibilityChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		Visible bool `json:"visible"` // True if the page is visible.
	} `json:"Params,omitempty"`
}

// Fired when a color has been picked.
type PageColorPickedEvent struct {
	Method string `json:"method"`
	Params struct {
		Color *DOMRGBA `json:"color"` // RGBA of the picked color.
	} `json:"Params,omitempty"`
}

// Fired when a navigation is started if navigation throttles are enabled.  The navigation will be deferred until processNavigation is called.
type PageNavigationRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		IsInMainFrame bool   `json:"isInMainFrame"` // Whether the navigation is taking place in the main frame or in a subframe.
		IsRedirect    bool   `json:"isRedirect"`    // Whether the navigation has encountered a server redirect or not.
		NavigationId  int    `json:"navigationId"`  //
		Url           string `json:"url"`           // URL of requested navigation.
	} `json:"Params,omitempty"`
}

type Page struct {
	target gcdmessage.ChromeTargeter
}

func NewPage(target gcdmessage.ChromeTargeter) *Page {
	c := &Page{target: target}
	return c
}

// Enables page domain notifications.
func (c *Page) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.enable"})
}

// Disables page domain notifications.
func (c *Page) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.disable"})
}

// AddScriptToEvaluateOnLoad -
// scriptSource -
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnLoad(scriptSource string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scriptSource"] = scriptSource
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addScriptToEvaluateOnLoad", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Identifier string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Identifier, nil
}

// RemoveScriptToEvaluateOnLoad -
// identifier -
func (c *Page) RemoveScriptToEvaluateOnLoad(identifier string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["identifier"] = identifier
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.removeScriptToEvaluateOnLoad", Params: paramRequest})
}

// SetAutoAttachToCreatedPages - Controls whether browser will open a new inspector window for connected pages.
// autoAttach - If true, browser will open a new inspector window for every page created from this one.
func (c *Page) SetAutoAttachToCreatedPages(autoAttach bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["autoAttach"] = autoAttach
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setAutoAttachToCreatedPages", Params: paramRequest})
}

// Reload - Reloads given page optionally ignoring the cache.
// ignoreCache - If true, browser cache is ignored (as if the user pressed Shift+refresh).
// scriptToEvaluateOnLoad - If set, the script will be injected into all frames of the inspected page after reload.
func (c *Page) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["ignoreCache"] = ignoreCache
	paramRequest["scriptToEvaluateOnLoad"] = scriptToEvaluateOnLoad
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.reload", Params: paramRequest})
}

// Navigate - Navigates current page to the given URL.
// url - URL to navigate the page to.
// referrer - Referrer URL.
// Returns -  frameId - Frame id that will be navigated.
func (c *Page) Navigate(url string, referrer string) (string, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["url"] = url
	paramRequest["referrer"] = referrer
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigate", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			FrameId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.FrameId, nil
}

// Force the page stop all navigations and pending resource fetches.
func (c *Page) StopLoading() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopLoading"})
}

// GetNavigationHistory - Returns navigation history for the current page.
// Returns -  currentIndex - Index of the current navigation history entry. entries - Array of navigation history entries.
func (c *Page) GetNavigationHistory() (int, []*PageNavigationEntry, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getNavigationHistory"})
	if err != nil {
		return 0, nil, err
	}

	var chromeData struct {
		Result struct {
			CurrentIndex int
			Entries      []*PageNavigationEntry
		}
	}

	if resp == nil {
		return 0, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, nil, err
	}

	return chromeData.Result.CurrentIndex, chromeData.Result.Entries, nil
}

// NavigateToHistoryEntry - Navigates current page to the given history entry.
// entryId - Unique id of the entry to navigate to.
func (c *Page) NavigateToHistoryEntry(entryId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["entryId"] = entryId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigateToHistoryEntry", Params: paramRequest})
}

// GetCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Page) GetCookies() ([]*NetworkCookie, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getCookies"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Page) DeleteCookie(cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.deleteCookie", Params: paramRequest})
}

// GetResourceTree - Returns present frame / resource tree structure.
// Returns -  frameTree - Present frame / resource tree structure.
func (c *Page) GetResourceTree() (*PageFrameResourceTree, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceTree"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			FrameTree *PageFrameResourceTree
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

	return chromeData.Result.FrameTree, nil
}

// GetResourceContent - Returns content of the given resource.
// frameId - Frame id to get resource for.
// url - URL of the resource to get content for.
// Returns -  content - Resource content. base64Encoded - True, if content was served as base64.
func (c *Page) GetResourceContent(frameId string, url string) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceContent", Params: paramRequest})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Content       string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.Content, chromeData.Result.Base64Encoded, nil
}

// SearchInResource - Searches for given string in resource content.
// frameId - Frame id for resource to search in.
// url - URL of the resource to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Page) SearchInResource(frameId string, url string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	paramRequest["query"] = query
	paramRequest["caseSensitive"] = caseSensitive
	paramRequest["isRegex"] = isRegex
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.searchInResource", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
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

	return chromeData.Result.Result, nil
}

// SetDocumentContent - Sets given markup as the document's HTML.
// frameId - Frame id to set HTML for.
// html - HTML content to set.
func (c *Page) SetDocumentContent(frameId string, html string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["frameId"] = frameId
	paramRequest["html"] = html
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDocumentContent", Params: paramRequest})
}

// SetDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
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
// screenOrientation - Screen orientation override.
func (c *Page) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int, screenOrientation *EmulationScreenOrientation) (*gcdmessage.ChromeResponse, error) {
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
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceMetricsOverride", Params: paramRequest})
}

// Clears the overriden device metrics.
func (c *Page) ClearDeviceMetricsOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceMetricsOverride"})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Page) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["latitude"] = latitude
	paramRequest["longitude"] = longitude
	paramRequest["accuracy"] = accuracy
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setGeolocationOverride", Params: paramRequest})
}

// Clears the overriden Geolocation Position and Error.
func (c *Page) ClearGeolocationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearGeolocationOverride"})
}

// SetDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *Page) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["alpha"] = alpha
	paramRequest["beta"] = beta
	paramRequest["gamma"] = gamma
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceOrientationOverride", Params: paramRequest})
}

// Clears the overridden Device Orientation.
func (c *Page) ClearDeviceOrientationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceOrientationOverride"})
}

// SetTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Page) SetTouchEmulationEnabled(enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["enabled"] = enabled
	paramRequest["configuration"] = configuration
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setTouchEmulationEnabled", Params: paramRequest})
}

// CaptureScreenshot - Capture page screenshot.
// format - Image compression format (defaults to png).
// quality - Compression quality from range [0..100] (jpeg only).
// fromSurface - Capture the screenshot from the surface, rather than the view. Defaults to false.
// Returns -  data - Base64-encoded image data.
func (c *Page) CaptureScreenshot(format string, quality int, fromSurface bool) (string, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["format"] = format
	paramRequest["quality"] = quality
	paramRequest["fromSurface"] = fromSurface
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.captureScreenshot", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Data, nil
}

// PrintToPDF - Print page as pdf.
// Returns -  data - Base64-encoded pdf data.
func (c *Page) PrintToPDF() (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.printToPDF"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Data, nil
}

// StartScreencast - Starts sending each frame using the <code>screencastFrame</code> event.
// format - Image compression format.
// quality - Compression quality from range [0..100].
// maxWidth - Maximum screenshot width.
// maxHeight - Maximum screenshot height.
// everyNthFrame - Send every n-th frame.
func (c *Page) StartScreencast(format string, quality int, maxWidth int, maxHeight int, everyNthFrame int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["format"] = format
	paramRequest["quality"] = quality
	paramRequest["maxWidth"] = maxWidth
	paramRequest["maxHeight"] = maxHeight
	paramRequest["everyNthFrame"] = everyNthFrame
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.startScreencast", Params: paramRequest})
}

// Stops sending each frame in the <code>screencastFrame</code>.
func (c *Page) StopScreencast() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopScreencast"})
}

// ScreencastFrameAck - Acknowledges that a screencast frame has been received by the frontend.
// sessionId - Frame number.
func (c *Page) ScreencastFrameAck(sessionId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["sessionId"] = sessionId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.screencastFrameAck", Params: paramRequest})
}

// HandleJavaScriptDialog - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
// accept - Whether to accept or dismiss the dialog.
// promptText - The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (c *Page) HandleJavaScriptDialog(accept bool, promptText string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["accept"] = accept
	paramRequest["promptText"] = promptText
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.handleJavaScriptDialog", Params: paramRequest})
}

// SetColorPickerEnabled - Shows / hides color picker
// enabled - Shows / hides color picker
func (c *Page) SetColorPickerEnabled(enabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setColorPickerEnabled", Params: paramRequest})
}

// ConfigureOverlay - Configures overlay.
// suspended - Whether overlay should be suspended and not consume any resources.
// message - Overlay message to display.
func (c *Page) ConfigureOverlay(suspended bool, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["suspended"] = suspended
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.configureOverlay", Params: paramRequest})
}

// GetAppManifest -
// Returns -  url - Manifest location. errors -  data - Manifest content.
func (c *Page) GetAppManifest() (string, []*PageAppManifestError, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getAppManifest"})
	if err != nil {
		return "", nil, "", err
	}

	var chromeData struct {
		Result struct {
			Url    string
			Errors []*PageAppManifestError
			Data   string
		}
	}

	if resp == nil {
		return "", nil, "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, "", err
	}

	return chromeData.Result.Url, chromeData.Result.Errors, chromeData.Result.Data, nil
}

//
func (c *Page) RequestAppBanner() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.requestAppBanner"})
}

// SetControlNavigations - Toggles navigation throttling which allows programatic control over navigation and redirect response.
// enabled -
func (c *Page) SetControlNavigations(enabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setControlNavigations", Params: paramRequest})
}

// ProcessNavigation - Should be sent in response to a navigationRequested or a redirectRequested event, telling the browser how to handle the navigation.
// response -  enum values: Proceed, Cancel, CancelAndIgnore
// navigationId -
func (c *Page) ProcessNavigation(response string, navigationId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["response"] = response
	paramRequest["navigationId"] = navigationId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.processNavigation", Params: paramRequest})
}

// GetLayoutMetrics - Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
// Returns -  layoutViewport - Metrics relating to the layout viewport. visualViewport - Metrics relating to the visual viewport. contentSize - Size of scrollable area.
func (c *Page) GetLayoutMetrics() (*PageLayoutViewport, *PageVisualViewport, *DOMRect, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getLayoutMetrics"})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			LayoutViewport *PageLayoutViewport
			VisualViewport *PageVisualViewport
			ContentSize    *DOMRect
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.LayoutViewport, chromeData.Result.VisualViewport, chromeData.Result.ContentSize, nil
}
