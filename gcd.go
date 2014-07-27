package gcd

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net"
)

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
