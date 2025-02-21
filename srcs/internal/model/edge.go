package edge

import (
	"fmt"
	"regexp"
	"strconv"
)

type Edge struct {
	Start    uint
	End      uint
	Distance float64
}

// INFO float validation is handled later with ParseFloat for accuracy.
const edgePatternStr = `^[ ]*(\d+)[ ]*,[ ]*(\d+)[ ]*,[ ]*([\d.]+)[ ]*$`

var edgePattern = regexp.MustCompile(edgePatternStr)

func ParseGraphEdge(line string) (*Edge, error) {
	matches := edgePattern.FindStringSubmatch(line)
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
	// start and end must be positive integers (1 or greater); 0 is not included.
	if start <= 0 || end <= 0 {
		return nil, fmt.Errorf("start and end must be positive integers: %s", line)
	}

	return &Edge{Start: uint(start), End: uint(end), Distance: distance}, nil
}
