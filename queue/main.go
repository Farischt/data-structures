package queue

import "fmt"

type IQueue interface {
	IsEmpty() bool
	IsFull() bool
	Enqueue(interface{}) error
	Dequeue() (*interface{}, error)
	Size() int
}

type Queue struct {
	storage  []interface{}
	capacity int
}

func New(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
	}
}

func (q *Queue) String() string {
	queue := "\n------- Queue --------\n"

	for _, el := range q.storage {
		queue += fmt.Sprintf("%v\n--\n", el)
	}

	return queue + "-------- End ----------\n"
}

// IsEmpty returns either if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.storage) == 0
}

// IsFull returns either if the queue is full or not
func (q *Queue) IsFull() bool {
	return len(q.storage) == q.capacity
}

// Enqueue adds a new element to the queue
func (q *Queue) Enqueue(data interface{}) error {

	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}

	q.storage = append(q.storage, data)
	return nil
}

// Dequeue removes the first added element from the queue
func (q *Queue) Dequeue() (*interface{}, error) {

	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	item := q.storage[0]
	q.storage = q.storage[1:len(q.storage)]
	q.capacity--

	return &item, nil
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return len(q.storage)
}
