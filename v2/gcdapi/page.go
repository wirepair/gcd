// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Page functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type PagePermissionsPolicyBlockLocator struct {
	FrameId     string `json:"frameId"`     //
	BlockReason string `json:"blockReason"` //  enum values: Header, IframeAttribute
}

// No Description.
type PagePermissionsPolicyFeatureState struct {
	Feature string                             `json:"feature"`           //  enum values: accelerometer, ambient-light-sensor, autoplay, camera, ch-dpr, ch-device-memory, ch-downlink, ch-ect, ch-lang, ch-rtt, ch-ua, ch-ua-arch, ch-ua-platform, ch-ua-model, ch-ua-mobile, ch-ua-full-version, ch-ua-platform-version, ch-viewport-width, ch-width, clipboard-read, clipboard-write, conversion-measurement, cross-origin-isolated, display-capture, document-domain, encrypted-media, execution-while-out-of-viewport, execution-while-not-rendered, focus-without-user-activation, fullscreen, frobulate, gamepad, geolocation, gyroscope, hid, idle-detection, interest-cohort, magnetometer, microphone, midi, otp-credentials, payment, picture-in-picture, publickey-credentials-get, screen-wake-lock, serial, storage-access-api, sync-xhr, trust-token-redemption, usb, vertical-scroll, web-share, xr-spatial-tracking
	Allowed bool                               `json:"allowed"`           //
	Locator *PagePermissionsPolicyBlockLocator `json:"locator,omitempty"` //
}

// Information about the Frame on the page.
type PageFrame struct {
	Id                             string   `json:"id"`                             // Frame unique identifier.
	ParentId                       string   `json:"parentId,omitempty"`             // Parent frame identifier.
	LoaderId                       string   `json:"loaderId"`                       // Identifier of the loader associated with this frame.
	Name                           string   `json:"name,omitempty"`                 // Frame's name as specified in the tag.
	Url                            string   `json:"url"`                            // Frame document's URL without fragment.
	UrlFragment                    string   `json:"urlFragment,omitempty"`          // Frame document's URL fragment including the '#'.
	DomainAndRegistry              string   `json:"domainAndRegistry"`              // Frame document's registered domain, taking the public suffixes list into account. Extracted from the Frame's url. Example URLs: http://www.google.com/file.html -> "google.com"               http://a.b.co.uk/file.html      -> "b.co.uk"
	SecurityOrigin                 string   `json:"securityOrigin"`                 // Frame document's security origin.
	MimeType                       string   `json:"mimeType"`                       // Frame document's mimeType as determined by the browser.
	UnreachableUrl                 string   `json:"unreachableUrl,omitempty"`       // If the frame failed to load, this contains the URL that could not be loaded. Note that unlike url above, this URL may contain a fragment.
	AdFrameType                    string   `json:"adFrameType,omitempty"`          // Indicates whether this frame was tagged as an ad. enum values: none, child, root
	SecureContextType              string   `json:"secureContextType"`              // Indicates whether the main document is a secure context and explains why that is the case. enum values: Secure, SecureLocalhost, InsecureScheme, InsecureAncestor
	CrossOriginIsolatedContextType string   `json:"crossOriginIsolatedContextType"` // Indicates whether this is a cross origin isolated context. enum values: Isolated, NotIsolated, NotIsolatedFeatureDisabled
	GatedAPIFeatures               []string `json:"gatedAPIFeatures"`               // Indicated which gated APIs / features are available. enum values: SharedArrayBuffers, SharedArrayBuffersTransferAllowed, PerformanceMeasureMemory, PerformanceProfile
}

// Information about the Resource on the page.
type PageFrameResource struct {
	Url          string  `json:"url"`                    // Resource URL.
	Type         string  `json:"type"`                   // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
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

// Information about the Frame hierarchy.
type PageFrameTree struct {
	Frame       *PageFrame       `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*PageFrameTree `json:"childFrames,omitempty"` // Child frames.
}

// Navigation history entry.
type PageNavigationEntry struct {
	Id             int    `json:"id"`             // Unique id of the navigation history entry.
	Url            string `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string `json:"title"`          // Title of the navigation history entry.
	TransitionType string `json:"transitionType"` // Transition type. enum values: link, typed, address_bar, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
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

// Parsed app manifest properties.
type PageAppManifestParsedProperties struct {
	Scope string `json:"scope"` // Computed scope value
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
	OffsetX      float64 `json:"offsetX"`        // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`        // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`          // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`          // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`    // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"`   // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`          // Scale relative to the ideal viewport (size at width=device-width).
	Zoom         float64 `json:"zoom,omitempty"` // Page zoom factor (CSS to device independent pixels ratio).
}

// Viewport for capturing screenshot.
type PageViewport struct {
	X      float64 `json:"x"`      // X offset in device independent pixels (dip).
	Y      float64 `json:"y"`      // Y offset in device independent pixels (dip).
	Width  float64 `json:"width"`  // Rectangle width in device independent pixels (dip).
	Height float64 `json:"height"` // Rectangle height in device independent pixels (dip).
	Scale  float64 `json:"scale"`  // Page scale factor.
}

// Generic font families collection.
type PageFontFamilies struct {
	Standard   string `json:"standard,omitempty"`   // The standard font-family.
	Fixed      string `json:"fixed,omitempty"`      // The fixed font-family.
	Serif      string `json:"serif,omitempty"`      // The serif font-family.
	SansSerif  string `json:"sansSerif,omitempty"`  // The sansSerif font-family.
	Cursive    string `json:"cursive,omitempty"`    // The cursive font-family.
	Fantasy    string `json:"fantasy,omitempty"`    // The fantasy font-family.
	Pictograph string `json:"pictograph,omitempty"` // The pictograph font-family.
}

// Default font sizes.
type PageFontSizes struct {
	Standard int `json:"standard,omitempty"` // Default standard font size.
	Fixed    int `json:"fixed,omitempty"`    // Default fixed font size.
}

// No Description.
type PageInstallabilityErrorArgument struct {
	Name  string `json:"name"`  // Argument name (e.g. name:'minimum-icon-size-in-pixels').
	Value string `json:"value"` // Argument value (e.g. value:'64').
}

// The installability error
type PageInstallabilityError struct {
	ErrorId        string                             `json:"errorId"`        // The error id (e.g. 'manifest-missing-suitable-icon').
	ErrorArguments []*PageInstallabilityErrorArgument `json:"errorArguments"` // The list of error arguments (e.g. {name:'minimum-icon-size-in-pixels', value:'64'}).
}

// Per-script compilation cache parameters for `Page.produceCompilationCache`
type PageCompilationCacheParams struct {
	Url   string `json:"url"`             // The URL of the script to produce a compilation cache entry for.
	Eager bool   `json:"eager,omitempty"` // A hint to the backend whether eager compilation is recommended. (the actual compilation mode used is upon backend discretion).
}

//
type PageDomContentEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

// Emitted only when `page.interceptFileChooser` is enabled.
type PageFileChooserOpenedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId       string `json:"frameId"`       // Id of the frame containing input node.
		BackendNodeId int    `json:"backendNodeId"` // Input node id.
		Mode          string `json:"mode"`          // Input mode.
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

// Fired when frame no longer has a scheduled navigation.
type PageFrameClearedScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
	} `json:"Params,omitempty"`
}

