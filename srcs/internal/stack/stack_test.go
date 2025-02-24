package stack

import (
	"testing"

	"github.com/rikeda-cloud/longest-path-solver/internal/graph"
)

func TestStackOperations(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}

	var node1 graph.EdgeID = 1
	var node2 graph.EdgeID = 2
	s.Push(Item{Node: node1, Path: []graph.EdgeID{1}})
	s.Push(Item{Node: node2, Path: []graph.EdgeID{1, 2}})

	if s.IsEmpty() {
		t.Errorf("Expected stack to have elements, but it is empty")
	}

	poppedItem, ok := s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an item, but it failed")
	}
	if poppedItem.Node != node2 {
		t.Errorf("Expected popped item to be {Node: 2}, got {Node: %d}", poppedItem.Node)
	}

	poppedItem, ok = s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an item, but it failed")
	}
	if poppedItem.Node != node1 {
		t.Errorf("Expected popped item to be {Node: 1}, got {Node: %d}", poppedItem.Node)
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements, but it is not")
	}

	poppedItem, ok = s.Pop()
	if ok {
		t.Errorf("Expected Pop on an empty stack to return false, but it returned true")
	}
}
