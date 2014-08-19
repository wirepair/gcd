/*
The MIT License (MIT)

Copyright (c) 2014 isaac dawson

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
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
)

// default response object, contains the id and a result if applicable.
type ChromeResponse struct {
	Id     int64       `json:"id"`
	Result interface{} `json:"result"`
}

// default no-arg request
type ChromeRequest struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

// default chrome error response to an invalid request.
type ChromeErrorResponse struct {
	Id    int64        `json:"id"`    // the request Id that this is a response of
	Error *ChromeError `json:"error"` // the error object
}

// An error object returned from a request
type ChromeError struct {
	Code    int64  `json:"code"`    // the error code
	Message string `json:"message"` // the error message
}

// A gcd type for reporting chrome request errors
type ChromeRequestErr struct {
	Resp *ChromeErrorResponse // a ref to the error response to be used to generate the user friendly error string
}

// user friendly error response
func (cerr *ChromeRequestErr) Error() string {
	return "request " + strconv.FormatInt(cerr.Resp.Id, 10) + " failed, code: " + strconv.FormatInt(cerr.Resp.Error.Code, 10) + " msg: " + cerr.Resp.Error.Message
}

// default request object that has parameters.
type ParamRequest struct {
	Id     int64       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

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

// An internal message object used for components and ChromeTarget to communicate back and forth
type gcdMessage struct {
	ReplyCh chan *gcdMessage // json response channel
	Id      int64            // id to map response channels to send chans
	Data    []byte           // the data for the websocket to send/recv
	Method  string           // event name type.
	Target  *ChromeTarget    // reference to the ChromeTarget for events
}

// Our Chrome Target (Tab/Process). Messages are returned to callers via non-buffered channels. Helpfully,
// the remote debugger service uses id's so we can correlate which request should match which response.
// We use a map to store the id of the request which contains a reference to a gcdMessage that holds the
// reply channel for the ChromeTarget to return the response to.
// Events are handled by mapping the method name to a function which takes a target and byte output.
// For now, callers will need to unmarshall the types themselves.
type ChromeTarget struct {
	sync.RWMutex                                             // lock for dispatching/subscribing
	replyDispatcher   map[int64]chan *gcdMessage             // Replies to synch methods using a non-buffered channel
	eventDispatcher   map[string]func(*ChromeTarget, []byte) // calls the function when events match the subscribed method
	conn              *websocket.Conn                        // the connection to the chrome debugger service for this tab/process
	applicationcache  *ChromeApplicationCache                // application cache API
	canvas            *ChromeCanvas                          // canvas API
	console           *ChromeConsole                         // console API
	css               *ChromeCSS                             // CSS API
	database          *ChromeDatabase                        // Database API
	debugger          *ChromeDebugger                        // JS Debugger API
	deviceorientation *ChromeDeviceOrientation               // Device Orientation API
	dom               *ChromeDOM                             // DOM API
	domdebugger       *ChromeDOMDebugger                     // DOM Debugger API
	domstorage        *ChromeDOMStorage                      // DOM Storage API
	filesystem        *ChromeFileSystem                      // Is anyone still reading this? FileSystem API
	geolocation       *ChromeGeolocation                     // Geolocation API
	heapprofiler      *ChromeHeapProfiler                    // HeapProfiler API
	indexeddb         *ChromeIndexedDB                       // IndexedDB API
	input             *ChromeInput                           // Why am i doing this, it's obvious what they are, I quit.
	inspector         *ChromeInspector
	layertree         *ChromeLayerTree
	memory            *ChromeMemory
	network           *ChromeNetwork
	page              *ChromePage
	power             *ChromePower
	profiler          *ChromeProfiler
	runtime           *ChromeRuntime
	timeline          *ChromeTimeline
	tracing           *ChromeTracing
	worker            *ChromeWorker
	Target            *TargetInfo      // The target information see, TargetInfo
	sendCh            chan *gcdMessage // The channel used for API components to send back to use
	doneCh            chan bool        // we be donez.
	sendId            int64            // An Id which is atomically incremented per request.
}

// Creates a new Chrome Target by connecting to the service given the URL taken from initial connection.
func newChromeTarget(port string, target *TargetInfo) *ChromeTarget {
	conn := wsConnection(fmt.Sprintf("localhost:%s", port), target.WebSocketDebuggerUrl)
	replier := make(map[int64]chan *gcdMessage)
	eventer := make(map[string]func(*ChromeTarget, []byte))
	sendCh := make(chan *gcdMessage)
	doneCh := make(chan bool)
	chromeTarget := &ChromeTarget{conn: conn, Target: target, sendCh: sendCh, replyDispatcher: replier, eventDispatcher: eventer, doneCh: doneCh, sendId: 0}
	chromeTarget.listen()
	return chromeTarget
}

// Increments the Id so we can synchronize our request/responses internally
func (c *ChromeTarget) getId() int64 {
	return atomic.AddInt64(&c.sendId, 1)
}

/// Subscribes Events, you must know the method name, such as Page.loadFiredEvent, and bind a function
// which takes a ChromeTarget (us) and the raw JSON byte data for that event.
func (c *ChromeTarget) Subscribe(method string, callback func(*ChromeTarget, []byte)) {
	c.Lock()
	c.eventDispatcher[method] = callback
	c.Unlock()
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
			c.Lock()
			c.replyDispatcher[msg.Id] = msg.ReplyCh
			c.Unlock()
			//log.Println("Send:", string(msg.Data))
			websocket.Message.Send(c.conn, string(msg.Data))
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
				log.Printf("error in websocket read: %v\n", err)
				c.doneCh <- true
			} else {
				c.dispatchResponse([]byte(msg))
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

	c.Lock()
	if r, ok := c.replyDispatcher[f.Id]; ok {
		delete(c.replyDispatcher, f.Id)
		c.Unlock()
		r <- &gcdMessage{Id: f.Id, Data: msg}
		return
	}
	c.Unlock()

	c.RLock()
	if r, ok := c.eventDispatcher[f.Method]; ok {
		c.RUnlock()
		r(c, msg)
	} else {
		fmt.Printf("no event recvr bound for: %s", f.Method)
		fmt.Printf("data: %s\n", string(msg))
	}
	c.RUnlock()
}

// Connects to the tab/process for sending/recv'ing debug events
func wsConnection(addr, url string) *websocket.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	config, errConfig := websocket.NewConfig(url, "http://localhost")
	if errConfig != nil {
		log.Fatalf("error building config: %v\n", err)
	}
	client, errWS := websocket.NewClient(config, conn)
	if errWS != nil {
		log.Fatalf("WebSocket handshake error: %v\n", errWS)
	}
	return client
}

// Takes in a ParamRequest and gives back a response channel so the caller can decode as necessary.
func sendCustomReturn(sendCh chan<- *gcdMessage, paramRequest *ParamRequest) (<-chan *gcdMessage, error) {
	data, err := json.Marshal(paramRequest)
	if err != nil {
		return nil, err
	}
	log.Print(string(data))
	recvCh := make(chan *gcdMessage)
	sendMsg := &gcdMessage{ReplyCh: recvCh, Id: paramRequest.Id, Data: []byte(data)}
	sendCh <- sendMsg
	return recvCh, nil
}

// Sends a generic request that gets back a generic response, or error. This returns a ChromeResponse
// object.
func sendDefaultRequest(sendCh chan<- *gcdMessage, paramRequest *ParamRequest) (*ChromeResponse, error) {
	req := &ChromeRequest{Id: paramRequest.Id, Method: paramRequest.Method}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	recvCh := make(chan *gcdMessage)
	sendMsg := &gcdMessage{ReplyCh: recvCh, Id: paramRequest.Id, Data: []byte(data)}
	sendCh <- sendMsg
	resp := <-recvCh
	chromeResponse := &ChromeResponse{}
	err = json.Unmarshal(resp.Data, chromeResponse)
	if err != nil {
		return nil, err
	}
	return chromeResponse, nil
}
