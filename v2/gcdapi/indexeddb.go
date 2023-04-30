// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains IndexedDB functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Database with an array of object stores.
type IndexedDBDatabaseWithObjectStores struct {
	Name         string                  `json:"name"`         // Database name.
	Version      float64                 `json:"version"`      // Database version (type is not 'integer', as the standard requires the version number to be 'unsigned long long')
	ObjectStores []*IndexedDBObjectStore `json:"objectStores"` // Object stores in this database.
}

// Object store.
type IndexedDBObjectStore struct {
	Name          string                       `json:"name"`          // Object store name.
	KeyPath       *IndexedDBKeyPath            `json:"keyPath"`       // Object store key path.
	AutoIncrement bool                         `json:"autoIncrement"` // If true, object store has auto increment flag set.
	Indexes       []*IndexedDBObjectStoreIndex `json:"indexes"`       // Indexes in this object store.
}

// Object store index.
type IndexedDBObjectStoreIndex struct {
	Name       string            `json:"name"`       // Index name.
	KeyPath    *IndexedDBKeyPath `json:"keyPath"`    // Index key path.
	Unique     bool              `json:"unique"`     // If true, index is unique.
	MultiEntry bool              `json:"multiEntry"` // If true, index allows multiple entries for a key.
}

// Key.
type IndexedDBKey struct {
	Type   string          `json:"type"`             // Key type.
	Number float64         `json:"number,omitempty"` // Number value.
	String string          `json:"string,omitempty"` // String value.
	Date   float64         `json:"date,omitempty"`   // Date value.
	Array  []*IndexedDBKey `json:"array,omitempty"`  // Array value.
}

// Key range.
type IndexedDBKeyRange struct {
	Lower     *IndexedDBKey `json:"lower,omitempty"` // Lower bound.
	Upper     *IndexedDBKey `json:"upper,omitempty"` // Upper bound.
	LowerOpen bool          `json:"lowerOpen"`       // If true lower bound is open.
	UpperOpen bool          `json:"upperOpen"`       // If true upper bound is open.
}

// Data entry.
type IndexedDBDataEntry struct {
	Key        *RuntimeRemoteObject `json:"key"`        // Key object.
	PrimaryKey *RuntimeRemoteObject `json:"primaryKey"` // Primary key object.
	Value      *RuntimeRemoteObject `json:"value"`      // Value object.
}

// Key path.
type IndexedDBKeyPath struct {
	Type   string   `json:"type"`             // Key path type.
	String string   `json:"string,omitempty"` // String value.
	Array  []string `json:"array,omitempty"`  // Array value.
}

type IndexedDB struct {
	target gcdmessage.ChromeTargeter
}

func NewIndexedDB(target gcdmessage.ChromeTargeter) *IndexedDB {
	c := &IndexedDB{target: target}
	return c
}

type IndexedDBClearObjectStoreParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// ClearObjectStoreWithParams - Clears all entries from an object store.
func (c *IndexedDB) ClearObjectStoreWithParams(ctx context.Context, v *IndexedDBClearObjectStoreParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.clearObjectStore", Params: v})
}

// ClearObjectStore - Clears all entries from an object store.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName - Database name.
// objectStoreName - Object store name.
func (c *IndexedDB) ClearObjectStore(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string, objectStoreName string) (*gcdmessage.ChromeResponse, error) {
	var v IndexedDBClearObjectStoreParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	v.ObjectStoreName = objectStoreName
	return c.ClearObjectStoreWithParams(ctx, &v)
}

type IndexedDBDeleteDatabaseParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

// DeleteDatabaseWithParams - Deletes a database.
func (c *IndexedDB) DeleteDatabaseWithParams(ctx context.Context, v *IndexedDBDeleteDatabaseParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.deleteDatabase", Params: v})
}

// DeleteDatabase - Deletes a database.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName - Database name.
func (c *IndexedDB) DeleteDatabase(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string) (*gcdmessage.ChromeResponse, error) {
	var v IndexedDBDeleteDatabaseParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	return c.DeleteDatabaseWithParams(ctx, &v)
}

type IndexedDBDeleteObjectStoreEntriesParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	//
	DatabaseName string `json:"databaseName"`
	//
	ObjectStoreName string `json:"objectStoreName"`
	// Range of entry keys to delete
	KeyRange *IndexedDBKeyRange `json:"keyRange"`
}

// DeleteObjectStoreEntriesWithParams - Delete a range of entries from an object store
func (c *IndexedDB) DeleteObjectStoreEntriesWithParams(ctx context.Context, v *IndexedDBDeleteObjectStoreEntriesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.deleteObjectStoreEntries", Params: v})
}

// DeleteObjectStoreEntries - Delete a range of entries from an object store
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName -
// objectStoreName -
// keyRange - Range of entry keys to delete
func (c *IndexedDB) DeleteObjectStoreEntries(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string, objectStoreName string, keyRange *IndexedDBKeyRange) (*gcdmessage.ChromeResponse, error) {
	var v IndexedDBDeleteObjectStoreEntriesParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	v.ObjectStoreName = objectStoreName
	v.KeyRange = keyRange
	return c.DeleteObjectStoreEntriesWithParams(ctx, &v)
}

// Disables events from backend.
func (c *IndexedDB) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.disable"})
}

// Enables events from backend.
func (c *IndexedDB) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.enable"})
}

type IndexedDBRequestDataParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
	// Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`
	// Number of records to skip.
	SkipCount int `json:"skipCount"`
	// Number of records to fetch.
	PageSize int `json:"pageSize"`
	// Key range.
	KeyRange *IndexedDBKeyRange `json:"keyRange,omitempty"`
}

// RequestDataWithParams - Requests data from object store or index.
// Returns -  objectStoreDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *IndexedDB) RequestDataWithParams(ctx context.Context, v *IndexedDBRequestDataParams) ([]*IndexedDBDataEntry, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestData", Params: v})
	if err != nil {
		return nil, false, err
	}

	var chromeData struct {
		Result struct {
			ObjectStoreDataEntries []*IndexedDBDataEntry
			HasMore                bool
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

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, err
	}

	return chromeData.Result.ObjectStoreDataEntries, chromeData.Result.HasMore, nil
}

// RequestData - Requests data from object store or index.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName - Database name.
// objectStoreName - Object store name.
// indexName - Index name, empty string for object store data requests.
// skipCount - Number of records to skip.
// pageSize - Number of records to fetch.
// keyRange - Key range.
// Returns -  objectStoreDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *IndexedDB) RequestData(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string, objectStoreName string, indexName string, skipCount int, pageSize int, keyRange *IndexedDBKeyRange) ([]*IndexedDBDataEntry, bool, error) {
	var v IndexedDBRequestDataParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	v.ObjectStoreName = objectStoreName
	v.IndexName = indexName
	v.SkipCount = skipCount
	v.PageSize = pageSize
	v.KeyRange = keyRange
	return c.RequestDataWithParams(ctx, &v)
}

type IndexedDBGetMetadataParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// GetMetadataWithParams - Gets metadata of an object store.
// Returns -  entriesCount - the entries count keyGeneratorValue - the current value of key generator, to become the next inserted key into the object store. Valid if objectStore.autoIncrement is true.
func (c *IndexedDB) GetMetadataWithParams(ctx context.Context, v *IndexedDBGetMetadataParams) (float64, float64, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.getMetadata", Params: v})
	if err != nil {
		return 0, 0, err
	}

	var chromeData struct {
		Result struct {
			EntriesCount      float64
			KeyGeneratorValue float64
		}
	}

	if resp == nil {
		return 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, err
	}

	return chromeData.Result.EntriesCount, chromeData.Result.KeyGeneratorValue, nil
}

// GetMetadata - Gets metadata of an object store.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName - Database name.
// objectStoreName - Object store name.
// Returns -  entriesCount - the entries count keyGeneratorValue - the current value of key generator, to become the next inserted key into the object store. Valid if objectStore.autoIncrement is true.
func (c *IndexedDB) GetMetadata(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string, objectStoreName string) (float64, float64, error) {
	var v IndexedDBGetMetadataParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	v.ObjectStoreName = objectStoreName
	return c.GetMetadataWithParams(ctx, &v)
}

type IndexedDBRequestDatabaseParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

// RequestDatabaseWithParams - Requests database with given name in given frame.
// Returns -  databaseWithObjectStores - Database with an array of object stores.
func (c *IndexedDB) RequestDatabaseWithParams(ctx context.Context, v *IndexedDBRequestDatabaseParams) (*IndexedDBDatabaseWithObjectStores, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestDatabase", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			DatabaseWithObjectStores *IndexedDBDatabaseWithObjectStores
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

	return chromeData.Result.DatabaseWithObjectStores, nil
}

// RequestDatabase - Requests database with given name in given frame.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// databaseName - Database name.
// Returns -  databaseWithObjectStores - Database with an array of object stores.
func (c *IndexedDB) RequestDatabase(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket, databaseName string) (*IndexedDBDatabaseWithObjectStores, error) {
	var v IndexedDBRequestDatabaseParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	v.DatabaseName = databaseName
	return c.RequestDatabaseWithParams(ctx, &v)
}

type IndexedDBRequestDatabaseNamesParams struct {
	// At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	// Storage key.
	StorageKey string `json:"storageKey,omitempty"`
	// Storage bucket. If not specified, it uses the default bucket.
	StorageBucket *StorageStorageBucket `json:"storageBucket,omitempty"`
}

// RequestDatabaseNamesWithParams - Requests database names for given security origin.
// Returns -  databaseNames - Database names for origin.
func (c *IndexedDB) RequestDatabaseNamesWithParams(ctx context.Context, v *IndexedDBRequestDatabaseNamesParams) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestDatabaseNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			DatabaseNames []string
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

	return chromeData.Result.DatabaseNames, nil
}

// RequestDatabaseNames - Requests database names for given security origin.
// securityOrigin - At least and at most one of securityOrigin, storageKey, or storageBucket must be specified. Security origin.
// storageKey - Storage key.
// storageBucket - Storage bucket. If not specified, it uses the default bucket.
// Returns -  databaseNames - Database names for origin.
func (c *IndexedDB) RequestDatabaseNames(ctx context.Context, securityOrigin string, storageKey string, storageBucket *StorageStorageBucket) ([]string, error) {
	var v IndexedDBRequestDatabaseNamesParams
	v.SecurityOrigin = securityOrigin
	v.StorageKey = storageKey
	v.StorageBucket = storageBucket
	return c.RequestDatabaseNamesWithParams(ctx, &v)
}
