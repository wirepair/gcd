// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Browser functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Browser window bounds information
type BrowserBounds struct {
	Left        int    `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         int    `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       int    `json:"width,omitempty"`       // The window width in pixels.
	Height      int    `json:"height,omitempty"`      // The window height in pixels.
	WindowState string `json:"windowState,omitempty"` // The window state. Default to normal. enum values: normal, minimized, maximized, fullscreen
}

// Definition of PermissionDescriptor defined in the Permissions API: https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
type BrowserPermissionDescriptor struct {
	Name                     string `json:"name"`                               // Name of permission. See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Sysex                    bool   `json:"sysex,omitempty"`                    // For "midi" permission, may also specify sysex control.
	UserVisibleOnly          bool   `json:"userVisibleOnly,omitempty"`          // For "push" permission, may specify userVisibleOnly. Note that userVisibleOnly = true is the only currently supported type.
	Type                     string `json:"type,omitempty"`                     // For "wake-lock" permission, must specify type as either "screen" or "system".
	AllowWithoutSanitization bool   `json:"allowWithoutSanitization,omitempty"` // For "clipboard" permission, may specify allowWithoutSanitization.
}

// Chrome histogram bucket.
type BrowserBucket struct {
	Low   int `json:"low"`   // Minimum value (inclusive).
	High  int `json:"high"`  // Maximum value (exclusive).
	Count int `json:"count"` // Number of samples.
}

// Chrome histogram.
type BrowserHistogram struct {
	Name    string           `json:"name"`    // Name.
	Sum     int              `json:"sum"`     // Sum of sample values.
	Count   int              `json:"count"`   // Total number of samples.
	Buckets []*BrowserBucket `json:"buckets"` // Buckets.
}

type Browser struct {
	target gcdmessage.ChromeTargeter
}

func NewBrowser(target gcdmessage.ChromeTargeter) *Browser {
	c := &Browser{target: target}
	return c
}

type BrowserSetPermissionParams struct {
	// Descriptor of permission to override.
	Permission *BrowserPermissionDescriptor `json:"permission"`
	// Setting of the permission. enum values: granted, denied, prompt
	Setting string `json:"setting"`
	// Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`
	// Context to override. When omitted, default browser context is used.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// SetPermissionWithParams - Set permission settings for given origin.
func (c *Browser) SetPermissionWithParams(v *BrowserSetPermissionParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.setPermission", Params: v})
}

// SetPermission - Set permission settings for given origin.
// permission - Descriptor of permission to override.
// setting - Setting of the permission. enum values: granted, denied, prompt
// origin - Origin the permission applies to, all origins if not specified.
// browserContextId - Context to override. When omitted, default browser context is used.
func (c *Browser) SetPermission(permission *BrowserPermissionDescriptor, setting string, origin string, browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v BrowserSetPermissionParams
	v.Permission = permission
	v.Setting = setting
	v.Origin = origin
	v.BrowserContextId = browserContextId
	return c.SetPermissionWithParams(&v)
}

type BrowserGrantPermissionsParams struct {
	//  enum values: accessibilityEvents, audioCapture, backgroundSync, backgroundFetch, clipboardReadWrite, clipboardSanitizedWrite, durableStorage, flash, geolocation, midi, midiSysex, nfc, notifications, paymentHandler, periodicBackgroundSync, protectedMediaIdentifier, sensors, videoCapture, idleDetection, wakeLockScreen, wakeLockSystem
	Permissions []string `json:"permissions"`
	// Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`
	// BrowserContext to override permissions. When omitted, default browser context is used.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// GrantPermissionsWithParams - Grant specific permissions to the given origin and reject all others.
func (c *Browser) GrantPermissionsWithParams(v *BrowserGrantPermissionsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.grantPermissions", Params: v})
}

// GrantPermissions - Grant specific permissions to the given origin and reject all others.
// permissions -  enum values: accessibilityEvents, audioCapture, backgroundSync, backgroundFetch, clipboardReadWrite, clipboardSanitizedWrite, durableStorage, flash, geolocation, midi, midiSysex, nfc, notifications, paymentHandler, periodicBackgroundSync, protectedMediaIdentifier, sensors, videoCapture, idleDetection, wakeLockScreen, wakeLockSystem
// origin - Origin the permission applies to, all origins if not specified.
// browserContextId - BrowserContext to override permissions. When omitted, default browser context is used.
func (c *Browser) GrantPermissions(permissions []string, origin string, browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v BrowserGrantPermissionsParams
	v.Permissions = permissions
	v.Origin = origin
	v.BrowserContextId = browserContextId
	return c.GrantPermissionsWithParams(&v)
}

type BrowserResetPermissionsParams struct {
	// BrowserContext to reset permissions. When omitted, default browser context is used.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// ResetPermissionsWithParams - Reset all permission management for all origins.
func (c *Browser) ResetPermissionsWithParams(v *BrowserResetPermissionsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.resetPermissions", Params: v})
}

// ResetPermissions - Reset all permission management for all origins.
// browserContextId - BrowserContext to reset permissions. When omitted, default browser context is used.
func (c *Browser) ResetPermissions(browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v BrowserResetPermissionsParams
	v.BrowserContextId = browserContextId
	return c.ResetPermissionsWithParams(&v)
}

type BrowserSetDownloadBehaviorParams struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny). |allowAndName| allows download and names files according to their dowmload guids.
	Behavior string `json:"behavior"`
	// BrowserContext to set download behavior. When omitted, default browser context is used.
	BrowserContextId string `json:"browserContextId,omitempty"`
	// The default path to save downloaded files to. This is requred if behavior is set to 'allow' or 'allowAndName'.
	DownloadPath string `json:"downloadPath,omitempty"`
}

// SetDownloadBehaviorWithParams - Set the behavior when downloading a file.
func (c *Browser) SetDownloadBehaviorWithParams(v *BrowserSetDownloadBehaviorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.setDownloadBehavior", Params: v})
}

// SetDownloadBehavior - Set the behavior when downloading a file.
// behavior - Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny). |allowAndName| allows download and names files according to their dowmload guids.
// browserContextId - BrowserContext to set download behavior. When omitted, default browser context is used.
// downloadPath - The default path to save downloaded files to. This is requred if behavior is set to 'allow' or 'allowAndName'.
func (c *Browser) SetDownloadBehavior(behavior string, browserContextId string, downloadPath string) (*gcdmessage.ChromeResponse, error) {
	var v BrowserSetDownloadBehaviorParams
	v.Behavior = behavior
	v.BrowserContextId = browserContextId
	v.DownloadPath = downloadPath
	return c.SetDownloadBehaviorWithParams(&v)
}

// Close browser gracefully.
func (c *Browser) Close() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.close"})
}

