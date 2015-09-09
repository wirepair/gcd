// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the ServiceWorker types.
// API Version: 1.1
package types

// ServiceWorker registration.
type ChromeServiceWorkerServiceWorkerRegistration struct {
	RegistrationId string `json:"registrationId"`
	ScopeURL       string `json:"scopeURL"`
	IsDeleted      bool   `json:"isDeleted,omitempty"`
}

//
type ChromeServiceWorkerServiceWorkerVersionRunningStatus string // possible values: stopped, starting, running, stopping

//
type ChromeServiceWorkerServiceWorkerVersionStatus string // possible values: new, installing, installed, activating, activated, redundant

//
type ChromeServiceWorkerTargetID string

// ServiceWorker version.
type ChromeServiceWorkerServiceWorkerVersion struct {
	VersionId          string                                                `json:"versionId"`
	RegistrationId     string                                                `json:"registrationId"`
	ScriptURL          string                                                `json:"scriptURL"`
	RunningStatus      *ChromeServiceWorkerServiceWorkerVersionRunningStatus `json:"runningStatus"`
	Status             *ChromeServiceWorkerServiceWorkerVersionStatus        `json:"status"`
	ScriptLastModified float64                                               `json:"scriptLastModified,omitempty"` // The Last-Modified header value of the main script.
	ScriptResponseTime float64                                               `json:"scriptResponseTime,omitempty"` // The time at which the response headers of the main script were received from the server.  For cached script it is the last time the cache entry was validated.
	ControlledClients  []*ChromeServiceWorkerTargetID                        `json:"controlledClients,omitempty"`
}

// ServiceWorker error message.
type ChromeServiceWorkerServiceWorkerErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`
	RegistrationId string `json:"registrationId"`
	VersionId      string `json:"versionId"`
	SourceURL      string `json:"sourceURL"`
	LineNumber     int    `json:"lineNumber"`
	ColumnNumber   int    `json:"columnNumber"`
}

//
type ChromeServiceWorkerTargetInfo struct {
	Id    *ChromeServiceWorkerTargetID `json:"id"`
	Type  string                       `json:"type"`
	Title string                       `json:"title"`
	Url   string                       `json:"url"`
}
