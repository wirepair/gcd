// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains ServiceWorker functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
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
	ScriptResponseTime float64  `json:"scriptResponseTime,omitempty"` // The time at which the response headers of the main script were received from the server.  For cached script it is the last time the cache entry was validated.
	ControlledClients  []string `json:"controlledClients,omitempty"`  //
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

// No Description.
type ServiceWorkerTargetInfo struct {
	Id    string `json:"id"`    //
	Type  string `json:"type"`  //
	Title string `json:"title"` //
	Url   string `json:"url"`   //
}

//
type ServiceWorkerWorkerCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId  string `json:"workerId"`  //
		Url       string `json:"url"`       //
		VersionId string `json:"versionId"` //
	} `json:"Params,omitempty"`
}

//
type ServiceWorkerWorkerTerminatedEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId string `json:"workerId"` //
	} `json:"Params,omitempty"`
}

//
type ServiceWorkerDispatchMessageEvent struct {
	Method string `json:"method"`
	Params struct {
		WorkerId string `json:"workerId"` //
		Message  string `json:"message"`  //
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

//
type ServiceWorkerWorkerErrorReportedEvent struct {
	Method string `json:"method"`
	Params struct {
		ErrorMessage *ServiceWorkerServiceWorkerErrorMessage `json:"errorMessage"` //
	} `json:"Params,omitempty"`
}

type ServiceWorker struct {
	target gcdmessage.ChromeTargeter
}

func NewServiceWorker(target gcdmessage.ChromeTargeter) *ServiceWorker {
	c := &ServiceWorker{target: target}
	return c
}

//
func (c *ServiceWorker) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.enable"})
}

//
func (c *ServiceWorker) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.disable"})
}

// SendMessage -
// workerId -
// message -
func (c *ServiceWorker) SendMessage(workerId string, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.sendMessage", Params: paramRequest})
}

// Stop -
// workerId -
func (c *ServiceWorker) Stop(workerId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.stop", Params: paramRequest})
}

// Unregister -
// scopeURL -
func (c *ServiceWorker) Unregister(scopeURL string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.unregister", Params: paramRequest})
}

// UpdateRegistration -
// scopeURL -
func (c *ServiceWorker) UpdateRegistration(scopeURL string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.updateRegistration", Params: paramRequest})
}

// StartWorker -
// scopeURL -
func (c *ServiceWorker) StartWorker(scopeURL string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.startWorker", Params: paramRequest})
}

// SkipWaiting -
// scopeURL -
func (c *ServiceWorker) SkipWaiting(scopeURL string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.skipWaiting", Params: paramRequest})
}

// StopWorker -
// versionId -
func (c *ServiceWorker) StopWorker(versionId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["versionId"] = versionId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.stopWorker", Params: paramRequest})
}

// InspectWorker -
// versionId -
func (c *ServiceWorker) InspectWorker(versionId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["versionId"] = versionId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.inspectWorker", Params: paramRequest})
}

// SetForceUpdateOnPageLoad -
// forceUpdateOnPageLoad -
func (c *ServiceWorker) SetForceUpdateOnPageLoad(forceUpdateOnPageLoad bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["forceUpdateOnPageLoad"] = forceUpdateOnPageLoad
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.setForceUpdateOnPageLoad", Params: paramRequest})
}

// DeliverPushMessage -
// origin -
// registrationId -
// data -
func (c *ServiceWorker) DeliverPushMessage(origin string, registrationId string, data string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["origin"] = origin
	paramRequest["registrationId"] = registrationId
	paramRequest["data"] = data
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.deliverPushMessage", Params: paramRequest})
}

// DispatchSyncEvent -
// origin -
// registrationId -
// tag -
// lastChance -
func (c *ServiceWorker) DispatchSyncEvent(origin string, registrationId string, tag string, lastChance bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["origin"] = origin
	paramRequest["registrationId"] = registrationId
	paramRequest["tag"] = tag
	paramRequest["lastChance"] = lastChance
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.dispatchSyncEvent", Params: paramRequest})
}

// GetTargetInfo -
// targetId -
// Returns -  targetInfo -
func (c *ServiceWorker) GetTargetInfo(targetId string) (*ServiceWorkerTargetInfo, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.getTargetInfo", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfo *ServiceWorkerTargetInfo
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

	return chromeData.Result.TargetInfo, nil
}

// ActivateTarget -
// targetId -
func (c *ServiceWorker) ActivateTarget(targetId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ServiceWorker.activateTarget", Params: paramRequest})
}