// Fired when frame has been detached from its parent.
type PageFrameDetachedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has been detached.
		Reason  string `json:"reason"`  //
	} `json:"Params,omitempty"`
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
type PageFrameNavigatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Frame *PageFrame `json:"frame"` // Frame object.
		Type  string     `json:"type"`  //  enum values: Navigation, BackForwardCacheRestore
	} `json:"Params,omitempty"`
}

// Fired when opening document to write to.
type PageDocumentOpenedEvent struct {
	Method string `json:"method"`
	Params struct {
		Frame *PageFrame `json:"frame"` // Frame object.
	} `json:"Params,omitempty"`
}

// Fired when a renderer-initiated navigation is requested. Navigation may still be cancelled after the event is issued.
type PageFrameRequestedNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId     string `json:"frameId"`     // Id of the frame that is being navigated.
		Reason      string `json:"reason"`      // The reason for the navigation. enum values: formSubmissionGet, formSubmissionPost, httpHeaderRefresh, scriptInitiated, metaTagRefresh, pageBlockInterstitial, reload, anchorClick
		Url         string `json:"url"`         // The destination URL for the requested navigation.
		Disposition string `json:"disposition"` // The disposition for the navigation. enum values: currentTab, newTab, newWindow, download
	} `json:"Params,omitempty"`
}

// Fired when frame schedules a potential navigation.
type PageFrameScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string  `json:"frameId"` // Id of the frame that has scheduled a navigation.
		Delay   float64 `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
		Reason  string  `json:"reason"`  // The reason for the navigation. enum values: formSubmissionGet, formSubmissionPost, httpHeaderRefresh, scriptInitiated, metaTagRefresh, pageBlockInterstitial, reload, anchorClick
		Url     string  `json:"url"`     // The destination URL for the scheduled navigation.
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

// Fired when page is about to start a download. Deprecated. Use Browser.downloadWillBegin instead.
type PageDownloadWillBeginEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId           string `json:"frameId"`           // Id of the frame that caused download to begin.
		Guid              string `json:"guid"`              // Global unique identifier of the download.
		Url               string `json:"url"`               // URL of the resource being downloaded.
		SuggestedFilename string `json:"suggestedFilename"` // Suggested file name of the resource (the actual name of the file saved on disk may differ).
	} `json:"Params,omitempty"`
}

// Fired when download makes progress. Last call has |done| == true. Deprecated. Use Browser.downloadProgress instead.
type PageDownloadProgressEvent struct {
	Method string `json:"method"`
	Params struct {
		Guid          string  `json:"guid"`          // Global unique identifier of the download.
		TotalBytes    float64 `json:"totalBytes"`    // Total expected bytes to download.
		ReceivedBytes float64 `json:"receivedBytes"` // Total bytes received.
		State         string  `json:"state"`         // Download status.
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been closed.
type PageJavascriptDialogClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		Result    bool   `json:"result"`    // Whether dialog was confirmed.
		UserInput string `json:"userInput"` // User input in case of prompt.
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to open.
type PageJavascriptDialogOpeningEvent struct {
	Method string `json:"method"`
	Params struct {
		Url               string `json:"url"`                     // Frame url.
		Message           string `json:"message"`                 // Message that will be displayed by the dialog.
		Type              string `json:"type"`                    // Dialog type. enum values: alert, confirm, prompt, beforeunload
		HasBrowserHandler bool   `json:"hasBrowserHandler"`       // True iff browser is capable showing or acting on the given dialog. When browser has no dialog handler for given target, calling alert while Page domain is engaged will stall the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
		DefaultPrompt     string `json:"defaultPrompt,omitempty"` // Default dialog prompt.
	} `json:"Params,omitempty"`
}

// Fired for top level page lifecycle events such as navigation, load, paint, etc.
type PageLifecycleEventEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId   string  `json:"frameId"`   // Id of the frame.
		LoaderId  string  `json:"loaderId"`  // Loader identifier. Empty string if the request is fetched from worker.
		Name      string  `json:"name"`      //
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

// Fired for failed bfcache history navigations if BackForwardCache feature is enabled. Do not assume any ordering with the Page.frameNavigated event. This event is fired only for main-frame history navigation where the document changes (non-same-document navigations), when bfcache navigation fails.
type PageBackForwardCacheNotUsedEvent struct {
	Method string `json:"method"`
	Params struct {
		LoaderId string `json:"loaderId"` // The loader id for the associated navgation.
		FrameId  string `json:"frameId"`  // The frame id of the associated frame.
	} `json:"Params,omitempty"`
}

//
type PageLoadEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

// Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
type PageNavigatedWithinDocumentEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame.
		Url     string `json:"url"`     // Frame's new url.
	} `json:"Params,omitempty"`
}

// Compressed image data requested by the `startScreencast`.
type PageScreencastFrameEvent struct {
	Method string `json:"method"`
	Params struct {
		Data      string                       `json:"data"`      // Base64-encoded compressed image. (Encoded as a base64 string when passed over JSON)
		Metadata  *PageScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
		SessionId int                          `json:"sessionId"` // Frame number.
	} `json:"Params,omitempty"`
}

// Fired when the page with currently enabled screencast was shown or hidden `.
type PageScreencastVisibilityChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		Visible bool `json:"visible"` // True if the page is visible.
	} `json:"Params,omitempty"`
}

// Fired when a new window is going to be opened, via window.open(), link click, form submission, etc.
type PageWindowOpenEvent struct {
	Method string `json:"method"`
	Params struct {
		Url            string   `json:"url"`            // The URL for the new window.
		WindowName     string   `json:"windowName"`     // Window name.
		WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
		UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
	} `json:"Params,omitempty"`
}

// Issued for every compilation cache generated. Is only available if Page.setGenerateCompilationCache is enabled.
type PageCompilationCacheProducedEvent struct {
	Method string `json:"method"`
	Params struct {
		Url  string `json:"url"`  //
		Data string `json:"data"` // Base64-encoded data (Encoded as a base64 string when passed over JSON)
	} `json:"Params,omitempty"`
}

type Page struct {
	target gcdmessage.ChromeTargeter
}

func NewPage(target gcdmessage.ChromeTargeter) *Page {
	c := &Page{target: target}
	return c
}

type PageAddScriptToEvaluateOnLoadParams struct {
	//
	ScriptSource string `json:"scriptSource"`
}

