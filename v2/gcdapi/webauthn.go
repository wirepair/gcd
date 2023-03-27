// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains WebAuthn functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type WebAuthnVirtualAuthenticatorOptions struct {
	Protocol                    string `json:"protocol"`                              //  enum values: u2f, ctap2
	Ctap2Version                string `json:"ctap2Version,omitempty"`                // Defaults to ctap2_0. Ignored if |protocol| == u2f. enum values: ctap2_0, ctap2_1
	Transport                   string `json:"transport"`                             //  enum values: usb, nfc, ble, cable, internal
	HasResidentKey              bool   `json:"hasResidentKey,omitempty"`              // Defaults to false.
	HasUserVerification         bool   `json:"hasUserVerification,omitempty"`         // Defaults to false.
	HasLargeBlob                bool   `json:"hasLargeBlob,omitempty"`                // If set to true, the authenticator will support the largeBlob extension. https://w3c.github.io/webauthn#largeBlob Defaults to false.
	HasCredBlob                 bool   `json:"hasCredBlob,omitempty"`                 // If set to true, the authenticator will support the credBlob extension. https://fidoalliance.org/specs/fido-v2.1-rd-20201208/fido-client-to-authenticator-protocol-v2.1-rd-20201208.html#sctn-credBlob-extension Defaults to false.
	HasMinPinLength             bool   `json:"hasMinPinLength,omitempty"`             // If set to true, the authenticator will support the minPinLength extension. https://fidoalliance.org/specs/fido-v2.1-ps-20210615/fido-client-to-authenticator-protocol-v2.1-ps-20210615.html#sctn-minpinlength-extension Defaults to false.
	HasPrf                      bool   `json:"hasPrf,omitempty"`                      // If set to true, the authenticator will support the prf extension. https://w3c.github.io/webauthn/#prf-extension Defaults to false.
	AutomaticPresenceSimulation bool   `json:"automaticPresenceSimulation,omitempty"` // If set to true, tests of user presence will succeed immediately. Otherwise, they will not be resolved. Defaults to true.
	IsUserVerified              bool   `json:"isUserVerified,omitempty"`              // Sets whether User Verification succeeds or fails for an authenticator. Defaults to false.
}

// No Description.
type WebAuthnCredential struct {
	CredentialId         string `json:"credentialId"`         //
	IsResidentCredential bool   `json:"isResidentCredential"` //
	RpId                 string `json:"rpId,omitempty"`       // Relying Party ID the credential is scoped to. Must be set when adding a credential.
	PrivateKey           string `json:"privateKey"`           // The ECDSA P-256 private key in PKCS#8 format. (Encoded as a base64 string when passed over JSON)
	UserHandle           string `json:"userHandle,omitempty"` // An opaque byte sequence with a maximum size of 64 bytes mapping the credential to a specific user. (Encoded as a base64 string when passed over JSON)
	SignCount            int    `json:"signCount"`            // Signature counter. This is incremented by one for each successful assertion. See https://w3c.github.io/webauthn/#signature-counter
	LargeBlob            string `json:"largeBlob,omitempty"`  // The large blob associated with the credential. See https://w3c.github.io/webauthn/#sctn-large-blob-extension (Encoded as a base64 string when passed over JSON)
}

// Triggered when a credential is added to an authenticator.
type WebAuthnCredentialAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		AuthenticatorId string              `json:"authenticatorId"` //
		Credential      *WebAuthnCredential `json:"credential"`      //
	} `json:"Params,omitempty"`
}

// Triggered when a credential is used in a webauthn assertion.
type WebAuthnCredentialAssertedEvent struct {
	Method string `json:"method"`
	Params struct {
		AuthenticatorId string              `json:"authenticatorId"` //
		Credential      *WebAuthnCredential `json:"credential"`      //
	} `json:"Params,omitempty"`
}

type WebAuthn struct {
	target gcdmessage.ChromeTargeter
}

func NewWebAuthn(target gcdmessage.ChromeTargeter) *WebAuthn {
	c := &WebAuthn{target: target}
	return c
}

type WebAuthnEnableParams struct {
	// Whether to enable the WebAuthn user interface. Enabling the UI is recommended for debugging and demo purposes, as it is closer to the real experience. Disabling the UI is recommended for automated testing. Supported at the embedder's discretion if UI is available. Defaults to false.
	EnableUI bool `json:"enableUI,omitempty"`
}

// EnableWithParams - Enable the WebAuthn domain and start intercepting credential storage and retrieval with a virtual authenticator.
func (c *WebAuthn) EnableWithParams(ctx context.Context, v *WebAuthnEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.enable", Params: v})
}

