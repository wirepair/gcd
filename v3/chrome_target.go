/*
The MIT License (MIT)

Copyright (c) 2020 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package gcd

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wirepair/gcd/v3/observer"

	"github.com/wirepair/gcd/v3/gcdapi"
	"github.com/wirepair/gcd/v3/gcdmessage"
)

// TargetInfo defines the 'tab' or target for this chrome instance,
// can be multiple and background processes are included (not just visual tabs)
type TargetInfo struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	FaviconUrl           string `json:"faviconUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

// responseHeader from websocket with id and method
type responseHeader struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

// devtoolsEventResponse from websocket with id and method and msg data to dispatch to subscribers
type devtoolsEventResponse struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
	Msg    []byte `json:"msg"`
}

// ChromeTarget (Tab/Process). Messages are returned to callers via non-buffered channels. Helpfully,
// the remote debugger service uses id's so we can correlate which request should match which response.
// We use a map to store the id of the request which contains a reference to a gcdmessage.Message that holds the
// reply channel for the ChromeTarget to return the response to.
// Events are handled by mapping the method name to a function which takes a target and byte output.
// For now, callers will need to unmarshall the types themselves.
type ChromeTarget struct {
	ctx    context.Context
	sendId int64 // An Id which is atomically incremented per request.
	// must be at top because of alignement and atomic usage
	replyLock       sync.RWMutex                                  // lock for dispatching responses
	replyDispatcher map[int64]chan *gcdmessage.Message            // Replies to synch methods using a non-buffered channel
	eventLock       sync.RWMutex                                  // lock for dispatching events
	eventDispatcher map[string]func(*ChromeTarget, int64, []byte) // calls the function when events match the subscribed method
	conn            *WebSocket                                    // the connection to the chrome debugger service for this tab/process

	// Chrome Debugger Domains
	Accessibility        *gcdapi.Accessibility
	Animation            *gcdapi.Animation
	ApplicationCache     *gcdapi.ApplicationCache // application cache API
	Audits               *gcdapi.Audits
	BackgroundService    *gcdapi.BackgroundService
	Browser              *gcdapi.Browser
	CacheStorage         *gcdapi.CacheStorage
	Cast                 *gcdapi.Cast
	Console              *gcdapi.Console           // console API
	CSS                  *gcdapi.CSS               // CSS API
	Database             *gcdapi.Database          // Database API
	Debugger             *gcdapi.Debugger          // JS Debugger API
	DeviceOrientation    *gcdapi.DeviceOrientation // Device Orientation API
	DOM                  *gcdapi.DOM               // DOM API
	DOMDebugger          *gcdapi.DOMDebugger       // DOM Debugger API
	DOMSnapshot          *gcdapi.DOMSnapshot
	DOMStorage           *gcdapi.DOMStorage // DOM Storage API
	Emulation            *gcdapi.Emulation
	Fetch                *gcdapi.Fetch
	HeadlessExperimental *gcdapi.HeadlessExperimental
	HeapProfiler         *gcdapi.HeapProfiler // HeapProfiler API
	IndexedDB            *gcdapi.IndexedDB    // IndexedDB API
	Input                *gcdapi.Input        // Why am i doing this, it's obvious what they are, I quit.
	Inspector            *gcdapi.Inspector
	IO                   *gcdapi.IO
	LayerTree            *gcdapi.LayerTree
	Log                  *gcdapi.Log
	Memory               *gcdapi.Memory
	Network              *gcdapi.Network
	Overlay              *gcdapi.Overlay
	Page                 *gcdapi.Page
	Performance          *gcdapi.Performance // if stable channel you'll need to uncomment
	PerformanceTimeline  *gcdapi.PerformanceTimeline
	Profiler             *gcdapi.Profiler
	Runtime              *gcdapi.Runtime
	Schema               *gcdapi.Schema
	Security             *gcdapi.Security
	ServiceWorker        *gcdapi.ServiceWorker
	Storage              *gcdapi.Storage
	SystemInfo           *gcdapi.SystemInfo
	TargetApi            *gcdapi.Target // buh name collision
	Tracing              *gcdapi.Tracing
	Tethering            *gcdapi.Tethering
	Media                *gcdapi.Media
	WebAudio             *gcdapi.WebAudio
	WebAuthn             *gcdapi.WebAuthn

	Target          *TargetInfo                 // The target information see, TargetInfo
	sendCh          chan *gcdmessage.Message    // The channel used for API components to send back to use
	eventCh         chan *devtoolsEventResponse // The channel used for dispatching devtool events
	doneCh          chan struct{}               // we be donez.
	apiTimeout      time.Duration               // A customizable timeout for waiting on Chrome to respond to us
	logger          Log
	debugger        *Gcd
	stopped         bool // we are/have shutdown
	messageObserver observer.MessageObserver
}

// openChromeTarget creates a new Chrome Target by connecting to the service given the URL taken from initial connection.
func openChromeTarget(debugger *Gcd, target *TargetInfo, observer observer.MessageObserver) (*ChromeTarget, error) {
	conn, err := wsConnection(debugger.ctx, target.WebSocketDebuggerUrl)
	if err != nil {
		return nil, err
	}

	chromeTarget := &ChromeTarget{
		conn:            conn,
		ctx:             debugger.ctx,
		Target:          target,
		apiTimeout:      120 * time.Second, // default 120 seconds to wait for chrome to respond to us
		sendCh:          make(chan *gcdmessage.Message),
		replyDispatcher: make(map[int64]chan *gcdmessage.Message),
		eventDispatcher: make(map[string]func(*ChromeTarget, int64, []byte)),
		eventCh:         make(chan *devtoolsEventResponse),
		doneCh:          make(chan struct{}),
		logger:          debugger.logger,
		debugger:        debugger,
		sendId:          0,
		messageObserver: observer,
	}

	chromeTarget.Init()
	chromeTarget.listen()
	return chromeTarget, nil
}

// Init all api objects
func (c *ChromeTarget) Init() {
	c.Accessibility = gcdapi.NewAccessibility(c)
	c.Animation = gcdapi.NewAnimation(c)
	c.ApplicationCache = gcdapi.NewApplicationCache(c)
	c.Audits = gcdapi.NewAudits(c)
	c.Browser = gcdapi.NewBrowser(c)
	c.BackgroundService = gcdapi.NewBackgroundService(c)
	c.CacheStorage = gcdapi.NewCacheStorage(c)
	c.Cast = gcdapi.NewCast(c)
	c.Console = gcdapi.NewConsole(c)
	c.CSS = gcdapi.NewCSS(c)
	c.Database = gcdapi.NewDatabase(c)
	c.Debugger = gcdapi.NewDebugger(c)
	c.DeviceOrientation = gcdapi.NewDeviceOrientation(c)
	c.DOMDebugger = gcdapi.NewDOMDebugger(c)
	c.DOM = gcdapi.NewDOM(c)
	c.DOMSnapshot = gcdapi.NewDOMSnapshot(c)
	c.DOMStorage = gcdapi.NewDOMStorage(c)
	c.Emulation = gcdapi.NewEmulation(c)
	c.HeapProfiler = gcdapi.NewHeapProfiler(c)
	c.IndexedDB = gcdapi.NewIndexedDB(c)
	c.Input = gcdapi.NewInput(c)
	c.Inspector = gcdapi.NewInspector(c)
	c.IO = gcdapi.NewIO(c)
	c.LayerTree = gcdapi.NewLayerTree(c)
	c.Memory = gcdapi.NewMemory(c)
	c.Log = gcdapi.NewLog(c)
	c.Network = gcdapi.NewNetwork(c)
	c.Overlay = gcdapi.NewOverlay(c)
	c.Page = gcdapi.NewPage(c)
	c.Profiler = gcdapi.NewProfiler(c)
	c.Runtime = gcdapi.NewRuntime(c)
	c.Schema = gcdapi.NewSchema(c)
	c.Security = gcdapi.NewSecurity(c)
	c.SystemInfo = gcdapi.NewSystemInfo(c)
	c.ServiceWorker = gcdapi.NewServiceWorker(c)
	c.TargetApi = gcdapi.NewTarget(c)
	c.Tracing = gcdapi.NewTracing(c)
	c.Tethering = gcdapi.NewTethering(c)
	c.HeadlessExperimental = gcdapi.NewHeadlessExperimental(c)
	c.Performance = gcdapi.NewPerformance(c)
	c.PerformanceTimeline = gcdapi.NewPerformanceTimeline(c)
	c.Fetch = gcdapi.NewFetch(c)
	c.Cast = gcdapi.NewCast(c)
	c.Media = gcdapi.NewMedia(c)
	c.WebAudio = gcdapi.NewWebAudio(c)
	c.WebAuthn = gcdapi.NewWebAuthn(c)
	c.BackgroundService = gcdapi.NewBackgroundService(c)
}

// clean up this target
func (c *ChromeTarget) shutdown() {
	if c.stopped == true {
		return
	}
	c.stopped = true

	// close websocket read/write goroutines
	close(c.doneCh)

	// close websocket connection
	c.conn.close()
}

// SetApiTimeout for how long we should wait before giving up gcdmessages.
// In the highly unusable (but it has occurred) event that chrome
// does not respond to one of our messages, we should be able to return
// from gcdmessage functions.
func (c *ChromeTarget) SetApiTimeout(timeout time.Duration) {
	c.apiTimeout = timeout
}

// GetApiTimeout used by gcdmessage.SendCustomReturn and gcdmessage.SendDefaultRequest
// to timeout an API call if chrome hasn't responded to us in apiTimeout
// time.
func (c *ChromeTarget) GetApiTimeout() time.Duration {
	return c.apiTimeout
}

// Subscribe Events, you must know the method name, such as Page.loadFiredEvent, and bind a function
// which takes a ChromeTarget (us) and the raw JSON byte data for that event.
func (c *ChromeTarget) Subscribe(method string, callback func(*ChromeTarget, int64, []byte)) {
	c.eventLock.Lock()
	c.eventDispatcher[method] = func(chromeTarget *ChromeTarget, id int64, data []byte) {
		c.messageObserver.Event(method, id, data)
		callback(chromeTarget, id, data)
	}
	c.eventLock.Unlock()
}

// Unsubscribe the handler for no longer receiving events.
func (c *ChromeTarget) Unsubscribe(method string) {
	c.eventLock.Lock()
	delete(c.eventDispatcher, method)
	c.eventLock.Unlock()
}

// Listens for API components wanting to send, and recv'ing data from the Chrome Debugger Service
func (c *ChromeTarget) listen() {
	go c.listenRead()
	go c.listenWrite()
	go c.dispatchEvents()
}

// Listens for API components wishing to send requests to the Chrome Debugger Service
func (c *ChromeTarget) listenWrite() {
	for {
		select {
		// send message to the browser debugger client
		case msg := <-c.sendCh:
			c.replyLock.Lock()
			c.replyDispatcher[msg.Id] = msg.ReplyCh
			c.replyLock.Unlock()

			c.logger.Println(msg.Id, " sending to chrome: ", string(msg.Data))
			err := c.conn.Send(msg.Data)
			if err != nil {
				c.logger.Println("error sending websocket message: ", err)
				return
			}
		// receive done from listenRead
		case <-c.doneCh:
			return
		case <-c.ctx.Done():
			return
		}
	}

}

// Listens for responses coming in from the Chrome Debugger Service.
func (c *ChromeTarget) listenRead() {
	for {
		msg, err := c.conn.Read()
		if err != nil {
			c.shutdown()
			return
		}
		c.dispatchResponse(msg)
	}
}

// dispatchResponse takes in the json message and extracts
// the id or method fields to dispatch either responses or events
// to listeners.
func (c *ChromeTarget) dispatchResponse(msg []byte) {
	f := &responseHeader{}
	err := json.Unmarshal(msg, f)
	if err != nil {
		c.logger.Println("error reading response data from chrome target", err)
	}

	c.replyLock.Lock()

	if r, ok := c.replyDispatcher[f.Id]; ok {
		delete(c.replyDispatcher, f.Id)
		c.replyLock.Unlock()
		c.logDebug("dispatching response id:", f.Id)
		go c.dispatchWithTimeout(r, f.Id, msg)
		return
	}
	c.replyLock.Unlock()

	c.checkTargetDisconnected(f.Method)

	c.eventLock.RLock()
	_, ok := c.eventDispatcher[f.Method]
	c.eventLock.RUnlock()

	if ok {
		c.eventCh <- &devtoolsEventResponse{Id: f.Id, Method: f.Method, Msg: msg}
		return
	}

	if c.debugger.debugEvents {
		c.logger.Println("no event reciever bound: ", f.Method, " data: ", string(msg))
	}
}

func (c *ChromeTarget) dispatchEvents() {
	for {
		select {
		case <-c.doneCh:
			return
		case <-c.ctx.Done():
			return
		case m := <-c.eventCh:
			c.logDebug("dispatching", m.Method, "event: ")
			c.eventLock.Lock()
			cb, ok := c.eventDispatcher[m.Method]
			c.eventLock.Unlock()
			if !ok {
				break
			}
			cb(c, m.Id, m.Msg)
		}
	}
}

func (c *ChromeTarget) dispatchWithTimeout(r chan<- *gcdmessage.Message, id int64, msg []byte) {
	timeout := time.NewTimer(c.GetApiTimeout())
	defer timeout.Stop()

	select {
	case r <- &gcdmessage.Message{Id: id, Data: msg}:
		timeout.Stop()
	case <-timeout.C:
		c.logDebug("timed out dispatching request id: ", id, " of msg: ", msg)
		close(r)
		return
	}
	return
}

// check target detached/crashed
// close any replier channels that are open
// dispatch detached / crashed event as usual
func (c *ChromeTarget) checkTargetDisconnected(method string) {
	switch method {
	case "Inspector.targetCrashed", "Inspector.detached":
		c.replyLock.Lock()
		for _, replyCh := range c.replyDispatcher {
			close(replyCh)
		}
		// empty out the dispatcher
		c.replyDispatcher = make(map[int64]chan *gcdmessage.Message)
		c.replyLock.Unlock()
	}
}

// Connects to the tab/process for sending/recv'ing debug events
func wsConnection(ctx context.Context, url string) (*WebSocket, error) {
	client := &WebSocket{}
	if err := client.Connect(ctx, url, nil); err != nil {
		return nil, err
	}
	return client, nil
}

// gcdmessage.ChromeTargeter interface methods

// GetId increments the Id so we can synchronize our request/responses internally
func (c *ChromeTarget) GetId() int64 {
	return atomic.AddInt64(&c.sendId, 1)
}

// GetSendCh the channel used for API components to send back to use
func (c *ChromeTarget) GetSendCh() chan *gcdmessage.Message {
	return c.sendCh
}

// GetDoneCh channel used to signal any pending SendDefaultRequest and SendCustomReturn
// that we are exiting so we don't block goroutines from exiting.
func (c *ChromeTarget) GetDoneCh() chan struct{} {
	return c.doneCh
}

func (c *ChromeTarget) logDebug(args ...interface{}) {
	if c.debugger.debug {
		c.logger.Println(args...)
	}
}

// SendCustomReturn takes in a ParamRequest and gives back a response channel so the caller can decode as necessary.
func (c *ChromeTarget) SendCustomReturn(ctx context.Context, paramRequest *gcdmessage.ParamRequest) (*gcdmessage.Message, error) {
	data, err := json.Marshal(paramRequest)

	if err != nil {
		return nil, err
	}

	c.messageObserver.Request(paramRequest.Id, paramRequest.Method, data)

	response, err := c.sendData(ctx, paramRequest.Id, data)

	c.messageObserver.Response(paramRequest.Id, paramRequest.Method, observer.DigResponseData(response), err)
	return response, err
}

// SendDefaultRequest sends a generic request that gets back a generic response, or error. This returns a ChromeResponse
// object.
func (c *ChromeTarget) SendDefaultRequest(ctx context.Context, paramRequest *gcdmessage.ParamRequest) (*gcdmessage.ChromeResponse, error) {
	data, err := json.Marshal(&gcdmessage.ChromeRequest{Id: paramRequest.Id, Method: paramRequest.Method, Params: paramRequest.Params})

	if err != nil {
		return nil, err
	}

	c.messageObserver.Request(paramRequest.Id, paramRequest.Method, data)

	response, err := c.sendData(ctx, paramRequest.Id, data)

	c.messageObserver.Response(paramRequest.Id, paramRequest.Method, observer.DigResponseData(response), err)

	if err != nil {
		return nil, err
	}

	chromeResponse := &gcdmessage.ChromeResponse{}
	err = json.Unmarshal(response.Data, chromeResponse)

	if err != nil {
		return nil, err
	}

	return chromeResponse, nil
}

func (c *ChromeTarget) sendData(ctx context.Context, ID int64, data []byte) (*gcdmessage.Message, error) {
	recvCh := make(chan *gcdmessage.Message, 1)

	select {
	case <-ctx.Done():
		return nil, &gcdmessage.ChromeCtxDoneErr{}
	case c.sendCh <- &gcdmessage.Message{ReplyCh: recvCh, Id: ID, Data: data}:
	case <-time.After(c.GetApiTimeout()):
		return nil, &gcdmessage.ChromeApiTimeoutErr{}
	case <-c.GetDoneCh():
		return nil, &gcdmessage.ChromeDoneErr{}
	}

	var resp *gcdmessage.Message
	select {
	case <-ctx.Done():
		return nil, &gcdmessage.ChromeCtxDoneErr{}
	case <-time.After(c.GetApiTimeout()):
		return nil, &gcdmessage.ChromeApiTimeoutErr{}
	case resp = <-recvCh:
	case <-c.GetDoneCh():
		return nil, &gcdmessage.ChromeDoneErr{}
	}

	if resp == nil || resp.Data == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	return resp, nil
}
