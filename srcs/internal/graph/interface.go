package graph

type IGraph interface {
	AddEdge(EdgeID, EdgeID, float64) error
	FindDistance(EdgeID, EdgeID) (float64, bool)
	GetFromEdgeIDSlice() []EdgeID
	GetToEdgeIDSlice(EdgeID) []EdgeID
}
