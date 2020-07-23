// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains LayerTree functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/gcdmessage"
)

// Rectangle where scrolling happens on the main thread.
type LayerTreeScrollRect struct {
	Rect *DOMRect `json:"rect"` // Rectangle itself.
	Type string   `json:"type"` // Reason for rectangle to force scrolling on the main thread
}

// Sticky position constraints.
type LayerTreeStickyPositionConstraint struct {
	StickyBoxRect                       *DOMRect `json:"stickyBoxRect"`                                 // Layout rectangle of the sticky element before being shifted
	ContainingBlockRect                 *DOMRect `json:"containingBlockRect"`                           // Layout rectangle of the containing block of the sticky element
	NearestLayerShiftingStickyBox       string   `json:"nearestLayerShiftingStickyBox,omitempty"`       // The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingContainingBlock string   `json:"nearestLayerShiftingContainingBlock,omitempty"` // The nearest sticky layer that shifts the containing block
}

// Serialized fragment of layer picture along with its offset within the layer.
type LayerTreePictureTile struct {
	X       float64 `json:"x"`       // Offset from owning layer left boundary
	Y       float64 `json:"y"`       // Offset from owning layer top boundary
	Picture string  `json:"picture"` // Base64-encoded snapshot data.
}

// Information about a compositing layer.
type LayerTreeLayer struct {
	LayerId                  string                             `json:"layerId"`                            // The unique id for this layer.
	ParentLayerId            string                             `json:"parentLayerId,omitempty"`            // The id of parent (not present for root).
	BackendNodeId            int                                `json:"backendNodeId,omitempty"`            // The backend id for the node associated with this layer.
	OffsetX                  float64                            `json:"offsetX"`                            // Offset from parent layer, X coordinate.
	OffsetY                  float64                            `json:"offsetY"`                            // Offset from parent layer, Y coordinate.
	Width                    float64                            `json:"width"`                              // Layer width.
	Height                   float64                            `json:"height"`                             // Layer height.
	Transform                []float64                          `json:"transform,omitempty"`                // Transformation matrix for layer, default is identity matrix
	AnchorX                  float64                            `json:"anchorX,omitempty"`                  // Transform anchor point X, absent if no transform specified
	AnchorY                  float64                            `json:"anchorY,omitempty"`                  // Transform anchor point Y, absent if no transform specified
	AnchorZ                  float64                            `json:"anchorZ,omitempty"`                  // Transform anchor point Z, absent if no transform specified
	PaintCount               int                                `json:"paintCount"`                         // Indicates how many time this layer has painted.
	DrawsContent             bool                               `json:"drawsContent"`                       // Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	Invisible                bool                               `json:"invisible,omitempty"`                // Set if layer is not visible.
	ScrollRects              []*LayerTreeScrollRect             `json:"scrollRects,omitempty"`              // Rectangles scrolling on main thread only.
	StickyPositionConstraint *LayerTreeStickyPositionConstraint `json:"stickyPositionConstraint,omitempty"` // Sticky position constraint information
}

//
type LayerTreeLayerPaintedEvent struct {
	Method string `json:"method"`
	Params struct {
		LayerId string   `json:"layerId"` // The id of the painted layer.
		Clip    *DOMRect `json:"clip"`    // Clip rectangle.
	} `json:"Params,omitempty"`
}

//
type LayerTreeLayerTreeDidChangeEvent struct {
	Method string `json:"method"`
	Params struct {
		Layers []*LayerTreeLayer `json:"layers,omitempty"` // Layer tree, absent if not in the comspositing mode.
	} `json:"Params,omitempty"`
}

type LayerTree struct {
	target gcdmessage.ChromeTargeter
}

func NewLayerTree(target gcdmessage.ChromeTargeter) *LayerTree {
	c := &LayerTree{target: target}
	return c
}

