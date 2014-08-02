// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOM types.
// API Version: 1.1
package types

 
// Unique DOM node identifier.
type ChromeDOMNodeId int 
 
 
// Unique DOM node identifier used to reference a node that may not have been pushed to the front-end.
type ChromeDOMBackendNodeId int 
 
 
// Pseudo element type.
type ChromeDOMPseudoType string // possible values: before, after
 
 
// Shadow root type.
type ChromeDOMShadowRootType string // possible values: user-agent, author
 
 
// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type ChromeDOMNode struct {
	NodeId *ChromeDOMNodeId `json:"nodeId"` // Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeType int `json:"nodeType"` // <code>Node</code>'s nodeType.
	NodeName string `json:"nodeName"` // <code>Node</code>'s nodeName.
	LocalName string `json:"localName"` // <code>Node</code>'s localName.
	NodeValue string `json:"nodeValue"` // <code>Node</code>'s nodeValue.
	ChildNodeCount int `json:"childNodeCount,omitempty"` // Child count for <code>Container</code> nodes.
	Children []*ChromeDOMNode `json:"children,omitempty"` // Child nodes of this node when requested with children.
	Attributes []string `json:"attributes,omitempty"` // Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
	DocumentURL string `json:"documentURL,omitempty"` // Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL string `json:"baseURL,omitempty"` // Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	PublicId string `json:"publicId,omitempty"` // <code>DocumentType</code>'s publicId.
	SystemId string `json:"systemId,omitempty"` // <code>DocumentType</code>'s systemId.
	InternalSubset string `json:"internalSubset,omitempty"` // <code>DocumentType</code>'s internalSubset.
	XmlVersion string `json:"xmlVersion,omitempty"` // <code>Document</code>'s XML version in case of XML documents.
	Name string `json:"name,omitempty"` // <code>Attr</code>'s name.
	Value string `json:"value,omitempty"` // <code>Attr</code>'s value.
	PseudoType *ChromeDOMPseudoType `json:"pseudoType,omitempty"` // Pseudo element type for this node.
	ShadowRootType *ChromeDOMShadowRootType `json:"shadowRootType,omitempty"` // Shadow root type.
	FrameId *ChromePageFrameId `json:"frameId,omitempty"` // Frame ID for frame owner elements.
	ContentDocument *ChromeDOMNode `json:"contentDocument,omitempty"` // Content document for frame owner elements.
	ShadowRoots []*ChromeDOMNode `json:"shadowRoots,omitempty"` // Shadow root list for given element host.
	TemplateContent *ChromeDOMNode `json:"templateContent,omitempty"` // Content document fragment for template elements.
	PseudoElements []*ChromeDOMNode `json:"pseudoElements,omitempty"` // Pseudo elements associated with this node.
	ImportedDocument *ChromeDOMNode `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
}
 
 
// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type ChromeDOMEventListener struct {
	Type string `json:"type"` // <code>EventListener</code>'s type.
	UseCapture bool `json:"useCapture"` // <code>EventListener</code>'s useCapture.
	IsAttribute bool `json:"isAttribute"` // <code>EventListener</code>'s isAttribute.
	NodeId *ChromeDOMNodeId `json:"nodeId"` // Target <code>DOMNode</code> id.
	HandlerBody string `json:"handlerBody"` // Event handler function body.
	Location *ChromeDebuggerLocation `json:"location"` // Handler code location.
	SourceName string `json:"sourceName,omitempty"` // Source script URL.
	Handler *ChromeRuntimeRemoteObject `json:"handler,omitempty"` // Event handler function value.
}
 
 
// A structure holding an RGBA color.
type ChromeDOMRGBA struct {
	R int `json:"r"` // The red component, in the [0-255] range.
	G int `json:"g"` // The green component, in the [0-255] range.
	B int `json:"b"` // The blue component, in the [0-255] range.
	A float `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}
 
 
// An array of quad vertices, x immediately followed by y for each point, points clock-wise.
type ChromeDOMQuad []float 
 
 
// Box model.
type ChromeDOMBoxModel struct {
	Content *ChromeDOMQuad `json:"content"` // Content box
	Padding *ChromeDOMQuad `json:"padding"` // Padding box
	Border *ChromeDOMQuad `json:"border"` // Border box
	Margin *ChromeDOMQuad `json:"margin"` // Margin box
	Width int `json:"width"` // Node width
	Height int `json:"height"` // Node height
	ShapeOutside *ChromeDOMShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}
 
 
// CSS Shape Outside details.
type ChromeDOMShapeOutsideInfo struct {
	Bounds *ChromeDOMQuad `json:"bounds"` // Shape bounds
	Shape []string `json:"shape"` // Shape coordinate details
	MarginShape []string `json:"marginShape"` // Margin shape bounds
}
 
 
// Rectangle.
type ChromeDOMRect struct {
	X float `json:"x"` // X coordinate
	Y float `json:"y"` // Y coordinate
	Width float `json:"width"` // Rectangle width
	Height float `json:"height"` // Rectangle height
}
 
 
// Configuration data for the highlighting of page elements.
type ChromeDOMHighlightConfig struct {
	ShowInfo bool `json:"showInfo,omitempty"` // Whether the node info tooltip should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"` // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	ContentColor *ChromeDOMRGBA `json:"contentColor,omitempty"` // The content box highlight fill color (default: transparent).
	PaddingColor *ChromeDOMRGBA `json:"paddingColor,omitempty"` // The padding highlight fill color (default: transparent).
	BorderColor *ChromeDOMRGBA `json:"borderColor,omitempty"` // The border highlight fill color (default: transparent).
	MarginColor *ChromeDOMRGBA `json:"marginColor,omitempty"` // The margin highlight fill color (default: transparent).
	EventTargetColor *ChromeDOMRGBA `json:"eventTargetColor,omitempty"` // The event target element highlight fill color (default: transparent).
	ShapeColor *ChromeDOMRGBA `json:"shapeColor,omitempty"` // The shape outside fill color (default: transparent).
	ShapeMarginColor *ChromeDOMRGBA `json:"shapeMarginColor,omitempty"` // The shape margin fill color (default: transparent).
}
 
