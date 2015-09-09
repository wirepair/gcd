// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the CacheStorage types.
// API Version: 1.1
package types

// Unique identifier of the Cache object.
type ChromeCacheStorageCacheId string

// Data entry.
type ChromeCacheStorageDataEntry struct {
	Request  string `json:"request"`  // Request url spec.
	Response string `json:"response"` // Response stataus text.
}

// Cache identifier.
type ChromeCacheStorageCache struct {
	CacheId        *ChromeCacheStorageCacheId `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string                     `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string                     `json:"cacheName"`      // The name of the cache.
}
