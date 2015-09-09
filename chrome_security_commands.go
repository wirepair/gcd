// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Security commands.
// API Version: 1.1

package gcd

import ()

// add this API domain to ChromeTarget
func (c *ChromeTarget) Security() *ChromeSecurity {
	if c.security == nil {
		c.security = newChromeSecurity(c)
	}
	return c.security
}

type ChromeSecurity struct {
	target *ChromeTarget
}

func newChromeSecurity(target *ChromeTarget) *ChromeSecurity {
	c := &ChromeSecurity{target: target}
	return c
}

// Enables tracking security state changes.
func (c *ChromeSecurity) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Security.enable"})
}

// Disables tracking security state changes.
func (c *ChromeSecurity) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Security.disable"})
}
