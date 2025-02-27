package algorithm

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

type PathNode struct {
	Path     []graph.EdgeID
	Node     graph.EdgeID
	Distance float64
}
