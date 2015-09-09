// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Worker commands.
// API Version: 1.1

package gcd

import ()

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

//
func (c *ChromeWorker) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.enable"})
}

//
func (c *ChromeWorker) Disable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.disable"})
}

// sendMessageToWorker -
// workerId -
// message -
func (c *ChromeWorker) SendMessageToWorker(workerId string, message string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.sendMessageToWorker", Params: paramRequest})
}

// connectToWorker -
// workerId -
func (c *ChromeWorker) ConnectToWorker(workerId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.connectToWorker", Params: paramRequest})
}

// disconnectFromWorker -
// workerId -
func (c *ChromeWorker) DisconnectFromWorker(workerId string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.disconnectFromWorker", Params: paramRequest})
}

// setAutoconnectToWorkers -
// value -
func (c *ChromeWorker) SetAutoconnectToWorkers(value bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Worker.setAutoconnectToWorkers", Params: paramRequest})
}
