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
type AuditsSameSiteCookieIssueDetails struct {
	Cookie                 *AuditsAffectedCookie  `json:"cookie"`                   //
	CookieWarningReasons   []string               `json:"cookieWarningReasons"`     //  enum values: WarnSameSiteUnspecifiedCrossSiteContext, WarnSameSiteNoneInsecure, WarnSameSiteUnspecifiedLaxAllowUnsafe, WarnSameSiteStrictLaxDowngradeStrict, WarnSameSiteStrictCrossDowngradeStrict, WarnSameSiteStrictCrossDowngradeLax, WarnSameSiteLaxCrossDowngradeStrict, WarnSameSiteLaxCrossDowngradeLax
	CookieExclusionReasons []string               `json:"cookieExclusionReasons"`   //  enum values: ExcludeSameSiteUnspecifiedTreatedAsLax, ExcludeSameSiteNoneInsecure
	Operation              string                 `json:"operation"`                // Optionally identifies the site-for-cookies and the cookie url, which may be used by the front-end as additional context. enum values: SetCookie, ReadCookie
	SiteForCookies         string                 `json:"siteForCookies,omitempty"` //
	CookieUrl              string                 `json:"cookieUrl,omitempty"`      //
	Request                *AuditsAffectedRequest `json:"request,omitempty"`        //
}

// No Description.
type AuditsMixedContentIssueDetails struct {
	ResourceType     string                 `json:"resourceType,omitempty"` // The type of resource causing the mixed content issue (css, js, iframe, form,...). Marked as optional because it is mapped to from blink::mojom::RequestContextType, which will be replaced by network::mojom::RequestDestination enum values: Audio, Beacon, CSPReport, Download, EventSource, Favicon, Font, Form, Frame, Image, Import, Manifest, Ping, PluginData, PluginResource, Prefetch, Resource, Script, ServiceWorker, SharedWorker, Stylesheet, Track, Video, Worker, XMLHttpRequest, XSLT
	ResolutionStatus string                 `json:"resolutionStatus"`       // The way the mixed content issue is being resolved. enum values: MixedContentBlocked, MixedContentAutomaticallyUpgraded, MixedContentWarning
	InsecureURL      string                 `json:"insecureURL"`            // The unsafe http url causing the mixed content issue.
	MainResourceURL  string                 `json:"mainResourceURL"`        // The url responsible for the call to an unsafe url.
	Request          *AuditsAffectedRequest `json:"request,omitempty"`      // The mixed content request. Does not always exist (e.g. for unsafe form submission urls).
	Frame            *AuditsAffectedFrame   `json:"frame,omitempty"`        // Optional because not every mixed content issue is necessarily linked to a frame.
}

// Details for a request that has been blocked with the BLOCKED_BY_RESPONSE code. Currently only used for COEP/COOP, but may be extended to include some CSP errors in the future.
type AuditsBlockedByResponseIssueDetails struct {
	Request *AuditsAffectedRequest `json:"request"`         //
	Frame   *AuditsAffectedFrame   `json:"frame,omitempty"` //
	Reason  string                 `json:"reason"`          //  enum values: CoepFrameResourceNeedsCoepHeader, CoopSandboxedIFrameCannotNavigateToCoopPage, CorpNotSameOrigin, CorpNotSameOriginAfterDefaultedToSameOriginByCoep, CorpNotSameSite
}

// No Description.
type AuditsHeavyAdIssueDetails struct {
	Resolution string               `json:"resolution"` // The resolution status, either blocking the content or warning. enum values: HeavyAdBlocked, HeavyAdWarning
	Reason     string               `json:"reason"`     // The reason the ad was blocked, total network or cpu or peak cpu. enum values: NetworkTotalLimit, CpuTotalLimit, CpuPeakLimit
	Frame      *AuditsAffectedFrame `json:"frame"`      // The frame that was blocked.
}

// No Description.
type AuditsSourceCodeLocation struct {
	Url          string `json:"url"`          //
	LineNumber   int    `json:"lineNumber"`   //
	ColumnNumber int    `json:"columnNumber"` //
}

// No Description.
type AuditsContentSecurityPolicyIssueDetails struct {
	BlockedURL                         string                    `json:"blockedURL,omitempty"`               // The url not included in allowed sources.
	ViolatedDirective                  string                    `json:"violatedDirective"`                  // Specific directive that is violated, causing the CSP issue.
	ContentSecurityPolicyViolationType string                    `json:"contentSecurityPolicyViolationType"` //  enum values: kInlineViolation, kEvalViolation, kURLViolation, kTrustedTypesSinkViolation, kTrustedTypesPolicyViolation
	FrameAncestor                      *AuditsAffectedFrame      `json:"frameAncestor,omitempty"`            //
	SourceCodeLocation                 *AuditsSourceCodeLocation `json:"sourceCodeLocation,omitempty"`       //
}

// This struct holds a list of optional fields with additional information specific to the kind of issue. When adding a new issue code, please also add a new optional field to this type.
type AuditsInspectorIssueDetails struct {
	SameSiteCookieIssueDetails        *AuditsSameSiteCookieIssueDetails        `json:"sameSiteCookieIssueDetails,omitempty"`        //
	MixedContentIssueDetails          *AuditsMixedContentIssueDetails          `json:"mixedContentIssueDetails,omitempty"`          //
	BlockedByResponseIssueDetails     *AuditsBlockedByResponseIssueDetails     `json:"blockedByResponseIssueDetails,omitempty"`     //
	HeavyAdIssueDetails               *AuditsHeavyAdIssueDetails               `json:"heavyAdIssueDetails,omitempty"`               //
	ContentSecurityPolicyIssueDetails *AuditsContentSecurityPolicyIssueDetails `json:"contentSecurityPolicyIssueDetails,omitempty"` //
}

// An inspector issue reported from the back-end.
type AuditsInspectorIssue struct {
	Code    string                       `json:"code"`    //  enum values: SameSiteCookieIssue, MixedContentIssue, BlockedByResponseIssue, HeavyAdIssue, ContentSecurityPolicyIssue
	Details *AuditsInspectorIssueDetails `json:"details"` //
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
// Returns -  body - The encoded body as a base64 string. Omitted if sizeOnly is true. originalSize - Size before re-encoding. encodedSize - Size after re-encoding.
func (c *Audits) GetEncodedResponseWithParams(ctx context.Context, v *AuditsGetEncodedResponseParams) (string, int, int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.getEncodedResponse", Params: v})
	if err != nil {
		return "", 0, 0, err
	}

	var chromeData struct {
		Result struct {
			Body         string
			OriginalSize int
			EncodedSize  int
		}
	}

	if resp == nil {
		return "", 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", 0, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", 0, 0, err
	}

	return chromeData.Result.Body, chromeData.Result.OriginalSize, chromeData.Result.EncodedSize, nil
}

// GetEncodedResponse - Returns the response body and size if it were re-encoded with the specified settings. Only applies to images.
// requestId - Identifier of the network request to get content for.
// encoding - The encoding to use.
// quality - The quality of the encoding (0-1). (defaults to 1)
// sizeOnly - Whether to only return the size information (defaults to false).
// Returns -  body - The encoded body as a base64 string. Omitted if sizeOnly is true. originalSize - Size before re-encoding. encodedSize - Size after re-encoding.
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
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.disable"})
}

// Enables issues domain, sends the issues collected so far to the client by means of the `issueAdded` event.
func (c *Audits) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, ctx, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Audits.enable"})
}
