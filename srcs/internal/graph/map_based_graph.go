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

func (g *MapBasedGraph) AddEdge(id1, id2 EdgeID, distance float64) error {
	if _, exist := g.FindDistance(id1, id2); exist {
		return errors.New("edge already exists")
	}

	if g.Adj[id1] == nil {
		g.Adj[id1] = make(map[EdgeID]float64)
	}
	if g.Adj[id2] == nil {
		g.Adj[id2] = make(map[EdgeID]float64)
	}

	g.Adj[id1][id2] = distance
	g.Adj[id2][id1] = distance

	return nil
}

func (g *MapBasedGraph) FindDistance(fromID, toID EdgeID) (float64, bool) {
	if edges, existEdges := g.Adj[fromID]; existEdges {
		if distance, existDistance := edges[toID]; existDistance {
			return distance, true
		}
	}
	return 0.0, false
}

func (g *MapBasedGraph) GetFromEdgeIDSlice() []EdgeID {
	fromEdgeIDSlice := make([]EdgeID, 0, len(g.Adj))
	for fromID := range g.Adj {
		fromEdgeIDSlice = append(fromEdgeIDSlice, fromID)
	}
	return fromEdgeIDSlice
}

func (g *MapBasedGraph) GetToEdgeIDSlice(fromID EdgeID) []EdgeID {
	toEdgeIDSlice := make([]EdgeID, 0, len(g.Adj[fromID]))
	for key := range g.Adj[fromID] {
		toEdgeIDSlice = append(toEdgeIDSlice, key)
	}
	return toEdgeIDSlice
}
