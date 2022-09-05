// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Accessibility functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// A single source for a computed AX property.
type AccessibilityAXValueSource struct {
	Type              string                `json:"type"`                        // What type of source this is. enum values: attribute, implicit, style, contents, placeholder, relatedElement
	Value             *AccessibilityAXValue `json:"value,omitempty"`             // The value of this property source.
	Attribute         string                `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *AccessibilityAXValue `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        bool                  `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      string                `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a <label> element. enum values: description, figcaption, label, labelfor, labelwrapped, legend, rubyannotation, tablecaption, title, other
	NativeSourceValue *AccessibilityAXValue `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           bool                  `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     string                `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// No Description.
type AccessibilityAXRelatedNode struct {
	BackendDOMNodeId int    `json:"backendDOMNodeId"` // The BackendNodeId of the related DOM node.
	Idref            string `json:"idref,omitempty"`  // The IDRef value provided, if any.
	Text             string `json:"text,omitempty"`   // The text alternative of this node in the current context.
}

// No Description.
type AccessibilityAXProperty struct {
	Name  string                `json:"name"`  // The name of this property. enum values: busy, disabled, editable, focusable, focused, hidden, hiddenRoot, invalid, keyshortcuts, settable, roledescription, live, atomic, relevant, root, autocomplete, hasPopup, level, multiselectable, orientation, multiline, readonly, required, valuemin, valuemax, valuetext, checked, expanded, modal, pressed, selected, activedescendant, controls, describedby, details, errormessage, flowto, labelledby, owns
	Value *AccessibilityAXValue `json:"value"` // The value of this property.
}

// A single computed AX property.
type AccessibilityAXValue struct {
	Type         string                        `json:"type"`                   // The type of this value. enum values: boolean, tristate, booleanOrUndefined, idref, idrefList, integer, node, nodeList, number, string, computedString, token, tokenList, domRelation, role, internalRole, valueUndefined
	Value        interface{}                   `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []*AccessibilityAXRelatedNode `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []*AccessibilityAXValueSource `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
}

// A node in the accessibility tree.
type AccessibilityAXNode struct {
	NodeId           string                     `json:"nodeId"`                     // Unique identifier for this node.
	Ignored          bool                       `json:"ignored"`                    // Whether this node is ignored for accessibility
	IgnoredReasons   []*AccessibilityAXProperty `json:"ignoredReasons,omitempty"`   // Collection of reasons why this node is hidden.
	Role             *AccessibilityAXValue      `json:"role,omitempty"`             // This `Node`'s role, whether explicit or implicit.
	ChromeRole       *AccessibilityAXValue      `json:"chromeRole,omitempty"`       // This `Node`'s Chrome raw role.
	Name             *AccessibilityAXValue      `json:"name,omitempty"`             // The accessible name for this `Node`.
	Description      *AccessibilityAXValue      `json:"description,omitempty"`      // The accessible description for this `Node`.
	Value            *AccessibilityAXValue      `json:"value,omitempty"`            // The value for this `Node`.
	Properties       []*AccessibilityAXProperty `json:"properties,omitempty"`       // All other properties
	ParentId         string                     `json:"parentId,omitempty"`         // ID for this node's parent.
	ChildIds         []string                   `json:"childIds,omitempty"`         // IDs for each of this node's child nodes.
	BackendDOMNodeId int                        `json:"backendDOMNodeId,omitempty"` // The backend ID for the associated DOM node, if any.
	FrameId          string                     `json:"frameId,omitempty"`          // The frame ID for the frame associated with this nodes document.
}

// The loadComplete event mirrors the load complete event sent by the browser to assistive technology when the web page has finished loading.
type AccessibilityLoadCompleteEvent struct {
	Method string `json:"method"`
	Params struct {
		Root *AccessibilityAXNode `json:"root"` // New document root node.
	} `json:"Params,omitempty"`
}

// The nodesUpdated event is sent every time a previously requested node has changed the in tree.
type AccessibilityNodesUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Nodes []*AccessibilityAXNode `json:"nodes"` // Updated node data.
	} `json:"Params,omitempty"`
}

type Accessibility struct {
	target gcdmessage.ChromeTargeter
}

func NewAccessibility(target gcdmessage.ChromeTargeter) *Accessibility {
	c := &Accessibility{target: target}
	return c
}

