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
	StorageType string  `json:"storageType"` // Name of storage type. enum values: appcache, cookies, file_systems, indexeddb, local_storage, shader_cache, websql, service_workers, cache_storage, interest_groups, shared_storage, storage_buckets, all, other
	Usage       float64 `json:"usage"`       // Storage usage (bytes).
}

// Pair of issuer origin and number of available (signed, but not used) Trust Tokens from that issuer.
type StorageTrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"` //
	Count        float64 `json:"count"`        //
}

// Struct for a single key-value pair in an origin's shared storage.
type StorageSharedStorageEntry struct {
	Key   string `json:"key"`   //
	Value string `json:"value"` //
}

// Details for an origin's shared storage.
type StorageSharedStorageMetadata struct {
	CreationTime    float64 `json:"creationTime"`    // Time when the origin's shared storage was last created.
	Length          int     `json:"length"`          // Number of key-value pairs stored in origin's shared storage.
	RemainingBudget float64 `json:"remainingBudget"` // Current amount of bits of entropy remaining in the navigation budget.
	BytesUsed       int     `json:"bytesUsed"`       // Total number of bytes stored as key-value pairs in origin's shared storage.
}

// Pair of reporting metadata details for a candidate URL for `selectURL()`.
type StorageSharedStorageReportingMetadata struct {
	EventType    string `json:"eventType"`    //
	ReportingUrl string `json:"reportingUrl"` //
}

// Bundles a candidate URL with its reporting metadata.
type StorageSharedStorageUrlWithMetadata struct {
	Url               string                                   `json:"url"`               // Spec of candidate URL.
	ReportingMetadata []*StorageSharedStorageReportingMetadata `json:"reportingMetadata"` // Any associated reporting metadata.
}

// Bundles the parameters for shared storage access events whose presence/absence can vary according to SharedStorageAccessType.
type StorageSharedStorageAccessParams struct {
	ScriptSourceUrl  string                                 `json:"scriptSourceUrl,omitempty"`  // Spec of the module script URL. Present only for SharedStorageAccessType.documentAddModule.
	OperationName    string                                 `json:"operationName,omitempty"`    // Name of the registered operation to be run. Present only for SharedStorageAccessType.documentRun and SharedStorageAccessType.documentSelectURL.
	SerializedData   string                                 `json:"serializedData,omitempty"`   // The operation's serialized data in bytes (converted to a string). Present only for SharedStorageAccessType.documentRun and SharedStorageAccessType.documentSelectURL.
	UrlsWithMetadata []*StorageSharedStorageUrlWithMetadata `json:"urlsWithMetadata,omitempty"` // Array of candidate URLs' specs, along with any associated metadata. Present only for SharedStorageAccessType.documentSelectURL.
	Key              string                                 `json:"key,omitempty"`              // Key for a specific entry in an origin's shared storage. Present only for SharedStorageAccessType.documentSet, SharedStorageAccessType.documentAppend, SharedStorageAccessType.documentDelete, SharedStorageAccessType.workletSet, SharedStorageAccessType.workletAppend, SharedStorageAccessType.workletDelete, SharedStorageAccessType.workletGet, SharedStorageAccessType.headerSet, SharedStorageAccessType.headerAppend, and SharedStorageAccessType.headerDelete.
	Value            string                                 `json:"value,omitempty"`            // Value for a specific entry in an origin's shared storage. Present only for SharedStorageAccessType.documentSet, SharedStorageAccessType.documentAppend, SharedStorageAccessType.workletSet, SharedStorageAccessType.workletAppend, SharedStorageAccessType.headerSet, and SharedStorageAccessType.headerAppend.
	IgnoreIfPresent  bool                                   `json:"ignoreIfPresent,omitempty"`  // Whether or not to set an entry for a key if that key is already present. Present only for SharedStorageAccessType.documentSet, SharedStorageAccessType.workletSet, and SharedStorageAccessType.headerSet.
}

// No Description.
type StorageStorageBucket struct {
	StorageKey string `json:"storageKey"`     //
	Name       string `json:"name,omitempty"` // If not specified, it is the default bucket of the storageKey.
}

