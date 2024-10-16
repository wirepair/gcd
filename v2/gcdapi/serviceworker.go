// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains ServiceWorker functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// ServiceWorker registration.
type ServiceWorkerServiceWorkerRegistration struct {
	RegistrationId string `json:"registrationId"` //
	ScopeURL       string `json:"scopeURL"`       //
	IsDeleted      bool   `json:"isDeleted"`      //
}

// ServiceWorker version.
type ServiceWorkerServiceWorkerVersion struct {
	VersionId          string   `json:"versionId"`                    //
	RegistrationId     string   `json:"registrationId"`               //
	ScriptURL          string   `json:"scriptURL"`                    //
	RunningStatus      string   `json:"runningStatus"`                //  enum values: stopped, starting, running, stopping
	Status             string   `json:"status"`                       //  enum values: new, installing, installed, activating, activated, redundant
	ScriptLastModified float64  `json:"scriptLastModified,omitempty"` // The Last-Modified header value of the main script.
	ScriptResponseTime float64  `json:"scriptResponseTime,omitempty"` // The time at which the response headers of the main script were received from the server. For cached script it is the last time the cache entry was validated.
	ControlledClients  []string `json:"controlledClients,omitempty"`  //
	TargetId           string   `json:"targetId,omitempty"`           //
	RouterRules        string   `json:"routerRules,omitempty"`        //
}

