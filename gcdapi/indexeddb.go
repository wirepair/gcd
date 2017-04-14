// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains IndexedDB functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Database with an array of object stores.
type IndexedDBDatabaseWithObjectStores struct {
	Name         string                  `json:"name"`         // Database name.
	Version      int                     `json:"version"`      // Database version.
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

// Enables events from backend.
func (c *IndexedDB) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.enable"})
}

// Disables events from backend.
func (c *IndexedDB) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.disable"})
}

// RequestDatabaseNames - Requests database names for given security origin.
// securityOrigin - Security origin.
// Returns -  databaseNames - Database names for origin.
func (c *IndexedDB) RequestDatabaseNames(securityOrigin string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["securityOrigin"] = securityOrigin
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestDatabaseNames", Params: paramRequest})
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

// RequestDatabase - Requests database with given name in given frame.
// securityOrigin - Security origin.
// databaseName - Database name.
// Returns -  databaseWithObjectStores - Database with an array of object stores.
func (c *IndexedDB) RequestDatabase(securityOrigin string, databaseName string) (*IndexedDBDatabaseWithObjectStores, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestDatabase", Params: paramRequest})
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

// RequestData - Requests data from object store or index.
// securityOrigin - Security origin.
// databaseName - Database name.
// objectStoreName - Object store name.
// indexName - Index name, empty string for object store data requests.
// skipCount - Number of records to skip.
// pageSize - Number of records to fetch.
// keyRange - Key range.
// Returns -  objectStoreDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *IndexedDB) RequestData(securityOrigin string, databaseName string, objectStoreName string, indexName string, skipCount int, pageSize int, keyRange *IndexedDBKeyRange) ([]*IndexedDBDataEntry, bool, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	paramRequest["objectStoreName"] = objectStoreName
	paramRequest["indexName"] = indexName
	paramRequest["skipCount"] = skipCount
	paramRequest["pageSize"] = pageSize
	paramRequest["keyRange"] = keyRange
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.requestData", Params: paramRequest})
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

// ClearObjectStore - Clears all entries from an object store.
// securityOrigin - Security origin.
// databaseName - Database name.
// objectStoreName - Object store name.
func (c *IndexedDB) ClearObjectStore(securityOrigin string, databaseName string, objectStoreName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	paramRequest["objectStoreName"] = objectStoreName
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.clearObjectStore", Params: paramRequest})
}

// DeleteDatabase - Deletes a database.
// securityOrigin - Security origin.
// databaseName - Database name.
func (c *IndexedDB) DeleteDatabase(securityOrigin string, databaseName string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IndexedDB.deleteDatabase", Params: paramRequest})
}