// No Description.
type StorageStorageBucketInfo struct {
	Bucket     *StorageStorageBucket `json:"bucket"`     //
	Id         string                `json:"id"`         //
	Expiration float64               `json:"expiration"` //
	Quota      float64               `json:"quota"`      // Storage quota (bytes).
	Persistent bool                  `json:"persistent"` //
	Durability string                `json:"durability"` //  enum values: relaxed, strict
}

// No Description.
type StorageAttributionReportingFilterDataEntry struct {
	Key    string   `json:"key"`    //
	Values []string `json:"values"` //
}

// No Description.
type StorageAttributionReportingFilterConfig struct {
	FilterValues   []*StorageAttributionReportingFilterDataEntry `json:"filterValues"`             //
	LookbackWindow int                                           `json:"lookbackWindow,omitempty"` // duration in seconds
}

// No Description.
type StorageAttributionReportingFilterPair struct {
	Filters    []*StorageAttributionReportingFilterConfig `json:"filters"`    //
	NotFilters []*StorageAttributionReportingFilterConfig `json:"notFilters"` //
}

// No Description.
type StorageAttributionReportingAggregationKeysEntry struct {
	Key   string `json:"key"`   //
	Value string `json:"value"` //
}

// No Description.
type StorageAttributionReportingEventReportWindows struct {
	Start int   `json:"start"` // duration in seconds
	Ends  []int `json:"ends"`  // duration in seconds
}

// No Description.
type StorageAttributionReportingTriggerSpec struct {
	TriggerData        []float64                                      `json:"triggerData"`        // number instead of integer because not all uint32 can be represented by int
	EventReportWindows *StorageAttributionReportingEventReportWindows `json:"eventReportWindows"` //
}

// No Description.
type StorageAttributionReportingAggregatableDebugReportingData struct {
	KeyPiece string   `json:"keyPiece"` //
	Value    float64  `json:"value"`    // number instead of integer because not all uint32 can be represented by int
	Types    []string `json:"types"`    //
}

// No Description.
type StorageAttributionReportingAggregatableDebugReportingConfig struct {
	Budget                       float64                                                      `json:"budget,omitempty"`                       // number instead of integer because not all uint32 can be represented by int, only present for source registrations
	KeyPiece                     string                                                       `json:"keyPiece"`                               //
	DebugData                    []*StorageAttributionReportingAggregatableDebugReportingData `json:"debugData"`                              //
	AggregationCoordinatorOrigin string                                                       `json:"aggregationCoordinatorOrigin,omitempty"` //
}

// No Description.
type StorageAttributionReportingSourceRegistration struct {
	Time                             float64                                                      `json:"time"`                             //
	Expiry                           int                                                          `json:"expiry"`                           // duration in seconds
	TriggerSpecs                     []*StorageAttributionReportingTriggerSpec                    `json:"triggerSpecs"`                     //
	AggregatableReportWindow         int                                                          `json:"aggregatableReportWindow"`         // duration in seconds
	Type                             string                                                       `json:"type"`                             //  enum values: navigation, event
	SourceOrigin                     string                                                       `json:"sourceOrigin"`                     //
	ReportingOrigin                  string                                                       `json:"reportingOrigin"`                  //
	DestinationSites                 []string                                                     `json:"destinationSites"`                 //
	EventId                          string                                                       `json:"eventId"`                          //
	Priority                         string                                                       `json:"priority"`                         //
	FilterData                       []*StorageAttributionReportingFilterDataEntry                `json:"filterData"`                       //
	AggregationKeys                  []*StorageAttributionReportingAggregationKeysEntry           `json:"aggregationKeys"`                  //
	DebugKey                         string                                                       `json:"debugKey,omitempty"`               //
	TriggerDataMatching              string                                                       `json:"triggerDataMatching"`              //  enum values: exact, modulus
	DestinationLimitPriority         string                                                       `json:"destinationLimitPriority"`         //
	AggregatableDebugReportingConfig *StorageAttributionReportingAggregatableDebugReportingConfig `json:"aggregatableDebugReportingConfig"` //
}

