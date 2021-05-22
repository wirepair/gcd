// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Target functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type TargetTargetInfo struct {
	TargetId         string `json:"targetId"`                   //
	Type             string `json:"type"`                       //
	Title            string `json:"title"`                      //
	Url              string `json:"url"`                        //
	Attached         bool   `json:"attached"`                   // Whether the target has an attached client.
	OpenerId         string `json:"openerId,omitempty"`         // Opener target Id
	CanAccessOpener  bool   `json:"canAccessOpener"`            // Whether the target has access to the originating window.
	OpenerFrameId    string `json:"openerFrameId,omitempty"`    // Frame id of originating window (is only set if target has an opener).
	BrowserContextId string `json:"browserContextId,omitempty"` //
}

// No Description.
type TargetRemoteLocation struct {
	Host string `json:"host"` //
	Port int    `json:"port"` //
}

// Issued when attached to target because of auto-attach or `attachToTarget` command.
type TargetAttachedToTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		SessionId          string            `json:"sessionId"`          // Identifier assigned to the session used to send/receive messages.
		TargetInfo         *TargetTargetInfo `json:"targetInfo"`         //
		WaitingForDebugger bool              `json:"waitingForDebugger"` //
	} `json:"Params,omitempty"`
}

// Issued when detached from target for any reason (including `detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been attached to it.
type TargetDetachedFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		SessionId string `json:"sessionId"`          // Detached session identifier.
		TargetId  string `json:"targetId,omitempty"` // Deprecated.
	} `json:"Params,omitempty"`
}

// Notifies about a new protocol message received from the session (as reported in `attachedToTarget` event).
type TargetReceivedMessageFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		SessionId string `json:"sessionId"`          // Identifier of a session which sends a message.
		Message   string `json:"message"`            //
		TargetId  string `json:"targetId,omitempty"` // Deprecated.
	} `json:"Params,omitempty"`
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

// Issued when a target has crashed.
type TargetTargetCrashedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId  string `json:"targetId"`  //
		Status    string `json:"status"`    // Termination status type.
		ErrorCode int    `json:"errorCode"` // Termination error code.
	} `json:"Params,omitempty"`
}

// Issued when some information about a target has changed. This only happens between `targetCreated` and `targetDestroyed`.
type TargetTargetInfoChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo *TargetTargetInfo `json:"targetInfo"` //
	} `json:"Params,omitempty"`
}

type Target struct {
	target gcdmessage.ChromeTargeter
}

func NewTarget(target gcdmessage.ChromeTargeter) *Target {
	c := &Target{target: target}
	return c
}

type TargetActivateTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// ActivateTargetWithParams - Activates (focuses) the target.
func (c *Target) ActivateTargetWithParams(ctx context.Context, v *TargetActivateTargetParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.activateTarget", Params: v})
}

// ActivateTarget - Activates (focuses) the target.
// targetId -
func (c *Target) ActivateTarget(ctx context.Context, targetId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetActivateTargetParams
	v.TargetId = targetId
	return c.ActivateTargetWithParams(ctx, &v)
}

