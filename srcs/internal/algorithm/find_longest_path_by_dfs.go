package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func FindLongestPathByDfs(g *graph.Graph) []graph.EdgeID {
	longestPath := []graph.EdgeID{}
	maxDistance := 0.0

	for startEdgeID := range g.Adj {
		path, distance := dfs(g, startEdgeID)
		if distance > maxDistance {
			longestPath = path
			maxDistance = distance
		}
	}
	return longestPath
}
