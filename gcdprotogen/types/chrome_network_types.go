// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Network types.
// API Version: 1.1
package types

// Unique loader identifier.
type ChromeNetworkLoaderId string

// Unique request identifier.
type ChromeNetworkRequestId string

// Number of seconds since epoch.
type ChromeNetworkTimestamp float64

// Request / response headers as keys / values of JSON object.
type ChromeNetworkHeaders map[string]interface{}

// Timing information for the request.
type ChromeNetworkResourceTiming struct {
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

// Loading priority of a resource request.
type ChromeNetworkResourcePriority string // possible values: VeryLow, Low, Medium, High, VeryHigh

// HTTP request data.
type ChromeNetworkRequest struct {
	Url              string                         `json:"url"`                        // Request URL.
	Method           string                         `json:"method"`                     // HTTP request method.
	Headers          *ChromeNetworkHeaders          `json:"headers"`                    // HTTP request headers.
	PostData         string                         `json:"postData,omitempty"`         // HTTP POST request data.
	MixedContentType string                         `json:"mixedContentType,omitempty"` // The mixed content status of the request, as defined in http://www.w3.org/TR/mixed-content/
	InitialPriority  *ChromeNetworkResourcePriority `json:"initialPriority"`            // Priority of the resource request at the time request is sent.
}

// An internal certificate ID value.
type ChromeNetworkCertificateId int

// Subject of a certificate.
type ChromeNetworkCertificateSubject struct {
	Name           string   `json:"name"`           // Certificate subject name.
	SanDnsNames    []string `json:"sanDnsNames"`    // Subject Alternative Name (SAN) DNS names.
	SanIpAddresses []string `json:"sanIpAddresses"` // Subject Alternative Name (SAN) IP addresses.
}

// Details about a request's certificate.
type ChromeNetworkCertificateDetails struct {
	Subject   *ChromeNetworkCertificateSubject `json:"subject"`   // Certificate subject.
	Issuer    string                           `json:"issuer"`    // Name of the issuing CA.
	ValidFrom *ChromeNetworkTimestamp          `json:"validFrom"` // Certificate valid from date.
	ValidTo   *ChromeNetworkTimestamp          `json:"validTo"`   // Certificate valid to (expiration) date
}

// Security details about a request.
type ChromeNetworkSecurityDetails struct {
	Protocol      string                      `json:"protocol"`      // Protocol name (e.g. "TLS 1.2" or "QUIC".
	KeyExchange   string                      `json:"keyExchange"`   // Key Exchange used by the connection.
	Cipher        string                      `json:"cipher"`        // Cipher name.
	Mac           string                      `json:"mac,omitempty"` // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId *ChromeNetworkCertificateId `json:"certificateId"` // Certificate ID value.
}

// The reason why request was blocked.
type ChromeNetworkBlockedReason string // possible values: csp, mixed-content, origin, inspector, other

// HTTP response data.
type ChromeNetworkResponse struct {
	Url                string                        `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             float64                       `json:"status"`                       // HTTP response status code.
	StatusText         string                        `json:"statusText"`                   // HTTP response status text.
	Headers            *ChromeNetworkHeaders         `json:"headers"`                      // HTTP response headers.
	HeadersText        string                        `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string                        `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     *ChromeNetworkHeaders         `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText string                        `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool                          `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionId       float64                       `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    string                        `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         int                           `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      bool                          `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  bool                          `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64                       `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *ChromeNetworkResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           string                        `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      *ChromeSecuritySecurityState  `json:"securityState"`                // Security state of the request resource.
	SecurityDetails    *ChromeNetworkSecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocket request data.
type ChromeNetworkWebSocketRequest struct {
	Headers *ChromeNetworkHeaders `json:"headers"` // HTTP request headers.
}

// WebSocket response data.
type ChromeNetworkWebSocketResponse struct {
	Status             float64               `json:"status"`                       // HTTP response status code.
	StatusText         string                `json:"statusText"`                   // HTTP response status text.
	Headers            *ChromeNetworkHeaders `json:"headers"`                      // HTTP response headers.
	HeadersText        string                `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     *ChromeNetworkHeaders `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText string                `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocket frame data.
type ChromeNetworkWebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// Information about the cached resource.
type ChromeNetworkCachedResource struct {
	Url      string                  `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     *ChromePageResourceType `json:"type"`               // Type of this resource.
	Response *ChromeNetworkResponse  `json:"response,omitempty"` // Cached response data.
	BodySize float64                 `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type ChromeNetworkInitiator struct {
	Type            string                        `json:"type"`                      // Type of this initiator.
	StackTrace      *ChromeConsoleStackTrace      `json:"stackTrace,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	Url             string                        `json:"url,omitempty"`             // Initiator URL, set for Parser type only.
	LineNumber      float64                       `json:"lineNumber,omitempty"`      // Initiator line number, set for Parser type only.
	AsyncStackTrace *ChromeConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Initiator asynchronous JavaScript stack trace, if available.
}

// Cookie object
type ChromeNetworkCookie struct {
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
