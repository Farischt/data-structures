package stack

import "fmt"

type IStack interface {
	IsFull() bool
	IsEmpty() bool
	Push(interface{}) error
	Pop() (*interface{}, error)
	Size() int
	Peek() *interface{}
}

type Stack struct {
	Storage  []interface{}
	Capacity int
}

func New(capacity int) *Stack {
	return &Stack{
		Capacity: capacity,
	}
}

func (s *Stack) String() string {
	stack := "------- Stack --------\n"
	for i := s.Size() - 1; i >= 0; i-- {
		stack = stack + fmt.Sprintf("%d\n--\n", s.Storage[i])
	}
	stack = stack + "------- End --------"
	return stack
}

// IsFull returns either if the stack is full or not
func (s *Stack) IsFull() bool {
	return len(s.Storage) == s.Capacity
}

// IsEmpty returns either if the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.Storage) == 0
}

// Push adds a new element to the stack
func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	s.Storage = append(s.Storage, data)
	return nil
}

// Pop removes the last added element from the stack
func (s *Stack) Pop() (*interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	lastEl := s.Storage[s.Capacity-1]
	s.Storage = s.Storage[:len(s.Storage)-1]

	s.Capacity--
	return &lastEl, nil
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.Storage)
}

// Peek returns the last added element from the stack
func (s *Stack) Peek() *interface{} {
	if s.IsEmpty() {
		return nil
	}

	return &s.Storage[s.Capacity-1]
}
