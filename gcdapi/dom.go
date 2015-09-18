// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOM functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Backend node with a friendly name.
type DOMBackendNode struct {
	NodeType      int    `json:"nodeType"`      // <code>Node</code>'s nodeType.
	NodeName      string `json:"nodeName"`      // <code>Node</code>'s nodeName.
	BackendNodeId int    `json:"backendNodeId"` //
}

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type DOMNode struct {
	NodeId           int               `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeType         int               `json:"nodeType"`                   // <code>Node</code>'s nodeType.
	NodeName         string            `json:"nodeName"`                   // <code>Node</code>'s nodeName.
	LocalName        string            `json:"localName"`                  // <code>Node</code>'s localName.
	NodeValue        string            `json:"nodeValue"`                  // <code>Node</code>'s nodeValue.
	ChildNodeCount   int               `json:"childNodeCount,omitempty"`   // Child count for <code>Container</code> nodes.
	Children         []*DOMNode        `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string          `json:"attributes,omitempty"`       // Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
	DocumentURL      string            `json:"documentURL,omitempty"`      // Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL          string            `json:"baseURL,omitempty"`          // Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	PublicId         string            `json:"publicId,omitempty"`         // <code>DocumentType</code>'s publicId.
	SystemId         string            `json:"systemId,omitempty"`         // <code>DocumentType</code>'s systemId.
	InternalSubset   string            `json:"internalSubset,omitempty"`   // <code>DocumentType</code>'s internalSubset.
	XmlVersion       string            `json:"xmlVersion,omitempty"`       // <code>Document</code>'s XML version in case of XML documents.
	Name             string            `json:"name,omitempty"`             // <code>Attr</code>'s name.
	Value            string            `json:"value,omitempty"`            // <code>Attr</code>'s value.
	PseudoType       string            `json:"pseudoType,omitempty"`       // Pseudo element type for this node. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	ShadowRootType   string            `json:"shadowRootType,omitempty"`   // Shadow root type. enum values: user-agent, open, closed
	FrameId          string            `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *DOMNode          `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []*DOMNode        `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *DOMNode          `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []*DOMNode        `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *DOMNode          `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []*DOMBackendNode `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
}

// A structure holding an RGBA color.
type DOMRGBA struct {
	R int     `json:"r"`           // The red component, in the [0-255] range.
	G int     `json:"g"`           // The green component, in the [0-255] range.
	B int     `json:"b"`           // The blue component, in the [0-255] range.
	A float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// Box model.
type DOMBoxModel struct {
	Content      []float64            `json:"content"`                // Content box
	Padding      []float64            `json:"padding"`                // Padding box
	Border       []float64            `json:"border"`                 // Border box
	Margin       []float64            `json:"margin"`                 // Margin box
	Width        int                  `json:"width"`                  // Node width
	Height       int                  `json:"height"`                 // Node height
	ShapeOutside *DOMShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// CSS Shape Outside details.
type DOMShapeOutsideInfo struct {
	Bounds      []float64 `json:"bounds"`      // Shape bounds
	Shape       []string  `json:"shape"`       // Shape coordinate details
	MarginShape []string  `json:"marginShape"` // Margin shape bounds
}

// Rectangle.
type DOMRect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// Configuration data for the highlighting of page elements.
type DOMHighlightConfig struct {
	ShowInfo           bool     `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         bool     `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool     `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  bool     `json:"displayAsMaterial,omitempty"`  //
	ContentColor       *DOMRGBA `json:"contentColor,omitempty"`       // The content box highlight fill color (default: transparent).
	PaddingColor       *DOMRGBA `json:"paddingColor,omitempty"`       // The padding highlight fill color (default: transparent).
	BorderColor        *DOMRGBA `json:"borderColor,omitempty"`        // The border highlight fill color (default: transparent).
	MarginColor        *DOMRGBA `json:"marginColor,omitempty"`        // The margin highlight fill color (default: transparent).
	EventTargetColor   *DOMRGBA `json:"eventTargetColor,omitempty"`   // The event target element highlight fill color (default: transparent).
	ShapeColor         *DOMRGBA `json:"shapeColor,omitempty"`         // The shape outside fill color (default: transparent).
	ShapeMarginColor   *DOMRGBA `json:"shapeMarginColor,omitempty"`   // The shape margin fill color (default: transparent).
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code>.
type DOMInspectNodeRequestedEvent struct {
	BackendNodeId int `json:"backendNodeId"` // Id of the node to inspect.
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon most of the calls requesting node ids.
type DOMSetChildNodesEvent struct {
	ParentId int        `json:"parentId"` // Parent node id to populate with children.
	Nodes    []*DOMNode `json:"nodes"`    // Child nodes array.
}

// Fired when <code>Element</code>'s attribute is modified.
type DOMAttributeModifiedEvent struct {
	NodeId int    `json:"nodeId"` // Id of the node that has changed.
	Name   string `json:"name"`   // Attribute name.
	Value  string `json:"value"`  // Attribute value.
}

// Fired when <code>Element</code>'s attribute is removed.
type DOMAttributeRemovedEvent struct {
	NodeId int    `json:"nodeId"` // Id of the node that has changed.
	Name   string `json:"name"`   // A ttribute name.
}

// Fired when <code>Element</code>'s inline style is modified via a CSS property modification.
type DOMInlineStyleInvalidatedEvent struct {
	NodeIds []int `json:"nodeIds"` // Ids of the nodes for which the inline styles have been invalidated.
}

// Mirrors <code>DOMCharacterDataModified</code> event.
type DOMCharacterDataModifiedEvent struct {
	NodeId        int    `json:"nodeId"`        // Id of the node that has changed.
	CharacterData string `json:"characterData"` // New text value.
}

// Fired when <code>Container</code>'s child node count has changed.
type DOMChildNodeCountUpdatedEvent struct {
	NodeId         int `json:"nodeId"`         // Id of the node that has changed.
	ChildNodeCount int `json:"childNodeCount"` // New node count.
}

// Mirrors <code>DOMNodeInserted</code> event.
type DOMChildNodeInsertedEvent struct {
	ParentNodeId   int      `json:"parentNodeId"`   // Id of the node that has changed.
	PreviousNodeId int      `json:"previousNodeId"` // If of the previous siblint.
	Node           *DOMNode `json:"node"`           // Inserted node data.
}

// Mirrors <code>DOMNodeRemoved</code> event.
type DOMChildNodeRemovedEvent struct {
	ParentNodeId int `json:"parentNodeId"` // Parent id.
	NodeId       int `json:"nodeId"`       // Id of the node that has been removed.
}

// Called when shadow root is pushed into the element.
type DOMShadowRootPushedEvent struct {
	HostId int      `json:"hostId"` // Host element id.
	Root   *DOMNode `json:"root"`   // Shadow root.
}

// Called when shadow root is popped from the element.
type DOMShadowRootPoppedEvent struct {
	HostId int `json:"hostId"` // Host element id.
	RootId int `json:"rootId"` // Shadow root id.
}

// Called when a pseudo element is added to an element.
type DOMPseudoElementAddedEvent struct {
	ParentId      int      `json:"parentId"`      // Pseudo element's parent element id.
	PseudoElement *DOMNode `json:"pseudoElement"` // The added pseudo element.
}

// Called when a pseudo element is removed from an element.
type DOMPseudoElementRemovedEvent struct {
	ParentId        int `json:"parentId"`        // Pseudo element's parent element id.
	PseudoElementId int `json:"pseudoElementId"` // The removed pseudo element id.
}

// Called when distrubution is changed.
type DOMDistributedNodesUpdatedEvent struct {
	InsertionPointId int               `json:"insertionPointId"` // Insertion point where distrubuted nodes were updated.
	DistributedNodes []*DOMBackendNode `json:"distributedNodes"` // Distributed nodes for given insertion point.
}

type DOM struct {
	target gcdmessage.ChromeTargeter
}

func NewDOM(target gcdmessage.ChromeTargeter) *DOM {
	c := &DOM{target: target}
	return c
}

// Enables DOM agent for the given page.
func (c *DOM) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.enable"})
}

// Disables DOM agent for the given page.
func (c *DOM) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.disable"})
}

