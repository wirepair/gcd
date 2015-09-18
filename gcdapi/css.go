// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CSS functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// CSS rule collection for a single pseudo style.
type CSSPseudoIdMatches struct {
	PseudoId int             `json:"pseudoId"` // Pseudo style identifier (see <code>enum PseudoId</code> in <code>ComputedStyleConstants.h</code>).
	Matches  []*CSSRuleMatch `json:"matches"`  // Matches of CSS rules applicable to the pseudo style.
}

// Inherited CSS rule collection from ancestor node.
type CSSInheritedStyleEntry struct {
	InlineStyle     *CSSCSSStyle    `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*CSSRuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// Match data for a CSS rule.
type CSSRuleMatch struct {
	Rule              *CSSCSSRule `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int       `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Data for a simple selector (these are delimited by commas in a selector list).
type CSSSelector struct {
	Value string          `json:"value"`           // Selector text.
	Range *CSSSourceRange `json:"range,omitempty"` // Selector range in the underlying resource (if available).
}

// Selector list data.
type CSSSelectorList struct {
	Selectors []*CSSSelector `json:"selectors"` // Selectors in the list.
	Text      string         `json:"text"`      // Rule selector text.
}

// CSS stylesheet metainformation.
type CSSCSSStyleSheetHeader struct {
	StyleSheetId string  `json:"styleSheetId"`           // The stylesheet identifier.
	FrameId      string  `json:"frameId"`                // Owner frame identifier.
	SourceURL    string  `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL string  `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       string  `json:"origin"`                 // Stylesheet origin. enum values: injected, user-agent, inspector, regular
	Title        string  `json:"title"`                  // Stylesheet title.
	OwnerNode    int     `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool    `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL bool    `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool    `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64 `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64 `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
}

// CSS rule representation.
type CSSCSSRule struct {
	StyleSheetId string           `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList *CSSSelectorList `json:"selectorList"`           // Rule selector data.
	Origin       string           `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	Style        *CSSCSSStyle     `json:"style"`                  // Associated style declaration.
	Media        []*CSSCSSMedia   `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
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
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
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
	Important bool            `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Implicit  bool            `json:"implicit,omitempty"`  // Whether the property is implicit (implies <code>false</code> if absent).
	Text      string          `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool            `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies <code>true</code> if absent).
	Disabled  bool            `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *CSSSourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// CSS media rule descriptor.
type CSSCSSMedia struct {
	Text               string           `json:"text"`                         // Media query text.
	Source             string           `json:"source"`                       // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL          string           `json:"sourceURL,omitempty"`          // URL of the document containing the media query description.
	Range              *CSSSourceRange  `json:"range,omitempty"`              // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	ParentStyleSheetId string           `json:"parentStyleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList          []*CSSMediaQuery `json:"mediaList,omitempty"`          // Array of media queries.
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

// Information about amount of glyphs that were rendered with given font.
type CSSPlatformFontUsage struct {
	FamilyName string  `json:"familyName"` // Font's family name reported by platform.
	GlyphCount float64 `json:"glyphCount"` // Amount of glyphs that were rendered with this font.
}

// Fired whenever a stylesheet is changed as a result of the client operation.
type CSSStyleSheetChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` //
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is added.
type CSSStyleSheetAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Header *CSSCSSStyleSheetHeader `json:"header"` // Added stylesheet metainfo.
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is removed.
type CSSStyleSheetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` // Identifier of the removed stylesheet.
	} `json:"Params,omitempty"`
}

//
type CSSLayoutEditorChangeEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string          `json:"styleSheetId"` // Identifier of the stylesheet where the modification occurred.
		ChangeRange  *CSSSourceRange `json:"changeRange"`  // Range where the modification occurred.
	} `json:"Params,omitempty"`
}

type CSS struct {
	target gcdmessage.ChromeTargeter
}

func NewCSS(target gcdmessage.ChromeTargeter) *CSS {
	c := &CSS{target: target}
	return c
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *CSS) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.enable"})
}

// Disables the CSS agent for the given page.
func (c *CSS) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.disable"})
}

