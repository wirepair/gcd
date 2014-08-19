// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ApplicationCache commands.
// API Version: 1.1

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) ApplicationCache() *ChromeApplicationCache {
	if c.applicationcache == nil {
		c.applicationcache = newChromeApplicationCache(c)
	}
	return c.applicationcache
}

type ChromeApplicationCache struct {
	target *ChromeTarget
}

func newChromeApplicationCache(target *ChromeTarget) *ChromeApplicationCache {
	c := &ChromeApplicationCache{target: target}
	return c
}

// start non parameterized commands
// Enables application cache domain notifications.
func (c *ChromeApplicationCache) Enable() (*ChromeResponse, error) {
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ApplicationCache.enable"})
}

// end non parameterized commands

// start parameterized commands with no special return types

// end parameterized commands with no special return types

// start commands with no parameters but special return types

// getFramesWithManifests - Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
// Returns -
// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (c *ChromeApplicationCache) GetFramesWithManifests() ([]*types.ChromeApplicationCacheFrameWithManifest, error) {
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ApplicationCache.getFramesWithManifests"})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			FrameIds []*types.ChromeApplicationCacheFrameWithManifest
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.FrameIds, nil
}

// end commands with no parameters but special return types

// start commands with parameters and special return types

// getManifestForFrame - Returns manifest URL for document in the given frame.
// Returns -
// Manifest URL for document in the given frame.
func (c *ChromeApplicationCache) GetManifestForFrame(frameId *types.ChromePageFrameId) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ApplicationCache.getManifestForFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ManifestURL string
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return "", &ChromeRequestErr{Resp: cerr}
		}
		return "", err
	}

	return chromeData.Result.ManifestURL, nil
}

// getApplicationCacheForFrame - Returns relevant application cache data for the document in given frame.
// Returns -
// Relevant application cache data for the document in given frame.
func (c *ChromeApplicationCache) GetApplicationCacheForFrame(frameId *types.ChromePageFrameId) (*types.ChromeApplicationCacheApplicationCache, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ApplicationCache.getApplicationCacheForFrame", Params: paramRequest})
	resp := <-recvCh

	var chromeData struct {
		Result struct {
			ApplicationCache *types.ChromeApplicationCacheApplicationCache
		}
	}

	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return nil, &ChromeRequestErr{Resp: cerr}
		}
		return nil, err
	}

	return chromeData.Result.ApplicationCache, nil
}

// end commands with parameters and special return types