// Crashes browser on the main thread.
func (c *Browser) Crash() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.crash"})
}

// Crashes GPU process.
func (c *Browser) CrashGpuProcess() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.crashGpuProcess"})
}

// GetVersion - Returns version information.
// Returns -  protocolVersion - Protocol version. product - Product name. revision - Product revision. userAgent - User-Agent. jsVersion - V8 version.
func (c *Browser) GetVersion() (string, string, string, string, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getVersion"})
	if err != nil {
		return "", "", "", "", "", err
	}

	var chromeData struct {
		Result struct {
			ProtocolVersion string
			Product         string
			Revision        string
			UserAgent       string
			JsVersion       string
		}
	}

	if resp == nil {
		return "", "", "", "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", "", "", "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", "", "", "", "", err
	}

	return chromeData.Result.ProtocolVersion, chromeData.Result.Product, chromeData.Result.Revision, chromeData.Result.UserAgent, chromeData.Result.JsVersion, nil
}

// GetBrowserCommandLine - Returns the command line switches for the browser process if, and only if --enable-automation is on the commandline.
// Returns -  arguments - Commandline parameters
func (c *Browser) GetBrowserCommandLine() ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getBrowserCommandLine"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Arguments []string
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

	return chromeData.Result.Arguments, nil
}

type BrowserGetHistogramsParams struct {
	// Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
	Query string `json:"query,omitempty"`
	// If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// GetHistogramsWithParams - Get Chrome histograms.
// Returns -  histograms - Histograms.
func (c *Browser) GetHistogramsWithParams(v *BrowserGetHistogramsParams) ([]*BrowserHistogram, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getHistograms", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Histograms []*BrowserHistogram
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

	return chromeData.Result.Histograms, nil
}

// GetHistograms - Get Chrome histograms.
// query - Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
// delta - If true, retrieve delta since last call.
// Returns -  histograms - Histograms.
func (c *Browser) GetHistograms(query string, delta bool) ([]*BrowserHistogram, error) {
	var v BrowserGetHistogramsParams
	v.Query = query
	v.Delta = delta
	return c.GetHistogramsWithParams(&v)
}

type BrowserGetHistogramParams struct {
	// Requested histogram name.
	Name string `json:"name"`
	// If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// GetHistogramWithParams - Get a Chrome histogram by name.
// Returns -  histogram - Histogram.
func (c *Browser) GetHistogramWithParams(v *BrowserGetHistogramParams) (*BrowserHistogram, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getHistogram", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Histogram *BrowserHistogram
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

	return chromeData.Result.Histogram, nil
}

// GetHistogram - Get a Chrome histogram by name.
// name - Requested histogram name.
// delta - If true, retrieve delta since last call.
// Returns -  histogram - Histogram.
func (c *Browser) GetHistogram(name string, delta bool) (*BrowserHistogram, error) {
	var v BrowserGetHistogramParams
	v.Name = name
	v.Delta = delta
	return c.GetHistogramWithParams(&v)
}

type BrowserGetWindowBoundsParams struct {
	// Browser window id.
	WindowId int `json:"windowId"`
}

// GetWindowBoundsWithParams - Get position and size of the browser window.
// Returns -  bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowBoundsWithParams(v *BrowserGetWindowBoundsParams) (*BrowserBounds, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getWindowBounds", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Bounds *BrowserBounds
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

	return chromeData.Result.Bounds, nil
}

// GetWindowBounds - Get position and size of the browser window.
// windowId - Browser window id.
// Returns -  bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowBounds(windowId int) (*BrowserBounds, error) {
	var v BrowserGetWindowBoundsParams
	v.WindowId = windowId
	return c.GetWindowBoundsWithParams(&v)
}