type TargetAttachToTargetParams struct {
	//
	TargetId string `json:"targetId"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// AttachToTargetWithParams - Attaches to the target with given id.
// Returns -  sessionId - Id assigned to the session.
func (c *Target) AttachToTargetWithParams(ctx context.Context, v *TargetAttachToTargetParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.attachToTarget", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SessionId string
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

	return chromeData.Result.SessionId, nil
}

// AttachToTarget - Attaches to the target with given id.
// targetId -
// flatten - Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
// Returns -  sessionId - Id assigned to the session.
func (c *Target) AttachToTarget(ctx context.Context, targetId string, flatten bool) (string, error) {
	var v TargetAttachToTargetParams
	v.TargetId = targetId
	v.Flatten = flatten
	return c.AttachToTargetWithParams(ctx, &v)
}

// AttachToBrowserTarget - Attaches to the browser target, only uses flat sessionId mode.
// Returns -  sessionId - Id assigned to the session.
func (c *Target) AttachToBrowserTarget(ctx context.Context) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.attachToBrowserTarget"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SessionId string
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

	return chromeData.Result.SessionId, nil
}

type TargetCloseTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// CloseTargetWithParams - Closes the target. If the target is a page that gets closed too.
// Returns -  success - Always set to true. If an error occurs, the response indicates protocol error.
func (c *Target) CloseTargetWithParams(ctx context.Context, v *TargetCloseTargetParams) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.closeTarget", Params: v})
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

// CloseTarget - Closes the target. If the target is a page that gets closed too.
// targetId -
// Returns -  success - Always set to true. If an error occurs, the response indicates protocol error.
func (c *Target) CloseTarget(ctx context.Context, targetId string) (bool, error) {
	var v TargetCloseTargetParams
	v.TargetId = targetId
	return c.CloseTargetWithParams(ctx, &v)
}

type TargetExposeDevToolsProtocolParams struct {
	//
	TargetId string `json:"targetId"`
	// Binding name, 'cdp' if not specified.
	BindingName string `json:"bindingName,omitempty"`
}

// ExposeDevToolsProtocolWithParams - Inject object to the target's main frame that provides a communication channel with browser target.  Injected object will be available as `window[bindingName]`.  The object has the follwing API: - `binding.send(json)` - a method to send messages over the remote debugging protocol - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
func (c *Target) ExposeDevToolsProtocolWithParams(ctx context.Context, v *TargetExposeDevToolsProtocolParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.exposeDevToolsProtocol", Params: v})
}

// ExposeDevToolsProtocol - Inject object to the target's main frame that provides a communication channel with browser target.  Injected object will be available as `window[bindingName]`.  The object has the follwing API: - `binding.send(json)` - a method to send messages over the remote debugging protocol - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
// targetId -
// bindingName - Binding name, 'cdp' if not specified.
func (c *Target) ExposeDevToolsProtocol(ctx context.Context, targetId string, bindingName string) (*gcdmessage.ChromeResponse, error) {
	var v TargetExposeDevToolsProtocolParams
	v.TargetId = targetId
	v.BindingName = bindingName
	return c.ExposeDevToolsProtocolWithParams(ctx, &v)
}

type TargetCreateBrowserContextParams struct {
	// If specified, disposes this context when debugging session disconnects.
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"`
	// Proxy server, similar to the one passed to --proxy-server
	ProxyServer string `json:"proxyServer,omitempty"`
	// Proxy bypass list, similar to the one passed to --proxy-bypass-list
	ProxyBypassList string `json:"proxyBypassList,omitempty"`
}

// CreateBrowserContextWithParams - Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
// Returns -  browserContextId - The id of the context created.
func (c *Target) CreateBrowserContextWithParams(ctx context.Context, v *TargetCreateBrowserContextParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createBrowserContext", Params: v})
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

// CreateBrowserContext - Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
// disposeOnDetach - If specified, disposes this context when debugging session disconnects.
// proxyServer - Proxy server, similar to the one passed to --proxy-server
// proxyBypassList - Proxy bypass list, similar to the one passed to --proxy-bypass-list
// Returns -  browserContextId - The id of the context created.
func (c *Target) CreateBrowserContext(ctx context.Context, disposeOnDetach bool, proxyServer string, proxyBypassList string) (string, error) {
	var v TargetCreateBrowserContextParams
	v.DisposeOnDetach = disposeOnDetach
	v.ProxyServer = proxyServer
	v.ProxyBypassList = proxyBypassList
	return c.CreateBrowserContextWithParams(ctx, &v)
}

// GetBrowserContexts - Returns all browser contexts created with `Target.createBrowserContext` method.
// Returns -  browserContextIds - An array of browser context ids.
func (c *Target) GetBrowserContexts(ctx context.Context) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getBrowserContexts"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			BrowserContextIds []string
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

	return chromeData.Result.BrowserContextIds, nil
}

