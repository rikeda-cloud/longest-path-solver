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
	graphInuts, err := input.ParseGraphInputs(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	graph, err := input.ConvertGraphInputsToGraph(graphInuts, graph.NewGraph())
	if err != nil {
		log.Fatal(err)
	}
	result := algorithm.FindLongestPathByDfs(graph)
	output.PrintResult(result)
}
