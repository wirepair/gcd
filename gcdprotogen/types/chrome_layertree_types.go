// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the LayerTree types.
// API Version: 1.1
package main

 
// Unique RenderLayer identifier.
type ChromeLayerTreeLayerId string 
 
 
// Unique snapshot identifier.
type ChromeLayerTreeSnapshotId string 
 
 
// Rectangle where scrolling happens on the main thread.
type ChromeLayerTreeScrollRect struct {
	Rect *ChromeDOMRect `json:"rect"` // Rectangle itself.
	Type string `json:"type"` // Reason for rectangle to force scrolling on the main thread
}
 
 
// Information about a compositing layer.
type ChromeLayerTreeLayer struct {
	LayerId *ChromeLayerTreeLayerId `json:"layerId"` // The unique id for this layer.
	ParentLayerId *ChromeLayerTreeLayerId `json:"parentLayerId,omitempty"` // The id of parent (not present for root).
	BackendNodeId *ChromeDOMBackendNodeId `json:"backendNodeId,omitempty"` // The backend id for the node associated with this layer.
	OffsetX int `json:"offsetX"` // Offset from parent layer, X coordinate.
	OffsetY int `json:"offsetY"` // Offset from parent layer, Y coordinate.
	Width int `json:"width"` // Layer width.
	Height int `json:"height"` // Layer height.
	Transform []int `json:"transform,omitempty"` // Transformation matrix for layer, default is identity matrix
	AnchorX int `json:"anchorX,omitempty"` // Transform anchor point X, absent if no transform specified
	AnchorY int `json:"anchorY,omitempty"` // Transform anchor point Y, absent if no transform specified
	AnchorZ int `json:"anchorZ,omitempty"` // Transform anchor point Z, absent if no transform specified
	PaintCount int `json:"paintCount"` // Indicates how many time this layer has painted.
	Invisible bool `json:"invisible,omitempty"` // Set if layer is not visible.
	ScrollRects []*ChromeLayerTreeScrollRect `json:"scrollRects,omitempty"` // Rectangles scrolling on main thread only.
}
 
 
// Array of timings, one per paint step.
type ChromeLayerTreePaintProfile []int 
 
