package queue

import "testing"

func TestNew(t *testing.T) {
	q := New[int](1)
	if q == nil {
		t.Errorf("should create new queue, test new queue")
	}
}

func TestString(t *testing.T) {
	q := New[int](1)
	q.Enqueue(10)
	if q.String() != "\n------- Queue --------\n10\n--\n-------- End ----------\n" {
		t.Errorf("should return correct string, test string")
	}
}

func TestIsEmpy(t *testing.T) {
	q := New[int](1)
	if !q.IsEmpty() {
		t.Errorf("should be empty, test queue is empty")
	}
}

func TestIsFull(t *testing.T) {
	q := New[int](1)
	q.Enqueue(10)
	if !q.IsFull() {
		t.Errorf("should be full, test queue is full")
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int](2)
	q.Enqueue(10)
	q.Enqueue(20)

	if q.storage[0] != 10 {
		t.Errorf("should enqueue 10, test enqueue")
	}
}

func TestEnqueueFullQueue(t *testing.T) {
	q := New[int](1)
	q.Enqueue(10)
	size := q.Size()
	q.Enqueue(20)

	if q.storage[1] != 20 {
		t.Errorf("should enqueue 20, test enqueue full queue")
	}

	if q.Size() != size+1 {
		t.Errorf("should increase the size of the queue, test enqueue full queue")
	}
}

func TestDequeue(t *testing.T) {
	expectedDequeuedValue := 10
	q := New[int](2)
	q.Enqueue(expectedDequeuedValue)
	q.Enqueue(20)
	value, err := q.Dequeue()
	if err != nil {
		t.Error("should dequeue without error, test dequeue")
	} else if *value != expectedDequeuedValue {
		t.Errorf("should dequeue correctly, test dequeue ,expected %d, got %d", expectedDequeuedValue, *value)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := New[int](1)
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("should return error, test dequeue empty queue")
	}
}

func TestSize(t *testing.T) {
	q := New[int](2)
	q.Enqueue(10)
	q.Enqueue(20)
	if q.Size() != 2 {
		t.Errorf("should have size 2, got %d", q.Size())
	}
}