// No Description.
type StorageAttributionReportingAggregatableValueDictEntry struct {
	Key         string  `json:"key"`         //
	Value       float64 `json:"value"`       // number instead of integer because not all uint32 can be represented by int
	FilteringId string  `json:"filteringId"` //
}

// No Description.
type StorageAttributionReportingAggregatableValueEntry struct {
	Values  []*StorageAttributionReportingAggregatableValueDictEntry `json:"values"`  //
	Filters *StorageAttributionReportingFilterPair                   `json:"filters"` //
}

// No Description.
type StorageAttributionReportingEventTriggerData struct {
	Data     string                                 `json:"data"`               //
	Priority string                                 `json:"priority"`           //
	DedupKey string                                 `json:"dedupKey,omitempty"` //
	Filters  *StorageAttributionReportingFilterPair `json:"filters"`            //
}

// No Description.
type StorageAttributionReportingAggregatableTriggerData struct {
	KeyPiece   string                                 `json:"keyPiece"`   //
	SourceKeys []string                               `json:"sourceKeys"` //
	Filters    *StorageAttributionReportingFilterPair `json:"filters"`    //
}

// No Description.
type StorageAttributionReportingAggregatableDedupKey struct {
	DedupKey string                                 `json:"dedupKey,omitempty"` //
	Filters  *StorageAttributionReportingFilterPair `json:"filters"`            //
}

// No Description.
type StorageAttributionReportingTriggerRegistration struct {
	Filters                          *StorageAttributionReportingFilterPair                       `json:"filters"`                                //
	DebugKey                         string                                                       `json:"debugKey,omitempty"`                     //
	AggregatableDedupKeys            []*StorageAttributionReportingAggregatableDedupKey           `json:"aggregatableDedupKeys"`                  //
	EventTriggerData                 []*StorageAttributionReportingEventTriggerData               `json:"eventTriggerData"`                       //
	AggregatableTriggerData          []*StorageAttributionReportingAggregatableTriggerData        `json:"aggregatableTriggerData"`                //
	AggregatableValues               []*StorageAttributionReportingAggregatableValueEntry         `json:"aggregatableValues"`                     //
	AggregatableFilteringIdMaxBytes  int                                                          `json:"aggregatableFilteringIdMaxBytes"`        //
	DebugReporting                   bool                                                         `json:"debugReporting"`                         //
	AggregationCoordinatorOrigin     string                                                       `json:"aggregationCoordinatorOrigin,omitempty"` //
	SourceRegistrationTimeConfig     string                                                       `json:"sourceRegistrationTimeConfig"`           //  enum values: include, exclude
	TriggerContextId                 string                                                       `json:"triggerContextId,omitempty"`             //
	AggregatableDebugReportingConfig *StorageAttributionReportingAggregatableDebugReportingConfig `json:"aggregatableDebugReportingConfig"`       //
}

// A single Related Website Set object.
type StorageRelatedWebsiteSet struct {
	PrimarySites    []string `json:"primarySites"`    // The primary site of this set, along with the ccTLDs if there is any.
	AssociatedSites []string `json:"associatedSites"` // The associated sites of this set, along with the ccTLDs if there is any.
	ServiceSites    []string `json:"serviceSites"`    // The service sites of this set, along with the ccTLDs if there is any.
}

// A cache's contents have been modified.
type StorageCacheStorageContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin     string `json:"origin"`     // Origin to update.
		StorageKey string `json:"storageKey"` // Storage key to update.
		BucketId   string `json:"bucketId"`   // Storage bucket to update.
		CacheName  string `json:"cacheName"`  // Name of cache in origin.
	} `json:"Params,omitempty"`
}

// A cache has been added/deleted.
type StorageCacheStorageListUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin     string `json:"origin"`     // Origin to update.
		StorageKey string `json:"storageKey"` // Storage key to update.
		BucketId   string `json:"bucketId"`   // Storage bucket to update.
	} `json:"Params,omitempty"`
}

// The origin's IndexedDB object store has been modified.
type StorageIndexedDBContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin          string `json:"origin"`          // Origin to update.
		StorageKey      string `json:"storageKey"`      // Storage key to update.
		BucketId        string `json:"bucketId"`        // Storage bucket to update.
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
		BucketId   string `json:"bucketId"`   // Storage bucket to update.
	} `json:"Params,omitempty"`
}

