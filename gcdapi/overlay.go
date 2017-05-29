// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Overlay functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Configuration data for the highlighting of page elements.
type OverlayHighlightConfig struct {
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
	SelectorList       string   `json:"selectorList,omitempty"`       // Selectors to highlight relevant nodes.
}

// Fired when the node should be highlighted. This happens after call to <code>setInspectMode</code>.
type OverlayNodeHighlightRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId int `json:"nodeId"` //
	} `json:"Params,omitempty"`
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code> or when user manually inspects an element.
type OverlayInspectNodeRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		BackendNodeId int `json:"backendNodeId"` // Id of the node to inspect.
	} `json:"Params,omitempty"`
}

type Overlay struct {
	target gcdmessage.ChromeTargeter
}

func NewOverlay(target gcdmessage.ChromeTargeter) *Overlay {
	c := &Overlay{target: target}
	return c
}

// Enables domain notifications.
func (c *Overlay) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.enable"})
}

// Disables domain notifications.
func (c *Overlay) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.disable"})
}

// SetShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *Overlay) SetShowPaintRects(result bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["result"] = result
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowPaintRects", Params: paramRequest})
}

// SetShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *Overlay) SetShowDebugBorders(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowDebugBorders", Params: paramRequest})
}

// SetShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *Overlay) SetShowFPSCounter(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowFPSCounter", Params: paramRequest})
}

// SetShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *Overlay) SetShowScrollBottleneckRects(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowScrollBottleneckRects", Params: paramRequest})
}

// SetShowViewportSizeOnResize - Paints viewport size upon main frame resize.
// show - Whether to paint size or not.
func (c *Overlay) SetShowViewportSizeOnResize(show bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["show"] = show
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowViewportSizeOnResize", Params: paramRequest})
}

// SetPausedInDebuggerMessage -
// message - The message to display, also triggers resume and step over controls.
func (c *Overlay) SetPausedInDebuggerMessage(message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setPausedInDebuggerMessage", Params: paramRequest})
}

// SetSuspended -
// suspended - Whether overlay should be suspended and not consume any resources until resumed.
func (c *Overlay) SetSuspended(suspended bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["suspended"] = suspended
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setSuspended", Params: paramRequest})
}

// SetInspectMode - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
// mode - Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, none
// highlightConfig - A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
func (c *Overlay) SetInspectMode(mode string, highlightConfig *OverlayHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["mode"] = mode
	paramRequest["highlightConfig"] = highlightConfig
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setInspectMode", Params: paramRequest})
}

// HighlightRect - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
// x - X coordinate
// y - Y coordinate
// width - Rectangle width
// height - Rectangle height
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightRect(x int, y int, width int, height int, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 6)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightRect", Params: paramRequest})
}

// HighlightQuad - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
// quad - Quad to highlight
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightQuad(quad []float64, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["quad"] = quad
	paramRequest["color"] = color
	paramRequest["outlineColor"] = outlineColor
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightQuad", Params: paramRequest})
}

// HighlightNode - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// highlightConfig - A descriptor for the highlight appearance.
// nodeId - Identifier of the node to highlight.
// backendNodeId - Identifier of the backend node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
func (c *Overlay) HighlightNode(highlightConfig *OverlayHighlightConfig, nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["highlightConfig"] = highlightConfig
	paramRequest["nodeId"] = nodeId
	paramRequest["backendNodeId"] = backendNodeId
	paramRequest["objectId"] = objectId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightNode", Params: paramRequest})
}

// HighlightFrame - Highlights owner element of the frame with given id.
// frameId - Identifier of the frame to highlight.
// contentColor - The content box highlight fill color (default: transparent).
// contentOutlineColor - The content box highlight outline color (default: transparent).
func (c *Overlay) HighlightFrame(frameId string, contentColor *DOMRGBA, contentOutlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["frameId"] = frameId
	paramRequest["contentColor"] = contentColor
	paramRequest["contentOutlineColor"] = contentOutlineColor
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightFrame", Params: paramRequest})
}

// Hides any highlight.
func (c *Overlay) HideHighlight() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.hideHighlight"})
}

// GetHighlightObjectForTest - For testing.
// nodeId - Id of the node to get highlight object for.
// Returns -  highlight - Highlight data for the node.
func (c *Overlay) GetHighlightObjectForTest(nodeId int) (map[string]interface{}, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.getHighlightObjectForTest", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Highlight map[string]interface{}
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

	return chromeData.Result.Highlight, nil
}
