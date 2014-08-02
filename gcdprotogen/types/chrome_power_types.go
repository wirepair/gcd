// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Power types.
// API Version: 1.1
package types

 
// PowerEvent item
type ChromePowerPowerEvent struct {
	Type string `json:"type"` // Power Event Type.
	Timestamp float `json:"timestamp"` // Power Event Time, in milliseconds.
	Value float `json:"value"` // Power Event Value.
}
 
