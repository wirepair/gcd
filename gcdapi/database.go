// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Database functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Database object.
type DatabaseDatabase struct {
	Id      string `json:"id"`      // Database ID.
	Domain  string `json:"domain"`  // Database domain.
	Name    string `json:"name"`    // Database name.
	Version string `json:"version"` // Database version.
}

// Database error.
type DatabaseError struct {
	Message string `json:"message"` // Error message.
	Code    int    `json:"code"`    // Error code.
}

//
type DatabaseAddDatabaseEvent struct {
	Database *DatabaseDatabase `json:"database"` //
}

type Database struct {
	target gcdmessage.ChromeTargeter
}

func NewDatabase(target gcdmessage.ChromeTargeter) *Database {
	c := &Database{target: target}
	return c
}

// Enables database tracking, database events will now be delivered to the client.
func (c *Database) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.enable"})
}

// Disables database tracking, prevents database events from being sent to the client.
func (c *Database) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.disable"})
}

// GetDatabaseTableNames -
// databaseId -
// Returns -  tableNames -
func (c *Database) GetDatabaseTableNames(databaseId string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["databaseId"] = databaseId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.getDatabaseTableNames"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			TableNames []string
		}
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

	return chromeData.Result.TableNames, nil
}

// ExecuteSQL -
// databaseId -
// query -
// Returns -  columnNames -  values -  sqlError -
func (c *Database) ExecuteSQL(databaseId string, query string) ([]string, []string, *DatabaseError, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["databaseId"] = databaseId
	paramRequest["query"] = query
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.executeSQL"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ColumnNames []string
			Values      []string
			SqlError    *DatabaseError
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.ColumnNames, chromeData.Result.Values, chromeData.Result.SqlError, nil
}
