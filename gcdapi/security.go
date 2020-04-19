// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Security functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Details about the security state of the page certificate.
type SecurityCertificateSecurityState struct {
	Protocol                    string   `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                 string   `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup            string   `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                      string   `json:"cipher"`                            // Cipher name.
	Mac                         string   `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Certificate                 []string `json:"certificate"`                       // Page certificate.
	SubjectName                 string   `json:"subjectName"`                       // Certificate subject name.
	Issuer                      string   `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                   float64  `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                     float64  `json:"validTo"`                           // Certificate valid to (expiration) date
	CertificateNetworkError     string   `json:"certificateNetworkError,omitempty"` // The highest priority network error code, if the certificate has an error.
	CertificateHasWeakSignature bool     `json:"certificateHasWeakSignature"`       // True if the certificate uses a weak signature aglorithm.
	CertificateHasSha1Signature bool     `json:"certificateHasSha1Signature"`       // True if the certificate has a SHA1 signature in the chain.
	ModernSSL                   bool     `json:"modernSSL"`                         // True if modern SSL
	ObsoleteSslProtocol         bool     `json:"obsoleteSslProtocol"`               // True if the connection is using an obsolete SSL protocol.
	ObsoleteSslKeyExchange      bool     `json:"obsoleteSslKeyExchange"`            // True if the connection is using an obsolete SSL key exchange.
	ObsoleteSslCipher           bool     `json:"obsoleteSslCipher"`                 // True if the connection is using an obsolete SSL cipher.
	ObsoleteSslSignature        bool     `json:"obsoleteSslSignature"`              // True if the connection is using an obsolete SSL signature.
}

// No Description.
type SecuritySafetyTipInfo struct {
	SafetyTipStatus string `json:"safetyTipStatus"`   // Describes whether the page triggers any safety tips or reputation warnings. Default is unknown. enum values: badReputation, lookalike
	SafeUrl         string `json:"safeUrl,omitempty"` // The URL the safety tip suggested ("Did you mean?"). Only filled in for lookalike matches.
}

// Security state information about the page.
type SecurityVisibleSecurityState struct {
	SecurityState            string                            `json:"securityState"`                      // The security level of the page. enum values: unknown, neutral, insecure, secure, info, insecure-broken
	CertificateSecurityState *SecurityCertificateSecurityState `json:"certificateSecurityState,omitempty"` // Security state details about the page certificate.
	SafetyTipInfo            *SecuritySafetyTipInfo            `json:"safetyTipInfo,omitempty"`            // The type of Safety Tip triggered on the page. Note that this field will be set even if the Safety Tip UI was not actually shown.
	SecurityStateIssueIds    []string                          `json:"securityStateIssueIds"`              // Array of security state issues ids.
}

// An explanation of an factor contributing to the security state.
type SecuritySecurityStateExplanation struct {
	SecurityState    string   `json:"securityState"`             // Security state representing the severity of the factor being explained. enum values: unknown, neutral, insecure, secure, info, insecure-broken
	Title            string   `json:"title"`                     // Title describing the type of factor.
	Summary          string   `json:"summary"`                   // Short phrase describing the type of factor.
	Description      string   `json:"description"`               // Full text explanation of the factor.
	MixedContentType string   `json:"mixedContentType"`          // The type of mixed content described by the explanation. enum values: blockable, optionally-blockable, none
	Certificate      []string `json:"certificate"`               // Page certificate.
	Recommendations  []string `json:"recommendations,omitempty"` // Recommendations to fix any issues.
}

// Information about insecure content on the page.
type SecurityInsecureContentStatus struct {
	RanMixedContent                bool   `json:"ranMixedContent"`                // Always false.
	DisplayedMixedContent          bool   `json:"displayedMixedContent"`          // Always false.
	ContainedMixedForm             bool   `json:"containedMixedForm"`             // Always false.
	RanContentWithCertErrors       bool   `json:"ranContentWithCertErrors"`       // Always false.
	DisplayedContentWithCertErrors bool   `json:"displayedContentWithCertErrors"` // Always false.
<<<<<<< HEAD
	RanInsecureContentStyle        string `json:"ranInsecureContentStyle"`        // Always set to unknown. enum values: unknown, neutral, insecure, secure, info, insecure-broken
	DisplayedInsecureContentStyle  string `json:"displayedInsecureContentStyle"`  // Always set to unknown. enum values: unknown, neutral, insecure, secure, info, insecure-broken
=======
	RanInsecureContentStyle        string `json:"ranInsecureContentStyle"`        // Always set to unknown. enum values: unknown, neutral, insecure, secure, info
	DisplayedInsecureContentStyle  string `json:"displayedInsecureContentStyle"`  // Always set to unknown. enum values: unknown, neutral, insecure, secure, info
>>>>>>> master
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the `handleCertificateError` command. Note: this event does not fire if the certificate error has been allowed internally. Only one client per target should override certificate errors at the same time.
type SecurityCertificateErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		EventId    int    `json:"eventId"`    // The ID of the event.
		ErrorType  string `json:"errorType"`  // The type of the error.
		RequestURL string `json:"requestURL"` // The url that was requested.
	} `json:"Params,omitempty"`
}

