// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Network functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Timing information for the request.
type NetworkResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`       // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart        float64 `json:"proxyStart"`        // Started resolving proxy.
	ProxyEnd          float64 `json:"proxyEnd"`          // Finished resolving proxy.
	DnsStart          float64 `json:"dnsStart"`          // Started DNS address resolve.
	DnsEnd            float64 `json:"dnsEnd"`            // Finished DNS address resolve.
	ConnectStart      float64 `json:"connectStart"`      // Started connecting to the remote host.
	ConnectEnd        float64 `json:"connectEnd"`        // Connected to the remote host.
	SslStart          float64 `json:"sslStart"`          // Started SSL handshake.
	SslEnd            float64 `json:"sslEnd"`            // Finished SSL handshake.
	WorkerStart       float64 `json:"workerStart"`       // Started running ServiceWorker.
	WorkerReady       float64 `json:"workerReady"`       // Finished Starting ServiceWorker.
	SendStart         float64 `json:"sendStart"`         // Started sending request.
	SendEnd           float64 `json:"sendEnd"`           // Finished sending request.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// HTTP request data.
type NetworkRequest struct {
	Url              string                 `json:"url"`                        // Request URL.
	Method           string                 `json:"method"`                     // HTTP request method.
	Headers          map[string]interface{} `json:"headers"`                    // HTTP request headers.
	PostData         string                 `json:"postData,omitempty"`         // HTTP POST request data.
	MixedContentType string                 `json:"mixedContentType,omitempty"` // The mixed content status of the request, as defined in http://www.w3.org/TR/mixed-content/
	InitialPriority  string                 `json:"initialPriority"`            // Priority of the resource request at the time request is sent. enum values: VeryLow, Low, Medium, High, VeryHigh
}

// Subject of a certificate.
type NetworkCertificateSubject struct {
	Name           string   `json:"name"`           // Certificate subject name.
	SanDnsNames    []string `json:"sanDnsNames"`    // Subject Alternative Name (SAN) DNS names.
	SanIpAddresses []string `json:"sanIpAddresses"` // Subject Alternative Name (SAN) IP addresses.
}

// Details about a request's certificate.
type NetworkCertificateDetails struct {
	Subject   *NetworkCertificateSubject `json:"subject"`   // Certificate subject.
	Issuer    string                     `json:"issuer"`    // Name of the issuing CA.
	ValidFrom float64                    `json:"validFrom"` // Certificate valid from date.
	ValidTo   float64                    `json:"validTo"`   // Certificate valid to (expiration) date
}

// Security details about a request.
type NetworkSecurityDetails struct {
	Protocol      string `json:"protocol"`      // Protocol name (e.g. "TLS 1.2" or "QUIC".
	KeyExchange   string `json:"keyExchange"`   // Key Exchange used by the connection.
	Cipher        string `json:"cipher"`        // Cipher name.
	Mac           string `json:"mac,omitempty"` // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId int    `json:"certificateId"` // Certificate ID value.
}

// HTTP response data.
type NetworkResponse struct {
	Url                string                  `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             float64                 `json:"status"`                       // HTTP response status code.
	StatusText         string                  `json:"statusText"`                   // HTTP response status text.
	Headers            map[string]interface{}  `json:"headers"`                      // HTTP response headers.
	HeadersText        string                  `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string                  `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     map[string]interface{}  `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText string                  `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool                    `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionId       float64                 `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    string                  `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         int                     `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      bool                    `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  bool                    `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64                 `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *NetworkResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           string                  `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      string                  `json:"securityState"`                // Security state of the request resource. enum values: unknown, neutral, insecure, warning, secure
	SecurityDetails    *NetworkSecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocket request data.
type NetworkWebSocketRequest struct {
	Headers map[string]interface{} `json:"headers"` // HTTP request headers.
}

// WebSocket response data.
type NetworkWebSocketResponse struct {
	Status             float64                `json:"status"`                       // HTTP response status code.
	StatusText         string                 `json:"statusText"`                   // HTTP response status text.
	Headers            map[string]interface{} `json:"headers"`                      // HTTP response headers.
	HeadersText        string                 `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     map[string]interface{} `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText string                 `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocket frame data.
type NetworkWebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// Information about the cached resource.
type NetworkCachedResource struct {
	Url      string           `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     string           `json:"type"`               // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other
	Response *NetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize float64          `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type NetworkInitiator struct {
	Type            string                  `json:"type"`                      // Type of this initiator.
	StackTrace      []*ConsoleCallFrame     `json:"stackTrace,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	Url             string                  `json:"url,omitempty"`             // Initiator URL, set for Parser type only.
	LineNumber      float64                 `json:"lineNumber,omitempty"`      // Initiator line number, set for Parser type only.
	AsyncStackTrace *ConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Initiator asynchronous JavaScript stack trace, if available.
}

