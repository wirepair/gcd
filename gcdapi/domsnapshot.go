// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMSnapshot functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// A Node in the DOM tree.
type DOMSnapshotDOMNode struct {
	NodeType             int                         `json:"nodeType"`                       // `Node`'s nodeType.
	NodeName             string                      `json:"nodeName"`                       // `Node`'s nodeName.
	NodeValue            string                      `json:"nodeValue"`                      // `Node`'s nodeValue.
	TextValue            string                      `json:"textValue,omitempty"`            // Only set for textarea elements, contains the text value.
	InputValue           string                      `json:"inputValue,omitempty"`           // Only set for input elements, contains the input's associated text value.
	InputChecked         bool                        `json:"inputChecked,omitempty"`         // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected       bool                        `json:"optionSelected,omitempty"`       // Only set for option elements, indicates if the element has been selected
	BackendNodeId        int                         `json:"backendNodeId"`                  // `Node`'s id, corresponds to DOM.Node.backendNodeId.
	ChildNodeIndexes     []int                       `json:"childNodeIndexes,omitempty"`     // The indexes of the node's child nodes in the `domNodes` array returned by `getSnapshot`, if any.
	Attributes           []*DOMSnapshotNameValue     `json:"attributes,omitempty"`           // Attributes of an `Element` node.
	PseudoElementIndexes []int                       `json:"pseudoElementIndexes,omitempty"` // Indexes of pseudo elements associated with this node in the `domNodes` array returned by `getSnapshot`, if any.
	LayoutNodeIndex      int                         `json:"layoutNodeIndex,omitempty"`      // The index of the node's related layout tree node in the `layoutTreeNodes` array returned by `getSnapshot`, if any.
	DocumentURL          string                      `json:"documentURL,omitempty"`          // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL              string                      `json:"baseURL,omitempty"`              // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	ContentLanguage      string                      `json:"contentLanguage,omitempty"`      // Only set for documents, contains the document's content language.
	DocumentEncoding     string                      `json:"documentEncoding,omitempty"`     // Only set for documents, contains the document's character set encoding.
	PublicId             string                      `json:"publicId,omitempty"`             // `DocumentType` node's publicId.
	SystemId             string                      `json:"systemId,omitempty"`             // `DocumentType` node's systemId.
	FrameId              string                      `json:"frameId,omitempty"`              // Frame ID for frame owner elements and also for the document node.
	ContentDocumentIndex int                         `json:"contentDocumentIndex,omitempty"` // The index of a frame owner element's content document in the `domNodes` array returned by `getSnapshot`, if any.
	PseudoType           string                      `json:"pseudoType,omitempty"`           // Type of a pseudo element node. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	ShadowRootType       string                      `json:"shadowRootType,omitempty"`       // Shadow root type. enum values: user-agent, open, closed
	IsClickable          bool                        `json:"isClickable,omitempty"`          // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	EventListeners       []*DOMDebuggerEventListener `json:"eventListeners,omitempty"`       // Details of the node's event listeners, if any.
	CurrentSourceURL     string                      `json:"currentSourceURL,omitempty"`     // The selected url for nodes with a srcset attribute.
	OriginURL            string                      `json:"originURL,omitempty"`            // The url of the script (if any) that generates this node.
	ScrollOffsetX        float64                     `json:"scrollOffsetX,omitempty"`        // Scroll offsets, set when this node is a Document.
	ScrollOffsetY        float64                     `json:"scrollOffsetY,omitempty"`        //
}

// Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type DOMSnapshotInlineTextBox struct {
	BoundingBox         *DOMRect `json:"boundingBox"`         // The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	StartCharacterIndex int      `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters       int      `json:"numCharacters"`       // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

// Details of an element in the DOM tree with a LayoutObject.
type DOMSnapshotLayoutTreeNode struct {
	DomNodeIndex      int                         `json:"domNodeIndex"`                // The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	BoundingBox       *DOMRect                    `json:"boundingBox"`                 // The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	LayoutText        string                      `json:"layoutText,omitempty"`        // Contents of the LayoutText, if any.
	InlineTextNodes   []*DOMSnapshotInlineTextBox `json:"inlineTextNodes,omitempty"`   // The post-layout inline text nodes, if any.
	StyleIndex        int                         `json:"styleIndex,omitempty"`        // Index into the `computedStyles` array returned by `getSnapshot`.
	PaintOrder        int                         `json:"paintOrder,omitempty"`        // Global paint order index, which is determined by the stacking order of the nodes. Nodes that are painted together will have the same index. Only provided if includePaintOrder in getSnapshot was true.
	IsStackingContext bool                        `json:"isStackingContext,omitempty"` // Set to true to indicate the element begins a new stacking context.
}

