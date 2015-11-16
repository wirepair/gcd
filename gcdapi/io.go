// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains IO functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

type IO struct {
	target gcdmessage.ChromeTargeter
}

func NewIO(target gcdmessage.ChromeTargeter) *IO {
	c := &IO{target: target}
	return c
}

// Read - Read a chunk of the stream
// handle - Handle of the stream to read.
// offset - Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
// size - Maximum number of bytes to read (left upon the agent discretion if not specified).
// Returns -  data - Data that were read. eof - Set if the end-of-file condition occured while reading.
func (c *IO) Read(handle string, offset int, size int) (string, bool, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["handle"] = handle
	paramRequest["offset"] = offset
	paramRequest["size"] = size
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.read", Params: paramRequest})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Data string
			Eof  bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.Data, chromeData.Result.Eof, nil
}

// Close - Close the stream, discard any temporary backing storage.
// handle - Handle of the stream to close.
func (c *IO) Close(handle string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["handle"] = handle
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.close", Params: paramRequest})
}
