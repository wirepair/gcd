// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Database types.
// API Version: 1.1
package types

 
// Unique identifier of Database object.
type ChromeDatabaseDatabaseId string 
 
 
// Database object.
type ChromeDatabaseDatabase struct {
	Id *ChromeDatabaseDatabaseId `json:"id"` // Database ID.
	Domain string `json:"domain"` // Database domain.
	Name string `json:"name"` // Database name.
	Version string `json:"version"` // Database version.
}
 
 
// Database error.
type ChromeDatabaseError struct {
	Message string `json:"message"` // Error message.
	Code int `json:"code"` // Error code.
}
 
