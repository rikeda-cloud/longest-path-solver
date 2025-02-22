package input

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func TestConvertGraphInputsToGraph(t *testing.T) {
	tests := []struct {
		graphInputs []*GraphInput
		expected    *graph.Graph
	}{
		{
			graphInputs: []*GraphInput{
				{Start: 1, End: 2, Distance: 8.54},
				{Start: 2, End: 3, Distance: 3.11},
				{Start: 3, End: 1, Distance: 2.19},
				{Start: 3, End: 4, Distance: 4.0},
				{Start: 4, End: 1, Distance: 1.4},
			},
			expected: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 8.54}, {To: 3, Distance: 2.19}, {To: 4, Distance: 1.4}},
				2: {{To: 1, Distance: 8.54}, {To: 3, Distance: 3.11}},
				3: {{To: 2, Distance: 3.11}, {To: 1, Distance: 2.19}, {To: 4, Distance: 4.0}},
				4: {{To: 3, Distance: 4.0}, {To: 1, Distance: 1.4}},
			}),
		},
		{
			graphInputs: []*GraphInput{
				{Start: 1, End: 2, Distance: 1.0},
				{Start: 2, End: 3, Distance: 2.0},
				{Start: 3, End: 4, Distance: 3.0},
				{Start: 4, End: 5, Distance: 4.0},
				{Start: 5, End: 6, Distance: 5.0},
				{Start: 6, End: 1, Distance: 6.0},
			},
			expected: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				1: {{To: 2, Distance: 1.0}, {To: 6, Distance: 6.0}},
				2: {{To: 1, Distance: 1.0}, {To: 3, Distance: 2.0}},
				3: {{To: 2, Distance: 2.0}, {To: 4, Distance: 3.0}},
				4: {{To: 3, Distance: 3.0}, {To: 5, Distance: 4.0}},
				5: {{To: 4, Distance: 4.0}, {To: 6, Distance: 5.0}},
				6: {{To: 5, Distance: 5.0}, {To: 1, Distance: 6.0}},
			}),
		},
		{
			graphInputs: []*GraphInput{{Start: 4, End: 5, Distance: 6.0}},
			expected: graph.NewGraphWithEdges(map[graph.EdgeID][]graph.Edge{
				4: {{To: 5, Distance: 6.0}},
				5: {{To: 4, Distance: 6.0}},
			}),
		},
	}

	for _, test := range tests {
		result, _ := ConvertGraphInputsToGraph(test.graphInputs)
		assertGraphEqual(t, test.expected, result)
	}
}

func assertGraphEqual(t *testing.T, expected, actual *graph.Graph) {
	if expected == nil && actual == nil {
		return // Both are nil, considered equal
	}
	if expected == nil || actual == nil {
		t.Errorf("one of the graphs is nil: expected %v, actual %v", expected, actual)
		return
	}
	if expected.Equal(actual) == false {
		t.Errorf("graphs are not equal: expected %+v, actual %+v", expected, actual)
	}
}