type LayerTreeCompositingReasonsParams struct {
	// The id of the layer for which we want to get the reasons it was composited.
	LayerId string `json:"layerId"`
}

// CompositingReasonsWithParams - Provides the reasons why the given layer was composited.
// Returns -  compositingReasons - A list of strings specifying reasons for the given layer to become composited. compositingReasonIds - A list of strings specifying reason IDs for the given layer to become composited.
func (c *LayerTree) CompositingReasonsWithParams(ctx context.Context, v *LayerTreeCompositingReasonsParams) ([]string, []string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.compositingReasons", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CompositingReasons   []string
			CompositingReasonIds []string
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.CompositingReasons, chromeData.Result.CompositingReasonIds, nil
}

// CompositingReasons - Provides the reasons why the given layer was composited.
// layerId - The id of the layer for which we want to get the reasons it was composited.
// Returns -  compositingReasons - A list of strings specifying reasons for the given layer to become composited. compositingReasonIds - A list of strings specifying reason IDs for the given layer to become composited.
func (c *LayerTree) CompositingReasons(ctx context.Context, layerId string) ([]string, []string, error) {
	var v LayerTreeCompositingReasonsParams
	v.LayerId = layerId
	return c.CompositingReasonsWithParams(ctx, &v)
}

// Disables compositing tree inspection.
func (c *LayerTree) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.disable"})
}

// Enables compositing tree inspection.
func (c *LayerTree) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.enable"})
}

type LayerTreeLoadSnapshotParams struct {
	// An array of tiles composing the snapshot.
	Tiles []*LayerTreePictureTile `json:"tiles"`
}

