package stack

import "fmt"

type IStack[T any] interface {
	IsFull() bool
	IsEmpty() bool
	Push(T) error
	// Pop method should return a ptr of T
	Pop() (T, error)
	Size() int
	// Peek method should also return a ptr of T
	Peek() (T, error)
}

type Stack[T any] struct {
	storage  []T
	capacity int
}

func New[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		capacity: capacity,
	}
}

func (s *Stack[T]) String() string {
	stack := "\n------- Stack --------\n"
	for i := s.Size() - 1; i >= 0; i-- {
		stack = stack + fmt.Sprintf("%v\n--\n", s.storage[i])
	}
	stack = stack + "------- End --------\n"
	return stack
}

// IsFull returns either if the stack is full or not
func (s *Stack[T]) IsFull() bool {
	return len(s.storage) == s.capacity
}

// IsEmpty returns either if the stack is empty or not
func (s *Stack[T]) IsEmpty() bool {
	return len(s.storage) == 0
}

// Push adds a new element to the stack
func (s *Stack[T]) Push(data T) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	s.storage = append(s.storage, data)
	return nil
}

// Pop removes the last added element from the stack
func (s *Stack[T]) Pop() (T, error) {
	var nilValue T
	if s.IsEmpty() {
		return nilValue, fmt.Errorf("stack is empty")
	}

	lastEl := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	s.capacity--
	return lastEl, nil
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	return len(s.storage)
}

// Peek returns the last added element from the stack
func (s *Stack[T]) Peek() (T, error) {
	var nilValue T
	if s.IsEmpty() {
		return nilValue, fmt.Errorf("stack is empty")
	}

	return s.storage[len(s.storage)-1], nil
}
