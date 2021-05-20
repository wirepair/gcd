// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Overlay functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Configuration data for drawing the source order of an elements children.
type OverlaySourceOrderConfig struct {
	ParentOutlineColor *DOMRGBA `json:"parentOutlineColor"` // the color to outline the givent element in.
	ChildOutlineColor  *DOMRGBA `json:"childOutlineColor"`  // the color to outline the child elements in.
}

// Configuration data for the highlighting of Grid elements.
type OverlayGridHighlightConfig struct {
	ShowGridExtensionLines  bool     `json:"showGridExtensionLines,omitempty"`  // Whether the extension lines from grid cells to the rulers should be shown (default: false).
	ShowPositiveLineNumbers bool     `json:"showPositiveLineNumbers,omitempty"` // Show Positive line number labels (default: false).
	ShowNegativeLineNumbers bool     `json:"showNegativeLineNumbers,omitempty"` // Show Negative line number labels (default: false).
	ShowAreaNames           bool     `json:"showAreaNames,omitempty"`           // Show area name labels (default: false).
	ShowLineNames           bool     `json:"showLineNames,omitempty"`           // Show line name labels (default: false).
	ShowTrackSizes          bool     `json:"showTrackSizes,omitempty"`          // Show track size labels (default: false).
	GridBorderColor         *DOMRGBA `json:"gridBorderColor,omitempty"`         // The grid container border highlight color (default: transparent).
	CellBorderColor         *DOMRGBA `json:"cellBorderColor,omitempty"`         // The cell border color (default: transparent). Deprecated, please use rowLineColor and columnLineColor instead.
	RowLineColor            *DOMRGBA `json:"rowLineColor,omitempty"`            // The row line color (default: transparent).
	ColumnLineColor         *DOMRGBA `json:"columnLineColor,omitempty"`         // The column line color (default: transparent).
	GridBorderDash          bool     `json:"gridBorderDash,omitempty"`          // Whether the grid border is dashed (default: false).
	CellBorderDash          bool     `json:"cellBorderDash,omitempty"`          // Whether the cell border is dashed (default: false). Deprecated, please us rowLineDash and columnLineDash instead.
	RowLineDash             bool     `json:"rowLineDash,omitempty"`             // Whether row lines are dashed (default: false).
	ColumnLineDash          bool     `json:"columnLineDash,omitempty"`          // Whether column lines are dashed (default: false).
	RowGapColor             *DOMRGBA `json:"rowGapColor,omitempty"`             // The row gap highlight fill color (default: transparent).
	RowHatchColor           *DOMRGBA `json:"rowHatchColor,omitempty"`           // The row gap hatching fill color (default: transparent).
	ColumnGapColor          *DOMRGBA `json:"columnGapColor,omitempty"`          // The column gap highlight fill color (default: transparent).
	ColumnHatchColor        *DOMRGBA `json:"columnHatchColor,omitempty"`        // The column gap hatching fill color (default: transparent).
	AreaBorderColor         *DOMRGBA `json:"areaBorderColor,omitempty"`         // The named grid areas border color (Default: transparent).
	GridBackgroundColor     *DOMRGBA `json:"gridBackgroundColor,omitempty"`     // The grid container background color (Default: transparent).
}

// Configuration data for the highlighting of Flex container elements.
type OverlayFlexContainerHighlightConfig struct {
	ContainerBorder       *OverlayLineStyle `json:"containerBorder,omitempty"`       // The style of the container border
	LineSeparator         *OverlayLineStyle `json:"lineSeparator,omitempty"`         // The style of the separator between lines
	ItemSeparator         *OverlayLineStyle `json:"itemSeparator,omitempty"`         // The style of the separator between items
	MainDistributedSpace  *OverlayBoxStyle  `json:"mainDistributedSpace,omitempty"`  // Style of content-distribution space on the main axis (justify-content).
	CrossDistributedSpace *OverlayBoxStyle  `json:"crossDistributedSpace,omitempty"` // Style of content-distribution space on the cross axis (align-content).
	RowGapSpace           *OverlayBoxStyle  `json:"rowGapSpace,omitempty"`           // Style of empty space caused by row gaps (gap/row-gap).
	ColumnGapSpace        *OverlayBoxStyle  `json:"columnGapSpace,omitempty"`        // Style of empty space caused by columns gaps (gap/column-gap).
	CrossAlignment        *OverlayLineStyle `json:"crossAlignment,omitempty"`        // Style of the self-alignment line (align-items).
}

