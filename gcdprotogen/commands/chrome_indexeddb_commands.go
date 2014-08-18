// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the IndexedDB commands.
// API Version: 1.1

package gcd


import (
	
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) IndexedDB() *ChromeIndexedDB {
	if c.indexeddb == nil {
		c.indexeddb = newChromeIndexedDB(c)
	}
	return c.indexeddb
}


type ChromeIndexedDB struct {
	target *ChromeTarget
}

func newChromeIndexedDB(target *ChromeTarget) *ChromeIndexedDB {
	c := &ChromeIndexedDB{target: target}
	return c
}

// start non parameterized commands 
// Enables events from backend.
func (c *ChromeIndexedDB) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.enable"})
}
 
// Disables events from backend.
func (c *ChromeIndexedDB) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// clearObjectStore - Clears all entries from an object store.
// securityOrigin - Security origin.
// databaseName - Database name.
// objectStoreName - Object store name.
func (c *ChromeIndexedDB) ClearObjectStore(securityOrigin string, databaseName string, objectStoreName string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	paramRequest["objectStoreName"] = objectStoreName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.clearObjectStore", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

