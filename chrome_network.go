package gcd

import (
	"encoding/json"
)

type NetworkRequest struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

type ChromeNetwork struct {
	target *ChromeTarget
}

func newChromeNetwork(target *ChromeTarget) *ChromeNetwork {
	c := &ChromeNetwork{target: target}
	return c
}

func (c *ChromeNetwork) CanClearBrowserCache() error {
	req := &KeyEvent{Id: c.target.getId(), Method: "Network.CanClearBrowserCache"}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.target.sendCh <- data
	return nil
}

func (c *ChromeConsole) CanClearBrowserCookies() error {
	req := &MouseEvent{Id: c.target.getId(), Method: "Network.canClearBrowserCookies"}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.target.sendCh <- data
	return nil
}