// Disables the accessibility domain.
func (c *Accessibility) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.disable"})
}

// Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls. This turns on accessibility for the page, which can impact performance until accessibility is disabled.
func (c *Accessibility) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.enable"})
}

type AccessibilityGetPartialAXTreeParams struct {
	// Identifier of the node to get the partial accessibility tree for.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectId string `json:"objectId,omitempty"`
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives bool `json:"fetchRelatives,omitempty"`
}

// GetPartialAXTreeWithParams - Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
// Returns -  nodes - The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
func (c *Accessibility) GetPartialAXTreeWithParams(ctx context.Context, v *AccessibilityGetPartialAXTreeParams) ([]*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getPartialAXTree", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*AccessibilityAXNode
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

	return chromeData.Result.Nodes, nil
}

// GetPartialAXTree - Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
// nodeId - Identifier of the node to get the partial accessibility tree for.
// backendNodeId - Identifier of the backend node to get the partial accessibility tree for.
// objectId - JavaScript object id of the node wrapper to get the partial accessibility tree for.
// fetchRelatives - Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
// Returns -  nodes - The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
func (c *Accessibility) GetPartialAXTree(ctx context.Context, nodeId int, backendNodeId int, objectId string, fetchRelatives bool) ([]*AccessibilityAXNode, error) {
	var v AccessibilityGetPartialAXTreeParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	v.FetchRelatives = fetchRelatives
	return c.GetPartialAXTreeWithParams(ctx, &v)
}

type AccessibilityGetFullAXTreeParams struct {
	// The maximum depth at which descendants of the root node should be retrieved. If omitted, the full tree is returned.
	Depth int `json:"depth,omitempty"`
	// The frame for whose document the AX tree should be retrieved. If omited, the root frame is used.
	FrameId string `json:"frameId,omitempty"`
}

// GetFullAXTreeWithParams - Fetches the entire accessibility tree for the root Document
// Returns -  nodes -
func (c *Accessibility) GetFullAXTreeWithParams(ctx context.Context, v *AccessibilityGetFullAXTreeParams) ([]*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getFullAXTree", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*AccessibilityAXNode
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

	return chromeData.Result.Nodes, nil
}

// GetFullAXTree - Fetches the entire accessibility tree for the root Document
// depth - The maximum depth at which descendants of the root node should be retrieved. If omitted, the full tree is returned.
// frameId - The frame for whose document the AX tree should be retrieved. If omited, the root frame is used.
// Returns -  nodes -
func (c *Accessibility) GetFullAXTree(ctx context.Context, depth int, frameId string) ([]*AccessibilityAXNode, error) {
	var v AccessibilityGetFullAXTreeParams
	v.Depth = depth
	v.FrameId = frameId
	return c.GetFullAXTreeWithParams(ctx, &v)
}

type AccessibilityGetRootAXNodeParams struct {
	// The frame in whose document the node resides. If omitted, the root frame is used.
	FrameId string `json:"frameId,omitempty"`
}

// GetRootAXNodeWithParams - Fetches the root node. Requires `enable()` to have been called previously.
// Returns -  node -
func (c *Accessibility) GetRootAXNodeWithParams(ctx context.Context, v *AccessibilityGetRootAXNodeParams) (*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getRootAXNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Node *AccessibilityAXNode
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

	return chromeData.Result.Node, nil
}

// GetRootAXNode - Fetches the root node. Requires `enable()` to have been called previously.
// frameId - The frame in whose document the node resides. If omitted, the root frame is used.
// Returns -  node -
func (c *Accessibility) GetRootAXNode(ctx context.Context, frameId string) (*AccessibilityAXNode, error) {
	var v AccessibilityGetRootAXNodeParams
	v.FrameId = frameId
	return c.GetRootAXNodeWithParams(ctx, &v)
}

type AccessibilityGetAXNodeAndAncestorsParams struct {
	// Identifier of the node to get.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node to get.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper to get.
	ObjectId string `json:"objectId,omitempty"`
}

// GetAXNodeAndAncestorsWithParams - Fetches a node and all ancestors up to and including the root. Requires `enable()` to have been called previously.
// Returns -  nodes -
func (c *Accessibility) GetAXNodeAndAncestorsWithParams(ctx context.Context, v *AccessibilityGetAXNodeAndAncestorsParams) ([]*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getAXNodeAndAncestors", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*AccessibilityAXNode
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

	return chromeData.Result.Nodes, nil
}