// AddScriptToEvaluateOnLoadWithParams - Deprecated, please use addScriptToEvaluateOnNewDocument instead.
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnLoadWithParams(ctx context.Context, v *PageAddScriptToEvaluateOnLoadParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addScriptToEvaluateOnLoad", Params: v})
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

// AddScriptToEvaluateOnLoad - Deprecated, please use addScriptToEvaluateOnNewDocument instead.
// scriptSource -
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnLoad(ctx context.Context, scriptSource string) (string, error) {
	var v PageAddScriptToEvaluateOnLoadParams
	v.ScriptSource = scriptSource
	return c.AddScriptToEvaluateOnLoadWithParams(ctx, &v)
}

type PageAddScriptToEvaluateOnNewDocumentParams struct {
	//
	Source string `json:"source"`
	// If specified, creates an isolated world with the given name and evaluates given script in it. This world name will be used as the ExecutionContextDescription::name when the corresponding event is emitted.
	WorldName string `json:"worldName,omitempty"`
}

// AddScriptToEvaluateOnNewDocumentWithParams - Evaluates given script in every frame upon creation (before loading frame's scripts).
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnNewDocumentWithParams(ctx context.Context, v *PageAddScriptToEvaluateOnNewDocumentParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addScriptToEvaluateOnNewDocument", Params: v})
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

// AddScriptToEvaluateOnNewDocument - Evaluates given script in every frame upon creation (before loading frame's scripts).
// source -
// worldName - If specified, creates an isolated world with the given name and evaluates given script in it. This world name will be used as the ExecutionContextDescription::name when the corresponding event is emitted.
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnNewDocument(ctx context.Context, source string, worldName string) (string, error) {
	var v PageAddScriptToEvaluateOnNewDocumentParams
	v.Source = source
	v.WorldName = worldName
	return c.AddScriptToEvaluateOnNewDocumentWithParams(ctx, &v)
}

// Brings page to front (activates tab).
func (c *Page) BringToFront(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.bringToFront"})
}

type PageCaptureScreenshotParams struct {
	// Image compression format (defaults to png).
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
	// Capture the screenshot of a given region only.
	Clip *PageViewport `json:"clip,omitempty"`
	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	FromSurface bool `json:"fromSurface,omitempty"`
	// Capture the screenshot beyond the viewport. Defaults to false.
	CaptureBeyondViewport bool `json:"captureBeyondViewport,omitempty"`
}

// CaptureScreenshotWithParams - Capture page screenshot.
// Returns -  data - Base64-encoded image data. (Encoded as a base64 string when passed over JSON)
func (c *Page) CaptureScreenshotWithParams(ctx context.Context, v *PageCaptureScreenshotParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.captureScreenshot", Params: v})
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

// CaptureScreenshot - Capture page screenshot.
// format - Image compression format (defaults to png).
// quality - Compression quality from range [0..100] (jpeg only).
// clip - Capture the screenshot of a given region only.
// fromSurface - Capture the screenshot from the surface, rather than the view. Defaults to true.
// captureBeyondViewport - Capture the screenshot beyond the viewport. Defaults to false.
// Returns -  data - Base64-encoded image data. (Encoded as a base64 string when passed over JSON)
func (c *Page) CaptureScreenshot(ctx context.Context, format string, quality int, clip *PageViewport, fromSurface bool, captureBeyondViewport bool) (string, error) {
	var v PageCaptureScreenshotParams
	v.Format = format
	v.Quality = quality
	v.Clip = clip
	v.FromSurface = fromSurface
	v.CaptureBeyondViewport = captureBeyondViewport
	return c.CaptureScreenshotWithParams(ctx, &v)
}

type PageCaptureSnapshotParams struct {
	// Format (defaults to mhtml).
	Format string `json:"format,omitempty"`
}

// CaptureSnapshotWithParams - Returns a snapshot of the page as a string. For MHTML format, the serialization includes iframes, shadow DOM, external resources, and element-inline styles.
// Returns -  data - Serialized page data.
func (c *Page) CaptureSnapshotWithParams(ctx context.Context, v *PageCaptureSnapshotParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.captureSnapshot", Params: v})
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

// CaptureSnapshot - Returns a snapshot of the page as a string. For MHTML format, the serialization includes iframes, shadow DOM, external resources, and element-inline styles.
// format - Format (defaults to mhtml).
// Returns -  data - Serialized page data.
func (c *Page) CaptureSnapshot(ctx context.Context, format string) (string, error) {
	var v PageCaptureSnapshotParams
	v.Format = format
	return c.CaptureSnapshotWithParams(ctx, &v)
}

// Clears the overridden device metrics.
func (c *Page) ClearDeviceMetricsOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceMetricsOverride"})
}

// Clears the overridden Device Orientation.
func (c *Page) ClearDeviceOrientationOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceOrientationOverride"})
}

// Clears the overridden Geolocation Position and Error.
func (c *Page) ClearGeolocationOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearGeolocationOverride"})
}

type PageCreateIsolatedWorldParams struct {
	// Id of the frame in which the isolated world should be created.
	FrameId string `json:"frameId"`
	// An optional name which is reported in the Execution Context.
	WorldName string `json:"worldName,omitempty"`
	// Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

// CreateIsolatedWorldWithParams - Creates an isolated world for the given frame.
// Returns -  executionContextId - Execution context of the isolated world.
func (c *Page) CreateIsolatedWorldWithParams(ctx context.Context, v *PageCreateIsolatedWorldParams) (int, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.createIsolatedWorld", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			ExecutionContextId int
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

	return chromeData.Result.ExecutionContextId, nil
}

// CreateIsolatedWorld - Creates an isolated world for the given frame.
// frameId - Id of the frame in which the isolated world should be created.
// worldName - An optional name which is reported in the Execution Context.
// grantUniveralAccess - Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
// Returns -  executionContextId - Execution context of the isolated world.
func (c *Page) CreateIsolatedWorld(ctx context.Context, frameId string, worldName string, grantUniveralAccess bool) (int, error) {
	var v PageCreateIsolatedWorldParams
	v.FrameId = frameId
	v.WorldName = worldName
	v.GrantUniveralAccess = grantUniveralAccess
	return c.CreateIsolatedWorldWithParams(ctx, &v)
}

type PageDeleteCookieParams struct {
	// Name of the cookie to remove.
	CookieName string `json:"cookieName"`
	// URL to match cooke domain and path.
	Url string `json:"url"`
}

// DeleteCookieWithParams - Deletes browser cookie with given name, domain and path.
func (c *Page) DeleteCookieWithParams(ctx context.Context, v *PageDeleteCookieParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.deleteCookie", Params: v})
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Page) DeleteCookie(ctx context.Context, cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	var v PageDeleteCookieParams
	v.CookieName = cookieName
	v.Url = url
	return c.DeleteCookieWithParams(ctx, &v)
}

// Disables page domain notifications.
func (c *Page) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.disable"})
}

// Enables page domain notifications.
func (c *Page) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.enable"})
}

