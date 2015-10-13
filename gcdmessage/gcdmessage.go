/*
The MIT License (MIT)

Copyright (c) 2015 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// This package contains messaging types and functions between the API and our gcd library.

package gcdmessage

import (
	"encoding/json"
	"strconv"
)

type ChromeTargeter interface {
	GetId() int64
	GetSendCh() chan *Message
}

// An internal message object used for components and ChromeTarget to communicate back and forth
type Message struct {
	ReplyCh chan *Message  // json response channel
	Id      int64          // id to map response channels to send chans
	Data    []byte         // the data for the websocket to send/recv
	Method  string         // event name type.
	Target  ChromeTargeter // reference to the ChromeTarget for events
}

// default response object, contains the id and a result if applicable.
type ChromeResponse struct {
	Id     int64       `json:"id"`
	Result interface{} `json:"result"`
}

// default no-arg request
type ChromeRequest struct {
	Id     int64       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

// default chrome error response to an invalid request.
type ChromeErrorResponse struct {
	Id    int64        `json:"id"`    // the request Id that this is a response of
	Error *ChromeError `json:"error"` // the error object
}

// An error object returned from a request
type ChromeError struct {
	Code    int64  `json:"code"`    // the error code
	Message string `json:"message"` // the error message
}

// A gcd type for reporting chrome request errors
type ChromeRequestErr struct {
	Resp *ChromeErrorResponse // a ref to the error response to be used to generate the user friendly error string
}

// user friendly error response
func (cerr *ChromeRequestErr) Error() string {
	return "request " + strconv.FormatInt(cerr.Resp.Id, 10) + " failed, code: " + strconv.FormatInt(cerr.Resp.Error.Code, 10) + " msg: " + cerr.Resp.Error.Message
}

// When a chrometarget crashes and we have to close response channels and return nil
type ChromeEmptyResponseErr struct {
}

func (cerr *ChromeEmptyResponseErr) Error() string {
	return "nil response received"
}

// default request object that has parameters.
type ParamRequest struct {
	Id     int64       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

// Takes in a ParamRequest and gives back a response channel so the caller can decode as necessary.
func SendCustomReturn(sendCh chan<- *Message, paramRequest *ParamRequest) (<-chan *Message, error) {
	data, err := json.Marshal(paramRequest)
	if err != nil {
		return nil, err
	}

	recvCh := make(chan *Message)
	sendMsg := &Message{ReplyCh: recvCh, Id: paramRequest.Id, Data: []byte(data)}
	sendCh <- sendMsg
	return recvCh, nil
}

// Sends a generic request that gets back a generic response, or error. This returns a ChromeResponse
// object.
func SendDefaultRequest(sendCh chan<- *Message, paramRequest *ParamRequest) (*ChromeResponse, error) {
	req := &ChromeRequest{Id: paramRequest.Id, Method: paramRequest.Method, Params: paramRequest.Params}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	recvCh := make(chan *Message)
	sendMsg := &Message{ReplyCh: recvCh, Id: paramRequest.Id, Data: []byte(data)}
	sendCh <- sendMsg
	resp := <-recvCh
	chromeResponse := &ChromeResponse{}
	err = json.Unmarshal(resp.Data, chromeResponse)
	if err != nil {
		return nil, err
	}
	return chromeResponse, nil
}
