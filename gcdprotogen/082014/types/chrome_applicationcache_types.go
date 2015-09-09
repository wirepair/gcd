// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ApplicationCache types.
// API Version: 1.1
package types

 
// Detailed application cache resource information.
type ChromeApplicationCacheApplicationCacheResource struct {
	Url string `json:"url"` // Resource url.
	Size int `json:"size"` // Resource size.
	Type string `json:"type"` // Resource type.
}
 
 
// Detailed application cache information.
type ChromeApplicationCacheApplicationCache struct {
	ManifestURL string `json:"manifestURL"` // Manifest URL.
	Size float64 `json:"size"` // Application cache size.
	CreationTime float64 `json:"creationTime"` // Application cache creation time.
	UpdateTime float64 `json:"updateTime"` // Application cache update time.
	Resources []*ChromeApplicationCacheApplicationCacheResource `json:"resources"` // Application cache resources.
}
 
 
// Frame identifier - manifest URL pair.
type ChromeApplicationCacheFrameWithManifest struct {
	FrameId *ChromePageFrameId `json:"frameId"` // Frame identifier.
	ManifestURL string `json:"manifestURL"` // Manifest URL.
	Status int `json:"status"` // Application cache status.
}
 