// LoadSnapshotWithParams - Returns the snapshot identifier.
// Returns -  snapshotId - The id of the snapshot.
func (c *LayerTree) LoadSnapshotWithParams(ctx context.Context, v *LayerTreeLoadSnapshotParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.loadSnapshot", Params: v})
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
func (c *LayerTree) LoadSnapshot(ctx context.Context, tiles []*LayerTreePictureTile) (string, error) {
	var v LayerTreeLoadSnapshotParams
	v.Tiles = tiles
	return c.LoadSnapshotWithParams(ctx, &v)
}

type LayerTreeMakeSnapshotParams struct {
	// The id of the layer.
	LayerId string `json:"layerId"`
}

// MakeSnapshotWithParams - Returns the layer snapshot identifier.
// Returns -  snapshotId - The id of the layer snapshot.
func (c *LayerTree) MakeSnapshotWithParams(ctx context.Context, v *LayerTreeMakeSnapshotParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.makeSnapshot", Params: v})
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

// MakeSnapshot - Returns the layer snapshot identifier.
// layerId - The id of the layer.
// Returns -  snapshotId - The id of the layer snapshot.
func (c *LayerTree) MakeSnapshot(ctx context.Context, layerId string) (string, error) {
	var v LayerTreeMakeSnapshotParams
	v.LayerId = layerId
	return c.MakeSnapshotWithParams(ctx, &v)
}

type LayerTreeProfileSnapshotParams struct {
	// The id of the layer snapshot.
	SnapshotId string `json:"snapshotId"`
	// The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount int `json:"minRepeatCount,omitempty"`
	// The minimum duration (in seconds) to replay the snapshot.
	MinDuration float64 `json:"minDuration,omitempty"`
	// The clip rectangle to apply when replaying the snapshot.
	ClipRect *DOMRect `json:"clipRect,omitempty"`
}

// ProfileSnapshotWithParams -
// Returns -  timings - The array of paint profiles, one per run.
func (c *LayerTree) ProfileSnapshotWithParams(ctx context.Context, v *LayerTreeProfileSnapshotParams) ([]float64, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.profileSnapshot", Params: v})
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

// ProfileSnapshot -
// snapshotId - The id of the layer snapshot.
// minRepeatCount - The maximum number of times to replay the snapshot (1, if not specified).
// minDuration - The minimum duration (in seconds) to replay the snapshot.
// clipRect - The clip rectangle to apply when replaying the snapshot.
// Returns -  timings - The array of paint profiles, one per run.
func (c *LayerTree) ProfileSnapshot(ctx context.Context, snapshotId string, minRepeatCount int, minDuration float64, clipRect *DOMRect) ([]float64, error) {
	var v LayerTreeProfileSnapshotParams
	v.SnapshotId = snapshotId
	v.MinRepeatCount = minRepeatCount
	v.MinDuration = minDuration
	v.ClipRect = clipRect
	return c.ProfileSnapshotWithParams(ctx, &v)
}

type LayerTreeReleaseSnapshotParams struct {
	// The id of the layer snapshot.
	SnapshotId string `json:"snapshotId"`
}

// ReleaseSnapshotWithParams - Releases layer snapshot captured by the back-end.
func (c *LayerTree) ReleaseSnapshotWithParams(ctx context.Context, v *LayerTreeReleaseSnapshotParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.releaseSnapshot", Params: v})
}

// ReleaseSnapshot - Releases layer snapshot captured by the back-end.
// snapshotId - The id of the layer snapshot.
func (c *LayerTree) ReleaseSnapshot(ctx context.Context, snapshotId string) (*gcdmessage.ChromeResponse, error) {
	var v LayerTreeReleaseSnapshotParams
	v.SnapshotId = snapshotId
	return c.ReleaseSnapshotWithParams(ctx, &v)
}

type LayerTreeReplaySnapshotParams struct {
	// The id of the layer snapshot.
	SnapshotId string `json:"snapshotId"`
	// The first step to replay from (replay from the very start if not specified).
	FromStep int `json:"fromStep,omitempty"`
	// The last step to replay to (replay till the end if not specified).
	ToStep int `json:"toStep,omitempty"`
	// The scale to apply while replaying (defaults to 1).
	Scale float64 `json:"scale,omitempty"`
}

// ReplaySnapshotWithParams - Replays the layer snapshot and returns the resulting bitmap.
// Returns -  dataURL - A data: URL for resulting image.
func (c *LayerTree) ReplaySnapshotWithParams(ctx context.Context, v *LayerTreeReplaySnapshotParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.replaySnapshot", Params: v})
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

// ReplaySnapshot - Replays the layer snapshot and returns the resulting bitmap.
// snapshotId - The id of the layer snapshot.
// fromStep - The first step to replay from (replay from the very start if not specified).
// toStep - The last step to replay to (replay till the end if not specified).
// scale - The scale to apply while replaying (defaults to 1).
// Returns -  dataURL - A data: URL for resulting image.
func (c *LayerTree) ReplaySnapshot(ctx context.Context, snapshotId string, fromStep int, toStep int, scale float64) (string, error) {
	var v LayerTreeReplaySnapshotParams
	v.SnapshotId = snapshotId
	v.FromStep = fromStep
	v.ToStep = toStep
	v.Scale = scale
	return c.ReplaySnapshotWithParams(ctx, &v)
}

type LayerTreeSnapshotCommandLogParams struct {
	// The id of the layer snapshot.
	SnapshotId string `json:"snapshotId"`
}

// SnapshotCommandLogWithParams - Replays the layer snapshot and returns canvas log.
// Returns -
func (c *LayerTree) SnapshotCommandLogWithParams(ctx context.Context, v *LayerTreeSnapshotCommandLogParams) error {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "LayerTree.snapshotCommandLog", Params: v})
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

// SnapshotCommandLog - Replays the layer snapshot and returns canvas log.
// snapshotId - The id of the layer snapshot.
// Returns -
func (c *LayerTree) SnapshotCommandLog(ctx context.Context, snapshotId string) error {
	var v LayerTreeSnapshotCommandLogParams
	v.SnapshotId = snapshotId
	return c.SnapshotCommandLogWithParams(ctx, &v)
}