// Enable - Enable the WebAuthn domain and start intercepting credential storage and retrieval with a virtual authenticator.
// enableUI - Whether to enable the WebAuthn user interface. Enabling the UI is recommended for debugging and demo purposes, as it is closer to the real experience. Disabling the UI is recommended for automated testing. Supported at the embedder's discretion if UI is available. Defaults to false.
func (c *WebAuthn) Enable(ctx context.Context, enableUI bool) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnEnableParams
	v.EnableUI = enableUI
	return c.EnableWithParams(ctx, &v)
}

// Disable the WebAuthn domain.
func (c *WebAuthn) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.disable"})
}

type WebAuthnAddVirtualAuthenticatorParams struct {
	//
	Options *WebAuthnVirtualAuthenticatorOptions `json:"options"`
}

// AddVirtualAuthenticatorWithParams - Creates and adds a virtual authenticator.
// Returns -  authenticatorId -
func (c *WebAuthn) AddVirtualAuthenticatorWithParams(ctx context.Context, v *WebAuthnAddVirtualAuthenticatorParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.addVirtualAuthenticator", Params: v})
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
func (c *WebAuthn) AddVirtualAuthenticator(ctx context.Context, options *WebAuthnVirtualAuthenticatorOptions) (string, error) {
	var v WebAuthnAddVirtualAuthenticatorParams
	v.Options = options
	return c.AddVirtualAuthenticatorWithParams(ctx, &v)
}

type WebAuthnSetResponseOverrideBitsParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	// If isBogusSignature is set, overrides the signature in the authenticator response to be zero. Defaults to false.
	IsBogusSignature bool `json:"isBogusSignature,omitempty"`
	// If isBadUV is set, overrides the UV bit in the flags in the authenticator response to be zero. Defaults to false.
	IsBadUV bool `json:"isBadUV,omitempty"`
	// If isBadUP is set, overrides the UP bit in the flags in the authenticator response to be zero. Defaults to false.
	IsBadUP bool `json:"isBadUP,omitempty"`
}

// SetResponseOverrideBitsWithParams - Resets parameters isBogusSignature, isBadUV, isBadUP to false if they are not present.
func (c *WebAuthn) SetResponseOverrideBitsWithParams(ctx context.Context, v *WebAuthnSetResponseOverrideBitsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.setResponseOverrideBits", Params: v})
}

// SetResponseOverrideBits - Resets parameters isBogusSignature, isBadUV, isBadUP to false if they are not present.
// authenticatorId -
// isBogusSignature - If isBogusSignature is set, overrides the signature in the authenticator response to be zero. Defaults to false.
// isBadUV - If isBadUV is set, overrides the UV bit in the flags in the authenticator response to be zero. Defaults to false.
// isBadUP - If isBadUP is set, overrides the UP bit in the flags in the authenticator response to be zero. Defaults to false.
func (c *WebAuthn) SetResponseOverrideBits(ctx context.Context, authenticatorId string, isBogusSignature bool, isBadUV bool, isBadUP bool) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnSetResponseOverrideBitsParams
	v.AuthenticatorId = authenticatorId
	v.IsBogusSignature = isBogusSignature
	v.IsBadUV = isBadUV
	v.IsBadUP = isBadUP
	return c.SetResponseOverrideBitsWithParams(ctx, &v)
}

type WebAuthnRemoveVirtualAuthenticatorParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// RemoveVirtualAuthenticatorWithParams - Removes the given authenticator.
func (c *WebAuthn) RemoveVirtualAuthenticatorWithParams(ctx context.Context, v *WebAuthnRemoveVirtualAuthenticatorParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.removeVirtualAuthenticator", Params: v})
}

// RemoveVirtualAuthenticator - Removes the given authenticator.
// authenticatorId -
func (c *WebAuthn) RemoveVirtualAuthenticator(ctx context.Context, authenticatorId string) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnRemoveVirtualAuthenticatorParams
	v.AuthenticatorId = authenticatorId
	return c.RemoveVirtualAuthenticatorWithParams(ctx, &v)
}

type WebAuthnAddCredentialParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	Credential *WebAuthnCredential `json:"credential"`
}

// AddCredentialWithParams - Adds the credential to the specified authenticator.
func (c *WebAuthn) AddCredentialWithParams(ctx context.Context, v *WebAuthnAddCredentialParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.addCredential", Params: v})
}

// AddCredential - Adds the credential to the specified authenticator.
// authenticatorId -
// credential -
func (c *WebAuthn) AddCredential(ctx context.Context, authenticatorId string, credential *WebAuthnCredential) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnAddCredentialParams
	v.AuthenticatorId = authenticatorId
	v.Credential = credential
	return c.AddCredentialWithParams(ctx, &v)
}

type WebAuthnGetCredentialParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	CredentialId string `json:"credentialId"`
}

