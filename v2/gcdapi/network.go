// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Network functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Timing information for the request.
type NetworkResourceTiming struct {
	RequestTime                 float64 `json:"requestTime"`                           // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart                  float64 `json:"proxyStart"`                            // Started resolving proxy.
	ProxyEnd                    float64 `json:"proxyEnd"`                              // Finished resolving proxy.
	DnsStart                    float64 `json:"dnsStart"`                              // Started DNS address resolve.
	DnsEnd                      float64 `json:"dnsEnd"`                                // Finished DNS address resolve.
	ConnectStart                float64 `json:"connectStart"`                          // Started connecting to the remote host.
	ConnectEnd                  float64 `json:"connectEnd"`                            // Connected to the remote host.
	SslStart                    float64 `json:"sslStart"`                              // Started SSL handshake.
	SslEnd                      float64 `json:"sslEnd"`                                // Finished SSL handshake.
	WorkerStart                 float64 `json:"workerStart"`                           // Started running ServiceWorker.
	WorkerReady                 float64 `json:"workerReady"`                           // Finished Starting ServiceWorker.
	WorkerFetchStart            float64 `json:"workerFetchStart"`                      // Started fetch event.
	WorkerRespondWithSettled    float64 `json:"workerRespondWithSettled"`              // Settled fetch event respondWith promise.
	WorkerRouterEvaluationStart float64 `json:"workerRouterEvaluationStart,omitempty"` // Started ServiceWorker static routing source evaluation.
	WorkerCacheLookupStart      float64 `json:"workerCacheLookupStart,omitempty"`      // Started cache lookup when the source was evaluated to `cache`.
	SendStart                   float64 `json:"sendStart"`                             // Started sending request.
	SendEnd                     float64 `json:"sendEnd"`                               // Finished sending request.
	PushStart                   float64 `json:"pushStart"`                             // Time the server started pushing request.
	PushEnd                     float64 `json:"pushEnd"`                               // Time the server finished pushing request.
	ReceiveHeadersStart         float64 `json:"receiveHeadersStart"`                   // Started receiving response headers.
	ReceiveHeadersEnd           float64 `json:"receiveHeadersEnd"`                     // Finished receiving response headers.
}

// Post data entry for HTTP request
type NetworkPostDataEntry struct {
}

// HTTP request data.
type NetworkRequest struct {
	Url              string                   `json:"url"`                        // Request URL (without fragment).
	UrlFragment      string                   `json:"urlFragment,omitempty"`      // Fragment of the requested URL starting with hash, if present.
	Method           string                   `json:"method"`                     // HTTP request method.
	Headers          map[string]interface{}   `json:"headers"`                    // HTTP request headers.
	PostData         string                   `json:"postData,omitempty"`         // HTTP POST request data. Use postDataEntries instead.
	HasPostData      bool                     `json:"hasPostData,omitempty"`      // True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	PostDataEntries  []*NetworkPostDataEntry  `json:"postDataEntries,omitempty"`  // Request body elements (post data broken into individual entries).
	MixedContentType string                   `json:"mixedContentType,omitempty"` // The mixed content type of the request. enum values: blockable, optionally-blockable, none
	InitialPriority  string                   `json:"initialPriority"`            // Priority of the resource request at the time request is sent. enum values: VeryLow, Low, Medium, High, VeryHigh
	ReferrerPolicy   string                   `json:"referrerPolicy"`             // The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload    bool                     `json:"isLinkPreload,omitempty"`    // Whether is loaded via link preload.
	TrustTokenParams *NetworkTrustTokenParams `json:"trustTokenParams,omitempty"` // Set for requests when the TrustToken API is used. Contains the parameters passed by the developer (e.g. via "fetch") as understood by the backend.
	IsSameSite       bool                     `json:"isSameSite,omitempty"`       // True if this resource request is considered to be the 'same site' as the request corresponding to the main frame.
}

// Details of a signed certificate timestamp (SCT).
type NetworkSignedCertificateTimestamp struct {
	Status             string  `json:"status"`             // Validation status.
	Origin             string  `json:"origin"`             // Origin.
	LogDescription     string  `json:"logDescription"`     // Log name / description.
	LogId              string  `json:"logId"`              // Log ID.
	Timestamp          float64 `json:"timestamp"`          // Issuance date. Unlike TimeSinceEpoch, this contains the number of milliseconds since January 1, 1970, UTC, not the number of seconds.
	HashAlgorithm      string  `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string  `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string  `json:"signatureData"`      // Signature data.
}

// Security details about a request.
type NetworkSecurityDetails struct {
	Protocol                          string                               `json:"protocol"`                           // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string                               `json:"keyExchange"`                        // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  string                               `json:"keyExchangeGroup,omitempty"`         // (EC)DH group used by the connection, if applicable.
	Cipher                            string                               `json:"cipher"`                             // Cipher name.
	Mac                               string                               `json:"mac,omitempty"`                      // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId                     int                                  `json:"certificateId"`                      // Certificate ID value.
	SubjectName                       string                               `json:"subjectName"`                        // Certificate subject name.
	SanList                           []string                             `json:"sanList"`                            // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string                               `json:"issuer"`                             // Name of the issuing CA.
	ValidFrom                         float64                              `json:"validFrom"`                          // Certificate valid from date.
	ValidTo                           float64                              `json:"validTo"`                            // Certificate valid to (expiration) date
	SignedCertificateTimestampList    []*NetworkSignedCertificateTimestamp `json:"signedCertificateTimestampList"`     // List of signed certificate timestamps (SCTs).
	CertificateTransparencyCompliance string                               `json:"certificateTransparencyCompliance"`  // Whether the request complied with Certificate Transparency policy enum values: unknown, not-compliant, compliant
	ServerSignatureAlgorithm          int                                  `json:"serverSignatureAlgorithm,omitempty"` // The signature algorithm used by the server in the TLS server signature, represented as a TLS SignatureScheme code point. Omitted if not applicable or not known.
	EncryptedClientHello              bool                                 `json:"encryptedClientHello"`               // Whether the connection used Encrypted ClientHello
}

// No Description.
type NetworkCorsErrorStatus struct {
	CorsError       string `json:"corsError"`       //  enum values: DisallowedByMode, InvalidResponse, WildcardOriginNotAllowed, MissingAllowOriginHeader, MultipleAllowOriginValues, InvalidAllowOriginValue, AllowOriginMismatch, InvalidAllowCredentials, CorsDisabledScheme, PreflightInvalidStatus, PreflightDisallowedRedirect, PreflightWildcardOriginNotAllowed, PreflightMissingAllowOriginHeader, PreflightMultipleAllowOriginValues, PreflightInvalidAllowOriginValue, PreflightAllowOriginMismatch, PreflightInvalidAllowCredentials, PreflightMissingAllowExternal, PreflightInvalidAllowExternal, PreflightMissingAllowPrivateNetwork, PreflightInvalidAllowPrivateNetwork, InvalidAllowMethodsPreflightResponse, InvalidAllowHeadersPreflightResponse, MethodDisallowedByPreflightResponse, HeaderDisallowedByPreflightResponse, RedirectContainsCredentials, InsecurePrivateNetwork, InvalidPrivateNetworkAccess, UnexpectedPrivateNetworkAccess, NoCorsRedirectModeNotFollow, PreflightMissingPrivateNetworkAccessId, PreflightMissingPrivateNetworkAccessName, PrivateNetworkAccessPermissionUnavailable, PrivateNetworkAccessPermissionDenied
	FailedParameter string `json:"failedParameter"` //
}

// Determines what type of Trust Token operation is executed and depending on the type, some additional parameters. The values are specified in third_party/blink/renderer/core/fetch/trust_token.idl.
type NetworkTrustTokenParams struct {
	Operation     string   `json:"operation"`         //  enum values: Issuance, Redemption, Signing
	RefreshPolicy string   `json:"refreshPolicy"`     // Only set for "token-redemption" operation and determine whether to request a fresh SRR or use a still valid cached SRR.
	Issuers       []string `json:"issuers,omitempty"` // Origins of issuers from whom to request tokens or redemption records.
}

// No Description.
type NetworkServiceWorkerRouterInfo struct {
	RuleIdMatched     int    `json:"ruleIdMatched,omitempty"`     // ID of the rule matched. If there is a matched rule, this field will be set, otherwiser no value will be set.
	MatchedSourceType string `json:"matchedSourceType,omitempty"` // The router source of the matched rule. If there is a matched rule, this field will be set, otherwise no value will be set. enum values: network, cache, fetch-event, race-network-and-fetch-handler
	ActualSourceType  string `json:"actualSourceType,omitempty"`  // The actual router source used. enum values: network, cache, fetch-event, race-network-and-fetch-handler
}

