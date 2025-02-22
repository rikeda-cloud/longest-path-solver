package input

import (
	"strings"
	"testing"
)

func TestParseGraphInputs(t *testing.T) {
	// INFO The input is assumed to be in the specified format.
	// ex) "始点の ID(正の整数値), 終点の ID(正の整数値), 距離(浮動小数点数)¥r¥n"
	tests := []struct {
		input    string
		expected []*GraphInput
	}{
		{
			input: "1, 2, 8.54\r\n2, 3, 3.11\r\n3, 1, 2.19\r\n3, 4, 4\r\n4, 1, 1.4\r\n",
			expected: []*GraphInput{
				{Start: 1, End: 2, Distance: 8.54},
				{Start: 2, End: 3, Distance: 3.11},
				{Start: 3, End: 1, Distance: 2.19},
				{Start: 3, End: 4, Distance: 4.0},
				{Start: 4, End: 1, Distance: 1.4},
			},
		},
		{
			input: "1,2,1.0\r\n2,3,2.0\r\n3,4,3.0\r\n4,5,4.0\r\n5,6,5.0\r\n6,1,6.0\r\n",
			expected: []*GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 2, End: 3, Distance: 2.0},
				{Start: 3, End: 4, Distance: 3.0},
				{Start: 4, End: 5, Distance: 4.0},
				{Start: 5, End: 6, Distance: 5.0},
				{Start: 6, End: 1, Distance: 6.0},
			},
		},
		{
			input: " 4 , 5 , 6 \r\n",
			expected: []*GraphInput{
				{Start: 4, End: 5, Distance: 6.0},
			},
		},
		{input: "", expected: nil},
		{input: "1,2,3\r\n\r\n4,5,6\r\n", expected: nil},
		{input: "1,2,3\r\n4,5,6\r\n\r\n", expected: nil},
		{input: "1,2,8.54,\r\n", expected: nil},
		{input: "1,2,String\r\n", expected: nil},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		result, _ := ParseGraphInputs(reader)
		assertEdgesEqual(t, test.expected, result)
	}
}

func assertEdgesEqual(t *testing.T, expected, actual []*GraphInput) {
	if expected == nil && actual == nil {
		return // Both are nil, considered equal
	}
	if expected == nil || actual == nil {
		t.Errorf("expected %v, got %v", expected, actual)
		return
	}
	if len(expected) != len(actual) {
		t.Errorf("expected %d edges, got %d", len(expected), len(actual))
		return
	}
	for i := 0; i < len(expected); i++ {
		if *expected[i] != *actual[i] {
			t.Errorf("unexpected edge got %+v, want %+v", actual[i], expected[i])
			return
		}
	}
}
