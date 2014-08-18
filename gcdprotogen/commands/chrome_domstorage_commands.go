// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOMStorage commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) DOMStorage() *ChromeDOMStorage {
	if c.domstorage == nil {
		c.domstorage = newChromeDOMStorage(c)
	}
	return c.domstorage
}


type ChromeDOMStorage struct {
	target *ChromeTarget
}

func newChromeDOMStorage(target *ChromeTarget) *ChromeDOMStorage {
	c := &ChromeDOMStorage{target: target}
	return c
}

// start non parameterized commands 
// Enables storage tracking, storage events will now be delivered to the client.
func (c *ChromeDOMStorage) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.enable"})
}
 
// Disables storage tracking, prevents storage events from being sent to the client.
func (c *ChromeDOMStorage) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// setDOMStorageItem - 
// storageId - 
// key - 
// value - 
func (c *ChromeDOMStorage) SetDOMStorageItem(storageId *types.ChromeDOMStorageStorageId, key string, value string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["storageId"] = storageId
	paramRequest["key"] = key
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.setDOMStorageItem", Params: paramRequest})
}

// removeDOMStorageItem - 
// storageId - 
// key - 
func (c *ChromeDOMStorage) RemoveDOMStorageItem(storageId *types.ChromeDOMStorageStorageId, key string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["storageId"] = storageId
	paramRequest["key"] = key
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.removeDOMStorageItem", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types


// end commands with no parameters but special return types

