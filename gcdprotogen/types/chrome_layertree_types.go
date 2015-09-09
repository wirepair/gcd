// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the LayerTree types.
// API Version: 1.1
package types

// Unique Layer identifier.
type ChromeLayerTreeLayerId string

// Unique snapshot identifier.
type ChromeLayerTreeSnapshotId string

// Rectangle where scrolling happens on the main thread.
type ChromeLayerTreeScrollRect struct {
	Rect *ChromeDOMRect `json:"rect"` // Rectangle itself.
	Type string         `json:"type"` // Reason for rectangle to force scrolling on the main thread
}

// Serialized fragment of layer picture along with its offset within the layer.
type ChromeLayerTreePictureTile struct {
	X       float64 `json:"x"`       // Offset from owning layer left boundary
	Y       float64 `json:"y"`       // Offset from owning layer top boundary
	Picture string  `json:"picture"` // Base64-encoded snapshot data.
}

// Information about a compositing layer.
type ChromeLayerTreeLayer struct {
	LayerId       *ChromeLayerTreeLayerId      `json:"layerId"`                 // The unique id for this layer.
	ParentLayerId *ChromeLayerTreeLayerId      `json:"parentLayerId,omitempty"` // The id of parent (not present for root).
	BackendNodeId *ChromeDOMBackendNodeId      `json:"backendNodeId,omitempty"` // The backend id for the node associated with this layer.
	OffsetX       float64                      `json:"offsetX"`                 // Offset from parent layer, X coordinate.
	OffsetY       float64                      `json:"offsetY"`                 // Offset from parent layer, Y coordinate.
	Width         float64                      `json:"width"`                   // Layer width.
	Height        float64                      `json:"height"`                  // Layer height.
	Transform     []float64                    `json:"transform,omitempty"`     // Transformation matrix for layer, default is identity matrix
	AnchorX       float64                      `json:"anchorX,omitempty"`       // Transform anchor point X, absent if no transform specified
	AnchorY       float64                      `json:"anchorY,omitempty"`       // Transform anchor point Y, absent if no transform specified
	AnchorZ       float64                      `json:"anchorZ,omitempty"`       // Transform anchor point Z, absent if no transform specified
	PaintCount    int                          `json:"paintCount"`              // Indicates how many time this layer has painted.
	DrawsContent  bool                         `json:"drawsContent"`            // Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	Invisible     bool                         `json:"invisible,omitempty"`     // Set if layer is not visible.
	ScrollRects   []*ChromeLayerTreeScrollRect `json:"scrollRects,omitempty"`   // Rectangles scrolling on main thread only.
}

// Array of timings, one per paint step.
type ChromeLayerTreePaintProfile []float64