// GetAXNodeAndAncestors - Fetches a node and all ancestors up to and including the root. Requires `enable()` to have been called previously.
// nodeId - Identifier of the node to get.
// backendNodeId - Identifier of the backend node to get.
// objectId - JavaScript object id of the node wrapper to get.
// Returns -  nodes -
func (c *Accessibility) GetAXNodeAndAncestors(ctx context.Context, nodeId int, backendNodeId int, objectId string) ([]*AccessibilityAXNode, error) {
	var v AccessibilityGetAXNodeAndAncestorsParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.GetAXNodeAndAncestorsWithParams(ctx, &v)
}

type AccessibilityGetChildAXNodesParams struct {
	//
	Id string `json:"id"`
	// The frame in whose document the node resides. If omitted, the root frame is used.
	FrameId string `json:"frameId,omitempty"`
}

// GetChildAXNodesWithParams - Fetches a particular accessibility node by AXNodeId. Requires `enable()` to have been called previously.
// Returns -  nodes -
func (c *Accessibility) GetChildAXNodesWithParams(ctx context.Context, v *AccessibilityGetChildAXNodesParams) ([]*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.getChildAXNodes", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*AccessibilityAXNode
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

	return chromeData.Result.Nodes, nil
}

// GetChildAXNodes - Fetches a particular accessibility node by AXNodeId. Requires `enable()` to have been called previously.
// id -
// frameId - The frame in whose document the node resides. If omitted, the root frame is used.
// Returns -  nodes -
func (c *Accessibility) GetChildAXNodes(ctx context.Context, id string, frameId string) ([]*AccessibilityAXNode, error) {
	var v AccessibilityGetChildAXNodesParams
	v.Id = id
	v.FrameId = frameId
	return c.GetChildAXNodesWithParams(ctx, &v)
}

type AccessibilityQueryAXTreeParams struct {
	// Identifier of the node for the root to query.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node for the root to query.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper for the root to query.
	ObjectId string `json:"objectId,omitempty"`
	// Find nodes with this computed name.
	AccessibleName string `json:"accessibleName,omitempty"`
	// Find nodes with this computed role.
	Role string `json:"role,omitempty"`
}

// QueryAXTreeWithParams - Query a DOM node's accessibility subtree for accessible name and role. This command computes the name and role for all nodes in the subtree, including those that are ignored for accessibility, and returns those that mactch the specified name and role. If no DOM node is specified, or the DOM node does not exist, the command returns an error. If neither `accessibleName` or `role` is specified, it returns all the accessibility nodes in the subtree.
// Returns -  nodes - A list of `Accessibility.AXNode` matching the specified attributes, including nodes that are ignored for accessibility.
func (c *Accessibility) QueryAXTreeWithParams(ctx context.Context, v *AccessibilityQueryAXTreeParams) ([]*AccessibilityAXNode, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Accessibility.queryAXTree", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*AccessibilityAXNode
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

	return chromeData.Result.Nodes, nil
}

// QueryAXTree - Query a DOM node's accessibility subtree for accessible name and role. This command computes the name and role for all nodes in the subtree, including those that are ignored for accessibility, and returns those that mactch the specified name and role. If no DOM node is specified, or the DOM node does not exist, the command returns an error. If neither `accessibleName` or `role` is specified, it returns all the accessibility nodes in the subtree.
// nodeId - Identifier of the node for the root to query.
// backendNodeId - Identifier of the backend node for the root to query.
// objectId - JavaScript object id of the node wrapper for the root to query.
// accessibleName - Find nodes with this computed name.
// role - Find nodes with this computed role.
// Returns -  nodes - A list of `Accessibility.AXNode` matching the specified attributes, including nodes that are ignored for accessibility.
func (c *Accessibility) QueryAXTree(ctx context.Context, nodeId int, backendNodeId int, objectId string, accessibleName string, role string) ([]*AccessibilityAXNode, error) {
	var v AccessibilityQueryAXTreeParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	v.AccessibleName = accessibleName
	v.Role = role
	return c.QueryAXTreeWithParams(ctx, &v)
}