// GetDocument - Returns the root DOM node to the caller.
// Returns -  root - Resulting node.
func (c *DOM) GetDocument() (*DOMNode, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getDocument"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Root *DOMNode
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

	return chromeData.Result.Root, nil
}

// RequestChildNodes - Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
// nodeId - Id of the node to get children for.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
func (c *DOM) RequestChildNodes(nodeId int, depth int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["depth"] = depth
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.requestChildNodes", Params: paramRequest})
}

// QuerySelector - Executes <code>querySelector</code> on a given node.
// nodeId - Id of the node to query upon.
// selector - Selector string.
// Returns -  nodeId - Query selector result.
func (c *DOM) QuerySelector(nodeId int, selector string) (int, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["selector"] = selector
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.querySelector", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// QuerySelectorAll - Executes <code>querySelectorAll</code> on a given node.
// nodeId - Id of the node to query upon.
// selector - Selector string.
// Returns -  nodeIds - Query selector result.
func (c *DOM) QuerySelectorAll(nodeId int, selector string) ([]int, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["selector"] = selector
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.querySelectorAll", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// SetNodeName - Sets node name for a node with given id.
// nodeId - Id of the node to set name for.
// name - New node's name.
// Returns -  nodeId - New node's id.
func (c *DOM) SetNodeName(nodeId int, name string) (int, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setNodeName", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// SetNodeValue - Sets node value for a node with given id.
// nodeId - Id of the node to set value for.
// value - New node's value.
func (c *DOM) SetNodeValue(nodeId int, value string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setNodeValue", Params: paramRequest})
}

// RemoveNode - Removes node with given id.
// nodeId - Id of the node to remove.
func (c *DOM) RemoveNode(nodeId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.removeNode", Params: paramRequest})
}

// SetAttributeValue - Sets attribute for an element with given id.
// nodeId - Id of the element to set attribute for.
// name - Attribute name.
// value - Attribute value.
func (c *DOM) SetAttributeValue(nodeId int, name string, value string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setAttributeValue", Params: paramRequest})
}

// SetAttributesAsText - Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
// nodeId - Id of the element to set attributes for.
// text - Text with a number of attributes. Will parse this text using HTML parser.
// name - Attribute name to replace with new attributes derived from text in case text parsed successfully.
func (c *DOM) SetAttributesAsText(nodeId int, text string, name string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["text"] = text
	paramRequest["name"] = name
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setAttributesAsText", Params: paramRequest})
}

// RemoveAttribute - Removes attribute with given name from an element with given id.
// nodeId - Id of the element to remove attribute from.
// name - Name of the attribute to remove.
func (c *DOM) RemoveAttribute(nodeId int, name string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.removeAttribute", Params: paramRequest})
}

// GetOuterHTML - Returns node's HTML markup.
// nodeId - Id of the node to get markup for.
// Returns -  outerHTML - Outer HTML markup.
func (c *DOM) GetOuterHTML(nodeId int) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getOuterHTML", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			OuterHTML string
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.OuterHTML, nil
}

// SetOuterHTML - Sets node HTML markup, returns new node id.
// nodeId - Id of the node to set markup for.
// outerHTML - Outer HTML markup to set.
func (c *DOM) SetOuterHTML(nodeId int, outerHTML string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["outerHTML"] = outerHTML
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setOuterHTML", Params: paramRequest})
}

// PerformSearch - Searches for a given string in the DOM tree. Use <code>getSearchResults</code> to access search results or <code>cancelSearch</code> to end this search session.
// query - Plain text or query selector or XPath search query.
// includeUserAgentShadowDOM - True to search in user agent shadow DOM.
// Returns -  searchId - Unique search session identifier. resultCount - Number of search results.
func (c *DOM) PerformSearch(query string, includeUserAgentShadowDOM bool) (string, int, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["query"] = query
	paramRequest["includeUserAgentShadowDOM"] = includeUserAgentShadowDOM
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.performSearch", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			SearchId    string
			ResultCount int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", 0, err
	}

	return chromeData.Result.SearchId, chromeData.Result.ResultCount, nil
}

// GetSearchResults - Returns search results from given <code>fromIndex</code> to given <code>toIndex</code> from the sarch with the given identifier.
// searchId - Unique search session identifier.
// fromIndex - Start index of the search result to be returned.
// toIndex - End index of the search result to be returned.
// Returns -  nodeIds - Ids of the search result nodes.
func (c *DOM) GetSearchResults(searchId string, fromIndex int, toIndex int) ([]int, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["searchId"] = searchId
	paramRequest["fromIndex"] = fromIndex
	paramRequest["toIndex"] = toIndex
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getSearchResults", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// DiscardSearchResults - Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search.
// searchId - Unique search session identifier.
func (c *DOM) DiscardSearchResults(searchId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["searchId"] = searchId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.discardSearchResults", Params: paramRequest})
}

// RequestNode - Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of <code>setChildNodes</code> notifications.
// objectId - JavaScript object id to convert into node.
// Returns -  nodeId - Node id for given object.
func (c *DOM) RequestNode(objectId string) (int, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.requestNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// SetInspectMode - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
// mode - Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, showLayoutEditor, none
// highlightConfig - A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
func (c *DOM) SetInspectMode(mode string, highlightConfig *DOMHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["mode"] = mode
	paramRequest["highlightConfig"] = highlightConfig
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setInspectMode", Params: paramRequest})
}

// HighlightRect - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
// x - X coordinate
// y - Y coordinate
// width - Rectangle width
// height - Rectangle height
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *DOM) HighlightRect(x int, y int, width int, height int, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 6)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightRect", Params: paramRequest})
}

// HighlightQuad - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
// quad - Quad to highlight
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *DOM) HighlightQuad(quad float64, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["quad"] = quad
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightQuad", Params: paramRequest})
}

