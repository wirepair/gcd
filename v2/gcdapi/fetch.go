// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Fetch functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type FetchRequestPattern struct {
	UrlPattern   string `json:"urlPattern,omitempty"`   // Wildcards (`'*'` -> zero or more, `'?'` -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to `"*"`.
	ResourceType string `json:"resourceType,omitempty"` // If set, only requests for matching resource types will be intercepted. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
	RequestStage string `json:"requestStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request. enum values: Request, Response
}

// Response HTTP header entry
type FetchHeaderEntry struct {
	Name  string `json:"name"`  //
	Value string `json:"value"` //
}

// Authorization challenge for HTTP status code 401 or 407.
type FetchAuthChallenge struct {
	Source string `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string `json:"origin"`           // Origin of the challenger.
	Scheme string `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string `json:"realm"`            // The realm of the challenge. May be empty.
}

// Response to an AuthChallenge.
type FetchAuthChallengeResponse struct {
	Response string `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// Issued when the domain is enabled and the request URL matches the specified filter. The request is paused until the client responds with one of continueRequest, failRequest or fulfillRequest. The stage of the request can be determined by presence of responseErrorReason and responseStatusCode -- the request is at the response stage if either of these fields is present and in the request stage otherwise.
type FetchRequestPausedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId           string              `json:"requestId"`                     // Each request the page makes will have a unique id.
		Request             *NetworkRequest     `json:"request"`                       // The details of the request.
		FrameId             string              `json:"frameId"`                       // The id of the frame that initiated the request.
		ResourceType        string              `json:"resourceType"`                  // How the requested resource will be used. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		ResponseErrorReason string              `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
		ResponseStatusCode  int                 `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage.
		ResponseStatusText  string              `json:"responseStatusText,omitempty"`  // Response status text if intercepted at response stage.
		ResponseHeaders     []*FetchHeaderEntry `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage.
		NetworkId           string              `json:"networkId,omitempty"`           // If the intercepted request had a corresponding Network.requestWillBeSent event fired for it, then this networkId will be the same as the requestId present in the requestWillBeSent event.
	} `json:"Params,omitempty"`
}

// Issued when the domain is enabled with handleAuthRequests set to true. The request is paused until client responds with continueWithAuth.
type FetchAuthRequiredEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId     string              `json:"requestId"`     // Each request the page makes will have a unique id.
		Request       *NetworkRequest     `json:"request"`       // The details of the request.
		FrameId       string              `json:"frameId"`       // The id of the frame that initiated the request.
		ResourceType  string              `json:"resourceType"`  // How the requested resource will be used. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		AuthChallenge *FetchAuthChallenge `json:"authChallenge"` // Details of the Authorization Challenge encountered. If this is set, client should respond with continueRequest that contains AuthChallengeResponse.
	} `json:"Params,omitempty"`
}

type Fetch struct {
	target gcdmessage.ChromeTargeter
}

func NewFetch(target gcdmessage.ChromeTargeter) *Fetch {
	c := &Fetch{target: target}
	return c
}

// Disables the fetch domain.
func (c *Fetch) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.disable"})
}

type FetchEnableParams struct {
	// If specified, only requests matching any of these patterns will produce fetchRequested event and will be paused until clients response. If not set, all requests will be affected.
	Patterns []*FetchRequestPattern `json:"patterns,omitempty"`
	// If true, authRequired events will be issued and requests will be paused expecting a call to continueWithAuth.
	HandleAuthRequests bool `json:"handleAuthRequests,omitempty"`
}

// EnableWithParams - Enables issuing of requestPaused events. A request will be paused until client calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
func (c *Fetch) EnableWithParams(ctx context.Context, v *FetchEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.enable", Params: v})
}

// Enable - Enables issuing of requestPaused events. A request will be paused until client calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
// patterns - If specified, only requests matching any of these patterns will produce fetchRequested event and will be paused until clients response. If not set, all requests will be affected.
// handleAuthRequests - If true, authRequired events will be issued and requests will be paused expecting a call to continueWithAuth.
func (c *Fetch) Enable(ctx context.Context, patterns []*FetchRequestPattern, handleAuthRequests bool) (*gcdmessage.ChromeResponse, error) {
	var v FetchEnableParams
	v.Patterns = patterns
	v.HandleAuthRequests = handleAuthRequests
	return c.EnableWithParams(ctx, &v)
}

type FetchFailRequestParams struct {
	// An id the client received in requestPaused event.
	RequestId string `json:"requestId"`
	// Causes the request to fail with the given reason. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
	ErrorReason string `json:"errorReason"`
}

// FailRequestWithParams - Causes the request to fail with specified reason.
func (c *Fetch) FailRequestWithParams(ctx context.Context, v *FetchFailRequestParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.failRequest", Params: v})
}

// FailRequest - Causes the request to fail with specified reason.
// requestId - An id the client received in requestPaused event.
// errorReason - Causes the request to fail with the given reason. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
func (c *Fetch) FailRequest(ctx context.Context, requestId string, errorReason string) (*gcdmessage.ChromeResponse, error) {
	var v FetchFailRequestParams
	v.RequestId = requestId
	v.ErrorReason = errorReason
	return c.FailRequestWithParams(ctx, &v)
}

