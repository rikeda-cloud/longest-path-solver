package input

import (
	"bufio"
	"fmt"
	"io"
)

func ParseGraphInputs(reader io.Reader) ([]*GraphInput, error) {
	var graphInputs []*GraphInput

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		graphInput, err := ParseGraphInput(line)
		if err != nil {
			return nil, err
		}
		graphInputs = append(graphInputs, graphInput)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return graphInputs, nil
}
