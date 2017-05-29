// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Network functionality.
// API Version: 1.2

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
	PushStart         float64 `json:"pushStart"`         // Time the server started pushing request.
	PushEnd           float64 `json:"pushEnd"`           // Time the server finished pushing request.
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
	ReferrerPolicy   string                 `json:"referrerPolicy"`             // The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload    bool                   `json:"isLinkPreload,omitempty"`    // Whether is loaded via link preload.
}

// Details of a signed certificate timestamp (SCT).
type NetworkSignedCertificateTimestamp struct {
	Status             string  `json:"status"`             // Validation status.
	Origin             string  `json:"origin"`             // Origin.
	LogDescription     string  `json:"logDescription"`     // Log name / description.
	LogId              string  `json:"logId"`              // Log ID.
	Timestamp          float64 `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string  `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string  `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string  `json:"signatureData"`      // Signature data.
}

// Security details about a request.
type NetworkSecurityDetails struct {
	Protocol                       string                               `json:"protocol"`                       // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                    string                               `json:"keyExchange"`                    // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup               string                               `json:"keyExchangeGroup,omitempty"`     // (EC)DH group used by the connection, if applicable.
	Cipher                         string                               `json:"cipher"`                         // Cipher name.
	Mac                            string                               `json:"mac,omitempty"`                  // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId                  int                                  `json:"certificateId"`                  // Certificate ID value.
	SubjectName                    string                               `json:"subjectName"`                    // Certificate subject name.
	SanList                        []string                             `json:"sanList"`                        // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                         string                               `json:"issuer"`                         // Name of the issuing CA.
	ValidFrom                      float64                              `json:"validFrom"`                      // Certificate valid from date.
	ValidTo                        float64                              `json:"validTo"`                        // Certificate valid to (expiration) date
	SignedCertificateTimestampList []*NetworkSignedCertificateTimestamp `json:"signedCertificateTimestampList"` // List of signed certificate timestamps (SCTs).
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
	SecurityState      string                  `json:"securityState"`                // Security state of the request resource. enum values: unknown, neutral, insecure, warning, secure, info
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
	Type     string           `json:"type"`               // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
	Response *NetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize float64          `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type NetworkInitiator struct {
	Type       string             `json:"type"`                 // Type of this initiator.
	Stack      *RuntimeStackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	Url        string             `json:"url,omitempty"`        // Initiator URL, set for Parser type only.
	LineNumber float64            `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type only (0-based).
}

// Cookie object
type NetworkCookie struct {
	Name     string  `json:"name"`               // Cookie name.
	Value    string  `json:"value"`              // Cookie value.
	Domain   string  `json:"domain"`             // Cookie domain.
	Path     string  `json:"path"`               // Cookie path.
	Expires  float64 `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int     `json:"size"`               // Cookie size.
	HttpOnly bool    `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool    `json:"secure"`             // True if cookie is secure.
	Session  bool    `json:"session"`            // True in case of session cookie.
	SameSite string  `json:"sameSite,omitempty"` // Cookie SameSite type. enum values: Strict, Lax
}

// Fired when resource loading priority is changed
type NetworkResourceChangedPriorityEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId   string  `json:"requestId"`   // Request identifier.
		NewPriority string  `json:"newPriority"` // New priority enum values: VeryLow, Low, Medium, High, VeryHigh
		Timestamp   float64 `json:"timestamp"`   // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when page is about to send HTTP request.
type NetworkRequestWillBeSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId        string            `json:"requestId"`                  // Request identifier.
		FrameId          string            `json:"frameId"`                    // Frame identifier.
		LoaderId         string            `json:"loaderId"`                   // Loader identifier.
		DocumentURL      string            `json:"documentURL"`                // URL of the document this request is loaded for.
		Request          *NetworkRequest   `json:"request"`                    // Request data.
		Timestamp        float64           `json:"timestamp"`                  // Timestamp.
		WallTime         float64           `json:"wallTime"`                   // UTC Timestamp.
		Initiator        *NetworkInitiator `json:"initiator"`                  // Request initiator.
		RedirectResponse *NetworkResponse  `json:"redirectResponse,omitempty"` // Redirect response data.
		Type             string            `json:"type,omitempty"`             // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
	} `json:"Params,omitempty"`
}

// Fired if request ended up loading from cache.
type NetworkRequestServedFromCacheEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string `json:"requestId"` // Request identifier.
	} `json:"Params,omitempty"`
}

// Fired when HTTP response is available.
type NetworkResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string           `json:"requestId"` // Request identifier.
		FrameId   string           `json:"frameId"`   // Frame identifier.
		LoaderId  string           `json:"loaderId"`  // Loader identifier.
		Timestamp float64          `json:"timestamp"` // Timestamp.
		Type      string           `json:"type"`      // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		Response  *NetworkResponse `json:"response"`  // Response data.
	} `json:"Params,omitempty"`
}

// Fired when data chunk was received over the network.
type NetworkDataReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId         string  `json:"requestId"`         // Request identifier.
		Timestamp         float64 `json:"timestamp"`         // Timestamp.
		DataLength        int     `json:"dataLength"`        // Data chunk length.
		EncodedDataLength int     `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
	} `json:"Params,omitempty"`
}

