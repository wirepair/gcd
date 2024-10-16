// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Audits functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Information about a cookie that is affected by an inspector issue.
type AuditsAffectedCookie struct {
	Name   string `json:"name"`   // The following three properties uniquely identify a cookie
	Path   string `json:"path"`   //
	Domain string `json:"domain"` //
}

// Information about a request that is affected by an inspector issue.
type AuditsAffectedRequest struct {
	RequestId string `json:"requestId"`     // The unique request id.
	Url       string `json:"url,omitempty"` //
}

// Information about the frame affected by an inspector issue.
type AuditsAffectedFrame struct {
	FrameId string `json:"frameId"` //
}

// This information is currently necessary, as the front-end has a difficult time finding a specific cookie. With this, we can convey specific error information without the cookie.
type AuditsCookieIssueDetails struct {
	Cookie                 *AuditsAffectedCookie  `json:"cookie,omitempty"`         // If AffectedCookie is not set then rawCookieLine contains the raw Set-Cookie header string. This hints at a problem where the cookie line is syntactically or semantically malformed in a way that no valid cookie could be created.
	RawCookieLine          string                 `json:"rawCookieLine,omitempty"`  //
	CookieWarningReasons   []string               `json:"cookieWarningReasons"`     //  enum values: WarnSameSiteUnspecifiedCrossSiteContext, WarnSameSiteNoneInsecure, WarnSameSiteUnspecifiedLaxAllowUnsafe, WarnSameSiteStrictLaxDowngradeStrict, WarnSameSiteStrictCrossDowngradeStrict, WarnSameSiteStrictCrossDowngradeLax, WarnSameSiteLaxCrossDowngradeStrict, WarnSameSiteLaxCrossDowngradeLax, WarnAttributeValueExceedsMaxSize, WarnDomainNonASCII
	CookieExclusionReasons []string               `json:"cookieExclusionReasons"`   //  enum values: ExcludeSameSiteUnspecifiedTreatedAsLax, ExcludeSameSiteNoneInsecure, ExcludeSameSiteLax, ExcludeSameSiteStrict, ExcludeInvalidSameParty, ExcludeSamePartyCrossPartyContext, ExcludeDomainNonASCII, ExcludeThirdPartyCookieBlockedInFirstPartySet
	Operation              string                 `json:"operation"`                // Optionally identifies the site-for-cookies and the cookie url, which may be used by the front-end as additional context. enum values: SetCookie, ReadCookie
	SiteForCookies         string                 `json:"siteForCookies,omitempty"` //
	CookieUrl              string                 `json:"cookieUrl,omitempty"`      //
	Request                *AuditsAffectedRequest `json:"request,omitempty"`        //
}

// No Description.
type AuditsMixedContentIssueDetails struct {
	ResourceType     string                 `json:"resourceType,omitempty"` // The type of resource causing the mixed content issue (css, js, iframe, form,...). Marked as optional because it is mapped to from blink::mojom::RequestContextType, which will be replaced by network::mojom::RequestDestination enum values: AttributionSrc, Audio, Beacon, CSPReport, Download, EventSource, Favicon, Font, Form, Frame, Image, Import, Manifest, Ping, PluginData, PluginResource, Prefetch, Resource, Script, ServiceWorker, SharedWorker, Stylesheet, Track, Video, Worker, XMLHttpRequest, XSLT
	ResolutionStatus string                 `json:"resolutionStatus"`       // The way the mixed content issue is being resolved. enum values: MixedContentBlocked, MixedContentAutomaticallyUpgraded, MixedContentWarning
	InsecureURL      string                 `json:"insecureURL"`            // The unsafe http url causing the mixed content issue.
	MainResourceURL  string                 `json:"mainResourceURL"`        // The url responsible for the call to an unsafe url.
	Request          *AuditsAffectedRequest `json:"request,omitempty"`      // The mixed content request. Does not always exist (e.g. for unsafe form submission urls).
	Frame            *AuditsAffectedFrame   `json:"frame,omitempty"`        // Optional because not every mixed content issue is necessarily linked to a frame.
}

