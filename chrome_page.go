package gcd

import (
	"code.google.com/p/go.net/websocket"
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

func newChromePage(page *InspectablePage) (*ChromePage, error) {
	return nil, nil
}
