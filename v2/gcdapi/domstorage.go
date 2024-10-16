// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMStorage functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// DOM Storage identifier.
type DOMStorageStorageId struct {
	SecurityOrigin string `json:"securityOrigin,omitempty"` // Security origin for the storage.
	StorageKey     string `json:"storageKey,omitempty"`     // Represents a key by which DOM Storage keys its CachedStorageAreas
	IsLocalStorage bool   `json:"isLocalStorage"`           // Whether the storage is local storage (not session storage).
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
type DOMStorageDomStorageItemRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
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

//
type DOMStorageDomStorageItemsClearedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
	} `json:"Params,omitempty"`
}

type DOMStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMStorage(target gcdmessage.ChromeTargeter) *DOMStorage {
	c := &DOMStorage{target: target}
	return c
}

type DOMStorageClearParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
}

// ClearWithParams -
func (c *DOMStorage) ClearWithParams(ctx context.Context, v *DOMStorageClearParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.clear", Params: v})
}

// Clear -
// storageId -
func (c *DOMStorage) Clear(ctx context.Context, storageId *DOMStorageStorageId) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageClearParams
	v.StorageId = storageId
	return c.ClearWithParams(ctx, &v)
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (c *DOMStorage) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.disable"})
}

// Enables storage tracking, storage events will now be delivered to the client.
func (c *DOMStorage) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.enable"})
}

type DOMStorageGetDOMStorageItemsParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
}

// GetDOMStorageItemsWithParams -
// Returns -  entries -
func (c *DOMStorage) GetDOMStorageItemsWithParams(ctx context.Context, v *DOMStorageGetDOMStorageItemsParams) ([][]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.getDOMStorageItems", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Entries [][]string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Entries, nil
}

// GetDOMStorageItems -
// storageId -
// Returns -  entries -
func (c *DOMStorage) GetDOMStorageItems(ctx context.Context, storageId *DOMStorageStorageId) ([][]string, error) {
	var v DOMStorageGetDOMStorageItemsParams
	v.StorageId = storageId
	return c.GetDOMStorageItemsWithParams(ctx, &v)
}

type DOMStorageRemoveDOMStorageItemParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
	//
	Key string `json:"key"`
}

// RemoveDOMStorageItemWithParams -
func (c *DOMStorage) RemoveDOMStorageItemWithParams(ctx context.Context, v *DOMStorageRemoveDOMStorageItemParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.removeDOMStorageItem", Params: v})
}

// RemoveDOMStorageItem -
// storageId -
// key -
func (c *DOMStorage) RemoveDOMStorageItem(ctx context.Context, storageId *DOMStorageStorageId, key string) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageRemoveDOMStorageItemParams
	v.StorageId = storageId
	v.Key = key
	return c.RemoveDOMStorageItemWithParams(ctx, &v)
}

type DOMStorageSetDOMStorageItemParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
	//
	Key string `json:"key"`
	//
	Value string `json:"value"`
}

// SetDOMStorageItemWithParams -
func (c *DOMStorage) SetDOMStorageItemWithParams(ctx context.Context, v *DOMStorageSetDOMStorageItemParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.setDOMStorageItem", Params: v})
}

// SetDOMStorageItem -
// storageId -
// key -
// value -
func (c *DOMStorage) SetDOMStorageItem(ctx context.Context, storageId *DOMStorageStorageId, key string, value string) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageSetDOMStorageItemParams
	v.StorageId = storageId
	v.Key = key
	v.Value = value
	return c.SetDOMStorageItemWithParams(ctx, &v)
}
