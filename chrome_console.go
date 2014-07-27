package gcd

import (
	"encoding/json"
)

type ConsoleRequest struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

type ChromeConsole struct {
	Page *ChromePage
}

func newChromeConsole(page *ChromePage) *ChromeConsole {
	c := &ChromeConsole{Page: page}
	return c
}

func (c *ChromeConsole) Enable() error {
	req := &ConsoleRequest{Id: c.Page.getId(), Method: "Console.enable"}
	return sendRequest(c.Page.sendCh, req)
}

func (c *ChromeConsole) Disable() error {
	req := &ConsoleRequest{Id: c.Page.getId(), Method: "Console.disable"}
	return sendRequest(c.Page.sendCh, req)
}

func (c *ChromeConsole) ClearMessages() error {
	req := &ConsoleRequest{Id: c.Page.getId(), Method: "Console.clearMessages"}
	return sendRequest(c.Page.sendCh, req)
}

func sendRequest(sendCh chan<- []byte, req *ConsoleRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	sendCh <- data
	return nil
}