// Configuration data for the highlighting of Flex item elements.
type OverlayFlexItemHighlightConfig struct {
	BaseSizeBox      *OverlayBoxStyle  `json:"baseSizeBox,omitempty"`      // Style of the box representing the item's base size
	BaseSizeBorder   *OverlayLineStyle `json:"baseSizeBorder,omitempty"`   // Style of the border around the box representing the item's base size
	FlexibilityArrow *OverlayLineStyle `json:"flexibilityArrow,omitempty"` // Style of the arrow representing if the item grew or shrank
}

// Style information for drawing a line.
type OverlayLineStyle struct {
	Color   *DOMRGBA `json:"color,omitempty"`   // The color of the line (default: transparent)
	Pattern string   `json:"pattern,omitempty"` // The line pattern (default: solid)
}

// Style information for drawing a box.
type OverlayBoxStyle struct {
	FillColor  *DOMRGBA `json:"fillColor,omitempty"`  // The background color for the box (default: transparent)
	HatchColor *DOMRGBA `json:"hatchColor,omitempty"` // The hatching color for the box (default: transparent)
}

// Configuration data for the highlighting of page elements.
type OverlayHighlightConfig struct {
	ShowInfo                     bool                                 `json:"showInfo,omitempty"`                     // Whether the node info tooltip should be shown (default: false).
	ShowStyles                   bool                                 `json:"showStyles,omitempty"`                   // Whether the node styles in the tooltip (default: false).
	ShowRulers                   bool                                 `json:"showRulers,omitempty"`                   // Whether the rulers should be shown (default: false).
	ShowAccessibilityInfo        bool                                 `json:"showAccessibilityInfo,omitempty"`        // Whether the a11y info should be shown (default: true).
	ShowExtensionLines           bool                                 `json:"showExtensionLines,omitempty"`           // Whether the extension lines from node to the rulers should be shown (default: false).
	ContentColor                 *DOMRGBA                             `json:"contentColor,omitempty"`                 // The content box highlight fill color (default: transparent).
	PaddingColor                 *DOMRGBA                             `json:"paddingColor,omitempty"`                 // The padding highlight fill color (default: transparent).
	BorderColor                  *DOMRGBA                             `json:"borderColor,omitempty"`                  // The border highlight fill color (default: transparent).
	MarginColor                  *DOMRGBA                             `json:"marginColor,omitempty"`                  // The margin highlight fill color (default: transparent).
	EventTargetColor             *DOMRGBA                             `json:"eventTargetColor,omitempty"`             // The event target element highlight fill color (default: transparent).
	ShapeColor                   *DOMRGBA                             `json:"shapeColor,omitempty"`                   // The shape outside fill color (default: transparent).
	ShapeMarginColor             *DOMRGBA                             `json:"shapeMarginColor,omitempty"`             // The shape margin fill color (default: transparent).
	CssGridColor                 *DOMRGBA                             `json:"cssGridColor,omitempty"`                 // The grid layout color (default: transparent).
	ColorFormat                  string                               `json:"colorFormat,omitempty"`                  // The color format used to format color styles (default: hex). enum values: rgb, hsl, hex
	GridHighlightConfig          *OverlayGridHighlightConfig          `json:"gridHighlightConfig,omitempty"`          // The grid layout highlight configuration (default: all transparent).
	FlexContainerHighlightConfig *OverlayFlexContainerHighlightConfig `json:"flexContainerHighlightConfig,omitempty"` // The flex container highlight configuration (default: all transparent).
	FlexItemHighlightConfig      *OverlayFlexItemHighlightConfig      `json:"flexItemHighlightConfig,omitempty"`      // The flex item highlight configuration (default: all transparent).
	ContrastAlgorithm            string                               `json:"contrastAlgorithm,omitempty"`            // The contrast algorithm to use for the contrast ratio (default: aa). enum values: aa, aaa, apca
}

// Configurations for Persistent Grid Highlight
type OverlayGridNodeHighlightConfig struct {
	GridHighlightConfig *OverlayGridHighlightConfig `json:"gridHighlightConfig"` // A descriptor for the highlight appearance.
	NodeId              int                         `json:"nodeId"`              // Identifier of the node to highlight.
}

