// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Worker commands.
// API Version: 1.1

package gcd


import (
	
	"encoding/json"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Worker() *ChromeWorker {
	if c.worker == nil {
		c.worker = newChromeWorker(c)
	}
	return c.worker
}


type ChromeWorker struct {
	target *ChromeTarget
}

func newChromeWorker(target *ChromeTarget) *ChromeWorker {
	c := &ChromeWorker{target: target}
	return c
}

// start non parameterized commands 
// 
func (c *ChromeWorker) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.enable"})
}
 
// 
func (c *ChromeWorker) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.disable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// sendMessageToWorker - 
// workerId - 
// message - 
func (c *ChromeWorker) SendMessageToWorker(workerId int, message interface{}, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.sendMessageToWorker", Params: paramRequest})
}

// connectToWorker - 
// workerId - 
func (c *ChromeWorker) ConnectToWorker(workerId int, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.connectToWorker", Params: paramRequest})
}

// disconnectFromWorker - 
// workerId - 
func (c *ChromeWorker) DisconnectFromWorker(workerId int, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.disconnectFromWorker", Params: paramRequest})
}

// setAutoconnectToWorkers - 
// value - 
func (c *ChromeWorker) SetAutoconnectToWorkers(value bool, ) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.setAutoconnectToWorkers", Params: paramRequest})
}


// end parameterized commands with no special return types


// start commands with no parameters but special return types

// canInspectWorkers - Tells whether browser supports workers inspection.
// Returns - 
// True if browser has workers support.
func (c *ChromeWorker) CanInspectWorkers() (bool, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.canInspectWorkers"})
	resp := <-recvCh

	var chromeData struct {
		Result struct { 
			Result bool 
		}
	}
		
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return false, &ChromeRequestErr{Resp: cerr}
		}
		return false, err
	}

	return chromeData.Result.Result, nil
}


// end commands with no parameters but special return types


// start commands with parameters and special return types


// end commands with parameters and special return types

