// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Storage functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Usage for a storage type.
type StorageUsageForType struct {
	StorageType string  `json:"storageType"` // Name of storage type. enum values: appcache, cookies, file_systems, indexeddb, local_storage, shader_cache, websql, service_workers, cache_storage, interest_groups, all, other
	Usage       float64 `json:"usage"`       // Storage usage (bytes).
}

// Pair of issuer origin and number of available (signed, but not used) Trust Tokens from that issuer.
type StorageTrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"` //
	Count        float64 `json:"count"`        //
}

// Ad advertising element inside an interest group.
type StorageInterestGroupAd struct {
	RenderUrl string `json:"renderUrl"`          //
	Metadata  string `json:"metadata,omitempty"` //
}

// The full details of an interest group.
type StorageInterestGroupDetails struct {
	OwnerOrigin               string                    `json:"ownerOrigin"`                        //
	Name                      string                    `json:"name"`                               //
	ExpirationTime            float64                   `json:"expirationTime"`                     //
	JoiningOrigin             string                    `json:"joiningOrigin"`                      //
	BiddingUrl                string                    `json:"biddingUrl,omitempty"`               //
	BiddingWasmHelperUrl      string                    `json:"biddingWasmHelperUrl,omitempty"`     //
	UpdateUrl                 string                    `json:"updateUrl,omitempty"`                //
	TrustedBiddingSignalsUrl  string                    `json:"trustedBiddingSignalsUrl,omitempty"` //
	TrustedBiddingSignalsKeys []string                  `json:"trustedBiddingSignalsKeys"`          //
	UserBiddingSignals        string                    `json:"userBiddingSignals,omitempty"`       //
	Ads                       []*StorageInterestGroupAd `json:"ads"`                                //
	AdComponents              []*StorageInterestGroupAd `json:"adComponents"`                       //
}

// A cache's contents have been modified.
type StorageCacheStorageContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin    string `json:"origin"`    // Origin to update.
		CacheName string `json:"cacheName"` // Name of cache in origin.
	} `json:"Params,omitempty"`
}

// A cache has been added/deleted.
type StorageCacheStorageListUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin string `json:"origin"` // Origin to update.
	} `json:"Params,omitempty"`
}

// The origin's IndexedDB object store has been modified.
type StorageIndexedDBContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin          string `json:"origin"`          // Origin to update.
		StorageKey      string `json:"storageKey"`      // Storage key to update.
		DatabaseName    string `json:"databaseName"`    // Database to update.
		ObjectStoreName string `json:"objectStoreName"` // ObjectStore to update.
	} `json:"Params,omitempty"`
}

// The origin's IndexedDB database list has been modified.
type StorageIndexedDBListUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin     string `json:"origin"`     // Origin to update.
		StorageKey string `json:"storageKey"` // Storage key to update.
	} `json:"Params,omitempty"`
}

// One of the interest groups was accessed by the associated page.
type StorageInterestGroupAccessedEvent struct {
	Method string `json:"method"`
	Params struct {
		AccessTime  float64 `json:"accessTime"`  //
		Type        string  `json:"type"`        //  enum values: join, leave, update, bid, win
		OwnerOrigin string  `json:"ownerOrigin"` //
		Name        string  `json:"name"`        //
	} `json:"Params,omitempty"`
}

type Storage struct {
	target gcdmessage.ChromeTargeter
}

func NewStorage(target gcdmessage.ChromeTargeter) *Storage {
	c := &Storage{target: target}
	return c
}

type StorageGetStorageKeyForFrameParams struct {
	//
	FrameId string `json:"frameId"`
}

// GetStorageKeyForFrameWithParams - Returns a storage key given a frame id.
// Returns -  storageKey -
func (c *Storage) GetStorageKeyForFrameWithParams(ctx context.Context, v *StorageGetStorageKeyForFrameParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getStorageKeyForFrame", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			StorageKey string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.StorageKey, nil
}

// GetStorageKeyForFrame - Returns a storage key given a frame id.
// frameId -
// Returns -  storageKey -
func (c *Storage) GetStorageKeyForFrame(ctx context.Context, frameId string) (string, error) {
	var v StorageGetStorageKeyForFrameParams
	v.FrameId = frameId
	return c.GetStorageKeyForFrameWithParams(ctx, &v)
}

type StorageClearDataForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
	// Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// ClearDataForOriginWithParams - Clears storage for origin.
func (c *Storage) ClearDataForOriginWithParams(ctx context.Context, v *StorageClearDataForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearDataForOrigin", Params: v})
}

