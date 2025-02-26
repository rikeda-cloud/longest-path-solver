package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/stack"
)

func dfs(g graph.IGraph, startEdgeID graph.EdgeID) ([]graph.EdgeID, float64) {
	longestPath := []graph.EdgeID{}
	maxDistance := 0.0
	s := stack.NewStack[graph.PathNode]()
	s.Push(graph.PathNode{
		Path: []graph.EdgeID{startEdgeID}, Node: startEdgeID, Distance: 0.0,
	})

	for !s.IsEmpty() {
		top, _ := s.Pop()
		if top.Distance > maxDistance {
			longestPath = top.Path
			maxDistance = top.Distance
		}

		// If the start and end points are the same, no further search is performed.
		if isLoop(top.Path) {
			continue
		}

		for _, neighborID := range g.GetToEdgeIDSlice(top.Node) {
			if !contains(top.Path, neighborID) || canCreateLoop(top.Path, neighborID) {
				newPath := make([]graph.EdgeID, len(top.Path)+1)
				copy(newPath, top.Path)
				newPath[len(top.Path)] = neighborID

				distanceToNeighbor, _ := g.FindDistance(top.Node, neighborID)
				newDistance := top.Distance + distanceToNeighbor
				s.Push(graph.PathNode{
					Path: newPath, Node: neighborID, Distance: newDistance,
				})
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

func contains(path []graph.EdgeID, edgeID graph.EdgeID) bool {
	for _, n := range path {
		if n == edgeID {
			return true
		}
	}
	return false
}
