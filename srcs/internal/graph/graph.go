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
	if len(g.GetFromEdgeIDs()) != len(other.GetFromEdgeIDs()) {
		return false
	}

	for _, fromID := range g.GetFromEdgeIDs() {
		gToIDs := g.GetToEdgeIDs(fromID)
		otherToIDs := other.GetToEdgeIDs(fromID)

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
		return errors.New("Duplicate edges")
	}

	// Adds an edge between id1 and id2 with the given distance,
	// ensuring both directions are stored for the undirected graph.
	g.Adj[id1] = append(g.Adj[id1], Edge{To: id2, Distance: distance})
	g.Adj[id2] = append(g.Adj[id2], Edge{To: id1, Distance: distance})

	return nil
}

func (g *Graph) FindDistance(fromID, toID EdgeID) (float64, bool) {
	for _, edge := range g.Adj[fromID] {
		if edge.To == toID {
			return edge.Distance, true
		}
	}
	return 0.0, false
}

func (g *Graph) GetFromEdgeIDs() []EdgeID {
	fromEdgeIDs := make([]EdgeID, 0, len(g.Adj))
	for fromID := range g.Adj {
		fromEdgeIDs = append(fromEdgeIDs, fromID)
	}
	return fromEdgeIDs
}

func (g *Graph) GetToEdgeIDs(fromID EdgeID) []EdgeID {
	edges := g.Adj[fromID]
	toEdgeIDs := make([]EdgeID, 0, len(edges))
	for _, edge := range edges {
		toEdgeIDs = append(toEdgeIDs, edge.To)
	}
	return toEdgeIDs
}
