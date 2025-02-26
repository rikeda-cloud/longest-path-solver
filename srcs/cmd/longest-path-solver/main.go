package main

import (
	"log"
	"os"

	"github.com/rikeda-cloud/longest-path-solver/internal/algorithm"
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
	"github.com/rikeda-cloud/longest-path-solver/internal/input"
	"github.com/rikeda-cloud/longest-path-solver/internal/output"
)

func main() {
	graphInputs, err := input.ParseGraphInputs(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	graph, err := input.ConvertGraphInputsToGraph(graphInputs, graph.NewGraph())
	if err != nil {
		log.Fatal(err)
	}
	result := algorithm.FindLongestPathByDfsGoroutine(graph)
	output.PrintResult(result)
}
