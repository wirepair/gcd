// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Testing functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

type Testing struct {
	target gcdmessage.ChromeTargeter
}

func NewTesting(target gcdmessage.ChromeTargeter) *Testing {
	c := &Testing{target: target}
	return c
}

type TestingGenerateTestReportParams struct {
	// Message to be displayed in the report.
	Message string `json:"message"`
	// Specifies the endpoint group to deliver the report to.
	Group string `json:"group,omitempty"`
}

// GenerateTestReportWithParams - Generates a report for testing.
func (c *Testing) GenerateTestReportWithParams(v *TestingGenerateTestReportParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Testing.generateTestReport", Params: v})
}

// GenerateTestReport - Generates a report for testing.
// message - Message to be displayed in the report.
// group - Specifies the endpoint group to deliver the report to.
func (c *Testing) GenerateTestReport(message string, group string) (*gcdmessage.ChromeResponse, error) {
	var v TestingGenerateTestReportParams
	v.Message = message
	v.Group = group
	return c.GenerateTestReportWithParams(&v)
}
