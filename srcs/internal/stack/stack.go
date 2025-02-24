package stack

import (
	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

type Stack struct {
	data []Item
}

type Item struct {
	Node graph.EdgeID
	Path []graph.EdgeID
}

func NewStack() *Stack {
	return &Stack{data: []Item{}}
}

func (s *Stack) Push(item Item) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (Item, bool) {
	dataSize := len(s.data)
	if dataSize == 0 {
		return Item{}, false
	}
	item := s.data[dataSize-1]
	s.data = s.data[:dataSize-1]
	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
