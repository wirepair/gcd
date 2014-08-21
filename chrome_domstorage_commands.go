// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOMStorage commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
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

// Enables storage tracking, storage events will now be delivered to the client.
func (c *ChromeDOMStorage) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.enable"})
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (c *ChromeDOMStorage) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.disable"})
}

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

// getDOMStorageItems -
// Returns -
func (c *ChromeDOMStorage) GetDOMStorageItems(storageId *types.ChromeDOMStorageStorageId) ([]*types.ChromeDOMStorageItem, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["storageId"] = storageId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.getDOMStorageItems", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Entries []*types.ChromeDOMStorageItem
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.Entries, nil
}