type TargetCreateTargetParams struct {
	// The initial URL the page will be navigated to. An empty string indicates about:blank.
	Url string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width int `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height int `json:"height,omitempty"`
	// The browser context to create the page in.
	BrowserContextId string `json:"browserContextId,omitempty"`
	// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
	// Whether to create a new Window or Tab (chrome-only, false by default).
	NewWindow bool `json:"newWindow,omitempty"`
	// Whether to create the target in background or foreground (chrome-only, false by default).
	Background bool `json:"background,omitempty"`
}

// CreateTargetWithParams - Creates a new page.
// Returns -  targetId - The id of the page opened.
func (c *Target) CreateTargetWithParams(ctx context.Context, v *TargetCreateTargetParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createTarget", Params: v})
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

// CreateTarget - Creates a new page.
// url - The initial URL the page will be navigated to. An empty string indicates about:blank.
// width - Frame width in DIP (headless chrome only).
// height - Frame height in DIP (headless chrome only).
// browserContextId - The browser context to create the page in.
// enableBeginFrameControl - Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
// newWindow - Whether to create a new Window or Tab (chrome-only, false by default).
// background - Whether to create the target in background or foreground (chrome-only, false by default).
// Returns -  targetId - The id of the page opened.
func (c *Target) CreateTarget(ctx context.Context, url string, width int, height int, browserContextId string, enableBeginFrameControl bool, newWindow bool, background bool) (string, error) {
	var v TargetCreateTargetParams
	v.Url = url
	v.Width = width
	v.Height = height
	v.BrowserContextId = browserContextId
	v.EnableBeginFrameControl = enableBeginFrameControl
	v.NewWindow = newWindow
	v.Background = background
	return c.CreateTargetWithParams(ctx, &v)
}

type TargetDetachFromTargetParams struct {
	// Session to detach.
	SessionId string `json:"sessionId,omitempty"`
	// Deprecated.
	TargetId string `json:"targetId,omitempty"`
}

// DetachFromTargetWithParams - Detaches session with given id.
func (c *Target) DetachFromTargetWithParams(ctx context.Context, v *TargetDetachFromTargetParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.detachFromTarget", Params: v})
}

// DetachFromTarget - Detaches session with given id.
// sessionId - Session to detach.
// targetId - Deprecated.
func (c *Target) DetachFromTarget(ctx context.Context, sessionId string, targetId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetDetachFromTargetParams
	v.SessionId = sessionId
	v.TargetId = targetId
	return c.DetachFromTargetWithParams(ctx, &v)
}

type TargetDisposeBrowserContextParams struct {
	//
	BrowserContextId string `json:"browserContextId"`
}

// DisposeBrowserContextWithParams - Deletes a BrowserContext. All the belonging pages will be closed without calling their beforeunload hooks.
func (c *Target) DisposeBrowserContextWithParams(ctx context.Context, v *TargetDisposeBrowserContextParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.disposeBrowserContext", Params: v})
}

// DisposeBrowserContext - Deletes a BrowserContext. All the belonging pages will be closed without calling their beforeunload hooks.
// browserContextId -
func (c *Target) DisposeBrowserContext(ctx context.Context, browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetDisposeBrowserContextParams
	v.BrowserContextId = browserContextId
	return c.DisposeBrowserContextWithParams(ctx, &v)
}

type TargetGetTargetInfoParams struct {
	//
	TargetId string `json:"targetId,omitempty"`
}

// GetTargetInfoWithParams - Returns information about a target.
// Returns -  targetInfo -
func (c *Target) GetTargetInfoWithParams(ctx context.Context, v *TargetGetTargetInfoParams) (*TargetTargetInfo, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargetInfo", Params: v})
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

// GetTargetInfo - Returns information about a target.
// targetId -
// Returns -  targetInfo -
func (c *Target) GetTargetInfo(ctx context.Context, targetId string) (*TargetTargetInfo, error) {
	var v TargetGetTargetInfoParams
	v.TargetId = targetId
	return c.GetTargetInfoWithParams(ctx, &v)
}

// GetTargets - Retrieves a list of available targets.
// Returns -  targetInfos - The list of targets.
func (c *Target) GetTargets(ctx context.Context) ([]*TargetTargetInfo, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargets"})
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

type TargetSendMessageToTargetParams struct {
	//
	Message string `json:"message"`
	// Identifier of the session.
	SessionId string `json:"sessionId,omitempty"`
	// Deprecated.
	TargetId string `json:"targetId,omitempty"`
}

// SendMessageToTargetWithParams - Sends protocol message over session with given id. Consider using flat mode instead; see commands attachToTarget, setAutoAttach, and crbug.com/991325.
func (c *Target) SendMessageToTargetWithParams(ctx context.Context, v *TargetSendMessageToTargetParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.sendMessageToTarget", Params: v})
}

// SendMessageToTarget - Sends protocol message over session with given id. Consider using flat mode instead; see commands attachToTarget, setAutoAttach, and crbug.com/991325.
// message -
// sessionId - Identifier of the session.
// targetId - Deprecated.
func (c *Target) SendMessageToTarget(ctx context.Context, message string, sessionId string, targetId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetSendMessageToTargetParams
	v.Message = message
	v.SessionId = sessionId
	v.TargetId = targetId
	return c.SendMessageToTargetWithParams(ctx, &v)
}

type TargetSetAutoAttachParams struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger` to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// SetAutoAttachWithParams - Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (c *Target) SetAutoAttachWithParams(ctx context.Context, v *TargetSetAutoAttachParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setAutoAttach", Params: v})
}

