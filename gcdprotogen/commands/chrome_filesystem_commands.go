// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the FileSystem commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) FileSystem() *ChromeFileSystem {
	if c.filesystem == nil {
		c.filesystem = newChromeFileSystem(c)
	}
	return c.filesystem
}


type ChromeFileSystem struct {
	target *ChromeTarget
}

func newChromeFileSystem(target *ChromeTarget) *ChromeFileSystem {
	c := &ChromeFileSystem{target: target}
	return c
}

// start non parameterized commands 
// Enables events from backend.
func (c *ChromeFileSystem) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.enable"})
}
 
// Disables events from backend.
func (c *ChromeFileSystem) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "FileSystem.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

