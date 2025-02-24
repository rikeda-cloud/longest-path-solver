package output

import (
	"fmt"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func PrintResult(result []graph.EdgeID) {
	if len(result) <= 1 {
		return
	}

	for _, id := range result {
		fmt.Print(id, "\r\n")
	}
}