// ClearDataForOrigin - Clears storage for origin.
// origin - Security origin.
// storageTypes - Comma separated list of StorageType to clear.
func (c *Storage) ClearDataForOrigin(ctx context.Context, origin string, storageTypes string) (*gcdmessage.ChromeResponse, error) {
	var v StorageClearDataForOriginParams
	v.Origin = origin
	v.StorageTypes = storageTypes
	return c.ClearDataForOriginWithParams(ctx, &v)
}

type StorageClearDataForStorageKeyParams struct {
	// Storage key.
	StorageKey string `json:"storageKey"`
	// Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// ClearDataForStorageKeyWithParams - Clears storage for storage key.
func (c *Storage) ClearDataForStorageKeyWithParams(ctx context.Context, v *StorageClearDataForStorageKeyParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearDataForStorageKey", Params: v})
}

// ClearDataForStorageKey - Clears storage for storage key.
// storageKey - Storage key.
// storageTypes - Comma separated list of StorageType to clear.
func (c *Storage) ClearDataForStorageKey(ctx context.Context, storageKey string, storageTypes string) (*gcdmessage.ChromeResponse, error) {
	var v StorageClearDataForStorageKeyParams
	v.StorageKey = storageKey
	v.StorageTypes = storageTypes
	return c.ClearDataForStorageKeyWithParams(ctx, &v)
}

type StorageGetCookiesParams struct {
	// Browser context to use when called on the browser endpoint.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// GetCookiesWithParams - Returns all browser cookies.
// Returns -  cookies - Array of cookie objects.
func (c *Storage) GetCookiesWithParams(ctx context.Context, v *StorageGetCookiesParams) ([]*NetworkCookie, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getCookies", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

// GetCookies - Returns all browser cookies.
// browserContextId - Browser context to use when called on the browser endpoint.
// Returns -  cookies - Array of cookie objects.
func (c *Storage) GetCookies(ctx context.Context, browserContextId string) ([]*NetworkCookie, error) {
	var v StorageGetCookiesParams
	v.BrowserContextId = browserContextId
	return c.GetCookiesWithParams(ctx, &v)
}

type StorageSetCookiesParams struct {
	// Cookies to be set.
	Cookies []*NetworkCookieParam `json:"cookies"`
	// Browser context to use when called on the browser endpoint.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// SetCookiesWithParams - Sets given cookies.
func (c *Storage) SetCookiesWithParams(ctx context.Context, v *StorageSetCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setCookies", Params: v})
}

// SetCookies - Sets given cookies.
// cookies - Cookies to be set.
// browserContextId - Browser context to use when called on the browser endpoint.
func (c *Storage) SetCookies(ctx context.Context, cookies []*NetworkCookieParam, browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetCookiesParams
	v.Cookies = cookies
	v.BrowserContextId = browserContextId
	return c.SetCookiesWithParams(ctx, &v)
}

type StorageClearCookiesParams struct {
	// Browser context to use when called on the browser endpoint.
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// ClearCookiesWithParams - Clears cookies.
func (c *Storage) ClearCookiesWithParams(ctx context.Context, v *StorageClearCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearCookies", Params: v})
}

// ClearCookies - Clears cookies.
// browserContextId - Browser context to use when called on the browser endpoint.
func (c *Storage) ClearCookies(ctx context.Context, browserContextId string) (*gcdmessage.ChromeResponse, error) {
	var v StorageClearCookiesParams
	v.BrowserContextId = browserContextId
	return c.ClearCookiesWithParams(ctx, &v)
}

type StorageGetUsageAndQuotaParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// GetUsageAndQuotaWithParams - Returns usage and quota in bytes.
// Returns -  usage - Storage usage (bytes). quota - Storage quota (bytes). overrideActive - Whether or not the origin has an active storage quota override usageBreakdown - Storage usage per type (bytes).
func (c *Storage) GetUsageAndQuotaWithParams(ctx context.Context, v *StorageGetUsageAndQuotaParams) (float64, float64, bool, []*StorageUsageForType, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getUsageAndQuota", Params: v})
	if err != nil {
		return 0, 0, false, nil, err
	}

	var chromeData struct {
		Result struct {
			Usage          float64
			Quota          float64
			OverrideActive bool
			UsageBreakdown []*StorageUsageForType
		}
	}

	if resp == nil {
		return 0, 0, false, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, false, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, false, nil, err
	}

	return chromeData.Result.Usage, chromeData.Result.Quota, chromeData.Result.OverrideActive, chromeData.Result.UsageBreakdown, nil
}

// GetUsageAndQuota - Returns usage and quota in bytes.
// origin - Security origin.
// Returns -  usage - Storage usage (bytes). quota - Storage quota (bytes). overrideActive - Whether or not the origin has an active storage quota override usageBreakdown - Storage usage per type (bytes).
func (c *Storage) GetUsageAndQuota(ctx context.Context, origin string) (float64, float64, bool, []*StorageUsageForType, error) {
	var v StorageGetUsageAndQuotaParams
	v.Origin = origin
	return c.GetUsageAndQuotaWithParams(ctx, &v)
}

type StorageOverrideQuotaForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
	// The quota size (in bytes) to override the original quota with. If this is called multiple times, the overridden quota will be equal to the quotaSize provided in the final call. If this is called without specifying a quotaSize, the quota will be reset to the default value for the specified origin. If this is called multiple times with different origins, the override will be maintained for each origin until it is disabled (called without a quotaSize).
	QuotaSize float64 `json:"quotaSize,omitempty"`
}

// OverrideQuotaForOriginWithParams - Override quota for the specified origin
func (c *Storage) OverrideQuotaForOriginWithParams(ctx context.Context, v *StorageOverrideQuotaForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.overrideQuotaForOrigin", Params: v})
}

// OverrideQuotaForOrigin - Override quota for the specified origin
// origin - Security origin.
// quotaSize - The quota size (in bytes) to override the original quota with. If this is called multiple times, the overridden quota will be equal to the quotaSize provided in the final call. If this is called without specifying a quotaSize, the quota will be reset to the default value for the specified origin. If this is called multiple times with different origins, the override will be maintained for each origin until it is disabled (called without a quotaSize).
func (c *Storage) OverrideQuotaForOrigin(ctx context.Context, origin string, quotaSize float64) (*gcdmessage.ChromeResponse, error) {
	var v StorageOverrideQuotaForOriginParams
	v.Origin = origin
	v.QuotaSize = quotaSize
	return c.OverrideQuotaForOriginWithParams(ctx, &v)
}

type StorageTrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// TrackCacheStorageForOriginWithParams - Registers origin to be notified when an update occurs to its cache storage list.
func (c *Storage) TrackCacheStorageForOriginWithParams(ctx context.Context, v *StorageTrackCacheStorageForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackCacheStorageForOrigin", Params: v})
}

// TrackCacheStorageForOrigin - Registers origin to be notified when an update occurs to its cache storage list.
// origin - Security origin.
func (c *Storage) TrackCacheStorageForOrigin(ctx context.Context, origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackCacheStorageForOriginParams
	v.Origin = origin
	return c.TrackCacheStorageForOriginWithParams(ctx, &v)
}

type StorageTrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// TrackIndexedDBForOriginWithParams - Registers origin to be notified when an update occurs to its IndexedDB.
func (c *Storage) TrackIndexedDBForOriginWithParams(ctx context.Context, v *StorageTrackIndexedDBForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackIndexedDBForOrigin", Params: v})
}

// TrackIndexedDBForOrigin - Registers origin to be notified when an update occurs to its IndexedDB.
// origin - Security origin.
func (c *Storage) TrackIndexedDBForOrigin(ctx context.Context, origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackIndexedDBForOriginParams
	v.Origin = origin
	return c.TrackIndexedDBForOriginWithParams(ctx, &v)
}

type StorageTrackIndexedDBForStorageKeyParams struct {
	// Storage key.
	StorageKey string `json:"storageKey"`
}

// TrackIndexedDBForStorageKeyWithParams - Registers storage key to be notified when an update occurs to its IndexedDB.
func (c *Storage) TrackIndexedDBForStorageKeyWithParams(ctx context.Context, v *StorageTrackIndexedDBForStorageKeyParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackIndexedDBForStorageKey", Params: v})
}

// TrackIndexedDBForStorageKey - Registers storage key to be notified when an update occurs to its IndexedDB.
// storageKey - Storage key.
func (c *Storage) TrackIndexedDBForStorageKey(ctx context.Context, storageKey string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackIndexedDBForStorageKeyParams
	v.StorageKey = storageKey
	return c.TrackIndexedDBForStorageKeyWithParams(ctx, &v)
}

type StorageUntrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// UntrackCacheStorageForOriginWithParams - Unregisters origin from receiving notifications for cache storage.
func (c *Storage) UntrackCacheStorageForOriginWithParams(ctx context.Context, v *StorageUntrackCacheStorageForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackCacheStorageForOrigin", Params: v})
}

// UntrackCacheStorageForOrigin - Unregisters origin from receiving notifications for cache storage.
// origin - Security origin.
func (c *Storage) UntrackCacheStorageForOrigin(ctx context.Context, origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackCacheStorageForOriginParams
	v.Origin = origin
	return c.UntrackCacheStorageForOriginWithParams(ctx, &v)
}

type StorageUntrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// UntrackIndexedDBForOriginWithParams - Unregisters origin from receiving notifications for IndexedDB.
func (c *Storage) UntrackIndexedDBForOriginWithParams(ctx context.Context, v *StorageUntrackIndexedDBForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackIndexedDBForOrigin", Params: v})
}

// UntrackIndexedDBForOrigin - Unregisters origin from receiving notifications for IndexedDB.
// origin - Security origin.
func (c *Storage) UntrackIndexedDBForOrigin(ctx context.Context, origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackIndexedDBForOriginParams
	v.Origin = origin
	return c.UntrackIndexedDBForOriginWithParams(ctx, &v)
}

type StorageUntrackIndexedDBForStorageKeyParams struct {
	// Storage key.
	StorageKey string `json:"storageKey"`
}

// UntrackIndexedDBForStorageKeyWithParams - Unregisters storage key from receiving notifications for IndexedDB.
func (c *Storage) UntrackIndexedDBForStorageKeyWithParams(ctx context.Context, v *StorageUntrackIndexedDBForStorageKeyParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackIndexedDBForStorageKey", Params: v})
}

// UntrackIndexedDBForStorageKey - Unregisters storage key from receiving notifications for IndexedDB.
// storageKey - Storage key.
func (c *Storage) UntrackIndexedDBForStorageKey(ctx context.Context, storageKey string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackIndexedDBForStorageKeyParams
	v.StorageKey = storageKey
	return c.UntrackIndexedDBForStorageKeyWithParams(ctx, &v)
}

// GetTrustTokens - Returns the number of stored Trust Tokens per issuer for the current browsing context.
// Returns -  tokens -
func (c *Storage) GetTrustTokens(ctx context.Context) ([]*StorageTrustTokens, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getTrustTokens"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Tokens []*StorageTrustTokens
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

	return chromeData.Result.Tokens, nil
}

type StorageClearTrustTokensParams struct {
	//
	IssuerOrigin string `json:"issuerOrigin"`
}

// ClearTrustTokensWithParams - Removes all Trust Tokens issued by the provided issuerOrigin. Leaves other stored data, including the issuer's Redemption Records, intact.
// Returns -  didDeleteTokens - True if any tokens were deleted, false otherwise.
func (c *Storage) ClearTrustTokensWithParams(ctx context.Context, v *StorageClearTrustTokensParams) (bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearTrustTokens", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			DidDeleteTokens bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.DidDeleteTokens, nil
}

// ClearTrustTokens - Removes all Trust Tokens issued by the provided issuerOrigin. Leaves other stored data, including the issuer's Redemption Records, intact.
// issuerOrigin -
// Returns -  didDeleteTokens - True if any tokens were deleted, false otherwise.
func (c *Storage) ClearTrustTokens(ctx context.Context, issuerOrigin string) (bool, error) {
	var v StorageClearTrustTokensParams
	v.IssuerOrigin = issuerOrigin
	return c.ClearTrustTokensWithParams(ctx, &v)
}

type StorageGetInterestGroupDetailsParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
	//
	Name string `json:"name"`
}

// GetInterestGroupDetailsWithParams - Gets details for a named interest group.
// Returns -  details -
func (c *Storage) GetInterestGroupDetailsWithParams(ctx context.Context, v *StorageGetInterestGroupDetailsParams) (*StorageInterestGroupDetails, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getInterestGroupDetails", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Details *StorageInterestGroupDetails
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

	return chromeData.Result.Details, nil
}

// GetInterestGroupDetails - Gets details for a named interest group.
// ownerOrigin -
// name -
// Returns -  details -
func (c *Storage) GetInterestGroupDetails(ctx context.Context, ownerOrigin string, name string) (*StorageInterestGroupDetails, error) {
	var v StorageGetInterestGroupDetailsParams
	v.OwnerOrigin = ownerOrigin
	v.Name = name
	return c.GetInterestGroupDetailsWithParams(ctx, &v)
}

type StorageSetInterestGroupTrackingParams struct {
	//
	Enable bool `json:"enable"`
}

// SetInterestGroupTrackingWithParams - Enables/Disables issuing of interestGroupAccessed events.
func (c *Storage) SetInterestGroupTrackingWithParams(ctx context.Context, v *StorageSetInterestGroupTrackingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setInterestGroupTracking", Params: v})
}

// SetInterestGroupTracking - Enables/Disables issuing of interestGroupAccessed events.
// enable -
func (c *Storage) SetInterestGroupTracking(ctx context.Context, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetInterestGroupTrackingParams
	v.Enable = enable
	return c.SetInterestGroupTrackingWithParams(ctx, &v)
}
