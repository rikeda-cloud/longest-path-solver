package algorithm

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func setupGraph(g graph.IGraph) graph.IGraph {
	_ = g.AddEdge(2, 1, 2.0)
	_ = g.AddEdge(3, 1, 3.0)
	_ = g.AddEdge(3, 2, 3.0)
	_ = g.AddEdge(4, 1, 4.0)
	_ = g.AddEdge(4, 2, 4.0)
	_ = g.AddEdge(4, 3, 4.0)
	_ = g.AddEdge(5, 1, 5.0)
	_ = g.AddEdge(5, 2, 5.0)
	_ = g.AddEdge(5, 3, 5.0)
	_ = g.AddEdge(5, 4, 5.0)
	_ = g.AddEdge(6, 1, 6.0)
	_ = g.AddEdge(6, 2, 6.0)
	_ = g.AddEdge(6, 3, 6.0)
	_ = g.AddEdge(6, 4, 6.0)
	_ = g.AddEdge(6, 5, 6.0)
	_ = g.AddEdge(7, 1, 7.0)
	_ = g.AddEdge(7, 2, 7.0)
	_ = g.AddEdge(7, 3, 7.0)
	_ = g.AddEdge(7, 4, 7.0)
	_ = g.AddEdge(7, 5, 7.0)
	_ = g.AddEdge(7, 6, 7.0)
	_ = g.AddEdge(8, 1, 8.0)
	_ = g.AddEdge(8, 2, 8.0)
	_ = g.AddEdge(8, 3, 8.0)
	_ = g.AddEdge(8, 4, 8.0)
	_ = g.AddEdge(8, 5, 8.0)
	_ = g.AddEdge(8, 6, 8.0)
	_ = g.AddEdge(8, 7, 8.0)
	_ = g.AddEdge(1, 2, 2.0)
	_ = g.AddEdge(1, 3, 3.0)
	_ = g.AddEdge(2, 3, 3.0)
	_ = g.AddEdge(1, 4, 4.0)
	_ = g.AddEdge(2, 4, 4.0)
	_ = g.AddEdge(3, 4, 4.0)
	_ = g.AddEdge(1, 5, 5.0)
	_ = g.AddEdge(2, 5, 5.0)
	_ = g.AddEdge(3, 5, 5.0)
	_ = g.AddEdge(4, 5, 5.0)
	_ = g.AddEdge(1, 6, 6.0)
	_ = g.AddEdge(2, 6, 6.0)
	_ = g.AddEdge(3, 6, 6.0)
	_ = g.AddEdge(4, 6, 6.0)
	_ = g.AddEdge(5, 6, 6.0)
	_ = g.AddEdge(1, 7, 7.0)
	_ = g.AddEdge(2, 7, 7.0)
	_ = g.AddEdge(3, 7, 7.0)
	_ = g.AddEdge(4, 7, 7.0)
	_ = g.AddEdge(5, 7, 7.0)
	_ = g.AddEdge(6, 7, 7.0)
	_ = g.AddEdge(1, 8, 8.0)
	_ = g.AddEdge(2, 8, 8.0)
	_ = g.AddEdge(3, 8, 8.0)
	_ = g.AddEdge(4, 8, 8.0)
	_ = g.AddEdge(5, 8, 8.0)
	_ = g.AddEdge(6, 8, 8.0)
	_ = g.AddEdge(7, 8, 8.0)
	return g
}

func BenchmarkFindLongestPathByDfs(b *testing.B) {
	g := setupGraph(graph.NewGraph())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindLongestPathByDfs(g)
	}
}

func BenchmarkFindLongestPathByDfsGoroutine(b *testing.B) {
	g := setupGraph(graph.NewGraph())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindLongestPathByDfsGoroutine(g)
	}
}

func BenchmarkFindLongestPathByDfsMapBasedGraph(b *testing.B) {
	g := setupGraph(graph.NewMapBasedGraph())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindLongestPathByDfs(g)
	}
}

func BenchmarkFindLongestPathByDfsGoroutineMapBasedGraph(b *testing.B) {
	g := setupGraph(graph.NewMapBasedGraph())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindLongestPathByDfsGoroutine(g)
	}
}
