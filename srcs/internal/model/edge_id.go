package edge

import (
	"errors"
)

type EdgeID uint

var ErrInvalidEdgeID = errors.New("EdgeID must be 1 or greater")

func NewEdgeID(number int) (EdgeID, error) {
	// EdgeID must be positive integers (1 or greater); 0 is not included.
	if number < 1 {
		return 0, ErrInvalidEdgeID
	}
	return EdgeID(number), nil
}
