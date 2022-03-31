// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CSS functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// CSS rule collection for a single pseudo style.
type CSSPseudoElementMatches struct {
	PseudoType string          `json:"pseudoType"` // Pseudo element type. enum values: first-line, first-letter, before, after, marker, backdrop, selection, target-text, spelling-error, grammar-error, highlight, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button, page-transition, page-transition-container, page-transition-image-wrapper, page-transition-outgoing-image, page-transition-incoming-image
	Matches    []*CSSRuleMatch `json:"matches"`    // Matches of CSS rules applicable to the pseudo style.
}

// Inherited CSS rule collection from ancestor node.
type CSSInheritedStyleEntry struct {
	InlineStyle     *CSSCSSStyle    `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*CSSRuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// Inherited pseudo element matches from pseudos of an ancestor node.
type CSSInheritedPseudoElementMatches struct {
	PseudoElements []*CSSPseudoElementMatches `json:"pseudoElements"` // Matches of pseudo styles from the pseudos of an ancestor node.
}

// Match data for a CSS rule.
type CSSRuleMatch struct {
	Rule              *CSSCSSRule `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int       `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Data for a simple selector (these are delimited by commas in a selector list).
type CSSValue struct {
	Text  string          `json:"text"`            // Value text.
	Range *CSSSourceRange `json:"range,omitempty"` // Value range in the underlying resource (if available).
}

// Selector list data.
type CSSSelectorList struct {
	Selectors []*CSSValue `json:"selectors"` // Selectors in the list.
	Text      string      `json:"text"`      // Rule selector text.
}

// CSS stylesheet metainformation.
type CSSCSSStyleSheetHeader struct {
	StyleSheetId  string  `json:"styleSheetId"`           // The stylesheet identifier.
	FrameId       string  `json:"frameId"`                // Owner frame identifier.
	SourceURL     string  `json:"sourceURL"`              // Stylesheet resource URL. Empty if this is a constructed stylesheet created using new CSSStyleSheet() (but non-empty if this is a constructed sylesheet imported as a CSS module script).
	SourceMapURL  string  `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin        string  `json:"origin"`                 // Stylesheet origin. enum values: injected, user-agent, inspector, regular
	Title         string  `json:"title"`                  // Stylesheet title.
	OwnerNode     int     `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled      bool    `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL  bool    `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline      bool    `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	IsMutable     bool    `json:"isMutable"`              // Whether this stylesheet is mutable. Inline stylesheets become mutable after they have been modified via CSSOM API. <link> element's stylesheets become mutable only if DevTools modifies them. Constructed stylesheets (new CSSStyleSheet()) are mutable immediately after creation.
	IsConstructed bool    `json:"isConstructed"`          // True if this stylesheet is created through new CSSStyleSheet() or imported as a CSS module script.
	StartLine     float64 `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn   float64 `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length        float64 `json:"length"`                 // Size of the content (in characters).
	EndLine       float64 `json:"endLine"`                // Line offset of the end of the stylesheet within the resource (zero based).
	EndColumn     float64 `json:"endColumn"`              // Column offset of the end of the stylesheet within the resource (zero based).
}

// CSS rule representation.
type CSSCSSRule struct {
	StyleSheetId     string                  `json:"styleSheetId,omitempty"`     // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList     *CSSSelectorList        `json:"selectorList"`               // Rule selector data.
	Origin           string                  `json:"origin"`                     // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	Style            *CSSCSSStyle            `json:"style"`                      // Associated style declaration.
	Media            []*CSSCSSMedia          `json:"media,omitempty"`            // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
	ContainerQueries []*CSSCSSContainerQuery `json:"containerQueries,omitempty"` // Container query list array (for rules involving container queries). The array enumerates container queries starting with the innermost one, going outwards.
	Supports         []*CSSCSSSupports       `json:"supports,omitempty"`         // @supports CSS at-rule array. The array enumerates @supports at-rules starting with the innermost one, going outwards.
	Layers           []*CSSCSSLayer          `json:"layers,omitempty"`           // Cascade layer array. Contains the layer hierarchy that this rule belongs to starting with the innermost layer and going outwards.
}