// Cookie object
type NetworkCookie struct {
	Name     string  `json:"name"`     // Cookie name.
	Value    string  `json:"value"`    // Cookie value.
	Domain   string  `json:"domain"`   // Cookie domain.
	Path     string  `json:"path"`     // Cookie path.
	Expires  float64 `json:"expires"`  // Cookie expires.
	Size     int     `json:"size"`     // Cookie size.
	HttpOnly bool    `json:"httpOnly"` // True if cookie is http-only.
	Secure   bool    `json:"secure"`   // True if cookie is secure.
	Session  bool    `json:"session"`  // True in case of session cookie.
}

// Fired when page is about to send HTTP request.
type NetworkRequestWillBeSentEvent struct {
	RequestId        string            `json:"requestId"`                  // Request identifier.
	FrameId          string            `json:"frameId"`                    // Frame identifier.
	LoaderId         string            `json:"loaderId"`                   // Loader identifier.
	DocumentURL      string            `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          *NetworkRequest   `json:"request"`                    // Request data.
	Timestamp        float64           `json:"timestamp"`                  // Timestamp.
	WallTime         float64           `json:"wallTime"`                   // UTC Timestamp.
	Initiator        *NetworkInitiator `json:"initiator"`                  // Request initiator.
	RedirectResponse *NetworkResponse  `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             string            `json:"type,omitempty"`             // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other
}

// Fired if request ended up loading from cache.
type NetworkRequestServedFromCacheEvent struct {
	RequestId string `json:"requestId"` // Request identifier.
}

// Fired when HTTP response is available.
type NetworkResponseReceivedEvent struct {
	RequestId string           `json:"requestId"` // Request identifier.
	FrameId   string           `json:"frameId"`   // Frame identifier.
	LoaderId  string           `json:"loaderId"`  // Loader identifier.
	Timestamp float64          `json:"timestamp"` // Timestamp.
	Type      string           `json:"type"`      // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other
	Response  *NetworkResponse `json:"response"`  // Response data.
}

// Fired when data chunk was received over the network.
type NetworkDataReceivedEvent struct {
	RequestId         string  `json:"requestId"`         // Request identifier.
	Timestamp         float64 `json:"timestamp"`         // Timestamp.
	DataLength        int     `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int     `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// Fired when HTTP request has finished loading.
type NetworkLoadingFinishedEvent struct {
	RequestId         string  `json:"requestId"`         // Request identifier.
	Timestamp         float64 `json:"timestamp"`         // Timestamp.
	EncodedDataLength float64 `json:"encodedDataLength"` // Total number of bytes received for this request.
}

// Fired when HTTP request has failed to load.
type NetworkLoadingFailedEvent struct {
	RequestId     string  `json:"requestId"`               // Request identifier.
	Timestamp     float64 `json:"timestamp"`               // Timestamp.
	Type          string  `json:"type"`                    // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Other
	ErrorText     string  `json:"errorText"`               // User friendly error message.
	Canceled      bool    `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason string  `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any. enum values: csp, mixed-content, origin, inspector, other
}

// Fired when WebSocket is about to initiate handshake.
type NetworkWebSocketWillSendHandshakeRequestEvent struct {
	RequestId string                   `json:"requestId"` // Request identifier.
	Timestamp float64                  `json:"timestamp"` // Timestamp.
	WallTime  float64                  `json:"wallTime"`  // UTC Timestamp.
	Request   *NetworkWebSocketRequest `json:"request"`   // WebSocket request data.
}

// Fired when WebSocket handshake response becomes available.
type NetworkWebSocketHandshakeResponseReceivedEvent struct {
	RequestId string                    `json:"requestId"` // Request identifier.
	Timestamp float64                   `json:"timestamp"` // Timestamp.
	Response  *NetworkWebSocketResponse `json:"response"`  // WebSocket response data.
}

// Fired upon WebSocket creation.
type NetworkWebSocketCreatedEvent struct {
	RequestId string `json:"requestId"` // Request identifier.
	Url       string `json:"url"`       // WebSocket request URL.
}

// Fired when WebSocket is closed.
type NetworkWebSocketClosedEvent struct {
	RequestId string  `json:"requestId"` // Request identifier.
	Timestamp float64 `json:"timestamp"` // Timestamp.
}

// Fired when WebSocket frame is received.
type NetworkWebSocketFrameReceivedEvent struct {
	RequestId string                 `json:"requestId"` // Request identifier.
	Timestamp float64                `json:"timestamp"` // Timestamp.
	Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
}

// Fired when WebSocket frame error occurs.
type NetworkWebSocketFrameErrorEvent struct {
	RequestId    string  `json:"requestId"`    // Request identifier.
	Timestamp    float64 `json:"timestamp"`    // Timestamp.
	ErrorMessage string  `json:"errorMessage"` // WebSocket frame error message.
}

// Fired when WebSocket frame is sent.
type NetworkWebSocketFrameSentEvent struct {
	RequestId string                 `json:"requestId"` // Request identifier.
	Timestamp float64                `json:"timestamp"` // Timestamp.
	Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
}

// Fired when EventSource message is received.
type NetworkEventSourceMessageReceivedEvent struct {
	RequestId string  `json:"requestId"` // Request identifier.
	Timestamp float64 `json:"timestamp"` // Timestamp.
	EventName string  `json:"eventName"` // Message type.
	EventId   string  `json:"eventId"`   // Message identifier.
	Data      string  `json:"data"`      // Message content.
}

type Network struct {
	target gcdmessage.ChromeTargeter
}

func NewNetwork(target gcdmessage.ChromeTargeter) *Network {
	c := &Network{target: target}
	return c
}

// Enables network tracking, network events will now be delivered to the client.
func (c *Network) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enable"})
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *Network) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.disable"})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
func (c *Network) SetUserAgentOverride(userAgent string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["userAgent"] = userAgent
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setUserAgentOverride", Params: paramRequest})
}