// HTTP response data.
type NetworkResponse struct {
	Url                         string                          `json:"url"`                                   // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status                      int                             `json:"status"`                                // HTTP response status code.
	StatusText                  string                          `json:"statusText"`                            // HTTP response status text.
	Headers                     map[string]interface{}          `json:"headers"`                               // HTTP response headers.
	HeadersText                 string                          `json:"headersText,omitempty"`                 // HTTP response headers text. This has been replaced by the headers in Network.responseReceivedExtraInfo.
	MimeType                    string                          `json:"mimeType"`                              // Resource mimeType as determined by the browser.
	Charset                     string                          `json:"charset"`                               // Resource charset as determined by the browser (if applicable).
	RequestHeaders              map[string]interface{}          `json:"requestHeaders,omitempty"`              // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText          string                          `json:"requestHeadersText,omitempty"`          // HTTP request headers text. This has been replaced by the headers in Network.requestWillBeSentExtraInfo.
	ConnectionReused            bool                            `json:"connectionReused"`                      // Specifies whether physical connection was actually reused for this request.
	ConnectionId                float64                         `json:"connectionId"`                          // Physical connection id that was actually used for this request.
	RemoteIPAddress             string                          `json:"remoteIPAddress,omitempty"`             // Remote IP address.
	RemotePort                  int                             `json:"remotePort,omitempty"`                  // Remote port.
	FromDiskCache               bool                            `json:"fromDiskCache,omitempty"`               // Specifies that the request was served from the disk cache.
	FromServiceWorker           bool                            `json:"fromServiceWorker,omitempty"`           // Specifies that the request was served from the ServiceWorker.
	FromPrefetchCache           bool                            `json:"fromPrefetchCache,omitempty"`           // Specifies that the request was served from the prefetch cache.
	FromEarlyHints              bool                            `json:"fromEarlyHints,omitempty"`              // Specifies that the request was served from the prefetch cache.
	ServiceWorkerRouterInfo     *NetworkServiceWorkerRouterInfo `json:"serviceWorkerRouterInfo,omitempty"`     // Information about how ServiceWorker Static Router API was used. If this field is set with `matchedSourceType` field, a matching rule is found. If this field is set without `matchedSource`, no matching rule is found. Otherwise, the API is not used.
	EncodedDataLength           float64                         `json:"encodedDataLength"`                     // Total number of bytes received for this request so far.
	Timing                      *NetworkResourceTiming          `json:"timing,omitempty"`                      // Timing information for the given request.
	ServiceWorkerResponseSource string                          `json:"serviceWorkerResponseSource,omitempty"` // Response source of response from ServiceWorker. enum values: cache-storage, http-cache, fallback-code, network
	ResponseTime                float64                         `json:"responseTime,omitempty"`                // The time at which the returned response was generated.
	CacheStorageCacheName       string                          `json:"cacheStorageCacheName,omitempty"`       // Cache Storage Cache Name.
	Protocol                    string                          `json:"protocol,omitempty"`                    // Protocol used to fetch this request.
	AlternateProtocolUsage      string                          `json:"alternateProtocolUsage,omitempty"`      // The reason why Chrome uses a specific transport protocol for HTTP semantics. enum values: alternativeJobWonWithoutRace, alternativeJobWonRace, mainJobWonRace, mappingMissing, broken, dnsAlpnH3JobWonWithoutRace, dnsAlpnH3JobWonRace, unspecifiedReason
	SecurityState               string                          `json:"securityState"`                         // Security state of the request resource. enum values: unknown, neutral, insecure, secure, info, insecure-broken
	SecurityDetails             *NetworkSecurityDetails         `json:"securityDetails,omitempty"`             // Security details for the request.
}

// WebSocket request data.
type NetworkWebSocketRequest struct {
	Headers map[string]interface{} `json:"headers"` // HTTP request headers.
}

