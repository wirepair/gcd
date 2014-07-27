package gcd

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net"
	"sync/atomic"
)

type InspectablePage struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	FaviconUrl           string `json:"faviconUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

type ChromePage struct {
	conn    *websocket.Conn
	Console *ChromeConsole
	Page    *InspectablePage
	sendCh  chan []byte
	recvCh  chan []byte
	doneCh  chan bool
	sendId  int64
	/*
		Debugger    *ChromeDebugger
		Dom         *ChromeDom
		DomDebugger *ChromeDomDebugger
		Input       *ChromeInput
		Network     *ChromeNetwork
		Page        *ChromePage
		Runtime     *ChromeRuntime
		Timeline    *ChromeTimeline
	*/
}

func newChromePage(port string, page *InspectablePage) *ChromePage {
	conn := wsConnection(fmt.Sprintf("localhost:%s", port), page.WebSocketDebuggerUrl)
	sendCh := make(chan []byte)
	recvCh := make(chan []byte)
	doneCh := make(chan bool)
	chromePage := &ChromePage{conn: conn, Console: nil, Page: page, sendCh: sendCh, recvCh: recvCh, doneCh: doneCh, sendId: 0}
	chromePage.Console = NewChromeConsole(chromePage)
	chromePage.listen()
	return chromePage
}

func (c *ChromePage) getId() int64 {
	fmt.Printf("sendId: %d\n", c.sendId)
	return atomic.AddInt64(&c.sendId, 1)
}

func (c *ChromePage) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {
		// send message to the client
		case msg := <-c.sendCh:
			log.Println("Send:", string(msg))
			websocket.Message.Send(c.conn, string(msg))
		// receive done request
		case <-c.doneCh:
			fmt.Println("listenWrite doneCh true...")
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

func (c *ChromePage) listen() {
	go c.listenRead()
	go c.listenWrite()
}

// Listen read request via chanel
func (c *ChromePage) listenRead() {
	log.Println("Listening read from client")
	for {
		select {
		// receive done request
		case <-c.doneCh:
			fmt.Println("listenRead doneCh true")
			c.doneCh <- true // for listenWrite method
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
				c.parseResponse(msg)
			}
		}
	}
}

func (c *ChromePage) parseResponse(msg string) {
	log.Printf("parseResponse got msg: %s\n", msg)
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
