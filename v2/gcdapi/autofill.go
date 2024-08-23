// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Autofill functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type AutofillCreditCard struct {
	Number      string `json:"number"`      // 16-digit credit card number.
	Name        string `json:"name"`        // Name of the credit card owner.
	ExpiryMonth string `json:"expiryMonth"` // 2-digit expiry month.
	ExpiryYear  string `json:"expiryYear"`  // 4-digit expiry year.
	Cvc         string `json:"cvc"`         // 3-digit card verification code.
}

// No Description.
type AutofillAddressField struct {
	Name  string `json:"name"`  // address field name, for example GIVEN_NAME.
	Value string `json:"value"` // address field value, for example Jon Doe.
}

// A list of address fields.
type AutofillAddressFields struct {
	Fields []*AutofillAddressField `json:"fields"` //
}

// No Description.
type AutofillAddress struct {
	Fields []*AutofillAddressField `json:"fields"` // fields and values defining an address.
}

// Defines how an address can be displayed like in chrome://settings/addresses. Address UI is a two dimensional array, each inner array is an "address information line", and when rendered in a UI surface should be displayed as such. The following address UI for instance: [[{name: "GIVE_NAME", value: "Jon"}, {name: "FAMILY_NAME", value: "Doe"}], [{name: "CITY", value: "Munich"}, {name: "ZIP", value: "81456"}]] should allow the receiver to render: Jon Doe Munich 81456
type AutofillAddressUI struct {
	AddressFields []*AutofillAddressFields `json:"addressFields"` // A two dimension array containing the representation of values from an address profile.
}

// No Description.
type AutofillFilledField struct {
	HtmlType        string `json:"htmlType"`        // The type of the field, e.g text, password etc.
	Id              string `json:"id"`              // the html id
	Name            string `json:"name"`            // the html name
	Value           string `json:"value"`           // the field value
	AutofillType    string `json:"autofillType"`    // The actual field type, e.g FAMILY_NAME
	FillingStrategy string `json:"fillingStrategy"` // The filling strategy enum values: autocompleteAttribute, autofillInferred
	FrameId         string `json:"frameId"`         // The frame the field belongs to
	FieldId         int    `json:"fieldId"`         // The form field's DOM node
}

// Emitted when an address form is filled.
type AutofillAddressFormFilledEvent struct {
	Method string `json:"method"`
	Params struct {
		FilledFields []*AutofillFilledField `json:"filledFields"` // Information about the fields that were filled
		AddressUi    *AutofillAddressUI     `json:"addressUi"`    // An UI representation of the address used to fill the form. Consists of a 2D array where each child represents an address/profile line.
	} `json:"Params,omitempty"`
}

type Autofill struct {
	target gcdmessage.ChromeTargeter
}

func NewAutofill(target gcdmessage.ChromeTargeter) *Autofill {
	c := &Autofill{target: target}
	return c
}

type AutofillTriggerParams struct {
	// Identifies a field that serves as an anchor for autofill.
	FieldId int `json:"fieldId"`
	// Identifies the frame that field belongs to.
	FrameId string `json:"frameId,omitempty"`
	// Credit card information to fill out the form. Credit card data is not saved.
	Card *AutofillCreditCard `json:"card"`
}

// TriggerWithParams - Trigger autofill on a form identified by the fieldId. If the field and related form cannot be autofilled, returns an error.
func (c *Autofill) TriggerWithParams(ctx context.Context, v *AutofillTriggerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Autofill.trigger", Params: v})
}

// Trigger - Trigger autofill on a form identified by the fieldId. If the field and related form cannot be autofilled, returns an error.
// fieldId - Identifies a field that serves as an anchor for autofill.
// frameId - Identifies the frame that field belongs to.
// card - Credit card information to fill out the form. Credit card data is not saved.
func (c *Autofill) Trigger(ctx context.Context, fieldId int, frameId string, card *AutofillCreditCard) (*gcdmessage.ChromeResponse, error) {
	var v AutofillTriggerParams
	v.FieldId = fieldId
	v.FrameId = frameId
	v.Card = card
	return c.TriggerWithParams(ctx, &v)
}

type AutofillSetAddressesParams struct {
	//
	Addresses []*AutofillAddress `json:"addresses"`
}

// SetAddressesWithParams - Set addresses so that developers can verify their forms implementation.
func (c *Autofill) SetAddressesWithParams(ctx context.Context, v *AutofillSetAddressesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Autofill.setAddresses", Params: v})
}

// SetAddresses - Set addresses so that developers can verify their forms implementation.
// addresses -
func (c *Autofill) SetAddresses(ctx context.Context, addresses []*AutofillAddress) (*gcdmessage.ChromeResponse, error) {
	var v AutofillSetAddressesParams
	v.Addresses = addresses
	return c.SetAddressesWithParams(ctx, &v)
}

// Disables autofill domain notifications.
func (c *Autofill) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Autofill.disable"})
}

// Enables autofill domain notifications.
func (c *Autofill) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Autofill.enable"})
}