// Details for a request that has been blocked with the BLOCKED_BY_RESPONSE code. Currently only used for COEP/COOP, but may be extended to include some CSP errors in the future.
type AuditsBlockedByResponseIssueDetails struct {
	Request      *AuditsAffectedRequest `json:"request"`                //
	ParentFrame  *AuditsAffectedFrame   `json:"parentFrame,omitempty"`  //
	BlockedFrame *AuditsAffectedFrame   `json:"blockedFrame,omitempty"` //
	Reason       string                 `json:"reason"`                 //  enum values: CoepFrameResourceNeedsCoepHeader, CoopSandboxedIFrameCannotNavigateToCoopPage, CorpNotSameOrigin, CorpNotSameOriginAfterDefaultedToSameOriginByCoep, CorpNotSameSite
}

// No Description.
type AuditsHeavyAdIssueDetails struct {
	Resolution string               `json:"resolution"` // The resolution status, either blocking the content or warning. enum values: HeavyAdBlocked, HeavyAdWarning
	Reason     string               `json:"reason"`     // The reason the ad was blocked, total network or cpu or peak cpu. enum values: NetworkTotalLimit, CpuTotalLimit, CpuPeakLimit
	Frame      *AuditsAffectedFrame `json:"frame"`      // The frame that was blocked.
}

// No Description.
type AuditsSourceCodeLocation struct {
	ScriptId     string `json:"scriptId,omitempty"` //
	Url          string `json:"url"`                //
	LineNumber   int    `json:"lineNumber"`         //
	ColumnNumber int    `json:"columnNumber"`       //
}

// No Description.
type AuditsContentSecurityPolicyIssueDetails struct {
	BlockedURL                         string                    `json:"blockedURL,omitempty"`               // The url not included in allowed sources.
	ViolatedDirective                  string                    `json:"violatedDirective"`                  // Specific directive that is violated, causing the CSP issue.
	IsReportOnly                       bool                      `json:"isReportOnly"`                       //
	ContentSecurityPolicyViolationType string                    `json:"contentSecurityPolicyViolationType"` //  enum values: kInlineViolation, kEvalViolation, kURLViolation, kTrustedTypesSinkViolation, kTrustedTypesPolicyViolation, kWasmEvalViolation
	FrameAncestor                      *AuditsAffectedFrame      `json:"frameAncestor,omitempty"`            //
	SourceCodeLocation                 *AuditsSourceCodeLocation `json:"sourceCodeLocation,omitempty"`       //
	ViolatingNodeId                    int                       `json:"violatingNodeId,omitempty"`          //
}

// Details for a issue arising from an SAB being instantiated in, or transferred to a context that is not cross-origin isolated.
type AuditsSharedArrayBufferIssueDetails struct {
	SourceCodeLocation *AuditsSourceCodeLocation `json:"sourceCodeLocation"` //
	IsWarning          bool                      `json:"isWarning"`          //
	Type               string                    `json:"type"`               //  enum values: TransferIssue, CreationIssue
}

// No Description.
type AuditsTrustedWebActivityIssueDetails struct {
	Url            string `json:"url"`                      // The url that triggers the violation.
	ViolationType  string `json:"violationType"`            //  enum values: kHttpError, kUnavailableOffline, kDigitalAssetLinks
	HttpStatusCode int    `json:"httpStatusCode,omitempty"` //
	PackageName    string `json:"packageName,omitempty"`    // The package name of the Trusted Web Activity client app. This field is only used when violation type is kDigitalAssetLinks.
	Signature      string `json:"signature,omitempty"`      // The signature of the Trusted Web Activity client app. This field is only used when violation type is kDigitalAssetLinks.
}

// No Description.
type AuditsLowTextContrastIssueDetails struct {
	ViolatingNodeId       int     `json:"violatingNodeId"`       //
	ViolatingNodeSelector string  `json:"violatingNodeSelector"` //
	ContrastRatio         float64 `json:"contrastRatio"`         //
	ThresholdAA           float64 `json:"thresholdAA"`           //
	ThresholdAAA          float64 `json:"thresholdAAA"`          //
	FontSize              string  `json:"fontSize"`              //
	FontWeight            string  `json:"fontWeight"`            //
}

