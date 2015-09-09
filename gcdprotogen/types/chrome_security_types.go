// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Security types.
// API Version: 1.1
package types

// The security level of a page or resource.
type ChromeSecuritySecurityState string // possible values: unknown, neutral, insecure, warning, secure

// An explanation of an factor contributing to the security state.
type ChromeSecuritySecurityStateExplanation struct {
	SecurityState *ChromeSecuritySecurityState `json:"securityState"`           // Security state representing the severity of the factor being explained.
	Summary       string                       `json:"summary"`                 // Short phrase describing the type of factor.
	Description   string                       `json:"description"`             // Full text explanation of the factor.
	CertificateId *ChromeNetworkCertificateId  `json:"certificateId,omitempty"` // Associated certificate id.
}

// Information about mixed content on the page.
type ChromeSecurityMixedContentStatus struct {
	RanInsecureContent            bool                         `json:"ranInsecureContent"`            // True if the page ran insecure content such as scripts.
	DisplayedInsecureContent      bool                         `json:"displayedInsecureContent"`      // True if the page displayed insecure content such as images.
	RanInsecureContentStyle       *ChromeSecuritySecurityState `json:"ranInsecureContentStyle"`       // Security state representing a page that ran insecure content.
	DisplayedInsecureContentStyle *ChromeSecuritySecurityState `json:"displayedInsecureContentStyle"` // Security state representing a page that displayed insecure content.
}