type BrowserGetWindowForTargetParams struct {
	// Devtools agent host id. If called as a part of the session, associated targetId is used.
	TargetId string `json:"targetId,omitempty"`
}

// GetWindowForTargetWithParams - Get the browser window that contains the devtools target.
// Returns -  windowId - Browser window id. bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowForTargetWithParams(v *BrowserGetWindowForTargetParams) (int, *BrowserBounds, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getWindowForTarget", Params: v})
	if err != nil {
		return 0, nil, err
	}

	var chromeData struct {
		Result struct {
			WindowId int
			Bounds   *BrowserBounds
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

	return chromeData.Result.WindowId, chromeData.Result.Bounds, nil
}

// GetWindowForTarget - Get the browser window that contains the devtools target.
// targetId - Devtools agent host id. If called as a part of the session, associated targetId is used.
// Returns -  windowId - Browser window id. bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowForTarget(targetId string) (int, *BrowserBounds, error) {
	var v BrowserGetWindowForTargetParams
	v.TargetId = targetId
	return c.GetWindowForTargetWithParams(&v)
}

type BrowserSetWindowBoundsParams struct {
	// Browser window id.
	WindowId int `json:"windowId"`
	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds *BrowserBounds `json:"bounds"`
}

// SetWindowBoundsWithParams - Set position and/or size of the browser window.
func (c *Browser) SetWindowBoundsWithParams(v *BrowserSetWindowBoundsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.setWindowBounds", Params: v})
}

// SetWindowBounds - Set position and/or size of the browser window.
// windowId - Browser window id.
// bounds - New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
func (c *Browser) SetWindowBounds(windowId int, bounds *BrowserBounds) (*gcdmessage.ChromeResponse, error) {
	var v BrowserSetWindowBoundsParams
	v.WindowId = windowId
	v.Bounds = bounds
	return c.SetWindowBoundsWithParams(&v)
}

type BrowserSetDockTileParams struct {
	//
	BadgeLabel string `json:"badgeLabel,omitempty"`
	// Png encoded image.
	Image string `json:"image,omitempty"`
}

// SetDockTileWithParams - Set dock tile details, platform-specific.
func (c *Browser) SetDockTileWithParams(v *BrowserSetDockTileParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.setDockTile", Params: v})
}

// SetDockTile - Set dock tile details, platform-specific.
// badgeLabel -
// image - Png encoded image.
func (c *Browser) SetDockTile(badgeLabel string, image string) (*gcdmessage.ChromeResponse, error) {
	var v BrowserSetDockTileParams
	v.BadgeLabel = badgeLabel
	v.Image = image
	return c.SetDockTileWithParams(&v)
}
