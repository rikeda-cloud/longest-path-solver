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

func NewGraphWithEdges(edges map[EdgeID][]Edge) *Graph {
	adj := make(map[EdgeID][]Edge)
	for key, edges := range edges {
		adj[key] = edges
	}
	return &Graph{Adj: adj}
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
	if g.edgeExists(id1, id2) {
		return errors.New("edge already exists")
	}

	// Adds an edge between id1 and id2 with the given distance,
	// ensuring both directions are stored for the undirected graph.
	g.Adj[id1] = append(g.Adj[id1], Edge{To: id2, Distance: distance})
	g.Adj[id2] = append(g.Adj[id2], Edge{To: id1, Distance: distance})

	return nil
}

func (g *Graph) CalcTotalDistance(path []EdgeID) (float64, bool) {
	totalDistance := 0.0
	for i := 0; i < len(path)-1; i++ {
		fromID := path[i]
		toID := path[i+1]
		distance, ok := g.getDistance(fromID, toID)
		if !ok {
			return 0.0, false
		}
		totalDistance += distance
	}
	return totalDistance, true
}

func (g *Graph) getDistance(fromEdgeID, toEdgeID EdgeID) (float64, bool) {
	for _, edge := range g.Adj[fromEdgeID] {
		if edge.To == toEdgeID {
			return edge.Distance, true
		}
	}
	return 0.0, false
}

func (g *Graph) edgeExists(id1, id2 EdgeID) bool {
	for _, edge := range g.Adj[id1] {
		if edge.To == id2 {
			return true
		}
	}
	return false
}
