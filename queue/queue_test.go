package queue

import "testing"

func TestIsEmpy(t *testing.T) {
	q := New[int](1)
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
}

func TestIsFull(t *testing.T) {
	q := New[int](1)
	q.Enqueue(10)
	if !q.IsFull() {
		t.Errorf("Expected queue to be full")
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int](2)
	q.Enqueue(10)
	q.Enqueue(20)
	if !q.IsFull() {
		t.Errorf("Expected queue to be full")
	}
}

func TestEnqueueFullQueue(t *testing.T) {
	q := New[int](1)
	q.Enqueue(10)
	size := q.Size()
	q.Enqueue(20)
	if q.Size() != size+1 {
		t.Errorf("Should increase the size of the queue")
	}
}

func TestDequeue(t *testing.T) {
	expectedDequeuedValue := 10
	q := New[int](2)
	q.Enqueue(expectedDequeuedValue)
	q.Enqueue(20)
	value, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if *value != expectedDequeuedValue {
		t.Errorf("Expected %d, got %d", expectedDequeuedValue, *value)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := New[int](1)
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected dequeue to return error, got nil")
	}
}

func TestSize(t *testing.T) {
	q := New[int](2)
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}