// HighlightNode - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// highlightConfig - A descriptor for the highlight appearance.
// nodeId - Identifier of the node to highlight.
// backendNodeId - Identifier of the backend node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
func (c *DOM) HighlightNode(highlightConfig *DOMHighlightConfig, nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["highlightConfig"] = highlightConfig
	paramRequest["nodeId"] = nodeId
	paramRequest["backendNodeId"] = backendNodeId
	paramRequest["objectId"] = objectId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightNode", Params: paramRequest})
}

// Hides DOM node highlight.
func (c *DOM) HideHighlight() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.hideHighlight"})
}

// HighlightFrame - Highlights owner element of the frame with given id.
// frameId - Identifier of the frame to highlight.
// contentColor - The content box highlight fill color (default: transparent).
// contentOutlineColor - The content box highlight outline color (default: transparent).
func (c *DOM) HighlightFrame(frameId string, contentColor *DOMRGBA, contentOutlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["frameId"] = frameId
	paramRequest["contentColor"] = contentColor
	paramRequest["contentOutlineColor"] = contentOutlineColor
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightFrame", Params: paramRequest})
}

// PushNodeByPathToFrontend - Requests that the node is sent to the caller given its path. // FIXME, use XPath
// path - Path to node in the proprietary format.
// Returns -  nodeId - Id of the node for given path.
func (c *DOM) PushNodeByPathToFrontend(path string) (int, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["path"] = path
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.pushNodeByPathToFrontend", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// PushNodesByBackendIdsToFrontend - Requests that a batch of nodes is sent to the caller given their backend node ids.
// backendNodeIds - The array of backend node ids.
// Returns -  nodeIds - The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
func (c *DOM) PushNodesByBackendIdsToFrontend(backendNodeIds int) ([]int, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["backendNodeIds"] = backendNodeIds
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.pushNodesByBackendIdsToFrontend", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// SetInspectedNode - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// nodeId - DOM node id to be accessible by means of $x command line API.
func (c *DOM) SetInspectedNode(nodeId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setInspectedNode", Params: paramRequest})
}

// ResolveNode - Resolves JavaScript node object for given node id.
// nodeId - Id of the node to resolve.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  object - JavaScript object wrapper for given node.
func (c *DOM) ResolveNode(nodeId int, objectGroup string) (*RuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.resolveNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Object *RuntimeRemoteObject
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

	return chromeData.Result.Object, nil
}

// GetAttributes - Returns attributes for the specified node.
// nodeId - Id of the node to retrieve attibutes for.
// Returns -  attributes - An interleaved array of node attribute names and values.
func (c *DOM) GetAttributes(nodeId int) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getAttributes", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Attributes []string
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

	return chromeData.Result.Attributes, nil
}

// CopyTo - Creates a deep copy of the specified node and places it into the target container before the given anchor.
// nodeId - Id of the node to copy.
// targetNodeId - Id of the element to drop the copy into.
// insertBeforeNodeId - Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>).
// Returns -  nodeId - Id of the node clone.
func (c *DOM) CopyTo(nodeId int, targetNodeId int, insertBeforeNodeId int) (int, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["targetNodeId"] = targetNodeId
	paramRequest["insertBeforeNodeId"] = insertBeforeNodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.copyTo", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// MoveTo - Moves node into the new container, places it before the given anchor.
// nodeId - Id of the node to move.
// targetNodeId - Id of the element to drop the moved node into.
// insertBeforeNodeId - Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>).
// Returns -  nodeId - New id of the moved node.
func (c *DOM) MoveTo(nodeId int, targetNodeId int, insertBeforeNodeId int) (int, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["targetNodeId"] = targetNodeId
	paramRequest["insertBeforeNodeId"] = insertBeforeNodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.moveTo", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// Undoes the last performed action.
func (c *DOM) Undo() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.undo"})
}

// Re-does the last undone action.
func (c *DOM) Redo() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.redo"})
}

// Marks last undoable state.
func (c *DOM) MarkUndoableState() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.markUndoableState"})
}

