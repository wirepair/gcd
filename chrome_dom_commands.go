// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOM commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) DOM() *ChromeDOM {
	if c.dom == nil {
		c.dom = newChromeDOM(c)
	}
	return c.dom
}

type ChromeDOM struct {
	target *ChromeTarget
}

func newChromeDOM(target *ChromeTarget) *ChromeDOM {
	c := &ChromeDOM{target: target}
	return c
}

// Enables DOM agent for the given page.
func (c *ChromeDOM) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.enable"})
}

// Disables DOM agent for the given page.
func (c *ChromeDOM) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.disable"})
}

// Hides DOM node highlight.
func (c *ChromeDOM) HideHighlight() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.hideHighlight"})
}

// Undoes the last performed action.
func (c *ChromeDOM) Undo() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.undo"})
}

// Re-does the last undone action.
func (c *ChromeDOM) Redo() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.redo"})
}

// Marks last undoable state.
func (c *ChromeDOM) MarkUndoableState() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.markUndoableState"})
}

// requestChildNodes - Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
// nodeId - Id of the node to get children for.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
func (c *ChromeDOM) RequestChildNodes(nodeId *types.ChromeDOMNodeId, depth int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["depth"] = depth
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.requestChildNodes", Params: paramRequest})
}

// setNodeValue - Sets node value for a node with given id.
// nodeId - Id of the node to set value for.
// value - New node's value.
func (c *ChromeDOM) SetNodeValue(nodeId *types.ChromeDOMNodeId, value string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setNodeValue", Params: paramRequest})
}

// removeNode - Removes node with given id.
// nodeId - Id of the node to remove.
func (c *ChromeDOM) RemoveNode(nodeId *types.ChromeDOMNodeId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.removeNode", Params: paramRequest})
}

// setAttributeValue - Sets attribute for an element with given id.
// nodeId - Id of the element to set attribute for.
// name - Attribute name.
// value - Attribute value.
func (c *ChromeDOM) SetAttributeValue(nodeId *types.ChromeDOMNodeId, name string, value string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setAttributeValue", Params: paramRequest})
}

// setAttributesAsText - Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
// nodeId - Id of the element to set attributes for.
// text - Text with a number of attributes. Will parse this text using HTML parser.
// name - Attribute name to replace with new attributes derived from text in case text parsed successfully.
func (c *ChromeDOM) SetAttributesAsText(nodeId *types.ChromeDOMNodeId, text string, name string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["text"] = text
	paramRequest["name"] = name
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setAttributesAsText", Params: paramRequest})
}

// removeAttribute - Removes attribute with given name from an element with given id.
// nodeId - Id of the element to remove attribute from.
// name - Name of the attribute to remove.
func (c *ChromeDOM) RemoveAttribute(nodeId *types.ChromeDOMNodeId, name string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.removeAttribute", Params: paramRequest})
}

// setOuterHTML - Sets node HTML markup, returns new node id.
// nodeId - Id of the node to set markup for.
// outerHTML - Outer HTML markup to set.
func (c *ChromeDOM) SetOuterHTML(nodeId *types.ChromeDOMNodeId, outerHTML string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["outerHTML"] = outerHTML
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setOuterHTML", Params: paramRequest})
}

// discardSearchResults - Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search.
// searchId - Unique search session identifier.
func (c *ChromeDOM) DiscardSearchResults(searchId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["searchId"] = searchId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.discardSearchResults", Params: paramRequest})
}

// setInspectModeEnabled - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
// enabled - True to enable inspection mode, false to disable it.
// inspectUAShadowDOM - True to enable inspection mode for user agent shadow DOM.
// highlightConfig - A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
func (c *ChromeDOM) SetInspectModeEnabled(enabled bool, inspectUAShadowDOM bool, highlightConfig *types.ChromeDOMHighlightConfig) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["enabled"] = enabled
	paramRequest["inspectUAShadowDOM"] = inspectUAShadowDOM
	paramRequest["highlightConfig"] = highlightConfig
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setInspectModeEnabled", Params: paramRequest})
}

// highlightRect - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
// x - X coordinate
// y - Y coordinate
// width - Rectangle width
// height - Rectangle height
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *ChromeDOM) HighlightRect(x int, y int, width int, height int, color *types.ChromeDOMRGBA, outlineColor *types.ChromeDOMRGBA) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 6)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.highlightRect", Params: paramRequest})
}

