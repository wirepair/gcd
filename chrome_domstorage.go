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
	return sendDOMStorageRequest(c.target.sendCh, c.target.getId(), "DOMStorage.enable")
}

func (c *ChromeDOMStorage) Disable() (*ChromeResponse, error) {
	return sendDOMStorageRequest(c.target.sendCh, c.target.getId(), "DOMStorage.disable")
}

func (c *ChromeDOMStorage) GetDOMStorageItems(storageID *types.ChromeDOMStorageStorageId) (*types.ChromeDOMStorageItem, error) {
	sendCh := c.target.sendCh
	id := c.target.getId()
	paramRequest := make(map[string]interface{}, 0)
	paramRequest["storageId"] = storageID
	req := &ParamRequest{Id: id, Method: "DOMStorage.getDOMStorageItems", Params: paramRequest}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Will send: %s\n", data)
	recvCh := make(chan *gcdMessage)
	sendMsg := &gcdMessage{ReplyCh: recvCh, Id: id, Data: []byte(data)}
	sendCh <- sendMsg
	resp := <-recvCh
	response := &types.ChromeDOMStorageItem{}
	fmt.Printf("GetDOMStorageItems: %s\n", string(resp.Data))
	err = json.Unmarshal(resp.Data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func sendDOMStorageRequest(sendCh chan<- *gcdMessage, id int64, method string) (*ChromeResponse, error) {
	req := &ChromeRequest{Id: id, Method: method}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	recvCh := make(chan *gcdMessage)
	sendMsg := &gcdMessage{ReplyCh: recvCh, Id: id, Data: []byte(data)}
	sendCh <- sendMsg
	resp := <-recvCh
	consoleResponse := &ChromeResponse{}
	err = json.Unmarshal(resp.Data, consoleResponse)
	if err != nil {
		return nil, err
	}
	return consoleResponse, nil
}
