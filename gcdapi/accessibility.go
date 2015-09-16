// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Accessibility functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// A single source for a computed AX property.
type AccessibilityAXPropertySource struct {
	Name          string `json:"name"`                    // The name/label of this source.
	SourceType    string `json:"sourceType"`              // What type of source this is. enum values: attribute, implicit, style
	Value         string `json:"value"`                   // The value of this property source.
	Type          string `json:"type"`                    // What type the value should be interpreted as. enum values: boolean, tristate, booleanOrUndefined, idref, idrefList, integer, number, string, token, tokenList, domRelation, role, internalRole
	Invalid       bool   `json:"invalid,omitempty"`       // Whether the value for this property is invalid.
	InvalidReason string `json:"invalidReason,omitempty"` // Reason for the value being invalid, if it is.
}

// No Description.
type AccessibilityAXRelatedNode struct {
	Idref         string `json:"idref,omitempty"` // The IDRef value provided, if any.
	BackendNodeId int    `json:"backendNodeId"`   // The BackendNodeId of the related node.
}

// No Description.
type AccessibilityAXProperty struct {
	Name  string                `json:"name"`  // The name of this property.
	Value *AccessibilityAXValue `json:"value"` // The value of this property.
}

// A single computed AX property.
type AccessibilityAXValue struct {
	Type                  string                           `json:"type"`                            // The type of this value. enum values: boolean, tristate, booleanOrUndefined, idref, idrefList, integer, number, string, token, tokenList, domRelation, role, internalRole
	Value                 string                           `json:"value,omitempty"`                 // The computed value of this property.
	RelatedNodeValue      *AccessibilityAXRelatedNode      `json:"relatedNodeValue,omitempty"`      // The related node value, if any.
	RelatedNodeArrayValue []*AccessibilityAXRelatedNode    `json:"relatedNodeArrayValue,omitempty"` // Multiple relted nodes, if applicable.
	Sources               []*AccessibilityAXPropertySource `json:"sources,omitempty"`               // The sources which contributed to the computation of this property.
}

// A node in the accessibility tree.
type AccessibilityAXNode struct {
	NodeId         string                     `json:"nodeId"`                   // Unique identifier for this node.
	Ignored        bool                       `json:"ignored"`                  // Whether this node is ignored for accessibility
	IgnoredReasons []*AccessibilityAXProperty `json:"ignoredReasons,omitempty"` // Collection of reasons why this node is hidden.
	Role           *AccessibilityAXValue      `json:"role,omitempty"`           // This <code>Node</code>'s role, whether explicit or implicit.
	Name           *AccessibilityAXValue      `json:"name,omitempty"`           // The accessible name for this <code>Node</code>.
	Description    *AccessibilityAXValue      `json:"description,omitempty"`    // The accessible description for this <code>Node</code>.
	Value          *AccessibilityAXValue      `json:"value,omitempty"`          // The value for this <code>Node</code>.
	Help           *AccessibilityAXValue      `json:"help,omitempty"`           // Help.
	Properties     []*AccessibilityAXProperty `json:"properties,omitempty"`     // All other properties
}

type Accessibility struct {
	target gcdmessage.ChromeTargeter
}

func NewAccessibility(target gcdmessage.ChromeTargeter) *Accessibility {
	c := &Accessibility{target: target}
	return c
}

// GetAXNode - Fetches the accessibility node for this DOM node, if it exists.
// nodeId - ID of node to get accessibility node for.
// Returns -  accessibilityNode - The <code>Accessibility.AXNode</code> for this DOM node, if it exists.
func (c *Accessibility) GetAXNode(nodeId int) (*AccessibilityAXNode, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getAXNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			AccessibilityNode *AccessibilityAXNode
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.AccessibilityNode, nil
}