// GetAppManifest -
// Returns -  url - Manifest location. errors -  data - Manifest content. parsed - Parsed manifest properties
func (c *Page) GetAppManifest(ctx context.Context) (string, []*PageAppManifestError, string, *PageAppManifestParsedProperties, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getAppManifest"})
	if err != nil {
		return "", nil, "", nil, err
	}

	var chromeData struct {
		Result struct {
			Url    string
			Errors []*PageAppManifestError
			Data   string
			Parsed *PageAppManifestParsedProperties
		}
	}

	if resp == nil {
		return "", nil, "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, "", nil, err
	}

	return chromeData.Result.Url, chromeData.Result.Errors, chromeData.Result.Data, chromeData.Result.Parsed, nil
}

// GetInstallabilityErrors -
// Returns -  installabilityErrors -
func (c *Page) GetInstallabilityErrors(ctx context.Context) ([]*PageInstallabilityError, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getInstallabilityErrors"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			InstallabilityErrors []*PageInstallabilityError
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

	return chromeData.Result.InstallabilityErrors, nil
}

// GetManifestIcons -
// Returns -  primaryIcon -
func (c *Page) GetManifestIcons(ctx context.Context) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getManifestIcons"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			PrimaryIcon string
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

	return chromeData.Result.PrimaryIcon, nil
}

// GetCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the `cookies` field.
// Returns -  cookies - Array of cookie objects.
func (c *Page) GetCookies(ctx context.Context) ([]*NetworkCookie, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getCookies"})
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

// GetFrameTree - Returns present frame tree structure.
// Returns -  frameTree - Present frame tree structure.
func (c *Page) GetFrameTree(ctx context.Context) (*PageFrameTree, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getFrameTree"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			FrameTree *PageFrameTree
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

// GetLayoutMetrics - Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
// Returns -  layoutViewport - Deprecated metrics relating to the layout viewport. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssLayoutViewport` instead. visualViewport - Deprecated metrics relating to the visual viewport. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssVisualViewport` instead. contentSize - Deprecated size of scrollable area. Can be in DP or in CSS pixels depending on the `enable-use-zoom-for-dsf` flag. Use `cssContentSize` instead. cssLayoutViewport - Metrics relating to the layout viewport in CSS pixels. cssVisualViewport - Metrics relating to the visual viewport in CSS pixels. cssContentSize - Size of scrollable area in CSS pixels.
func (c *Page) GetLayoutMetrics(ctx context.Context) (*PageLayoutViewport, *PageVisualViewport, *DOMRect, *PageLayoutViewport, *PageVisualViewport, *DOMRect, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getLayoutMetrics"})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			LayoutViewport    *PageLayoutViewport
			VisualViewport    *PageVisualViewport
			ContentSize       *DOMRect
			CssLayoutViewport *PageLayoutViewport
			CssVisualViewport *PageVisualViewport
			CssContentSize    *DOMRect
		}
	}

	if resp == nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return chromeData.Result.LayoutViewport, chromeData.Result.VisualViewport, chromeData.Result.ContentSize, chromeData.Result.CssLayoutViewport, chromeData.Result.CssVisualViewport, chromeData.Result.CssContentSize, nil
}

