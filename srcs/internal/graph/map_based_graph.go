package graph

import (
	"errors"
)

type MapBasedGraph struct {
	Adj map[EdgeID]map[EdgeID]float64 // adj = "Adjacency Map"
}

func NewMapBasedGraph() *MapBasedGraph {
	return &MapBasedGraph{Adj: make(map[EdgeID]map[EdgeID]float64)}
}

func (g *MapBasedGraph) Equal(other *MapBasedGraph) bool {
	if len(g.Adj) != len(other.Adj) {
		return false
	}

	for from, edges := range g.Adj {
		otherEdges, exists := other.Adj[from]
		if !exists {
			return false
		}
		if len(edges) != len(otherEdges) {
			return false
		}
		for to, weight := range edges {
			if otherWeight, exists := otherEdges[to]; !exists || weight != otherWeight {
				return false
			}
		}
	}
	return true
}

func (g *MapBasedGraph) AddEdge(fromID, toID EdgeID, distance float64) error {
	if _, exist := g.FindDistance(fromID, toID); exist {
		return errors.New("Duplicate edges")
	}

	if g.Adj[fromID] == nil {
		g.Adj[fromID] = make(map[EdgeID]float64)
	}
	g.Adj[fromID][toID] = distance

	return nil
}

func (g *MapBasedGraph) FindDistance(fromID, toID EdgeID) (float64, bool) {
	edges, existEdges := g.Adj[fromID]
	if !existEdges {
		return 0.0, false
	}
	distance, existDistance := edges[toID]
	if !existDistance {
		return 0.0, false
	}
	return distance, true
}

func (g *MapBasedGraph) GetFromEdgeIDs() []EdgeID {
	fromEdgeIDs := make([]EdgeID, 0, len(g.Adj))
	for fromID := range g.Adj {
		fromEdgeIDs = append(fromEdgeIDs, fromID)
	}
	return fromEdgeIDs
}

func (g *MapBasedGraph) GetToEdgeIDs(fromID EdgeID) []EdgeID {
	toEdgeIDs := make([]EdgeID, 0, len(g.Adj[fromID]))
	for key := range g.Adj[fromID] {
		toEdgeIDs = append(toEdgeIDs, key)
	}
	return toEdgeIDs
}
