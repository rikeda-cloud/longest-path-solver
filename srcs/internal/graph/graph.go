package graph

import (
	"errors"
)

type Edge struct {
	To       EdgeID
	Distance float64
}

type Graph struct {
	adj map[EdgeID][]Edge // adj = "Adjacency List"
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[EdgeID][]Edge)}
}

func newGraphWithEdges(edges map[EdgeID][]Edge) *Graph {
	adj := make(map[EdgeID][]Edge)
	for key, edges := range edges {
		adj[key] = edges
	}
	return &Graph{adj: adj}
}

func (g *Graph) AddEdge(id1, id2 EdgeID, distance float64) error {
	if g.edgeExists(id1, id2) {
		return errors.New("edge already exists")
	}

	// Adds an edge between id1 and id2 with the given distance,
	// ensuring both directions are stored for the undirected graph.
	g.adj[id1] = append(g.adj[id1], Edge{To: id2, Distance: distance})
	g.adj[id2] = append(g.adj[id2], Edge{To: id1, Distance: distance})

	return nil
}

func (g *Graph) edgeExists(id1, id2 EdgeID) bool {
	for _, edge := range g.adj[id1] {
		if edge.To == id2 {
			return true
		}
	}
	return false
}
