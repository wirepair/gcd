// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Console commands.
// API Version: 1.1

package gcd

import ()

// add this API domain to ChromeTarget
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

// Enables console domain, sends the messages collected so far to the client by means of the <code>messageAdded</code> notification.
func (c *ChromeConsole) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.enable"})
}

// Disables console domain, prevents further console messages from being reported to the client.
func (c *ChromeConsole) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.disable"})
}

// Clears console messages collected in the browser.
func (c *ChromeConsole) ClearMessages() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.clearMessages"})
}