// WebSocket response data.
type NetworkWebSocketResponse struct {
	Status             int                    `json:"status"`                       // HTTP response status code.
	StatusText         string                 `json:"statusText"`                   // HTTP response status text.
	Headers            map[string]interface{} `json:"headers"`                      // HTTP response headers.
	HeadersText        string                 `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     map[string]interface{} `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText string                 `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocket message data. This represents an entire WebSocket message, not just a fragmented frame as the name suggests.
type NetworkWebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket message opcode.
	Mask        bool    `json:"mask"`        // WebSocket message mask.
	PayloadData string  `json:"payloadData"` // WebSocket message payload data. If the opcode is 1, this is a text message and payloadData is a UTF-8 string. If the opcode isn't 1, then payloadData is a base64 encoded string representing binary data.
}

// Information about the cached resource.
type NetworkCachedResource struct {
	Url      string           `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     string           `json:"type"`               // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
	Response *NetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize float64          `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type NetworkInitiator struct {
	Type         string             `json:"type"`                   // Type of this initiator.
	Stack        *RuntimeStackTrace `json:"stack,omitempty"`        // Initiator JavaScript stack trace, set for Script only.
	Url          string             `json:"url,omitempty"`          // Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
	LineNumber   float64            `json:"lineNumber,omitempty"`   // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
	ColumnNumber float64            `json:"columnNumber,omitempty"` // Initiator column number, set for Parser type or for Script type (when script is importing module) (0-based).
	RequestId    string             `json:"requestId,omitempty"`    // Set if another request triggered this request (e.g. preflight).
}

// cookiePartitionKey object The representation of the components of the key that are created by the cookiePartitionKey class contained in net/cookies/cookie_partition_key.h.
type NetworkCookiePartitionKey struct {
	TopLevelSite         string `json:"topLevelSite"`         // The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie.
	HasCrossSiteAncestor bool   `json:"hasCrossSiteAncestor"` // Indicates if the cookie has any ancestors that are cross-site to the topLevelSite.
}

// Cookie object
type NetworkCookie struct {
	Name               string                     `json:"name"`                         // Cookie name.
	Value              string                     `json:"value"`                        // Cookie value.
	Domain             string                     `json:"domain"`                       // Cookie domain.
	Path               string                     `json:"path"`                         // Cookie path.
	Expires            float64                    `json:"expires"`                      // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size               int                        `json:"size"`                         // Cookie size.
	HttpOnly           bool                       `json:"httpOnly"`                     // True if cookie is http-only.
	Secure             bool                       `json:"secure"`                       // True if cookie is secure.
	Session            bool                       `json:"session"`                      // True in case of session cookie.
	SameSite           string                     `json:"sameSite,omitempty"`           // Cookie SameSite type. enum values: Strict, Lax, None
	Priority           string                     `json:"priority"`                     // Cookie Priority enum values: Low, Medium, High
	SameParty          bool                       `json:"sameParty"`                    // True if cookie is SameParty.
	SourceScheme       string                     `json:"sourceScheme"`                 // Cookie source scheme type. enum values: Unset, NonSecure, Secure
	SourcePort         int                        `json:"sourcePort"`                   // Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
	PartitionKey       *NetworkCookiePartitionKey `json:"partitionKey,omitempty"`       // Cookie partition key.
	PartitionKeyOpaque bool                       `json:"partitionKeyOpaque,omitempty"` // True if cookie partition key is opaque.
}

// A cookie which was not stored from a response with the corresponding reason.
type NetworkBlockedSetCookieWithReason struct {
	BlockedReasons []string       `json:"blockedReasons"`   // The reason(s) this cookie was blocked. enum values: SecureOnly, SameSiteStrict, SameSiteLax, SameSiteUnspecifiedTreatedAsLax, SameSiteNoneInsecure, UserPreferences, ThirdPartyPhaseout, ThirdPartyBlockedInFirstPartySet, SyntaxError, SchemeNotSupported, OverwriteSecure, InvalidDomain, InvalidPrefix, UnknownError, SchemefulSameSiteStrict, SchemefulSameSiteLax, SchemefulSameSiteUnspecifiedTreatedAsLax, SamePartyFromCrossPartyContext, SamePartyConflictsWithOtherAttributes, NameValuePairExceedsMaxSize, DisallowedCharacter, NoCookieContent
	CookieLine     string         `json:"cookieLine"`       // The string representing this individual cookie as it would appear in the header. This is not the entire "cookie" or "set-cookie" header which could have multiple cookies.
	Cookie         *NetworkCookie `json:"cookie,omitempty"` // The cookie object which represents the cookie which was not stored. It is optional because sometimes complete cookie information is not available, such as in the case of parsing errors.
}

// A cookie should have been blocked by 3PCD but is exempted and stored from a response with the corresponding reason. A cookie could only have at most one exemption reason.
type NetworkExemptedSetCookieWithReason struct {
	ExemptionReason string         `json:"exemptionReason"` // The reason the cookie was exempted. enum values: None, UserSetting, TPCDMetadata, TPCDDeprecationTrial, TPCDHeuristics, EnterprisePolicy, StorageAccess, TopLevelStorageAccess, CorsOptIn, Scheme
	CookieLine      string         `json:"cookieLine"`      // The string representing this individual cookie as it would appear in the header.
	Cookie          *NetworkCookie `json:"cookie"`          // The cookie object representing the cookie.
}

// A cookie associated with the request which may or may not be sent with it. Includes the cookies itself and reasons for blocking or exemption.
type NetworkAssociatedCookie struct {
	Cookie          *NetworkCookie `json:"cookie"`                    // The cookie object representing the cookie which was not sent.
	BlockedReasons  []string       `json:"blockedReasons"`            // The reason(s) the cookie was blocked. If empty means the cookie is included. enum values: SecureOnly, NotOnPath, DomainMismatch, SameSiteStrict, SameSiteLax, SameSiteUnspecifiedTreatedAsLax, SameSiteNoneInsecure, UserPreferences, ThirdPartyPhaseout, ThirdPartyBlockedInFirstPartySet, UnknownError, SchemefulSameSiteStrict, SchemefulSameSiteLax, SchemefulSameSiteUnspecifiedTreatedAsLax, SamePartyFromCrossPartyContext, NameValuePairExceedsMaxSize
	ExemptionReason string         `json:"exemptionReason,omitempty"` // The reason the cookie should have been blocked by 3PCD but is exempted. A cookie could only have at most one exemption reason. enum values: None, UserSetting, TPCDMetadata, TPCDDeprecationTrial, TPCDHeuristics, EnterprisePolicy, StorageAccess, TopLevelStorageAccess, CorsOptIn, Scheme
}

// Cookie parameter object
type NetworkCookieParam struct {
	Name         string                     `json:"name"`                   // Cookie name.
	Value        string                     `json:"value"`                  // Cookie value.
	Url          string                     `json:"url,omitempty"`          // The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
	Domain       string                     `json:"domain,omitempty"`       // Cookie domain.
	Path         string                     `json:"path,omitempty"`         // Cookie path.
	Secure       bool                       `json:"secure,omitempty"`       // True if cookie is secure.
	HttpOnly     bool                       `json:"httpOnly,omitempty"`     // True if cookie is http-only.
	SameSite     string                     `json:"sameSite,omitempty"`     // Cookie SameSite type. enum values: Strict, Lax, None
	Expires      float64                    `json:"expires,omitempty"`      // Cookie expiration date, session cookie if not set
	Priority     string                     `json:"priority,omitempty"`     // Cookie Priority. enum values: Low, Medium, High
	SameParty    bool                       `json:"sameParty,omitempty"`    // True if cookie is SameParty.
	SourceScheme string                     `json:"sourceScheme,omitempty"` // Cookie source scheme type. enum values: Unset, NonSecure, Secure
	SourcePort   int                        `json:"sourcePort,omitempty"`   // Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
	PartitionKey *NetworkCookiePartitionKey `json:"partitionKey,omitempty"` // Cookie partition key. If not set, the cookie will be set as not partitioned.
}

// Authorization challenge for HTTP status code 401 or 407.
type NetworkAuthChallenge struct {
	Source string `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string `json:"origin"`           // Origin of the challenger.
	Scheme string `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string `json:"realm"`            // The realm of the challenge. May be empty.
}

// Response to an AuthChallenge.
type NetworkAuthChallengeResponse struct {
	Response string `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// Request pattern for interception.
type NetworkRequestPattern struct {
	UrlPattern        string `json:"urlPattern,omitempty"`        // Wildcards (`'*'` -> zero or more, `'?'` -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to `"*"`.
	ResourceType      string `json:"resourceType,omitempty"`      // If set, only requests for matching resource types will be intercepted. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
	InterceptionStage string `json:"interceptionStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request. enum values: Request, HeadersReceived
}

// Information about a signed exchange signature. https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
type NetworkSignedExchangeSignature struct {
	Label        string   `json:"label"`                  // Signed exchange signature label.
	Signature    string   `json:"signature"`              // The hex string of signed exchange signature.
	Integrity    string   `json:"integrity"`              // Signed exchange signature integrity.
	CertUrl      string   `json:"certUrl,omitempty"`      // Signed exchange signature cert Url.
	CertSha256   string   `json:"certSha256,omitempty"`   // The hex string of signed exchange signature cert sha256.
	ValidityUrl  string   `json:"validityUrl"`            // Signed exchange signature validity Url.
	Date         int      `json:"date"`                   // Signed exchange signature date.
	Expires      int      `json:"expires"`                // Signed exchange signature expires.
	Certificates []string `json:"certificates,omitempty"` // The encoded certificates.
}

// Information about a signed exchange header. https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
type NetworkSignedExchangeHeader struct {
	RequestUrl      string                            `json:"requestUrl"`      // Signed exchange request URL.
	ResponseCode    int                               `json:"responseCode"`    // Signed exchange response code.
	ResponseHeaders map[string]interface{}            `json:"responseHeaders"` // Signed exchange response headers.
	Signatures      []*NetworkSignedExchangeSignature `json:"signatures"`      // Signed exchange response signature.
	HeaderIntegrity string                            `json:"headerIntegrity"` // Signed exchange header integrity hash in the form of `sha256-<base64-hash-value>`.
}

// Information about a signed exchange response.
type NetworkSignedExchangeError struct {
	Message        string `json:"message"`                  // Error message.
	SignatureIndex int    `json:"signatureIndex,omitempty"` // The index of the signature which caused the error.
	ErrorField     string `json:"errorField,omitempty"`     // The field which caused the error. enum values: signatureSig, signatureIntegrity, signatureCertUrl, signatureCertSha256, signatureValidityUrl, signatureTimestamps
}

// Information about a signed exchange response.
type NetworkSignedExchangeInfo struct {
	OuterResponse   *NetworkResponse              `json:"outerResponse"`             // The outer response of signed HTTP exchange which was received from network.
	Header          *NetworkSignedExchangeHeader  `json:"header,omitempty"`          // Information about the signed exchange header.
	SecurityDetails *NetworkSecurityDetails       `json:"securityDetails,omitempty"` // Security details for the signed exchange header.
	Errors          []*NetworkSignedExchangeError `json:"errors,omitempty"`          // Errors occurred while handling the signed exchange.
}

// No Description.
type NetworkConnectTiming struct {
	RequestTime float64 `json:"requestTime"` // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime. Matches ResourceTiming's requestTime for the same request (but not for redirected requests).
}

// No Description.
type NetworkClientSecurityState struct {
	InitiatorIsSecureContext    bool   `json:"initiatorIsSecureContext"`    //
	InitiatorIPAddressSpace     string `json:"initiatorIPAddressSpace"`     //  enum values: Local, Private, Public, Unknown
	PrivateNetworkRequestPolicy string `json:"privateNetworkRequestPolicy"` //  enum values: Allow, BlockFromInsecureToMorePrivate, WarnFromInsecureToMorePrivate, PreflightBlock, PreflightWarn
}

// No Description.
type NetworkCrossOriginOpenerPolicyStatus struct {
	Value                       string `json:"value"`                                 //  enum values: SameOrigin, SameOriginAllowPopups, RestrictProperties, UnsafeNone, SameOriginPlusCoep, RestrictPropertiesPlusCoep, NoopenerAllowPopups
	ReportOnlyValue             string `json:"reportOnlyValue"`                       //  enum values: SameOrigin, SameOriginAllowPopups, RestrictProperties, UnsafeNone, SameOriginPlusCoep, RestrictPropertiesPlusCoep, NoopenerAllowPopups
	ReportingEndpoint           string `json:"reportingEndpoint,omitempty"`           //
	ReportOnlyReportingEndpoint string `json:"reportOnlyReportingEndpoint,omitempty"` //
}

// No Description.
type NetworkCrossOriginEmbedderPolicyStatus struct {
	Value                       string `json:"value"`                                 //  enum values: None, Credentialless, RequireCorp
	ReportOnlyValue             string `json:"reportOnlyValue"`                       //  enum values: None, Credentialless, RequireCorp
	ReportingEndpoint           string `json:"reportingEndpoint,omitempty"`           //
	ReportOnlyReportingEndpoint string `json:"reportOnlyReportingEndpoint,omitempty"` //
}

// No Description.
type NetworkContentSecurityPolicyStatus struct {
	EffectiveDirectives string `json:"effectiveDirectives"` //
	IsEnforced          bool   `json:"isEnforced"`          //
	Source              string `json:"source"`              //  enum values: HTTP, Meta
}

// No Description.
type NetworkSecurityIsolationStatus struct {
	Coop *NetworkCrossOriginOpenerPolicyStatus   `json:"coop,omitempty"` //
	Coep *NetworkCrossOriginEmbedderPolicyStatus `json:"coep,omitempty"` //
	Csp  []*NetworkContentSecurityPolicyStatus   `json:"csp,omitempty"`  //
}

// An object representing a report generated by the Reporting API.
type NetworkReportingApiReport struct {
	Id                string                 `json:"id"`                //
	InitiatorUrl      string                 `json:"initiatorUrl"`      // The URL of the document that triggered the report.
	Destination       string                 `json:"destination"`       // The name of the endpoint group that should be used to deliver the report.
	Type              string                 `json:"type"`              // The type of the report (specifies the set of data that is contained in the report body).
	Timestamp         float64                `json:"timestamp"`         // When the report was generated.
	Depth             int                    `json:"depth"`             // How many uploads deep the related request was.
	CompletedAttempts int                    `json:"completedAttempts"` // The number of delivery attempts made so far, not including an active attempt.
	Body              map[string]interface{} `json:"body"`              //
	Status            string                 `json:"status"`            //  enum values: Queued, Pending, MarkedForRemoval, Success
}

// No Description.
type NetworkReportingApiEndpoint struct {
	Url       string `json:"url"`       // The URL of the endpoint to which reports may be delivered.
	GroupName string `json:"groupName"` // Name of the endpoint group.
}

// An object providing the result of a network resource load.
type NetworkLoadNetworkResourcePageResult struct {
	Success        bool                   `json:"success"`                  //
	NetError       float64                `json:"netError,omitempty"`       // Optional values used for error reporting.
	NetErrorName   string                 `json:"netErrorName,omitempty"`   //
	HttpStatusCode float64                `json:"httpStatusCode,omitempty"` //
	Stream         string                 `json:"stream,omitempty"`         // If successful, one of the following two fields holds the result.
	Headers        map[string]interface{} `json:"headers,omitempty"`        // Response headers.
}

// An options object that may be extended later to better support CORS, CORB and streaming.
type NetworkLoadNetworkResourceOptions struct {
	DisableCache       bool `json:"disableCache"`       //
	IncludeCredentials bool `json:"includeCredentials"` //
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

// Fired when HTTP request has failed to load.
type NetworkLoadingFailedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId       string                  `json:"requestId"`                 // Request identifier.
		Timestamp       float64                 `json:"timestamp"`                 // Timestamp.
		Type            string                  `json:"type"`                      // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		ErrorText       string                  `json:"errorText"`                 // Error message. List of network errors: https://cs.chromium.org/chromium/src/net/base/net_error_list.h
		Canceled        bool                    `json:"canceled,omitempty"`        // True if loading was canceled.
		BlockedReason   string                  `json:"blockedReason,omitempty"`   // The reason why loading was blocked, if any. enum values: other, csp, mixed-content, origin, inspector, subresource-filter, content-type, coep-frame-resource-needs-coep-header, coop-sandboxed-iframe-cannot-navigate-to-coop-page, corp-not-same-origin, corp-not-same-origin-after-defaulted-to-same-origin-by-coep, corp-not-same-origin-after-defaulted-to-same-origin-by-dip, corp-not-same-origin-after-defaulted-to-same-origin-by-coep-and-dip, corp-not-same-site
		CorsErrorStatus *NetworkCorsErrorStatus `json:"corsErrorStatus,omitempty"` // The reason why loading was blocked by CORS, if any.
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

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or mocked. Deprecated, use Fetch.requestPaused instead.
type NetworkRequestInterceptedEvent struct {
	Method string `json:"method"`
	Params struct {
		InterceptionId      string                 `json:"interceptionId"`                // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
		Request             *NetworkRequest        `json:"request"`                       //
		FrameId             string                 `json:"frameId"`                       // The id of the frame that initiated the request.
		ResourceType        string                 `json:"resourceType"`                  // How the requested resource will be used. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		IsNavigationRequest bool                   `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
		IsDownload          bool                   `json:"isDownload,omitempty"`          // Set if the request is a navigation that will result in a download. Only present after response is received from the server (i.e. HeadersReceived stage).
		RedirectUrl         string                 `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
		AuthChallenge       *NetworkAuthChallenge  `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
		ResponseErrorReason string                 `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
		ResponseStatusCode  int                    `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
		ResponseHeaders     map[string]interface{} `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
		RequestId           string                 `json:"requestId,omitempty"`           // If the intercepted request had a corresponding requestWillBeSent event fired for it, then this requestId will be the same as the requestId present in the requestWillBeSent event.
	} `json:"Params,omitempty"`
}

// Fired if request ended up loading from cache.
type NetworkRequestServedFromCacheEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string `json:"requestId"` // Request identifier.
	} `json:"Params,omitempty"`
}

// Fired when page is about to send HTTP request.
type NetworkRequestWillBeSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId            string            `json:"requestId"`                  // Request identifier.
		LoaderId             string            `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
		DocumentURL          string            `json:"documentURL"`                // URL of the document this request is loaded for.
		Request              *NetworkRequest   `json:"request"`                    // Request data.
		Timestamp            float64           `json:"timestamp"`                  // Timestamp.
		WallTime             float64           `json:"wallTime"`                   // Timestamp.
		Initiator            *NetworkInitiator `json:"initiator"`                  // Request initiator.
		RedirectHasExtraInfo bool              `json:"redirectHasExtraInfo"`       // In the case that redirectResponse is populated, this flag indicates whether requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be or were emitted for the request which was just redirected.
		RedirectResponse     *NetworkResponse  `json:"redirectResponse,omitempty"` // Redirect response data.
		Type                 string            `json:"type,omitempty"`             // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		FrameId              string            `json:"frameId,omitempty"`          // Frame identifier.
		HasUserGesture       bool              `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
	} `json:"Params,omitempty"`
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

// Fired when a signed exchange was received over the network
type NetworkSignedExchangeReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                     `json:"requestId"` // Request identifier.
		Info      *NetworkSignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
	} `json:"Params,omitempty"`
}

// Fired when HTTP response is available.
type NetworkResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId    string           `json:"requestId"`         // Request identifier.
		LoaderId     string           `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
		Timestamp    float64          `json:"timestamp"`         // Timestamp.
		Type         string           `json:"type"`              // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Preflight, Other
		Response     *NetworkResponse `json:"response"`          // Response data.
		HasExtraInfo bool             `json:"hasExtraInfo"`      // Indicates whether requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be or were emitted for this request.
		FrameId      string           `json:"frameId,omitempty"` // Frame identifier.
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

// Fired upon WebSocket creation.
type NetworkWebSocketCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string            `json:"requestId"`           // Request identifier.
		Url       string            `json:"url"`                 // WebSocket request URL.
		Initiator *NetworkInitiator `json:"initiator,omitempty"` // Request initiator.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket message error occurs.
type NetworkWebSocketFrameErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId    string  `json:"requestId"`    // Request identifier.
		Timestamp    float64 `json:"timestamp"`    // Timestamp.
		ErrorMessage string  `json:"errorMessage"` // WebSocket error message.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket message is received.
type NetworkWebSocketFrameReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket message is sent.
type NetworkWebSocketFrameSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
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

// Fired upon WebTransport creation.
type NetworkWebTransportCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		TransportId string            `json:"transportId"`         // WebTransport identifier.
		Url         string            `json:"url"`                 // WebTransport request URL.
		Timestamp   float64           `json:"timestamp"`           // Timestamp.
		Initiator   *NetworkInitiator `json:"initiator,omitempty"` // Request initiator.
	} `json:"Params,omitempty"`
}

