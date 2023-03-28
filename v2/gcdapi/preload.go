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
	Id         string `json:"id"`         //
	LoaderId   string `json:"loaderId"`   // Identifies a document which the rule set is associated with.
	SourceText string `json:"sourceText"` // Source text of JSON representing the rule set. If it comes from <script> tag, it is the textContent of the node. Note that it is a JSON for valid case.  See also: - https://wicg.github.io/nav-speculation/speculation-rules.html - https://github.com/WICG/nav-speculation/blob/main/triggers.md
}

// A key that identifies a preloading attempt.  The url used is the url specified by the trigger (i.e. the initial URL), and not the final url that is navigated to. For example, prerendering allows same-origin main frame navigations during the attempt, but the attempt is still keyed with the initial URL.
type PreloadPreloadingAttemptKey struct {
	LoaderId   string `json:"loaderId"`             //
	Action     string `json:"action"`               //  enum values: Prefetch, Prerender
	Url        string `json:"url"`                  //
	TargetHint string `json:"targetHint,omitempty"` //  enum values: Blank, Self
}

// Lists sources for a preloading attempt, specifically the ids of rule sets that had a speculation rule that triggered the attempt, and the BackendNodeIds of <a href> or <area href> elements that triggered the attempt (in the case of attempts triggered by a document rule). It is possible for mulitple rule sets and links to trigger a single attempt.
type PreloadPreloadingAttemptSource struct {
	Key        *PreloadPreloadingAttemptKey `json:"key"`        //
	RuleSetIds []string                     `json:"ruleSetIds"` //
	NodeIds    []int                        `json:"nodeIds"`    //
}

// Upsert. Currently, it is only emitted when a rule set added.
type PreloadRuleSetUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RuleSet *PreloadRuleSet `json:"ruleSet"` //
	} `json:"Params,omitempty"`
}

//
type PreloadRuleSetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id string `json:"id"` //
	} `json:"Params,omitempty"`
}

// Fired when a prerender attempt is completed.
type PreloadPrerenderAttemptCompletedEvent struct {
	Method string `json:"method"`
	Params struct {
		InitiatingFrameId   string `json:"initiatingFrameId"`             // The frame id of the frame initiating prerendering.
		PrerenderingUrl     string `json:"prerenderingUrl"`               //
		FinalStatus         string `json:"finalStatus"`                   //  enum values: Activated, Destroyed, LowEndDevice, InvalidSchemeRedirect, InvalidSchemeNavigation, InProgressNavigation, NavigationRequestBlockedByCsp, MainFrameNavigation, MojoBinderPolicy, RendererProcessCrashed, RendererProcessKilled, Download, TriggerDestroyed, NavigationNotCommitted, NavigationBadHttpStatus, ClientCertRequested, NavigationRequestNetworkError, MaxNumOfRunningPrerendersExceeded, CancelAllHostsForTesting, DidFailLoad, Stop, SslCertificateError, LoginAuthRequested, UaChangeRequiresReload, BlockedByClient, AudioOutputDeviceRequested, MixedContent, TriggerBackgrounded, EmbedderTriggeredAndCrossOriginRedirected, MemoryLimitExceeded, FailToGetMemoryUsage, DataSaverEnabled, HasEffectiveUrl, ActivatedBeforeStarted, InactivePageRestriction, StartFailed, TimeoutBackgrounded, CrossSiteRedirectInInitialNavigation, CrossSiteNavigationInInitialNavigation, SameSiteCrossOriginRedirectNotOptInInInitialNavigation, SameSiteCrossOriginNavigationNotOptInInInitialNavigation, ActivationNavigationParameterMismatch, ActivatedInBackground, EmbedderHostDisallowed, ActivationNavigationDestroyedBeforeSuccess, TabClosedByUserGesture, TabClosedWithoutUserGesture, PrimaryMainFrameRendererProcessCrashed, PrimaryMainFrameRendererProcessKilled, ActivationFramePolicyNotCompatible, PreloadingDisabled, BatterySaverEnabled, ActivatedDuringMainFrameNavigation, PreloadingUnsupportedByWebContents, CrossSiteRedirectInMainFrameNavigation, CrossSiteNavigationInMainFrameNavigation, SameSiteCrossOriginRedirectNotOptInInMainFrameNavigation, SameSiteCrossOriginNavigationNotOptInInMainFrameNavigation
		DisallowedApiMethod string `json:"disallowedApiMethod,omitempty"` // This is used to give users more information about the name of the API call that is incompatible with prerender and has caused the cancellation of the attempt
	} `json:"Params,omitempty"`
}

// Fired when a prefetch attempt is updated.
type PreloadPrefetchStatusUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		InitiatingFrameId string `json:"initiatingFrameId"` // The frame id of the frame initiating prefetch.
		PrefetchUrl       string `json:"prefetchUrl"`       //
		Status            string `json:"status"`            //  enum values: Pending, Running, Ready, Success, Failure, NotSupported
	} `json:"Params,omitempty"`
}

// Fired when a prerender attempt is updated.
type PreloadPrerenderStatusUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		InitiatingFrameId string `json:"initiatingFrameId"` // The frame id of the frame initiating prerender.
		PrerenderingUrl   string `json:"prerenderingUrl"`   //
		Status            string `json:"status"`            //  enum values: Pending, Running, Ready, Success, Failure, NotSupported
	} `json:"Params,omitempty"`
}

// Send a list of sources for all preloading attempts.
type PreloadPreloadingAttemptSourcesUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
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

//
func (c *Preload) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Preload.enable"})
}

//
func (c *Preload) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Preload.disable"})
}
