package output

import (
	"fmt"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func PrintResult(result []graph.EdgeID) {
	for _, id := range result {
		fmt.Print(id, "\r\n")
	}
}
