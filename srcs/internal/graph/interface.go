package graph

type IGraph interface {
	CalcTotalDistance([]EdgeID) (float64, bool)
	GetFromEdgeIDSlice() []EdgeID
	GetToEdgeIDSlice(EdgeID) []EdgeID
	AddEdge(EdgeID, EdgeID, float64) error
	Equal(IGraph) bool
}