// highlightQuad - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
// quad - Quad to highlight
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *ChromeDOM) HighlightQuad(quad *types.ChromeDOMQuad, color *types.ChromeDOMRGBA, outlineColor *types.ChromeDOMRGBA) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["quad"] = quad
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.highlightQuad", Params: paramRequest})
}

// highlightNode - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// highlightConfig - A descriptor for the highlight appearance.
// nodeId - Identifier of the node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
func (c *ChromeDOM) HighlightNode(highlightConfig *types.ChromeDOMHighlightConfig, nodeId *types.ChromeDOMNodeId, objectId *types.ChromeRuntimeRemoteObjectId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["highlightConfig"] = highlightConfig
	paramRequest["nodeId"] = nodeId
	paramRequest["objectId"] = objectId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.highlightNode", Params: paramRequest})
}

// highlightFrame - Highlights owner element of the frame with given id.
// frameId - Identifier of the frame to highlight.
// contentColor - The content box highlight fill color (default: transparent).
// contentOutlineColor - The content box highlight outline color (default: transparent).
func (c *ChromeDOM) HighlightFrame(frameId *types.ChromePageFrameId, contentColor *types.ChromeDOMRGBA, contentOutlineColor *types.ChromeDOMRGBA) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["frameId"] = frameId
	paramRequest["contentColor"] = contentColor
	paramRequest["contentOutlineColor"] = contentOutlineColor
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.highlightFrame", Params: paramRequest})
}

// focus - Focuses the given element.
// nodeId - Id of the node to focus.
func (c *ChromeDOM) Focus(nodeId *types.ChromeDOMNodeId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.focus", Params: paramRequest})
}

// setFileInputFiles - Sets files for the given file input element.
// nodeId - Id of the file input node to set files for.
// files - Array of file paths to set.
func (c *ChromeDOM) SetFileInputFiles(nodeId *types.ChromeDOMNodeId, files []string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["files"] = files
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setFileInputFiles", Params: paramRequest})
}

// getDocument - Returns the root DOM node to the caller.
// Returns -
// Resulting node.
func (c *ChromeDOM) GetDocument() (*types.ChromeDOMNode, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getDocument"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Root *types.ChromeDOMNode
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

	return chromeData.Result.Root, nil
}

// querySelector - Executes <code>querySelector</code> on a given node.
// Returns -
// Query selector result.
func (c *ChromeDOM) QuerySelector(nodeId *types.ChromeDOMNodeId, selector string) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["selector"] = selector
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.querySelector", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// querySelectorAll - Executes <code>querySelectorAll</code> on a given node.
// Returns -
// Query selector result.
func (c *ChromeDOM) QuerySelectorAll(nodeId *types.ChromeDOMNodeId, selector string) ([]*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["selector"] = selector
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.querySelectorAll", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []*types.ChromeDOMNodeId
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

	return chromeData.Result.NodeIds, nil
}

// setNodeName - Sets node name for a node with given id.
// Returns -
// New node's id.
func (c *ChromeDOM) SetNodeName(nodeId *types.ChromeDOMNodeId, name string) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["name"] = name
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setNodeName", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// getEventListenersForNode - Returns event listeners relevant to the node.
// Returns -
// Array of relevant listeners.
func (c *ChromeDOM) GetEventListenersForNode(nodeId *types.ChromeDOMNodeId, objectGroup string) ([]*types.ChromeDOMEventListener, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getEventListenersForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Listeners []*types.ChromeDOMEventListener
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

	return chromeData.Result.Listeners, nil
}

// getOuterHTML - Returns node's HTML markup.
// Returns -
// Outer HTML markup.
func (c *ChromeDOM) GetOuterHTML(nodeId *types.ChromeDOMNodeId) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getOuterHTML", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			OuterHTML string
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

	return chromeData.Result.OuterHTML, nil
}

// performSearch - Searches for a given string in the DOM tree. Use <code>getSearchResults</code> to access search results or <code>cancelSearch</code> to end this search session.
// Returns -
// Unique search session identifier.
// Number of search results.
func (c *ChromeDOM) PerformSearch(query string) (string, float64, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["query"] = query
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.performSearch", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			SearchId    string
			ResultCount float64
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", 0, &ChromeRequestErr{Resp: cerr}
		}
		return "", 0, err
	}

	return chromeData.Result.SearchId, chromeData.Result.ResultCount, nil
}

