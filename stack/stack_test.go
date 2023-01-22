package stack

import "testing"

func TestIsFull(t *testing.T) {
	s := New[int64](2)
	s.Push(10)
	s.Push(20)
	if !s.IsFull() {
		t.Errorf("Expected true, got false")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int64](1)
	if !s.IsEmpty() {
		t.Errorf("Expected true, got false")
	}
}

func TestPush(t *testing.T) {
	expectedLastElement := 10
	s := New[int](2)

	s.Push(20)
	s.Push(expectedLastElement)

	lastElement, _ := s.Peek()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	}
}

func TestPushFullStack(t *testing.T) {
	s := New[int64](1)
	s.Push(10)
	size := s.Size()
	s.Push(20)

	if s.Size() != size+1 {
		t.Errorf("should increase the size of the stack")
	}
}

func TestPop(t *testing.T) {
	expectedLastElement := 10
	initialCapacity := 2
	expectedCapacity := initialCapacity - 1
	s := New[int](initialCapacity)
	s.Push(15)
	s.Push(expectedLastElement)

	lastElement, _ := s.Pop()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	} else if s.capacity != expectedCapacity {
		t.Errorf("Expected capacity %d, got %d", expectedCapacity, s.Size())
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := New[int64](1)
	_, err := s.Pop()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSize(t *testing.T) {
	expectedSize := 2
	s := New[int](expectedSize)
	s.Push(10)
	s.Push(20)

	if s.Size() != expectedSize {
		t.Errorf("Expected %d, got %d", expectedSize, s.Size())
	}
}

func TestPeek(t *testing.T) {
	expectedLastElement := 10
	s := New[int](1)
	s.Push(expectedLastElement)

	lastElement, _ := s.Peek()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, lastElement)
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := New[int](1)
	_, err := s.Peek()

	if err == nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
