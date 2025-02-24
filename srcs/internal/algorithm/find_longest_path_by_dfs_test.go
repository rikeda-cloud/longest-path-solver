package algorithm

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func TestFindLongestPathByDfs(t *testing.T) {
	tests := []struct {
		input    *graph.Graph
		expected []graph.EdgeID
	}{
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {
					{To: 2, Distance: 8.54},
					{To: 3, Distance: 2.19},
					{To: 4, Distance: 1.4},
				},
				2: {{To: 1, Distance: 8.54}, {To: 3, Distance: 3.11}},
				3: {
					{To: 2, Distance: 3.11},
					{To: 1, Distance: 2.19},
					{To: 4, Distance: 4.0},
				},
				4: {{To: 3, Distance: 4.0}, {To: 1, Distance: 1.4}},
			}),
			expected: []graph.EdgeID{2, 1, 4, 3, 2},
		},
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 1.0}, {To: 3, Distance: 2.0}},
				2: {{To: 1, Distance: 1.0}},
				3: {{To: 1, Distance: 2.0}},
			}),
			expected: []graph.EdgeID{3, 1, 2},
		},
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 1.0}},
				2: {{To: 1, Distance: 1.0}},
			}),
			expected: []graph.EdgeID{1, 2},
		},
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 1.0}, {To: 5, Distance: 1.0}},
				2: {{To: 1, Distance: 1.0}, {To: 3, Distance: 1.0}},
				3: {{To: 2, Distance: 1.0}, {To: 4, Distance: 1.0}},
				4: {{To: 3, Distance: 1.0}, {To: 5, Distance: 1.0}},
				5: {{To: 4, Distance: 1.0}, {To: 1, Distance: 1.0}},
			}),
			expected: []graph.EdgeID{1, 2, 3, 4, 5, 1},
		},
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 1.0}, {To: 3, Distance: 1.0}},
				2: {
					{To: 1, Distance: 1.0},
					{To: 4, Distance: 1.0},
					{To: 5, Distance: 1.0},
				},
				3: {
					{To: 1, Distance: 1.0},
					{To: 6, Distance: 1.0},
					{To: 7, Distance: 1.0},
				},
				4: {
					{To: 2, Distance: 1.0},
					{To: 8, Distance: 1.0},
					{To: 9, Distance: 5.0},
				},
				5: {
					{To: 2, Distance: 1.0},
					{To: 10, Distance: 5.0},
					{To: 11, Distance: 1.0},
				},
				6: {
					{To: 3, Distance: 1.0},
					{To: 12, Distance: 1.0},
					{To: 13, Distance: 1.0},
				},
				7: {
					{To: 3, Distance: 1.0},
					{To: 14, Distance: 1.0},
					{To: 15, Distance: 1.0},
				},
				8:  {{To: 4, Distance: 1.0}, {To: 15, Distance: 1.0}},
				9:  {{To: 4, Distance: 5.0}},
				10: {{To: 5, Distance: 5.0}},
				11: {{To: 5, Distance: 1.0}},
				12: {{To: 6, Distance: 1.0}},
				13: {{To: 6, Distance: 1.0}},
				14: {{To: 7, Distance: 1.0}},
				15: {{To: 7, Distance: 1.0}, {To: 8, Distance: 1.0}},
			}),
			expected: []graph.EdgeID{10, 5, 2, 1, 3, 7, 15, 8, 4, 9},
		},
		{
			input: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {
					{To: 2, Distance: 1.0},
					{To: 3, Distance: 1.0},
					{To: 4, Distance: 1.0},
					{To: 5, Distance: 2.0},
					{To: 6, Distance: 2.0},
					{To: 7, Distance: 1.0},
				},
				2: {
					{To: 1, Distance: 1.0},
					{To: 4, Distance: 1.0},
					{To: 5, Distance: 1.0},
				},
				3: {
					{To: 1, Distance: 1.0},
					{To: 6, Distance: 1.0},
					{To: 7, Distance: 1.0},
				},
				4: {{To: 2, Distance: 1.0}, {To: 1, Distance: 1.0}},
				5: {{To: 2, Distance: 1.0}, {To: 1, Distance: 2.0}},
				6: {{To: 3, Distance: 1.0}, {To: 1, Distance: 2.0}},
				7: {{To: 3, Distance: 1.0}, {To: 1, Distance: 1.0}},
			}),
			expected: []graph.EdgeID{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, test := range tests {
		result := FindLongestPathByDfs(test.input)
		assertEdgeIDSlice(t, test.expected, result)
	}
}

func assertEdgeIDSlice(t *testing.T, expected, actual []graph.EdgeID) {
	if len(expected) != len(actual) {
		for _, id := range actual {
			println(id)
		}
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

	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		reversed[i] = slice[j]
	}
	return reversed
}
