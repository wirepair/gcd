// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the CSS types.
// API Version: 1.1
package types

//
type ChromeCSSStyleSheetId string

// Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via inspector" rules), "regular" for regular stylesheets.
type ChromeCSSStyleSheetOrigin string // possible values: injected, user-agent, inspector, regular

// CSS rule collection for a single pseudo style.
type ChromeCSSPseudoIdMatches struct {
	PseudoId int                   `json:"pseudoId"` // Pseudo style identifier (see <code>enum PseudoId</code> in <code>ComputedStyleConstants.h</code>).
	Matches  []*ChromeCSSRuleMatch `json:"matches"`  // Matches of CSS rules applicable to the pseudo style.
}

// Inherited CSS rule collection from ancestor node.
type ChromeCSSInheritedStyleEntry struct {
	InlineStyle     *ChromeCSSCSSStyle    `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*ChromeCSSRuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// Match data for a CSS rule.
type ChromeCSSRuleMatch struct {
	Rule              *ChromeCSSCSSRule `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int             `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Data for a simple selector (these are delimited by commas in a selector list).
type ChromeCSSSelector struct {
	Value string                `json:"value"`           // Selector text.
	Range *ChromeCSSSourceRange `json:"range,omitempty"` // Selector range in the underlying resource (if available).
}

// Selector list data.
type ChromeCSSSelectorList struct {
	Selectors []*ChromeCSSSelector `json:"selectors"` // Selectors in the list.
	Text      string               `json:"text"`      // Rule selector text.
}

// CSS stylesheet metainformation.
type ChromeCSSCSSStyleSheetHeader struct {
	StyleSheetId *ChromeCSSStyleSheetId     `json:"styleSheetId"`           // The stylesheet identifier.
	FrameId      *ChromePageFrameId         `json:"frameId"`                // Owner frame identifier.
	SourceURL    string                     `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL string                     `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       *ChromeCSSStyleSheetOrigin `json:"origin"`                 // Stylesheet origin.
	Title        string                     `json:"title"`                  // Stylesheet title.
	OwnerNode    *ChromeDOMBackendNodeId    `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool                       `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL bool                       `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool                       `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64                    `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64                    `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
}

// CSS rule representation.
type ChromeCSSCSSRule struct {
	StyleSheetId *ChromeCSSStyleSheetId     `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList *ChromeCSSSelectorList     `json:"selectorList"`           // Rule selector data.
	Origin       *ChromeCSSStyleSheetOrigin `json:"origin"`                 // Parent stylesheet's origin.
	Style        *ChromeCSSCSSStyle         `json:"style"`                  // Associated style declaration.
	Media        []*ChromeCSSCSSMedia       `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}

// Text range within a resource. All numbers are zero-based.
type ChromeCSSSourceRange struct {
	StartLine   int `json:"startLine"`   // Start line of range.
	StartColumn int `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int `json:"endLine"`     // End line of range
	EndColumn   int `json:"endColumn"`   // End column of range (exclusive).
}

//
type ChromeCSSShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
}

//
type ChromeCSSCSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// CSS style representation.
type ChromeCSSCSSStyle struct {
	StyleSheetId     *ChromeCSSStyleSheetId     `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CssProperties    []*ChromeCSSCSSProperty    `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []*ChromeCSSShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CssText          string                     `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *ChromeCSSSourceRange      `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// CSS property declaration data.
type ChromeCSSCSSProperty struct {
	Name      string                `json:"name"`                // The property name.
	Value     string                `json:"value"`               // The property value.
	Important bool                  `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Implicit  bool                  `json:"implicit,omitempty"`  // Whether the property is implicit (implies <code>false</code> if absent).
	Text      string                `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool                  `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies <code>true</code> if absent).
	Disabled  bool                  `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *ChromeCSSSourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// CSS media rule descriptor.
type ChromeCSSCSSMedia struct {
	Text               string                 `json:"text"`                         // Media query text.
	Source             string                 `json:"source"`                       // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL          string                 `json:"sourceURL,omitempty"`          // URL of the document containing the media query description.
	Range              *ChromeCSSSourceRange  `json:"range,omitempty"`              // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	ParentStyleSheetId *ChromeCSSStyleSheetId `json:"parentStyleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList          []*ChromeCSSMediaQuery `json:"mediaList,omitempty"`          // Array of media queries.
}

// Media query descriptor.
type ChromeCSSMediaQuery struct {
	Expressions []*ChromeCSSMediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                             `json:"active"`      // Whether the media query condition is satisfied.
}

// Media query expression descriptor.
type ChromeCSSMediaQueryExpression struct {
	Value          float64               `json:"value"`                    // Media query expression value.
	Unit           string                `json:"unit"`                     // Media query expression units.
	Feature        string                `json:"feature"`                  // Media query expression feature.
	ValueRange     *ChromeCSSSourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength float64               `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// Information about amount of glyphs that were rendered with given font.
type ChromeCSSPlatformFontUsage struct {
	FamilyName string  `json:"familyName"` // Font's family name reported by platform.
	GlyphCount float64 `json:"glyphCount"` // Amount of glyphs that were rendered with this font.
}
