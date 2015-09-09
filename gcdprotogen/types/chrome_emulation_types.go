// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Emulation types.
// API Version: 1.1
package types

// Visible page viewport
type ChromeEmulationViewport struct {
	ScrollX                float64 `json:"scrollX"`                // X scroll offset in CSS pixels.
	ScrollY                float64 `json:"scrollY"`                // Y scroll offset in CSS pixels.
	ContentsWidth          float64 `json:"contentsWidth"`          // Contents width in CSS pixels.
	ContentsHeight         float64 `json:"contentsHeight"`         // Contents height in CSS pixels.
	PageScaleFactor        float64 `json:"pageScaleFactor"`        // Page scale factor.
	MinimumPageScaleFactor float64 `json:"minimumPageScaleFactor"` // Minimum page scale factor.
	MaximumPageScaleFactor float64 `json:"maximumPageScaleFactor"` // Maximum page scale factor.
}
