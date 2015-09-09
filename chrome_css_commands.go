// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the CSS commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) CSS() *ChromeCSS {
	if c.css == nil {
		c.css = newChromeCSS(c)
	}
	return c.css
}

type ChromeCSS struct {
	target *ChromeTarget
}

func newChromeCSS(target *ChromeTarget) *ChromeCSS {
	c := &ChromeCSS{target: target}
	return c
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *ChromeCSS) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.enable"})
}

// Disables the CSS agent for the given page.
func (c *ChromeCSS) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.disable"})
}

// setStyleSheetText - Sets the new stylesheet text.
// styleSheetId -
// text -
func (c *ChromeCSS) SetStyleSheetText(styleSheetId *types.ChromeCSSStyleSheetId, text string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["text"] = text
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setStyleSheetText", Params: paramRequest})
}

// forcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *ChromeCSS) ForcePseudoState(nodeId *types.ChromeDOMNodeId, forcedPseudoClasses []string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["forcedPseudoClasses"] = forcedPseudoClasses
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.forcePseudoState", Params: paramRequest})
}

// setEffectivePropertyValueForNode - Find a rule with the given active property for the given node and set the new value for this property
// nodeId - The element id for which to set property.
// propertyName -
// value -
func (c *ChromeCSS) SetEffectivePropertyValueForNode(nodeId *types.ChromeDOMNodeId, propertyName string, value string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["propertyName"] = propertyName
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setEffectivePropertyValueForNode", Params: paramRequest})
}

// getMediaQueries - Returns all media queries parsed by the rendering engine.
// Returns -
func (c *ChromeCSS) GetMediaQueries() ([]*types.ChromeCSSCSSMedia, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getMediaQueries"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Medias []*types.ChromeCSSCSSMedia
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Medias, nil
}

// getMatchedStylesForNode - Returns requested styles for a DOM node identified by <code>nodeId</code>.
// Returns -
// CSS rules matching this node, from all applicable stylesheets.
// Pseudo style matches for this node.
// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
func (c *ChromeCSS) GetMatchedStylesForNode(nodeId *types.ChromeDOMNodeId, excludePseudo bool, excludeInherited bool) ([]*types.ChromeCSSRuleMatch, []*types.ChromeCSSPseudoIdMatches, []*types.ChromeCSSInheritedStyleEntry, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["excludePseudo"] = excludePseudo
	paramRequest["excludeInherited"] = excludeInherited
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getMatchedStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			MatchedCSSRules []*types.ChromeCSSRuleMatch
			PseudoElements  []*types.ChromeCSSPseudoIdMatches
			Inherited       []*types.ChromeCSSInheritedStyleEntry
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, nil
}

// getInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by <code>nodeId</code>.
// Returns -
// Inline style for the specified DOM node.
// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *ChromeCSS) GetInlineStylesForNode(nodeId *types.ChromeDOMNodeId) (*types.ChromeCSSCSSStyle, *types.ChromeCSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getInlineStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			InlineStyle     *types.ChromeCSSCSSStyle
			AttributesStyle *types.ChromeCSSCSSStyle
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// getComputedStyleForNode - Returns the computed style for a DOM node identified by <code>nodeId</code>.
// Returns -
// Computed style for the specified DOM node.
func (c *ChromeCSS) GetComputedStyleForNode(nodeId *types.ChromeDOMNodeId) ([]*types.ChromeCSSCSSComputedStyleProperty, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getComputedStyleForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ComputedStyle []*types.ChromeCSSCSSComputedStyleProperty
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.ComputedStyle, nil
}

// getPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// Returns -
// Usage statistics for every employed platform font.
func (c *ChromeCSS) GetPlatformFontsForNode(nodeId *types.ChromeDOMNodeId) ([]*types.ChromeCSSPlatformFontUsage, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getPlatformFontsForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Fonts []*types.ChromeCSSPlatformFontUsage
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Fonts, nil
}

// getStyleSheetText - Returns the current textual content and the URL for a stylesheet.
// Returns -
// The stylesheet text.
func (c *ChromeCSS) GetStyleSheetText(styleSheetId *types.ChromeCSSStyleSheetId) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["styleSheetId"] = styleSheetId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getStyleSheetText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Text string
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.Text, nil
}

// setRuleSelector - Modifies the rule selector.
// Returns -
// The resulting rule after the selector modification.
func (c *ChromeCSS) SetRuleSelector(styleSheetId *types.ChromeCSSStyleSheetId, theRange *types.ChromeCSSSourceRange, selector string) (*types.ChromeCSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["selector"] = selector
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setRuleSelector", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Rule *types.ChromeCSSCSSRule
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Rule, nil
}

// setStyleText - Modifies the style text.
// Returns -
// The resulting style after the selector modification.
func (c *ChromeCSS) SetStyleText(styleSheetId *types.ChromeCSSStyleSheetId, theRange *types.ChromeCSSSourceRange, text string) (*types.ChromeCSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setStyleText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Style *types.ChromeCSSCSSStyle
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Style, nil
}

// setMediaText - Modifies the rule selector.
// Returns -
// The resulting CSS media rule after modification.
func (c *ChromeCSS) SetMediaText(styleSheetId *types.ChromeCSSStyleSheetId, theRange *types.ChromeCSSSourceRange, text string) (*types.ChromeCSSCSSMedia, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setMediaText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Media *types.ChromeCSSCSSMedia
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Media, nil
}

// createStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given <code>frameId</code>.
// Returns -
// Identifier of the created "via-inspector" stylesheet.
func (c *ChromeCSS) CreateStyleSheet(frameId *types.ChromePageFrameId) (*types.ChromeCSSStyleSheetId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.createStyleSheet", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			StyleSheetId *types.ChromeCSSStyleSheetId
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.StyleSheetId, nil
}

// addRule - Inserts a new rule with the given <code>ruleText</code> in a stylesheet with given <code>styleSheetId</code>, at the position specified by <code>location</code>.
// Returns -
// The newly created rule.
func (c *ChromeCSS) AddRule(styleSheetId *types.ChromeCSSStyleSheetId, ruleText string, location *types.ChromeCSSSourceRange) (*types.ChromeCSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["ruleText"] = ruleText
	paramRequest["location"] = location
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.addRule", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Rule *types.ChromeCSSCSSRule
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Rule, nil
}
