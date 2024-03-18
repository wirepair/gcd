package gcdapi_test

import (
	"github.com/wirepair/gcd/v2/gcdapi"
	"reflect"
	"testing"
)

func TestReturnValuesAreGeneratedProperlyForCommands(t *testing.T) {
	tests := []struct {
		name               string
		function           interface{}
		expectedReturnType string
	}{
		{
			name:               "a base type",
			function:           (&gcdapi.DOM{}).GetOuterHTML,
			expectedReturnType: "string",
		},
		{
			name:               "a reference to a complex type",
			function:           (&gcdapi.DOM{}).DescribeNode,
			expectedReturnType: "*gcdapi.DOMNode",
		},
		{
			name:               "an array of a base type",
			function:           (&gcdapi.DOM{}).CollectClassNamesFromSubtree,
			expectedReturnType: "[]string",
		},
		{
			name:               "an array to a reference type that references a base type",
			function:           (&gcdapi.DOM{}).GetNodesForSubtreeByStyle,
			expectedReturnType: "[]int",
		},
		{
			name:               "an array to a reference type that references an array of a base type",
			function:           (&gcdapi.DOMStorage{}).GetDOMStorageItems,
			expectedReturnType: "[][]string",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fnType := reflect.TypeOf(test.function)
			returnType := fnType.Out(0).String()

			if test.expectedReturnType != returnType {
				t.Logf("test failed, expected %v, got %v", test.expectedReturnType, returnType)
				t.Fail()
			}
		})
	}
}
