package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/stack"
)

func dfs(g *graph.Graph, startEdgeID graph.EdgeID) ([]graph.EdgeID, float64) {
	longestPath := []graph.EdgeID{}
	maxDistance := 0.0
	s := stack.NewStack()
	s.Push(stack.Item{Node: startEdgeID, Path: []graph.EdgeID{startEdgeID}})

	for !s.IsEmpty() {
		top, _ := s.Pop()
		distance, _ := g.CalcTotalDistance(top.Path)
		if distance > maxDistance {
			longestPath = top.Path
			maxDistance = distance
		}

		// If the start and end points are the same, no further search is performed.
		if isLoop(top.Path) {
			continue
		}

		for _, neighbor := range g.Adj[top.Node] {
			if !contains(top.Path, neighbor) || canCreateLoop(top.Path, neighbor.To) {
				newPath := append([]graph.EdgeID{}, append(top.Path, neighbor.To)...)
				s.Push(stack.Item{Node: neighbor.To, Path: newPath})
			}
		}
	}
	return longestPath, maxDistance
}

func isLoop(path []graph.EdgeID) bool {
	if len(path) <= 2 {
		return false
	}
	startEdgeID := path[0]
	endEdgeID := path[len(path)-1]
	return startEdgeID == endEdgeID
}

func canCreateLoop(path []graph.EdgeID, nextEdgeID graph.EdgeID) bool {
	// INFO The act of returning with a one-way ticket is against the rules.
	// INFO It is allowed for the start and end points to be the same due to loops.
	if len(path) <= 2 {
		return false
	}
	startEdgeID := path[0]
	return startEdgeID == nextEdgeID
}

func contains(path []graph.EdgeID, node graph.Edge) bool {
	for _, n := range path {
		if n == node.To {
			return true
		}
	}
	return false
}
