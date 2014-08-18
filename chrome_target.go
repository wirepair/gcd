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

type ChromeTarget struct {
	sync.RWMutex
	replyDispatcher map[int64]chan *gcdMessage
	eventDispatcher map[string]func(*ChromeTarget, []byte)
	conn            *websocket.Conn
	console         *ChromeConsole
	input           *ChromeInput
	network         *ChromeNetwork
	domstorage      *ChromeDOMStorage
	page            *ChromePage
	memory          *ChromeMemory
	Target          *TargetInfo
	sendCh          chan *gcdMessage
	doneCh          chan bool
	sendId          int64
	/*
		debugger    *ChromeDebugger
		dom         *ChromeDom
		domDebugger *ChromeDomDebugger

		runtime     *ChromeRuntime
		timeline    *ChromeTimeline
	*/
}

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

func (c *ChromeTarget) getId() int64 {
	fmt.Printf("sendId: %d\n", c.sendId)
	return atomic.AddInt64(&c.sendId, 1)
}

func (c *ChromeTarget) Subscribe(method string, callback func(*ChromeTarget, []byte)) {
	c.Lock()
	c.eventDispatcher[method] = callback
	c.Unlock()
	fmt.Printf("finished subscribing to %s\n", method)
}

func (c *ChromeTarget) listen() {
	go c.listenRead()
	go c.listenWrite()
}

func (c *ChromeTarget) listenWrite() {
	log.Println("Listening write for client")
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
			fmt.Println("listenWrite doneCh true, returning...")
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via channel
func (c *ChromeTarget) listenRead() {
	log.Println("Listening read from client")
	for {
		select {
		// receive done from listenWrite
		case <-c.doneCh:
			fmt.Println("listenRead doneCh true, returning")
			c.doneCh <- true
			return
		// read data from websocket connection
		default:
			var msg string
			err := websocket.Message.Receive(c.conn, &msg)
			if err == io.EOF {
				fmt.Println("EOF, sending doneCh")
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
		r <- &gcdMessage{Id: f.Id, Data: msg}
		c.Unlock()
		return
	}
	c.Unlock()

	c.RLock()
	fmt.Printf(string(msg))
	if r, ok := c.eventDispatcher[f.Method]; ok {
		r(c, msg)
	} else {
		fmt.Printf("no event recvr bound for: %s", f.Method)
	}
	c.RUnlock()
}

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