// A subset of the full ComputedStyle as defined by the request whitelist.
type DOMSnapshotComputedStyle struct {
	Properties []*DOMSnapshotNameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// A name/value pair.
type DOMSnapshotNameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}

// Data that is only present on rare nodes.
type DOMSnapshotRareStringData struct {
	Index []int `json:"index"` //
	Value []int `json:"value"` //
}

// No Description.
type DOMSnapshotRareBooleanData struct {
	Index []int `json:"index"` //
}

// No Description.
type DOMSnapshotRareIntegerData struct {
	Index []int `json:"index"` //
	Value []int `json:"value"` //
}

// Document snapshot.
type DOMSnapshotDocumentSnapshot struct {
	DocumentURL     int                            `json:"documentURL"`             // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL         int                            `json:"baseURL"`                 // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	ContentLanguage int                            `json:"contentLanguage"`         // Contains the document's content language.
	EncodingName    int                            `json:"encodingName"`            // Contains the document's character set encoding.
	PublicId        int                            `json:"publicId"`                // `DocumentType` node's publicId.
	SystemId        int                            `json:"systemId"`                // `DocumentType` node's systemId.
	FrameId         int                            `json:"frameId"`                 // Frame ID for frame owner elements and also for the document node.
	Nodes           *DOMSnapshotNodeTreeSnapshot   `json:"nodes"`                   // A table with dom nodes.
	Layout          *DOMSnapshotLayoutTreeSnapshot `json:"layout"`                  // The nodes in the layout tree.
	TextBoxes       *DOMSnapshotTextBoxSnapshot    `json:"textBoxes"`               // The post-layout inline text nodes.
	ScrollOffsetX   float64                        `json:"scrollOffsetX,omitempty"` // Horizontal scroll offset.
	ScrollOffsetY   float64                        `json:"scrollOffsetY,omitempty"` // Vertical scroll offset.
}

// Table containing nodes.
type DOMSnapshotNodeTreeSnapshot struct {
	ParentIndex          []int                       `json:"parentIndex,omitempty"`          // Parent node index.
	NodeType             []int                       `json:"nodeType,omitempty"`             // `Node`'s nodeType.
	NodeName             []int                       `json:"nodeName,omitempty"`             // `Node`'s nodeName.
	NodeValue            []int                       `json:"nodeValue,omitempty"`            // `Node`'s nodeValue.
	BackendNodeId        []int                       `json:"backendNodeId,omitempty"`        // `Node`'s id, corresponds to DOM.Node.backendNodeId.
	Attributes           []int                       `json:"attributes,omitempty"`           // Attributes of an `Element` node. Flatten name, value pairs.
	TextValue            *DOMSnapshotRareStringData  `json:"textValue,omitempty"`            // Only set for textarea elements, contains the text value.
	InputValue           *DOMSnapshotRareStringData  `json:"inputValue,omitempty"`           // Only set for input elements, contains the input's associated text value.
	InputChecked         *DOMSnapshotRareBooleanData `json:"inputChecked,omitempty"`         // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected       *DOMSnapshotRareBooleanData `json:"optionSelected,omitempty"`       // Only set for option elements, indicates if the element has been selected
	ContentDocumentIndex *DOMSnapshotRareIntegerData `json:"contentDocumentIndex,omitempty"` // The index of the document in the list of the snapshot documents.
	PseudoType           *DOMSnapshotRareStringData  `json:"pseudoType,omitempty"`           // Type of a pseudo element node.
	IsClickable          *DOMSnapshotRareBooleanData `json:"isClickable,omitempty"`          // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	CurrentSourceURL     *DOMSnapshotRareStringData  `json:"currentSourceURL,omitempty"`     // The selected url for nodes with a srcset attribute.
	OriginURL            *DOMSnapshotRareStringData  `json:"originURL,omitempty"`            // The url of the script (if any) that generates this node.
}

// Table of details of an element in the DOM tree with a LayoutObject.
type DOMSnapshotLayoutTreeSnapshot struct {
	NodeIndex        []int                       `json:"nodeIndex"`             // Index of the corresponding node in the `NodeTreeSnapshot` array returned by `captureSnapshot`.
	Styles           []int                       `json:"styles"`                // Array of indexes specifying computed style strings, filtered according to the `computedStyles` parameter passed to `captureSnapshot`.
	Bounds           []float64                   `json:"bounds"`                // The absolute position bounding box.
	Text             []int                       `json:"text"`                  // Contents of the LayoutText, if any.
	StackingContexts *DOMSnapshotRareBooleanData `json:"stackingContexts"`      // Stacking context information.
	OffsetRects      []float64                   `json:"offsetRects,omitempty"` // The offset rect of nodes. Only available when includeDOMRects is set to true
	ScrollRects      []float64                   `json:"scrollRects,omitempty"` // The scroll rect of nodes. Only available when includeDOMRects is set to true
	ClientRects      []float64                   `json:"clientRects,omitempty"` // The client rect of nodes. Only available when includeDOMRects is set to true
}