type FetchFulfillRequestParams struct {
	// An id the client received in requestPaused event.
	RequestId string `json:"requestId"`
	// An HTTP response code.
	ResponseCode int `json:"responseCode"`
	// Response headers.
	ResponseHeaders []*FetchHeaderEntry `json:"responseHeaders,omitempty"`
	// Alternative way of specifying response headers as a \0-separated series of name: value pairs. Prefer the above method unless you need to represent some non-UTF8 values that can't be transmitted over the protocol as text. (Encoded as a base64 string when passed over JSON)
	BinaryResponseHeaders string `json:"binaryResponseHeaders,omitempty"`
	// A response body. If absent, original response body will be used if the request is intercepted at the response stage and empty body will be used if the request is intercepted at the request stage. (Encoded as a base64 string when passed over JSON)
	Body string `json:"body,omitempty"`
	// A textual representation of responseCode. If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`
}

// FulfillRequestWithParams - Provides response to the request.
func (c *Fetch) FulfillRequestWithParams(ctx context.Context, v *FetchFulfillRequestParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.fulfillRequest", Params: v})
}

// FulfillRequest - Provides response to the request.
// requestId - An id the client received in requestPaused event.
// responseCode - An HTTP response code.
// responseHeaders - Response headers.
// binaryResponseHeaders - Alternative way of specifying response headers as a \0-separated series of name: value pairs. Prefer the above method unless you need to represent some non-UTF8 values that can't be transmitted over the protocol as text. (Encoded as a base64 string when passed over JSON)
// body - A response body. If absent, original response body will be used if the request is intercepted at the response stage and empty body will be used if the request is intercepted at the request stage. (Encoded as a base64 string when passed over JSON)
// responsePhrase - A textual representation of responseCode. If absent, a standard phrase matching responseCode is used.
func (c *Fetch) FulfillRequest(ctx context.Context, requestId string, responseCode int, responseHeaders []*FetchHeaderEntry, binaryResponseHeaders string, body string, responsePhrase string) (*gcdmessage.ChromeResponse, error) {
	var v FetchFulfillRequestParams
	v.RequestId = requestId
	v.ResponseCode = responseCode
	v.ResponseHeaders = responseHeaders
	v.BinaryResponseHeaders = binaryResponseHeaders
	v.Body = body
	v.ResponsePhrase = responsePhrase
	return c.FulfillRequestWithParams(ctx, &v)
}

type FetchContinueRequestParams struct {
	// An id the client received in requestPaused event.
	RequestId string `json:"requestId"`
	// If set, the request url will be modified in a way that's not observable by page.
	Url string `json:"url,omitempty"`
	// If set, the request method is overridden.
	Method string `json:"method,omitempty"`
	// If set, overrides the post data in the request. (Encoded as a base64 string when passed over JSON)
	PostData string `json:"postData,omitempty"`
	// If set, overrides the request headers.
	Headers []*FetchHeaderEntry `json:"headers,omitempty"`
	// If set, overrides response interception behavior for this request.
	InterceptResponse bool `json:"interceptResponse,omitempty"`
}

// ContinueRequestWithParams - Continues the request, optionally modifying some of its parameters.
func (c *Fetch) ContinueRequestWithParams(ctx context.Context, v *FetchContinueRequestParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.continueRequest", Params: v})
}

// ContinueRequest - Continues the request, optionally modifying some of its parameters.
// requestId - An id the client received in requestPaused event.
// url - If set, the request url will be modified in a way that's not observable by page.
// method - If set, the request method is overridden.
// postData - If set, overrides the post data in the request. (Encoded as a base64 string when passed over JSON)
// headers - If set, overrides the request headers.
// interceptResponse - If set, overrides response interception behavior for this request.
func (c *Fetch) ContinueRequest(ctx context.Context, requestId string, url string, method string, postData string, headers []*FetchHeaderEntry, interceptResponse bool) (*gcdmessage.ChromeResponse, error) {
	var v FetchContinueRequestParams
	v.RequestId = requestId
	v.Url = url
	v.Method = method
	v.PostData = postData
	v.Headers = headers
	v.InterceptResponse = interceptResponse
	return c.ContinueRequestWithParams(ctx, &v)
}

type FetchContinueWithAuthParams struct {
	// An id the client received in authRequired event.
	RequestId string `json:"requestId"`
	// Response to  with an authChallenge.
	AuthChallengeResponse *FetchAuthChallengeResponse `json:"authChallengeResponse"`
}

// ContinueWithAuthWithParams - Continues a request supplying authChallengeResponse following authRequired event.
func (c *Fetch) ContinueWithAuthWithParams(ctx context.Context, v *FetchContinueWithAuthParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.continueWithAuth", Params: v})
}

// ContinueWithAuth - Continues a request supplying authChallengeResponse following authRequired event.
// requestId - An id the client received in authRequired event.
// authChallengeResponse - Response to  with an authChallenge.
func (c *Fetch) ContinueWithAuth(ctx context.Context, requestId string, authChallengeResponse *FetchAuthChallengeResponse) (*gcdmessage.ChromeResponse, error) {
	var v FetchContinueWithAuthParams
	v.RequestId = requestId
	v.AuthChallengeResponse = authChallengeResponse
	return c.ContinueWithAuthWithParams(ctx, &v)
}

