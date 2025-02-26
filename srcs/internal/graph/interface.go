package graph

type IGraph interface {
	GetFromEdgeIDSlice() []EdgeID
	GetToEdgeIDSlice(EdgeID) []EdgeID
	AddEdge(EdgeID, EdgeID, float64) error
	FindDistance(EdgeID, EdgeID) (float64, bool)
}
