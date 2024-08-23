// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Preload functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Corresponds to SpeculationRuleSet
type PreloadRuleSet struct {
	Id            string `json:"id"`                      //
	LoaderId      string `json:"loaderId"`                // Identifies a document which the rule set is associated with.
	SourceText    string `json:"sourceText"`              // Source text of JSON representing the rule set. If it comes from `<script>` tag, it is the textContent of the node. Note that it is a JSON for valid case.  See also: - https://wicg.github.io/nav-speculation/speculation-rules.html - https://github.com/WICG/nav-speculation/blob/main/triggers.md
	BackendNodeId int    `json:"backendNodeId,omitempty"` // A speculation rule set is either added through an inline `<script>` tag or through an external resource via the 'Speculation-Rules' HTTP header. For the first case, we include the BackendNodeId of the relevant `<script>` tag. For the second case, we include the external URL where the rule set was loaded from, and also RequestId if Network domain is enabled.  See also: - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-script - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-header
	Url           string `json:"url,omitempty"`           //
	RequestId     string `json:"requestId,omitempty"`     //
	ErrorType     string `json:"errorType,omitempty"`     // Error information `errorMessage` is null iff `errorType` is null. enum values: SourceIsNotJsonObject, InvalidRulesSkipped
	ErrorMessage  string `json:"errorMessage,omitempty"`  // TODO(https://crbug.com/1425354): Replace this property with structured error.
}

// A key that identifies a preloading attempt.  The url used is the url specified by the trigger (i.e. the initial URL), and not the final url that is navigated to. For example, prerendering allows same-origin main frame navigations during the attempt, but the attempt is still keyed with the initial URL.
type PreloadPreloadingAttemptKey struct {
	LoaderId   string `json:"loaderId"`             //
	Action     string `json:"action"`               //  enum values: Prefetch, Prerender
	Url        string `json:"url"`                  //
	TargetHint string `json:"targetHint,omitempty"` //  enum values: Blank, Self
}

// Lists sources for a preloading attempt, specifically the ids of rule sets that had a speculation rule that triggered the attempt, and the BackendNodeIds of <a href> or <area href> elements that triggered the attempt (in the case of attempts triggered by a document rule). It is possible for multiple rule sets and links to trigger a single attempt.
type PreloadPreloadingAttemptSource struct {
	Key        *PreloadPreloadingAttemptKey `json:"key"`        //
	RuleSetIds []string                     `json:"ruleSetIds"` //
	NodeIds    []int                        `json:"nodeIds"`    //
}

// Information of headers to be displayed when the header mismatch occurred.
type PreloadPrerenderMismatchedHeaders struct {
	HeaderName      string `json:"headerName"`                //
	InitialValue    string `json:"initialValue,omitempty"`    //
	ActivationValue string `json:"activationValue,omitempty"` //
}

// Upsert. Currently, it is only emitted when a rule set added.
type PreloadRuleSetUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RuleSet *PreloadRuleSet `json:"ruleSet"` //
	} `json:"Params,omitempty"`
}

type PreloadRuleSetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` //
	} `json:"Params,omitempty"`
}

// Fired when a preload enabled state is updated.
type PreloadPreloadEnabledStateUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		DisabledByPreference                        bool `json:"disabledByPreference"`                        //
		DisabledByDataSaver                         bool `json:"disabledByDataSaver"`                         //
		DisabledByBatterySaver                      bool `json:"disabledByBatterySaver"`                      //
		DisabledByHoldbackPrefetchSpeculationRules  bool `json:"disabledByHoldbackPrefetchSpeculationRules"`  //
		DisabledByHoldbackPrerenderSpeculationRules bool `json:"disabledByHoldbackPrerenderSpeculationRules"` //
	} `json:"Params,omitempty"`
}

// Fired when a prefetch attempt is updated.
type PreloadPrefetchStatusUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Key               *PreloadPreloadingAttemptKey `json:"key"`               //
		InitiatingFrameId string                       `json:"initiatingFrameId"` // The frame id of the frame initiating prefetch.
		PrefetchUrl       string                       `json:"prefetchUrl"`       //
		Status            string                       `json:"status"`            //  enum values: Pending, Running, Ready, Success, Failure, NotSupported
		PrefetchStatus    string                       `json:"prefetchStatus"`    //  enum values: PrefetchAllowed, PrefetchFailedIneligibleRedirect, PrefetchFailedInvalidRedirect, PrefetchFailedMIMENotSupported, PrefetchFailedNetError, PrefetchFailedNon2XX, PrefetchFailedPerPageLimitExceeded, PrefetchEvictedAfterCandidateRemoved, PrefetchEvictedForNewerPrefetch, PrefetchHeldback, PrefetchIneligibleRetryAfter, PrefetchIsPrivacyDecoy, PrefetchIsStale, PrefetchNotEligibleBrowserContextOffTheRecord, PrefetchNotEligibleDataSaverEnabled, PrefetchNotEligibleExistingProxy, PrefetchNotEligibleHostIsNonUnique, PrefetchNotEligibleNonDefaultStoragePartition, PrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy, PrefetchNotEligibleSchemeIsNotHttps, PrefetchNotEligibleUserHasCookies, PrefetchNotEligibleUserHasServiceWorker, PrefetchNotEligibleBatterySaverEnabled, PrefetchNotEligiblePreloadingDisabled, PrefetchNotFinishedInTime, PrefetchNotStarted, PrefetchNotUsedCookiesChanged, PrefetchProxyNotAvailable, PrefetchResponseUsed, PrefetchSuccessfulButNotUsed, PrefetchNotUsedProbeFailed
		RequestId         string                       `json:"requestId"`         //
	} `json:"Params,omitempty"`
}

// Fired when a prerender attempt is updated.
type PreloadPrerenderStatusUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Key                     *PreloadPreloadingAttemptKey         `json:"key"`                               //
		Status                  string                               `json:"status"`                            //  enum values: Pending, Running, Ready, Success, Failure, NotSupported
		PrerenderStatus         string                               `json:"prerenderStatus,omitempty"`         //  enum values: Activated, Destroyed, LowEndDevice, InvalidSchemeRedirect, InvalidSchemeNavigation, NavigationRequestBlockedByCsp, MainFrameNavigation, MojoBinderPolicy, RendererProcessCrashed, RendererProcessKilled, Download, TriggerDestroyed, NavigationNotCommitted, NavigationBadHttpStatus, ClientCertRequested, NavigationRequestNetworkError, CancelAllHostsForTesting, DidFailLoad, Stop, SslCertificateError, LoginAuthRequested, UaChangeRequiresReload, BlockedByClient, AudioOutputDeviceRequested, MixedContent, TriggerBackgrounded, MemoryLimitExceeded, DataSaverEnabled, TriggerUrlHasEffectiveUrl, ActivatedBeforeStarted, InactivePageRestriction, StartFailed, TimeoutBackgrounded, CrossSiteRedirectInInitialNavigation, CrossSiteNavigationInInitialNavigation, SameSiteCrossOriginRedirectNotOptInInInitialNavigation, SameSiteCrossOriginNavigationNotOptInInInitialNavigation, ActivationNavigationParameterMismatch, ActivatedInBackground, EmbedderHostDisallowed, ActivationNavigationDestroyedBeforeSuccess, TabClosedByUserGesture, TabClosedWithoutUserGesture, PrimaryMainFrameRendererProcessCrashed, PrimaryMainFrameRendererProcessKilled, ActivationFramePolicyNotCompatible, PreloadingDisabled, BatterySaverEnabled, ActivatedDuringMainFrameNavigation, PreloadingUnsupportedByWebContents, CrossSiteRedirectInMainFrameNavigation, CrossSiteNavigationInMainFrameNavigation, SameSiteCrossOriginRedirectNotOptInInMainFrameNavigation, SameSiteCrossOriginNavigationNotOptInInMainFrameNavigation, MemoryPressureOnTrigger, MemoryPressureAfterTriggered, PrerenderingDisabledByDevTools, SpeculationRuleRemoved, ActivatedWithAuxiliaryBrowsingContexts, MaxNumOfRunningEagerPrerendersExceeded, MaxNumOfRunningNonEagerPrerendersExceeded, MaxNumOfRunningEmbedderPrerendersExceeded, PrerenderingUrlHasEffectiveUrl, RedirectedPrerenderingUrlHasEffectiveUrl, ActivationUrlHasEffectiveUrl, JavaScriptInterfaceAdded, JavaScriptInterfaceRemoved, AllPrerenderingCanceled, WindowClosed, SlowNetwork, OtherPrerenderedPageActivated
		DisallowedMojoInterface string                               `json:"disallowedMojoInterface,omitempty"` // This is used to give users more information about the name of Mojo interface that is incompatible with prerender and has caused the cancellation of the attempt.
		MismatchedHeaders       []*PreloadPrerenderMismatchedHeaders `json:"mismatchedHeaders,omitempty"`       //
	} `json:"Params,omitempty"`
}

// Send a list of sources for all preloading attempts in a document.
type PreloadPreloadingAttemptSourcesUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		LoaderId                 string                            `json:"loaderId"`                 //
		PreloadingAttemptSources []*PreloadPreloadingAttemptSource `json:"preloadingAttemptSources"` //
	} `json:"Params,omitempty"`
}

type Preload struct {
	target gcdmessage.ChromeTargeter
}

func NewPreload(target gcdmessage.ChromeTargeter) *Preload {
	c := &Preload{target: target}
	return c
}

func (c *Preload) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Preload.enable"})
}

func (c *Preload) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Preload.disable"})
}
