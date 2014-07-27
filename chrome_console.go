package gcd

type ConsoleRequest struct {
	Id     int    `json:"id"`
	Method string `json:"method"`
}

type ChromeConsole struct {
	Page *ChromePage
}

func NewChromeConsole(page *ChromePage) *ChromeConsole {
	c := &ChromeConsole{Page: page}
	return c
}

func (c *ChromeConsole) ClearMessages() error {
	return nil
}