// No Description.
type OverlayFlexNodeHighlightConfig struct {
	FlexContainerHighlightConfig *OverlayFlexContainerHighlightConfig `json:"flexContainerHighlightConfig"` // A descriptor for the highlight appearance of flex containers.
	NodeId                       int                                  `json:"nodeId"`                       // Identifier of the node to highlight.
}

// No Description.
type OverlayScrollSnapContainerHighlightConfig struct {
	SnapportBorder     *OverlayLineStyle `json:"snapportBorder,omitempty"`     // The style of the snapport border (default: transparent)
	SnapAreaBorder     *OverlayLineStyle `json:"snapAreaBorder,omitempty"`     // The style of the snap area border (default: transparent)
	ScrollMarginColor  *DOMRGBA          `json:"scrollMarginColor,omitempty"`  // The margin highlight fill color (default: transparent).
	ScrollPaddingColor *DOMRGBA          `json:"scrollPaddingColor,omitempty"` // The padding highlight fill color (default: transparent).
}

// No Description.
type OverlayScrollSnapHighlightConfig struct {
	ScrollSnapContainerHighlightConfig *OverlayScrollSnapContainerHighlightConfig `json:"scrollSnapContainerHighlightConfig"` // A descriptor for the highlight appearance of scroll snap containers.
	NodeId                             int                                        `json:"nodeId"`                             // Identifier of the node to highlight.
}

// Configuration for dual screen hinge
type OverlayHingeConfig struct {
	Rect         *DOMRect `json:"rect"`                   // A rectangle represent hinge
	ContentColor *DOMRGBA `json:"contentColor,omitempty"` // The content box highlight fill color (default: a dark color).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"` // The content box highlight outline color (default: transparent).
}

// Fired when the node should be inspected. This happens after call to `setInspectMode` or when user manually inspects an element.
type OverlayInspectNodeRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		BackendNodeId int `json:"backendNodeId"` // Id of the node to inspect.
	} `json:"Params,omitempty"`
}

// Fired when the node should be highlighted. This happens after call to `setInspectMode`.
type OverlayNodeHighlightRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId int `json:"nodeId"` //
	} `json:"Params,omitempty"`
}

// Fired when user asks to capture screenshot of some area on the page.
type OverlayScreenshotRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		Viewport *PageViewport `json:"viewport"` // Viewport to capture, in device independent pixels (dip).
	} `json:"Params,omitempty"`
}

type Overlay struct {
	target gcdmessage.ChromeTargeter
}

func NewOverlay(target gcdmessage.ChromeTargeter) *Overlay {
	c := &Overlay{target: target}
	return c
}

// Disables domain notifications.
func (c *Overlay) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.disable"})
}

// Enables domain notifications.
func (c *Overlay) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.enable"})
}

type OverlayGetHighlightObjectForTestParams struct {
	// Id of the node to get highlight object for.
	NodeId int `json:"nodeId"`
	// Whether to include distance info.
	IncludeDistance bool `json:"includeDistance,omitempty"`
	// Whether to include style info.
	IncludeStyle bool `json:"includeStyle,omitempty"`
	// The color format to get config with (default: hex). enum values: rgb, hsl, hex
	ColorFormat string `json:"colorFormat,omitempty"`
	// Whether to show accessibility info (default: true).
	ShowAccessibilityInfo bool `json:"showAccessibilityInfo,omitempty"`
}