type FetchContinueResponseParams struct {
	// An id the client received in requestPaused event.
	RequestId string `json:"requestId"`
	// An HTTP response code. If absent, original response code will be used.
	ResponseCode int `json:"responseCode,omitempty"`
	// A textual representation of responseCode. If absent, a standard phrase matching responseCode is used.
	ResponsePhrase string `json:"responsePhrase,omitempty"`
	// Response headers. If absent, original response headers will be used.
	ResponseHeaders []*FetchHeaderEntry `json:"responseHeaders,omitempty"`
	// Alternative way of specifying response headers as a \0-separated series of name: value pairs. Prefer the above method unless you need to represent some non-UTF8 values that can't be transmitted over the protocol as text. (Encoded as a base64 string when passed over JSON)
	BinaryResponseHeaders string `json:"binaryResponseHeaders,omitempty"`
}

// ContinueResponseWithParams - Continues loading of the paused response, optionally modifying the response headers. If either responseCode or headers are modified, all of them must be present.
func (c *Fetch) ContinueResponseWithParams(ctx context.Context, v *FetchContinueResponseParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.continueResponse", Params: v})
}

// ContinueResponse - Continues loading of the paused response, optionally modifying the response headers. If either responseCode or headers are modified, all of them must be present.
// requestId - An id the client received in requestPaused event.
// responseCode - An HTTP response code. If absent, original response code will be used.
// responsePhrase - A textual representation of responseCode. If absent, a standard phrase matching responseCode is used.
// responseHeaders - Response headers. If absent, original response headers will be used.
// binaryResponseHeaders - Alternative way of specifying response headers as a \0-separated series of name: value pairs. Prefer the above method unless you need to represent some non-UTF8 values that can't be transmitted over the protocol as text. (Encoded as a base64 string when passed over JSON)
func (c *Fetch) ContinueResponse(ctx context.Context, requestId string, responseCode int, responsePhrase string, responseHeaders []*FetchHeaderEntry, binaryResponseHeaders string) (*gcdmessage.ChromeResponse, error) {
	var v FetchContinueResponseParams
	v.RequestId = requestId
	v.ResponseCode = responseCode
	v.ResponsePhrase = responsePhrase
	v.ResponseHeaders = responseHeaders
	v.BinaryResponseHeaders = binaryResponseHeaders
	return c.ContinueResponseWithParams(ctx, &v)
}

type FetchGetResponseBodyParams struct {
	// Identifier for the intercepted request to get body for.
	RequestId string `json:"requestId"`
}

// GetResponseBodyWithParams - Causes the body of the response to be received from the server and returned as a single string. May only be issued for a request that is paused in the Response stage and is mutually exclusive with takeResponseBodyForInterceptionAsStream. Calling other methods that affect the request or disabling fetch domain before body is received results in an undefined behavior.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Fetch) GetResponseBodyWithParams(ctx context.Context, v *FetchGetResponseBodyParams) (string, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.getResponseBody", Params: v})
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

// GetResponseBody - Causes the body of the response to be received from the server and returned as a single string. May only be issued for a request that is paused in the Response stage and is mutually exclusive with takeResponseBodyForInterceptionAsStream. Calling other methods that affect the request or disabling fetch domain before body is received results in an undefined behavior.
// requestId - Identifier for the intercepted request to get body for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Fetch) GetResponseBody(ctx context.Context, requestId string) (string, bool, error) {
	var v FetchGetResponseBodyParams
	v.RequestId = requestId
	return c.GetResponseBodyWithParams(ctx, &v)
}

type FetchTakeResponseBodyAsStreamParams struct {
	//
	RequestId string `json:"requestId"`
}

// TakeResponseBodyAsStreamWithParams - Returns a handle to the stream representing the response body. The request must be paused in the HeadersReceived stage. Note that after this command the request can't be continued as is -- client either needs to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified. This method is mutually exclusive with getResponseBody. Calling other methods that affect the request or disabling fetch domain before body is received results in an undefined behavior.
// Returns -  stream -
func (c *Fetch) TakeResponseBodyAsStreamWithParams(ctx context.Context, v *FetchTakeResponseBodyAsStreamParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Fetch.takeResponseBodyAsStream", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Stream string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Stream, nil
}

// TakeResponseBodyAsStream - Returns a handle to the stream representing the response body. The request must be paused in the HeadersReceived stage. Note that after this command the request can't be continued as is -- client either needs to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified. This method is mutually exclusive with getResponseBody. Calling other methods that affect the request or disabling fetch domain before body is received results in an undefined behavior.
// requestId -
// Returns -  stream -
func (c *Fetch) TakeResponseBodyAsStream(ctx context.Context, requestId string) (string, error) {
	var v FetchTakeResponseBodyAsStreamParams
	v.RequestId = requestId
	return c.TakeResponseBodyAsStreamWithParams(ctx, &v)
}