// GetNavigationHistory - Returns navigation history for the current page.
// Returns -  currentIndex - Index of the current navigation history entry. entries - Array of navigation history entries.
func (c *Page) GetNavigationHistory(ctx context.Context) (int, []*PageNavigationEntry, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getNavigationHistory"})
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

// Resets navigation history for the current page.
func (c *Page) ResetNavigationHistory(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.resetNavigationHistory"})
}

type PageGetResourceContentParams struct {
	// Frame id to get resource for.
	FrameId string `json:"frameId"`
	// URL of the resource to get content for.
	Url string `json:"url"`
}

// GetResourceContentWithParams - Returns content of the given resource.
// Returns -  content - Resource content. base64Encoded - True, if content was served as base64.
func (c *Page) GetResourceContentWithParams(ctx context.Context, v *PageGetResourceContentParams) (string, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceContent", Params: v})
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

// GetResourceContent - Returns content of the given resource.
// frameId - Frame id to get resource for.
// url - URL of the resource to get content for.
// Returns -  content - Resource content. base64Encoded - True, if content was served as base64.
func (c *Page) GetResourceContent(ctx context.Context, frameId string, url string) (string, bool, error) {
	var v PageGetResourceContentParams
	v.FrameId = frameId
	v.Url = url
	return c.GetResourceContentWithParams(ctx, &v)
}

// GetResourceTree - Returns present frame / resource tree structure.
// Returns -  frameTree - Present frame / resource tree structure.
func (c *Page) GetResourceTree(ctx context.Context) (*PageFrameResourceTree, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceTree"})
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

type PageHandleJavaScriptDialogParams struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`
	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
	PromptText string `json:"promptText,omitempty"`
}

// HandleJavaScriptDialogWithParams - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
func (c *Page) HandleJavaScriptDialogWithParams(ctx context.Context, v *PageHandleJavaScriptDialogParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.handleJavaScriptDialog", Params: v})
}

// HandleJavaScriptDialog - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
// accept - Whether to accept or dismiss the dialog.
// promptText - The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (c *Page) HandleJavaScriptDialog(ctx context.Context, accept bool, promptText string) (*gcdmessage.ChromeResponse, error) {
	var v PageHandleJavaScriptDialogParams
	v.Accept = accept
	v.PromptText = promptText
	return c.HandleJavaScriptDialogWithParams(ctx, &v)
}

type PageNavigateParams struct {
	// URL to navigate the page to.
	Url string `json:"url"`
	// Referrer URL.
	Referrer string `json:"referrer,omitempty"`
	// Intended transition type. enum values: link, typed, address_bar, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
	TransitionType string `json:"transitionType,omitempty"`
	// Frame id to navigate, if not specified navigates the top frame.
	FrameId string `json:"frameId,omitempty"`
	// Referrer-policy used for the navigation. enum values: noReferrer, noReferrerWhenDowngrade, origin, originWhenCrossOrigin, sameOrigin, strictOrigin, strictOriginWhenCrossOrigin, unsafeUrl
	ReferrerPolicy string `json:"referrerPolicy,omitempty"`
}

// NavigateWithParams - Navigates current page to the given URL.
// Returns -  frameId - Frame id that has navigated (or failed to navigate) loaderId - Loader identifier. errorText - User friendly error message, present if and only if navigation has failed.
func (c *Page) NavigateWithParams(ctx context.Context, v *PageNavigateParams) (string, string, string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigate", Params: v})
	if err != nil {
		return "", "", "", err
	}

	var chromeData struct {
		Result struct {
			FrameId   string
			LoaderId  string
			ErrorText string
		}
	}

	if resp == nil {
		return "", "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", "", "", err
	}

	return chromeData.Result.FrameId, chromeData.Result.LoaderId, chromeData.Result.ErrorText, nil
}

// Navigate - Navigates current page to the given URL.
// url - URL to navigate the page to.
// referrer - Referrer URL.
// transitionType - Intended transition type. enum values: link, typed, address_bar, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
// frameId - Frame id to navigate, if not specified navigates the top frame.
// referrerPolicy - Referrer-policy used for the navigation. enum values: noReferrer, noReferrerWhenDowngrade, origin, originWhenCrossOrigin, sameOrigin, strictOrigin, strictOriginWhenCrossOrigin, unsafeUrl
// Returns -  frameId - Frame id that has navigated (or failed to navigate) loaderId - Loader identifier. errorText - User friendly error message, present if and only if navigation has failed.
func (c *Page) Navigate(ctx context.Context, url string, referrer string, transitionType string, frameId string, referrerPolicy string) (string, string, string, error) {
	var v PageNavigateParams
	v.Url = url
	v.Referrer = referrer
	v.TransitionType = transitionType
	v.FrameId = frameId
	v.ReferrerPolicy = referrerPolicy
	return c.NavigateWithParams(ctx, &v)
}

type PageNavigateToHistoryEntryParams struct {
	// Unique id of the entry to navigate to.
	EntryId int `json:"entryId"`
}

// NavigateToHistoryEntryWithParams - Navigates current page to the given history entry.
func (c *Page) NavigateToHistoryEntryWithParams(ctx context.Context, v *PageNavigateToHistoryEntryParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigateToHistoryEntry", Params: v})
}

// NavigateToHistoryEntry - Navigates current page to the given history entry.
// entryId - Unique id of the entry to navigate to.
func (c *Page) NavigateToHistoryEntry(ctx context.Context, entryId int) (*gcdmessage.ChromeResponse, error) {
	var v PageNavigateToHistoryEntryParams
	v.EntryId = entryId
	return c.NavigateToHistoryEntryWithParams(ctx, &v)
}

type PagePrintToPDFParams struct {
	// Paper orientation. Defaults to false.
	Landscape bool `json:"landscape,omitempty"`
	// Display header and footer. Defaults to false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`
	// Print background graphics. Defaults to false.
	PrintBackground bool `json:"printBackground,omitempty"`
	// Scale of the webpage rendering. Defaults to 1.
	Scale float64 `json:"scale,omitempty"`
	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth float64 `json:"paperWidth,omitempty"`
	// Paper height in inches. Defaults to 11 inches.
	PaperHeight float64 `json:"paperHeight,omitempty"`
	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop float64 `json:"marginTop,omitempty"`
	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom float64 `json:"marginBottom,omitempty"`
	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft float64 `json:"marginLeft,omitempty"`
	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight float64 `json:"marginRight,omitempty"`
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	PageRanges string `json:"pageRanges,omitempty"`
	// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
	IgnoreInvalidPageRanges bool `json:"ignoreInvalidPageRanges,omitempty"`
	// HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - `date`: formatted print date - `title`: document title - `url`: document location - `pageNumber`: current page number - `totalPages`: total pages in the document  For example, `<span class=title></span>` would generate span containing the title.
	HeaderTemplate string `json:"headerTemplate,omitempty"`
	// HTML template for the print footer. Should use the same format as the `headerTemplate`.
	FooterTemplate string `json:"footerTemplate,omitempty"`
	// Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize bool `json:"preferCSSPageSize,omitempty"`
	// return as stream
	TransferMode string `json:"transferMode,omitempty"`
}

// PrintToPDFWithParams - Print page as PDF.
// Returns -  data - Base64-encoded pdf data. Empty if |returnAsStream| is specified. (Encoded as a base64 string when passed over JSON) stream - A handle of the stream that holds resulting PDF data.
func (c *Page) PrintToPDFWithParams(ctx context.Context, v *PagePrintToPDFParams) (string, string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.printToPDF", Params: v})
	if err != nil {
		return "", "", err
	}

	var chromeData struct {
		Result struct {
			Data   string
			Stream string
		}
	}

	if resp == nil {
		return "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", "", err
	}

	return chromeData.Result.Data, chromeData.Result.Stream, nil
}

// PrintToPDF - Print page as PDF.
// landscape - Paper orientation. Defaults to false.
// displayHeaderFooter - Display header and footer. Defaults to false.
// printBackground - Print background graphics. Defaults to false.
// scale - Scale of the webpage rendering. Defaults to 1.
// paperWidth - Paper width in inches. Defaults to 8.5 inches.
// paperHeight - Paper height in inches. Defaults to 11 inches.
// marginTop - Top margin in inches. Defaults to 1cm (~0.4 inches).
// marginBottom - Bottom margin in inches. Defaults to 1cm (~0.4 inches).
// marginLeft - Left margin in inches. Defaults to 1cm (~0.4 inches).
// marginRight - Right margin in inches. Defaults to 1cm (~0.4 inches).
// pageRanges - Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
// ignoreInvalidPageRanges - Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
// headerTemplate - HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - `date`: formatted print date - `title`: document title - `url`: document location - `pageNumber`: current page number - `totalPages`: total pages in the document  For example, `<span class=title></span>` would generate span containing the title.
// footerTemplate - HTML template for the print footer. Should use the same format as the `headerTemplate`.
// preferCSSPageSize - Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
// transferMode - return as stream
// Returns -  data - Base64-encoded pdf data. Empty if |returnAsStream| is specified. (Encoded as a base64 string when passed over JSON) stream - A handle of the stream that holds resulting PDF data.
func (c *Page) PrintToPDF(ctx context.Context, landscape bool, displayHeaderFooter bool, printBackground bool, scale float64, paperWidth float64, paperHeight float64, marginTop float64, marginBottom float64, marginLeft float64, marginRight float64, pageRanges string, ignoreInvalidPageRanges bool, headerTemplate string, footerTemplate string, preferCSSPageSize bool, transferMode string) (string, string, error) {
	var v PagePrintToPDFParams
	v.Landscape = landscape
	v.DisplayHeaderFooter = displayHeaderFooter
	v.PrintBackground = printBackground
	v.Scale = scale
	v.PaperWidth = paperWidth
	v.PaperHeight = paperHeight
	v.MarginTop = marginTop
	v.MarginBottom = marginBottom
	v.MarginLeft = marginLeft
	v.MarginRight = marginRight
	v.PageRanges = pageRanges
	v.IgnoreInvalidPageRanges = ignoreInvalidPageRanges
	v.HeaderTemplate = headerTemplate
	v.FooterTemplate = footerTemplate
	v.PreferCSSPageSize = preferCSSPageSize
	v.TransferMode = transferMode
	return c.PrintToPDFWithParams(ctx, &v)
}

type PageReloadParams struct {
	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache bool `json:"ignoreCache,omitempty"`
	// If set, the script will be injected into all frames of the inspected page after reload. Argument will be ignored if reloading dataURL origin.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

// ReloadWithParams - Reloads given page optionally ignoring the cache.
func (c *Page) ReloadWithParams(ctx context.Context, v *PageReloadParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.reload", Params: v})
}

// Reload - Reloads given page optionally ignoring the cache.
// ignoreCache - If true, browser cache is ignored (as if the user pressed Shift+refresh).
// scriptToEvaluateOnLoad - If set, the script will be injected into all frames of the inspected page after reload. Argument will be ignored if reloading dataURL origin.
func (c *Page) Reload(ctx context.Context, ignoreCache bool, scriptToEvaluateOnLoad string) (*gcdmessage.ChromeResponse, error) {
	var v PageReloadParams
	v.IgnoreCache = ignoreCache
	v.ScriptToEvaluateOnLoad = scriptToEvaluateOnLoad
	return c.ReloadWithParams(ctx, &v)
}

type PageRemoveScriptToEvaluateOnLoadParams struct {
	//
	Identifier string `json:"identifier"`
}

// RemoveScriptToEvaluateOnLoadWithParams - Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
func (c *Page) RemoveScriptToEvaluateOnLoadWithParams(ctx context.Context, v *PageRemoveScriptToEvaluateOnLoadParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.removeScriptToEvaluateOnLoad", Params: v})
}

// RemoveScriptToEvaluateOnLoad - Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
// identifier -
func (c *Page) RemoveScriptToEvaluateOnLoad(ctx context.Context, identifier string) (*gcdmessage.ChromeResponse, error) {
	var v PageRemoveScriptToEvaluateOnLoadParams
	v.Identifier = identifier
	return c.RemoveScriptToEvaluateOnLoadWithParams(ctx, &v)
}

type PageRemoveScriptToEvaluateOnNewDocumentParams struct {
	//
	Identifier string `json:"identifier"`
}

// RemoveScriptToEvaluateOnNewDocumentWithParams - Removes given script from the list.
func (c *Page) RemoveScriptToEvaluateOnNewDocumentWithParams(ctx context.Context, v *PageRemoveScriptToEvaluateOnNewDocumentParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.removeScriptToEvaluateOnNewDocument", Params: v})
}

// RemoveScriptToEvaluateOnNewDocument - Removes given script from the list.
// identifier -
func (c *Page) RemoveScriptToEvaluateOnNewDocument(ctx context.Context, identifier string) (*gcdmessage.ChromeResponse, error) {
	var v PageRemoveScriptToEvaluateOnNewDocumentParams
	v.Identifier = identifier
	return c.RemoveScriptToEvaluateOnNewDocumentWithParams(ctx, &v)
}

type PageScreencastFrameAckParams struct {
	// Frame number.
	SessionId int `json:"sessionId"`
}

// ScreencastFrameAckWithParams - Acknowledges that a screencast frame has been received by the frontend.
func (c *Page) ScreencastFrameAckWithParams(ctx context.Context, v *PageScreencastFrameAckParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.screencastFrameAck", Params: v})
}

// ScreencastFrameAck - Acknowledges that a screencast frame has been received by the frontend.
// sessionId - Frame number.
func (c *Page) ScreencastFrameAck(ctx context.Context, sessionId int) (*gcdmessage.ChromeResponse, error) {
	var v PageScreencastFrameAckParams
	v.SessionId = sessionId
	return c.ScreencastFrameAckWithParams(ctx, &v)
}

type PageSearchInResourceParams struct {
	// Frame id for resource to search in.
	FrameId string `json:"frameId"`
	// URL of the resource to search in.
	Url string `json:"url"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInResourceWithParams - Searches for given string in resource content.
// Returns -  result - List of search matches.
func (c *Page) SearchInResourceWithParams(ctx context.Context, v *PageSearchInResourceParams) ([]*DebuggerSearchMatch, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.searchInResource", Params: v})
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

// SearchInResource - Searches for given string in resource content.
// frameId - Frame id for resource to search in.
// url - URL of the resource to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Page) SearchInResource(ctx context.Context, frameId string, url string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v PageSearchInResourceParams
	v.FrameId = frameId
	v.Url = url
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInResourceWithParams(ctx, &v)
}

type PageSetAdBlockingEnabledParams struct {
	// Whether to block ads.
	Enabled bool `json:"enabled"`
}

// SetAdBlockingEnabledWithParams - Enable Chrome's experimental ad filter on all sites.
func (c *Page) SetAdBlockingEnabledWithParams(ctx context.Context, v *PageSetAdBlockingEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setAdBlockingEnabled", Params: v})
}

// SetAdBlockingEnabled - Enable Chrome's experimental ad filter on all sites.
// enabled - Whether to block ads.
func (c *Page) SetAdBlockingEnabled(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetAdBlockingEnabledParams
	v.Enabled = enabled
	return c.SetAdBlockingEnabledWithParams(ctx, &v)
}

type PageSetBypassCSPParams struct {
	// Whether to bypass page CSP.
	Enabled bool `json:"enabled"`
}

// SetBypassCSPWithParams - Enable page Content Security Policy by-passing.
func (c *Page) SetBypassCSPWithParams(ctx context.Context, v *PageSetBypassCSPParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setBypassCSP", Params: v})
}

// SetBypassCSP - Enable page Content Security Policy by-passing.
// enabled - Whether to bypass page CSP.
func (c *Page) SetBypassCSP(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetBypassCSPParams
	v.Enabled = enabled
	return c.SetBypassCSPWithParams(ctx, &v)
}

type PageGetPermissionsPolicyStateParams struct {
	//
	FrameId string `json:"frameId"`
}

// GetPermissionsPolicyStateWithParams - Get Permissions Policy state on given frame.
// Returns -  states -
func (c *Page) GetPermissionsPolicyStateWithParams(ctx context.Context, v *PageGetPermissionsPolicyStateParams) ([]*PagePermissionsPolicyFeatureState, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getPermissionsPolicyState", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			States []*PagePermissionsPolicyFeatureState
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

	return chromeData.Result.States, nil
}

// GetPermissionsPolicyState - Get Permissions Policy state on given frame.
// frameId -
// Returns -  states -
func (c *Page) GetPermissionsPolicyState(ctx context.Context, frameId string) ([]*PagePermissionsPolicyFeatureState, error) {
	var v PageGetPermissionsPolicyStateParams
	v.FrameId = frameId
	return c.GetPermissionsPolicyStateWithParams(ctx, &v)
}

type PageSetDeviceMetricsOverrideParams struct {
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
	// The viewport dimensions and scale. If not set, the override is cleared.
	Viewport *PageViewport `json:"viewport,omitempty"`
}

// SetDeviceMetricsOverrideWithParams - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (c *Page) SetDeviceMetricsOverrideWithParams(ctx context.Context, v *PageSetDeviceMetricsOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceMetricsOverride", Params: v})
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
// viewport - The viewport dimensions and scale. If not set, the override is cleared.
func (c *Page) SetDeviceMetricsOverride(ctx context.Context, width int, height int, deviceScaleFactor float64, mobile bool, scale float64, screenWidth int, screenHeight int, positionX int, positionY int, dontSetVisibleSize bool, screenOrientation *EmulationScreenOrientation, viewport *PageViewport) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDeviceMetricsOverrideParams
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
	return c.SetDeviceMetricsOverrideWithParams(ctx, &v)
}

type PageSetDeviceOrientationOverrideParams struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`
	// Mock beta
	Beta float64 `json:"beta"`
	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// SetDeviceOrientationOverrideWithParams - Overrides the Device Orientation.
func (c *Page) SetDeviceOrientationOverrideWithParams(ctx context.Context, v *PageSetDeviceOrientationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceOrientationOverride", Params: v})
}

// SetDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *Page) SetDeviceOrientationOverride(ctx context.Context, alpha float64, beta float64, gamma float64) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDeviceOrientationOverrideParams
	v.Alpha = alpha
	v.Beta = beta
	v.Gamma = gamma
	return c.SetDeviceOrientationOverrideWithParams(ctx, &v)
}

type PageSetFontFamiliesParams struct {
	// Specifies font families to set. If a font family is not specified, it won't be changed.
	FontFamilies *PageFontFamilies `json:"fontFamilies"`
}

// SetFontFamiliesWithParams - Set generic font families.
func (c *Page) SetFontFamiliesWithParams(ctx context.Context, v *PageSetFontFamiliesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setFontFamilies", Params: v})
}

// SetFontFamilies - Set generic font families.
// fontFamilies - Specifies font families to set. If a font family is not specified, it won't be changed.
func (c *Page) SetFontFamilies(ctx context.Context, fontFamilies *PageFontFamilies) (*gcdmessage.ChromeResponse, error) {
	var v PageSetFontFamiliesParams
	v.FontFamilies = fontFamilies
	return c.SetFontFamiliesWithParams(ctx, &v)
}

type PageSetFontSizesParams struct {
	// Specifies font sizes to set. If a font size is not specified, it won't be changed.
	FontSizes *PageFontSizes `json:"fontSizes"`
}

// SetFontSizesWithParams - Set default font sizes.
func (c *Page) SetFontSizesWithParams(ctx context.Context, v *PageSetFontSizesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setFontSizes", Params: v})
}

// SetFontSizes - Set default font sizes.
// fontSizes - Specifies font sizes to set. If a font size is not specified, it won't be changed.
func (c *Page) SetFontSizes(ctx context.Context, fontSizes *PageFontSizes) (*gcdmessage.ChromeResponse, error) {
	var v PageSetFontSizesParams
	v.FontSizes = fontSizes
	return c.SetFontSizesWithParams(ctx, &v)
}

type PageSetDocumentContentParams struct {
	// Frame id to set HTML for.
	FrameId string `json:"frameId"`
	// HTML content to set.
	Html string `json:"html"`
}

// SetDocumentContentWithParams - Sets given markup as the document's HTML.
func (c *Page) SetDocumentContentWithParams(ctx context.Context, v *PageSetDocumentContentParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDocumentContent", Params: v})
}

// SetDocumentContent - Sets given markup as the document's HTML.
// frameId - Frame id to set HTML for.
// html - HTML content to set.
func (c *Page) SetDocumentContent(ctx context.Context, frameId string, html string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDocumentContentParams
	v.FrameId = frameId
	v.Html = html
	return c.SetDocumentContentWithParams(ctx, &v)
}

type PageSetDownloadBehaviorParams struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
	Behavior string `json:"behavior"`
	// The default path to save downloaded files to. This is required if behavior is set to 'allow'
	DownloadPath string `json:"downloadPath,omitempty"`
}

// SetDownloadBehaviorWithParams - Set the behavior when downloading a file.
func (c *Page) SetDownloadBehaviorWithParams(ctx context.Context, v *PageSetDownloadBehaviorParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDownloadBehavior", Params: v})
}

// SetDownloadBehavior - Set the behavior when downloading a file.
// behavior - Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
// downloadPath - The default path to save downloaded files to. This is required if behavior is set to 'allow'
func (c *Page) SetDownloadBehavior(ctx context.Context, behavior string, downloadPath string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDownloadBehaviorParams
	v.Behavior = behavior
	v.DownloadPath = downloadPath
	return c.SetDownloadBehaviorWithParams(ctx, &v)
}

type PageSetGeolocationOverrideParams struct {
	// Mock latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy float64 `json:"accuracy,omitempty"`
}

// SetGeolocationOverrideWithParams - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (c *Page) SetGeolocationOverrideWithParams(ctx context.Context, v *PageSetGeolocationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setGeolocationOverride", Params: v})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Page) SetGeolocationOverride(ctx context.Context, latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	var v PageSetGeolocationOverrideParams
	v.Latitude = latitude
	v.Longitude = longitude
	v.Accuracy = accuracy
	return c.SetGeolocationOverrideWithParams(ctx, &v)
}

type PageSetLifecycleEventsEnabledParams struct {
	// If true, starts emitting lifecycle events.
	Enabled bool `json:"enabled"`
}

// SetLifecycleEventsEnabledWithParams - Controls whether page will emit lifecycle events.
func (c *Page) SetLifecycleEventsEnabledWithParams(ctx context.Context, v *PageSetLifecycleEventsEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setLifecycleEventsEnabled", Params: v})
}

// SetLifecycleEventsEnabled - Controls whether page will emit lifecycle events.
// enabled - If true, starts emitting lifecycle events.
func (c *Page) SetLifecycleEventsEnabled(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetLifecycleEventsEnabledParams
	v.Enabled = enabled
	return c.SetLifecycleEventsEnabledWithParams(ctx, &v)
}

type PageSetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration string `json:"configuration,omitempty"`
}

