// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Tracing functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// No Description.
type TracingTraceConfig struct {
	RecordMode           string                 `json:"recordMode,omitempty"`           // Controls how the trace buffer stores data.
	TraceBufferSizeInKb  float64                `json:"traceBufferSizeInKb,omitempty"`  // Size of the trace buffer in kilobytes. If not specified or zero is passed, a default value of 200 MB would be used.
	EnableSampling       bool                   `json:"enableSampling,omitempty"`       // Turns on JavaScript stack sampling.
	EnableSystrace       bool                   `json:"enableSystrace,omitempty"`       // Turns on system tracing.
	EnableArgumentFilter bool                   `json:"enableArgumentFilter,omitempty"` // Turns on argument filter.
	IncludedCategories   []string               `json:"includedCategories,omitempty"`   // Included category filters.
	ExcludedCategories   []string               `json:"excludedCategories,omitempty"`   // Excluded category filters.
	SyntheticDelays      []string               `json:"syntheticDelays,omitempty"`      // Configuration to synthesize the delays in tracing.
	MemoryDumpConfig     map[string]interface{} `json:"memoryDumpConfig,omitempty"`     // Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}

type TracingBufferUsageEvent struct {
	Method string `json:"method"`
	Params struct {
		PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
		EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
		Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	} `json:"Params,omitempty"`
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
type TracingTracingCompleteEvent struct {
	Method string `json:"method"`
	Params struct {
		DataLossOccurred  bool   `json:"dataLossOccurred"`            // Indicates whether some trace data is known to have been lost, e.g. because the trace ring buffer wrapped around.
		Stream            string `json:"stream,omitempty"`            // A handle of the stream that holds resulting trace data.
		TraceFormat       string `json:"traceFormat,omitempty"`       // Trace data format of returned stream. enum values: json, proto
		StreamCompression string `json:"streamCompression,omitempty"` // Compression format of returned stream. enum values: none, gzip
	} `json:"Params,omitempty"`
}

type Tracing struct {
	target gcdmessage.ChromeTargeter
}

func NewTracing(target gcdmessage.ChromeTargeter) *Tracing {
	c := &Tracing{target: target}
	return c
}

// Stop trace events collection.
func (c *Tracing) End(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.end"})
}

// GetCategories - Gets supported tracing categories.
// Returns -  categories - A list of supported tracing categories.
func (c *Tracing) GetCategories(ctx context.Context) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.getCategories"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			Categories []string
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

	return chromeData.Result.Categories, nil
}

type TracingRecordClockSyncMarkerParams struct {
	// The ID of this clock sync marker
	SyncId string `json:"syncId"`
}

// RecordClockSyncMarkerWithParams - Record a clock sync marker in the trace.
func (c *Tracing) RecordClockSyncMarkerWithParams(ctx context.Context, v *TracingRecordClockSyncMarkerParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.recordClockSyncMarker", Params: v})
}

// RecordClockSyncMarker - Record a clock sync marker in the trace.
// syncId - The ID of this clock sync marker
func (c *Tracing) RecordClockSyncMarker(ctx context.Context, syncId string) (*gcdmessage.ChromeResponse, error) {
	var v TracingRecordClockSyncMarkerParams
	v.SyncId = syncId
	return c.RecordClockSyncMarkerWithParams(ctx, &v)
}

type TracingRequestMemoryDumpParams struct {
	// Enables more deterministic results by forcing garbage collection
	Deterministic bool `json:"deterministic,omitempty"`
	// Specifies level of details in memory dump. Defaults to "detailed". enum values: background, light, detailed
	LevelOfDetail string `json:"levelOfDetail,omitempty"`
}

// RequestMemoryDumpWithParams - Request a global memory dump.
// Returns -  dumpGuid - GUID of the resulting global memory dump. success - True iff the global memory dump succeeded.
func (c *Tracing) RequestMemoryDumpWithParams(ctx context.Context, v *TracingRequestMemoryDumpParams) (string, bool, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.requestMemoryDump", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			DumpGuid string
			Success  bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	if chromeData.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.DumpGuid, chromeData.Result.Success, nil
}

// RequestMemoryDump - Request a global memory dump.
// deterministic - Enables more deterministic results by forcing garbage collection
// levelOfDetail - Specifies level of details in memory dump. Defaults to "detailed". enum values: background, light, detailed
// Returns -  dumpGuid - GUID of the resulting global memory dump. success - True iff the global memory dump succeeded.
func (c *Tracing) RequestMemoryDump(ctx context.Context, deterministic bool, levelOfDetail string) (string, bool, error) {
	var v TracingRequestMemoryDumpParams
	v.Deterministic = deterministic
	v.LevelOfDetail = levelOfDetail
	return c.RequestMemoryDumpWithParams(ctx, &v)
}

type TracingStartParams struct {
	// Category/tag filter
	Categories string `json:"categories,omitempty"`
	// Tracing options
	Options string `json:"options,omitempty"`
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`
	// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to `ReportEvents`).
	TransferMode string `json:"transferMode,omitempty"`
	// Trace data format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `json`). enum values: json, proto
	StreamFormat string `json:"streamFormat,omitempty"`
	// Compression format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `none`) enum values: none, gzip
	StreamCompression string `json:"streamCompression,omitempty"`
	//
	TraceConfig *TracingTraceConfig `json:"traceConfig,omitempty"`
	// Backend type (defaults to `auto`) enum values: auto, chrome, system
	TracingBackend string `json:"tracingBackend,omitempty"`
}

// StartWithParams - Start trace events collection.
func (c *Tracing) StartWithParams(ctx context.Context, v *TracingStartParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.start", Params: v})
}

// Start - Start trace events collection.
// categories - Category/tag filter
// options - Tracing options
// bufferUsageReportingInterval - If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
// transferMode - Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to `ReportEvents`).
// streamFormat - Trace data format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `json`). enum values: json, proto
// streamCompression - Compression format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `none`) enum values: none, gzip
// traceConfig -
// tracingBackend - Backend type (defaults to `auto`) enum values: auto, chrome, system
func (c *Tracing) Start(ctx context.Context, categories string, options string, bufferUsageReportingInterval float64, transferMode string, streamFormat string, streamCompression string, traceConfig *TracingTraceConfig, tracingBackend string) (*gcdmessage.ChromeResponse, error) {
	var v TracingStartParams
	v.Categories = categories
	v.Options = options
	v.BufferUsageReportingInterval = bufferUsageReportingInterval
	v.TransferMode = transferMode
	v.StreamFormat = streamFormat
	v.StreamCompression = streamCompression
	v.TraceConfig = traceConfig
	v.TracingBackend = tracingBackend
	return c.StartWithParams(ctx, &v)
}