// getSearchResults - Returns search results from given <code>fromIndex</code> to given <code>toIndex</code> from the sarch with the given identifier.
// Returns -
// Ids of the search result nodes.
func (c *ChromeDOM) GetSearchResults(searchId string, fromIndex int, toIndex int) ([]*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["searchId"] = searchId
	paramRequest["fromIndex"] = fromIndex
	paramRequest["toIndex"] = toIndex
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getSearchResults", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []*types.ChromeDOMNodeId
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

	return chromeData.Result.NodeIds, nil
}

// requestNode - Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of <code>setChildNodes</code> notifications.
// Returns -
// Node id for given object.
func (c *ChromeDOM) RequestNode(objectId *types.ChromeRuntimeRemoteObjectId) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.requestNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// pushNodeByPathToFrontend - Requests that the node is sent to the caller given its path. // FIXME, use XPath
// Returns -
// Id of the node for given path.
func (c *ChromeDOM) PushNodeByPathToFrontend(path string) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["path"] = path
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.pushNodeByPathToFrontend", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// pushNodesByBackendIdsToFrontend - Requests that a batch of nodes is sent to the caller given their backend node ids.
// Returns -
// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
func (c *ChromeDOM) PushNodesByBackendIdsToFrontend(backendNodeIds []*types.ChromeDOMBackendNodeId) ([]*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["backendNodeIds"] = backendNodeIds
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.pushNodesByBackendIdsToFrontend", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeIds []*types.ChromeDOMNodeId
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

	return chromeData.Result.NodeIds, nil
}

// resolveNode - Resolves JavaScript node object for given node id.
// Returns -
// JavaScript object wrapper for given node.
func (c *ChromeDOM) ResolveNode(nodeId *types.ChromeDOMNodeId, objectGroup string) (*types.ChromeRuntimeRemoteObject, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["objectGroup"] = objectGroup
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.resolveNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Object *types.ChromeRuntimeRemoteObject
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

	return chromeData.Result.Object, nil
}

// getAttributes - Returns attributes for the specified node.
// Returns -
// An interleaved array of node attribute names and values.
func (c *ChromeDOM) GetAttributes(nodeId *types.ChromeDOMNodeId) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getAttributes", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Attributes []string
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

	return chromeData.Result.Attributes, nil
}

// copyTo - Creates a deep copy of the specified node and places it into the target container before the given anchor.
// Returns -
// Id of the node clone.
func (c *ChromeDOM) CopyTo(nodeId *types.ChromeDOMNodeId, targetNodeId *types.ChromeDOMNodeId, insertBeforeNodeId *types.ChromeDOMNodeId) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["targetNodeId"] = targetNodeId
	paramRequest["insertBeforeNodeId"] = insertBeforeNodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.copyTo", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// moveTo - Moves node into the new container, places it before the given anchor.
// Returns -
// New id of the moved node.
func (c *ChromeDOM) MoveTo(nodeId *types.ChromeDOMNodeId, targetNodeId *types.ChromeDOMNodeId, insertBeforeNodeId *types.ChromeDOMNodeId) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["targetNodeId"] = targetNodeId
	paramRequest["insertBeforeNodeId"] = insertBeforeNodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.moveTo", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// getBoxModel - Returns boxes for the currently selected nodes.
// Returns -
// Box model for the node.
func (c *ChromeDOM) GetBoxModel(nodeId *types.ChromeDOMNodeId) (*types.ChromeDOMBoxModel, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getBoxModel", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Model *types.ChromeDOMBoxModel
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

	return chromeData.Result.Model, nil
}

// getNodeForLocation - Returns node id at given location.
// Returns -
// Id of the node at given coordinates.
func (c *ChromeDOM) GetNodeForLocation(x int, y int) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["x"] = x
	paramRequest["y"] = y
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getNodeForLocation", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}

// getRelayoutBoundary - Returns the id of the nearest ancestor that is a relayout boundary.
// Returns -
// Relayout boundary node id for the given node.
func (c *ChromeDOM) GetRelayoutBoundary(nodeId *types.ChromeDOMNodeId) (*types.ChromeDOMNodeId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getRelayoutBoundary", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			NodeId *types.ChromeDOMNodeId
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

	return chromeData.Result.NodeId, nil
}
