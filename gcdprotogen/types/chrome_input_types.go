// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Input types.
// API Version: 1.1
package main

 
// 
type ChromeInputTouchPoint struct {
	State string `json:"state"` // State of the touch point.
	X int `json:"x"` // X coordinate of the event relative to the main frame's viewport.
	Y int `json:"y"` // Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX int `json:"radiusX,omitempty"` // X radius of the touch area (default: 1).
	RadiusY int `json:"radiusY,omitempty"` // Y radius of the touch area (default: 1).
	RotationAngle int `json:"rotationAngle,omitempty"` // Rotation angle (default: 0.0).
	Force int `json:"force,omitempty"` // Force (default: 1.0).
	Id int `json:"id,omitempty"` // Identifier used to track touch sources between events, must be unique within an event.
}
 
