package algorithm

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/input"
)

type FindLongestPath func(graph.IGraph) []graph.EdgeID

func createTests() []struct {
	graphInputs []*input.GraphInput
	expected    []graph.EdgeID
} {
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
			expected: []graph.EdgeID{1, 2, 3, 4, 1},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 1, End: 3, Distance: 2.0},
			},
			expected: []graph.EdgeID{1, 3},
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
				{Start: 4, End: 9, Distance: 1.0},
				{Start: 5, End: 10, Distance: 1.0},
				{Start: 5, End: 11, Distance: 1.0},
				{Start: 6, End: 12, Distance: 1.0},
				{Start: 6, End: 13, Distance: 1.0},
				{Start: 7, End: 14, Distance: 1.0},
				{Start: 7, End: 15, Distance: 1.0},
				{Start: 8, End: 15, Distance: 1.0},
			},
			expected: []graph.EdgeID{1, 2, 4, 8, 15},
		},
		{
			graphInputs: []*input.GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 1, End: 3, Distance: 1.0},
				{Start: 1, End: 4, Distance: 1.0},
				{Start: 1, End: 5, Distance: 1.0},
				{Start: 1, End: 6, Distance: 1.0},
				{Start: 1, End: 7, Distance: 1.0},
				{Start: 2, End: 4, Distance: 1.0},
				{Start: 2, End: 5, Distance: 2.0},
				{Start: 3, End: 6, Distance: 1.0},
				{Start: 3, End: 7, Distance: 1.0},
			},
			expected: []graph.EdgeID{1, 2, 5},
		},
	}
	return tests
}

func TestFindLongestPathGraph(t *testing.T) {
	tests := createTests()
	algorithms := []FindLongestPath{FindLongestPathByDfs, FindLongestPathByDfsGoroutine}

	for _, algorithm := range algorithms {
		for _, test := range tests {
			g, _ := input.ConvertGraphInputsToGraph(test.graphInputs, graph.NewGraph())
			result := algorithm(g)
			assertEdgeIDs(t, test.expected, result)
		}
	}
}

func TestFindLongestPathMapBasedGraph(t *testing.T) {
	tests := createTests()
	algorithms := []FindLongestPath{FindLongestPathByDfs, FindLongestPathByDfsGoroutine}

	for _, algorithm := range algorithms {
		for _, test := range tests {
			g, _ := input.ConvertGraphInputsToGraph(test.graphInputs, graph.NewMapBasedGraph())
			result := algorithm(g)
			assertEdgeIDs(t, test.expected, result)
		}
	}
}

func assertEdgeIDs(t *testing.T, expected, actual []graph.EdgeID) {
	if len(expected) != len(actual) {
		t.Errorf("expected length: %d, actual length %d", len(expected), len(actual))
		return
	}

	if hasLoop(expected, actual) {
		if !checkLoopEdgeIDs(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	} else {
		if !checkEdgeIDs(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func checkEdgeIDs(expected, actual []graph.EdgeID) bool {
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}

func checkLoopEdgeIDs(expected, actual []graph.EdgeID) bool {
	// Remove the last element and get rid of the loop.
	expected = expected[:len(expected)-1]
	actual = actual[:len(actual)-1]

	// Any edge can be asserted as the starting point.
	// ex) if expected[1,2,3,4,5], actual[3,4,5,1,2] => baseIdx is 3
	//     if expected[1,2,3], actual[4,5,6] => baseIdx is -1
	baseIdx := -1
	for i, val := range actual {
		if expected[0] == val {
			baseIdx = i
			break
		}
	}
	if baseIdx == -1 {
		return false
	}

	// ex) if expected[1,2,3,4,5], actual[3,4,5,1,2] => true
	//     if expected[1,2,3,4,5], actual[0,1,2,3,4] => false
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[(baseIdx+i)%len(actual)] {
			return false
		}
	}
	return true
}

func hasLoop(expected, actual []graph.EdgeID) bool {
	// Determine if there is a loop in both expected and actual.
	if len(expected) <= 2 || len(actual) <= 2 {
		return false
	}
	startIdx := 0
	endIdx := len(expected) - 1
	hasLoopExpected := expected[startIdx] == expected[endIdx]
	hasLoopActual := actual[startIdx] == actual[endIdx]
	return hasLoopExpected && hasLoopActual
}
