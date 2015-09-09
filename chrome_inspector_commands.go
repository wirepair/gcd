// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Inspector commands.
// API Version: 1.1

package gcd

import ()

// add this API domain to ChromeTarget
func (c *ChromeTarget) Inspector() *ChromeInspector {
	if c.inspector == nil {
		c.inspector = newChromeInspector(c)
	}
	return c.inspector
}

type ChromeInspector struct {
	target *ChromeTarget
}

func newChromeInspector(target *ChromeTarget) *ChromeInspector {
	c := &ChromeInspector{target: target}
	return c
}

// Enables inspector domain notifications.
func (c *ChromeInspector) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Inspector.enable"})
}

// Disables inspector domain notifications.
func (c *ChromeInspector) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Inspector.disable"})
}
