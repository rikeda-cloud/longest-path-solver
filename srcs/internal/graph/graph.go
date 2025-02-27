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
	if len(g.Adj) != len(other.Adj) {
		return false
	}

	for fromID, edges := range g.Adj {
		otherEdges, ok := other.Adj[fromID]

		if !ok || len(edges) != len(otherEdges) {
			return false
		}

		for i, edge := range edges {
			if edge.To != otherEdges[i].To || edge.Distance != otherEdges[i].Distance {
				return false
			}
		}
	}
	return true
}

func (g *Graph) AddEdge(fromID, toID EdgeID, distance float64) error {
	if _, exist := g.FindDistance(fromID, toID); exist {
		return errors.New("Duplicate edges")
	}

	g.Adj[fromID] = append(g.Adj[fromID], Edge{To: toID, Distance: distance})

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
