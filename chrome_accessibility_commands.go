// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Accessibility commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Accessibility() *ChromeAccessibility {
	if c.accessibility == nil {
		c.accessibility = newChromeAccessibility(c)
	}
	return c.accessibility
}

type ChromeAccessibility struct {
	target *ChromeTarget
}

func newChromeAccessibility(target *ChromeTarget) *ChromeAccessibility {
	c := &ChromeAccessibility{target: target}
	return c
}

// getAXNode - Fetches the accessibility node for this DOM node, if it exists.
// Returns -
// The <code>Accessibility.AXNode</code> for this DOM node, if it exists.
func (c *ChromeAccessibility) GetAXNode(nodeId *types.ChromeDOMNodeId) (*types.ChromeAccessibilityAXNode, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Accessibility.getAXNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			AccessibilityNode *types.ChromeAccessibilityAXNode
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.AccessibilityNode, nil
}
