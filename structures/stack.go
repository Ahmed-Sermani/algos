package structures

type Stack[T any] struct {
	data []T
	top  int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: -1}
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
	s.top++
}

func (s *Stack[T]) Pop() T {
	if s.top == -1 {
		panic("stack underflow")
	}
	val := s.data[s.top]
	s.top--
	return val
}