// Fired when WebTransport handshake is finished.
type NetworkWebTransportConnectionEstablishedEvent struct {
	Method string `json:"method"`
	Params struct {
		TransportId string  `json:"transportId"` // WebTransport identifier.
		Timestamp   float64 `json:"timestamp"`   // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when WebTransport is disposed.
type NetworkWebTransportClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		TransportId string  `json:"transportId"` // WebTransport identifier.
		Timestamp   float64 `json:"timestamp"`   // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when additional information about a requestWillBeSent event is available from the network stack. Not every requestWillBeSent event will have an additional requestWillBeSentExtraInfo fired for it, and there is no guarantee whether requestWillBeSent or requestWillBeSentExtraInfo will be fired first for the same request.
type NetworkRequestWillBeSentExtraInfoEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId                     string                      `json:"requestId"`                               // Request identifier. Used to match this information to an existing requestWillBeSent event.
		AssociatedCookies             []*NetworkAssociatedCookie  `json:"associatedCookies"`                       // A list of cookies potentially associated to the requested URL. This includes both cookies sent with the request and the ones not sent; the latter are distinguished by having blockedReasons field set.
		Headers                       map[string]interface{}      `json:"headers"`                                 // Raw request headers as they will be sent over the wire.
		ConnectTiming                 *NetworkConnectTiming       `json:"connectTiming"`                           // Connection timing information for the request.
		ClientSecurityState           *NetworkClientSecurityState `json:"clientSecurityState,omitempty"`           // The client security state set for the request.
		SiteHasCookieInOtherPartition bool                        `json:"siteHasCookieInOtherPartition,omitempty"` // Whether the site has partitioned cookies stored in a partition different than the current one.
	} `json:"Params,omitempty"`
}

