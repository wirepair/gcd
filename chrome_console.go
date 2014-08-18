package gcd

import (
	"encoding/json"
)

func (c *ChromeTarget) Console() *ChromeConsole {
	if c.console == nil {
		c.console = newChromeConsole(c)
	}
	return c.console
}

type ChromeConsole struct {
	target *ChromeTarget
}

func newChromeConsole(target *ChromeTarget) *ChromeConsole {
	c := &ChromeConsole{target: target}
	return c
}

func (c *ChromeConsole) Enable() (*ChromeResponse, error) {
	return sendConsoleRequest(c.target.sendCh, c.target.getId(), "Console.enable")
}

func (c *ChromeConsole) Disable() (*ChromeResponse, error) {
	return sendConsoleRequest(c.target.sendCh, c.target.getId(), "Console.disable")
}

func (c *ChromeConsole) ClearMessages() (*ChromeResponse, error) {
	return sendConsoleRequest(c.target.sendCh, c.target.getId(), "Console.clearMessages")
}

func sendConsoleRequest(sendCh chan<- *gcdMessage, id int64, method string) (*ChromeResponse, error) {
	req := &ChromeRequest{Id: id, Method: method}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	recvCh := make(chan *gcdMessage)
	sendMsg := &gcdMessage{ReplyCh: recvCh, Id: id, Data: []byte(data)}
	sendCh <- sendMsg
	resp := <-recvCh
	consoleResponse := &ChromeResponse{}
	err = json.Unmarshal(resp.Data, consoleResponse)
	if err != nil {
		return nil, err
	}
	return consoleResponse, nil
}
