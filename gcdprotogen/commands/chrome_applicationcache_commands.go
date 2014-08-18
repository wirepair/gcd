// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ApplicationCache commands.
// API Version: 1.1

package gcd


import (
	
	"encoding/json"
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
	var frameIds []*types.ChromeApplicationCacheFrameWithManifest 
	recvCh, _ := sendCustomReturn(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "ApplicationCache.getFramesWithManifests"})
	resp := <-recvCh

	var chromeData interface{}
	err := json.Unmarshal(resp.Data, &chromeData)
	if err != nil {
		cerr := &ChromeErrorResponse{}
		chromeError := json.Unmarshal(resp.Data, cerr)
		if chromeError == nil {
			return frameIds, &ChromeRequestErr{Resp: cerr}
		}
		return frameIds, err
	}

	m := chromeData.(map[string]interface{})
	if r, ok := m["result"]; ok {
		results := r.(map[string]interface{})
		frameIds = results["frameIds"].([]*types.ChromeApplicationCacheFrameWithManifest)
	}
	return frameIds, nil
}


// end commands with no parameters but special return types

