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

// addBlockedURL - Blocks specific URL from loading.
// url - URL to block.
func (c *ChromeNetwork) AddBlockedURL(url string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.addBlockedURL", Params: paramRequest})
}

// removeBlockedURL - Cancels blocking of a specific URL from loading.
// url - URL to stop blocking.
func (c *ChromeNetwork) RemoveBlockedURL(url string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.removeBlockedURL", Params: paramRequest})
}

// replayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *ChromeNetwork) ReplayXHR(requestId *types.ChromeNetworkRequestId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.replayXHR", Params: paramRequest})
}

// setMonitoringXHREnabled - Toggles monitoring of XMLHttpRequest. If <code>true</code>, console will receive messages upon each XHR issued.
// enabled - Monitoring enabled state.
func (c *ChromeNetwork) SetMonitoringXHREnabled(enabled bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.setMonitoringXHREnabled", Params: paramRequest})
}

// deleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *ChromeNetwork) DeleteCookie(cookieName string, url string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.deleteCookie", Params: paramRequest})
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

// setDataSizeLimitsForTest - For testing.
// maxTotalSize - Maximum total buffer size.
// maxResourceSize - Maximum per-resource size.
func (c *ChromeNetwork) SetDataSizeLimitsForTest(maxTotalSize int, maxResourceSize int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["maxTotalSize"] = maxTotalSize
	paramRequest["maxResourceSize"] = maxResourceSize
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.setDataSizeLimitsForTest", Params: paramRequest})
}

// showCertificateViewer - Displays native dialog with the certificate details.
// certificateId - Certificate id.
func (c *ChromeNetwork) ShowCertificateViewer(certificateId *types.ChromeNetworkCertificateId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["certificateId"] = certificateId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.showCertificateViewer", Params: paramRequest})
}

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

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
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

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// getCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -
// Array of cookie objects.
func (c *ChromeNetwork) GetCookies() ([]*types.ChromeNetworkCookie, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.getCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Cookies []*types.ChromeNetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// canEmulateNetworkConditions - Tells whether emulation of network conditions is supported.
// Returns -
// True if emulation of network conditions is supported.
func (c *ChromeNetwork) CanEmulateNetworkConditions() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.canEmulateNetworkConditions"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

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

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// getCertificateDetails - Returns details for the given certificate.
// Returns -
// Certificate details.
func (c *ChromeNetwork) GetCertificateDetails(certificateId *types.ChromeNetworkCertificateId) (*types.ChromeNetworkCertificateDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["certificateId"] = certificateId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Network.getCertificateDetails", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result *types.ChromeNetworkCertificateDetails
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

	return chromeData.Result.Result, nil
}