// Focus - Focuses the given element.
// nodeId - Id of the node to focus.
func (c *DOM) Focus(nodeId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.focus", Params: paramRequest})
}

// SetFileInputFiles - Sets files for the given file input element.
// nodeId - Id of the file input node to set files for.
// files - Array of file paths to set.
func (c *DOM) SetFileInputFiles(nodeId int, files string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["files"] = files
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setFileInputFiles", Params: paramRequest})
}

// GetBoxModel - Returns boxes for the currently selected nodes.
// nodeId - Id of the node to get box model for.
// Returns -  model - Box model for the node.
func (c *DOM) GetBoxModel(nodeId int) (*DOMBoxModel, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getBoxModel", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Model *DOMBoxModel
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

	return chromeData.Result.Model, nil
}

// GetNodeForLocation - Returns node id at given location.
// x - X coordinate.
// y - Y coordinate.
// Returns -  nodeId - Id of the node at given coordinates.
func (c *DOM) GetNodeForLocation(x int, y int) (int, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["x"] = x
	paramRequest["y"] = y
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getNodeForLocation", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// GetRelayoutBoundary - Returns the id of the nearest ancestor that is a relayout boundary.
// nodeId - Id of the node.
// Returns -  nodeId - Relayout boundary node id for the given node.
func (c *DOM) GetRelayoutBoundary(nodeId int) (int, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getRelayoutBoundary", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId int
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return 0, err
	}

	return chromeData.Result.NodeId, nil
}

// GetHighlightObjectForTest - For testing.
// nodeId - Id of the node to get highlight object for.
// Returns -  highlight - Highlight data for the node.
func (c *DOM) GetHighlightObjectForTest(nodeId int) (map[string]interface{}, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getHighlightObjectForTest", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Highlight map[string]interface{}
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

	return chromeData.Result.Highlight, nil
}
