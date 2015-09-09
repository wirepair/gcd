// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Accessibility types.
// API Version: 1.1
package types

// Unique accessibility node identifier.
type ChromeAccessibilityAXNodeId string

// Enum of possible property types.
type ChromeAccessibilityAXValueType string // possible values: boolean, tristate, booleanOrUndefined, idref, idrefList, integer, number, string, token, tokenList, domRelation, role, internalRole

// Enum of possible property sources.
type ChromeAccessibilityAXPropertySourceType string // possible values: attribute, implicit, style

// A single source for a computed AX property.
type ChromeAccessibilityAXPropertySource struct {
	Name          string                                   `json:"name"`                    // The name/label of this source.
	SourceType    *ChromeAccessibilityAXPropertySourceType `json:"sourceType"`              // What type of source this is.
	Value         string                                   `json:"value"`                   // The value of this property source.
	Type          *ChromeAccessibilityAXValueType          `json:"type"`                    // What type the value should be interpreted as.
	Invalid       bool                                     `json:"invalid,omitempty"`       // Whether the value for this property is invalid.
	InvalidReason string                                   `json:"invalidReason,omitempty"` // Reason for the value being invalid, if it is.
}

//
type ChromeAccessibilityAXRelatedNode struct {
	Idref         string                  `json:"idref,omitempty"` // The IDRef value provided, if any.
	BackendNodeId *ChromeDOMBackendNodeId `json:"backendNodeId"`   // The BackendNodeId of the related node.
}

//
type ChromeAccessibilityAXProperty struct {
	Name  string                      `json:"name"`  // The name of this property.
	Value *ChromeAccessibilityAXValue `json:"value"` // The value of this property.
}

// A single computed AX property.
type ChromeAccessibilityAXValue struct {
	Type                  *ChromeAccessibilityAXValueType        `json:"type"`                            // The type of this value.
	Value                 string                                 `json:"value,omitempty"`                 // The computed value of this property.
	RelatedNodeValue      *ChromeAccessibilityAXRelatedNode      `json:"relatedNodeValue,omitempty"`      // The related node value, if any.
	RelatedNodeArrayValue []*ChromeAccessibilityAXRelatedNode    `json:"relatedNodeArrayValue,omitempty"` // Multiple relted nodes, if applicable.
	Sources               []*ChromeAccessibilityAXPropertySource `json:"sources,omitempty"`               // The sources which contributed to the computation of this property.
}

// States which apply to every AX node.
type ChromeAccessibilityAXGlobalStates string // possible values: disabled, hidden, hiddenRoot, invalid

// Attributes which apply to nodes in live regions.
type ChromeAccessibilityAXLiveRegionAttributes string // possible values: live, atomic, relevant, busy, root

// Attributes which apply to widgets.
type ChromeAccessibilityAXWidgetAttributes string // possible values: autocomplete, haspopup, level, multiselectable, orientation, multiline, readonly, required, valuemin, valuemax, valuetext

// States which apply to widgets.
type ChromeAccessibilityAXWidgetStates string // possible values: checked, expanded, pressed, selected

// Relationships between elements other than parent/child/sibling.
type ChromeAccessibilityAXRelationshipAttributes string // possible values: activedescendant, flowto, controls, describedby, labelledby, owns

// A node in the accessibility tree.
type ChromeAccessibilityAXNode struct {
	NodeId         *ChromeAccessibilityAXNodeId     `json:"nodeId"`                   // Unique identifier for this node.
	Ignored        bool                             `json:"ignored"`                  // Whether this node is ignored for accessibility
	IgnoredReasons []*ChromeAccessibilityAXProperty `json:"ignoredReasons,omitempty"` // Collection of reasons why this node is hidden.
	Role           *ChromeAccessibilityAXValue      `json:"role,omitempty"`           // This <code>Node</code>'s role, whether explicit or implicit.
	Name           *ChromeAccessibilityAXValue      `json:"name,omitempty"`           // The accessible name for this <code>Node</code>.
	Description    *ChromeAccessibilityAXValue      `json:"description,omitempty"`    // The accessible description for this <code>Node</code>.
	Value          *ChromeAccessibilityAXValue      `json:"value,omitempty"`          // The value for this <code>Node</code>.
	Help           *ChromeAccessibilityAXValue      `json:"help,omitempty"`           // Help.
	Properties     []*ChromeAccessibilityAXProperty `json:"properties,omitempty"`     // All other properties
}
