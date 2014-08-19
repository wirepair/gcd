// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Network commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Network() *ChromeNetwork {
	if c.network == nil {
		c.network = newChromeNetwork(c)
	}
	return c.network
}

type ChromeNetwork struct {
	target *ChromeTarget
}

func newChromeNetwork(target *ChromeTarget) *ChromeNetwork {
	c := &ChromeNetwork{target: target}
	return c
}

// start non parameterized commands
// Enables network tracking, network events will now be delivered to the client.
func (c *ChromeNetwork) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.enable"})
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *ChromeNetwork) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.disable"})
}

// Clears browser cache.
func (c *ChromeNetwork) ClearBrowserCache() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.clearBrowserCache"})
}

// Clears browser cookies.
func (c *ChromeNetwork) ClearBrowserCookies() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.clearBrowserCookies"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
func (c *ChromeNetwork) SetUserAgentOverride(userAgent string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["userAgent"] = userAgent
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.setUserAgentOverride", Params: paramRequest})
}

// setExtraHTTPHeaders - Specifies whether to always send extra HTTP headers with the requests from this page.
// headers - Map with extra HTTP headers.
func (c *ChromeNetwork) SetExtraHTTPHeaders(headers *types.ChromeNetworkHeaders) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["headers"] = headers
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.setExtraHTTPHeaders", Params: paramRequest})
}

// replayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *ChromeNetwork) ReplayXHR(requestId *types.ChromeNetworkRequestId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.replayXHR", Params: paramRequest})
}

// emulateNetworkConditions - Activates emulation of network conditions.
// offline - True to emulate internet disconnection.
// latency - Additional latency (ms).
// downloadThroughput - Maximal aggregated download throughput.
// uploadThroughput - Maximal aggregated upload throughput.
func (c *ChromeNetwork) EmulateNetworkConditions(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["offline"] = offline
	paramRequest["latency"] = latency
	paramRequest["downloadThroughput"] = downloadThroughput
	paramRequest["uploadThroughput"] = uploadThroughput
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.emulateNetworkConditions", Params: paramRequest})
}

// setCacheDisabled - Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *ChromeNetwork) SetCacheDisabled(cacheDisabled bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["cacheDisabled"] = cacheDisabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.setCacheDisabled", Params: paramRequest})
}

// end parameterized commands with no special return types

// start commands with no parameters but special return types

// canClearBrowserCache - Tells whether clearing browser cache is supported.
// Returns -
// True if browser cache can be cleared.
func (c *ChromeNetwork) CanClearBrowserCache() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.canClearBrowserCache"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}

// canClearBrowserCookies - Tells whether clearing browser cookies is supported.
// Returns -
// True if browser cookies can be cleared.
func (c *ChromeNetwork) CanClearBrowserCookies() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.canClearBrowserCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}

// end commands with no parameters but special return types

// start commands with parameters and special return types

// getResponseBody - Returns content served for the given request.
// Returns -
// Response body.
// True, if content was sent as base64.
func (c *ChromeNetwork) GetResponseBody(requestId *types.ChromeNetworkRequestId) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.getResponseBody", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", false, &ChromeRequestErr{Resp: cerr}
		}
		return "", false, err
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// loadResourceForFrontend - Loads a resource in the context of a frame on the inspected page without cross origin checks.
// Returns -
// HTTP status code.
// Response headers.
// Resource content.
func (c *ChromeNetwork) LoadResourceForFrontend(frameId *types.ChromePageFrameId, url string, requestHeaders *types.ChromeNetworkHeaders) (float64, *types.ChromeNetworkHeaders, string, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["frameId"] = frameId
	paramRequest["url"] = url
	paramRequest["requestHeaders"] = requestHeaders
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.loadResourceForFrontend", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			StatusCode      float64
			ResponseHeaders *types.ChromeNetworkHeaders
			Content         string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return 0, nil, "", &ChromeRequestErr{Resp: cerr}
		}
		return 0, nil, "", err
	}

	return chromeData.Result.StatusCode, chromeData.Result.ResponseHeaders, chromeData.Result.Content, nil
}

// end commands with parameters and special return types
