// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains LayerTree functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Rectangle where scrolling happens on the main thread.
type LayerTreeScrollRect struct {
	Rect *DOMRect `json:"rect"` // Rectangle itself.
	Type string   `json:"type"` // Reason for rectangle to force scrolling on the main thread
}

// Serialized fragment of layer picture along with its offset within the layer.
type LayerTreePictureTile struct {
	X       float64 `json:"x"`       // Offset from owning layer left boundary
	Y       float64 `json:"y"`       // Offset from owning layer top boundary
	Picture string  `json:"picture"` // Base64-encoded snapshot data.
}

// Information about a compositing layer.
type LayerTreeLayer struct {
	LayerId       string                 `json:"layerId"`                 // The unique id for this layer.
	ParentLayerId string                 `json:"parentLayerId,omitempty"` // The id of parent (not present for root).
	BackendNodeId int                    `json:"backendNodeId,omitempty"` // The backend id for the node associated with this layer.
	OffsetX       float64                `json:"offsetX"`                 // Offset from parent layer, X coordinate.
	OffsetY       float64                `json:"offsetY"`                 // Offset from parent layer, Y coordinate.
	Width         float64                `json:"width"`                   // Layer width.
	Height        float64                `json:"height"`                  // Layer height.
	Transform     []float64              `json:"transform,omitempty"`     // Transformation matrix for layer, default is identity matrix
	AnchorX       float64                `json:"anchorX,omitempty"`       // Transform anchor point X, absent if no transform specified
	AnchorY       float64                `json:"anchorY,omitempty"`       // Transform anchor point Y, absent if no transform specified
	AnchorZ       float64                `json:"anchorZ,omitempty"`       // Transform anchor point Z, absent if no transform specified
	PaintCount    int                    `json:"paintCount"`              // Indicates how many time this layer has painted.
	DrawsContent  bool                   `json:"drawsContent"`            // Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	Invisible     bool                   `json:"invisible,omitempty"`     // Set if layer is not visible.
	ScrollRects   []*LayerTreeScrollRect `json:"scrollRects,omitempty"`   // Rectangles scrolling on main thread only.
}

//
type LayerTreeLayerTreeDidChangeEvent struct {
	Method string `json:"method"`
	Params struct {
		Layers []*LayerTreeLayer `json:"layers,omitempty"` // Layer tree, absent if not in the comspositing mode.
	} `json:"Params,omitempty"`
}

//
type LayerTreeLayerPaintedEvent struct {
	Method string `json:"method"`
	Params struct {
		LayerId string   `json:"layerId"` // The id of the painted layer.
		Clip    *DOMRect `json:"clip"`    // Clip rectangle.
	} `json:"Params,omitempty"`
}

type LayerTree struct {
	target gcdmessage.ChromeTargeter
}

func NewLayerTree(target gcdmessage.ChromeTargeter) *LayerTree {
	c := &LayerTree{target: target}
	return c
}

// Enables compositing tree inspection.
func (c *LayerTree) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.enable"})
}

// Disables compositing tree inspection.
func (c *LayerTree) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.disable"})
}

// CompositingReasons - Provides the reasons why the given layer was composited.
// layerId - The id of the layer for which we want to get the reasons it was composited.
// Returns -  compositingReasons - A list of strings specifying reasons for the given layer to become composited.
func (c *LayerTree) CompositingReasons(layerId string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["layerId"] = layerId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.compositingReasons", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			CompositingReasons []string
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

	return chromeData.Result.CompositingReasons, nil
}

// MakeSnapshot - Returns the layer snapshot identifier.
// layerId - The id of the layer.
// Returns -  snapshotId - The id of the layer snapshot.
func (c *LayerTree) MakeSnapshot(layerId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["layerId"] = layerId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.makeSnapshot", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SnapshotId string
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

	return chromeData.Result.SnapshotId, nil
}

// LoadSnapshot - Returns the snapshot identifier.
// tiles - An array of tiles composing the snapshot.
// Returns -  snapshotId - The id of the snapshot.
func (c *LayerTree) LoadSnapshot(tiles *LayerTreePictureTile) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["tiles"] = tiles
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.loadSnapshot", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SnapshotId string
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

	return chromeData.Result.SnapshotId, nil
}

// ReleaseSnapshot - Releases layer snapshot captured by the back-end.
// snapshotId - The id of the layer snapshot.
func (c *LayerTree) ReleaseSnapshot(snapshotId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["snapshotId"] = snapshotId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.releaseSnapshot", Params: paramRequest})
}

// ProfileSnapshot -
// snapshotId - The id of the layer snapshot.
// minRepeatCount - The maximum number of times to replay the snapshot (1, if not specified).
// minDuration - The minimum duration (in seconds) to replay the snapshot.
// clipRect - The clip rectangle to apply when replaying the snapshot.
// Returns -  timings - The array of paint profiles, one per run.
func (c *LayerTree) ProfileSnapshot(snapshotId string, minRepeatCount int, minDuration float64, clipRect *DOMRect) ([]float64, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["snapshotId"] = snapshotId
	paramRequest["minRepeatCount"] = minRepeatCount
	paramRequest["minDuration"] = minDuration
	paramRequest["clipRect"] = clipRect
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.profileSnapshot", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Timings []float64
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

	return chromeData.Result.Timings, nil
}

// ReplaySnapshot - Replays the layer snapshot and returns the resulting bitmap.
// snapshotId - The id of the layer snapshot.
// fromStep - The first step to replay from (replay from the very start if not specified).
// toStep - The last step to replay to (replay till the end if not specified).
// scale - The scale to apply while replaying (defaults to 1).
// Returns -  dataURL - A data: URL for resulting image.
func (c *LayerTree) ReplaySnapshot(snapshotId string, fromStep int, toStep int, scale float64) (string, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["snapshotId"] = snapshotId
	paramRequest["fromStep"] = fromStep
	paramRequest["toStep"] = toStep
	paramRequest["scale"] = scale
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.replaySnapshot", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			DataURL string
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

	return chromeData.Result.DataURL, nil
}

// SnapshotCommandLog - Replays the layer snapshot and returns canvas log.
// snapshotId - The id of the layer snapshot.
// Returns -
func (c *LayerTree) SnapshotCommandLog(snapshotId string) error {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["snapshotId"] = snapshotId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.snapshotCommandLog", Params: paramRequest})
	if err != nil {
		return err
	}

	var chromeData struct {
		Result struct {
		}
	}

	if resp == nil {
		return &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return err
	}

	return nil
}
