// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the LayerTree commands.
// API Version: 1.1

package gcd


import (
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

// start non parameterized commands 
// Enables compositing tree inspection.
func (c *ChromeLayerTree) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.enable"})
}
 
// Disables compositing tree inspection.
func (c *ChromeLayerTree) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// releaseSnapshot - Releases layer snapshot captured by the back-end.
// snapshotId - The id of the layer snapshot.
func (c *ChromeLayerTree) ReleaseSnapshot(snapshotId *types.ChromeLayerTreeSnapshotId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["snapshotId"] = snapshotId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "LayerTree.releaseSnapshot", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

