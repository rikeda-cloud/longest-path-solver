package input

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func ConvertGraphInputsToGraph(graphInputs []*GraphInput) (*graph.Graph, error) {
	g := graph.NewGraph()

	for _, graphInput := range graphInputs {
		err := g.AddEdge(graphInput.Start, graphInput.End, graphInput.Distance)
		if err != nil {
			return nil, err
		}
	}
	return g, nil
}
