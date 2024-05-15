package blocks

import "testing"

func TestHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "Empty headers",
			headers:  map[string]string{},
			expected: "headers {\n}",
		},
		{
			name: "Single header",
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			expected: "headers {\n  Content-Type: application/json\n}",
		},
		{
			name: "Multiple headers",
			headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "BearerToken",
			},
			expected: "headers {\n  Content-Type: application/json\n  Authorization: BearerToken\n}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Headers(test.headers)
			if result != test.expected {
				t.Errorf("got '%s', expected '%s'", result, test.expected)
			}
		})
	}
}

func TestMeta(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		nameStr  string
		seq      int
		expected string
	}{
		{
			name:     "Basic meta",
			action:   "GET",
			nameStr:  "example",
			seq:      1,
			expected: "meta {\n  name: GET example\n  type: http\n  seq: 1\n}\n\n",
		},
		{
			name:     "Meta with special characters",
			action:   "DELETE",
			nameStr:  "resource123",
			seq:      5,
			expected: "meta {\n  name: DELETE resource123\n  type: http\n  seq: 5\n}\n\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Meta(test.action, test.nameStr, test.seq)
			if result != test.expected {
				t.Errorf("got '%s', expected '%s'", result, test.expected)
			}
		})
	}
}

func TestMethod(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		url      string
		nameStr  string
		path     string
		expected string
	}{
		{
			name:     "Basic method",
			method:   "GET",
			url:      "http://example.com",
			nameStr:  "endpoint",
			path:     "",
			expected: "GET {\n  url: http://example.com/endpoint\n  body: json\n  auth: none\n}\n\n",
		},
		{
			name:     "Method with name and path",
			method:   "PUT",
			url:      "http://example.com",
			nameStr:  "endpoint",
			path:     "/resource/123",
			expected: "PUT {\n  url: http://example.com/endpoint/resource/123\n  body: json\n  auth: none\n}\n\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Method(test.method, test.url, test.nameStr, test.path)
			if result != test.expected {
				t.Errorf("got '%s', expected '%s'", result, test.expected)
			}
		})
	}
}
