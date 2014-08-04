package gcd

import (
	"encoding/json"
)

type ConsoleRequest struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

type ConsoleMessage struct {
	Column int    `json:"column,omitempty"`
	Level  string `json:"level"`
	Line   int
	NetworkRequestId
}

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

func (c *ChromeConsole) Enable() error {
	req := &ConsoleRequest{Id: c.target.getId(), Method: "Console.enable"}
	return sendRequest(c.target.sendCh, req)
}

func (c *ChromeConsole) Disable() error {
	req := &ConsoleRequest{Id: c.target.getId(), Method: "Console.disable"}
	return sendRequest(c.target.sendCh, req)
}

func (c *ChromeConsole) ClearMessages() error {
	req := &ConsoleRequest{Id: c.target.getId(), Method: "Console.clearMessages"}
	return sendRequest(c.target.sendCh, req)
}

func sendRequest(sendCh chan<- []byte, req *ConsoleRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	sendCh <- data
	return nil
}