// SetTouchEmulationEnabledWithParams - Toggles mouse event-based touch event emulation.
func (c *Page) SetTouchEmulationEnabledWithParams(ctx context.Context, v *PageSetTouchEmulationEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setTouchEmulationEnabled", Params: v})
}

// SetTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Page) SetTouchEmulationEnabled(ctx context.Context, enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetTouchEmulationEnabledParams
	v.Enabled = enabled
	v.Configuration = configuration
	return c.SetTouchEmulationEnabledWithParams(ctx, &v)
}

type PageStartScreencastParams struct {
	// Image compression format.
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100].
	Quality int `json:"quality,omitempty"`
	// Maximum screenshot width.
	MaxWidth int `json:"maxWidth,omitempty"`
	// Maximum screenshot height.
	MaxHeight int `json:"maxHeight,omitempty"`
	// Send every n-th frame.
	EveryNthFrame int `json:"everyNthFrame,omitempty"`
}

// StartScreencastWithParams - Starts sending each frame using the `screencastFrame` event.
func (c *Page) StartScreencastWithParams(ctx context.Context, v *PageStartScreencastParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.startScreencast", Params: v})
}

// StartScreencast - Starts sending each frame using the `screencastFrame` event.
// format - Image compression format.
// quality - Compression quality from range [0..100].
// maxWidth - Maximum screenshot width.
// maxHeight - Maximum screenshot height.
// everyNthFrame - Send every n-th frame.
func (c *Page) StartScreencast(ctx context.Context, format string, quality int, maxWidth int, maxHeight int, everyNthFrame int) (*gcdmessage.ChromeResponse, error) {
	var v PageStartScreencastParams
	v.Format = format
	v.Quality = quality
	v.MaxWidth = maxWidth
	v.MaxHeight = maxHeight
	v.EveryNthFrame = everyNthFrame
	return c.StartScreencastWithParams(ctx, &v)
}

