// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Browser functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type BrowserTargetInfo struct {
	TargetId string `json:"targetId"` //
	Type     string `json:"type"`     //
	Title    string `json:"title"`    //
	Url      string `json:"url"`      //
}

// Dispatches protocol message from the target with given id.
type BrowserDispatchMessageEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
		Message  string `json:"message"`  //
	} `json:"Params,omitempty"`
}

type Browser struct {
	target gcdmessage.ChromeTargeter
}

func NewBrowser(target gcdmessage.ChromeTargeter) *Browser {
	c := &Browser{target: target}
	return c
}

// GetTargets - Returns target information for all potential targets.
// Returns -  targetInfo -
func (c *Browser) GetTargets() ([]*BrowserTargetInfo, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getTargets"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfo []*BrowserTargetInfo
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

// Attach - Attaches to the target with given id.
// targetId - Target id.
func (c *Browser) Attach(targetId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.attach", Params: paramRequest})
}

// Detach - Detaches from the target with given id.
// targetId -
func (c *Browser) Detach(targetId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.detach", Params: paramRequest})
}

// SendMessage - Sends protocol message to the target with given id.
// targetId -
// message -
func (c *Browser) SendMessage(targetId string, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["targetId"] = targetId
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.sendMessage", Params: paramRequest})
}