// ServiceWorker error message.
type ServiceWorkerServiceWorkerErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`   //
	RegistrationId string `json:"registrationId"` //
	VersionId      string `json:"versionId"`      //
	SourceURL      string `json:"sourceURL"`      //
	LineNumber     int    `json:"lineNumber"`     //
	ColumnNumber   int    `json:"columnNumber"`   //
}

//
type ServiceWorkerWorkerErrorReportedEvent struct {
	Method string `json:"method"`
	Params struct {
		ErrorMessage *ServiceWorkerServiceWorkerErrorMessage `json:"errorMessage"` //
	} `json:"Params,omitempty"`
}

//
type ServiceWorkerWorkerRegistrationUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Registrations []*ServiceWorkerServiceWorkerRegistration `json:"registrations"` //
	} `json:"Params,omitempty"`
}

//
type ServiceWorkerWorkerVersionUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Versions []*ServiceWorkerServiceWorkerVersion `json:"versions"` //
	} `json:"Params,omitempty"`
}

type ServiceWorker struct {
	target gcdmessage.ChromeTargeter
}

func NewServiceWorker(target gcdmessage.ChromeTargeter) *ServiceWorker {
	c := &ServiceWorker{target: target}
	return c
}

type ServiceWorkerDeliverPushMessageParams struct {
	//
	Origin string `json:"origin"`
	//
	RegistrationId string `json:"registrationId"`
	//
	Data string `json:"data"`
}

// DeliverPushMessageWithParams -
func (c *ServiceWorker) DeliverPushMessageWithParams(ctx context.Context, v *ServiceWorkerDeliverPushMessageParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.deliverPushMessage", Params: v})
}

// DeliverPushMessage -
// origin -
// registrationId -
// data -
func (c *ServiceWorker) DeliverPushMessage(ctx context.Context, origin string, registrationId string, data string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerDeliverPushMessageParams
	v.Origin = origin
	v.RegistrationId = registrationId
	v.Data = data
	return c.DeliverPushMessageWithParams(ctx, &v)
}

//
func (c *ServiceWorker) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.disable"})
}

type ServiceWorkerDispatchSyncEventParams struct {
	//
	Origin string `json:"origin"`
	//
	RegistrationId string `json:"registrationId"`
	//
	Tag string `json:"tag"`
	//
	LastChance bool `json:"lastChance"`
}

// DispatchSyncEventWithParams -
func (c *ServiceWorker) DispatchSyncEventWithParams(ctx context.Context, v *ServiceWorkerDispatchSyncEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.dispatchSyncEvent", Params: v})
}

// DispatchSyncEvent -
// origin -
// registrationId -
// tag -
// lastChance -
func (c *ServiceWorker) DispatchSyncEvent(ctx context.Context, origin string, registrationId string, tag string, lastChance bool) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerDispatchSyncEventParams
	v.Origin = origin
	v.RegistrationId = registrationId
	v.Tag = tag
	v.LastChance = lastChance
	return c.DispatchSyncEventWithParams(ctx, &v)
}

type ServiceWorkerDispatchPeriodicSyncEventParams struct {
	//
	Origin string `json:"origin"`
	//
	RegistrationId string `json:"registrationId"`
	//
	Tag string `json:"tag"`
}

// DispatchPeriodicSyncEventWithParams -
func (c *ServiceWorker) DispatchPeriodicSyncEventWithParams(ctx context.Context, v *ServiceWorkerDispatchPeriodicSyncEventParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.dispatchPeriodicSyncEvent", Params: v})
}

// DispatchPeriodicSyncEvent -
// origin -
// registrationId -
// tag -
func (c *ServiceWorker) DispatchPeriodicSyncEvent(ctx context.Context, origin string, registrationId string, tag string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerDispatchPeriodicSyncEventParams
	v.Origin = origin
	v.RegistrationId = registrationId
	v.Tag = tag
	return c.DispatchPeriodicSyncEventWithParams(ctx, &v)
}

//
func (c *ServiceWorker) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.enable"})
}

type ServiceWorkerInspectWorkerParams struct {
	//
	VersionId string `json:"versionId"`
}

// InspectWorkerWithParams -
func (c *ServiceWorker) InspectWorkerWithParams(ctx context.Context, v *ServiceWorkerInspectWorkerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.inspectWorker", Params: v})
}

// InspectWorker -
// versionId -
func (c *ServiceWorker) InspectWorker(ctx context.Context, versionId string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerInspectWorkerParams
	v.VersionId = versionId
	return c.InspectWorkerWithParams(ctx, &v)
}

type ServiceWorkerSetForceUpdateOnPageLoadParams struct {
	//
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

// SetForceUpdateOnPageLoadWithParams -
func (c *ServiceWorker) SetForceUpdateOnPageLoadWithParams(ctx context.Context, v *ServiceWorkerSetForceUpdateOnPageLoadParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.setForceUpdateOnPageLoad", Params: v})
}

// SetForceUpdateOnPageLoad -
// forceUpdateOnPageLoad -
func (c *ServiceWorker) SetForceUpdateOnPageLoad(ctx context.Context, forceUpdateOnPageLoad bool) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerSetForceUpdateOnPageLoadParams
	v.ForceUpdateOnPageLoad = forceUpdateOnPageLoad
	return c.SetForceUpdateOnPageLoadWithParams(ctx, &v)
}

type ServiceWorkerSkipWaitingParams struct {
	//
	ScopeURL string `json:"scopeURL"`
}

// SkipWaitingWithParams -
func (c *ServiceWorker) SkipWaitingWithParams(ctx context.Context, v *ServiceWorkerSkipWaitingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.skipWaiting", Params: v})
}

// SkipWaiting -
// scopeURL -
func (c *ServiceWorker) SkipWaiting(ctx context.Context, scopeURL string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerSkipWaitingParams
	v.ScopeURL = scopeURL
	return c.SkipWaitingWithParams(ctx, &v)
}

type ServiceWorkerStartWorkerParams struct {
	//
	ScopeURL string `json:"scopeURL"`
}

// StartWorkerWithParams -
func (c *ServiceWorker) StartWorkerWithParams(ctx context.Context, v *ServiceWorkerStartWorkerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.startWorker", Params: v})
}

// StartWorker -
// scopeURL -
func (c *ServiceWorker) StartWorker(ctx context.Context, scopeURL string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerStartWorkerParams
	v.ScopeURL = scopeURL
	return c.StartWorkerWithParams(ctx, &v)
}

//
func (c *ServiceWorker) StopAllWorkers(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.stopAllWorkers"})
}

type ServiceWorkerStopWorkerParams struct {
	//
	VersionId string `json:"versionId"`
}

// StopWorkerWithParams -
func (c *ServiceWorker) StopWorkerWithParams(ctx context.Context, v *ServiceWorkerStopWorkerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.stopWorker", Params: v})
}

// StopWorker -
// versionId -
func (c *ServiceWorker) StopWorker(ctx context.Context, versionId string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerStopWorkerParams
	v.VersionId = versionId
	return c.StopWorkerWithParams(ctx, &v)
}

type ServiceWorkerUnregisterParams struct {
	//
	ScopeURL string `json:"scopeURL"`
}

// UnregisterWithParams -
func (c *ServiceWorker) UnregisterWithParams(ctx context.Context, v *ServiceWorkerUnregisterParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.unregister", Params: v})
}

// Unregister -
// scopeURL -
func (c *ServiceWorker) Unregister(ctx context.Context, scopeURL string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerUnregisterParams
	v.ScopeURL = scopeURL
	return c.UnregisterWithParams(ctx, &v)
}

type ServiceWorkerUpdateRegistrationParams struct {
	//
	ScopeURL string `json:"scopeURL"`
}

// UpdateRegistrationWithParams -
func (c *ServiceWorker) UpdateRegistrationWithParams(ctx context.Context, v *ServiceWorkerUpdateRegistrationParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.updateRegistration", Params: v})
}

// UpdateRegistration -
// scopeURL -
func (c *ServiceWorker) UpdateRegistration(ctx context.Context, scopeURL string) (*gcdmessage.ChromeResponse, error) {
	var v ServiceWorkerUpdateRegistrationParams
	v.ScopeURL = scopeURL
	return c.UpdateRegistrationWithParams(ctx, &v)
}
