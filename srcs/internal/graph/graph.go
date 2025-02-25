package graph

import (
	"errors"
)

type Edge struct {
	To       EdgeID
	Distance float64
}

type Graph struct {
	Adj map[EdgeID][]Edge // adj = "Adjacency List"
}

func NewGraph() *Graph {
	return &Graph{Adj: make(map[EdgeID][]Edge)}
}

func (g *Graph) Equal(other *Graph) bool {
	// Compare the number of edges in both graphs
	if len(g.Adj) != len(other.Adj) {
		return false
	}

	for key, edges := range g.Adj {
		otherEdges, exists := other.Adj[key]
		// Check if the other graph has the same key and the same number of edges
		if !exists || len(edges) != len(otherEdges) {
			return false
		}

		// Compare the edges for each key
		for i := range edges {
			if edges[i] != otherEdges[i] {
				return false
			}
		}
	}
	return true
}

func (g *Graph) AddEdge(id1, id2 EdgeID, distance float64) error {
	if _, exist := g.findEdge(id1, id2); exist {
		return errors.New("edge already exists")
	}

	// Adds an edge between id1 and id2 with the given distance,
	// ensuring both directions are stored for the undirected graph.
	g.Adj[id1] = append(g.Adj[id1], Edge{To: id2, Distance: distance})
	g.Adj[id2] = append(g.Adj[id2], Edge{To: id1, Distance: distance})

	return nil
}

func (g *Graph) GetFromEdgeIDSlice() []EdgeID {
	fromEdgeIDSlice := make([]EdgeID, 0, len(g.Adj))
	for fromEdgeID := range g.Adj {
		fromEdgeIDSlice = append(fromEdgeIDSlice, fromEdgeID)
	}
	return fromEdgeIDSlice
}

func (g *Graph) GetToEdgeIDSlice(fromEdgeID EdgeID) []EdgeID {
	edges := g.Adj[fromEdgeID]
	toEdgeIDSlice := make([]EdgeID, 0, len(edges))
	for _, edge := range edges {
		toEdgeIDSlice = append(toEdgeIDSlice, edge.To)
	}
	return toEdgeIDSlice
}

func (g *Graph) CalcTotalDistance(path []EdgeID) (float64, bool) {
	totalDistance := 0.0
	for i := 0; i < len(path)-1; i++ {
		fromID := path[i]
		toID := path[i+1]
		edge, ok := g.findEdge(fromID, toID)
		if !ok {
			return 0.0, false
		}
		totalDistance += edge.Distance
	}
	return totalDistance, true
}

func (g *Graph) findEdge(fromEdgeID, toEdgeID EdgeID) (Edge, bool) {
	for _, edge := range g.Adj[fromEdgeID] {
		if edge.To == toEdgeID {
			return edge, true
		}
	}
	return Edge{}, false
}