// Table of details of the post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type DOMSnapshotTextBoxSnapshot struct {
	LayoutIndex []int     `json:"layoutIndex"` // Index of the layout tree node that owns this box collection.
	Bounds      []float64 `json:"bounds"`      // The absolute position bounding box.
	Start       []int     `json:"start"`       // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	Length      []int     `json:"length"`      // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

type DOMSnapshot struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMSnapshot(target gcdmessage.ChromeTargeter) *DOMSnapshot {
	c := &DOMSnapshot{target: target}
	return c
}

// Disables DOM snapshot agent for the given page.
func (c *DOMSnapshot) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMSnapshot.disable"})
}

// Enables DOM snapshot agent for the given page.
func (c *DOMSnapshot) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMSnapshot.enable"})
}

type DOMSnapshotGetSnapshotParams struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
	// Whether or not to retrieve details of DOM listeners (default false).
	IncludeEventListeners bool `json:"includeEventListeners,omitempty"`
	// Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludePaintOrder bool `json:"includePaintOrder,omitempty"`
	// Whether to include UA shadow tree in the snapshot (default false).
	IncludeUserAgentShadowTree bool `json:"includeUserAgentShadowTree,omitempty"`
}

// GetSnapshotWithParams - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// Returns -  domNodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. layoutTreeNodes - The nodes in the layout tree. computedStyles - Whitelisted ComputedStyle properties for each node in the layout tree.
func (c *DOMSnapshot) GetSnapshotWithParams(v *DOMSnapshotGetSnapshotParams) ([]*DOMSnapshotDOMNode, []*DOMSnapshotLayoutTreeNode, []*DOMSnapshotComputedStyle, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMSnapshot.getSnapshot", Params: v})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			DomNodes        []*DOMSnapshotDOMNode
			LayoutTreeNodes []*DOMSnapshotLayoutTreeNode
			ComputedStyles  []*DOMSnapshotComputedStyle
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.DomNodes, chromeData.Result.LayoutTreeNodes, chromeData.Result.ComputedStyles, nil
}

// GetSnapshot - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// computedStyleWhitelist - Whitelist of computed styles to return.
// includeEventListeners - Whether or not to retrieve details of DOM listeners (default false).
// includePaintOrder - Whether to determine and include the paint order index of LayoutTreeNodes (default false).
// includeUserAgentShadowTree - Whether to include UA shadow tree in the snapshot (default false).
// Returns -  domNodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. layoutTreeNodes - The nodes in the layout tree. computedStyles - Whitelisted ComputedStyle properties for each node in the layout tree.
func (c *DOMSnapshot) GetSnapshot(computedStyleWhitelist []string, includeEventListeners bool, includePaintOrder bool, includeUserAgentShadowTree bool) ([]*DOMSnapshotDOMNode, []*DOMSnapshotLayoutTreeNode, []*DOMSnapshotComputedStyle, error) {
	var v DOMSnapshotGetSnapshotParams
	v.ComputedStyleWhitelist = computedStyleWhitelist
	v.IncludeEventListeners = includeEventListeners
	v.IncludePaintOrder = includePaintOrder
	v.IncludeUserAgentShadowTree = includeUserAgentShadowTree
	return c.GetSnapshotWithParams(&v)
}

type DOMSnapshotCaptureSnapshotParams struct {
	// Whitelist of computed styles to return.
	ComputedStyles []string `json:"computedStyles"`
	// Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
	IncludeDOMRects bool `json:"includeDOMRects,omitempty"`
}

// CaptureSnapshotWithParams - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// Returns -  documents - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. strings - Shared string table that all string properties refer to with indexes.
func (c *DOMSnapshot) CaptureSnapshotWithParams(v *DOMSnapshotCaptureSnapshotParams) ([]*DOMSnapshotDocumentSnapshot, []string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMSnapshot.captureSnapshot", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Documents []*DOMSnapshotDocumentSnapshot
			Strings   []string
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Documents, chromeData.Result.Strings, nil
}

// CaptureSnapshot - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// computedStyles - Whitelist of computed styles to return.
// includeDOMRects - Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
// Returns -  documents - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. strings - Shared string table that all string properties refer to with indexes.
func (c *DOMSnapshot) CaptureSnapshot(computedStyles []string, includeDOMRects bool) ([]*DOMSnapshotDocumentSnapshot, []string, error) {
	var v DOMSnapshotCaptureSnapshotParams
	v.ComputedStyles = computedStyles
	v.IncludeDOMRects = includeDOMRects
	return c.CaptureSnapshotWithParams(&v)
}
