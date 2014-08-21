// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the IndexedDB commands.
// API Version: 1.1

package gcd


import (
	"github.com/wirepair/gcd/gcdprotogen/types"
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) IndexedDB() *ChromeIndexedDB {
	if c.indexeddb == nil {
		c.indexeddb = newChromeIndexedDB(c)
	}
	return c.indexeddb
}


type ChromeIndexedDB struct {
	target *ChromeTarget
}

func newChromeIndexedDB(target *ChromeTarget) *ChromeIndexedDB {
	c := &ChromeIndexedDB{target: target}
	return c
}

 
// Enables events from backend.
func (c *ChromeIndexedDB) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.enable"})
}
 
// Disables events from backend.
func (c *ChromeIndexedDB) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.disable"})
}


// clearObjectStore - Clears all entries from an object store.
// securityOrigin - Security origin.
// databaseName - Database name.
// objectStoreName - Object store name.
func (c *ChromeIndexedDB) ClearObjectStore(securityOrigin string, databaseName string, objectStoreName string, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	paramRequest["objectStoreName"] = objectStoreName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.clearObjectStore", Params: paramRequest})
}





// requestDatabaseNames - Requests database names for given security origin.
// Returns - 
// Database names for origin.
func (c *ChromeIndexedDB) RequestDatabaseNames(securityOrigin string, ) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["securityOrigin"] = securityOrigin
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.requestDatabaseNames", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			DatabaseNames []string 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.DatabaseNames, nil
}

// requestDatabase - Requests database with given name in given frame.
// Returns - 
// Database with an array of object stores.
func (c *ChromeIndexedDB) RequestDatabase(securityOrigin string, databaseName string, ) (*types.ChromeIndexedDBDatabaseWithObjectStores, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.requestDatabase", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			DatabaseWithObjectStores *types.ChromeIndexedDBDatabaseWithObjectStores 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.DatabaseWithObjectStores, nil
}

// requestData - Requests data from object store or index.
// Returns - 
// Array of object store data entries.
// If true, there are more entries to fetch in the given range.
func (c *ChromeIndexedDB) RequestData(securityOrigin string, databaseName string, objectStoreName string, indexName string, skipCount int, pageSize int, keyRange *types.ChromeIndexedDBKeyRange, ) ([]*types.ChromeIndexedDBDataEntry, bool, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["securityOrigin"] = securityOrigin
	paramRequest["databaseName"] = databaseName
	paramRequest["objectStoreName"] = objectStoreName
	paramRequest["indexName"] = indexName
	paramRequest["skipCount"] = skipCount
	paramRequest["pageSize"] = pageSize
	paramRequest["keyRange"] = keyRange
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IndexedDB.requestData", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			ObjectStoreDataEntries []*types.ChromeIndexedDBDataEntry 
			HasMore bool 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, false, &ChromeRequestErr{Resp: cerr}
		}
		return nil, false, err
	}

	return chromeData.Result.ObjectStoreDataEntries, chromeData.Result.HasMore, nil
}


