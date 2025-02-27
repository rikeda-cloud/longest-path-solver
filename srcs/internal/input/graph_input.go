package input

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

type GraphInput struct {
	Start    graph.EdgeID
	End      graph.EdgeID
	Distance float64
}

// INFO float validation is handled later with ParseFloat for accuracy.
const inputPatternStr = `^[ ]*([+]?\d+)[ ]*,[ ]*([+]?\d+)[ ]*,[ ]*([+-]?[\d.]+)[ ]*$`

var inputPattern = regexp.MustCompile(inputPatternStr)

func ParseGraphInput(line string) (*GraphInput, error) {
	matches := inputPattern.FindStringSubmatch(line)
	// matches[0] is the full match.
	// matches[1:4] are the captured groups: start, end, and distance.
	if matches == nil || len(matches) != 4 {
		return nil, fmt.Errorf("invalid line format: %s", line)
	}

	start, errStart := strconv.Atoi(matches[1])
	end, errEnd := strconv.Atoi(matches[2])
	distance, errDistance := strconv.ParseFloat(matches[3], 64)

	if errStart != nil || errEnd != nil || errDistance != nil {
		return nil, fmt.Errorf("invalid data in line: %s", line)
	}

	startEdgeID, errStartID := graph.NewEdgeID(start)
	endEdgeID, errEndID := graph.NewEdgeID(end)
	if errStartID != nil || errEndID != nil {
		return nil, fmt.Errorf("start and end must be positive integers: %s", line)
	}

	if startEdgeID == endEdgeID {
		return nil, fmt.Errorf("Start and End cannot be the same in line: %s", line)
	}

	return &GraphInput{Start: startEdgeID, End: endEdgeID, Distance: distance}, nil
}