// SetAutoAttach - Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
// autoAttach - Whether to auto-attach to related targets.
// waitForDebuggerOnStart - Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger` to run paused targets.
// flatten - Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
func (c *Target) SetAutoAttach(ctx context.Context, autoAttach bool, waitForDebuggerOnStart bool, flatten bool) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetAutoAttachParams
	v.AutoAttach = autoAttach
	v.WaitForDebuggerOnStart = waitForDebuggerOnStart
	v.Flatten = flatten
	return c.SetAutoAttachWithParams(ctx, &v)
}

type TargetSetDiscoverTargetsParams struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

// SetDiscoverTargetsWithParams - Controls whether to discover available targets and notify via `targetCreated/targetInfoChanged/targetDestroyed` events.
func (c *Target) SetDiscoverTargetsWithParams(ctx context.Context, v *TargetSetDiscoverTargetsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setDiscoverTargets", Params: v})
}

// SetDiscoverTargets - Controls whether to discover available targets and notify via `targetCreated/targetInfoChanged/targetDestroyed` events.
// discover - Whether to discover available targets.
func (c *Target) SetDiscoverTargets(ctx context.Context, discover bool) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetDiscoverTargetsParams
	v.Discover = discover
	return c.SetDiscoverTargetsWithParams(ctx, &v)
}

type TargetSetRemoteLocationsParams struct {
	// List of remote locations.
	Locations []*TargetRemoteLocation `json:"locations"`
}

// SetRemoteLocationsWithParams - Enables target discovery for the specified locations, when `setDiscoverTargets` was set to `true`.
func (c *Target) SetRemoteLocationsWithParams(ctx context.Context, v *TargetSetRemoteLocationsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setRemoteLocations", Params: v})
}

// SetRemoteLocations - Enables target discovery for the specified locations, when `setDiscoverTargets` was set to `true`.
// locations - List of remote locations.
func (c *Target) SetRemoteLocations(ctx context.Context, locations []*TargetRemoteLocation) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetRemoteLocationsParams
	v.Locations = locations
	return c.SetRemoteLocationsWithParams(ctx, &v)
}
