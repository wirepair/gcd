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
func (c *ChromeCSS) SetStyleSheetText(styleSheetId *types.ChromeCSSStyleSheetId, text string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["text"] = text
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.setStyleSheetText", Params: paramRequest})
}

// forcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *ChromeCSS) ForcePseudoState(nodeId *types.ChromeDOMNodeId, forcedPseudoClasses []types.string) (*ChromeResponse, error) {
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
	var medias []*types.ChromeCSSCSSMedia 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CSS.getMediaQueries"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return medias, &ChromeRequestErr{Resp: cerr}
		}
		return medias, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		medias = results["medias"].([]*types.ChromeCSSCSSMedia)
	}
	return medias, nil
}


// end commands with no parameters but special return types