// Fired when additional information about a responseReceived event is available from the network stack. Not every responseReceived event will have an additional responseReceivedExtraInfo for it, and responseReceivedExtraInfo may be fired before or after responseReceived.
type NetworkResponseReceivedExtraInfoEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId                string                                `json:"requestId"`                          // Request identifier. Used to match this information to another responseReceived event.
		BlockedCookies           []*NetworkBlockedSetCookieWithReason  `json:"blockedCookies"`                     // A list of cookies which were not stored from the response along with the corresponding reasons for blocking. The cookies here may not be valid due to syntax errors, which are represented by the invalid cookie line string instead of a proper cookie.
		Headers                  map[string]interface{}                `json:"headers"`                            // Raw response headers as they were received over the wire.
		ResourceIPAddressSpace   string                                `json:"resourceIPAddressSpace"`             // The IP address space of the resource. The address space can only be determined once the transport established the connection, so we can't send it in `requestWillBeSentExtraInfo`. enum values: Local, Private, Public, Unknown
		StatusCode               int                                   `json:"statusCode"`                         // The status code of the response. This is useful in cases the request failed and no responseReceived event is triggered, which is the case for, e.g., CORS errors. This is also the correct status code for cached requests, where the status in responseReceived is a 200 and this will be 304.
		HeadersText              string                                `json:"headersText,omitempty"`              // Raw response header text as it was received over the wire. The raw text may not always be available, such as in the case of HTTP/2 or QUIC.
		CookiePartitionKey       *NetworkCookiePartitionKey            `json:"cookiePartitionKey,omitempty"`       // The cookie partition key that will be used to store partitioned cookies set in this response. Only sent when partitioned cookies are enabled.
		CookiePartitionKeyOpaque bool                                  `json:"cookiePartitionKeyOpaque,omitempty"` // True if partitioned cookies are enabled, but the partition key is not serializable to string.
		ExemptedCookies          []*NetworkExemptedSetCookieWithReason `json:"exemptedCookies,omitempty"`          // A list of cookies which should have been blocked by 3PCD but are exempted and stored from the response with the corresponding reason.
	} `json:"Params,omitempty"`
}

// Fired when 103 Early Hints headers is received in addition to the common response. Not every responseReceived event will have an responseReceivedEarlyHints fired. Only one responseReceivedEarlyHints may be fired for eached responseReceived event.
type NetworkResponseReceivedEarlyHintsEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier. Used to match this information to another responseReceived event.
		Headers   map[string]interface{} `json:"headers"`   // Raw response headers as they were received over the wire.
	} `json:"Params,omitempty"`
}

// Fired exactly once for each Trust Token operation. Depending on the type of the operation and whether the operation succeeded or failed, the event is fired before the corresponding request was sent or after the response was received.
type NetworkTrustTokenOperationDoneEvent struct {
	Method string `json:"method"`
	Params struct {
		Status           string `json:"status"`                     // Detailed success or error status of the operation. 'AlreadyExists' also signifies a successful operation, as the result of the operation already exists und thus, the operation was abort preemptively (e.g. a cache hit).
		Type             string `json:"type"`                       //  enum values: Issuance, Redemption, Signing
		RequestId        string `json:"requestId"`                  //
		TopLevelOrigin   string `json:"topLevelOrigin,omitempty"`   // Top level origin. The context in which the operation was attempted.
		IssuerOrigin     string `json:"issuerOrigin,omitempty"`     // Origin of the issuer in case of a "Issuance" or "Redemption" operation.
		IssuedTokenCount int    `json:"issuedTokenCount,omitempty"` // The number of obtained Trust Tokens on a successful "Issuance" operation.
	} `json:"Params,omitempty"`
}

// Fired once when parsing the .wbn file has succeeded. The event contains the information about the web bundle contents.
type NetworkSubresourceWebBundleMetadataReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string   `json:"requestId"` // Request identifier. Used to match this information to another event.
		Urls      []string `json:"urls"`      // A list of URLs of resources in the subresource Web Bundle.
	} `json:"Params,omitempty"`
}

// Fired once when parsing the .wbn file has failed.
type NetworkSubresourceWebBundleMetadataErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId    string `json:"requestId"`    // Request identifier. Used to match this information to another event.
		ErrorMessage string `json:"errorMessage"` // Error message
	} `json:"Params,omitempty"`
}

// Fired when handling requests for resources within a .wbn file. Note: this will only be fired for resources that are requested by the webpage.
type NetworkSubresourceWebBundleInnerResponseParsedEvent struct {
	Method string `json:"method"`
	Params struct {
		InnerRequestId  string `json:"innerRequestId"`            // Request identifier of the subresource request
		InnerRequestURL string `json:"innerRequestURL"`           // URL of the subresource resource.
		BundleRequestId string `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
	} `json:"Params,omitempty"`
}

// Fired when request for resources within a .wbn file failed.
type NetworkSubresourceWebBundleInnerResponseErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		InnerRequestId  string `json:"innerRequestId"`            // Request identifier of the subresource request
		InnerRequestURL string `json:"innerRequestURL"`           // URL of the subresource resource.
		ErrorMessage    string `json:"errorMessage"`              // Error message
		BundleRequestId string `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
	} `json:"Params,omitempty"`
}

// Is sent whenever a new report is added. And after 'enableReportingApi' for all existing reports.
type NetworkReportingApiReportAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Report *NetworkReportingApiReport `json:"report"` //
	} `json:"Params,omitempty"`
}

type NetworkReportingApiReportUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Report *NetworkReportingApiReport `json:"report"` //
	} `json:"Params,omitempty"`
}

type NetworkReportingApiEndpointsChangedForOriginEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin    string                         `json:"origin"`    // Origin of the document(s) which configured the endpoints.
		Endpoints []*NetworkReportingApiEndpoint `json:"endpoints"` //
	} `json:"Params,omitempty"`
}

type Network struct {
	target gcdmessage.ChromeTargeter
}

func NewNetwork(target gcdmessage.ChromeTargeter) *Network {
	c := &Network{target: target}
	return c
}

type NetworkSetAcceptedEncodingsParams struct {
	// List of accepted content encodings. enum values: deflate, gzip, br, zstd
	Encodings []string `json:"encodings"`
}

// SetAcceptedEncodingsWithParams - Sets a list of content encodings that will be accepted. Empty list means no encoding is accepted.
func (c *Network) SetAcceptedEncodingsWithParams(ctx context.Context, v *NetworkSetAcceptedEncodingsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setAcceptedEncodings", Params: v})
}

// SetAcceptedEncodings - Sets a list of content encodings that will be accepted. Empty list means no encoding is accepted.
// encodings - List of accepted content encodings. enum values: deflate, gzip, br, zstd
func (c *Network) SetAcceptedEncodings(ctx context.Context, encodings []string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetAcceptedEncodingsParams
	v.Encodings = encodings
	return c.SetAcceptedEncodingsWithParams(ctx, &v)
}

// Clears accepted encodings set by setAcceptedEncodings
func (c *Network) ClearAcceptedEncodingsOverride(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearAcceptedEncodingsOverride"})
}

// CanClearBrowserCache - Tells whether clearing browser cache is supported.
// Returns -  result - True if browser cache can be cleared.
func (c *Network) CanClearBrowserCache(ctx context.Context) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCache"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// CanClearBrowserCookies - Tells whether clearing browser cookies is supported.
// Returns -  result - True if browser cookies can be cleared.
func (c *Network) CanClearBrowserCookies(ctx context.Context) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCookies"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// CanEmulateNetworkConditions - Tells whether emulation of network conditions is supported.
// Returns -  result - True if emulation of network conditions is supported.
func (c *Network) CanEmulateNetworkConditions(ctx context.Context) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canEmulateNetworkConditions"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// Clears browser cache.
func (c *Network) ClearBrowserCache(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCache"})
}

// Clears browser cookies.
func (c *Network) ClearBrowserCookies(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCookies"})
}

