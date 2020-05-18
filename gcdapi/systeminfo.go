// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains SystemInfo functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Describes a single graphics processor (GPU).
type SystemInfoGPUDevice struct {
	VendorId      float64 `json:"vendorId"`           // PCI ID of the GPU vendor, if available; 0 otherwise.
	DeviceId      float64 `json:"deviceId"`           // PCI ID of the GPU device, if available; 0 otherwise.
	SubSysId      float64 `json:"subSysId,omitempty"` // Sub sys ID of the GPU, only available on Windows.
	Revision      float64 `json:"revision,omitempty"` // Revision of the GPU, only available on Windows.
	VendorString  string  `json:"vendorString"`       // String description of the GPU vendor, if the PCI ID is not available.
	DeviceString  string  `json:"deviceString"`       // String description of the GPU device, if the PCI ID is not available.
	DriverVendor  string  `json:"driverVendor"`       // String description of the GPU driver vendor.
	DriverVersion string  `json:"driverVersion"`      // String description of the GPU driver version.
}

// Describes the width and height dimensions of an entity.
type SystemInfoSize struct {
	Width  int `json:"width"`  // Width in pixels.
	Height int `json:"height"` // Height in pixels.
}

// Describes a supported video decoding profile with its associated minimum and maximum resolutions.
type SystemInfoVideoDecodeAcceleratorCapability struct {
	Profile       string          `json:"profile"`       // Video codec profile that is supported, e.g. VP9 Profile 2.
	MaxResolution *SystemInfoSize `json:"maxResolution"` // Maximum video dimensions in pixels supported for this |profile|.
	MinResolution *SystemInfoSize `json:"minResolution"` // Minimum video dimensions in pixels supported for this |profile|.
}

// Describes a supported video encoding profile with its associated maximum resolution and maximum framerate.
type SystemInfoVideoEncodeAcceleratorCapability struct {
	Profile                 string          `json:"profile"`                 // Video codec profile that is supported, e.g H264 Main.
	MaxResolution           *SystemInfoSize `json:"maxResolution"`           // Maximum video dimensions in pixels supported for this |profile|.
	MaxFramerateNumerator   int             `json:"maxFramerateNumerator"`   // Maximum encoding framerate in frames per second supported for this |profile|, as fraction's numerator and denominator, e.g. 24/1 fps, 24000/1001 fps, etc.
	MaxFramerateDenominator int             `json:"maxFramerateDenominator"` //
}

// Describes a supported image decoding profile with its associated minimum and maximum resolutions and subsampling.
type SystemInfoImageDecodeAcceleratorCapability struct {
	ImageType     string          `json:"imageType"`     // Image coded, e.g. Jpeg. enum values: jpeg, webp, unknown
	MaxDimensions *SystemInfoSize `json:"maxDimensions"` // Maximum supported dimensions of the image in pixels.
	MinDimensions *SystemInfoSize `json:"minDimensions"` // Minimum supported dimensions of the image in pixels.
	Subsamplings  []string        `json:"subsamplings"`  // Optional array of supported subsampling formats, e.g. 4:2:0, if known. enum values: yuv420, yuv422, yuv444
}

// Provides information about the GPU(s) on the system.
type SystemInfoGPUInfo struct {
	Devices              []*SystemInfoGPUDevice                        `json:"devices"`                 // The graphics devices on the system. Element 0 is the primary GPU.
	AuxAttributes        map[string]interface{}                        `json:"auxAttributes,omitempty"` // An optional dictionary of additional GPU related attributes.
	FeatureStatus        map[string]interface{}                        `json:"featureStatus,omitempty"` // An optional dictionary of graphics features and their status.
	DriverBugWorkarounds []string                                      `json:"driverBugWorkarounds"`    // An optional array of GPU driver bug workarounds.
	VideoDecoding        []*SystemInfoVideoDecodeAcceleratorCapability `json:"videoDecoding"`           // Supported accelerated video decoding capabilities.
	VideoEncoding        []*SystemInfoVideoEncodeAcceleratorCapability `json:"videoEncoding"`           // Supported accelerated video encoding capabilities.
	ImageDecoding        []*SystemInfoImageDecodeAcceleratorCapability `json:"imageDecoding"`           // Supported accelerated image decoding capabilities.
}

// Represents process info.
type SystemInfoProcessInfo struct {
	Type    string  `json:"type"`    // Specifies process type.
	Id      int     `json:"id"`      // Specifies process id.
	CpuTime float64 `json:"cpuTime"` // Specifies cumulative CPU usage in seconds across all threads of the process since the process start.
}

type SystemInfo struct {
	target gcdmessage.ChromeTargeter
}

func NewSystemInfo(target gcdmessage.ChromeTargeter) *SystemInfo {
	c := &SystemInfo{target: target}
	return c
}

// GetInfo - Returns information about the system.
// Returns -  gpu - Information about the GPUs on the system. modelName - A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported. modelVersion - A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported. commandLine - The command line string used to launch the browser. Will be the empty string if not supported.
func (c *SystemInfo) GetInfo() (*SystemInfoGPUInfo, string, string, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "SystemInfo.getInfo"})
	if err != nil {
		return nil, "", "", "", err
	}

	var chromeData struct {
		Result struct {
			Gpu          *SystemInfoGPUInfo
			ModelName    string
			ModelVersion string
			CommandLine  string
		}
	}

	if resp == nil {
		return nil, "", "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, "", "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, "", "", "", err
	}

	return chromeData.Result.Gpu, chromeData.Result.ModelName, chromeData.Result.ModelVersion, chromeData.Result.CommandLine, nil
}

// GetProcessInfo - Returns information about all running processes.
// Returns -  processInfo - An array of process info blocks.
func (c *SystemInfo) GetProcessInfo() ([]*SystemInfoProcessInfo, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "SystemInfo.getProcessInfo"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ProcessInfo []*SystemInfoProcessInfo
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.ProcessInfo, nil
}
