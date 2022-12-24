package queue

import "fmt"

type IQueue[T any] interface {
	IsEmpty() bool
	IsFull() bool
	Enqueue(T) error
	Dequeue() (*T, error)
	Size() int
}

type Queue[T any] struct {
	storage  []T
	capacity int
}

func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		capacity: capacity,
	}
}

func (q *Queue[T]) String() string {
	queue := "\n------- Queue --------\n"

	for _, el := range q.storage {
		queue += fmt.Sprintf("%v\n--\n", el)
	}

	return queue + "-------- End ----------\n"
}

// IsEmpty returns either if the queue is empty or not
func (q *Queue[T]) IsEmpty() bool {
	return len(q.storage) == 0
}

// IsFull returns either if the queue is full or not
func (q *Queue[T]) IsFull() bool {
	return len(q.storage) == q.capacity
}

// Enqueue adds a new element to the queue
func (q *Queue[T]) Enqueue(data T) error {

	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}

	q.storage = append(q.storage, data)
	return nil
}

// Dequeue removes the first added element from the queue
func (q *Queue[T]) Dequeue() (*T, error) {

	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	item := q.storage[0]
	q.storage = q.storage[1:len(q.storage)]
	q.capacity--

	return &item, nil
}

// Size returns the size of the queue
func (q *Queue[T]) Size() int {
	return len(q.storage)
}