// One of the interest groups was accessed. Note that these events are global to all targets sharing an interest group store.
type StorageInterestGroupAccessedEvent struct {
	Method string `json:"method"`
	Params struct {
		AccessTime            float64 `json:"accessTime"`                      //
		Type                  string  `json:"type"`                            //  enum values: join, leave, update, loaded, bid, win, additionalBid, additionalBidWin, topLevelBid, topLevelAdditionalBid, clear
		OwnerOrigin           string  `json:"ownerOrigin"`                     //
		Name                  string  `json:"name"`                            //
		ComponentSellerOrigin string  `json:"componentSellerOrigin,omitempty"` // For topLevelBid/topLevelAdditionalBid, and when appropriate, win and additionalBidWin
		Bid                   float64 `json:"bid,omitempty"`                   // For bid or somethingBid event, if done locally and not on a server.
		BidCurrency           string  `json:"bidCurrency,omitempty"`           //
		UniqueAuctionId       string  `json:"uniqueAuctionId,omitempty"`       // For non-global events --- links to interestGroupAuctionEvent
	} `json:"Params,omitempty"`
}

// An auction involving interest groups is taking place. These events are target-specific.
type StorageInterestGroupAuctionEventOccurredEvent struct {
	Method string `json:"method"`
	Params struct {
		EventTime       float64                `json:"eventTime"`                 //
		Type            string                 `json:"type"`                      //  enum values: started, configResolved
		UniqueAuctionId string                 `json:"uniqueAuctionId"`           //
		ParentAuctionId string                 `json:"parentAuctionId,omitempty"` // Set for child auctions.
		AuctionConfig   map[string]interface{} `json:"auctionConfig,omitempty"`   // Set for started and configResolved
	} `json:"Params,omitempty"`
}

// Specifies which auctions a particular network fetch may be related to, and in what role. Note that it is not ordered with respect to Network.requestWillBeSent (but will happen before loadingFinished loadingFailed).
type StorageInterestGroupAuctionNetworkRequestCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Type      string   `json:"type"`      //  enum values: bidderJs, bidderWasm, sellerJs, bidderTrustedSignals, sellerTrustedSignals
		RequestId string   `json:"requestId"` //
		Auctions  []string `json:"auctions"`  // This is the set of the auctions using the worklet that issued this request.  In the case of trusted signals, it's possible that only some of them actually care about the keys being queried.
	} `json:"Params,omitempty"`
}

// Shared storage was accessed by the associated page. The following parameters are included in all events.
type StorageSharedStorageAccessedEvent struct {
	Method string `json:"method"`
	Params struct {
		AccessTime  float64                           `json:"accessTime"`  // Time of the access.
		Type        string                            `json:"type"`        // Enum value indicating the Shared Storage API method invoked. enum values: documentAddModule, documentSelectURL, documentRun, documentSet, documentAppend, documentDelete, documentClear, documentGet, workletSet, workletAppend, workletDelete, workletClear, workletGet, workletKeys, workletEntries, workletLength, workletRemainingBudget, headerSet, headerAppend, headerDelete, headerClear
		MainFrameId string                            `json:"mainFrameId"` // DevTools Frame Token for the primary frame tree's root.
		OwnerOrigin string                            `json:"ownerOrigin"` // Serialized origin for the context that invoked the Shared Storage API.
		Params      *StorageSharedStorageAccessParams `json:"params"`      // The sub-parameters wrapped by `params` are all optional and their presence/absence depends on `type`.
	} `json:"Params,omitempty"`
}

type StorageStorageBucketCreatedOrUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		BucketInfo *StorageStorageBucketInfo `json:"bucketInfo"` //
	} `json:"Params,omitempty"`
}

type StorageStorageBucketDeletedEvent struct {
	Method string `json:"method"`
	Params struct {
		BucketId string `json:"bucketId"` //
	} `json:"Params,omitempty"`
}

