package stack

import "fmt"

type IStack[T any] interface {
	// IsFull returns either if the stack is full or not.
	// Time complexity: O(1).
	// Space complexity: O(1).
	IsFull() bool
	// IsEmpty returns either if the stack is empty or not.
	// Time complexity: O(1).
	// Space complexity: O(1).
	IsEmpty() bool
	// Push adds a new element to the top of the stack.
	// Returns an error if the stack is full.
	// Time complexity: O(1).
	// Space complexity: O(1).
	Push(T) error
	// Pop removes the top element from the stack.
	// Returns an error if the stack is empty.
	// Time complexity: O(1).
	// Space complexity: O(1).
	Pop() (*T, error)
	// Size returns the size of the stack.
	// Time complexity: O(1).
	// Space complexity: O(1).
	Size() int
	// Peek returns the top element from the stack.
	// Returns an error if the stack is empty.
	// Time complexity: O(1).
	// Space complexity: O(1).
	Peek() (*T, error)
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

// IsFull returns either if the stack is full or not.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) IsFull() bool {
	return len(s.storage) == s.capacity
}

// IsEmpty returns either if the stack is empty or not.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) IsEmpty() bool {
	return len(s.storage) == 0
}

// Push adds a new element to the top of the stack.
// Returns an error if the stack is full.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) Push(data T) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	s.storage = append(s.storage, data)
	return nil
}

// Size returns the size of the stack.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) Pop() (*T, error) {
	var nilValue *T
	if s.IsEmpty() {
		return nilValue, fmt.Errorf("stack is empty")
	}

	lastEl := s.storage[len(s.storage)-1]
	s.storage = s.storage[0 : len(s.storage)-1]
	s.capacity--
	return &lastEl, nil
}

// Size returns the size of the stack.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) Size() int {
	return len(s.storage)
}

// Peek returns the top element from the stack.
// Returns an error if the stack is empty.
// Time complexity: O(1).
// Space complexity: O(1).
func (s *Stack[T]) Peek() (*T, error) {
	var nilValue *T
	if s.IsEmpty() {
		return nilValue, fmt.Errorf("stack is empty")
	}

	return &s.storage[len(s.storage)-1], nil
}
