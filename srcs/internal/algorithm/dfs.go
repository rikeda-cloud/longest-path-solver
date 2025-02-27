package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/stack"
)

func dfs(g graph.IGraph, startEdgeID graph.EdgeID) ([]graph.EdgeID, float64) {
	longestPath := []graph.EdgeID{}
	maxDistance := 0.0
	s := stack.NewStack[PathNode]()
	s.Push(PathNode{
		Path: []graph.EdgeID{startEdgeID}, Node: startEdgeID, Distance: 0.0,
	})

	for !s.IsEmpty() { // Continue processing until stack is empty.
		top, _ := s.Pop()
		// Update the value if the distance exceeds the maximum distance.
		if top.Distance > maxDistance {
			longestPath = top.Path
			maxDistance = top.Distance
		}

		// If the start and end points are the same, no further search is performed.
		if isLoop(top.Path, top.Path[len(top.Path)-1]) {
			continue
		}

		for _, neighborID := range g.GetToEdgeIDs(top.Node) {
			// If it's an unexplored edge or the starting edge, add it to the stack.
			if !contains(top.Path, neighborID) || canCreateLoop(top.Path, neighborID) {
				// Create a new path by appending the neighborID to the current path.
				newPath := make([]graph.EdgeID, len(top.Path)+1)
				copy(newPath, top.Path)
				newPath[len(top.Path)] = neighborID

				// Add the distance to the adjacent edge to the total distance.
				distanceToNeighbor, _ := g.FindDistance(top.Node, neighborID)
				newDistance := top.Distance + distanceToNeighbor
				s.Push(PathNode{
					Path: newPath, Node: neighborID, Distance: newDistance,
				})
			}
		}
	}
	return longestPath, maxDistance
}

func isLoop(path []graph.EdgeID, endEdgeID graph.EdgeID) bool {
	if len(path) <= 1 {
		return false
	}
	startEdgeID := path[0]
	return startEdgeID == endEdgeID
}

func canCreateLoop(path []graph.EdgeID, nextEdgeID graph.EdgeID) bool {
	// INFO The act of returning with a one-way ticket is against the rules.
	// INFO It is allowed for the start and end points to be the same due to loops.
	return isLoop(path, nextEdgeID)
}

func contains(path []graph.EdgeID, edgeID graph.EdgeID) bool {
	for _, n := range path {
		if n == edgeID {
			return true
		}
	}
	return false
}
