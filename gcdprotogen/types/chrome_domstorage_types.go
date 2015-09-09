// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOMStorage types.
// API Version: 1.1
package types

// DOM Storage identifier.
type ChromeDOMStorageStorageId struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin for the storage.
	IsLocalStorage bool   `json:"isLocalStorage"` // Whether the storage is local storage (not session storage).
}

// DOM Storage item.
type ChromeDOMStorageItem []string
