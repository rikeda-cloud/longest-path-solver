package graph

import (
	"errors"
)

type MapBasedGraph struct {
	Adj map[EdgeID]map[EdgeID]float64
}

func NewMapBasedGraph() *MapBasedGraph {
	return &MapBasedGraph{Adj: make(map[EdgeID]map[EdgeID]float64)}
}

func (g *MapBasedGraph) Equal(other *MapBasedGraph) bool {
	// グラフのエッジ数を比較
	if len(g.Adj) != len(other.Adj) {
		return false
	}

	for key, edges := range g.Adj {
		otherEdges, exists := other.Adj[key]
		// 他のグラフに同じキーと同じ数のエッジがあるか確認
		if !exists || len(edges) != len(otherEdges) {
			return false
		}

		// 各エッジを比較
		for toID, distance := range edges {
			if otherDistance, found := otherEdges[toID]; !found || distance != otherDistance {
				return false
			}
		}
	}
	return true
}

func (g *MapBasedGraph) AddEdge(id1, id2 EdgeID, distance float64) error {
	if _, exist := g.findEdge(id1, id2); exist {
		return errors.New("edge already exists")
	}

	if g.Adj[id1] == nil {
		g.Adj[id1] = make(map[EdgeID]float64)
	}
	g.Adj[id1][id2] = distance

	if g.Adj[id2] == nil {
		g.Adj[id2] = make(map[EdgeID]float64)
	}
	g.Adj[id2][id1] = distance

	return nil
}

func (g *MapBasedGraph) GetFromEdgeIDSlice() []EdgeID {
	fromEdgeIDSlice := make([]EdgeID, 0, len(g.Adj))
	for fromEdgeID := range g.Adj {
		fromEdgeIDSlice = append(fromEdgeIDSlice, fromEdgeID)
	}
	return fromEdgeIDSlice
}

func (g *MapBasedGraph) GetToEdgeIDSlice(fromEdgeID EdgeID) []EdgeID {
	toEdgeIDSlice := make([]EdgeID, 0, len(g.Adj[fromEdgeID]))
	for key := range g.Adj[fromEdgeID] {
		toEdgeIDSlice = append(toEdgeIDSlice, key)
	}
	return toEdgeIDSlice
}

func (g *MapBasedGraph) CalcTotalDistance(path []EdgeID) (float64, bool) {
	totalDistance := 0.0
	for i := 0; i < len(path)-1; i++ {
		fromID := path[i]
		toID := path[i+1]
		distance, ok := g.findEdge(fromID, toID)
		if !ok {
			return 0.0, false
		}
		totalDistance += distance
	}
	return totalDistance, true
}

func (g *MapBasedGraph) findEdge(fromEdgeID, toEdgeID EdgeID) (float64, bool) {
	if edges, exists := g.Adj[fromEdgeID]; exists {
		if distance, found := edges[toEdgeID]; found {
			return distance, true
		}
	}
	return 0.0, false
}
