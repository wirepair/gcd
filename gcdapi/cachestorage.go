// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CacheStorage functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Data entry.
type CacheStorageDataEntry struct {
	Request  string `json:"request"`  // Request url spec.
	Response string `json:"response"` // Response stataus text.
}

// Cache identifier.
type CacheStorageCache struct {
	CacheId        string `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string `json:"cacheName"`      // The name of the cache.
}

type CacheStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewCacheStorage(target gcdmessage.ChromeTargeter) *CacheStorage {
	c := &CacheStorage{target: target}
	return c
}

// RequestCacheNames - Requests cache names.
// securityOrigin - Security origin.
// Returns -  caches - Caches for the security origin.
func (c *CacheStorage) RequestCacheNames(securityOrigin string) ([]*CacheStorageCache, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["securityOrigin"] = securityOrigin
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestCacheNames", Params: paramRequest})
	resp := <-recvCh

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

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Caches, nil
}

// RequestEntries - Requests data from cache.
// cacheId - ID of cache to get entries from.
// skipCount - Number of records to skip.
// pageSize - Number of records to fetch.
// Returns -  cacheDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *CacheStorage) RequestEntries(cacheId string, skipCount int, pageSize int) ([]*CacheStorageDataEntry, bool, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["cacheId"] = cacheId
	paramRequest["skipCount"] = skipCount
	paramRequest["pageSize"] = pageSize
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestEntries", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CacheDataEntries []*CacheStorageDataEntry
			HasMore          bool
		}
	}

	if resp == nil {
		return nil, false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, false, err
	}

	return chromeData.Result.CacheDataEntries, chromeData.Result.HasMore, nil
}

// DeleteCache - Deletes a cache.
// cacheId - Id of cache for deletion.
func (c *CacheStorage) DeleteCache(cacheId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["cacheId"] = cacheId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteCache", Params: paramRequest})
}

// DeleteEntry - Deletes a cache entry.
// cacheId - Id of cache where the entry will be deleted.
// request - URL spec of the request.
func (c *CacheStorage) DeleteEntry(cacheId string, request string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cacheId"] = cacheId
	paramRequest["request"] = request
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteEntry", Params: paramRequest})
}
