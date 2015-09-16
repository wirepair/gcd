// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains ApplicationCache functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Detailed application cache resource information.
type ApplicationCacheApplicationCacheResource struct {
	Url  string `json:"url"`  // Resource url.
	Size int    `json:"size"` // Resource size.
	Type string `json:"type"` // Resource type.
}

// Detailed application cache information.
type ApplicationCacheApplicationCache struct {
	ManifestURL  string                                      `json:"manifestURL"`  // Manifest URL.
	Size         float64                                     `json:"size"`         // Application cache size.
	CreationTime float64                                     `json:"creationTime"` // Application cache creation time.
	UpdateTime   float64                                     `json:"updateTime"`   // Application cache update time.
	Resources    []*ApplicationCacheApplicationCacheResource `json:"resources"`    // Application cache resources.
}

// Frame identifier - manifest URL pair.
type ApplicationCacheFrameWithManifest struct {
	FrameId     string `json:"frameId"`     // Frame identifier.
	ManifestURL string `json:"manifestURL"` // Manifest URL.
	Status      int    `json:"status"`      // Application cache status.
}

//
type ApplicationCacheApplicationCacheStatusUpdatedEvent struct {
	FrameId     string `json:"frameId"`     // Identifier of the frame containing document whose application cache updated status.
	ManifestURL string `json:"manifestURL"` // Manifest URL.
	Status      int    `json:"status"`      // Updated application cache status.
}

//
type ApplicationCacheNetworkStateUpdatedEvent struct {
	IsNowOnline bool `json:"isNowOnline"` //
}

type ApplicationCache struct {
	target gcdmessage.ChromeTargeter
}

func NewApplicationCache(target gcdmessage.ChromeTargeter) *ApplicationCache {
	c := &ApplicationCache{target: target}
	return c
}

// GetFramesWithManifests - Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
// Returns -  frameIds - Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (c *ApplicationCache) GetFramesWithManifests() ([]*ApplicationCacheFrameWithManifest, error) {
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getFramesWithManifests"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameIds []*ApplicationCacheFrameWithManifest
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.FrameIds, nil
}

// Enables application cache domain notifications.
func (c *ApplicationCache) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.enable"})
}

// GetManifestForFrame - Returns manifest URL for document in the given frame.
// frameId - Identifier of the frame containing document whose manifest is retrieved.
// Returns -  manifestURL - Manifest URL for document in the given frame.
func (c *ApplicationCache) GetManifestForFrame(frameId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getManifestForFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ManifestURL string
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return "", err
	}

	return chromeData.Result.ManifestURL, nil
}

// GetApplicationCacheForFrame - Returns relevant application cache data for the document in given frame.
// frameId - Identifier of the frame containing document whose application cache is retrieved.
// Returns -  applicationCache - Relevant application cache data for the document in given frame.
func (c *ApplicationCache) GetApplicationCacheForFrame(frameId string) (*ApplicationCacheApplicationCache, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := gcdmessage.SendCustomReturn(c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getApplicationCacheForFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ApplicationCache *ApplicationCacheApplicationCache
		}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		return nil, err
	}

	return chromeData.Result.ApplicationCache, nil
}
