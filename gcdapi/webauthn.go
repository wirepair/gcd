// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains WebAuthn functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type WebAuthnVirtualAuthenticatorOptions struct {
	Protocol                    string `json:"protocol"`                              //  enum values: u2f, ctap2
	Transport                   string `json:"transport"`                             //  enum values: usb, nfc, ble, cable, internal
	HasResidentKey              bool   `json:"hasResidentKey"`                        //
	HasUserVerification         bool   `json:"hasUserVerification"`                   //
	AutomaticPresenceSimulation bool   `json:"automaticPresenceSimulation,omitempty"` // If set to true, tests of user presence will succeed immediately. Otherwise, they will not be resolved. Defaults to true.
}

// No Description.
type WebAuthnCredential struct {
	CredentialId string `json:"credentialId"` //
	RpIdHash     string `json:"rpIdHash"`     // SHA-256 hash of the Relying Party ID the credential is scoped to. Must be 32 bytes long. See https://w3c.github.io/webauthn/#rpidhash
	PrivateKey   string `json:"privateKey"`   // The private key in PKCS#8 format.
	SignCount    int    `json:"signCount"`    // Signature counter. This is incremented by one for each successful assertion. See https://w3c.github.io/webauthn/#signature-counter
}

type WebAuthn struct {
	target gcdmessage.ChromeTargeter
}

func NewWebAuthn(target gcdmessage.ChromeTargeter) *WebAuthn {
	c := &WebAuthn{target: target}
	return c
}

// Enable the WebAuthn domain and start intercepting credential storage and retrieval with a virtual authenticator.
func (c *WebAuthn) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.enable"})
}

// Disable the WebAuthn domain.
func (c *WebAuthn) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.disable"})
}

type WebAuthnAddVirtualAuthenticatorParams struct {
	//
	Options *WebAuthnVirtualAuthenticatorOptions `json:"options"`
}

// AddVirtualAuthenticatorWithParams - Creates and adds a virtual authenticator.
// Returns -  authenticatorId -
func (c *WebAuthn) AddVirtualAuthenticatorWithParams(v *WebAuthnAddVirtualAuthenticatorParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.addVirtualAuthenticator", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			AuthenticatorId string
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

	return chromeData.Result.AuthenticatorId, nil
}

// AddVirtualAuthenticator - Creates and adds a virtual authenticator.
// options -
// Returns -  authenticatorId -
func (c *WebAuthn) AddVirtualAuthenticator(options *WebAuthnVirtualAuthenticatorOptions) (string, error) {
	var v WebAuthnAddVirtualAuthenticatorParams
	v.Options = options
	return c.AddVirtualAuthenticatorWithParams(&v)
}

type WebAuthnRemoveVirtualAuthenticatorParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// RemoveVirtualAuthenticatorWithParams - Removes the given authenticator.
func (c *WebAuthn) RemoveVirtualAuthenticatorWithParams(v *WebAuthnRemoveVirtualAuthenticatorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.removeVirtualAuthenticator", Params: v})
}

// RemoveVirtualAuthenticator - Removes the given authenticator.
// authenticatorId -
func (c *WebAuthn) RemoveVirtualAuthenticator(authenticatorId string) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnRemoveVirtualAuthenticatorParams
	v.AuthenticatorId = authenticatorId
	return c.RemoveVirtualAuthenticatorWithParams(&v)
}

type WebAuthnAddCredentialParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	Credential *WebAuthnCredential `json:"credential"`
}

// AddCredentialWithParams - Adds the credential to the specified authenticator.
func (c *WebAuthn) AddCredentialWithParams(v *WebAuthnAddCredentialParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.addCredential", Params: v})
}

// AddCredential - Adds the credential to the specified authenticator.
// authenticatorId -
// credential -
func (c *WebAuthn) AddCredential(authenticatorId string, credential *WebAuthnCredential) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnAddCredentialParams
	v.AuthenticatorId = authenticatorId
	v.Credential = credential
	return c.AddCredentialWithParams(&v)
}

type WebAuthnGetCredentialsParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// GetCredentialsWithParams - Returns all the credentials stored in the given virtual authenticator.
// Returns -  credentials -
func (c *WebAuthn) GetCredentialsWithParams(v *WebAuthnGetCredentialsParams) ([]*WebAuthnCredential, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.getCredentials", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Credentials []*WebAuthnCredential
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

	return chromeData.Result.Credentials, nil
}

// GetCredentials - Returns all the credentials stored in the given virtual authenticator.
// authenticatorId -
// Returns -  credentials -
func (c *WebAuthn) GetCredentials(authenticatorId string) ([]*WebAuthnCredential, error) {
	var v WebAuthnGetCredentialsParams
	v.AuthenticatorId = authenticatorId
	return c.GetCredentialsWithParams(&v)
}

type WebAuthnClearCredentialsParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// ClearCredentialsWithParams - Clears all the credentials from the specified device.
func (c *WebAuthn) ClearCredentialsWithParams(v *WebAuthnClearCredentialsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.clearCredentials", Params: v})
}

// ClearCredentials - Clears all the credentials from the specified device.
// authenticatorId -
func (c *WebAuthn) ClearCredentials(authenticatorId string) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnClearCredentialsParams
	v.AuthenticatorId = authenticatorId
	return c.ClearCredentialsWithParams(&v)
}

type WebAuthnSetUserVerifiedParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	IsUserVerified bool `json:"isUserVerified"`
}

// SetUserVerifiedWithParams - Sets whether User Verification succeeds or fails for an authenticator. The default is true.
func (c *WebAuthn) SetUserVerifiedWithParams(v *WebAuthnSetUserVerifiedParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.setUserVerified", Params: v})
}

// SetUserVerified - Sets whether User Verification succeeds or fails for an authenticator. The default is true.
// authenticatorId -
// isUserVerified -
func (c *WebAuthn) SetUserVerified(authenticatorId string, isUserVerified bool) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnSetUserVerifiedParams
	v.AuthenticatorId = authenticatorId
	v.IsUserVerified = isUserVerified
	return c.SetUserVerifiedWithParams(&v)
}