// Details for a CORS related issue, e.g. a warning or error related to CORS RFC1918 enforcement.
type AuditsCorsIssueDetails struct {
	CorsErrorStatus        *NetworkCorsErrorStatus     `json:"corsErrorStatus"`                  //
	IsWarning              bool                        `json:"isWarning"`                        //
	Request                *AuditsAffectedRequest      `json:"request"`                          //
	Location               *AuditsSourceCodeLocation   `json:"location,omitempty"`               //
	InitiatorOrigin        string                      `json:"initiatorOrigin,omitempty"`        //
	ResourceIPAddressSpace string                      `json:"resourceIPAddressSpace,omitempty"` //  enum values: Local, Private, Public, Unknown
	ClientSecurityState    *NetworkClientSecurityState `json:"clientSecurityState,omitempty"`    //
}

// Details for issues around "Attribution Reporting API" usage. Explainer: https://github.com/WICG/attribution-reporting-api
type AuditsAttributionReportingIssueDetails struct {
	ViolationType    string                 `json:"violationType"`              //  enum values: PermissionPolicyDisabled, UntrustworthyReportingOrigin, InsecureContext, InvalidHeader, InvalidRegisterTriggerHeader, InvalidEligibleHeader, SourceAndTriggerHeaders, SourceIgnored, TriggerIgnored, OsSourceIgnored, OsTriggerIgnored, InvalidRegisterOsSourceHeader, InvalidRegisterOsTriggerHeader, WebAndOsHeaders, NoWebOrOsSupport
	Request          *AuditsAffectedRequest `json:"request,omitempty"`          //
	ViolatingNodeId  int                    `json:"violatingNodeId,omitempty"`  //
	InvalidParameter string                 `json:"invalidParameter,omitempty"` //
}

// Details for issues about documents in Quirks Mode or Limited Quirks Mode that affects page layouting.
type AuditsQuirksModeIssueDetails struct {
	IsLimitedQuirksMode bool   `json:"isLimitedQuirksMode"` // If false, it means the document's mode is "quirks" instead of "limited-quirks".
	DocumentNodeId      int    `json:"documentNodeId"`      //
	Url                 string `json:"url"`                 //
	FrameId             string `json:"frameId"`             //
	LoaderId            string `json:"loaderId"`            //
}

// No Description.
type AuditsNavigatorUserAgentIssueDetails struct {
	Url      string                    `json:"url"`                //
	Location *AuditsSourceCodeLocation `json:"location,omitempty"` //
}

// Depending on the concrete errorType, different properties are set.
type AuditsGenericIssueDetails struct {
	ErrorType              string `json:"errorType"`                        // Issues with the same errorType are aggregated in the frontend. enum values: CrossOriginPortalPostMessageError, FormLabelForNameError, FormDuplicateIdForInputError, FormInputWithNoLabelError, FormAutocompleteAttributeEmptyError, FormEmptyIdAndNameAttributesForInputError, FormAriaLabelledByToNonExistingId, FormInputAssignedAutocompleteValueToIdOrNameAttributeError, FormLabelHasNeitherForNorNestedInput, FormLabelForMatchesNonExistingIdError, FormInputHasWrongButWellIntendedAutocompleteValueError
	FrameId                string `json:"frameId,omitempty"`                //
	ViolatingNodeId        int    `json:"violatingNodeId,omitempty"`        //
	ViolatingNodeAttribute string `json:"violatingNodeAttribute,omitempty"` //
}

// This issue tracks information needed to print a deprecation message. https://source.chromium.org/chromium/chromium/src/+/main:third_party/blink/renderer/core/frame/third_party/blink/renderer/core/frame/deprecation/README.md
type AuditsDeprecationIssueDetails struct {
	AffectedFrame      *AuditsAffectedFrame      `json:"affectedFrame,omitempty"` //
	SourceCodeLocation *AuditsSourceCodeLocation `json:"sourceCodeLocation"`      //
	Type               string                    `json:"type"`                    // One of the deprecation names from third_party/blink/renderer/core/frame/deprecation/deprecation.json5
}

// This issue warns about sites in the redirect chain of a finished navigation that may be flagged as trackers and have their state cleared if they don't receive a user interaction. Note that in this context 'site' means eTLD+1. For example, if the URL `https://example.test:80/bounce` was in the redirect chain, the site reported would be `example.test`.
type AuditsBounceTrackingIssueDetails struct {
	TrackingSites []string `json:"trackingSites"` //
}

