// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the CacheStorage commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) CacheStorage() *ChromeCacheStorage {
	if c.cachestorage == nil {
		c.cachestorage = newChromeCacheStorage(c)
	}
	return c.cachestorage
}

type ChromeCacheStorage struct {
	target *ChromeTarget
}

func newChromeCacheStorage(target *ChromeTarget) *ChromeCacheStorage {
	c := &ChromeCacheStorage{target: target}
	return c
}

// deleteCache - Deletes a cache.
// cacheId - Id of cache for deletion.
func (c *ChromeCacheStorage) DeleteCache(cacheId *types.ChromeCacheStorageCacheId) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["cacheId"] = cacheId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CacheStorage.deleteCache", Params: paramRequest})
}

// deleteEntry - Deletes a cache entry.
// cacheId - Id of cache where the entry will be deleted.
// request - URL spec of the request.
func (c *ChromeCacheStorage) DeleteEntry(cacheId *types.ChromeCacheStorageCacheId, request string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["cacheId"] = cacheId
	paramRequest["request"] = request
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CacheStorage.deleteEntry", Params: paramRequest})
}

// requestCacheNames - Requests cache names.
// Returns -
// Caches for the security origin.
func (c *ChromeCacheStorage) RequestCacheNames(securityOrigin string) ([]*types.ChromeCacheStorageCache, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["securityOrigin"] = securityOrigin
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CacheStorage.requestCacheNames", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Caches []*types.ChromeCacheStorageCache
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.Caches, nil
}

// requestEntries - Requests data from cache.
// Returns -
// Array of object store data entries.
// If true, there are more entries to fetch in the given range.
func (c *ChromeCacheStorage) RequestEntries(cacheId *types.ChromeCacheStorageCacheId, skipCount int, pageSize int) ([]*types.ChromeCacheStorageDataEntry, bool, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["cacheId"] = cacheId
	paramRequest["skipCount"] = skipCount
	paramRequest["pageSize"] = pageSize
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "CacheStorage.requestEntries", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			CacheDataEntries []*types.ChromeCacheStorageDataEntry
			HasMore          bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, false, err
	}

	return chromeData.Result.CacheDataEntries, chromeData.Result.HasMore, nil
}