// GetHighlightObjectForTestWithParams - For testing.
// Returns -  highlight - Highlight data for the node.
func (c *Overlay) GetHighlightObjectForTestWithParams(ctx context.Context, v *OverlayGetHighlightObjectForTestParams) (map[string]interface{}, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.getHighlightObjectForTest", Params: v})
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

// GetHighlightObjectForTest - For testing.
// nodeId - Id of the node to get highlight object for.
// includeDistance - Whether to include distance info.
// includeStyle - Whether to include style info.
// colorFormat - The color format to get config with (default: hex). enum values: rgb, hsl, hex
// showAccessibilityInfo - Whether to show accessibility info (default: true).
// Returns -  highlight - Highlight data for the node.
func (c *Overlay) GetHighlightObjectForTest(ctx context.Context, nodeId int, includeDistance bool, includeStyle bool, colorFormat string, showAccessibilityInfo bool) (map[string]interface{}, error) {
	var v OverlayGetHighlightObjectForTestParams
	v.NodeId = nodeId
	v.IncludeDistance = includeDistance
	v.IncludeStyle = includeStyle
	v.ColorFormat = colorFormat
	v.ShowAccessibilityInfo = showAccessibilityInfo
	return c.GetHighlightObjectForTestWithParams(ctx, &v)
}

type OverlayGetGridHighlightObjectsForTestParams struct {
	// Ids of the node to get highlight object for.
	NodeIds []int `json:"nodeIds"`
}

// GetGridHighlightObjectsForTestWithParams - For Persistent Grid testing.
// Returns -  highlights - Grid Highlight data for the node ids provided.
func (c *Overlay) GetGridHighlightObjectsForTestWithParams(ctx context.Context, v *OverlayGetGridHighlightObjectsForTestParams) (map[string]interface{}, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.getGridHighlightObjectsForTest", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Highlights map[string]interface{}
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

	return chromeData.Result.Highlights, nil
}

// GetGridHighlightObjectsForTest - For Persistent Grid testing.
// nodeIds - Ids of the node to get highlight object for.
// Returns -  highlights - Grid Highlight data for the node ids provided.
func (c *Overlay) GetGridHighlightObjectsForTest(ctx context.Context, nodeIds []int) (map[string]interface{}, error) {
	var v OverlayGetGridHighlightObjectsForTestParams
	v.NodeIds = nodeIds
	return c.GetGridHighlightObjectsForTestWithParams(ctx, &v)
}

type OverlayGetSourceOrderHighlightObjectForTestParams struct {
	// Id of the node to highlight.
	NodeId int `json:"nodeId"`
}

// GetSourceOrderHighlightObjectForTestWithParams - For Source Order Viewer testing.
// Returns -  highlight - Source order highlight data for the node id provided.
func (c *Overlay) GetSourceOrderHighlightObjectForTestWithParams(ctx context.Context, v *OverlayGetSourceOrderHighlightObjectForTestParams) (map[string]interface{}, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.getSourceOrderHighlightObjectForTest", Params: v})
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

// GetSourceOrderHighlightObjectForTest - For Source Order Viewer testing.
// nodeId - Id of the node to highlight.
// Returns -  highlight - Source order highlight data for the node id provided.
func (c *Overlay) GetSourceOrderHighlightObjectForTest(ctx context.Context, nodeId int) (map[string]interface{}, error) {
	var v OverlayGetSourceOrderHighlightObjectForTestParams
	v.NodeId = nodeId
	return c.GetSourceOrderHighlightObjectForTestWithParams(ctx, &v)
}

// Hides any highlight.
func (c *Overlay) HideHighlight(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.hideHighlight"})
}

type OverlayHighlightFrameParams struct {
	// Identifier of the frame to highlight.
	FrameId string `json:"frameId"`
	// The content box highlight fill color (default: transparent).
	ContentColor *DOMRGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor *DOMRGBA `json:"contentOutlineColor,omitempty"`
}

// HighlightFrameWithParams - Highlights owner element of the frame with given id.
func (c *Overlay) HighlightFrameWithParams(ctx context.Context, v *OverlayHighlightFrameParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightFrame", Params: v})
}

// HighlightFrame - Highlights owner element of the frame with given id.
// frameId - Identifier of the frame to highlight.
// contentColor - The content box highlight fill color (default: transparent).
// contentOutlineColor - The content box highlight outline color (default: transparent).
func (c *Overlay) HighlightFrame(ctx context.Context, frameId string, contentColor *DOMRGBA, contentOutlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightFrameParams
	v.FrameId = frameId
	v.ContentColor = contentColor
	v.ContentOutlineColor = contentOutlineColor
	return c.HighlightFrameWithParams(ctx, &v)
}

type OverlayHighlightNodeParams struct {
	// A descriptor for the highlight appearance.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig"`
	// Identifier of the node to highlight.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectId string `json:"objectId,omitempty"`
	// Selectors to highlight relevant nodes.
	Selector string `json:"selector,omitempty"`
}

// HighlightNodeWithParams - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (c *Overlay) HighlightNodeWithParams(ctx context.Context, v *OverlayHighlightNodeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightNode", Params: v})
}

// HighlightNode - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// highlightConfig - A descriptor for the highlight appearance.
// nodeId - Identifier of the node to highlight.
// backendNodeId - Identifier of the backend node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
// selector - Selectors to highlight relevant nodes.
func (c *Overlay) HighlightNode(ctx context.Context, highlightConfig *OverlayHighlightConfig, nodeId int, backendNodeId int, objectId string, selector string) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightNodeParams
	v.HighlightConfig = highlightConfig
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	v.Selector = selector
	return c.HighlightNodeWithParams(ctx, &v)
}

