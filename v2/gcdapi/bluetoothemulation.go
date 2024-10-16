// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains BluetoothEmulation functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Stores the manufacturer data
type BluetoothEmulationManufacturerData struct {
	Key  int    `json:"key"`  // Company identifier https://bitbucket.org/bluetooth-SIG/public/src/main/assigned_numbers/company_identifiers/company_identifiers.yaml https://usb.org/developers
	Data string `json:"data"` // Manufacturer-specific data (Encoded as a base64 string when passed over JSON)
}

// Stores the byte data of the advertisement packet sent by a Bluetooth device.
type BluetoothEmulationScanRecord struct {
	Name             string                                `json:"name,omitempty"`             //
	Uuids            []string                              `json:"uuids,omitempty"`            //
	Appearance       int                                   `json:"appearance,omitempty"`       // Stores the external appearance description of the device.
	TxPower          int                                   `json:"txPower,omitempty"`          // Stores the transmission power of a broadcasting device.
	ManufacturerData []*BluetoothEmulationManufacturerData `json:"manufacturerData,omitempty"` // Key is the company identifier and the value is an array of bytes of manufacturer specific data.
}

// Stores the advertisement packet information that is sent by a Bluetooth device.
type BluetoothEmulationScanEntry struct {
	DeviceAddress string                        `json:"deviceAddress"` //
	Rssi          int                           `json:"rssi"`          //
	ScanRecord    *BluetoothEmulationScanRecord `json:"scanRecord"`    //
}

type BluetoothEmulation struct {
	target gcdmessage.ChromeTargeter
}

func NewBluetoothEmulation(target gcdmessage.ChromeTargeter) *BluetoothEmulation {
	c := &BluetoothEmulation{target: target}
	return c
}

type BluetoothEmulationEnableParams struct {
	// State of the simulated central. enum values: absent, powered-off, powered-on
	State string `json:"state"`
}

// EnableWithParams - Enable the BluetoothEmulation domain.
func (c *BluetoothEmulation) EnableWithParams(ctx context.Context, v *BluetoothEmulationEnableParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "BluetoothEmulation.enable", Params: v})
}

// Enable - Enable the BluetoothEmulation domain.
// state - State of the simulated central. enum values: absent, powered-off, powered-on
func (c *BluetoothEmulation) Enable(ctx context.Context, state string) (*gcdmessage.ChromeResponse, error) {
	var v BluetoothEmulationEnableParams
	v.State = state
	return c.EnableWithParams(ctx, &v)
}

// Disable the BluetoothEmulation domain.
func (c *BluetoothEmulation) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "BluetoothEmulation.disable"})
}

type BluetoothEmulationSimulatePreconnectedPeripheralParams struct {
	//
	Address string `json:"address"`
	//
	Name string `json:"name"`
	//
	ManufacturerData []*BluetoothEmulationManufacturerData `json:"manufacturerData"`
	//
	KnownServiceUuids []string `json:"knownServiceUuids"`
}

// SimulatePreconnectedPeripheralWithParams - Simulates a peripheral with |address|, |name| and |knownServiceUuids| that has already been connected to the system.
func (c *BluetoothEmulation) SimulatePreconnectedPeripheralWithParams(ctx context.Context, v *BluetoothEmulationSimulatePreconnectedPeripheralParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "BluetoothEmulation.simulatePreconnectedPeripheral", Params: v})
}

// SimulatePreconnectedPeripheral - Simulates a peripheral with |address|, |name| and |knownServiceUuids| that has already been connected to the system.
// address -
// name -
// manufacturerData -
// knownServiceUuids -
func (c *BluetoothEmulation) SimulatePreconnectedPeripheral(ctx context.Context, address string, name string, manufacturerData []*BluetoothEmulationManufacturerData, knownServiceUuids []string) (*gcdmessage.ChromeResponse, error) {
	var v BluetoothEmulationSimulatePreconnectedPeripheralParams
	v.Address = address
	v.Name = name
	v.ManufacturerData = manufacturerData
	v.KnownServiceUuids = knownServiceUuids
	return c.SimulatePreconnectedPeripheralWithParams(ctx, &v)
}

type BluetoothEmulationSimulateAdvertisementParams struct {
	//
	Entry *BluetoothEmulationScanEntry `json:"entry"`
}

// SimulateAdvertisementWithParams - Simulates an advertisement packet described in |entry| being received by the central.
func (c *BluetoothEmulation) SimulateAdvertisementWithParams(ctx context.Context, v *BluetoothEmulationSimulateAdvertisementParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "BluetoothEmulation.simulateAdvertisement", Params: v})
}

// SimulateAdvertisement - Simulates an advertisement packet described in |entry| being received by the central.
// entry -
func (c *BluetoothEmulation) SimulateAdvertisement(ctx context.Context, entry *BluetoothEmulationScanEntry) (*gcdmessage.ChromeResponse, error) {
	var v BluetoothEmulationSimulateAdvertisementParams
	v.Entry = entry
	return c.SimulateAdvertisementWithParams(ctx, &v)
}