// SetExtraHTTPHeaders - Specifies whether to always send extra HTTP headers with the requests from this page.
// headers - Map with extra HTTP headers.
func (c *Network) SetExtraHTTPHeaders(headers map[string]interface{}) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["headers"] = headers
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setExtraHTTPHeaders", Params: paramRequest})
}

// GetResponseBody - Returns content served for the given request.
// requestId - Identifier of the network request to get content for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBody(requestId string) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBody", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// AddBlockedURL - Blocks specific URL from loading.
// url - URL to block.
func (c *Network) AddBlockedURL(url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.addBlockedURL", Params: paramRequest})
}

// RemoveBlockedURL - Cancels blocking of a specific URL from loading.
// url - URL to stop blocking.
func (c *Network) RemoveBlockedURL(url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.removeBlockedURL", Params: paramRequest})
}

// ReplayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *Network) ReplayXHR(requestId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.replayXHR", Params: paramRequest})
}

// SetMonitoringXHREnabled - Toggles monitoring of XMLHttpRequest. If <code>true</code>, console will receive messages upon each XHR issued.
// enabled - Monitoring enabled state.
func (c *Network) SetMonitoringXHREnabled(enabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setMonitoringXHREnabled", Params: paramRequest})
}

// CanClearBrowserCache - Tells whether clearing browser cache is supported.
// Returns -  result - True if browser cache can be cleared.
func (c *Network) CanClearBrowserCache() (bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCache"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cache.
func (c *Network) ClearBrowserCache() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCache"})
}

// CanClearBrowserCookies - Tells whether clearing browser cookies is supported.
// Returns -  result - True if browser cookies can be cleared.
func (c *Network) CanClearBrowserCookies() (bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cookies.
func (c *Network) ClearBrowserCookies() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCookies"})
}

// GetCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookies() ([]*NetworkCookie, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCookies"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Network) DeleteCookie(cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.deleteCookie", Params: paramRequest})
}

// CanEmulateNetworkConditions - Tells whether emulation of network conditions is supported.
// Returns -  result - True if emulation of network conditions is supported.
func (c *Network) CanEmulateNetworkConditions() (bool, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canEmulateNetworkConditions"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// EmulateNetworkConditions - Activates emulation of network conditions.
// offline - True to emulate internet disconnection.
// latency - Additional latency (ms).
// downloadThroughput - Maximal aggregated download throughput.
// uploadThroughput - Maximal aggregated upload throughput.
func (c *Network) EmulateNetworkConditions(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["offline"] = offline
	paramRequest["latency"] = latency
	paramRequest["downloadThroughput"] = downloadThroughput
	paramRequest["uploadThroughput"] = uploadThroughput
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.emulateNetworkConditions", Params: paramRequest})
}

// SetCacheDisabled - Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *Network) SetCacheDisabled(cacheDisabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["cacheDisabled"] = cacheDisabled
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCacheDisabled", Params: paramRequest})
}

// SetDataSizeLimitsForTest - For testing.
// maxTotalSize - Maximum total buffer size.
// maxResourceSize - Maximum per-resource size.
func (c *Network) SetDataSizeLimitsForTest(maxTotalSize int, maxResourceSize int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["maxTotalSize"] = maxTotalSize
	paramRequest["maxResourceSize"] = maxResourceSize
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setDataSizeLimitsForTest", Params: paramRequest})
}

// GetCertificateDetails - Returns details for the given certificate.
// certificateId - ID of the certificate to get details for.
// Returns -  result - Certificate details.
func (c *Network) GetCertificateDetails(certificateId int) (*NetworkCertificateDetails, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["certificateId"] = certificateId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCertificateDetails", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Result *NetworkCertificateDetails
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

	return chromeData.Result.Result, nil
}

// ShowCertificateViewer - Displays native dialog with the certificate details.
// certificateId - Certificate id.
func (c *Network) ShowCertificateViewer(certificateId int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["certificateId"] = certificateId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.showCertificateViewer", Params: paramRequest})
}
