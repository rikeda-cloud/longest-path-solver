package parser

import (
	"bufio"
	"fmt"
	"io"

	edge "github.com/rikeda-cloud/longest-path-solver/internal/model"
)

func ParseEdges(reader io.Reader) ([]*edge.Edge, error) {
	var edges []*edge.Edge

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		edge, err := edge.ParseGraphEdge(line)
		if err != nil {
			return nil, err
		}
		edges = append(edges, edge)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return edges, nil
}
