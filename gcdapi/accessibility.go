// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Accessibility functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// A single source for a computed AX property.
type AccessibilityAXValueSource struct {
	Type              string                `json:"type"`                        // What type of source this is. enum values: attribute, implicit, style, contents, placeholder, relatedElement
	Value             *AccessibilityAXValue `json:"value,omitempty"`             // The value of this property source.
	Attribute         string                `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *AccessibilityAXValue `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        bool                  `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      string                `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a <label> element. enum values: figcaption, label, labelfor, labelwrapped, legend, tablecaption, title, other
	NativeSourceValue *AccessibilityAXValue `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           bool                  `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     string                `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// No Description.
type AccessibilityAXRelatedNode struct {
	BackendNodeId int    `json:"backendNodeId"`   // The BackendNodeId of the related node.
	Idref         string `json:"idref,omitempty"` // The IDRef value provided, if any.
	Text          string `json:"text,omitempty"`  // The text alternative of this node in the current context.
}

// No Description.
type AccessibilityAXProperty struct {
	Name  string                `json:"name"`  // The name of this property.
	Value *AccessibilityAXValue `json:"value"` // The value of this property.
}

// A single computed AX property.
type AccessibilityAXValue struct {
	Type         string                        `json:"type"`                   // The type of this value. enum values: boolean, tristate, booleanOrUndefined, idref, idrefList, integer, node, nodeList, number, string, computedString, token, tokenList, domRelation, role, internalRole, valueUndefined
	Value        string                        `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []*AccessibilityAXRelatedNode `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []*AccessibilityAXValueSource `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
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
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getAXNode", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			AccessibilityNode *AccessibilityAXNode
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

	return chromeData.Result.AccessibilityNode, nil
}
