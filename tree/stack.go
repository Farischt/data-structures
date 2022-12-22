package tree

import "fmt"

type stacker interface {
	IsFull() bool
	IsEmpty() bool
	Push(*Node) error
	Pop() (*Node, error)
	Size() int
	Peek() *Node
}

type stack struct {
	storage  []*Node
	capacity int
}

func NewStack(capacity int) *stack {
	return &stack{
		capacity: capacity,
	}
}

func (s *stack) String() string {
	stack := "\n------- Stack --------\n"
	for i := s.Size() - 1; i >= 0; i-- {
		stack = stack + fmt.Sprintf("%v\n--\n", s.storage[i])
	}
	stack = stack + "------- End --------\n"
	return stack
}

// IsFull returns either if the stack is full or not
func (s *stack) IsFull() bool {
	return len(s.storage) == s.capacity
}

// IsEmpty returns either if the stack is empty or not
func (s *stack) IsEmpty() bool {
	return len(s.storage) == 0
}

// Push adds a new element to the stack
func (s *stack) Push(data *Node) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	s.storage = append(s.storage, data)
	return nil
}

// Pop removes the last added element from the stack
func (s *stack) Pop() (*Node, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	lastEl := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	s.capacity--
	return lastEl, nil
}

// Size returns the size of the stack
func (s *stack) Size() int {
	return len(s.storage)
}

// Peek returns the last added element from the stack
func (s *stack) Peek() *Node {
	if s.IsEmpty() {
		return nil
	}

	return s.storage[len(s.storage)-1]
}
