// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Database commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Database() *ChromeDatabase {
	if c.database == nil {
		c.database = newChromeDatabase(c)
	}
	return c.database
}

type ChromeDatabase struct {
	target *ChromeTarget
}

func newChromeDatabase(target *ChromeTarget) *ChromeDatabase {
	c := &ChromeDatabase{target: target}
	return c
}

// start non parameterized commands
// Enables database tracking, database events will now be delivered to the client.
func (c *ChromeDatabase) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.enable"})
}

// Disables database tracking, prevents database events from being sent to the client.
func (c *ChromeDatabase) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// end parameterized commands with no special return types

// start commands with no parameters but special return types

// end commands with no parameters but special return types

// start commands with parameters and special return types

// getDatabaseTableNames -
// Returns -
func (c *ChromeDatabase) GetDatabaseTableNames(databaseId *types.ChromeDatabaseDatabaseId) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["databaseId"] = databaseId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.getDatabaseTableNames", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			TableNames []string
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

	return chromeData.Result.TableNames, nil
}

// executeSQL -
// Returns -
func (c *ChromeDatabase) ExecuteSQL(databaseId *types.ChromeDatabaseDatabaseId, query string) ([]string, []string, *types.ChromeDatabaseError, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["databaseId"] = databaseId
	paramRequest["query"] = query
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Database.executeSQL", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ColumnNames []string
			Values      []string
			SqlError    *types.ChromeDatabaseError
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, nil, nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, nil, nil, err
	}

	return chromeData.Result.ColumnNames, chromeData.Result.Values, chromeData.Result.SqlError, nil
}

// end commands with parameters and special return types
