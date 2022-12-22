package stack

import "fmt"

type IStack interface {
	IsFull() bool
	IsEmpty() bool
	Push(interface{}) error
	Pop() (interface{}, error)
	Size() int
	Peek() interface{}
}

type Stack struct {
	storage  []interface{}
	capacity int
}

func New(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
	}
}

func (s *Stack) String() string {
	stack := "\n------- Stack --------\n"
	for i := s.Size() - 1; i >= 0; i-- {
		stack = stack + fmt.Sprintf("%v\n--\n", s.storage[i])
	}
	stack = stack + "------- End --------\n"
	return stack
}

// IsFull returns either if the stack is full or not
func (s *Stack) IsFull() bool {
	return len(s.storage) == s.capacity
}

// IsEmpty returns either if the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.storage) == 0
}

// Push adds a new element to the stack
func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	s.storage = append(s.storage, data)
	return nil
}

// Pop removes the last added element from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	lastEl := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	s.capacity--
	return lastEl, nil
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.storage)
}

// Peek returns the last added element from the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.storage[len(s.storage)-1]
}