type OverlayHighlightQuadParams struct {
	// Quad to highlight
	Quad []float64 `json:"quad"`
	// The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// HighlightQuadWithParams - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
func (c *Overlay) HighlightQuadWithParams(ctx context.Context, v *OverlayHighlightQuadParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightQuad", Params: v})
}

// HighlightQuad - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
// quad - Quad to highlight
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightQuad(ctx context.Context, quad []float64, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightQuadParams
	v.Quad = quad
	v.Color = color
	v.OutlineColor = outlineColor
	return c.HighlightQuadWithParams(ctx, &v)
}

type OverlayHighlightRectParams struct {
	// X coordinate
	X int `json:"x"`
	// Y coordinate
	Y int `json:"y"`
	// Rectangle width
	Width int `json:"width"`
	// Rectangle height
	Height int `json:"height"`
	// The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// HighlightRectWithParams - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (c *Overlay) HighlightRectWithParams(ctx context.Context, v *OverlayHighlightRectParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightRect", Params: v})
}

// HighlightRect - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
// x - X coordinate
// y - Y coordinate
// width - Rectangle width
// height - Rectangle height
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightRect(ctx context.Context, x int, y int, width int, height int, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightRectParams
	v.X = x
	v.Y = y
	v.Width = width
	v.Height = height
	v.Color = color
	v.OutlineColor = outlineColor
	return c.HighlightRectWithParams(ctx, &v)
}

type OverlayHighlightSourceOrderParams struct {
	// A descriptor for the appearance of the overlay drawing.
	SourceOrderConfig *OverlaySourceOrderConfig `json:"sourceOrderConfig"`
	// Identifier of the node to highlight.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectId string `json:"objectId,omitempty"`
}

// HighlightSourceOrderWithParams - Highlights the source order of the children of the DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (c *Overlay) HighlightSourceOrderWithParams(ctx context.Context, v *OverlayHighlightSourceOrderParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightSourceOrder", Params: v})
}

// HighlightSourceOrder - Highlights the source order of the children of the DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// sourceOrderConfig - A descriptor for the appearance of the overlay drawing.
// nodeId - Identifier of the node to highlight.
// backendNodeId - Identifier of the backend node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
func (c *Overlay) HighlightSourceOrder(ctx context.Context, sourceOrderConfig *OverlaySourceOrderConfig, nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightSourceOrderParams
	v.SourceOrderConfig = sourceOrderConfig
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.HighlightSourceOrderWithParams(ctx, &v)
}

type OverlaySetInspectModeParams struct {
	// Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, captureAreaScreenshot, showDistances, none
	Mode string `json:"mode"`
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled == false`.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig,omitempty"`
}

// SetInspectModeWithParams - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
func (c *Overlay) SetInspectModeWithParams(ctx context.Context, v *OverlaySetInspectModeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setInspectMode", Params: v})
}

// SetInspectMode - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
// mode - Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, captureAreaScreenshot, showDistances, none
// highlightConfig - A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled == false`.
func (c *Overlay) SetInspectMode(ctx context.Context, mode string, highlightConfig *OverlayHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetInspectModeParams
	v.Mode = mode
	v.HighlightConfig = highlightConfig
	return c.SetInspectModeWithParams(ctx, &v)
}

type OverlaySetShowAdHighlightsParams struct {
	// True for showing ad highlights
	Show bool `json:"show"`
}

// SetShowAdHighlightsWithParams - Highlights owner element of all frames detected to be ads.
func (c *Overlay) SetShowAdHighlightsWithParams(ctx context.Context, v *OverlaySetShowAdHighlightsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowAdHighlights", Params: v})
}

// SetShowAdHighlights - Highlights owner element of all frames detected to be ads.
// show - True for showing ad highlights
func (c *Overlay) SetShowAdHighlights(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowAdHighlightsParams
	v.Show = show
	return c.SetShowAdHighlightsWithParams(ctx, &v)
}

type OverlaySetPausedInDebuggerMessageParams struct {
	// The message to display, also triggers resume and step over controls.
	Message string `json:"message,omitempty"`
}

// SetPausedInDebuggerMessageWithParams -
func (c *Overlay) SetPausedInDebuggerMessageWithParams(ctx context.Context, v *OverlaySetPausedInDebuggerMessageParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setPausedInDebuggerMessage", Params: v})
}

