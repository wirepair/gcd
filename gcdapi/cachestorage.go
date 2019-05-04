// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CacheStorage functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Data entry.
type CacheStorageDataEntry struct {
	RequestURL         string                `json:"requestURL"`         // Request URL.
	RequestMethod      string                `json:"requestMethod"`      // Request method.
	RequestHeaders     []*CacheStorageHeader `json:"requestHeaders"`     // Request headers
	ResponseTime       float64               `json:"responseTime"`       // Number of seconds since epoch.
	ResponseStatus     int                   `json:"responseStatus"`     // HTTP response status code.
	ResponseStatusText string                `json:"responseStatusText"` // HTTP response status text.
	ResponseType       string                `json:"responseType"`       // HTTP response type enum values: basic, cors, default, error, opaqueResponse, opaqueRedirect
	ResponseHeaders    []*CacheStorageHeader `json:"responseHeaders"`    // Response headers
}

// Cache identifier.
type CacheStorageCache struct {
	CacheId        string `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string `json:"cacheName"`      // The name of the cache.
}

// No Description.
type CacheStorageHeader struct {
	Name  string `json:"name"`  //
	Value string `json:"value"` //
}

// Cached response
type CacheStorageCachedResponse struct {
	Body string `json:"body"` // Entry content, base64-encoded.
}

type CacheStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewCacheStorage(target gcdmessage.ChromeTargeter) *CacheStorage {
	c := &CacheStorage{target: target}
	return c
}

type CacheStorageDeleteCacheParams struct {
	// Id of cache for deletion.
	CacheId string `json:"cacheId"`
}

// DeleteCacheWithParams - Deletes a cache.
func (c *CacheStorage) DeleteCacheWithParams(v *CacheStorageDeleteCacheParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteCache", Params: v})
}

// DeleteCache - Deletes a cache.
// cacheId - Id of cache for deletion.
func (c *CacheStorage) DeleteCache(cacheId string) (*gcdmessage.ChromeResponse, error) {
	var v CacheStorageDeleteCacheParams
	v.CacheId = cacheId
	return c.DeleteCacheWithParams(&v)
}

type CacheStorageDeleteEntryParams struct {
	// Id of cache where the entry will be deleted.
	CacheId string `json:"cacheId"`
	// URL spec of the request.
	Request string `json:"request"`
}

// DeleteEntryWithParams - Deletes a cache entry.
func (c *CacheStorage) DeleteEntryWithParams(v *CacheStorageDeleteEntryParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteEntry", Params: v})
}

// DeleteEntry - Deletes a cache entry.
// cacheId - Id of cache where the entry will be deleted.
// request - URL spec of the request.
func (c *CacheStorage) DeleteEntry(cacheId string, request string) (*gcdmessage.ChromeResponse, error) {
	var v CacheStorageDeleteEntryParams
	v.CacheId = cacheId
	v.Request = request
	return c.DeleteEntryWithParams(&v)
}

type CacheStorageRequestCacheNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

// RequestCacheNamesWithParams - Requests cache names.
// Returns -  caches - Caches for the security origin.
func (c *CacheStorage) RequestCacheNamesWithParams(v *CacheStorageRequestCacheNamesParams) ([]*CacheStorageCache, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestCacheNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Caches []*CacheStorageCache
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

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Caches, nil
}

// RequestCacheNames - Requests cache names.
// securityOrigin - Security origin.
// Returns -  caches - Caches for the security origin.
func (c *CacheStorage) RequestCacheNames(securityOrigin string) ([]*CacheStorageCache, error) {
	var v CacheStorageRequestCacheNamesParams
	v.SecurityOrigin = securityOrigin
	return c.RequestCacheNamesWithParams(&v)
}

type CacheStorageRequestCachedResponseParams struct {
	// Id of cache that contains the entry.
	CacheId string `json:"cacheId"`
	// URL spec of the request.
	RequestURL string `json:"requestURL"`
	// headers of the request.
	RequestHeaders []*CacheStorageHeader `json:"requestHeaders"`
}

// RequestCachedResponseWithParams - Fetches cache entry.
// Returns -  response - Response read from the cache.
func (c *CacheStorage) RequestCachedResponseWithParams(v *CacheStorageRequestCachedResponseParams) (*CacheStorageCachedResponse, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestCachedResponse", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Response *CacheStorageCachedResponse
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

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Response, nil
}

// RequestCachedResponse - Fetches cache entry.
// cacheId - Id of cache that contains the entry.
// requestURL - URL spec of the request.
// requestHeaders - headers of the request.
// Returns -  response - Response read from the cache.
func (c *CacheStorage) RequestCachedResponse(cacheId string, requestURL string, requestHeaders []*CacheStorageHeader) (*CacheStorageCachedResponse, error) {
	var v CacheStorageRequestCachedResponseParams
	v.CacheId = cacheId
	v.RequestURL = requestURL
	v.RequestHeaders = requestHeaders
	return c.RequestCachedResponseWithParams(&v)
}

type CacheStorageRequestEntriesParams struct {
	// ID of cache to get entries from.
	CacheId string `json:"cacheId"`
	// Number of records to skip.
	SkipCount int `json:"skipCount"`
	// Number of records to fetch.
	PageSize int `json:"pageSize"`
	// If present, only return the entries containing this substring in the path
	PathFilter string `json:"pathFilter,omitempty"`
}

// RequestEntriesWithParams - Requests data from cache.
// Returns -  cacheDataEntries - Array of object store data entries. returnCount - Count of returned entries from this storage. If pathFilter is empty, it is the count of all entries from this storage.
func (c *CacheStorage) RequestEntriesWithParams(v *CacheStorageRequestEntriesParams) ([]*CacheStorageDataEntry, float64, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestEntries", Params: v})
	if err != nil {
		return nil, 0, err
	}

	var chromeData struct {
		Result struct {
			CacheDataEntries []*CacheStorageDataEntry
			ReturnCount      float64
		}
	}

	if resp == nil {
		return nil, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, 0, err
	}

	return chromeData.Result.CacheDataEntries, chromeData.Result.ReturnCount, nil
}

// RequestEntries - Requests data from cache.
// cacheId - ID of cache to get entries from.
// skipCount - Number of records to skip.
// pageSize - Number of records to fetch.
// pathFilter - If present, only return the entries containing this substring in the path
// Returns -  cacheDataEntries - Array of object store data entries. returnCount - Count of returned entries from this storage. If pathFilter is empty, it is the count of all entries from this storage.
func (c *CacheStorage) RequestEntries(cacheId string, skipCount int, pageSize int, pathFilter string) ([]*CacheStorageDataEntry, float64, error) {
	var v CacheStorageRequestEntriesParams
	v.CacheId = cacheId
	v.SkipCount = skipCount
	v.PageSize = pageSize
	v.PathFilter = pathFilter
	return c.RequestEntriesWithParams(&v)
}
