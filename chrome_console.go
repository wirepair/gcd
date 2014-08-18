package gcd

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
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.enable"})
}

func (c *ChromeConsole) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.disable"})
}

func (c *ChromeConsole) ClearMessages() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Console.clearMessages"})
}
