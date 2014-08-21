// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the LayerTree commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) LayerTree() *ChromeLayerTree {
	if c.layertree == nil {
		c.layertree = newChromeLayerTree(c)
	}
	return c.layertree
}

type ChromeLayerTree struct {
	target *ChromeTarget
}

func newChromeLayerTree(target *ChromeTarget) *ChromeLayerTree {
	c := &ChromeLayerTree{target: target}
	return c
}

// Enables compositing tree inspection.
func (c *ChromeLayerTree) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.enable"})
}

// Disables compositing tree inspection.
func (c *ChromeLayerTree) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.disable"})
}

// releaseSnapshot - Releases layer snapshot captured by the back-end.
// snapshotId - The id of the layer snapshot.
func (c *ChromeLayerTree) ReleaseSnapshot(snapshotId *types.ChromeLayerTreeSnapshotId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["snapshotId"] = snapshotId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.releaseSnapshot", Params: paramRequest})
}

// compositingReasons - Provides the reasons why the given layer was composited.
// Returns -
// A list of strings specifying reasons for the given layer to become composited.
func (c *ChromeLayerTree) CompositingReasons(layerId *types.ChromeLayerTreeLayerId) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["layerId"] = layerId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.compositingReasons", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CompositingReasons []string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.CompositingReasons, nil
}

// makeSnapshot - Returns the layer snapshot identifier.
// Returns -
// The id of the layer snapshot.
func (c *ChromeLayerTree) MakeSnapshot(layerId *types.ChromeLayerTreeLayerId) (*types.ChromeLayerTreeSnapshotId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["layerId"] = layerId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.makeSnapshot", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			SnapshotId *types.ChromeLayerTreeSnapshotId
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.SnapshotId, nil
}

// loadSnapshot - Returns the snapshot identifier.
// Returns -
// The id of the snapshot.
func (c *ChromeLayerTree) LoadSnapshot(data string) (*types.ChromeLayerTreeSnapshotId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["data"] = data
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.loadSnapshot", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			SnapshotId *types.ChromeLayerTreeSnapshotId
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.SnapshotId, nil
}

// profileSnapshot -
// Returns -
// The array of paint profiles, one per run.
func (c *ChromeLayerTree) ProfileSnapshot(snapshotId *types.ChromeLayerTreeSnapshotId, minRepeatCount int, minDuration float64) ([]*types.ChromeLayerTreePaintProfile, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["snapshotId"] = snapshotId
	paramRequest["minRepeatCount"] = minRepeatCount
	paramRequest["minDuration"] = minDuration
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.profileSnapshot", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Timings []*types.ChromeLayerTreePaintProfile
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Timings, nil
}

// replaySnapshot - Replays the layer snapshot and returns the resulting bitmap.
// Returns -
// A data: URL for resulting image.
func (c *ChromeLayerTree) ReplaySnapshot(snapshotId *types.ChromeLayerTreeSnapshotId, fromStep int, toStep int, scale float64) (string, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["snapshotId"] = snapshotId
	paramRequest["fromStep"] = fromStep
	paramRequest["toStep"] = toStep
	paramRequest["scale"] = scale
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.replaySnapshot", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			DataURL string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.DataURL, nil
}

// snapshotCommandLog - Replays the layer snapshot and returns canvas log.
// Returns -
// The array of canvas function calls.
func (c *ChromeLayerTree) SnapshotCommandLog(snapshotId *types.ChromeLayerTreeSnapshotId) ([]interface{}, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["snapshotId"] = snapshotId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.snapshotCommandLog", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CommandLog []interface{}
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.CommandLog, nil
}
