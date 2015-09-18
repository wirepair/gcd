// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Security functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// An explanation of an factor contributing to the security state.
type SecuritySecurityStateExplanation struct {
	SecurityState string `json:"securityState"`           // Security state representing the severity of the factor being explained. enum values: unknown, neutral, insecure, warning, secure
	Summary       string `json:"summary"`                 // Short phrase describing the type of factor.
	Description   string `json:"description"`             // Full text explanation of the factor.
	CertificateId int    `json:"certificateId,omitempty"` // Associated certificate id.
}

// Information about mixed content on the page.
type SecurityMixedContentStatus struct {
	RanInsecureContent            bool   `json:"ranInsecureContent"`            // True if the page ran insecure content such as scripts.
	DisplayedInsecureContent      bool   `json:"displayedInsecureContent"`      // True if the page displayed insecure content such as images.
	RanInsecureContentStyle       string `json:"ranInsecureContentStyle"`       // Security state representing a page that ran insecure content. enum values: unknown, neutral, insecure, warning, secure
	DisplayedInsecureContentStyle string `json:"displayedInsecureContentStyle"` // Security state representing a page that displayed insecure content. enum values: unknown, neutral, insecure, warning, secure
}

// The security state of the page changed.
type SecuritySecurityStateChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		SecurityState         string                              `json:"securityState"`                   // Security state. enum values: unknown, neutral, insecure, warning, secure
		Explanations          []*SecuritySecurityStateExplanation `json:"explanations,omitempty"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
		MixedContentStatus    *SecurityMixedContentStatus         `json:"mixedContentStatus,omitempty"`    // Information about mixed content on the page.
		SchemeIsCryptographic bool                                `json:"schemeIsCryptographic,omitempty"` // True if the page was loaded over cryptographic transport such as HTTPS.
	} `json:"Params,omitempty"`
}

type Security struct {
	target gcdmessage.ChromeTargeter
}

func NewSecurity(target gcdmessage.ChromeTargeter) *Security {
	c := &Security{target: target}
	return c
}

// Enables tracking security state changes.
func (c *Security) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.enable"})
}

// Disables tracking security state changes.
func (c *Security) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.disable"})
}
