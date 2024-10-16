// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Extensions functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

type Extensions struct {
	target gcdmessage.ChromeTargeter
}

func NewExtensions(target gcdmessage.ChromeTargeter) *Extensions {
	c := &Extensions{target: target}
	return c
}

type ExtensionsLoadUnpackedParams struct {
	// Absolute file path.
	Path string `json:"path"`
}

// LoadUnpackedWithParams - Installs an unpacked extension from the filesystem similar to --load-extension CLI flags. Returns extension ID once the extension has been installed. Available if the client is connected using the --remote-debugging-pipe flag and the --enable-unsafe-extension-debugging flag is set.
// Returns -  id - Extension id.
func (c *Extensions) LoadUnpackedWithParams(ctx context.Context, v *ExtensionsLoadUnpackedParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Extensions.loadUnpacked", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Id string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Id, nil
}

// LoadUnpacked - Installs an unpacked extension from the filesystem similar to --load-extension CLI flags. Returns extension ID once the extension has been installed. Available if the client is connected using the --remote-debugging-pipe flag and the --enable-unsafe-extension-debugging flag is set.
// path - Absolute file path.
// Returns -  id - Extension id.
func (c *Extensions) LoadUnpacked(ctx context.Context, path string) (string, error) {
	var v ExtensionsLoadUnpackedParams
	v.Path = path
	return c.LoadUnpackedWithParams(ctx, &v)
}

type ExtensionsGetStorageItemsParams struct {
	// ID of extension.
	Id string `json:"id"`
	// StorageArea to retrieve data from. enum values: session, local, sync, managed
	StorageArea string `json:"storageArea"`
	// Keys to retrieve.
	Keys []string `json:"keys,omitempty"`
}

// GetStorageItemsWithParams - Gets data from extension storage in the given `storageArea`. If `keys` is specified, these are used to filter the result.
// Returns -  data -
func (c *Extensions) GetStorageItemsWithParams(ctx context.Context, v *ExtensionsGetStorageItemsParams) (map[string]interface{}, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Extensions.getStorageItems", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Data map[string]interface{}
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

	return chromeData.Result.Data, nil
}

// GetStorageItems - Gets data from extension storage in the given `storageArea`. If `keys` is specified, these are used to filter the result.
// id - ID of extension.
// storageArea - StorageArea to retrieve data from. enum values: session, local, sync, managed
// keys - Keys to retrieve.
// Returns -  data -
func (c *Extensions) GetStorageItems(ctx context.Context, id string, storageArea string, keys []string) (map[string]interface{}, error) {
	var v ExtensionsGetStorageItemsParams
	v.Id = id
	v.StorageArea = storageArea
	v.Keys = keys
	return c.GetStorageItemsWithParams(ctx, &v)
}

type ExtensionsRemoveStorageItemsParams struct {
	// ID of extension.
	Id string `json:"id"`
	// StorageArea to remove data from. enum values: session, local, sync, managed
	StorageArea string `json:"storageArea"`
	// Keys to remove.
	Keys []string `json:"keys"`
}

// RemoveStorageItemsWithParams - Removes `keys` from extension storage in the given `storageArea`.
func (c *Extensions) RemoveStorageItemsWithParams(ctx context.Context, v *ExtensionsRemoveStorageItemsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Extensions.removeStorageItems", Params: v})
}

// RemoveStorageItems - Removes `keys` from extension storage in the given `storageArea`.
// id - ID of extension.
// storageArea - StorageArea to remove data from. enum values: session, local, sync, managed
// keys - Keys to remove.
func (c *Extensions) RemoveStorageItems(ctx context.Context, id string, storageArea string, keys []string) (*gcdmessage.ChromeResponse, error) {
	var v ExtensionsRemoveStorageItemsParams
	v.Id = id
	v.StorageArea = storageArea
	v.Keys = keys
	return c.RemoveStorageItemsWithParams(ctx, &v)
}

type ExtensionsClearStorageItemsParams struct {
	// ID of extension.
	Id string `json:"id"`
	// StorageArea to remove data from. enum values: session, local, sync, managed
	StorageArea string `json:"storageArea"`
}

// ClearStorageItemsWithParams - Clears extension storage in the given `storageArea`.
func (c *Extensions) ClearStorageItemsWithParams(ctx context.Context, v *ExtensionsClearStorageItemsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Extensions.clearStorageItems", Params: v})
}

// ClearStorageItems - Clears extension storage in the given `storageArea`.
// id - ID of extension.
// storageArea - StorageArea to remove data from. enum values: session, local, sync, managed
func (c *Extensions) ClearStorageItems(ctx context.Context, id string, storageArea string) (*gcdmessage.ChromeResponse, error) {
	var v ExtensionsClearStorageItemsParams
	v.Id = id
	v.StorageArea = storageArea
	return c.ClearStorageItemsWithParams(ctx, &v)
}

type ExtensionsSetStorageItemsParams struct {
	// ID of extension.
	Id string `json:"id"`
	// StorageArea to set data in. enum values: session, local, sync, managed
	StorageArea string `json:"storageArea"`
	// Values to set.
	Values map[string]interface{} `json:"values"`
}

// SetStorageItemsWithParams - Sets `values` in extension storage in the given `storageArea`. The provided `values` will be merged with existing values in the storage area.
func (c *Extensions) SetStorageItemsWithParams(ctx context.Context, v *ExtensionsSetStorageItemsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Extensions.setStorageItems", Params: v})
}

// SetStorageItems - Sets `values` in extension storage in the given `storageArea`. The provided `values` will be merged with existing values in the storage area.
// id - ID of extension.
// storageArea - StorageArea to set data in. enum values: session, local, sync, managed
// values - Values to set.
func (c *Extensions) SetStorageItems(ctx context.Context, id string, storageArea string, values map[string]interface{}) (*gcdmessage.ChromeResponse, error) {
	var v ExtensionsSetStorageItemsParams
	v.Id = id
	v.StorageArea = storageArea
	v.Values = values
	return c.SetStorageItemsWithParams(ctx, &v)
}
