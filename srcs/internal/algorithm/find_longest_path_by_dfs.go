package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func FindLongestPathByDfs(g graph.IGraph) []graph.EdgeID {
	longestPath := []graph.EdgeID{}
	maxDistance := 0.0

	for _, startEdgeID := range g.GetFromEdgeIDs() {
		path, distance := dfs(g, startEdgeID)
		if distance > maxDistance {
			longestPath = path
			maxDistance = distance
		}
	}
	return longestPath
}
