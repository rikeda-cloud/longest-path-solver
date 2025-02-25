package stack

import (
	"testing"
)

func TestStackOperations(t *testing.T) {
	s := NewStack[int]()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}

	item1 := 1
	item2 := 2
	item3 := 3
	s.Push(item1)
	s.Push(item2)

	if s.IsEmpty() {
		t.Errorf("Expected stack to have elements, but it is empty")
	}

	poppedItem, ok := s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an item, but it failed")
	}
	if poppedItem != item2 {
		t.Errorf("Expected popped item to be %d, got %d", item2, poppedItem)
	}

	s.Push(item3)

	poppedItem, ok = s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an item, but it failed")
	}
	if poppedItem != item3 {
		t.Errorf("Expected popped item to be %d, %d", item3, poppedItem)
	}

	poppedItem, ok = s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an item, but it failed")
	}
	if poppedItem != item1 {
		t.Errorf("Expected popped item to be %d, %d", item1, poppedItem)
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements, but it is not")
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("Expected Pop on an empty stack to return false, but it returned true")
	}
}
