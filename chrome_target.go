/*
The MIT License (MIT)

Copyright (c) 2015 isaac dawson

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
	"encoding/json"
	"github.com/wirepair/gcd/gcdapi"
	"github.com/wirepair/gcd/gcdmessage"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

// Defines the 'tab' or target for this chrome instance, can be multiple and background processes
// are included (not just visual tabs)
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

// Our Chrome Target (Tab/Process). Messages are returned to callers via non-buffered channels. Helpfully,
// the remote debugger service uses id's so we can correlate which request should match which response.
// We use a map to store the id of the request which contains a reference to a gcdmessage.Message that holds the
// reply channel for the ChromeTarget to return the response to.
// Events are handled by mapping the method name to a function which takes a target and byte output.
// For now, callers will need to unmarshall the types themselves.
type ChromeTarget struct {
	replyLock       sync.RWMutex                           // lock for dispatching responses
	replyDispatcher map[int64]chan *gcdmessage.Message     // Replies to synch methods using a non-buffered channel
	eventLock       sync.RWMutex                           // lock for dispatching events
	eventDispatcher map[string]func(*ChromeTarget, []byte) // calls the function when events match the subscribed method
	conn            *websocket.Conn                        // the connection to the chrome debugger service for this tab/process

	// Chrome Debugger Domains
	ApplicationCache  *gcdapi.ApplicationCache  // application cache API
	Console           *gcdapi.Console           // console API
	CSS               *gcdapi.CSS               // CSS API
	Database          *gcdapi.Database          // Database API
	Debugger          *gcdapi.Debugger          // JS Debugger API
	DeviceOrientation *gcdapi.DeviceOrientation // Device Orientation API
	DOM               *gcdapi.DOM               // DOM API
	DOMDebugger       *gcdapi.DOMDebugger       // DOM Debugger API
	DOMStorage        *gcdapi.DOMStorage        // DOM Storage API
	FileSystem        *gcdapi.FileSystem        // Is anyone still reading this? FileSystem API
	HeapProfiler      *gcdapi.HeapProfiler      // HeapProfiler API
	IndexedDB         *gcdapi.IndexedDB         // IndexedDB API
	Input             *gcdapi.Input             // Why am i doing this, it's obvious what they are, I quit.
	Inspector         *gcdapi.Inspector
	LayerTree         *gcdapi.LayerTree
	Memory            *gcdapi.Memory
	Network           *gcdapi.Network
	Page              *gcdapi.Page
	Power             *gcdapi.Power
	Profiler          *gcdapi.Profiler
	Runtime           *gcdapi.Runtime
	Timeline          *gcdapi.Timeline
	Tracing           *gcdapi.Tracing
	Worker            *gcdapi.Worker
	Accessibility     *gcdapi.Accessibility
	Animation         *gcdapi.Animation
	CacheStorage      *gcdapi.CacheStorage
	Emulation         *gcdapi.Emulation
	IO                *gcdapi.IO
	Rendering         *gcdapi.Rendering
	ScreenOrientation *gcdapi.ScreenOrientation
	Security          *gcdapi.Security
	ServiceWorker     *gcdapi.ServiceWorker

	Target       *TargetInfo              // The target information see, TargetInfo
	sendCh       chan *gcdmessage.Message // The channel used for API components to send back to use
	doneCh       chan bool                // we be donez.
	sendId       int64                    // An Id which is atomically incremented per request.
	debugEvents  bool                     // flag for spitting out event data as a string which we have not subscribed to.
	shuttingdown bool
}

// Creates a new Chrome Target by connecting to the service given the URL taken from initial connection.
func newChromeTarget(addr string, target *TargetInfo) *ChromeTarget {
	conn := wsConnection(addr, target.WebSocketDebuggerUrl)
	replier := make(map[int64]chan *gcdmessage.Message)
	var replyLock sync.RWMutex
	var eventLock sync.RWMutex
	eventer := make(map[string]func(*ChromeTarget, []byte))
	sendCh := make(chan *gcdmessage.Message)
	doneCh := make(chan bool)
	chromeTarget := &ChromeTarget{conn: conn, eventLock: eventLock, replyLock: replyLock, Target: target, sendCh: sendCh, replyDispatcher: replier, eventDispatcher: eventer, doneCh: doneCh, sendId: 0}
	chromeTarget.Init()
	chromeTarget.listen()
	return chromeTarget
}

// Initialize all api objects
func (c *ChromeTarget) Init() {
	c.ApplicationCache = gcdapi.NewApplicationCache(c)
	c.Console = gcdapi.NewConsole(c)
	c.CSS = gcdapi.NewCSS(c)
	c.Database = gcdapi.NewDatabase(c)
	c.Debugger = gcdapi.NewDebugger(c)
	c.DeviceOrientation = gcdapi.NewDeviceOrientation(c)
	c.DOM = gcdapi.NewDOM(c)
	c.DOMDebugger = gcdapi.NewDOMDebugger(c)
	c.DOMStorage = gcdapi.NewDOMStorage(c)
	c.FileSystem = gcdapi.NewFileSystem(c)
	c.HeapProfiler = gcdapi.NewHeapProfiler(c)
	c.IndexedDB = gcdapi.NewIndexedDB(c)
	c.Input = gcdapi.NewInput(c)
	c.Inspector = gcdapi.NewInspector(c)
	c.LayerTree = gcdapi.NewLayerTree(c)
	c.Memory = gcdapi.NewMemory(c)
	c.Network = gcdapi.NewNetwork(c)
	c.Page = gcdapi.NewPage(c)
	c.Power = gcdapi.NewPower(c)
	c.Profiler = gcdapi.NewProfiler(c)
	c.Runtime = gcdapi.NewRuntime(c)
	c.Timeline = gcdapi.NewTimeline(c)
	c.Tracing = gcdapi.NewTracing(c)
	c.Worker = gcdapi.NewWorker(c)
	c.Accessibility = gcdapi.NewAccessibility(c)
	c.Animation = gcdapi.NewAnimation(c)
	c.CacheStorage = gcdapi.NewCacheStorage(c)
	c.Emulation = gcdapi.NewEmulation(c)
	c.IO = gcdapi.NewIO(c)
	c.Rendering = gcdapi.NewRendering(c)
	c.ScreenOrientation = gcdapi.NewScreenOrientation(c)
	c.Security = gcdapi.NewSecurity(c)
	c.ServiceWorker = gcdapi.NewServiceWorker(c)
}

func (c *ChromeTarget) shutdown() {
	c.shuttingdown = true
	c.conn.Close()
}

/// Subscribes Events, you must know the method name, such as Page.loadFiredEvent, and bind a function
// which takes a ChromeTarget (us) and the raw JSON byte data for that event.
func (c *ChromeTarget) Subscribe(method string, callback func(*ChromeTarget, []byte)) {
	c.eventLock.Lock()
	c.eventDispatcher[method] = callback
	c.eventLock.Unlock()
}

// Unsubscribes the handler for no longer recieving events.
func (c *ChromeTarget) Unsubscribe(method string) {
	c.eventLock.Lock()
	delete(c.eventDispatcher, method)
	c.eventLock.Unlock()
}

// Whether to print out raw JSON event data.
func (c *ChromeTarget) DebugEvents(debug bool) {
	c.debugEvents = debug
}

// Listens for API components wanting to send, and recv'ing data from the Chrome Debugger Service
func (c *ChromeTarget) listen() {
	go c.listenRead()
	go c.listenWrite()
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
			//log.Println("Send:", string(msg.Data))
			websocket.Message.Send(c.conn, string(msg.Data))
			//log.Println("Done sending to WS")
		// receive done from listenRead
		case <-c.doneCh:
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listens for responses coming in from the Chrome Debugger Service.
func (c *ChromeTarget) listenRead() {
	for {
		select {
		// receive done from listenWrite
		case <-c.doneCh:
			c.doneCh <- true
			return
		// read data from websocket connection
		default:
			var msg string
			err := websocket.Message.Receive(c.conn, &msg)
			if err == io.EOF {
				c.doneCh <- true
				return
			} else if err != nil {
				if c.shuttingdown == false {
					log.Printf("error in websocket read: %v\n", err)
				}
				c.doneCh <- true
			} else {
				go c.dispatchResponse([]byte(msg))
			}
		}
	}
}

type responseHeader struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

// dispatchResponse takes in the json message and extracts
// the id or method fields to dispatch either responses or events
// to listeners.
func (c *ChromeTarget) dispatchResponse(msg []byte) {
	f := &responseHeader{}
	err := json.Unmarshal(msg, f)
	if err != nil {
		log.Fatalf("error reading response data from chrome target: %v\n", err)
	}

	c.replyLock.Lock()
	if r, ok := c.replyDispatcher[f.Id]; ok {
		delete(c.replyDispatcher, f.Id)
		c.replyLock.Unlock()
		r <- &gcdmessage.Message{Id: f.Id, Data: msg}
		return
	}
	c.replyLock.Unlock()

	c.eventLock.RLock()
	if r, ok := c.eventDispatcher[f.Method]; ok {
		c.eventLock.RUnlock()
		go r(c, msg)
		return

	}
	c.eventLock.RUnlock()

	if c.debugEvents == true {
		log.Printf("\n\nno event recv bound for: %s\n", f.Method)
		log.Printf("data: %s\n\n", string(msg))
	}
}

// Connects to the tab/process for sending/recv'ing debug events
func wsConnection(addr, url string) *websocket.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("error dialing ChromeTarget websocket: ", err)
	}
	config, errConfig := websocket.NewConfig(url, "http://localhost")
	if errConfig != nil {
		log.Fatalf("error building websocket config: addr: %s url: %s %v\n", addr, url, err)
	}
	client, errWS := websocket.NewClient(config, conn)
	if errWS != nil {
		log.Fatalf("error during websocket handshake: %v\n", errWS)
	}
	return client
}

// gcdmessage.ChromeTargeter interface methods

// Increments the Id so we can synchronize our request/responses internally
func (c *ChromeTarget) GetId() int64 {
	return atomic.AddInt64(&c.sendId, 1)
}

// The channel used for API components to send back to use
func (c *ChromeTarget) GetSendCh() chan *gcdmessage.Message {
	return c.sendCh
}
