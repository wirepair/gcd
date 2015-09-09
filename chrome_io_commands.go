// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the IO commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) IO() *ChromeIO {
	if c.io == nil {
		c.io = newChromeIO(c)
	}
	return c.io
}

type ChromeIO struct {
	target *ChromeTarget
}

func newChromeIO(target *ChromeTarget) *ChromeIO {
	c := &ChromeIO{target: target}
	return c
}

// close - Close the stream, discard any temporary backing storage.
// handle - Handle of the stream to close.
func (c *ChromeIO) Close(handle *types.ChromeIOStreamHandle) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["handle"] = handle
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IO.close", Params: paramRequest})
}

// read - Read a chunk of the stream
// Returns -
// Data that were read.
// Set if the end-of-file condition occured while reading.
func (c *ChromeIO) Read(handle *types.ChromeIOStreamHandle, offset int, size int) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["handle"] = handle
	paramRequest["offset"] = offset
	paramRequest["size"] = size
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "IO.read", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			Data string
			Eof  bool
		}
	}

	// test if error first
	cerr := &ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", false, err
	}

	return chromeData.Result.Data, chromeData.Result.Eof, nil
}
