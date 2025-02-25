package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zeroValue T
	dataSize := len(s.data)
	if dataSize == 0 {
		return zeroValue, false
	}
	item := s.data[dataSize-1]
	s.data = s.data[:dataSize-1]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