type StorageAttributionReportingSourceRegisteredEvent struct {
	Method string `json:"method"`
	Params struct {
		Registration *StorageAttributionReportingSourceRegistration `json:"registration"` //
		Result       string                                         `json:"result"`       //  enum values: success, internalError, insufficientSourceCapacity, insufficientUniqueDestinationCapacity, excessiveReportingOrigins, prohibitedByBrowserPolicy, successNoised, destinationReportingLimitReached, destinationGlobalLimitReached, destinationBothLimitsReached, reportingOriginsPerSiteLimitReached, exceedsMaxChannelCapacity, exceedsMaxTriggerStateCardinality, destinationPerDayReportingLimitReached
	} `json:"Params,omitempty"`
}

type StorageAttributionReportingTriggerRegisteredEvent struct {
	Method string `json:"method"`
	Params struct {
		Registration *StorageAttributionReportingTriggerRegistration `json:"registration"` //
		EventLevel   string                                          `json:"eventLevel"`   //  enum values: success, successDroppedLowerPriority, internalError, noCapacityForAttributionDestination, noMatchingSources, deduplicated, excessiveAttributions, priorityTooLow, neverAttributedSource, excessiveReportingOrigins, noMatchingSourceFilterData, prohibitedByBrowserPolicy, noMatchingConfigurations, excessiveReports, falselyAttributedSource, reportWindowPassed, notRegistered, reportWindowNotStarted, noMatchingTriggerData
		Aggregatable string                                          `json:"aggregatable"` //  enum values: success, internalError, noCapacityForAttributionDestination, noMatchingSources, excessiveAttributions, excessiveReportingOrigins, noHistograms, insufficientBudget, noMatchingSourceFilterData, notRegistered, prohibitedByBrowserPolicy, deduplicated, reportWindowPassed, excessiveReports
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
		gcdmessage.ChromeErrorResponse
		Result struct {
			StorageKey string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
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
		gcdmessage.ChromeErrorResponse
		Result struct {
			Cookies []*NetworkCookie
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
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
		gcdmessage.ChromeErrorResponse
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

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, false, nil, err
	}

	if chromeData.Error != nil {
		return 0, 0, false, nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
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

type StorageTrackCacheStorageForStorageKeyParams struct {
	// Storage key.
	StorageKey string `json:"storageKey"`
}

// TrackCacheStorageForStorageKeyWithParams - Registers storage key to be notified when an update occurs to its cache storage list.
func (c *Storage) TrackCacheStorageForStorageKeyWithParams(ctx context.Context, v *StorageTrackCacheStorageForStorageKeyParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackCacheStorageForStorageKey", Params: v})
}

// TrackCacheStorageForStorageKey - Registers storage key to be notified when an update occurs to its cache storage list.
// storageKey - Storage key.
func (c *Storage) TrackCacheStorageForStorageKey(ctx context.Context, storageKey string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackCacheStorageForStorageKeyParams
	v.StorageKey = storageKey
	return c.TrackCacheStorageForStorageKeyWithParams(ctx, &v)
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

type StorageUntrackCacheStorageForStorageKeyParams struct {
	// Storage key.
	StorageKey string `json:"storageKey"`
}

// UntrackCacheStorageForStorageKeyWithParams - Unregisters storage key from receiving notifications for cache storage.
func (c *Storage) UntrackCacheStorageForStorageKeyWithParams(ctx context.Context, v *StorageUntrackCacheStorageForStorageKeyParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackCacheStorageForStorageKey", Params: v})
}

// UntrackCacheStorageForStorageKey - Unregisters storage key from receiving notifications for cache storage.
// storageKey - Storage key.
func (c *Storage) UntrackCacheStorageForStorageKey(ctx context.Context, storageKey string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackCacheStorageForStorageKeyParams
	v.StorageKey = storageKey
	return c.UntrackCacheStorageForStorageKeyWithParams(ctx, &v)
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
		gcdmessage.ChromeErrorResponse
		Result struct {
			Tokens []*StorageTrustTokens
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
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
		gcdmessage.ChromeErrorResponse
		Result struct {
			DidDeleteTokens bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	if chromeData.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
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
// Returns -  details - This largely corresponds to: https://wicg.github.io/turtledove/#dictdef-generatebidinterestgroup but has absolute expirationTime instead of relative lifetimeMs and also adds joiningOrigin.
func (c *Storage) GetInterestGroupDetailsWithParams(ctx context.Context, v *StorageGetInterestGroupDetailsParams) (map[string]interface{}, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getInterestGroupDetails", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Details map[string]interface{}
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Details, nil
}

// GetInterestGroupDetails - Gets details for a named interest group.
// ownerOrigin -
// name -
// Returns -  details - This largely corresponds to: https://wicg.github.io/turtledove/#dictdef-generatebidinterestgroup but has absolute expirationTime instead of relative lifetimeMs and also adds joiningOrigin.
func (c *Storage) GetInterestGroupDetails(ctx context.Context, ownerOrigin string, name string) (map[string]interface{}, error) {
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

type StorageSetInterestGroupAuctionTrackingParams struct {
	//
	Enable bool `json:"enable"`
}

// SetInterestGroupAuctionTrackingWithParams - Enables/Disables issuing of interestGroupAuctionEventOccurred and interestGroupAuctionNetworkRequestCreated.
func (c *Storage) SetInterestGroupAuctionTrackingWithParams(ctx context.Context, v *StorageSetInterestGroupAuctionTrackingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setInterestGroupAuctionTracking", Params: v})
}

// SetInterestGroupAuctionTracking - Enables/Disables issuing of interestGroupAuctionEventOccurred and interestGroupAuctionNetworkRequestCreated.
// enable -
func (c *Storage) SetInterestGroupAuctionTracking(ctx context.Context, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetInterestGroupAuctionTrackingParams
	v.Enable = enable
	return c.SetInterestGroupAuctionTrackingWithParams(ctx, &v)
}

type StorageGetSharedStorageMetadataParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
}

// GetSharedStorageMetadataWithParams - Gets metadata for an origin's shared storage.
// Returns -  metadata -
func (c *Storage) GetSharedStorageMetadataWithParams(ctx context.Context, v *StorageGetSharedStorageMetadataParams) (*StorageSharedStorageMetadata, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getSharedStorageMetadata", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Metadata *StorageSharedStorageMetadata
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Metadata, nil
}

// GetSharedStorageMetadata - Gets metadata for an origin's shared storage.
// ownerOrigin -
// Returns -  metadata -
func (c *Storage) GetSharedStorageMetadata(ctx context.Context, ownerOrigin string) (*StorageSharedStorageMetadata, error) {
	var v StorageGetSharedStorageMetadataParams
	v.OwnerOrigin = ownerOrigin
	return c.GetSharedStorageMetadataWithParams(ctx, &v)
}

type StorageGetSharedStorageEntriesParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
}

// GetSharedStorageEntriesWithParams - Gets the entries in an given origin's shared storage.
// Returns -  entries -
func (c *Storage) GetSharedStorageEntriesWithParams(ctx context.Context, v *StorageGetSharedStorageEntriesParams) ([]*StorageSharedStorageEntry, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getSharedStorageEntries", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Entries []*StorageSharedStorageEntry
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Entries, nil
}

// GetSharedStorageEntries - Gets the entries in an given origin's shared storage.
// ownerOrigin -
// Returns -  entries -
func (c *Storage) GetSharedStorageEntries(ctx context.Context, ownerOrigin string) ([]*StorageSharedStorageEntry, error) {
	var v StorageGetSharedStorageEntriesParams
	v.OwnerOrigin = ownerOrigin
	return c.GetSharedStorageEntriesWithParams(ctx, &v)
}

type StorageSetSharedStorageEntryParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
	//
	Key string `json:"key"`
	//
	Value string `json:"value"`
	// If `ignoreIfPresent` is included and true, then only sets the entry if `key` doesn't already exist.
	IgnoreIfPresent bool `json:"ignoreIfPresent,omitempty"`
}

// SetSharedStorageEntryWithParams - Sets entry with `key` and `value` for a given origin's shared storage.
func (c *Storage) SetSharedStorageEntryWithParams(ctx context.Context, v *StorageSetSharedStorageEntryParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setSharedStorageEntry", Params: v})
}

// SetSharedStorageEntry - Sets entry with `key` and `value` for a given origin's shared storage.
// ownerOrigin -
// key -
// value -
// ignoreIfPresent - If `ignoreIfPresent` is included and true, then only sets the entry if `key` doesn't already exist.
func (c *Storage) SetSharedStorageEntry(ctx context.Context, ownerOrigin string, key string, value string, ignoreIfPresent bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetSharedStorageEntryParams
	v.OwnerOrigin = ownerOrigin
	v.Key = key
	v.Value = value
	v.IgnoreIfPresent = ignoreIfPresent
	return c.SetSharedStorageEntryWithParams(ctx, &v)
}

type StorageDeleteSharedStorageEntryParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
	//
	Key string `json:"key"`
}

// DeleteSharedStorageEntryWithParams - Deletes entry for `key` (if it exists) for a given origin's shared storage.
func (c *Storage) DeleteSharedStorageEntryWithParams(ctx context.Context, v *StorageDeleteSharedStorageEntryParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.deleteSharedStorageEntry", Params: v})
}

// DeleteSharedStorageEntry - Deletes entry for `key` (if it exists) for a given origin's shared storage.
// ownerOrigin -
// key -
func (c *Storage) DeleteSharedStorageEntry(ctx context.Context, ownerOrigin string, key string) (*gcdmessage.ChromeResponse, error) {
	var v StorageDeleteSharedStorageEntryParams
	v.OwnerOrigin = ownerOrigin
	v.Key = key
	return c.DeleteSharedStorageEntryWithParams(ctx, &v)
}

type StorageClearSharedStorageEntriesParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
}

// ClearSharedStorageEntriesWithParams - Clears all entries for a given origin's shared storage.
func (c *Storage) ClearSharedStorageEntriesWithParams(ctx context.Context, v *StorageClearSharedStorageEntriesParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearSharedStorageEntries", Params: v})
}

// ClearSharedStorageEntries - Clears all entries for a given origin's shared storage.
// ownerOrigin -
func (c *Storage) ClearSharedStorageEntries(ctx context.Context, ownerOrigin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageClearSharedStorageEntriesParams
	v.OwnerOrigin = ownerOrigin
	return c.ClearSharedStorageEntriesWithParams(ctx, &v)
}

type StorageResetSharedStorageBudgetParams struct {
	//
	OwnerOrigin string `json:"ownerOrigin"`
}

// ResetSharedStorageBudgetWithParams - Resets the budget for `ownerOrigin` by clearing all budget withdrawals.
func (c *Storage) ResetSharedStorageBudgetWithParams(ctx context.Context, v *StorageResetSharedStorageBudgetParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.resetSharedStorageBudget", Params: v})
}

// ResetSharedStorageBudget - Resets the budget for `ownerOrigin` by clearing all budget withdrawals.
// ownerOrigin -
func (c *Storage) ResetSharedStorageBudget(ctx context.Context, ownerOrigin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageResetSharedStorageBudgetParams
	v.OwnerOrigin = ownerOrigin
	return c.ResetSharedStorageBudgetWithParams(ctx, &v)
}

type StorageSetSharedStorageTrackingParams struct {
	//
	Enable bool `json:"enable"`
}

// SetSharedStorageTrackingWithParams - Enables/disables issuing of sharedStorageAccessed events.
func (c *Storage) SetSharedStorageTrackingWithParams(ctx context.Context, v *StorageSetSharedStorageTrackingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setSharedStorageTracking", Params: v})
}

// SetSharedStorageTracking - Enables/disables issuing of sharedStorageAccessed events.
// enable -
func (c *Storage) SetSharedStorageTracking(ctx context.Context, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetSharedStorageTrackingParams
	v.Enable = enable
	return c.SetSharedStorageTrackingWithParams(ctx, &v)
}

type StorageSetStorageBucketTrackingParams struct {
	//
	StorageKey string `json:"storageKey"`
	//
	Enable bool `json:"enable"`
}

// SetStorageBucketTrackingWithParams - Set tracking for a storage key's buckets.
func (c *Storage) SetStorageBucketTrackingWithParams(ctx context.Context, v *StorageSetStorageBucketTrackingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setStorageBucketTracking", Params: v})
}

// SetStorageBucketTracking - Set tracking for a storage key's buckets.
// storageKey -
// enable -
func (c *Storage) SetStorageBucketTracking(ctx context.Context, storageKey string, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetStorageBucketTrackingParams
	v.StorageKey = storageKey
	v.Enable = enable
	return c.SetStorageBucketTrackingWithParams(ctx, &v)
}

type StorageDeleteStorageBucketParams struct {
	//
	Bucket *StorageStorageBucket `json:"bucket"`
}

// DeleteStorageBucketWithParams - Deletes the Storage Bucket with the given storage key and bucket name.
func (c *Storage) DeleteStorageBucketWithParams(ctx context.Context, v *StorageDeleteStorageBucketParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.deleteStorageBucket", Params: v})
}

// DeleteStorageBucket - Deletes the Storage Bucket with the given storage key and bucket name.
// bucket -
func (c *Storage) DeleteStorageBucket(ctx context.Context, bucket *StorageStorageBucket) (*gcdmessage.ChromeResponse, error) {
	var v StorageDeleteStorageBucketParams
	v.Bucket = bucket
	return c.DeleteStorageBucketWithParams(ctx, &v)
}

// RunBounceTrackingMitigations - Deletes state for sites identified as potential bounce trackers, immediately.
// Returns -  deletedSites -
func (c *Storage) RunBounceTrackingMitigations(ctx context.Context) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.runBounceTrackingMitigations"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			DeletedSites []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.DeletedSites, nil
}

type StorageSetAttributionReportingLocalTestingModeParams struct {
	// If enabled, noise is suppressed and reports are sent immediately.
	Enabled bool `json:"enabled"`
}

// SetAttributionReportingLocalTestingModeWithParams - https://wicg.github.io/attribution-reporting-api/
func (c *Storage) SetAttributionReportingLocalTestingModeWithParams(ctx context.Context, v *StorageSetAttributionReportingLocalTestingModeParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setAttributionReportingLocalTestingMode", Params: v})
}

// SetAttributionReportingLocalTestingMode - https://wicg.github.io/attribution-reporting-api/
// enabled - If enabled, noise is suppressed and reports are sent immediately.
func (c *Storage) SetAttributionReportingLocalTestingMode(ctx context.Context, enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetAttributionReportingLocalTestingModeParams
	v.Enabled = enabled
	return c.SetAttributionReportingLocalTestingModeWithParams(ctx, &v)
}

type StorageSetAttributionReportingTrackingParams struct {
	//
	Enable bool `json:"enable"`
}

// SetAttributionReportingTrackingWithParams - Enables/disables issuing of Attribution Reporting events.
func (c *Storage) SetAttributionReportingTrackingWithParams(ctx context.Context, v *StorageSetAttributionReportingTrackingParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.setAttributionReportingTracking", Params: v})
}

// SetAttributionReportingTracking - Enables/disables issuing of Attribution Reporting events.
// enable -
func (c *Storage) SetAttributionReportingTracking(ctx context.Context, enable bool) (*gcdmessage.ChromeResponse, error) {
	var v StorageSetAttributionReportingTrackingParams
	v.Enable = enable
	return c.SetAttributionReportingTrackingWithParams(ctx, &v)
}

// SendPendingAttributionReports - Sends all pending Attribution Reports immediately, regardless of their scheduled report time.
// Returns -  numSent - The number of reports that were sent.
func (c *Storage) SendPendingAttributionReports(ctx context.Context) (int, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.sendPendingAttributionReports"})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			NumSent int
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	if chromeData.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.NumSent, nil
}

// GetRelatedWebsiteSets - Returns the effective Related Website Sets in use by this profile for the browser session. The effective Related Website Sets will not change during a browser session.
// Returns -  sets -
func (c *Storage) GetRelatedWebsiteSets(ctx context.Context) ([]*StorageRelatedWebsiteSet, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getRelatedWebsiteSets"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Sets []*StorageRelatedWebsiteSet
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.Sets, nil
}