type NetworkContinueInterceptedRequestParams struct {
	//
	InterceptionId string `json:"interceptionId"`
	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
	ErrorReason string `json:"errorReason,omitempty"`
	// If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
	Url string `json:"url,omitempty"`
	// If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
	Method string `json:"method,omitempty"`
	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData string `json:"postData,omitempty"`
	// If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse *NetworkAuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

// ContinueInterceptedRequestWithParams - Response to Network.requestIntercepted which either modifies the request to continue with any modifications, or blocks it, or completes it with the provided response bytes. If a network fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted event will be sent with the same InterceptionId. Deprecated, use Fetch.continueRequest, Fetch.fulfillRequest and Fetch.failRequest instead.
func (c *Network) ContinueInterceptedRequestWithParams(ctx context.Context, v *NetworkContinueInterceptedRequestParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.continueInterceptedRequest", Params: v})
}

// ContinueInterceptedRequest - Response to Network.requestIntercepted which either modifies the request to continue with any modifications, or blocks it, or completes it with the provided response bytes. If a network fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted event will be sent with the same InterceptionId. Deprecated, use Fetch.continueRequest, Fetch.fulfillRequest and Fetch.failRequest instead.
// interceptionId -
// errorReason - If set this causes the request to fail with the given reason. Passing `Aborted` for requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
// url - If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
// method - If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
// postData - If set this allows postData to be set. Must not be set in response to an authChallenge.
// headers - If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
// authChallengeResponse - Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
func (c *Network) ContinueInterceptedRequest(ctx context.Context, interceptionId string, errorReason string, url string, method string, postData string, headers map[string]interface{}, authChallengeResponse *NetworkAuthChallengeResponse) (*gcdmessage.ChromeResponse, error) {
	var v NetworkContinueInterceptedRequestParams
	v.InterceptionId = interceptionId
	v.ErrorReason = errorReason
	v.Url = url
	v.Method = method
	v.PostData = postData
	v.Headers = headers
	v.AuthChallengeResponse = authChallengeResponse
	return c.ContinueInterceptedRequestWithParams(ctx, &v)
}

type NetworkDeleteCookiesParams struct {
	// Name of the cookies to remove.
	Name string `json:"name"`
	// If specified, deletes all the cookies with the given name where domain and path match provided URL.
	Url string `json:"url,omitempty"`
	// If specified, deletes only cookies with the exact domain.
	Domain string `json:"domain,omitempty"`
	// If specified, deletes only cookies with the exact path.
	Path string `json:"path,omitempty"`
	// If specified, deletes only cookies with the the given name and partitionKey where all partition key attributes match the cookie partition key attribute.
	PartitionKey *NetworkCookiePartitionKey `json:"partitionKey,omitempty"`
}

// DeleteCookiesWithParams - Deletes browser cookies with matching name and url or domain/path/partitionKey pair.
func (c *Network) DeleteCookiesWithParams(ctx context.Context, v *NetworkDeleteCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.deleteCookies", Params: v})
}

// DeleteCookies - Deletes browser cookies with matching name and url or domain/path/partitionKey pair.
// name - Name of the cookies to remove.
// url - If specified, deletes all the cookies with the given name where domain and path match provided URL.
// domain - If specified, deletes only cookies with the exact domain.
// path - If specified, deletes only cookies with the exact path.
// partitionKey - If specified, deletes only cookies with the the given name and partitionKey where all partition key attributes match the cookie partition key attribute.
func (c *Network) DeleteCookies(ctx context.Context, name string, url string, domain string, path string, partitionKey *NetworkCookiePartitionKey) (*gcdmessage.ChromeResponse, error) {
	var v NetworkDeleteCookiesParams
	v.Name = name
	v.Url = url
	v.Domain = domain
	v.Path = path
	v.PartitionKey = partitionKey
	return c.DeleteCookiesWithParams(ctx, &v)
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *Network) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.disable"})
}

type NetworkEmulateNetworkConditionsParams struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`
	// Minimum latency from request sent to response headers received (ms).
	Latency float64 `json:"latency"`
	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput float64 `json:"downloadThroughput"`
	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	UploadThroughput float64 `json:"uploadThroughput"`
	// Connection type if known. enum values: none, cellular2g, cellular3g, cellular4g, bluetooth, ethernet, wifi, wimax, other
	ConnectionType string `json:"connectionType,omitempty"`
	// WebRTC packet loss (percent, 0-100). 0 disables packet loss emulation, 100 drops all the packets.
	PacketLoss float64 `json:"packetLoss,omitempty"`
	// WebRTC packet queue length (packet). 0 removes any queue length limitations.
	PacketQueueLength int `json:"packetQueueLength,omitempty"`
	// WebRTC packetReordering feature.
	PacketReordering bool `json:"packetReordering,omitempty"`
}

// EmulateNetworkConditionsWithParams - Activates emulation of network conditions.
func (c *Network) EmulateNetworkConditionsWithParams(ctx context.Context, v *NetworkEmulateNetworkConditionsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.emulateNetworkConditions", Params: v})
}

// EmulateNetworkConditions - Activates emulation of network conditions.
// offline - True to emulate internet disconnection.
// latency - Minimum latency from request sent to response headers received (ms).
// downloadThroughput - Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
// uploadThroughput - Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
// connectionType - Connection type if known. enum values: none, cellular2g, cellular3g, cellular4g, bluetooth, ethernet, wifi, wimax, other
// packetLoss - WebRTC packet loss (percent, 0-100). 0 disables packet loss emulation, 100 drops all the packets.
// packetQueueLength - WebRTC packet queue length (packet). 0 removes any queue length limitations.
// packetReordering - WebRTC packetReordering feature.
func (c *Network) EmulateNetworkConditions(ctx context.Context, offline bool, latency float64, downloadThroughput float64, uploadThroughput float64, connectionType string, packetLoss float64, packetQueueLength int, packetReordering bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEmulateNetworkConditionsParams
	v.Offline = offline
	v.Latency = latency
	v.DownloadThroughput = downloadThroughput
	v.UploadThroughput = uploadThroughput
	v.ConnectionType = connectionType
	v.PacketLoss = packetLoss
	v.PacketQueueLength = packetQueueLength
	v.PacketReordering = packetReordering
	return c.EmulateNetworkConditionsWithParams(ctx, &v)
}

type NetworkEnableParams struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize int `json:"maxPostDataSize,omitempty"`
}

// EnableWithParams - Enables network tracking, network events will now be delivered to the client.
func (c *Network) EnableWithParams(ctx context.Context, v *NetworkEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enable", Params: v})
}

// Enable - Enables network tracking, network events will now be delivered to the client.
// maxTotalBufferSize - Buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxResourceBufferSize - Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxPostDataSize - Longest post body size (in bytes) that would be included in requestWillBeSent notification
func (c *Network) Enable(ctx context.Context, maxTotalBufferSize int, maxResourceBufferSize int, maxPostDataSize int) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEnableParams
	v.MaxTotalBufferSize = maxTotalBufferSize
	v.MaxResourceBufferSize = maxResourceBufferSize
	v.MaxPostDataSize = maxPostDataSize
	return c.EnableWithParams(ctx, &v)
}

// GetAllCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the `cookies` field. Deprecated. Use Storage.getCookies instead.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetAllCookies(ctx context.Context) ([]*NetworkCookie, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getAllCookies"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Cookies []*NetworkCookie
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Cookies, nil
}

type NetworkGetCertificateParams struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

// GetCertificateWithParams - Returns the DER-encoded certificate.
// Returns -  tableNames -
func (c *Network) GetCertificateWithParams(ctx context.Context, v *NetworkGetCertificateParams) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCertificate", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			TableNames []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.TableNames, nil
}

// GetCertificate - Returns the DER-encoded certificate.
// origin - Origin to get certificate for.
// Returns -  tableNames -
func (c *Network) GetCertificate(ctx context.Context, origin string) ([]string, error) {
	var v NetworkGetCertificateParams
	v.Origin = origin
	return c.GetCertificateWithParams(ctx, &v)
}

type NetworkGetCookiesParams struct {
	// The list of URLs for which applicable cookies will be fetched. If not specified, it's assumed to be set to the list containing the URLs of the page and all of its subframes.
	Urls []string `json:"urls,omitempty"`
}

// GetCookiesWithParams - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the `cookies` field.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookiesWithParams(ctx context.Context, v *NetworkGetCookiesParams) ([]*NetworkCookie, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCookies", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Cookies []*NetworkCookie
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Cookies, nil
}

// GetCookies - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the `cookies` field.
// urls - The list of URLs for which applicable cookies will be fetched. If not specified, it's assumed to be set to the list containing the URLs of the page and all of its subframes.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookies(ctx context.Context, urls []string) ([]*NetworkCookie, error) {
	var v NetworkGetCookiesParams
	v.Urls = urls
	return c.GetCookiesWithParams(ctx, &v)
}

type NetworkGetResponseBodyParams struct {
	// Identifier of the network request to get content for.
	RequestId string `json:"requestId"`
}

// GetResponseBodyWithParams - Returns content served for the given request.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyWithParams(ctx context.Context, v *NetworkGetResponseBodyParams) (string, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBody", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	if chromeData.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// GetResponseBody - Returns content served for the given request.
// requestId - Identifier of the network request to get content for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBody(ctx context.Context, requestId string) (string, bool, error) {
	var v NetworkGetResponseBodyParams
	v.RequestId = requestId
	return c.GetResponseBodyWithParams(ctx, &v)
}

type NetworkGetRequestPostDataParams struct {
	// Identifier of the network request to get content for.
	RequestId string `json:"requestId"`
}

// GetRequestPostDataWithParams - Returns post data sent with the request. Returns an error when no data was sent with the request.
// Returns -  postData - Request body string, omitting files from multipart requests
func (c *Network) GetRequestPostDataWithParams(ctx context.Context, v *NetworkGetRequestPostDataParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getRequestPostData", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			PostData string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.PostData, nil
}

// GetRequestPostData - Returns post data sent with the request. Returns an error when no data was sent with the request.
// requestId - Identifier of the network request to get content for.
// Returns -  postData - Request body string, omitting files from multipart requests
func (c *Network) GetRequestPostData(ctx context.Context, requestId string) (string, error) {
	var v NetworkGetRequestPostDataParams
	v.RequestId = requestId
	return c.GetRequestPostDataWithParams(ctx, &v)
}

type NetworkGetResponseBodyForInterceptionParams struct {
	// Identifier for the intercepted request to get body for.
	InterceptionId string `json:"interceptionId"`
}

// GetResponseBodyForInterceptionWithParams - Returns content served for the given currently intercepted request.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyForInterceptionWithParams(ctx context.Context, v *NetworkGetResponseBodyForInterceptionParams) (string, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBodyForInterception", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	if chromeData.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// GetResponseBodyForInterception - Returns content served for the given currently intercepted request.
// interceptionId - Identifier for the intercepted request to get body for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyForInterception(ctx context.Context, interceptionId string) (string, bool, error) {
	var v NetworkGetResponseBodyForInterceptionParams
	v.InterceptionId = interceptionId
	return c.GetResponseBodyForInterceptionWithParams(ctx, &v)
}

type NetworkTakeResponseBodyForInterceptionAsStreamParams struct {
	//
	InterceptionId string `json:"interceptionId"`
}

// TakeResponseBodyForInterceptionAsStreamWithParams - Returns a handle to the stream representing the response body. Note that after this command, the intercepted request can't be continued as is -- you either need to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified.
// Returns -  stream -
func (c *Network) TakeResponseBodyForInterceptionAsStreamWithParams(ctx context.Context, v *NetworkTakeResponseBodyForInterceptionAsStreamParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.takeResponseBodyForInterceptionAsStream", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Stream string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Stream, nil
}

// TakeResponseBodyForInterceptionAsStream - Returns a handle to the stream representing the response body. Note that after this command, the intercepted request can't be continued as is -- you either need to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified.
// interceptionId -
// Returns -  stream -
func (c *Network) TakeResponseBodyForInterceptionAsStream(ctx context.Context, interceptionId string) (string, error) {
	var v NetworkTakeResponseBodyForInterceptionAsStreamParams
	v.InterceptionId = interceptionId
	return c.TakeResponseBodyForInterceptionAsStreamWithParams(ctx, &v)
}

type NetworkReplayXHRParams struct {
	// Identifier of XHR to replay.
	RequestId string `json:"requestId"`
}

// ReplayXHRWithParams - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
func (c *Network) ReplayXHRWithParams(ctx context.Context, v *NetworkReplayXHRParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.replayXHR", Params: v})
}

// ReplayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *Network) ReplayXHR(ctx context.Context, requestId string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkReplayXHRParams
	v.RequestId = requestId
	return c.ReplayXHRWithParams(ctx, &v)
}

type NetworkSearchInResponseBodyParams struct {
	// Identifier of the network response to search.
	RequestId string `json:"requestId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInResponseBodyWithParams - Searches for given string in response content.
