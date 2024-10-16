// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains FedCm functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Corresponds to IdentityRequestAccount
type FedCmAccount struct {
	AccountId         string `json:"accountId"`                   //
	Email             string `json:"email"`                       //
	Name              string `json:"name"`                        //
	GivenName         string `json:"givenName"`                   //
	PictureUrl        string `json:"pictureUrl"`                  //
	IdpConfigUrl      string `json:"idpConfigUrl"`                //
	IdpLoginUrl       string `json:"idpLoginUrl"`                 //
	LoginState        string `json:"loginState"`                  //  enum values: SignIn, SignUp
	TermsOfServiceUrl string `json:"termsOfServiceUrl,omitempty"` // These two are only set if the loginState is signUp
	PrivacyPolicyUrl  string `json:"privacyPolicyUrl,omitempty"`  //
}

//
type FedCmDialogShownEvent struct {
	Method string `json:"method"`
	Params struct {
		DialogId   string          `json:"dialogId"`           //
		DialogType string          `json:"dialogType"`         //  enum values: AccountChooser, AutoReauthn, ConfirmIdpLogin, Error
		Accounts   []*FedCmAccount `json:"accounts"`           //
		Title      string          `json:"title"`              // These exist primarily so that the caller can verify the RP context was used appropriately.
		Subtitle   string          `json:"subtitle,omitempty"` //
	} `json:"Params,omitempty"`
}

// Triggered when a dialog is closed, either by user action, JS abort, or a command below.
type FedCmDialogClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		DialogId string `json:"dialogId"` //
	} `json:"Params,omitempty"`
}

type FedCm struct {
	target gcdmessage.ChromeTargeter
}

func NewFedCm(target gcdmessage.ChromeTargeter) *FedCm {
	c := &FedCm{target: target}
	return c
}

type FedCmEnableParams struct {
	// Allows callers to disable the promise rejection delay that would normally happen, if this is unimportant to what's being tested. (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
	DisableRejectionDelay bool `json:"disableRejectionDelay,omitempty"`
}

// EnableWithParams -
func (c *FedCm) EnableWithParams(ctx context.Context, v *FedCmEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.enable", Params: v})
}

// Enable -
// disableRejectionDelay - Allows callers to disable the promise rejection delay that would normally happen, if this is unimportant to what's being tested. (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
func (c *FedCm) Enable(ctx context.Context, disableRejectionDelay bool) (*gcdmessage.ChromeResponse, error) {
	var v FedCmEnableParams
	v.DisableRejectionDelay = disableRejectionDelay
	return c.EnableWithParams(ctx, &v)
}

//
func (c *FedCm) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.disable"})
}

type FedCmSelectAccountParams struct {
	//
	DialogId string `json:"dialogId"`
	//
	AccountIndex int `json:"accountIndex"`
}

// SelectAccountWithParams -
func (c *FedCm) SelectAccountWithParams(ctx context.Context, v *FedCmSelectAccountParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.selectAccount", Params: v})
}

// SelectAccount -
// dialogId -
// accountIndex -
func (c *FedCm) SelectAccount(ctx context.Context, dialogId string, accountIndex int) (*gcdmessage.ChromeResponse, error) {
	var v FedCmSelectAccountParams
	v.DialogId = dialogId
	v.AccountIndex = accountIndex
	return c.SelectAccountWithParams(ctx, &v)
}

type FedCmClickDialogButtonParams struct {
	//
	DialogId string `json:"dialogId"`
	//  enum values: ConfirmIdpLoginContinue, ErrorGotIt, ErrorMoreDetails
	DialogButton string `json:"dialogButton"`
}

// ClickDialogButtonWithParams -
func (c *FedCm) ClickDialogButtonWithParams(ctx context.Context, v *FedCmClickDialogButtonParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.clickDialogButton", Params: v})
}

// ClickDialogButton -
// dialogId -
// dialogButton -  enum values: ConfirmIdpLoginContinue, ErrorGotIt, ErrorMoreDetails
func (c *FedCm) ClickDialogButton(ctx context.Context, dialogId string, dialogButton string) (*gcdmessage.ChromeResponse, error) {
	var v FedCmClickDialogButtonParams
	v.DialogId = dialogId
	v.DialogButton = dialogButton
	return c.ClickDialogButtonWithParams(ctx, &v)
}

type FedCmOpenUrlParams struct {
	//
	DialogId string `json:"dialogId"`
	//
	AccountIndex int `json:"accountIndex"`
	//  enum values: TermsOfService, PrivacyPolicy
	AccountUrlType string `json:"accountUrlType"`
}

// OpenUrlWithParams -
func (c *FedCm) OpenUrlWithParams(ctx context.Context, v *FedCmOpenUrlParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.openUrl", Params: v})
}

// OpenUrl -
// dialogId -
// accountIndex -
// accountUrlType -  enum values: TermsOfService, PrivacyPolicy
func (c *FedCm) OpenUrl(ctx context.Context, dialogId string, accountIndex int, accountUrlType string) (*gcdmessage.ChromeResponse, error) {
	var v FedCmOpenUrlParams
	v.DialogId = dialogId
	v.AccountIndex = accountIndex
	v.AccountUrlType = accountUrlType
	return c.OpenUrlWithParams(ctx, &v)
}

type FedCmDismissDialogParams struct {
	//
	DialogId string `json:"dialogId"`
	//
	TriggerCooldown bool `json:"triggerCooldown,omitempty"`
}

// DismissDialogWithParams -
func (c *FedCm) DismissDialogWithParams(ctx context.Context, v *FedCmDismissDialogParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.dismissDialog", Params: v})
}

// DismissDialog -
// dialogId -
// triggerCooldown -
func (c *FedCm) DismissDialog(ctx context.Context, dialogId string, triggerCooldown bool) (*gcdmessage.ChromeResponse, error) {
	var v FedCmDismissDialogParams
	v.DialogId = dialogId
	v.TriggerCooldown = triggerCooldown
	return c.DismissDialogWithParams(ctx, &v)
}

// Resets the cooldown time, if any, to allow the next FedCM call to show a dialog even if one was recently dismissed by the user.
func (c *FedCm) ResetCooldown(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "FedCm.resetCooldown"})
}
