package graph

type IGraph interface {
	CalcTotalDistance([]EdgeID) (float64, bool)
	GetFromEdgeIDSlice() []EdgeID
	GetToEdgeIDSlice(EdgeID) []EdgeID
}
