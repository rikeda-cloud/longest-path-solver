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
	if len(g.GetFromEdgeIDSlice()) != len(other.GetFromEdgeIDSlice()) {
		return false
	}

	for _, fromEdgeID := range g.GetFromEdgeIDSlice() {
		gToIDs := g.GetToEdgeIDSlice(fromEdgeID)
		otherToIDs := other.GetToEdgeIDSlice(fromEdgeID)

		if len(gToIDs) != len(otherToIDs) {
			return false
		}
		for i := range gToIDs {
			if gToIDs[i] != otherToIDs[i] {
				return false
			}
		}
	}
	return true
}

func (g *Graph) AddEdge(id1, id2 EdgeID, distance float64) error {
	if _, exist := g.FindDistance(id1, id2); exist {
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

func (g *Graph) FindDistance(fromID, toID EdgeID) (float64, bool) {
	for _, edge := range g.Adj[fromID] {
		if edge.To == toID {
			return edge.Distance, true
		}
	}
	return 0.0, false
}
