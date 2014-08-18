// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOM commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
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

// start non parameterized commands 
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

// end non parameterized commands

// start parameterized commands with no special return types

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
func (c *ChromeDOM) SetFileInputFiles(nodeId *types.ChromeDOMNodeId, files []types.string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["files"] = files
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.setFileInputFiles", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// getDocument - Returns the root DOM node to the caller.
// Returns - 
// Resulting node.
func (c *ChromeDOM) GetDocument() (*types.ChromeDOMNode, error) {	
	var root *types.ChromeDOMNode 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOM.getDocument"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return root, &ChromeRequestErr{Resp: cerr}
		}
		return root, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		root = results["root"].(*types.ChromeDOMNode)
	}
	return root, nil
}


// end commands with no parameters but special return types

