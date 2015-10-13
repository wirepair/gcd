// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMStorage functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// DOM Storage identifier.
type DOMStorageStorageId struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin for the storage.
	IsLocalStorage bool   `json:"isLocalStorage"` // Whether the storage is local storage (not session storage).
}

//
type DOMStorageDomStorageItemsClearedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
		NewValue  string               `json:"newValue"`  //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
		OldValue  string               `json:"oldValue"`  //
		NewValue  string               `json:"newValue"`  //
	} `json:"Params,omitempty"`
}

type DOMStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMStorage(target gcdmessage.ChromeTargeter) *DOMStorage {
	c := &DOMStorage{target: target}
	return c
}

// Enables storage tracking, storage events will now be delivered to the client.
func (c *DOMStorage) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.enable"})
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (c *DOMStorage) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.disable"})
}

// GetDOMStorageItems -
// storageId -
// Returns -  entries -
func (c *DOMStorage) GetDOMStorageItems(storageId *DOMStorageStorageId) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["storageId"] = storageId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.getDOMStorageItems", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Entries []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Entries, nil
}

// SetDOMStorageItem -
// storageId -
// key -
// value -
func (c *DOMStorage) SetDOMStorageItem(storageId *DOMStorageStorageId, key string, value string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["storageId"] = storageId
	paramRequest["key"] = key
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.setDOMStorageItem", Params: paramRequest})
}

// RemoveDOMStorageItem -
// storageId -
// key -
func (c *DOMStorage) RemoveDOMStorageItem(storageId *DOMStorageStorageId, key string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["storageId"] = storageId
	paramRequest["key"] = key
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.removeDOMStorageItem", Params: paramRequest})
}
