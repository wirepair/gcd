// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Database commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Database() *ChromeDatabase {
	if c.database == nil {
		c.database = newChromeDatabase(c)
	}
	return c.database
}


type ChromeDatabase struct {
	target *ChromeTarget
}

func newChromeDatabase(target *ChromeTarget) *ChromeDatabase {
	c := &ChromeDatabase{target: target}
	return c
}

// start non parameterized commands 
// Enables database tracking, database events will now be delivered to the client.
func (c *ChromeDatabase) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.enable"})
}
 
// Disables database tracking, prevents database events from being sent to the client.
func (c *ChromeDatabase) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

