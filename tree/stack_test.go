package tree

import "testing"

func TestIsFull(t *testing.T) {
	s := NewStack(2)
	s.Push(NewNode(20))
	s.Push(NewNode(20))
	if !s.IsFull() {
		t.Errorf("Expected true, got false")
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack(1)
	if !s.IsEmpty() {
		t.Errorf("Expected true, got false")
	}
}

func TestPush(t *testing.T) {
	expectedLastElement := NewNode(10)
	s := NewStack(2)

	s.Push(NewNode(20))
	s.Push(expectedLastElement)

	lastElement := s.Peek()

	if lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	}
}

func TestPushFullStack(t *testing.T) {
	s := NewStack(1)
	s.Push(NewNode(20))
	err := s.Push(NewNode(20))

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPop(t *testing.T) {
	expectedLastElement := NewNode(10)
	initialCapacity := 2
	expectedCapacity := initialCapacity - 1
	s := NewStack(initialCapacity)
	s.Push(NewNode(15))
	s.Push(expectedLastElement)

	lastElement, _ := s.Pop()

	if lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	} else if s.capacity != expectedCapacity {
		t.Errorf("Expected capacity %d, got %d", expectedCapacity, s.Size())
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStack(1)
	_, err := s.Pop()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSize(t *testing.T) {
	expectedSize := 2
	s := NewStack(expectedSize)
	s.Push(NewNode(10))
	s.Push(NewNode(20))

	if s.Size() != expectedSize {
		t.Errorf("Expected %d, got %d", expectedSize, s.Size())
	}
}

func TestPeek(t *testing.T) {
	expectedLastElement := NewNode(10)
	s := NewStack(1)
	s.Push(expectedLastElement)

	lastElement := s.Peek()

	if lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := NewStack(1)
	lastElement := s.Peek()

	if lastElement != nil {
		t.Errorf("Expected nil, got %v", lastElement)
	}
}
