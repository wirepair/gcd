// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the IndexedDB types.
// API Version: 1.1
package types

 
// Database with an array of object stores.
type ChromeIndexedDBDatabaseWithObjectStores struct {
	Name string `json:"name"` // Database name.
	Version string `json:"version"` // Deprecated string database version.
	IntVersion int `json:"intVersion"` // Integer database version.
	ObjectStores []*ChromeIndexedDBObjectStore `json:"objectStores"` // Object stores in this database.
}
 
 
// Object store.
type ChromeIndexedDBObjectStore struct {
	Name string `json:"name"` // Object store name.
	KeyPath *ChromeIndexedDBKeyPath `json:"keyPath"` // Object store key path.
	AutoIncrement bool `json:"autoIncrement"` // If true, object store has auto increment flag set.
	Indexes []*ChromeIndexedDBObjectStoreIndex `json:"indexes"` // Indexes in this object store.
}
 
 
// Object store index.
type ChromeIndexedDBObjectStoreIndex struct {
	Name string `json:"name"` // Index name.
	KeyPath *ChromeIndexedDBKeyPath `json:"keyPath"` // Index key path.
	Unique bool `json:"unique"` // If true, index is unique.
	MultiEntry bool `json:"multiEntry"` // If true, index allows multiple entries for a key.
}
 
 
// Key.
type ChromeIndexedDBKey struct {
	Type string `json:"type"` // Key type.
	Number float64 `json:"number,omitempty"` // Number value.
	String string `json:"string,omitempty"` // String value.
	Date float64 `json:"date,omitempty"` // Date value.
	Array []*ChromeIndexedDBKey `json:"array,omitempty"` // Array value.
}
 
 
// Key range.
type ChromeIndexedDBKeyRange struct {
	Lower *ChromeIndexedDBKey `json:"lower,omitempty"` // Lower bound.
	Upper *ChromeIndexedDBKey `json:"upper,omitempty"` // Upper bound.
	LowerOpen bool `json:"lowerOpen"` // If true lower bound is open.
	UpperOpen bool `json:"upperOpen"` // If true upper bound is open.
}
 
 
// Data entry.
type ChromeIndexedDBDataEntry struct {
	Key string `json:"key"` // JSON-stringified key object.
	PrimaryKey string `json:"primaryKey"` // JSON-stringified primary key object.
	Value string `json:"value"` // JSON-stringified value object.
}
 
 
// Key path.
type ChromeIndexedDBKeyPath struct {
	Type string `json:"type"` // Key path type.
	String string `json:"string,omitempty"` // String value.
	Array []string `json:"array,omitempty"` // Array value.
}
 