// Fired when HTTP request has finished loading.
type NetworkLoadingFinishedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId         string  `json:"requestId"`         // Request identifier.
		Timestamp         float64 `json:"timestamp"`         // Timestamp.
		EncodedDataLength float64 `json:"encodedDataLength"` // Total number of bytes received for this request.
	} `json:"Params,omitempty"`
}

// Fired when HTTP request has failed to load.
type NetworkLoadingFailedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId     string  `json:"requestId"`               // Request identifier.
		Timestamp     float64 `json:"timestamp"`               // Timestamp.
		Type          string  `json:"type"`                    // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		ErrorText     string  `json:"errorText"`               // User friendly error message.
		Canceled      bool    `json:"canceled,omitempty"`      // True if loading was canceled.
		BlockedReason string  `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any. enum values: csp, mixed-content, origin, inspector, subresource-filter, other
	} `json:"Params,omitempty"`
}

// Fired when WebSocket is about to initiate handshake.
type NetworkWebSocketWillSendHandshakeRequestEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                   `json:"requestId"` // Request identifier.
		Timestamp float64                  `json:"timestamp"` // Timestamp.
		WallTime  float64                  `json:"wallTime"`  // UTC Timestamp.
		Request   *NetworkWebSocketRequest `json:"request"`   // WebSocket request data.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket handshake response becomes available.
type NetworkWebSocketHandshakeResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                    `json:"requestId"` // Request identifier.
		Timestamp float64                   `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketResponse `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired upon WebSocket creation.
type NetworkWebSocketCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string            `json:"requestId"`           // Request identifier.
		Url       string            `json:"url"`                 // WebSocket request URL.
		Initiator *NetworkInitiator `json:"initiator,omitempty"` // Request initiator.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket is closed.
type NetworkWebSocketClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string  `json:"requestId"` // Request identifier.
		Timestamp float64 `json:"timestamp"` // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame is received.
type NetworkWebSocketFrameReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame error occurs.
type NetworkWebSocketFrameErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId    string  `json:"requestId"`    // Request identifier.
		Timestamp    float64 `json:"timestamp"`    // Timestamp.
		ErrorMessage string  `json:"errorMessage"` // WebSocket frame error message.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame is sent.
type NetworkWebSocketFrameSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired when EventSource message is received.
type NetworkEventSourceMessageReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string  `json:"requestId"` // Request identifier.
		Timestamp float64 `json:"timestamp"` // Timestamp.
		EventName string  `json:"eventName"` // Message type.
		EventId   string  `json:"eventId"`   // Message identifier.
		Data      string  `json:"data"`      // Message content.
	} `json:"Params,omitempty"`
}

type Network struct {
	target gcdmessage.ChromeTargeter
}

func NewNetwork(target gcdmessage.ChromeTargeter) *Network {
	c := &Network{target: target}
	return c
}

// Enable - Enables network tracking, network events will now be delivered to the client.
// maxTotalBufferSize - Buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxResourceBufferSize - Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
func (c *Network) Enable(maxTotalBufferSize int, maxResourceBufferSize int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["maxTotalBufferSize"] = maxTotalBufferSize
	paramRequest["maxResourceBufferSize"] = maxResourceBufferSize
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enable", Params: paramRequest})
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *Network) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.disable"})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
func (c *Network) SetUserAgentOverride(userAgent string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["userAgent"] = userAgent
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setUserAgentOverride", Params: paramRequest})
}

// SetExtraHTTPHeaders - Specifies whether to always send extra HTTP headers with the requests from this page.
// headers - Map with extra HTTP headers.
func (c *Network) SetExtraHTTPHeaders(headers map[string]interface{}) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["headers"] = headers
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setExtraHTTPHeaders", Params: paramRequest})
}

// GetResponseBody - Returns content served for the given request.
// requestId - Identifier of the network request to get content for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBody(requestId string) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBody", Params: paramRequest})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// SetBlockedURLs - Blocks URLs from loading.
// urls - URL patterns to block. Wildcards ('*') are allowed.
func (c *Network) SetBlockedURLs(urls []string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["urls"] = urls
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBlockedURLs", Params: paramRequest})
}

// ReplayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *Network) ReplayXHR(requestId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["requestId"] = requestId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.replayXHR", Params: paramRequest})
}

// CanClearBrowserCache - Tells whether clearing browser cache is supported.
// Returns -  result - True if browser cache can be cleared.
func (c *Network) CanClearBrowserCache() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCache"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cache.
func (c *Network) ClearBrowserCache() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCache"})
}

// CanClearBrowserCookies - Tells whether clearing browser cookies is supported.
// Returns -  result - True if browser cookies can be cleared.
func (c *Network) CanClearBrowserCookies() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCookies"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cookies.
func (c *Network) ClearBrowserCookies() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCookies"})
}

// GetCookies - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// urls - The list of URLs for which applicable cookies will be fetched
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookies(urls []string) ([]*NetworkCookie, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["urls"] = urls
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCookies", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// GetAllCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetAllCookies() ([]*NetworkCookie, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getAllCookies"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Network) DeleteCookie(cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cookieName"] = cookieName
	paramRequest["url"] = url
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.deleteCookie", Params: paramRequest})
}

// SetCookie - Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
// url - The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
// name - The name of the cookie.
// value - The value of the cookie.
// domain - If omitted, the cookie becomes a host-only cookie.
// path - Defaults to the path portion of the url parameter.
// secure - Defaults ot false.
// httpOnly - Defaults to false.
// sameSite - Defaults to browser default behavior. enum values: Strict, Lax
// expirationDate - If omitted, the cookie becomes a session cookie.
// Returns -  success - True if successfully set cookie.
func (c *Network) SetCookie(url string, name string, value string, domain string, path string, secure bool, httpOnly bool, sameSite string, expirationDate float64) (bool, error) {
	paramRequest := make(map[string]interface{}, 9)
	paramRequest["url"] = url
	paramRequest["name"] = name
	paramRequest["value"] = value
	paramRequest["domain"] = domain
	paramRequest["path"] = path
	paramRequest["secure"] = secure
	paramRequest["httpOnly"] = httpOnly
	paramRequest["sameSite"] = sameSite
	paramRequest["expirationDate"] = expirationDate
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCookie", Params: paramRequest})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// CanEmulateNetworkConditions - Tells whether emulation of network conditions is supported.
// Returns -  result - True if emulation of network conditions is supported.
func (c *Network) CanEmulateNetworkConditions() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canEmulateNetworkConditions"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// EmulateNetworkConditions - Activates emulation of network conditions.
// offline - True to emulate internet disconnection.
// latency - Additional latency (ms).
// downloadThroughput - Maximal aggregated download throughput.
// uploadThroughput - Maximal aggregated upload throughput.
// connectionType - Connection type if known. enum values: none, cellular2g, cellular3g, cellular4g, bluetooth, ethernet, wifi, wimax, other
func (c *Network) EmulateNetworkConditions(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64, connectionType string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["offline"] = offline
	paramRequest["latency"] = latency
	paramRequest["downloadThroughput"] = downloadThroughput
	paramRequest["uploadThroughput"] = uploadThroughput
	paramRequest["connectionType"] = connectionType
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.emulateNetworkConditions", Params: paramRequest})
}

// SetCacheDisabled - Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *Network) SetCacheDisabled(cacheDisabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["cacheDisabled"] = cacheDisabled
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCacheDisabled", Params: paramRequest})
}

// SetBypassServiceWorker - Toggles ignoring of service worker for each request.
// bypass - Bypass service worker and load from network.
func (c *Network) SetBypassServiceWorker(bypass bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["bypass"] = bypass
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBypassServiceWorker", Params: paramRequest})
}

// SetDataSizeLimitsForTest - For testing.
// maxTotalSize - Maximum total buffer size.
// maxResourceSize - Maximum per-resource size.
func (c *Network) SetDataSizeLimitsForTest(maxTotalSize int, maxResourceSize int) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["maxTotalSize"] = maxTotalSize
	paramRequest["maxResourceSize"] = maxResourceSize
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setDataSizeLimitsForTest", Params: paramRequest})
}

// GetCertificate - Returns the DER-encoded certificate.
// origin - Origin to get certificate for.
// Returns -  tableNames -
func (c *Network) GetCertificate(origin string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["origin"] = origin
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCertificate", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TableNames []string
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

	return chromeData.Result.TableNames, nil
}