// SetPausedInDebuggerMessage -
// message - The message to display, also triggers resume and step over controls.
func (c *Overlay) SetPausedInDebuggerMessage(ctx context.Context, message string) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetPausedInDebuggerMessageParams
	v.Message = message
	return c.SetPausedInDebuggerMessageWithParams(ctx, &v)
}

type OverlaySetShowDebugBordersParams struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// SetShowDebugBordersWithParams - Requests that backend shows debug borders on layers
func (c *Overlay) SetShowDebugBordersWithParams(ctx context.Context, v *OverlaySetShowDebugBordersParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowDebugBorders", Params: v})
}

// SetShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *Overlay) SetShowDebugBorders(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowDebugBordersParams
	v.Show = show
	return c.SetShowDebugBordersWithParams(ctx, &v)
}

type OverlaySetShowFPSCounterParams struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// SetShowFPSCounterWithParams - Requests that backend shows the FPS counter
func (c *Overlay) SetShowFPSCounterWithParams(ctx context.Context, v *OverlaySetShowFPSCounterParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowFPSCounter", Params: v})
}

// SetShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *Overlay) SetShowFPSCounter(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowFPSCounterParams
	v.Show = show
	return c.SetShowFPSCounterWithParams(ctx, &v)
}

type OverlaySetShowGridOverlaysParams struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	GridNodeHighlightConfigs []*OverlayGridNodeHighlightConfig `json:"gridNodeHighlightConfigs"`
}

// SetShowGridOverlaysWithParams - Highlight multiple elements with the CSS Grid overlay.
func (c *Overlay) SetShowGridOverlaysWithParams(ctx context.Context, v *OverlaySetShowGridOverlaysParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowGridOverlays", Params: v})
}

// SetShowGridOverlays - Highlight multiple elements with the CSS Grid overlay.
// gridNodeHighlightConfigs - An array of node identifiers and descriptors for the highlight appearance.
func (c *Overlay) SetShowGridOverlays(ctx context.Context, gridNodeHighlightConfigs []*OverlayGridNodeHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowGridOverlaysParams
	v.GridNodeHighlightConfigs = gridNodeHighlightConfigs
	return c.SetShowGridOverlaysWithParams(ctx, &v)
}

type OverlaySetShowFlexOverlaysParams struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	FlexNodeHighlightConfigs []*OverlayFlexNodeHighlightConfig `json:"flexNodeHighlightConfigs"`
}

// SetShowFlexOverlaysWithParams -
func (c *Overlay) SetShowFlexOverlaysWithParams(ctx context.Context, v *OverlaySetShowFlexOverlaysParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowFlexOverlays", Params: v})
}

// SetShowFlexOverlays -
// flexNodeHighlightConfigs - An array of node identifiers and descriptors for the highlight appearance.
func (c *Overlay) SetShowFlexOverlays(ctx context.Context, flexNodeHighlightConfigs []*OverlayFlexNodeHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowFlexOverlaysParams
	v.FlexNodeHighlightConfigs = flexNodeHighlightConfigs
	return c.SetShowFlexOverlaysWithParams(ctx, &v)
}

type OverlaySetShowScrollSnapOverlaysParams struct {
	// An array of node identifiers and descriptors for the highlight appearance.
	ScrollSnapHighlightConfigs []*OverlayScrollSnapHighlightConfig `json:"scrollSnapHighlightConfigs"`
}

// SetShowScrollSnapOverlaysWithParams -
func (c *Overlay) SetShowScrollSnapOverlaysWithParams(ctx context.Context, v *OverlaySetShowScrollSnapOverlaysParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowScrollSnapOverlays", Params: v})
}

// SetShowScrollSnapOverlays -
// scrollSnapHighlightConfigs - An array of node identifiers and descriptors for the highlight appearance.
func (c *Overlay) SetShowScrollSnapOverlays(ctx context.Context, scrollSnapHighlightConfigs []*OverlayScrollSnapHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowScrollSnapOverlaysParams
	v.ScrollSnapHighlightConfigs = scrollSnapHighlightConfigs
	return c.SetShowScrollSnapOverlaysWithParams(ctx, &v)
}

