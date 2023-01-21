package heap

import (
	"fmt"

	"github.com/farischt/ds/utils"
)

type HeapType string

const (
	MinHeap HeapType = "min"
	MaxHeap HeapType = "max"
)

type IHeap[T utils.Number] interface {
	// Returns the size of the heap.
	size() int
	// Returns true if the heap is empty.
	isEmpty() bool
	// Less returns true if element at index i is lower or equal to element at index j.
	less(i int, j int) bool
	// Greater returns true if element at index i is greater or equal to element at index j.
	greater(i int, j int) bool
	// Compare two elements in the heap based on the heap type.
	// True means the heap property is satisfied.
	// For example, if the heap type is min heap, then the compare function
	// returns true if i <= j.
	// If the heap type is max heap, then the compare function returns true
	// if i >= j.
	compare(i int, j int) bool
	// Swap two elements in the heap.
	swap(i int, j int)
	// Returns the parent index of the element at index i.
	parent(i int) int
	// Returns the left child index of the element at index i.
	left(i int) int
	// Returns the right child index of the element at index i.
	right(i int) bool
	// Returns true if the element at index i is in the heap.
	isInBound(i int) bool
	// Up heapify the element at index i.
	up(i int)
	// Down heapify the element at index i.
	down(i int)
	// Pushes an element into the heap.
	// The element is inserted at the end of the heap and then up heapified.
	// The time complexity is O(log(n)).
	// The space complexity is O(1).
	Push(v T)
	// Returns the top element of the heap and removes it.
	// The top element is the element that satisfies the heap property.
	// Returns an error if the heap is empty.
	// The time complexity is O(log(n)).
	// The space complexity is O(1).
	Pop() (*T, error)
	// Returns the top element of the heap without removing it.
	// The top element is the element that satisfies the heap property.
	// Returns an error if the heap is empty.
	// The time complexity is O(1).
	// The space complexity is O(1).
	Top() (*T, error)
}

type Heap[T utils.Number] struct {
	data []Item[T]
	t    HeapType
}

func New[T utils.Number](t HeapType) *Heap[T] {
	return &Heap[T]{
		data: make([]Item[T], 0),
		t:    t,
	}
}

// Returns the size of the heap.
func (h *Heap[T]) size() int {
	return len(h.data)
}

// Returns true if the heap is empty.
func (h *Heap[T]) isEmpty() bool {
	return h.size() == 0
}

// Less returns true if element at index i is lower or equal to element at index j.
func (h *Heap[T]) less(i int, j int) bool {
	return h.data[i].Value <= h.data[j].Value
}

// Greater returns true if element at index i is greater or equal to element at index j.
func (h *Heap[T]) greater(i int, j int) bool {
	return h.data[i].Value >= h.data[j].Value
}

// Compare two elements in the heap based on the heap type.
// True means the heap property is satisfied.
// For example, if the heap type is min heap, then the compare function
// returns true if i <= j.
// If the heap type is max heap, then the compare function returns true
// if i >= j.
func (h *Heap[T]) compare(i int, j int) bool {

	if h.t == MaxHeap {
		return h.greater(i, j)
	}

	return h.less(i, j)
}

// Swap two elements in the heap.
func (h *Heap[T]) swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Returns the parent index of the element at index i.
func (h *Heap[T]) parent(i int) int {
	return (i - 1) / 2
}

// Returns the left child index of the element at index i.
func (h *Heap[T]) left(i int) int {
	return 2*i + 1
}

// Returns the right child index of the element at index i.
func (h *Heap[T]) right(i int) int {
	return 2*i + 2
}

// Returns true if the element at index i is in the heap.
func (h *Heap[T]) isInBound(i int) bool {
	return i >= 0 && i <= h.size()-1
}

// Up heapify the element at index i.
func (h *Heap[T]) up(i int) {
	parentIndex := h.parent(i)
	for !h.compare(parentIndex, i) {
		h.swap(i, parentIndex)
		i = parentIndex
		parentIndex = h.parent(i)
	}
}

// Down heapify the element at index i.
func (h *Heap[T]) down(i int) {
	leftIndex := h.left(i)
	isLeftValid := h.isInBound(leftIndex)

	for isLeftValid {
		smallerChildIndex := leftIndex
		rightIndex := h.right(i)
		isRightValid := h.isInBound(rightIndex)

		if isRightValid && h.compare(rightIndex, leftIndex) {
			smallerChildIndex = rightIndex
		}

		if !h.compare(i, smallerChildIndex) {
			h.swap(i, smallerChildIndex)
			i = smallerChildIndex
			leftIndex = h.left(i)
			isLeftValid = h.isInBound(leftIndex)
		} else {
			break
		}
	}
}

// Pushes an element into the heap.
// The element is inserted at the end of the heap and then up heapified.
// The time complexity is O(log(n)).
// The space complexity is O(1).
func (h *Heap[T]) Push(v T, info interface{}) {
	// Insert at end
	h.data = append(h.data, *NewItem(v, info))
	currentIndex := h.size() - 1
	// Heapify up
	h.up(currentIndex)
}

// Returns the top element of the heap and removes it.
// The top element is the element that satisfies the heap property.
// Returns an error if the heap is empty.
// The time complexity is O(log(n)).
// The space complexity is O(1).
func (h *Heap[T]) Pop() (*Item[T], error) {

	var lastElement *Item[T]

	if h.isEmpty() {
		return lastElement, fmt.Errorf("heap is empty")
	}

	if !h.isEmpty() {
		h.swap(0, h.size()-1)
		// Remove the last element
		lastElement = &h.data[h.size()-1]
		h.data = h.data[:h.size()-1]

		h.down(0)
	}

	return lastElement, nil
}

// Returns the top element of the heap without removing it.
// The top element is the element that satisfies the heap property.
// Returns an error if the heap is empty.
// The time complexity is O(1).
// The space complexity is O(1).
func (h *Heap[T]) Top() (*Item[T], error) {
	var topElement *Item[T]

	if h.isEmpty() {
		return topElement, fmt.Errorf("heap is empty")
	}

	if !h.isEmpty() {
		topElement = &h.data[0]
	}

	return topElement, nil
}
