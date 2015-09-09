// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Page types.
// API Version: 1.1
package types

 
// Resource type as it was perceived by the rendering engine.
type ChromePageResourceType string // possible values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, WebSocket, Other
 
 
// Unique frame identifier.
type ChromePageFrameId string 
 
 
// Information about the Frame on the page.
type ChromePageFrame struct {
	Id string `json:"id"` // Frame unique identifier.
	ParentId string `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderId *ChromeNetworkLoaderId `json:"loaderId"` // Identifier of the loader associated with this frame.
	Name string `json:"name,omitempty"` // Frame's name as specified in the tag.
	Url string `json:"url"` // Frame document's URL.
	SecurityOrigin string `json:"securityOrigin"` // Frame document's security origin.
	MimeType string `json:"mimeType"` // Frame document's mimeType as determined by the browser.
}
 
 
// Information about frame resources.
type ChromePageResources struct {
	Url string `json:"url"` // Resource URL.
	Type *ChromePageResourceType `json:"type"` // Type of this resource.
	MimeType string `json:"mimeType"` // Resource mimeType as determined by the browser.
	Failed bool `json:"failed,omitempty"` // True if the resource failed to load.
	Canceled bool `json:"canceled,omitempty"` // True if the resource was canceled during loading.
}
 
 
// Information about the Frame hierarchy along with their cached resources.
type ChromePageFrameResourceTree struct {
	Frame *ChromePageFrame `json:"frame"` // Frame information for this tree item.
	ChildFrames []*ChromePageFrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources []*ChromePageResources `json:"resources"` // Information about frame resources.
}
 
 
// Search match for resource.
type ChromePageSearchMatch struct {
	LineNumber float64 `json:"lineNumber"` // Line number in resource content.
	LineContent string `json:"lineContent"` // Line with match content.
}
 
 
// Cookie object
type ChromePageCookie struct {
	Name string `json:"name"` // Cookie name.
	Value string `json:"value"` // Cookie value.
	Domain string `json:"domain"` // Cookie domain.
	Path string `json:"path"` // Cookie path.
	Expires float64 `json:"expires"` // Cookie expires.
	Size int `json:"size"` // Cookie size.
	HttpOnly bool `json:"httpOnly"` // True if cookie is http-only.
	Secure bool `json:"secure"` // True if cookie is secure.
	Session bool `json:"session"` // True in case of session cookie.
}
 
 
// Unique script identifier.
type ChromePageScriptIdentifier string 
 
 
// Navigation history entry.
type ChromePageNavigationEntry struct {
	Id int `json:"id"` // Unique id of the navigation history entry.
	Url string `json:"url"` // URL of the navigation history entry.
	Title string `json:"title"` // Title of the navigation history entry.
}
 
 
// Quota information
type ChromePageQuota struct {
	Temporary float64 `json:"temporary"` // Quota for temporary storage shared among all security origins
	Persistent float64 `json:"persistent"` // Quota for persistent storage for the security origin.
}
 
 
// Usage information
type ChromePageUsage struct {
	Temporary []*ChromePageUsageItem `json:"temporary"` // Temporary storage usage.
	Persistent []*ChromePageUsageItem `json:"persistent"` // Persistent storage usage.
	Syncable []*ChromePageUsageItem `json:"syncable"` // Syncable storage.
}
 
 
// Usage information for a client and storage type
type ChromePageUsageItem struct {
	Id string `json:"id"` // Item id.
	Value float64 `json:"value"` // Item usage value.
}
 
 
// Visible page viewport
type ChromePageViewport struct {
	ScrollX float64 `json:"scrollX"` // X scroll offset in CSS pixels.
	ScrollY float64 `json:"scrollY"` // Y scroll offset in CSS pixels.
	ContentsWidth float64 `json:"contentsWidth"` // Contents width in CSS pixels.
	ContentsHeight float64 `json:"contentsHeight"` // Contents height in CSS pixels.
	PageScaleFactor float64 `json:"pageScaleFactor"` // Page scale factor.
}
 
 
// Screencast frame metadata
type ChromePageScreencastFrameMetadata struct {
	DeviceScaleFactor float64 `json:"deviceScaleFactor"` // Device scale factor.
	Viewport *ChromeDOMRect `json:"viewport"` // Viewport in CSS pixels.
	OffsetTop float64 `json:"offsetTop,omitempty"` // Top offset in DIP.
	OffsetBottom float64 `json:"offsetBottom,omitempty"` // Bottom offset in DIP.
	PageScaleFactor float64 `json:"pageScaleFactor"` // Page scale factor.
	PageScaleFactorMin float64 `json:"pageScaleFactorMin"` // Page scale factor min.
	PageScaleFactorMax float64 `json:"pageScaleFactorMax"` // Page scale factor max.
	DeviceWidth float64 `json:"deviceWidth"` // Device screen width in DIP.
	DeviceHeight float64 `json:"deviceHeight"` // Device screen height in DIP.
	ScrollOffsetX float64 `json:"scrollOffsetX"` // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY float64 `json:"scrollOffsetY"` // Position of vertical scroll in CSS pixels.
}
 
