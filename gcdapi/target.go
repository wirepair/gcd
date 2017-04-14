// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Target functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type TargetTargetInfo struct {
	TargetId string `json:"targetId"` //
	Type     string `json:"type"`     //
	Title    string `json:"title"`    //
	Url      string `json:"url"`      //
}

// No Description.
type TargetRemoteLocation struct {
	Host string `json:"host"` //
	Port int    `json:"port"` //
}

// Issued when a possible inspection target is created.
type TargetTargetCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo *TargetTargetInfo `json:"targetInfo"` //
	} `json:"Params,omitempty"`
}

// Issued when a target is destroyed.
type TargetTargetDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
	} `json:"Params,omitempty"`
}

// Issued when attached to target because of auto-attach or <code>attachToTarget</code> command.
type TargetAttachedToTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo         *TargetTargetInfo `json:"targetInfo"`         //
		WaitingForDebugger bool              `json:"waitingForDebugger"` //
	} `json:"Params,omitempty"`
}

// Issued when detached from target for any reason (including <code>detachFromTarget</code> command).
type TargetDetachedFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
	} `json:"Params,omitempty"`
}

// Notifies about new protocol message from attached target.
type TargetReceivedMessageFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
		Message  string `json:"message"`  //
	} `json:"Params,omitempty"`
}

type Target struct {
	target gcdmessage.ChromeTargeter
}

func NewTarget(target gcdmessage.ChromeTargeter) *Target {
	c := &Target{target: target}
	return c
}

// SetDiscoverTargets - Controls whether to discover available targets and notify via <code>targetCreated/targetDestroyed</code> events.
// discover - Whether to discover available targets.
func (c *Target) SetDiscoverTargets(discover bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["discover"] = discover
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setDiscoverTargets", Params: paramRequest})
}

// SetAutoAttach - Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
// autoAttach - Whether to auto-attach to related targets.
// waitForDebuggerOnStart - Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
func (c *Target) SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["autoAttach"] = autoAttach
	paramRequest["waitForDebuggerOnStart"] = waitForDebuggerOnStart
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setAutoAttach", Params: paramRequest})
}

// SetAttachToFrames -
// value - Whether to attach to frames.
func (c *Target) SetAttachToFrames(value bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setAttachToFrames", Params: paramRequest})
}

// SetRemoteLocations - Enables target discovery for the specified locations, when <code>setDiscoverTargets</code> was set to <code>true</code>.
// locations - List of remote locations.
func (c *Target) SetRemoteLocations(locations *TargetRemoteLocation) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["locations"] = locations
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setRemoteLocations", Params: paramRequest})
}

// SendMessageToTarget - Sends protocol message to the target with given id.
// targetId -
// message -
func (c *Target) SendMessageToTarget(targetId string, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["targetId"] = targetId
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.sendMessageToTarget", Params: paramRequest})
}

// GetTargetInfo - Returns information about a target.
// targetId -
// Returns -  targetInfo -
func (c *Target) GetTargetInfo(targetId string) (*TargetTargetInfo, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargetInfo", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfo *TargetTargetInfo
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

// ActivateTarget - Activates (focuses) the target.
// targetId -
func (c *Target) ActivateTarget(targetId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.activateTarget", Params: paramRequest})
}

// CloseTarget - Closes the target. If the target is a page that gets closed too.
// targetId -
// Returns -  success -
func (c *Target) CloseTarget(targetId string) (bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.closeTarget", Params: paramRequest})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// AttachToTarget - Attaches to the target with given id.
// targetId -
// Returns -  success - Whether attach succeeded.
func (c *Target) AttachToTarget(targetId string) (bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.attachToTarget", Params: paramRequest})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// DetachFromTarget - Detaches from the target with given id.
// targetId -
func (c *Target) DetachFromTarget(targetId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["targetId"] = targetId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.detachFromTarget", Params: paramRequest})
}

// CreateBrowserContext - Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
// Returns -  browserContextId - The id of the context created.
func (c *Target) CreateBrowserContext() (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createBrowserContext"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			BrowserContextId string
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

	return chromeData.Result.BrowserContextId, nil
}

// DisposeBrowserContext - Deletes a BrowserContext, will fail of any open page uses it.
// browserContextId -
// Returns -  success -
func (c *Target) DisposeBrowserContext(browserContextId string) (bool, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["browserContextId"] = browserContextId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.disposeBrowserContext", Params: paramRequest})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// CreateTarget - Creates a new page.
// url - The initial URL the page will be navigated to.
// width - Frame width in DIP (headless chrome only).
// height - Frame height in DIP (headless chrome only).
// browserContextId - The browser context to create the page in (headless chrome only).
// Returns -  targetId - The id of the page opened.
func (c *Target) CreateTarget(url string, width int, height int, browserContextId string) (string, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["url"] = url
	paramRequest["width"] = width
	paramRequest["height"] = height
	paramRequest["browserContextId"] = browserContextId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createTarget", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			TargetId string
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

	return chromeData.Result.TargetId, nil
}

// GetTargets - Retrieves a list of available targets.
// Returns -  targetInfos - The list of targets.
func (c *Target) GetTargets() ([]*TargetTargetInfo, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargets"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfos []*TargetTargetInfo
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

	return chromeData.Result.TargetInfos, nil
}
