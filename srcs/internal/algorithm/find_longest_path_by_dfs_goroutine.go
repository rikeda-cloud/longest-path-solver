package algorithm

import (
	"sync"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func FindLongestPathByDfsGoroutine(g *graph.Graph) []graph.EdgeID {
	var longestPath []graph.EdgeID
	var maxDistance float64
	var mu sync.Mutex
	var wg sync.WaitGroup

	for startEdgeID := range g.Adj {
		wg.Add(1)
		go func(startEdgeID graph.EdgeID) {
			defer wg.Done()
			path, distance := dfs(g, startEdgeID)

			mu.Lock()
			if distance > maxDistance {
				maxDistance = distance
				longestPath = path
			}
			mu.Unlock()
		}(startEdgeID)
	}

	wg.Wait()
	return longestPath
}
