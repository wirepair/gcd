// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ServiceWorker commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) ServiceWorker() *ChromeServiceWorker {
	if c.serviceworker == nil {
		c.serviceworker = newChromeServiceWorker(c)
	}
	return c.serviceworker
}

type ChromeServiceWorker struct {
	target *ChromeTarget
}

func newChromeServiceWorker(target *ChromeTarget) *ChromeServiceWorker {
	c := &ChromeServiceWorker{target: target}
	return c
}

//
func (c *ChromeServiceWorker) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.enable"})
}

//
func (c *ChromeServiceWorker) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.disable"})
}

// sendMessage -
// workerId -
// message -
func (c *ChromeServiceWorker) SendMessage(workerId string, message string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.sendMessage", Params: paramRequest})
}

// stop -
// workerId -
func (c *ChromeServiceWorker) Stop(workerId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.stop", Params: paramRequest})
}

// unregister -
// scopeURL -
func (c *ChromeServiceWorker) Unregister(scopeURL string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.unregister", Params: paramRequest})
}

// updateRegistration -
// scopeURL -
func (c *ChromeServiceWorker) UpdateRegistration(scopeURL string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.updateRegistration", Params: paramRequest})
}

// startWorker -
// scopeURL -
func (c *ChromeServiceWorker) StartWorker(scopeURL string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["scopeURL"] = scopeURL
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.startWorker", Params: paramRequest})
}

// stopWorker -
// versionId -
func (c *ChromeServiceWorker) StopWorker(versionId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["versionId"] = versionId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.stopWorker", Params: paramRequest})
}

// inspectWorker -
// versionId -
func (c *ChromeServiceWorker) InspectWorker(versionId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["versionId"] = versionId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.inspectWorker", Params: paramRequest})
}

// skipWaiting -
// versionId -
func (c *ChromeServiceWorker) SkipWaiting(versionId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["versionId"] = versionId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.skipWaiting", Params: paramRequest})
}

// setDebugOnStart -
// debugOnStart -
func (c *ChromeServiceWorker) SetDebugOnStart(debugOnStart bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["debugOnStart"] = debugOnStart
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.setDebugOnStart", Params: paramRequest})
}

// deliverPushMessage -
// origin -
// registrationId -
// data -
func (c *ChromeServiceWorker) DeliverPushMessage(origin string, registrationId string, data string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["origin"] = origin
	paramRequest["registrationId"] = registrationId
	paramRequest["data"] = data
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.deliverPushMessage", Params: paramRequest})
}

// activateTarget -
// targetId -
func (c *ChromeServiceWorker) ActivateTarget(targetId *types.ChromeServiceWorkerTargetID) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.activateTarget", Params: paramRequest})
}

// getTargetInfo -
// Returns -
func (c *ChromeServiceWorker) GetTargetInfo(targetId *types.ChromeServiceWorkerTargetID) (*types.ChromeServiceWorkerTargetInfo, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ServiceWorker.getTargetInfo", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			TargetInfo *types.ChromeServiceWorkerTargetInfo
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.TargetInfo, nil
}
