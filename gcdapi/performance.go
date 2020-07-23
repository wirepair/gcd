// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Performance functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Run-time execution metric.
type PerformanceMetric struct {
	Name  string  `json:"name"`  // Metric name.
	Value float64 `json:"value"` // Metric value.
}

// Current values of the metrics.
type PerformanceMetricsEvent struct {
	Method string `json:"method"`
	Params struct {
		Metrics []*PerformanceMetric `json:"metrics"` // Current values of the metrics.
		Title   string               `json:"title"`   // Timestamp title.
	} `json:"Params,omitempty"`
}

type Performance struct {
	target gcdmessage.ChromeTargeter
}

func NewPerformance(target gcdmessage.ChromeTargeter) *Performance {
	c := &Performance{target: target}
	return c
}

// Disable collecting and reporting metrics.
func (c *Performance) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.disable"})
}

type PerformanceEnableParams struct {
	// Time domain to use for collecting and reporting duration metrics.
	TimeDomain string `json:"timeDomain,omitempty"`
}

// EnableWithParams - Enable collecting and reporting metrics.
func (c *Performance) EnableWithParams(v *PerformanceEnableParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.enable", Params: v})
}

// Enable - Enable collecting and reporting metrics.
// timeDomain - Time domain to use for collecting and reporting duration metrics.
func (c *Performance) Enable(timeDomain string) (*gcdmessage.ChromeResponse, error) {
	var v PerformanceEnableParams
	v.TimeDomain = timeDomain
	return c.EnableWithParams(&v)
}

type PerformanceSetTimeDomainParams struct {
	// Time domain
	TimeDomain string `json:"timeDomain"`
}

// SetTimeDomainWithParams - Sets time domain to use for collecting and reporting duration metrics. Note that this must be called before enabling metrics collection. Calling this method while metrics collection is enabled returns an error.
func (c *Performance) SetTimeDomainWithParams(v *PerformanceSetTimeDomainParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.setTimeDomain", Params: v})
}

// SetTimeDomain - Sets time domain to use for collecting and reporting duration metrics. Note that this must be called before enabling metrics collection. Calling this method while metrics collection is enabled returns an error.
// timeDomain - Time domain
func (c *Performance) SetTimeDomain(timeDomain string) (*gcdmessage.ChromeResponse, error) {
	var v PerformanceSetTimeDomainParams
	v.TimeDomain = timeDomain
	return c.SetTimeDomainWithParams(&v)
}

// GetMetrics - Retrieve current values of run-time metrics.
// Returns -  metrics - Current values for run-time metrics.
func (c *Performance) GetMetrics() ([]*PerformanceMetric, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.getMetrics"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Metrics []*PerformanceMetric
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

	return chromeData.Result.Metrics, nil
}
