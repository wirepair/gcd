// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Power types.
// API Version: 1.1
package main

 
// PowerEvent item
type ChromePowerPowerEvent struct {
	Type string `json:"type"` // Power Event Type.
	Timestamp int `json:"timestamp"` // Power Event Time, in milliseconds.
	Value int `json:"value"` // Power Event Value.
}
 