// Returns -  result - List of search matches.
func (c *Network) SearchInResponseBodyWithParams(ctx context.Context, v *NetworkSearchInResponseBodyParams) ([]*DebuggerSearchMatch, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.searchInResponseBody", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Result []*DebuggerSearchMatch
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Result, nil
}

// SearchInResponseBody - Searches for given string in response content.
// requestId - Identifier of the network response to search.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Network) SearchInResponseBody(ctx context.Context, requestId string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v NetworkSearchInResponseBodyParams
	v.RequestId = requestId
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInResponseBodyWithParams(ctx, &v)
}

type NetworkSetBlockedURLsParams struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls []string `json:"urls"`
}

// SetBlockedURLsWithParams - Blocks URLs from loading.
func (c *Network) SetBlockedURLsWithParams(ctx context.Context, v *NetworkSetBlockedURLsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBlockedURLs", Params: v})
}

// SetBlockedURLs - Blocks URLs from loading.
// urls - URL patterns to block. Wildcards ('*') are allowed.
func (c *Network) SetBlockedURLs(ctx context.Context, urls []string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetBlockedURLsParams
	v.Urls = urls
	return c.SetBlockedURLsWithParams(ctx, &v)
}

type NetworkSetBypassServiceWorkerParams struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

// SetBypassServiceWorkerWithParams - Toggles ignoring of service worker for each request.
func (c *Network) SetBypassServiceWorkerWithParams(ctx context.Context, v *NetworkSetBypassServiceWorkerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBypassServiceWorker", Params: v})
}

// SetBypassServiceWorker - Toggles ignoring of service worker for each request.
// bypass - Bypass service worker and load from network.
func (c *Network) SetBypassServiceWorker(ctx context.Context, bypass bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetBypassServiceWorkerParams
	v.Bypass = bypass
	return c.SetBypassServiceWorkerWithParams(ctx, &v)
}

type NetworkSetCacheDisabledParams struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// SetCacheDisabledWithParams - Toggles ignoring cache for each request. If `true`, cache will not be used.
func (c *Network) SetCacheDisabledWithParams(ctx context.Context, v *NetworkSetCacheDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCacheDisabled", Params: v})
}

// SetCacheDisabled - Toggles ignoring cache for each request. If `true`, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *Network) SetCacheDisabled(ctx context.Context, cacheDisabled bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetCacheDisabledParams
	v.CacheDisabled = cacheDisabled
	return c.SetCacheDisabledWithParams(ctx, &v)
}