// GetMatchedStylesForNode - Returns requested styles for a DOM node identified by <code>nodeId</code>.
// nodeId -
// excludePseudo - Whether to exclude pseudo styles (default: false).
// excludeInherited - Whether to exclude inherited styles (default: false).
// Returns -  matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root).
func (c *CSS) GetMatchedStylesForNode(nodeId int, excludePseudo bool, excludeInherited bool) ([]*CSSRuleMatch, []*CSSPseudoIdMatches, []*CSSInheritedStyleEntry, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["excludePseudo"] = excludePseudo
	paramRequest["excludeInherited"] = excludeInherited
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMatchedStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			MatchedCSSRules []*CSSRuleMatch
			PseudoElements  []*CSSPseudoIdMatches
			Inherited       []*CSSInheritedStyleEntry
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, nil
}

// GetInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by <code>nodeId</code>.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNode(nodeId int) (*CSSCSSStyle, *CSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getInlineStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			InlineStyle     *CSSCSSStyle
			AttributesStyle *CSSCSSStyle
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// GetComputedStyleForNode - Returns the computed style for a DOM node identified by <code>nodeId</code>.
// nodeId -
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNode(nodeId int) ([]*CSSCSSComputedStyleProperty, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getComputedStyleForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ComputedStyle []*CSSCSSComputedStyleProperty
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

	return chromeData.Result.ComputedStyle, nil
}

// GetPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// nodeId -
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNode(nodeId int) ([]*CSSPlatformFontUsage, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getPlatformFontsForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Fonts []*CSSPlatformFontUsage
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

	return chromeData.Result.Fonts, nil
}

// GetStyleSheetText - Returns the current textual content and the URL for a stylesheet.
// styleSheetId -
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetText(styleSheetId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["styleSheetId"] = styleSheetId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getStyleSheetText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Text string
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

	return chromeData.Result.Text, nil
}

// SetStyleSheetText - Sets the new stylesheet text.
// styleSheetId -
// text -
func (c *CSS) SetStyleSheetText(styleSheetId string, text string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["text"] = text
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleSheetText", Params: paramRequest})
}

// SetRuleSelector - Modifies the rule selector.
// styleSheetId -
// range -
// selector -
// Returns -  rule - The resulting rule after the selector modification.
func (c *CSS) SetRuleSelector(styleSheetId string, theRange *CSSSourceRange, selector string) (*CSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["selector"] = selector
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setRuleSelector", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Rule *CSSCSSRule
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

	return chromeData.Result.Rule, nil
}

// SetStyleText - Modifies the style text.
// styleSheetId -
// range -
// text -
// Returns -  style - The resulting style after the selector modification.
func (c *CSS) SetStyleText(styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Style *CSSCSSStyle
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

	return chromeData.Result.Style, nil
}

// SetMediaText - Modifies the rule selector.
// styleSheetId -
// range -
// text -
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaText(styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSMedia, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setMediaText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Media *CSSCSSMedia
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

	return chromeData.Result.Media, nil
}

// CreateStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given <code>frameId</code>.
// frameId - Identifier of the frame where "via-inspector" stylesheet should be created.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheet(frameId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.createStyleSheet", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			StyleSheetId string
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

	return chromeData.Result.StyleSheetId, nil
}

// AddRule - Inserts a new rule with the given <code>ruleText</code> in a stylesheet with given <code>styleSheetId</code>, at the position specified by <code>location</code>.
// styleSheetId - The css style sheet identifier where a new rule should be inserted.
// ruleText - The text of a new rule.
// location - Text position of a new rule in the target style sheet.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRule(styleSheetId string, ruleText string, location *CSSSourceRange) (*CSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["ruleText"] = ruleText
	paramRequest["location"] = location
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.addRule", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Rule *CSSCSSRule
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

	return chromeData.Result.Rule, nil
}

// ForcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *CSS) ForcePseudoState(nodeId int, forcedPseudoClasses string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["forcedPseudoClasses"] = forcedPseudoClasses
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.forcePseudoState", Params: paramRequest})
}

// GetMediaQueries - Returns all media queries parsed by the rendering engine.
// Returns -  medias -
func (c *CSS) GetMediaQueries() ([]*CSSCSSMedia, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMediaQueries"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Medias []*CSSCSSMedia
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

	return chromeData.Result.Medias, nil
}

// SetEffectivePropertyValueForNode - Find a rule with the given active property for the given node and set the new value for this property
// nodeId - The element id for which to set property.
// propertyName -
// value -
func (c *CSS) SetEffectivePropertyValueForNode(nodeId int, propertyName string, value string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["propertyName"] = propertyName
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setEffectivePropertyValueForNode", Params: paramRequest})
}
