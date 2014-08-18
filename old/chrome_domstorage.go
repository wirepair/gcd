package gcd

import (
	"encoding/json"
	"fmt"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

func (c *ChromeTarget) DOMStorage() *ChromeDOMStorage {
	if c.domstorage == nil {
		c.domstorage = newChromeDOMStorage(c)
	}
	return c.domstorage
}

type ChromeDOMStorage struct {
	target *ChromeTarget
}

func newChromeDOMStorage(target *ChromeTarget) *ChromeDOMStorage {
	c := &ChromeDOMStorage{target: target}
	return c
}

func (c *ChromeDOMStorage) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.enable"})
}

func (c *ChromeDOMStorage) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "DOMStorage.enable"})
}

func (c *ChromeDOMStorage) GetDOMStorageItems(storageID *types.ChromeDOMStorageStorageId) (*types.ChromeDOMStorageItem, error) {
	id := c.target.getId()
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["storageId"] = storageID
	req := &ParamRequest{Id: id, Method: "DOMStorage.getDOMStorageItems", Params: paramRequest}
	recvCh, err := sendCustomReturn(c.target.sendCh, req)
	if err != nil {
		return nil, err
	}
	resp := <-recvCh
	fmt.Printf("%s\n", string(resp.Data))
	response := &types.ChromeDOMStorageItem{}
	err = json.Unmarshal(resp.Data, response)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}
	return response, nil
}