type OverlaySetShowPaintRectsParams struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// SetShowPaintRectsWithParams - Requests that backend shows paint rectangles
func (c *Overlay) SetShowPaintRectsWithParams(ctx context.Context, v *OverlaySetShowPaintRectsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowPaintRects", Params: v})
}

// SetShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *Overlay) SetShowPaintRects(ctx context.Context, result bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowPaintRectsParams
	v.Result = result
	return c.SetShowPaintRectsWithParams(ctx, &v)
}

type OverlaySetShowLayoutShiftRegionsParams struct {
	// True for showing layout shift regions
	Result bool `json:"result"`
}

// SetShowLayoutShiftRegionsWithParams - Requests that backend shows layout shift regions
func (c *Overlay) SetShowLayoutShiftRegionsWithParams(ctx context.Context, v *OverlaySetShowLayoutShiftRegionsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowLayoutShiftRegions", Params: v})
}

// SetShowLayoutShiftRegions - Requests that backend shows layout shift regions
// result - True for showing layout shift regions
func (c *Overlay) SetShowLayoutShiftRegions(ctx context.Context, result bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowLayoutShiftRegionsParams
	v.Result = result
	return c.SetShowLayoutShiftRegionsWithParams(ctx, &v)
}

type OverlaySetShowScrollBottleneckRectsParams struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// SetShowScrollBottleneckRectsWithParams - Requests that backend shows scroll bottleneck rects
func (c *Overlay) SetShowScrollBottleneckRectsWithParams(ctx context.Context, v *OverlaySetShowScrollBottleneckRectsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowScrollBottleneckRects", Params: v})
}

// SetShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *Overlay) SetShowScrollBottleneckRects(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowScrollBottleneckRectsParams
	v.Show = show
	return c.SetShowScrollBottleneckRectsWithParams(ctx, &v)
}

type OverlaySetShowHitTestBordersParams struct {
	// True for showing hit-test borders
	Show bool `json:"show"`
}

// SetShowHitTestBordersWithParams - Requests that backend shows hit-test borders on layers
func (c *Overlay) SetShowHitTestBordersWithParams(ctx context.Context, v *OverlaySetShowHitTestBordersParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowHitTestBorders", Params: v})
}

// SetShowHitTestBorders - Requests that backend shows hit-test borders on layers
// show - True for showing hit-test borders
func (c *Overlay) SetShowHitTestBorders(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowHitTestBordersParams
	v.Show = show
	return c.SetShowHitTestBordersWithParams(ctx, &v)
}

type OverlaySetShowWebVitalsParams struct {
	//
	Show bool `json:"show"`
}

// SetShowWebVitalsWithParams - Request that backend shows an overlay with web vital metrics.
func (c *Overlay) SetShowWebVitalsWithParams(ctx context.Context, v *OverlaySetShowWebVitalsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowWebVitals", Params: v})
}

// SetShowWebVitals - Request that backend shows an overlay with web vital metrics.
// show -
func (c *Overlay) SetShowWebVitals(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowWebVitalsParams
	v.Show = show
	return c.SetShowWebVitalsWithParams(ctx, &v)
}

type OverlaySetShowViewportSizeOnResizeParams struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// SetShowViewportSizeOnResizeWithParams - Paints viewport size upon main frame resize.
func (c *Overlay) SetShowViewportSizeOnResizeWithParams(ctx context.Context, v *OverlaySetShowViewportSizeOnResizeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowViewportSizeOnResize", Params: v})
}

// SetShowViewportSizeOnResize - Paints viewport size upon main frame resize.
// show - Whether to paint size or not.
func (c *Overlay) SetShowViewportSizeOnResize(ctx context.Context, show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowViewportSizeOnResizeParams
	v.Show = show
	return c.SetShowViewportSizeOnResizeWithParams(ctx, &v)
}

type OverlaySetShowHingeParams struct {
	// hinge data, null means hideHinge
	HingeConfig *OverlayHingeConfig `json:"hingeConfig,omitempty"`
}

// SetShowHingeWithParams - Add a dual screen device hinge
func (c *Overlay) SetShowHingeWithParams(ctx context.Context, v *OverlaySetShowHingeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowHinge", Params: v})
}

// SetShowHinge - Add a dual screen device hinge
// hingeConfig - hinge data, null means hideHinge
func (c *Overlay) SetShowHinge(ctx context.Context, hingeConfig *OverlayHingeConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowHingeParams
	v.HingeConfig = hingeConfig
	return c.SetShowHingeWithParams(ctx, &v)
}