// No Description.
type AuditsFederatedAuthRequestIssueDetails struct {
	FederatedAuthRequestIssueReason string `json:"federatedAuthRequestIssueReason"` //  enum values: ShouldEmbargo, TooManyRequests, WellKnownHttpNotFound, WellKnownNoResponse, WellKnownInvalidResponse, WellKnownListEmpty, WellKnownInvalidContentType, ConfigNotInWellKnown, WellKnownTooBig, ConfigHttpNotFound, ConfigNoResponse, ConfigInvalidResponse, ConfigInvalidContentType, ClientMetadataHttpNotFound, ClientMetadataNoResponse, ClientMetadataInvalidResponse, ClientMetadataInvalidContentType, DisabledInSettings, ErrorFetchingSignin, InvalidSigninResponse, AccountsHttpNotFound, AccountsNoResponse, AccountsInvalidResponse, AccountsListEmpty, AccountsInvalidContentType, IdTokenHttpNotFound, IdTokenNoResponse, IdTokenInvalidResponse, IdTokenInvalidRequest, IdTokenInvalidContentType, ErrorIdToken, Canceled, RpPageNotVisible
}

// This issue tracks client hints related issues. It's used to deprecate old features, encourage the use of new ones, and provide general guidance.
type AuditsClientHintIssueDetails struct {
	SourceCodeLocation    *AuditsSourceCodeLocation `json:"sourceCodeLocation"`    //
	ClientHintIssueReason string                    `json:"clientHintIssueReason"` //  enum values: MetaTagAllowListInvalidOrigin, MetaTagModifiedHTML
}

// This struct holds a list of optional fields with additional information specific to the kind of issue. When adding a new issue code, please also add a new optional field to this type.
type AuditsInspectorIssueDetails struct {
	CookieIssueDetails                *AuditsCookieIssueDetails                `json:"cookieIssueDetails,omitempty"`                //
	MixedContentIssueDetails          *AuditsMixedContentIssueDetails          `json:"mixedContentIssueDetails,omitempty"`          //
	BlockedByResponseIssueDetails     *AuditsBlockedByResponseIssueDetails     `json:"blockedByResponseIssueDetails,omitempty"`     //
	HeavyAdIssueDetails               *AuditsHeavyAdIssueDetails               `json:"heavyAdIssueDetails,omitempty"`               //
	ContentSecurityPolicyIssueDetails *AuditsContentSecurityPolicyIssueDetails `json:"contentSecurityPolicyIssueDetails,omitempty"` //
	SharedArrayBufferIssueDetails     *AuditsSharedArrayBufferIssueDetails     `json:"sharedArrayBufferIssueDetails,omitempty"`     //
	TwaQualityEnforcementDetails      *AuditsTrustedWebActivityIssueDetails    `json:"twaQualityEnforcementDetails,omitempty"`      //
	LowTextContrastIssueDetails       *AuditsLowTextContrastIssueDetails       `json:"lowTextContrastIssueDetails,omitempty"`       //
	CorsIssueDetails                  *AuditsCorsIssueDetails                  `json:"corsIssueDetails,omitempty"`                  //
	AttributionReportingIssueDetails  *AuditsAttributionReportingIssueDetails  `json:"attributionReportingIssueDetails,omitempty"`  //
	QuirksModeIssueDetails            *AuditsQuirksModeIssueDetails            `json:"quirksModeIssueDetails,omitempty"`            //
	NavigatorUserAgentIssueDetails    *AuditsNavigatorUserAgentIssueDetails    `json:"navigatorUserAgentIssueDetails,omitempty"`    //
	GenericIssueDetails               *AuditsGenericIssueDetails               `json:"genericIssueDetails,omitempty"`               //
	DeprecationIssueDetails           *AuditsDeprecationIssueDetails           `json:"deprecationIssueDetails,omitempty"`           //
	ClientHintIssueDetails            *AuditsClientHintIssueDetails            `json:"clientHintIssueDetails,omitempty"`            //
	FederatedAuthRequestIssueDetails  *AuditsFederatedAuthRequestIssueDetails  `json:"federatedAuthRequestIssueDetails,omitempty"`  //
	BounceTrackingIssueDetails        *AuditsBounceTrackingIssueDetails        `json:"bounceTrackingIssueDetails,omitempty"`        //
}