// Force the page stop all navigations and pending resource fetches.
func (c *Page) StopLoading(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopLoading"})
}

// Crashes renderer on the IO thread, generates minidumps.
func (c *Page) Crash(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.crash"})
}

// Tries to close page, running its beforeunload hooks, if any.
func (c *Page) Close(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.close"})
}

type PageSetWebLifecycleStateParams struct {
	// Target lifecycle state
	State string `json:"state"`
}

// SetWebLifecycleStateWithParams - Tries to update the web lifecycle state of the page. It will transition the page to the given state according to: https://github.com/WICG/web-lifecycle/
func (c *Page) SetWebLifecycleStateWithParams(ctx context.Context, v *PageSetWebLifecycleStateParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setWebLifecycleState", Params: v})
}

// SetWebLifecycleState - Tries to update the web lifecycle state of the page. It will transition the page to the given state according to: https://github.com/WICG/web-lifecycle/
// state - Target lifecycle state
func (c *Page) SetWebLifecycleState(ctx context.Context, state string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetWebLifecycleStateParams
	v.State = state
	return c.SetWebLifecycleStateWithParams(ctx, &v)
}

// Stops sending each frame in the `screencastFrame`.
func (c *Page) StopScreencast(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopScreencast"})
}

type PageSetProduceCompilationCacheParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetProduceCompilationCacheWithParams - Forces compilation cache to be generated for every subresource script. See also: `Page.produceCompilationCache`.
func (c *Page) SetProduceCompilationCacheWithParams(ctx context.Context, v *PageSetProduceCompilationCacheParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setProduceCompilationCache", Params: v})
}