// CSS coverage information.
type CSSRuleUsage struct {
	StyleSheetId string  `json:"styleSheetId"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset  float64 `json:"startOffset"`  // Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset    float64 `json:"endOffset"`    // Offset of the end of the rule body from the beginning of the stylesheet.
	Used         bool    `json:"used"`         // Indicates whether the rule was actually used by some element in the page.
}

// Text range within a resource. All numbers are zero-based.
type CSSSourceRange struct {
	StartLine   int `json:"startLine"`   // Start line of range.
	StartColumn int `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int `json:"endLine"`     // End line of range
	EndColumn   int `json:"endColumn"`   // End column of range (exclusive).
}

// No Description.
type CSSShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
}

// No Description.
type CSSCSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// CSS style representation.
type CSSCSSStyle struct {
	StyleSheetId     string               `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CssProperties    []*CSSCSSProperty    `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []*CSSShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CssText          string               `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *CSSSourceRange      `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// CSS property declaration data.
type CSSCSSProperty struct {
	Name      string          `json:"name"`                // The property name.
	Value     string          `json:"value"`               // The property value.
	Important bool            `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
	Implicit  bool            `json:"implicit,omitempty"`  // Whether the property is implicit (implies `false` if absent).
	Text      string          `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool            `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies `true` if absent).
	Disabled  bool            `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *CSSSourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// CSS media rule descriptor.
type CSSCSSMedia struct {
	Text         string           `json:"text"`                   // Media query text.
	Source       string           `json:"source"`                 // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL    string           `json:"sourceURL,omitempty"`    // URL of the document containing the media query description.
	Range        *CSSSourceRange  `json:"range,omitempty"`        // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetId string           `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList    []*CSSMediaQuery `json:"mediaList,omitempty"`    // Array of media queries.
}

// Media query descriptor.
type CSSMediaQuery struct {
	Expressions []*CSSMediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                       `json:"active"`      // Whether the media query condition is satisfied.
}

// Media query expression descriptor.
type CSSMediaQueryExpression struct {
	Value          float64         `json:"value"`                    // Media query expression value.
	Unit           string          `json:"unit"`                     // Media query expression units.
	Feature        string          `json:"feature"`                  // Media query expression feature.
	ValueRange     *CSSSourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength float64         `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// CSS container query rule descriptor.
type CSSCSSContainerQuery struct {
	Text         string          `json:"text"`                   // Container query text.
	Range        *CSSSourceRange `json:"range,omitempty"`        // The associated rule header range in the enclosing stylesheet (if available).
	StyleSheetId string          `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	Name         string          `json:"name,omitempty"`         // Optional name for the container.
}

// CSS Supports at-rule descriptor.
type CSSCSSSupports struct {
	Text         string          `json:"text"`                   // Supports rule text.
	Active       bool            `json:"active"`                 // Whether the supports condition is satisfied.
	Range        *CSSSourceRange `json:"range,omitempty"`        // The associated rule header range in the enclosing stylesheet (if available).
	StyleSheetId string          `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
}

// CSS Layer at-rule descriptor.
type CSSCSSLayer struct {
	Text         string          `json:"text"`                   // Layer name.
	Range        *CSSSourceRange `json:"range,omitempty"`        // The associated rule header range in the enclosing stylesheet (if available).
	StyleSheetId string          `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
}

// CSS Layer data.
type CSSCSSLayerData struct {
	Name      string             `json:"name"`                // Layer name.
	SubLayers []*CSSCSSLayerData `json:"subLayers,omitempty"` // Direct sub-layers
	Order     float64            `json:"order"`               // Layer order. The order determines the order of the layer in the cascade order. A higher number has higher priority in the cascade order.
}

// Information about amount of glyphs that were rendered with given font.
type CSSPlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`   // Font's family name reported by platform.
	IsCustomFont bool    `json:"isCustomFont"` // Indicates if the font was downloaded or resolved locally.
	GlyphCount   float64 `json:"glyphCount"`   // Amount of glyphs that were rendered with this font.
}

// Information about font variation axes for variable fonts
type CSSFontVariationAxis struct {
	Tag          string  `json:"tag"`          // The font-variation-setting tag (a.k.a. "axis tag").
	Name         string  `json:"name"`         // Human-readable variation name in the default language (normally, "en").
	MinValue     float64 `json:"minValue"`     // The minimum value (inclusive) the font supports for this tag.
	MaxValue     float64 `json:"maxValue"`     // The maximum value (inclusive) the font supports for this tag.
	DefaultValue float64 `json:"defaultValue"` // The default value.
}

// Properties of a web font: https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions and additional information such as platformFontFamily and fontVariationAxes.
type CSSFontFace struct {
	FontFamily         string                  `json:"fontFamily"`                  // The font-family.
	FontStyle          string                  `json:"fontStyle"`                   // The font-style.
	FontVariant        string                  `json:"fontVariant"`                 // The font-variant.
	FontWeight         string                  `json:"fontWeight"`                  // The font-weight.
	FontStretch        string                  `json:"fontStretch"`                 // The font-stretch.
	UnicodeRange       string                  `json:"unicodeRange"`                // The unicode-range.
	Src                string                  `json:"src"`                         // The src.
	PlatformFontFamily string                  `json:"platformFontFamily"`          // The resolved platform font family
	FontVariationAxes  []*CSSFontVariationAxis `json:"fontVariationAxes,omitempty"` // Available variation settings (a.k.a. "axes").
}

// CSS keyframes rule representation.
type CSSCSSKeyframesRule struct {
	AnimationName *CSSValue             `json:"animationName"` // Animation name.
	Keyframes     []*CSSCSSKeyframeRule `json:"keyframes"`     // List of keyframes.
}

// CSS keyframe rule representation.
type CSSCSSKeyframeRule struct {
	StyleSheetId string       `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin       string       `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	KeyText      *CSSValue    `json:"keyText"`                // Associated key text.
	Style        *CSSCSSStyle `json:"style"`                  // Associated style declaration.
}

// A descriptor of operation to mutate style declaration text.
type CSSStyleDeclarationEdit struct {
	StyleSheetId string          `json:"styleSheetId"` // The css style sheet identifier.
	Range        *CSSSourceRange `json:"range"`        // The range of the style text in the enclosing stylesheet.
	Text         string          `json:"text"`         // New style text.
}

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded web font
type CSSFontsUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Font *CSSFontFace `json:"font,omitempty"` // The web font that has loaded.
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is added.
type CSSStyleSheetAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Header *CSSCSSStyleSheetHeader `json:"header"` // Added stylesheet metainfo.
	} `json:"Params,omitempty"`
}

// Fired whenever a stylesheet is changed as a result of the client operation.
type CSSStyleSheetChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` //
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is removed.
type CSSStyleSheetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` // Identifier of the removed stylesheet.
	} `json:"Params,omitempty"`
}

type CSS struct {
	target gcdmessage.ChromeTargeter
}

func NewCSS(target gcdmessage.ChromeTargeter) *CSS {
	c := &CSS{target: target}
	return c
}

type CSSAddRuleParams struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId string `json:"styleSheetId"`
	// The text of a new rule.
	RuleText string `json:"ruleText"`
	// Text position of a new rule in the target style sheet.
	Location *CSSSourceRange `json:"location"`
}

// AddRuleWithParams - Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the position specified by `location`.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRuleWithParams(ctx context.Context, v *CSSAddRuleParams) (*CSSCSSRule, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.addRule", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Rule *CSSCSSRule
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

	return chromeData.Result.Rule, nil
}

// AddRule - Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the position specified by `location`.
// styleSheetId - The css style sheet identifier where a new rule should be inserted.
// ruleText - The text of a new rule.
// location - Text position of a new rule in the target style sheet.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRule(ctx context.Context, styleSheetId string, ruleText string, location *CSSSourceRange) (*CSSCSSRule, error) {
	var v CSSAddRuleParams
	v.StyleSheetId = styleSheetId
	v.RuleText = ruleText
	v.Location = location
	return c.AddRuleWithParams(ctx, &v)
}

type CSSCollectClassNamesParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
}

// CollectClassNamesWithParams - Returns all class names from specified stylesheet.
// Returns -  classNames - Class name list.
func (c *CSS) CollectClassNamesWithParams(ctx context.Context, v *CSSCollectClassNamesParams) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.collectClassNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ClassNames []string
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

	return chromeData.Result.ClassNames, nil
}

// CollectClassNames - Returns all class names from specified stylesheet.
// styleSheetId -
// Returns -  classNames - Class name list.
func (c *CSS) CollectClassNames(ctx context.Context, styleSheetId string) ([]string, error) {
	var v CSSCollectClassNamesParams
	v.StyleSheetId = styleSheetId
	return c.CollectClassNamesWithParams(ctx, &v)
}

type CSSCreateStyleSheetParams struct {
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId string `json:"frameId"`
}

// CreateStyleSheetWithParams - Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheetWithParams(ctx context.Context, v *CSSCreateStyleSheetParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.createStyleSheet", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			StyleSheetId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.StyleSheetId, nil
}

// CreateStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
// frameId - Identifier of the frame where "via-inspector" stylesheet should be created.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheet(ctx context.Context, frameId string) (string, error) {
	var v CSSCreateStyleSheetParams
	v.FrameId = frameId
	return c.CreateStyleSheetWithParams(ctx, &v)
}

// Disables the CSS agent for the given page.
func (c *CSS) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.disable"})
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *CSS) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.enable"})
}

type CSSForcePseudoStateParams struct {
	// The element id for which to force the pseudo state.
	NodeId int `json:"nodeId"`
	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

// ForcePseudoStateWithParams - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
func (c *CSS) ForcePseudoStateWithParams(ctx context.Context, v *CSSForcePseudoStateParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.forcePseudoState", Params: v})
}

// ForcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *CSS) ForcePseudoState(ctx context.Context, nodeId int, forcedPseudoClasses []string) (*gcdmessage.ChromeResponse, error) {
	var v CSSForcePseudoStateParams
	v.NodeId = nodeId
	v.ForcedPseudoClasses = forcedPseudoClasses
	return c.ForcePseudoStateWithParams(ctx, &v)
}

type CSSGetBackgroundColorsParams struct {
	// Id of the node to get background colors for.
	NodeId int `json:"nodeId"`
}

// GetBackgroundColorsWithParams -
// Returns -  backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load). computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px'). computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
func (c *CSS) GetBackgroundColorsWithParams(ctx context.Context, v *CSSGetBackgroundColorsParams) ([]string, string, string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getBackgroundColors", Params: v})
	if err != nil {
		return nil, "", "", err
	}

	var chromeData struct {
		Result struct {
			BackgroundColors   []string
			ComputedFontSize   string
			ComputedFontWeight string
		}
	}

	if resp == nil {
		return nil, "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, "", "", err
	}

	return chromeData.Result.BackgroundColors, chromeData.Result.ComputedFontSize, chromeData.Result.ComputedFontWeight, nil
}

// GetBackgroundColors -
// nodeId - Id of the node to get background colors for.
// Returns -  backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load). computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px'). computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
func (c *CSS) GetBackgroundColors(ctx context.Context, nodeId int) ([]string, string, string, error) {
	var v CSSGetBackgroundColorsParams
	v.NodeId = nodeId
	return c.GetBackgroundColorsWithParams(ctx, &v)
}

type CSSGetComputedStyleForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetComputedStyleForNodeWithParams - Returns the computed style for a DOM node identified by `nodeId`.
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNodeWithParams(ctx context.Context, v *CSSGetComputedStyleForNodeParams) ([]*CSSCSSComputedStyleProperty, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getComputedStyleForNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ComputedStyle []*CSSCSSComputedStyleProperty
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

	return chromeData.Result.ComputedStyle, nil
}

// GetComputedStyleForNode - Returns the computed style for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNode(ctx context.Context, nodeId int) ([]*CSSCSSComputedStyleProperty, error) {
	var v CSSGetComputedStyleForNodeParams
	v.NodeId = nodeId
	return c.GetComputedStyleForNodeWithParams(ctx, &v)
}

type CSSGetInlineStylesForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetInlineStylesForNodeWithParams - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by `nodeId`.
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNodeWithParams(ctx context.Context, v *CSSGetInlineStylesForNodeParams) (*CSSCSSStyle, *CSSCSSStyle, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getInlineStylesForNode", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle     *CSSCSSStyle
			AttributesStyle *CSSCSSStyle
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

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// GetInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNode(ctx context.Context, nodeId int) (*CSSCSSStyle, *CSSCSSStyle, error) {
	var v CSSGetInlineStylesForNodeParams
	v.NodeId = nodeId
	return c.GetInlineStylesForNodeWithParams(ctx, &v)
}

type CSSGetMatchedStylesForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetMatchedStylesForNodeWithParams - Returns requested styles for a DOM node identified by `nodeId`.
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%"). matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root). inheritedPseudoElements - A chain of inherited pseudo element styles (from the immediate node parent up to the DOM tree root). cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (c *CSS) GetMatchedStylesForNodeWithParams(ctx context.Context, v *CSSGetMatchedStylesForNodeParams) (*CSSCSSStyle, *CSSCSSStyle, []*CSSRuleMatch, []*CSSPseudoElementMatches, []*CSSInheritedStyleEntry, []*CSSInheritedPseudoElementMatches, []*CSSCSSKeyframesRule, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMatchedStylesForNode", Params: v})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle             *CSSCSSStyle
			AttributesStyle         *CSSCSSStyle
			MatchedCSSRules         []*CSSRuleMatch
			PseudoElements          []*CSSPseudoElementMatches
			Inherited               []*CSSInheritedStyleEntry
			InheritedPseudoElements []*CSSInheritedPseudoElementMatches
			CssKeyframesRules       []*CSSCSSKeyframesRule
		}
	}

	if resp == nil {
		return nil, nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, chromeData.Result.InheritedPseudoElements, chromeData.Result.CssKeyframesRules, nil
}

// GetMatchedStylesForNode - Returns requested styles for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%"). matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root). inheritedPseudoElements - A chain of inherited pseudo element styles (from the immediate node parent up to the DOM tree root). cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (c *CSS) GetMatchedStylesForNode(ctx context.Context, nodeId int) (*CSSCSSStyle, *CSSCSSStyle, []*CSSRuleMatch, []*CSSPseudoElementMatches, []*CSSInheritedStyleEntry, []*CSSInheritedPseudoElementMatches, []*CSSCSSKeyframesRule, error) {
	var v CSSGetMatchedStylesForNodeParams
	v.NodeId = nodeId
	return c.GetMatchedStylesForNodeWithParams(ctx, &v)
}

// GetMediaQueries - Returns all media queries parsed by the rendering engine.
// Returns -  medias -
func (c *CSS) GetMediaQueries(ctx context.Context) ([]*CSSCSSMedia, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMediaQueries"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Medias []*CSSCSSMedia
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

	return chromeData.Result.Medias, nil
}

type CSSGetPlatformFontsForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetPlatformFontsForNodeWithParams - Requests information about platform fonts which we used to render child TextNodes in the given node.
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNodeWithParams(ctx context.Context, v *CSSGetPlatformFontsForNodeParams) ([]*CSSPlatformFontUsage, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getPlatformFontsForNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Fonts []*CSSPlatformFontUsage
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

	return chromeData.Result.Fonts, nil
}

// GetPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// nodeId -
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNode(ctx context.Context, nodeId int) ([]*CSSPlatformFontUsage, error) {
	var v CSSGetPlatformFontsForNodeParams
	v.NodeId = nodeId
	return c.GetPlatformFontsForNodeWithParams(ctx, &v)
}

type CSSGetStyleSheetTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
}

// GetStyleSheetTextWithParams - Returns the current textual content for a stylesheet.
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetTextWithParams(ctx context.Context, v *CSSGetStyleSheetTextParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getStyleSheetText", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Text string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Text, nil
}

// GetStyleSheetText - Returns the current textual content for a stylesheet.
// styleSheetId -
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetText(ctx context.Context, styleSheetId string) (string, error) {
	var v CSSGetStyleSheetTextParams
	v.StyleSheetId = styleSheetId
	return c.GetStyleSheetTextWithParams(ctx, &v)
}

type CSSGetLayersForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetLayersForNodeWithParams - Returns all layers parsed by the rendering engine for the tree scope of a node. Given a DOM element identified by nodeId, getLayersForNode returns the root layer for the nearest ancestor document or shadow root. The layer root contains the full layer tree for the tree scope and their ordering.
// Returns -  rootLayer -
func (c *CSS) GetLayersForNodeWithParams(ctx context.Context, v *CSSGetLayersForNodeParams) (*CSSCSSLayerData, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getLayersForNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RootLayer *CSSCSSLayerData
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

	return chromeData.Result.RootLayer, nil
}

// GetLayersForNode - Returns all layers parsed by the rendering engine for the tree scope of a node. Given a DOM element identified by nodeId, getLayersForNode returns the root layer for the nearest ancestor document or shadow root. The layer root contains the full layer tree for the tree scope and their ordering.
// nodeId -
// Returns -  rootLayer -
func (c *CSS) GetLayersForNode(ctx context.Context, nodeId int) (*CSSCSSLayerData, error) {
	var v CSSGetLayersForNodeParams
	v.NodeId = nodeId
	return c.GetLayersForNodeWithParams(ctx, &v)
}

type CSSTrackComputedStyleUpdatesParams struct {
	//
	PropertiesToTrack []*CSSCSSComputedStyleProperty `json:"propertiesToTrack"`
}

// TrackComputedStyleUpdatesWithParams - Starts tracking the given computed styles for updates. The specified array of properties replaces the one previously specified. Pass empty array to disable tracking. Use takeComputedStyleUpdates to retrieve the list of nodes that had properties modified. The changes to computed style properties are only tracked for nodes pushed to the front-end by the DOM agent. If no changes to the tracked properties occur after the node has been pushed to the front-end, no updates will be issued for the node.
func (c *CSS) TrackComputedStyleUpdatesWithParams(ctx context.Context, v *CSSTrackComputedStyleUpdatesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.trackComputedStyleUpdates", Params: v})
}

// TrackComputedStyleUpdates - Starts tracking the given computed styles for updates. The specified array of properties replaces the one previously specified. Pass empty array to disable tracking. Use takeComputedStyleUpdates to retrieve the list of nodes that had properties modified. The changes to computed style properties are only tracked for nodes pushed to the front-end by the DOM agent. If no changes to the tracked properties occur after the node has been pushed to the front-end, no updates will be issued for the node.
// propertiesToTrack -
func (c *CSS) TrackComputedStyleUpdates(ctx context.Context, propertiesToTrack []*CSSCSSComputedStyleProperty) (*gcdmessage.ChromeResponse, error) {
	var v CSSTrackComputedStyleUpdatesParams
	v.PropertiesToTrack = propertiesToTrack
	return c.TrackComputedStyleUpdatesWithParams(ctx, &v)
}

// TakeComputedStyleUpdates - Polls the next batch of computed style updates.
// Returns -  nodeIds - The list of node Ids that have their tracked computed styles updated
func (c *CSS) TakeComputedStyleUpdates(ctx context.Context) ([]int, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.takeComputedStyleUpdates"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

type CSSSetEffectivePropertyValueForNodeParams struct {
	// The element id for which to set property.
	NodeId int `json:"nodeId"`
	//
	PropertyName string `json:"propertyName"`
	//
	Value string `json:"value"`
}

// SetEffectivePropertyValueForNodeWithParams - Find a rule with the given active property for the given node and set the new value for this property
func (c *CSS) SetEffectivePropertyValueForNodeWithParams(ctx context.Context, v *CSSSetEffectivePropertyValueForNodeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setEffectivePropertyValueForNode", Params: v})
}

// SetEffectivePropertyValueForNode - Find a rule with the given active property for the given node and set the new value for this property
// nodeId - The element id for which to set property.
// propertyName -
// value -
func (c *CSS) SetEffectivePropertyValueForNode(ctx context.Context, nodeId int, propertyName string, value string) (*gcdmessage.ChromeResponse, error) {
	var v CSSSetEffectivePropertyValueForNodeParams
	v.NodeId = nodeId
	v.PropertyName = propertyName
	v.Value = value
	return c.SetEffectivePropertyValueForNodeWithParams(ctx, &v)
}

type CSSSetKeyframeKeyParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	KeyText string `json:"keyText"`
}

// SetKeyframeKeyWithParams - Modifies the keyframe rule key text.
// Returns -  keyText - The resulting key text after modification.
func (c *CSS) SetKeyframeKeyWithParams(ctx context.Context, v *CSSSetKeyframeKeyParams) (*CSSValue, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setKeyframeKey", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			KeyText *CSSValue
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

	return chromeData.Result.KeyText, nil
}

// SetKeyframeKey - Modifies the keyframe rule key text.
// styleSheetId -
// range -
// keyText -
// Returns -  keyText - The resulting key text after modification.
func (c *CSS) SetKeyframeKey(ctx context.Context, styleSheetId string, theRange *CSSSourceRange, keyText string) (*CSSValue, error) {
	var v CSSSetKeyframeKeyParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.KeyText = keyText
	return c.SetKeyframeKeyWithParams(ctx, &v)
}

type CSSSetMediaTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Text string `json:"text"`
}

// SetMediaTextWithParams - Modifies the rule selector.
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaTextWithParams(ctx context.Context, v *CSSSetMediaTextParams) (*CSSCSSMedia, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setMediaText", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Media *CSSCSSMedia
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

	return chromeData.Result.Media, nil
}

// SetMediaText - Modifies the rule selector.
// styleSheetId -
// range -
// text -
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaText(ctx context.Context, styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSMedia, error) {
	var v CSSSetMediaTextParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Text = text
	return c.SetMediaTextWithParams(ctx, &v)
}

type CSSSetContainerQueryTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Text string `json:"text"`
}

// SetContainerQueryTextWithParams - Modifies the expression of a container query.
// Returns -  containerQuery - The resulting CSS container query rule after modification.
func (c *CSS) SetContainerQueryTextWithParams(ctx context.Context, v *CSSSetContainerQueryTextParams) (*CSSCSSContainerQuery, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setContainerQueryText", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ContainerQuery *CSSCSSContainerQuery
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

	return chromeData.Result.ContainerQuery, nil
}

// SetContainerQueryText - Modifies the expression of a container query.
// styleSheetId -
// range -
// text -
// Returns -  containerQuery - The resulting CSS container query rule after modification.
func (c *CSS) SetContainerQueryText(ctx context.Context, styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSContainerQuery, error) {
	var v CSSSetContainerQueryTextParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Text = text
	return c.SetContainerQueryTextWithParams(ctx, &v)
}

type CSSSetSupportsTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Text string `json:"text"`
}

// SetSupportsTextWithParams - Modifies the expression of a supports at-rule.
// Returns -  supports - The resulting CSS Supports rule after modification.
func (c *CSS) SetSupportsTextWithParams(ctx context.Context, v *CSSSetSupportsTextParams) (*CSSCSSSupports, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setSupportsText", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Supports *CSSCSSSupports
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

	return chromeData.Result.Supports, nil
}

// SetSupportsText - Modifies the expression of a supports at-rule.
// styleSheetId -
// range -
// text -
// Returns -  supports - The resulting CSS Supports rule after modification.
func (c *CSS) SetSupportsText(ctx context.Context, styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSSupports, error) {
	var v CSSSetSupportsTextParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Text = text
	return c.SetSupportsTextWithParams(ctx, &v)
}

type CSSSetRuleSelectorParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Selector string `json:"selector"`
}

// SetRuleSelectorWithParams - Modifies the rule selector.
// Returns -  selectorList - The resulting selector list after modification.
func (c *CSS) SetRuleSelectorWithParams(ctx context.Context, v *CSSSetRuleSelectorParams) (*CSSSelectorList, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setRuleSelector", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			SelectorList *CSSSelectorList
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

	return chromeData.Result.SelectorList, nil
}

// SetRuleSelector - Modifies the rule selector.
// styleSheetId -
// range -
// selector -
// Returns -  selectorList - The resulting selector list after modification.
func (c *CSS) SetRuleSelector(ctx context.Context, styleSheetId string, theRange *CSSSourceRange, selector string) (*CSSSelectorList, error) {
	var v CSSSetRuleSelectorParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Selector = selector
	return c.SetRuleSelectorWithParams(ctx, &v)
}

type CSSSetStyleSheetTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	Text string `json:"text"`
}

// SetStyleSheetTextWithParams - Sets the new stylesheet text.
// Returns -  sourceMapURL - URL of source map associated with script (if any).
func (c *CSS) SetStyleSheetTextWithParams(ctx context.Context, v *CSSSetStyleSheetTextParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleSheetText", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SourceMapURL string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.SourceMapURL, nil
}

// SetStyleSheetText - Sets the new stylesheet text.
// styleSheetId -
// text -
// Returns -  sourceMapURL - URL of source map associated with script (if any).
func (c *CSS) SetStyleSheetText(ctx context.Context, styleSheetId string, text string) (string, error) {
	var v CSSSetStyleSheetTextParams
	v.StyleSheetId = styleSheetId
	v.Text = text
	return c.SetStyleSheetTextWithParams(ctx, &v)
}

type CSSSetStyleTextsParams struct {
	//
	Edits []*CSSStyleDeclarationEdit `json:"edits"`
}

// SetStyleTextsWithParams - Applies specified style edits one after another in the given order.
// Returns -  styles - The resulting styles after modification.
func (c *CSS) SetStyleTextsWithParams(ctx context.Context, v *CSSSetStyleTextsParams) ([]*CSSCSSStyle, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleTexts", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Styles []*CSSCSSStyle
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

	return chromeData.Result.Styles, nil
}

// SetStyleTexts - Applies specified style edits one after another in the given order.
// edits -
// Returns -  styles - The resulting styles after modification.
func (c *CSS) SetStyleTexts(ctx context.Context, edits []*CSSStyleDeclarationEdit) ([]*CSSCSSStyle, error) {
	var v CSSSetStyleTextsParams
	v.Edits = edits
	return c.SetStyleTextsWithParams(ctx, &v)
}

// Enables the selector recording.
func (c *CSS) StartRuleUsageTracking(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.startRuleUsageTracking"})
}

// StopRuleUsageTracking - Stop tracking rule usage and return the list of rules that were used since last call to `takeCoverageDelta` (or since start of coverage instrumentation)
// Returns -  ruleUsage -
func (c *CSS) StopRuleUsageTracking(ctx context.Context) ([]*CSSRuleUsage, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.stopRuleUsageTracking"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RuleUsage []*CSSRuleUsage
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

	return chromeData.Result.RuleUsage, nil
}

// TakeCoverageDelta - Obtain list of rules that became used since last call to this method (or since start of coverage instrumentation)
// Returns -  coverage -  timestamp - Monotonically increasing time, in seconds.
func (c *CSS) TakeCoverageDelta(ctx context.Context) ([]*CSSRuleUsage, float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.takeCoverageDelta"})
	if err != nil {
		return nil, 0, err
	}

	var chromeData struct {
		Result struct {
			Coverage  []*CSSRuleUsage
			Timestamp float64
		}
	}

	if resp == nil {
		return nil, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, 0, err
	}

	return chromeData.Result.Coverage, chromeData.Result.Timestamp, nil
}

type CSSSetLocalFontsEnabledParams struct {
	// Whether rendering of local fonts is enabled.
	Enabled bool `json:"enabled"`
}

// SetLocalFontsEnabledWithParams - Enables/disables rendering of local CSS fonts (enabled by default).
func (c *CSS) SetLocalFontsEnabledWithParams(ctx context.Context, v *CSSSetLocalFontsEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setLocalFontsEnabled", Params: v})
}

// SetLocalFontsEnabled - Enables/disables rendering of local CSS fonts (enabled by default).
// enabled - Whether rendering of local fonts is enabled.
func (c *CSS) SetLocalFontsEnabled(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v CSSSetLocalFontsEnabledParams
	v.Enabled = enabled
	return c.SetLocalFontsEnabledWithParams(ctx, &v)
}
