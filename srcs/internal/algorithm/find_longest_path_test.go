package algorithm

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/input"
)

type FindLongestPath func(*graph.Graph) []graph.EdgeID

func TestFindLongestPath(t *testing.T) {
	tests := []struct {
		graphInputs []*input.GraphInput
		expected    []graph.EdgeID
	}{
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 8.54},
				{Start: 2, End: 3, Distance: 3.11},
				{Start: 3, End: 1, Distance: 2.19},
				{Start: 3, End: 4, Distance: 4.0},
				{Start: 4, End: 1, Distance: 1.4},
			},
			expected: []graph.EdgeID{2, 1, 4, 3, 2},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 1, End: 3, Distance: 2.0},
			},
			expected: []graph.EdgeID{3, 1, 2},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
			},
			expected: []graph.EdgeID{1, 2},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 2, End: 3, Distance: 1.0},
				{Start: 3, End: 4, Distance: 1.0},
				{Start: 4, End: 5, Distance: 1.0},
				{Start: 5, End: 1, Distance: 1.0},
			},
			expected: []graph.EdgeID{1, 2, 3, 4, 5, 1},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 1, End: 3, Distance: 1.0},
				{Start: 2, End: 4, Distance: 1.0},
				{Start: 2, End: 5, Distance: 1.0},
				{Start: 3, End: 6, Distance: 1.0},
				{Start: 3, End: 7, Distance: 1.0},
				{Start: 4, End: 8, Distance: 1.0},
				{Start: 4, End: 9, Distance: 5.0},
				{Start: 5, End: 10, Distance: 5.0},
				{Start: 5, End: 11, Distance: 1.0},
				{Start: 6, End: 12, Distance: 1.0},
				{Start: 6, End: 13, Distance: 1.0},
				{Start: 7, End: 14, Distance: 1.0},
				{Start: 7, End: 15, Distance: 1.0},
				{Start: 8, End: 15, Distance: 1.0},
			},
			expected: []graph.EdgeID{10, 5, 2, 1, 3, 7, 15, 8, 4, 9},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 1, End: 3, Distance: 1.0},
				{Start: 1, End: 4, Distance: 1.0},
				{Start: 1, End: 5, Distance: 2.0},
				{Start: 1, End: 6, Distance: 2.0},
				{Start: 1, End: 7, Distance: 1.0},
				{Start: 2, End: 4, Distance: 1.0},
				{Start: 2, End: 5, Distance: 1.0},
				{Start: 3, End: 6, Distance: 1.0},
				{Start: 3, End: 7, Distance: 1.0},
			},
			expected: []graph.EdgeID{4, 2, 5, 1, 6, 3, 7},
		},
	}

	algorithms := []FindLongestPath{FindLongestPathByDfs, FindLongestPathByDfsGoroutine}

	for _, algorithm := range algorithms {
		for _, test := range tests {
			g, _ := input.ConvertGraphInputsToGraph(test.graphInputs)
			result := algorithm(g)
			assertEdgeIDSlice(t, test.expected, result)
		}
	}
}

func assertEdgeIDSlice(t *testing.T, expected, actual []graph.EdgeID) {
	if len(expected) != len(actual) {
		t.Errorf("expected length: %d, actual length %d", len(expected), len(actual))
		return
	}

	// Check for loop patterns and remove elements from the slice if necessary.
	if isLoopPattern(expected, actual) {
		expected = expected[:len(expected)-1]
		actual = actual[:len(actual)-1]
	}

	// Check if slices match (even in reverse order)
	revExpected := reverse(expected)
	if !checkEdgeIDs(expected, actual) && !checkEdgeIDs(revExpected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func checkEdgeIDs(expected, actual []graph.EdgeID) bool {
	baseIdx := findBaseIndex(expected, actual)
	if baseIdx == -1 {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[(baseIdx+i)%len(actual)] {
			return false
		}
	}
	return true
}

func findBaseIndex(expected, actual []graph.EdgeID) int {
	for i, val := range actual {
		if expected[0] == val {
			return i
		}
	}
	return -1
}

func isLoopPattern(expected, actual []graph.EdgeID) bool {
	if len(expected) <= 2 || len(actual) <= 2 {
		return false
	}
	startIdx := 0
	endIdx := len(expected) - 1
	return expected[startIdx] == expected[endIdx] && actual[startIdx] == actual[endIdx]
}

func reverse(slice []graph.EdgeID) []graph.EdgeID {
	reversed := make([]graph.EdgeID, len(slice))

	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-i-1]
	}
	return reversed
}
