package graph

type IGraph interface {
	CalcTotalDistance([]EdgeID) (float64, bool)
	GetToEdgeIDSlice(EdgeID) []EdgeID
}
