package cmd

import (
	"errors"
	"reflect"
	"testing"
)

func TestCheckArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{
			name:     "No arguments provided",
			args:     []string{},
			expected: errors.New("no entity name provided"),
		},
		{
			name:     "Only entity name provided",
			args:     []string{"entity"},
			expected: errors.New("url not provided"),
		},
		{
			name:     "Both entity name and URL provided",
			args:     []string{"entity", "http://example.com"},
			expected: nil,
		},
		{
			name:     "More than 2 arguments provided",
			args:     []string{"entity", "http://example.com", "extra"},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := checkArgs(test.args)
			if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) ||
				(err != nil && test.expected != nil && err.Error() != test.expected.Error()) {
				t.Errorf("got error '%v', expected '%v'", err, test.expected)
			}
		})
	}
}

func TestMapHeaders(t *testing.T) {
	tests := []struct {
		name       string
		headersStr string
		headers    map[string]string
		expected   error
	}{
		{
			name:       "Empty headers string",
			headersStr: "",
			headers:    map[string]string{},
			expected:   errors.New("headers in wrong format!! use key:value"),
		},
		{
			name:       "Invalid format - odd number of elements",
			headersStr: "key:value:key",
			headers:    map[string]string{},
			expected:   errors.New("headers in wrong format!! use key:value"),
		},
		{
			name:       "Invalid format - no colon separator",
			headersStr: "keyvalue",
			headers:    map[string]string{},
			expected:   errors.New("headers in wrong format!! use key:value"),
		},
		{
			name:       "Valid headers",
			headersStr: "Content-Type:application/json",
			headers:    map[string]string{"Content-Type": "application/json"},
			expected:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers = make(map[string]string)
			err := mapHeaders(test.headersStr)
			if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) ||
				(err != nil && test.expected != nil && err.Error() != test.expected.Error()) {
				t.Errorf("got error '%v', expected '%v'", err, test.expected)
			}
			if !reflect.DeepEqual(headers, test.headers) {
				t.Errorf(
					"headers not as expected. Got: %v, Expected: %v",
					headers,
					test.headers,
				)
			}
		})
	}
}
