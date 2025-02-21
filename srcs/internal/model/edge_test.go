package edge

import (
	"testing"
)

func TestParseGraphEdge(t *testing.T) {
	tests := []struct {
		input    string
		expected *Edge
	}{
		{input: "1, 2, 8.54", expected: &Edge{Start: 1, End: 2, Distance: 8.54}},
		{input: "2, 3, 3.11", expected: &Edge{Start: 2, End: 3, Distance: 3.11}},
		{input: "3, 1, 2.19", expected: &Edge{Start: 3, End: 1, Distance: 2.19}},
		{input: "3, 4, 4", expected: &Edge{Start: 3, End: 4, Distance: 4.0}},
		{input: "4, 1, 1.4", expected: &Edge{Start: 4, End: 1, Distance: 1.4}},
		{input: "1,2,3.45", expected: &Edge{Start: 1, End: 2, Distance: 3.45}},
		{input: " 1 ,2 ,0.345 ", expected: &Edge{Start: 1, End: 2, Distance: 0.345}},
		{input: "  1,2  ,.123 ", expected: &Edge{Start: 1, End: 2, Distance: 0.123}},
		{input: "1,2,123.", expected: &Edge{Start: 1, End: 2, Distance: 123.0}},
		{input: " 01,002,0003 ", expected: &Edge{Start: 1, End: 2, Distance: 3.0}},
		{input: "0, 2, 1.5", expected: nil},   // Start is 0
		{input: "1, 0, 1.5", expected: nil},   // End is 0
		{input: "	1,	2,	8.54", expected: nil}, // Only whitespace is allowed (0x20)
		{input: "1, 2, .", expected: nil},
		{input: "1, 2, 1.1.1", expected: nil},
		{input: "1, 2, 1.1.", expected: nil},
		{input: "1, 2, .1.1", expected: nil},
		{input: "1, 2, 1..1", expected: nil},
		{input: "1, 2, ..1", expected: nil},
		{input: "1, 2, 1..", expected: nil},
		{input: "invalid input", expected: nil},
	}

	for _, test := range tests {
		result, _ := ParseGraphEdge(test.input)
		assertEdgeEqual(t, test.expected, result)
	}
}

func assertEdgeEqual(t *testing.T, expected, actual *Edge) {
	if expected == nil && actual == nil {
		return // Both are nil, considered equal
	}
	if expected == nil || actual == nil {
		t.Errorf("expected %v, got %v", expected, actual)
		return
	}
	if *expected != *actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