// GetCredentialWithParams - Returns a single credential stored in the given virtual authenticator that matches the credential ID.
// Returns -  credential -
func (c *WebAuthn) GetCredentialWithParams(ctx context.Context, v *WebAuthnGetCredentialParams) (*WebAuthnCredential, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.getCredential", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Credential *WebAuthnCredential
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

	return chromeData.Result.Credential, nil
}

// GetCredential - Returns a single credential stored in the given virtual authenticator that matches the credential ID.
// authenticatorId -
// credentialId -
// Returns -  credential -
func (c *WebAuthn) GetCredential(ctx context.Context, authenticatorId string, credentialId string) (*WebAuthnCredential, error) {
	var v WebAuthnGetCredentialParams
	v.AuthenticatorId = authenticatorId
	v.CredentialId = credentialId
	return c.GetCredentialWithParams(ctx, &v)
}

type WebAuthnGetCredentialsParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// GetCredentialsWithParams - Returns all the credentials stored in the given virtual authenticator.
// Returns -  credentials -
func (c *WebAuthn) GetCredentialsWithParams(ctx context.Context, v *WebAuthnGetCredentialsParams) ([]*WebAuthnCredential, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.getCredentials", Params: v})
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
func (c *WebAuthn) GetCredentials(ctx context.Context, authenticatorId string) ([]*WebAuthnCredential, error) {
	var v WebAuthnGetCredentialsParams
	v.AuthenticatorId = authenticatorId
	return c.GetCredentialsWithParams(ctx, &v)
}

type WebAuthnRemoveCredentialParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	CredentialId string `json:"credentialId"`
}

// RemoveCredentialWithParams - Removes a credential from the authenticator.
func (c *WebAuthn) RemoveCredentialWithParams(ctx context.Context, v *WebAuthnRemoveCredentialParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.removeCredential", Params: v})
}

// RemoveCredential - Removes a credential from the authenticator.
// authenticatorId -
// credentialId -
func (c *WebAuthn) RemoveCredential(ctx context.Context, authenticatorId string, credentialId string) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnRemoveCredentialParams
	v.AuthenticatorId = authenticatorId
	v.CredentialId = credentialId
	return c.RemoveCredentialWithParams(ctx, &v)
}

type WebAuthnClearCredentialsParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
}

// ClearCredentialsWithParams - Clears all the credentials from the specified device.
func (c *WebAuthn) ClearCredentialsWithParams(ctx context.Context, v *WebAuthnClearCredentialsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.clearCredentials", Params: v})
}

// ClearCredentials - Clears all the credentials from the specified device.
// authenticatorId -
func (c *WebAuthn) ClearCredentials(ctx context.Context, authenticatorId string) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnClearCredentialsParams
	v.AuthenticatorId = authenticatorId
	return c.ClearCredentialsWithParams(ctx, &v)
}

type WebAuthnSetUserVerifiedParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	IsUserVerified bool `json:"isUserVerified"`
}

// SetUserVerifiedWithParams - Sets whether User Verification succeeds or fails for an authenticator. The default is true.
func (c *WebAuthn) SetUserVerifiedWithParams(ctx context.Context, v *WebAuthnSetUserVerifiedParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.setUserVerified", Params: v})
}

// SetUserVerified - Sets whether User Verification succeeds or fails for an authenticator. The default is true.
// authenticatorId -
// isUserVerified -
func (c *WebAuthn) SetUserVerified(ctx context.Context, authenticatorId string, isUserVerified bool) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnSetUserVerifiedParams
	v.AuthenticatorId = authenticatorId
	v.IsUserVerified = isUserVerified
	return c.SetUserVerifiedWithParams(ctx, &v)
}

type WebAuthnSetAutomaticPresenceSimulationParams struct {
	//
	AuthenticatorId string `json:"authenticatorId"`
	//
	Enabled bool `json:"enabled"`
}

// SetAutomaticPresenceSimulationWithParams - Sets whether tests of user presence will succeed immediately (if true) or fail to resolve (if false) for an authenticator. The default is true.
func (c *WebAuthn) SetAutomaticPresenceSimulationWithParams(ctx context.Context, v *WebAuthnSetAutomaticPresenceSimulationParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "WebAuthn.setAutomaticPresenceSimulation", Params: v})
}

// SetAutomaticPresenceSimulation - Sets whether tests of user presence will succeed immediately (if true) or fail to resolve (if false) for an authenticator. The default is true.
// authenticatorId -
// enabled -
func (c *WebAuthn) SetAutomaticPresenceSimulation(ctx context.Context, authenticatorId string, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v WebAuthnSetAutomaticPresenceSimulationParams
	v.AuthenticatorId = authenticatorId
	v.Enabled = enabled
	return c.SetAutomaticPresenceSimulationWithParams(ctx, &v)
}
