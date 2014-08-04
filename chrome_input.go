package gcd

import (
	"encoding/json"
)

type KeyEvent struct {
	Id     int64         `json:"id"`
	Method string        `json:"method"`
	Params *KeyEventData `json:"params"`
}

type MouseEvent struct {
	Id     int64           `json:"id"`
	Method string          `json:"method"`
	Params *MouseEventData `json:"params"`
}

type KeyEventData struct {
	Type                  string `json:"type"`
	Modifiers             int    `json:"modifiers"`
	Timestamp             int    `json:"timestamp"`
	Text                  string `json:"text"`
	UnmodifiedText        string `json:"unmodifiedText"`
	KeyIdentifier         string `json:"keyIdentifier"`
	WindowsVirtualKeyCode int    `json:"windowsVirtualKeyCode"`
	NativeVirtualKeyCode  int    `json:"nativeVirtualKeyCode"`
	MacCharCode           int    `json:"macCharCode"`
	AutoRepeat            bool   `json:"autoRepeat"`
	IsKeypad              bool   `json:"isKeypad"`
	IsSystemKey           bool   `json:"isSystemKey"`
}

type MouseEventData struct {
	Type       string `json:"type"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Modifiers  int    `json:"modifiers"`
	Timestamp  int    `json:"timestamp"`
	Button     string `json:"button"`
	ClickCount int    `json:"clickCount"`
}

func (c *ChromeTarget) Input() *ChromeInput {
	if c.input == nil {
		c.input = newChromeInput(c)
	}
	return c.input
}

type ChromeInput struct {
	target *ChromeTarget
}

func newChromeInput(target *ChromeTarget) *ChromeInput {
	c := &ChromeInput{target: target}
	return c
}

func (c *ChromeInput) SendKeyEvent(keyData *KeyEventData) error {
	req := &KeyEvent{Id: c.target.getId(), Method: "Input.dispatchKeyEvent", Params: keyData}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.target.sendCh <- data
	return nil
}

func (c *ChromeConsole) SendMouseEvent(mouseData *MouseEventData) error {
	req := &MouseEvent{Id: c.target.getId(), Method: "Input.dispatchMouseEvent", Params: mouseData}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.target.sendCh <- data
	return nil
}

// TODO: Add helpers like TypeKeys(data string), ClickMouse(x, y) etc.
