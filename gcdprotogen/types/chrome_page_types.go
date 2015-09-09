// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Page types.
// API Version: 1.1
package types

// Resource type as it was perceived by the rendering engine.
type ChromePageResourceType string // possible values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other

// Unique frame identifier.
type ChromePageFrameId string

// Information about the Frame on the page.
type ChromePageFrame struct {
	Id             string                 `json:"id"`                 // Frame unique identifier.
	ParentId       string                 `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderId       *ChromeNetworkLoaderId `json:"loaderId"`           // Identifier of the loader associated with this frame.
	Name           string                 `json:"name,omitempty"`     // Frame's name as specified in the tag.
	Url            string                 `json:"url"`                // Frame document's URL.
	SecurityOrigin string                 `json:"securityOrigin"`     // Frame document's security origin.
	MimeType       string                 `json:"mimeType"`           // Frame document's mimeType as determined by the browser.
}

// Information about frame resources.
type ChromePageResources struct {
	Url      string                  `json:"url"`                // Resource URL.
	Type     *ChromePageResourceType `json:"type"`               // Type of this resource.
	MimeType string                  `json:"mimeType"`           // Resource mimeType as determined by the browser.
	Failed   bool                    `json:"failed,omitempty"`   // True if the resource failed to load.
	Canceled bool                    `json:"canceled,omitempty"` // True if the resource was canceled during loading.
}

// Information about the Frame hierarchy along with their cached resources.
type ChromePageFrameResourceTree struct {
	Frame       *ChromePageFrame               `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*ChromePageFrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []*ChromePageResources         `json:"resources"`             // Information about frame resources.
}

// Unique script identifier.
type ChromePageScriptIdentifier string

// Navigation history entry.
type ChromePageNavigationEntry struct {
	Id    int    `json:"id"`    // Unique id of the navigation history entry.
	Url   string `json:"url"`   // URL of the navigation history entry.
	Title string `json:"title"` // Title of the navigation history entry.
}

// Screencast frame metadata
type ChromePageScreencastFrameMetadata struct {
	OffsetTop       float64 `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64 `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64 `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64 `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64 `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64 `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       float64 `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// Javascript dialog type
type ChromePageDialogType string // possible values: alert, confirm, prompt, beforeunload