// An inspector issue reported from the back-end.
type AuditsInspectorIssue struct {
	Code    string                       `json:"code"`              //  enum values: CookieIssue, MixedContentIssue, BlockedByResponseIssue, HeavyAdIssue, ContentSecurityPolicyIssue, SharedArrayBufferIssue, TrustedWebActivityIssue, LowTextContrastIssue, CorsIssue, AttributionReportingIssue, QuirksModeIssue, NavigatorUserAgentIssue, GenericIssue, DeprecationIssue, ClientHintIssue, FederatedAuthRequestIssue, BounceTrackingIssue
	Details *AuditsInspectorIssueDetails `json:"details"`           //
	IssueId string                       `json:"issueId,omitempty"` // A unique id for this issue. May be omitted if no other entity (e.g. exception, CDP message, etc.) is referencing this issue.
}

//
type AuditsIssueAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Issue *AuditsInspectorIssue `json:"issue"` //
	} `json:"Params,omitempty"`
}

type Audits struct {
	target gcdmessage.ChromeTargeter
}

func NewAudits(target gcdmessage.ChromeTargeter) *Audits {
	c := &Audits{target: target}
	return c
}

type AuditsGetEncodedResponseParams struct {
	// Identifier of the network request to get content for.
	RequestId string `json:"requestId"`
	// The encoding to use.
	Encoding string `json:"encoding"`
	// The quality of the encoding (0-1). (defaults to 1)
	Quality float64 `json:"quality,omitempty"`
	// Whether to only return the size information (defaults to false).
	SizeOnly bool `json:"sizeOnly,omitempty"`
}

// GetEncodedResponseWithParams - Returns the response body and size if it were re-encoded with the specified settings. Only applies to images.
// Returns -  body - The encoded body as a base64 string. Omitted if sizeOnly is true. (Encoded as a base64 string when passed over JSON) originalSize - Size before re-encoding. encodedSize - Size after re-encoding.
func (c *Audits) GetEncodedResponseWithParams(ctx context.Context, v *AuditsGetEncodedResponseParams) (string, int, int, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.getEncodedResponse", Params: v})
	if err != nil {
		return "", 0, 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Body         string
			OriginalSize int
			EncodedSize  int
		}
	}

	if resp == nil {
		return "", 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", 0, 0, err
	}

	if chromeData.Error != nil {
		return "", 0, 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Body, chromeData.Result.OriginalSize, chromeData.Result.EncodedSize, nil
}

// GetEncodedResponse - Returns the response body and size if it were re-encoded with the specified settings. Only applies to images.
// requestId - Identifier of the network request to get content for.
// encoding - The encoding to use.
// quality - The quality of the encoding (0-1). (defaults to 1)
// sizeOnly - Whether to only return the size information (defaults to false).
// Returns -  body - The encoded body as a base64 string. Omitted if sizeOnly is true. (Encoded as a base64 string when passed over JSON) originalSize - Size before re-encoding. encodedSize - Size after re-encoding.
func (c *Audits) GetEncodedResponse(ctx context.Context, requestId string, encoding string, quality float64, sizeOnly bool) (string, int, int, error) {
	var v AuditsGetEncodedResponseParams
	v.RequestId = requestId
	v.Encoding = encoding
	v.Quality = quality
	v.SizeOnly = sizeOnly
	return c.GetEncodedResponseWithParams(ctx, &v)
}

// Disables issues domain, prevents further issues from being reported to the client.
func (c *Audits) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.disable"})
}

// Enables issues domain, sends the issues collected so far to the client by means of the `issueAdded` event.
func (c *Audits) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.enable"})
}

type AuditsCheckContrastParams struct {
	// Whether to report WCAG AAA level issues. Default is false.
	ReportAAA bool `json:"reportAAA,omitempty"`
}

// CheckContrastWithParams - Runs the contrast check for the target page. Found issues are reported using Audits.issueAdded event.
func (c *Audits) CheckContrastWithParams(ctx context.Context, v *AuditsCheckContrastParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.checkContrast", Params: v})
}

// CheckContrast - Runs the contrast check for the target page. Found issues are reported using Audits.issueAdded event.
// reportAAA - Whether to report WCAG AAA level issues. Default is false.
func (c *Audits) CheckContrast(ctx context.Context, reportAAA bool) (*gcdmessage.ChromeResponse, error) {
	var v AuditsCheckContrastParams
	v.ReportAAA = reportAAA
	return c.CheckContrastWithParams(ctx, &v)
}
