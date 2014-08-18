// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the CSS commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
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

// start non parameterized commands 
// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *ChromeCSS) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.enable"})
}
 
// Disables the CSS agent for the given page.
func (c *ChromeCSS) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setStyleSheetText - Sets the new stylesheet text.
// styleSheetId - 
// text - 
func (c *ChromeCSS) SetStyleSheetText(styleSheetId *types.ChromeCSSStyleSheetId, text string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["text"] = text
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setStyleSheetText", Params: paramRequest})
}

// forcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *ChromeCSS) ForcePseudoState(nodeId *types.ChromeDOMNodeId, forcedPseudoClasses []string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["forcedPseudoClasses"] = forcedPseudoClasses
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.forcePseudoState", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

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
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Medias, nil
}


// end commands with no parameters but special return types


// start commands with parameters and special return types

// getMatchedStylesForNode - Returns requested styles for a DOM node identified by <code>nodeId</code>.
// Returns - 
// CSS rules matching this node, from all applicable stylesheets.
// Pseudo style matches for this node.
// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
func (c *ChromeCSS) GetMatchedStylesForNode(nodeId *types.ChromeDOMNodeId, includePseudo bool, includeInherited bool, ) ([]*types.ChromeCSSRuleMatch, []*types.ChromeCSSPseudoIdMatches, []*types.ChromeCSSInheritedStyleEntry, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["includePseudo"] = includePseudo
	paramRequest["includeInherited"] = includeInherited
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getMatchedStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			MatchedCSSRules []*types.ChromeCSSRuleMatch 
			PseudoElements []*types.ChromeCSSPseudoIdMatches 
			Inherited []*types.ChromeCSSInheritedStyleEntry 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, nil, err
	}

	return chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, nil
}

// getInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by <code>nodeId</code>.
// Returns - 
// Inline style for the specified DOM node.
// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *ChromeCSS) GetInlineStylesForNode(nodeId *types.ChromeDOMNodeId, ) (*types.ChromeCSSCSSStyle, *types.ChromeCSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getInlineStylesForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			InlineStyle *types.ChromeCSSCSSStyle 
			AttributesStyle *types.ChromeCSSCSSStyle 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// getComputedStyleForNode - Returns the computed style for a DOM node identified by <code>nodeId</code>.
// Returns - 
// Computed style for the specified DOM node.
func (c *ChromeCSS) GetComputedStyleForNode(nodeId *types.ChromeDOMNodeId, ) ([]*types.ChromeCSSCSSComputedStyleProperty, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getComputedStyleForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			ComputedStyle []*types.ChromeCSSCSSComputedStyleProperty 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.ComputedStyle, nil
}

// getPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// Returns - 
// Font family name which is determined by computed style.
// Usage statistics for every employed platform font.
func (c *ChromeCSS) GetPlatformFontsForNode(nodeId *types.ChromeDOMNodeId, ) (string, []*types.ChromeCSSPlatformFontUsage, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getPlatformFontsForNode", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			CssFamilyName string 
			Fonts []*types.ChromeCSSPlatformFontUsage 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", nil, &ChromeRequestErr{Resp: cerr}
		}
		return "", nil, err
	}

	return chromeData.Result.CssFamilyName, chromeData.Result.Fonts, nil
}

// getStyleSheetText - Returns the current textual content and the URL for a stylesheet.
// Returns - 
// The stylesheet text.
func (c *ChromeCSS) GetStyleSheetText(styleSheetId *types.ChromeCSSStyleSheetId, ) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["styleSheetId"] = styleSheetId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getStyleSheetText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Text string 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.Text, nil
}

// setPropertyText - Either replaces a property identified by <code>styleSheetId</code> and <code>range</code> with <code>text</code> or inserts a new property <code>text</code> at the position identified by an empty <code>range</code>.
// Returns - 
// The resulting style after the property text modification.
func (c *ChromeCSS) SetPropertyText(styleSheetId *types.ChromeCSSStyleSheetId, theRange *types.ChromeCSSSourceRange, text string, ) (*types.ChromeCSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setPropertyText", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Style *types.ChromeCSSCSSStyle 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Style, nil
}

// setRuleSelector - Modifies the rule selector.
// Returns - 
// The resulting rule after the selector modification.
func (c *ChromeCSS) SetRuleSelector(styleSheetId *types.ChromeCSSStyleSheetId, theRange *types.ChromeCSSSourceRange, selector string, ) (*types.ChromeCSSCSSRule, error) {
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
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Rule, nil
}

// createStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given <code>frameId</code>.
// Returns - 
// Identifier of the created "via-inspector" stylesheet.
func (c *ChromeCSS) CreateStyleSheet(frameId *types.ChromePageFrameId, ) (*types.ChromeCSSStyleSheetId, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.createStyleSheet", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			StyleSheetId *types.ChromeCSSStyleSheetId 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.StyleSheetId, nil
}

// addRule - Creates a new empty rule with the given <code>selector</code> in a stylesheet with given <code>styleSheetId</code>.
// Returns - 
// The newly created rule.
func (c *ChromeCSS) AddRule(styleSheetId *types.ChromeCSSStyleSheetId, selector string, ) (*types.ChromeCSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["selector"] = selector
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.addRule", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Rule *types.ChromeCSSCSSRule 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Rule, nil
}


// end commands with parameters and special return types

