// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Network types.
// API Version: 1.1
package main

 
// Unique loader identifier.
type ChromeNetworkLoaderId string 
 
 
// Unique request identifier.
type ChromeNetworkRequestId string 
 
 
// Number of seconds since epoch.
type ChromeNetworkTimestamp int 
 
 
// Request / response headers as keys / values of JSON object.
type ChromeNetworkHeaders struct {
}
 
 
// Timing information for the request.
type ChromeNetworkResourceTiming struct {
	RequestTime int `json:"requestTime"` // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart int `json:"proxyStart"` // Started resolving proxy.
	ProxyEnd int `json:"proxyEnd"` // Finished resolving proxy.
	DnsStart int `json:"dnsStart"` // Started DNS address resolve.
	DnsEnd int `json:"dnsEnd"` // Finished DNS address resolve.
	ConnectStart int `json:"connectStart"` // Started connecting to the remote host.
	ConnectEnd int `json:"connectEnd"` // Connected to the remote host.
	SslStart int `json:"sslStart"` // Started SSL handshake.
	SslEnd int `json:"sslEnd"` // Finished SSL handshake.
	SendStart int `json:"sendStart"` // Started sending request.
	SendEnd int `json:"sendEnd"` // Finished sending request.
	ReceiveHeadersEnd int `json:"receiveHeadersEnd"` // Finished receiving response headers.
}
 
 
// HTTP request data.
type ChromeNetworkRequest struct {
	Url string `json:"url"` // Request URL.
	Method string `json:"method"` // HTTP request method.
	Headers *ChromeNetworkHeaders `json:"headers"` // HTTP request headers.
	PostData string `json:"postData,omitempty"` // HTTP POST request data.
}
 
 
// HTTP response data.
type ChromeNetworkResponse struct {
	Url string `json:"url"` // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status int `json:"status"` // HTTP response status code.
	StatusText string `json:"statusText"` // HTTP response status text.
	Headers *ChromeNetworkHeaders `json:"headers"` // HTTP response headers.
	HeadersText string `json:"headersText,omitempty"` // HTTP response headers text.
	MimeType string `json:"mimeType"` // Resource mimeType as determined by the browser.
	RequestHeaders *ChromeNetworkHeaders `json:"requestHeaders,omitempty"` // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText string `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused bool `json:"connectionReused"` // Specifies whether physical connection was actually reused for this request.
	ConnectionId int `json:"connectionId"` // Physical connection id that was actually used for this request.
	RemoteIPAddress string `json:"remoteIPAddress,omitempty"` // Remote IP address.
	RemotePort int `json:"remotePort,omitempty"` // Remote port.
	FromDiskCache bool `json:"fromDiskCache,omitempty"` // Specifies that the request was served from the disk cache.
	EncodedDataLength int `json:"encodedDataLength"` // Total number of bytes received for this request so far.
	Timing *ChromeNetworkResourceTiming `json:"timing,omitempty"` // Timing information for the given request.
}
 
 
// WebSocket request data.
type ChromeNetworkWebSocketRequest struct {
	Headers *ChromeNetworkHeaders `json:"headers"` // HTTP request headers.
}
 
 
// WebSocket response data.
type ChromeNetworkWebSocketResponse struct {
	Status int `json:"status"` // HTTP response status code.
	StatusText string `json:"statusText"` // HTTP response status text.
	Headers *ChromeNetworkHeaders `json:"headers"` // HTTP response headers.
	HeadersText string `json:"headersText,omitempty"` // HTTP response headers text.
	RequestHeaders *ChromeNetworkHeaders `json:"requestHeaders,omitempty"` // HTTP request headers.
	RequestHeadersText string `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}
 
 
// WebSocket frame data.
type ChromeNetworkWebSocketFrame struct {
	Opcode int `json:"opcode"` // WebSocket frame opcode.
	Mask bool `json:"mask"` // WebSocke frame mask.
	PayloadData string `json:"payloadData"` // WebSocke frame payload data.
}
 
 
// Information about the cached resource.
type ChromeNetworkCachedResource struct {
	Url string `json:"url"` // Resource URL. This is the url of the original network request.
	Type *ChromePageResourceType `json:"type"` // Type of this resource.
	Response *ChromeNetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize int `json:"bodySize"` // Cached response body size.
}
 
 
// Information about the request initiator.
type ChromeNetworkInitiator struct {
	Type string `json:"type"` // Type of this initiator.
	StackTrace *ChromeConsoleStackTrace `json:"stackTrace,omitempty"` // Initiator JavaScript stack trace, set for Script only.
	Url string `json:"url,omitempty"` // Initiator URL, set for Parser type only.
	LineNumber int `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type only.
	AsyncStackTrace *ChromeConsoleAsyncStackTrace `json:"asyncStackTrace,omitempty"` // Initiator asynchronous JavaScript stack trace, if available.
}
 