// SetProduceCompilationCache - Forces compilation cache to be generated for every subresource script. See also: `Page.produceCompilationCache`.
// enabled -
func (c *Page) SetProduceCompilationCache(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetProduceCompilationCacheParams
	v.Enabled = enabled
	return c.SetProduceCompilationCacheWithParams(ctx, &v)
}

type PageProduceCompilationCacheParams struct {
	//
	Scripts []*PageCompilationCacheParams `json:"scripts"`
}

// ProduceCompilationCacheWithParams - Requests backend to produce compilation cache for the specified scripts. Unlike setProduceCompilationCache, this allows client to only produce cache for specific scripts. `scripts` are appeneded to the list of scripts for which the cache for would produced. Disabling compilation cache with `setProduceCompilationCache` would reset all pending cache requests. The list may also be reset during page navigation. When script with a matching URL is encountered, the cache is optionally produced upon backend discretion, based on internal heuristics. See also: `Page.compilationCacheProduced`.
func (c *Page) ProduceCompilationCacheWithParams(ctx context.Context, v *PageProduceCompilationCacheParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.produceCompilationCache", Params: v})
}

// ProduceCompilationCache - Requests backend to produce compilation cache for the specified scripts. Unlike setProduceCompilationCache, this allows client to only produce cache for specific scripts. `scripts` are appeneded to the list of scripts for which the cache for would produced. Disabling compilation cache with `setProduceCompilationCache` would reset all pending cache requests. The list may also be reset during page navigation. When script with a matching URL is encountered, the cache is optionally produced upon backend discretion, based on internal heuristics. See also: `Page.compilationCacheProduced`.
// scripts -
func (c *Page) ProduceCompilationCache(ctx context.Context, scripts []*PageCompilationCacheParams) (*gcdmessage.ChromeResponse, error) {
	var v PageProduceCompilationCacheParams
	v.Scripts = scripts
	return c.ProduceCompilationCacheWithParams(ctx, &v)
}

type PageAddCompilationCacheParams struct {
	//
	Url string `json:"url"`
	// Base64-encoded data (Encoded as a base64 string when passed over JSON)
	Data string `json:"data"`
}

// AddCompilationCacheWithParams - Seeds compilation cache for given url. Compilation cache does not survive cross-process navigation.
func (c *Page) AddCompilationCacheWithParams(ctx context.Context, v *PageAddCompilationCacheParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addCompilationCache", Params: v})
}

// AddCompilationCache - Seeds compilation cache for given url. Compilation cache does not survive cross-process navigation.
// url -
// data - Base64-encoded data (Encoded as a base64 string when passed over JSON)
func (c *Page) AddCompilationCache(ctx context.Context, url string, data string) (*gcdmessage.ChromeResponse, error) {
	var v PageAddCompilationCacheParams
	v.Url = url
	v.Data = data
	return c.AddCompilationCacheWithParams(ctx, &v)
}

// Clears seeded compilation cache.
func (c *Page) ClearCompilationCache(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearCompilationCache"})
}

type PageGenerateTestReportParams struct {
	// Message to be displayed in the report.
	Message string `json:"message"`
	// Specifies the endpoint group to deliver the report to.
	Group string `json:"group,omitempty"`
}

// GenerateTestReportWithParams - Generates a report for testing.
func (c *Page) GenerateTestReportWithParams(ctx context.Context, v *PageGenerateTestReportParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.generateTestReport", Params: v})
}

// GenerateTestReport - Generates a report for testing.
// message - Message to be displayed in the report.
// group - Specifies the endpoint group to deliver the report to.
func (c *Page) GenerateTestReport(ctx context.Context, message string, group string) (*gcdmessage.ChromeResponse, error) {
	var v PageGenerateTestReportParams
	v.Message = message
	v.Group = group
	return c.GenerateTestReportWithParams(ctx, &v)
}

// Pauses page execution. Can be resumed using generic Runtime.runIfWaitingForDebugger.
func (c *Page) WaitForDebugger(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.waitForDebugger"})
}

type PageSetInterceptFileChooserDialogParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetInterceptFileChooserDialogWithParams - Intercept file chooser requests and transfer control to protocol clients. When file chooser interception is enabled, native file chooser dialog is not shown. Instead, a protocol event `Page.fileChooserOpened` is emitted.
func (c *Page) SetInterceptFileChooserDialogWithParams(ctx context.Context, v *PageSetInterceptFileChooserDialogParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setInterceptFileChooserDialog", Params: v})
}

// SetInterceptFileChooserDialog - Intercept file chooser requests and transfer control to protocol clients. When file chooser interception is enabled, native file chooser dialog is not shown. Instead, a protocol event `Page.fileChooserOpened` is emitted.
// enabled -
func (c *Page) SetInterceptFileChooserDialog(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetInterceptFileChooserDialogParams
	v.Enabled = enabled
	return c.SetInterceptFileChooserDialogWithParams(ctx, &v)
}
