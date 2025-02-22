package input

import (
	"bufio"
	"fmt"
	"io"
)

func ParseGraphInputs(reader io.Reader) ([]*GraphInput, error) {
	var edges []*GraphInput

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		edge, err := ParseGraphInput(line)
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