type NetworkSetCookieParams struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
	// The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
	Url string `json:"url,omitempty"`
	// Cookie domain.
	Domain string `json:"domain,omitempty"`
	// Cookie path.
	Path string `json:"path,omitempty"`
	// True if cookie is secure.
	Secure bool `json:"secure,omitempty"`
	// True if cookie is http-only.
	HttpOnly bool `json:"httpOnly,omitempty"`
	// Cookie SameSite type. enum values: Strict, Lax, None
	SameSite string `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires float64 `json:"expires,omitempty"`
	// Cookie Priority type. enum values: Low, Medium, High
	Priority string `json:"priority,omitempty"`
	// True if cookie is SameParty.
	SameParty bool `json:"sameParty,omitempty"`
	// Cookie source scheme type. enum values: Unset, NonSecure, Secure
	SourceScheme string `json:"sourceScheme,omitempty"`
	// Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
	SourcePort int `json:"sourcePort,omitempty"`
	// Cookie partition key. If not set, the cookie will be set as not partitioned.
	PartitionKey *NetworkCookiePartitionKey `json:"partitionKey,omitempty"`
}

// SetCookieWithParams - Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
// Returns -  success - Always set to true. If an error occurs, the response indicates protocol error.
func (c *Network) SetCookieWithParams(ctx context.Context, v *NetworkSetCookieParams) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCookie", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Success, nil
}

// SetCookie - Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
// name - Cookie name.
// value - Cookie value.
// url - The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
// domain - Cookie domain.
// path - Cookie path.
// secure - True if cookie is secure.
// httpOnly - True if cookie is http-only.
// sameSite - Cookie SameSite type. enum values: Strict, Lax, None
// expires - Cookie expiration date, session cookie if not set
// priority - Cookie Priority type. enum values: Low, Medium, High
// sameParty - True if cookie is SameParty.
// sourceScheme - Cookie source scheme type. enum values: Unset, NonSecure, Secure
// sourcePort - Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
// partitionKey - Cookie partition key. If not set, the cookie will be set as not partitioned.
// Returns -  success - Always set to true. If an error occurs, the response indicates protocol error.
func (c *Network) SetCookie(ctx context.Context, name string, value string, url string, domain string, path string, secure bool, httpOnly bool, sameSite string, expires float64, priority string, sameParty bool, sourceScheme string, sourcePort int, partitionKey *NetworkCookiePartitionKey) (bool, error) {
	var v NetworkSetCookieParams
	v.Name = name
	v.Value = value
	v.Url = url
	v.Domain = domain
	v.Path = path
	v.Secure = secure
	v.HttpOnly = httpOnly
	v.SameSite = sameSite
	v.Expires = expires
	v.Priority = priority
	v.SameParty = sameParty
	v.SourceScheme = sourceScheme
	v.SourcePort = sourcePort
	v.PartitionKey = partitionKey
	return c.SetCookieWithParams(ctx, &v)
}

type NetworkSetCookiesParams struct {
	// Cookies to be set.
	Cookies []*NetworkCookieParam `json:"cookies"`
}

// SetCookiesWithParams - Sets given cookies.
func (c *Network) SetCookiesWithParams(ctx context.Context, v *NetworkSetCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCookies", Params: v})
}

// SetCookies - Sets given cookies.
// cookies - Cookies to be set.
func (c *Network) SetCookies(ctx context.Context, cookies []*NetworkCookieParam) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetCookiesParams
	v.Cookies = cookies
	return c.SetCookiesWithParams(ctx, &v)
}

type NetworkSetExtraHTTPHeadersParams struct {
	// Map with extra HTTP headers.
	Headers map[string]interface{} `json:"headers"`
}

// SetExtraHTTPHeadersWithParams - Specifies whether to always send extra HTTP headers with the requests from this page.
func (c *Network) SetExtraHTTPHeadersWithParams(ctx context.Context, v *NetworkSetExtraHTTPHeadersParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setExtraHTTPHeaders", Params: v})
}

// SetExtraHTTPHeaders - Specifies whether to always send extra HTTP headers with the requests from this page.
// headers - Map with extra HTTP headers.
func (c *Network) SetExtraHTTPHeaders(ctx context.Context, headers map[string]interface{}) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetExtraHTTPHeadersParams
	v.Headers = headers
	return c.SetExtraHTTPHeadersWithParams(ctx, &v)
}

type NetworkSetAttachDebugStackParams struct {
	// Whether to attach a page script stack for debugging purpose.
	Enabled bool `json:"enabled"`
}

// SetAttachDebugStackWithParams - Specifies whether to attach a page script stack id in requests
func (c *Network) SetAttachDebugStackWithParams(ctx context.Context, v *NetworkSetAttachDebugStackParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setAttachDebugStack", Params: v})
}

// SetAttachDebugStack - Specifies whether to attach a page script stack id in requests
// enabled - Whether to attach a page script stack for debugging purpose.
func (c *Network) SetAttachDebugStack(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetAttachDebugStackParams
	v.Enabled = enabled
	return c.SetAttachDebugStackWithParams(ctx, &v)
}

type NetworkSetRequestInterceptionParams struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
	Patterns []*NetworkRequestPattern `json:"patterns"`
}

// SetRequestInterceptionWithParams - Sets the requests to intercept that match the provided patterns and optionally resource types. Deprecated, please use Fetch.enable instead.
func (c *Network) SetRequestInterceptionWithParams(ctx context.Context, v *NetworkSetRequestInterceptionParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setRequestInterception", Params: v})
}

// SetRequestInterception - Sets the requests to intercept that match the provided patterns and optionally resource types. Deprecated, please use Fetch.enable instead.
// patterns - Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
func (c *Network) SetRequestInterception(ctx context.Context, patterns []*NetworkRequestPattern) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetRequestInterceptionParams
	v.Patterns = patterns
	return c.SetRequestInterceptionWithParams(ctx, &v)
}

type NetworkSetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
	// Browser language to emulate.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	// The platform navigator.platform should return.
	Platform string `json:"platform,omitempty"`
	// To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
	UserAgentMetadata *EmulationUserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

// SetUserAgentOverrideWithParams - Allows overriding user agent with the given string.
func (c *Network) SetUserAgentOverrideWithParams(ctx context.Context, v *NetworkSetUserAgentOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setUserAgentOverride", Params: v})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
// acceptLanguage - Browser language to emulate.
// platform - The platform navigator.platform should return.
// userAgentMetadata - To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
func (c *Network) SetUserAgentOverride(ctx context.Context, userAgent string, acceptLanguage string, platform string, userAgentMetadata *EmulationUserAgentMetadata) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetUserAgentOverrideParams
	v.UserAgent = userAgent
	v.AcceptLanguage = acceptLanguage
	v.Platform = platform
	v.UserAgentMetadata = userAgentMetadata
	return c.SetUserAgentOverrideWithParams(ctx, &v)
}

type NetworkStreamResourceContentParams struct {
	// Identifier of the request to stream.
	RequestId string `json:"requestId"`
}

// StreamResourceContentWithParams - Enables streaming of the response for the given requestId. If enabled, the dataReceived event contains the data that was received during streaming.
// Returns -
func (c *Network) StreamResourceContentWithParams(ctx context.Context, v *NetworkStreamResourceContentParams) error {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.streamResourceContent", Params: v})
	if err != nil {
		return err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
		}
	}

	if resp == nil {
		return &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return err
	}

	if chromeData.Error != nil {
		return &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return nil
}

// StreamResourceContent - Enables streaming of the response for the given requestId. If enabled, the dataReceived event contains the data that was received during streaming.
// requestId - Identifier of the request to stream.
// Returns -
func (c *Network) StreamResourceContent(ctx context.Context, requestId string) error {
	var v NetworkStreamResourceContentParams
	v.RequestId = requestId
	return c.StreamResourceContentWithParams(ctx, &v)
}

type NetworkGetSecurityIsolationStatusParams struct {
	// If no frameId is provided, the status of the target is provided.
	FrameId string `json:"frameId,omitempty"`
}

// GetSecurityIsolationStatusWithParams - Returns information about the COEP/COOP isolation status.
// Returns -  status -
func (c *Network) GetSecurityIsolationStatusWithParams(ctx context.Context, v *NetworkGetSecurityIsolationStatusParams) (*NetworkSecurityIsolationStatus, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getSecurityIsolationStatus", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Status *NetworkSecurityIsolationStatus
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Status, nil
}

// GetSecurityIsolationStatus - Returns information about the COEP/COOP isolation status.
// frameId - If no frameId is provided, the status of the target is provided.
// Returns -  status -
func (c *Network) GetSecurityIsolationStatus(ctx context.Context, frameId string) (*NetworkSecurityIsolationStatus, error) {
	var v NetworkGetSecurityIsolationStatusParams
	v.FrameId = frameId
	return c.GetSecurityIsolationStatusWithParams(ctx, &v)
}

type NetworkEnableReportingApiParams struct {
	// Whether to enable or disable events for the Reporting API
	Enable bool `json:"enable"`
}

// EnableReportingApiWithParams - Enables tracking for the Reporting API, events generated by the Reporting API will now be delivered to the client. Enabling triggers 'reportingApiReportAdded' for all existing reports.
func (c *Network) EnableReportingApiWithParams(ctx context.Context, v *NetworkEnableReportingApiParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enableReportingApi", Params: v})
}

// EnableReportingApi - Enables tracking for the Reporting API, events generated by the Reporting API will now be delivered to the client. Enabling triggers 'reportingApiReportAdded' for all existing reports.
// enable - Whether to enable or disable events for the Reporting API
func (c *Network) EnableReportingApi(ctx context.Context, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEnableReportingApiParams
	v.Enable = enable
	return c.EnableReportingApiWithParams(ctx, &v)
}

type NetworkLoadNetworkResourceParams struct {
	// Frame id to get the resource for. Mandatory for frame targets, and should be omitted for worker targets.
	FrameId string `json:"frameId,omitempty"`
	// URL of the resource to get content for.
	Url string `json:"url"`
	// Options for the request.
	Options *NetworkLoadNetworkResourceOptions `json:"options"`
}

// LoadNetworkResourceWithParams - Fetches the resource and returns the content.
// Returns -  resource -
func (c *Network) LoadNetworkResourceWithParams(ctx context.Context, v *NetworkLoadNetworkResourceParams) (*NetworkLoadNetworkResourcePageResult, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.loadNetworkResource", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Resource *NetworkLoadNetworkResourcePageResult
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Resource, nil
}

// LoadNetworkResource - Fetches the resource and returns the content.
// frameId - Frame id to get the resource for. Mandatory for frame targets, and should be omitted for worker targets.
// url - URL of the resource to get content for.
// options - Options for the request.
// Returns -  resource -
func (c *Network) LoadNetworkResource(ctx context.Context, frameId string, url string, options *NetworkLoadNetworkResourceOptions) (*NetworkLoadNetworkResourcePageResult, error) {
	var v NetworkLoadNetworkResourceParams
	v.FrameId = frameId
	v.Url = url
	v.Options = options
	return c.LoadNetworkResourceWithParams(ctx, &v)
}