// The security state of the page changed.
type SecurityVisibleSecurityStateChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		VisibleSecurityState *SecurityVisibleSecurityState `json:"visibleSecurityState"` // Security state information about the page.
	} `json:"Params,omitempty"`
}

// The security state of the page changed.
type SecuritySecurityStateChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		SecurityState         string                              `json:"securityState"`         // Security state. enum values: unknown, neutral, insecure, secure, info, insecure-broken
		SchemeIsCryptographic bool                                `json:"schemeIsCryptographic"` // True if the page was loaded over cryptographic transport such as HTTPS.
		Explanations          []*SecuritySecurityStateExplanation `json:"explanations"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
		InsecureContentStatus *SecurityInsecureContentStatus      `json:"insecureContentStatus"` // Information about insecure content on the page.
		Summary               string                              `json:"summary,omitempty"`     // Overrides user-visible description of the state.
	} `json:"Params,omitempty"`
}

type Security struct {
	target gcdmessage.ChromeTargeter
}

func NewSecurity(target gcdmessage.ChromeTargeter) *Security {
	c := &Security{target: target}
	return c
}

// Disables tracking security state changes.
func (c *Security) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.disable"})
}

// Enables tracking security state changes.
func (c *Security) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.enable"})
}

type SecuritySetIgnoreCertificateErrorsParams struct {
	// If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
}

// SetIgnoreCertificateErrorsWithParams - Enable/disable whether all certificate errors should be ignored.
func (c *Security) SetIgnoreCertificateErrorsWithParams(v *SecuritySetIgnoreCertificateErrorsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.setIgnoreCertificateErrors", Params: v})
}

// SetIgnoreCertificateErrors - Enable/disable whether all certificate errors should be ignored.
// ignore - If true, all certificate errors will be ignored.
func (c *Security) SetIgnoreCertificateErrors(ignore bool) (*gcdmessage.ChromeResponse, error) {
	var v SecuritySetIgnoreCertificateErrorsParams
	v.Ignore = ignore
	return c.SetIgnoreCertificateErrorsWithParams(&v)
}

type SecurityHandleCertificateErrorParams struct {
	// The ID of the event.
	EventId int `json:"eventId"`
	// The action to take on the certificate error. enum values: continue, cancel
	Action string `json:"action"`
}

// HandleCertificateErrorWithParams - Handles a certificate error that fired a certificateError event.
func (c *Security) HandleCertificateErrorWithParams(v *SecurityHandleCertificateErrorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.handleCertificateError", Params: v})
}

// HandleCertificateError - Handles a certificate error that fired a certificateError event.
// eventId - The ID of the event.
// action - The action to take on the certificate error. enum values: continue, cancel
func (c *Security) HandleCertificateError(eventId int, action string) (*gcdmessage.ChromeResponse, error) {
	var v SecurityHandleCertificateErrorParams
	v.EventId = eventId
	v.Action = action
	return c.HandleCertificateErrorWithParams(&v)
}

type SecuritySetOverrideCertificateErrorsParams struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// SetOverrideCertificateErrorsWithParams - Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with `handleCertificateError` commands.
func (c *Security) SetOverrideCertificateErrorsWithParams(v *SecuritySetOverrideCertificateErrorsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.setOverrideCertificateErrors", Params: v})
}

// SetOverrideCertificateErrors - Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with `handleCertificateError` commands.
// override - If true, certificate errors will be overridden.
func (c *Security) SetOverrideCertificateErrors(override bool) (*gcdmessage.ChromeResponse, error) {
	var v SecuritySetOverrideCertificateErrorsParams
	v.Override = override
	return c.SetOverrideCertificateErrorsWithParams(&v)
}
