// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Worker functionality.
// API Version: 1.1

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

//
type WorkerWorkerCreatedEvent struct {
	WorkerId           string `json:"workerId"`           //
	Url                string `json:"url"`                //
	InspectorConnected bool   `json:"inspectorConnected"` //
}

//
type WorkerWorkerTerminatedEvent struct {
	WorkerId string `json:"workerId"` //
}

//
type WorkerDispatchMessageFromWorkerEvent struct {
	WorkerId string `json:"workerId"` //
	Message  string `json:"message"`  //
}

type Worker struct {
	target gcdmessage.ChromeTargeter
}

func NewWorker(target gcdmessage.ChromeTargeter) *Worker {
	c := &Worker{target: target}
	return c
}

//
func (c *Worker) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.enable"})
}

//
func (c *Worker) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.disable"})
}

// SendMessageToWorker -
// workerId -
// message -
func (c *Worker) SendMessageToWorker(workerId string, message string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["workerId"] = workerId
	paramRequest["message"] = message
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.sendMessageToWorker", Params: paramRequest})
}

// ConnectToWorker -
// workerId -
func (c *Worker) ConnectToWorker(workerId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.connectToWorker", Params: paramRequest})
}

// DisconnectFromWorker -
// workerId -
func (c *Worker) DisconnectFromWorker(workerId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["workerId"] = workerId
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.disconnectFromWorker", Params: paramRequest})
}

// SetAutoconnectToWorkers -
// value -
func (c *Worker) SetAutoconnectToWorkers(value bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Worker.setAutoconnectToWorkers", Params: paramRequest})
}
