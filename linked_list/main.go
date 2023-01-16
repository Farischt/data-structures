package ll

import (
	"fmt"
)

type ILinkedList[T comparable] interface {
	// isEmpty returns either if the linked list is empty or not.
	// Time Complexity: O(1).
	// Space Complexity: O(1).
	isEmpty() bool
	// InsertAtStart insert a new node based on the data parameter, at the begining of the linked list.
	// It returns a pointer to the inserted Node.
	// Time Complexity: O(1).
	// Space Complexity: O(1).
	InsertAtStart(T) *Node[T]
	// InsertAtEnd insert a new node based on the data parameter, at the end of the linked list.
	// It returns a pointer to the inserted Node.
	// Time Complexity: O(n).
	// Space Complexity: O(1).
	InsertAtEnd(T) *Node[T]
	// DeleteNode deletes a node from the linked list based on the data parameter.
	// Time Complexity: O(n).
	// Space Complexity: O(1).
	DeleteNode(T)
	// Contains returns either if the linked list contains the data parameter or not.
	// Time Complexity: O(n).
	// Space Complexity: O(1).
	Contains(T) bool
	// toArray transform the linked list to an array of Nodes.
	// Time Complexity: O(n).
	// Space Complexity: O(n).
	toArray() []Node[T]
	// Size returns the size of the linked list.
	// Time Complexity: O(n).
	// Space Complexity: O(n).
	Size() int
	// GetLast returns the last node of the linked list.
	// Time Complexity: O(n).
	// Space Complexity: O(1).
	GetLast(*Node[T]) *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	}
}

func (ll *LinkedList[T]) String() string {
	if ll.isEmpty() {
		return "nil"
	}
	var str string
	currentNode := ll.head
	for currentNode != nil {
		str += fmt.Sprintf("%v -> ", currentNode.data)
		currentNode = currentNode.next
	}
	str += "nil"
	return str
}

// isEmpty returns either if the linked list is empty or not.
// Time Complexity: O(1).
// Space Complexity: O(1).
func (ll *LinkedList[T]) isEmpty() bool {
	return ll.head == nil
}

// InsertAtStart insert a new node based on the data parameter, at the begining of the linked list.
// It returns a pointer to the inserted Node.
// Time Complexity: O(1).
// Space Complexity: O(1).
func (ll *LinkedList[T]) InsertAtStart(data T) *Node[T] {
	newNode := NewNode(data)
	// Check if the list is empty
	if ll.isEmpty() {
		ll.head = newNode
		return newNode
	}

	newNode.next = ll.head
	ll.head = newNode
	return ll.head
}

// InsertAtEnd insert a new node based on the data parameter, at the end of the linked list.
// It returns a pointer to the inserted Node.
// Time Complexity: O(n).
// Space Complexity: O(1).
func (ll *LinkedList[T]) InsertAtEnd(data T) *Node[T] {
	newNode := NewNode(data)

	if ll.isEmpty() {
		ll.head = newNode
		return newNode
	}

	lastNode := GetLast(ll.head)
	lastNode.next = newNode
	return newNode
}

// DeleteNode deletes a node from the linked list based on the data parameter.
// Time Complexity: O(n).
// Space Complexity: O(1).
func (ll *LinkedList[T]) DeleteNode(data T) {
	if ll.isEmpty() {
		return
	}

	var previousNode *Node[T] = nil
	currentNode := ll.head

	for currentNode != nil {
		if currentNode == ll.head && currentNode.data == data {
			ll.head = currentNode.next
		} else if currentNode.data == data {
			previousNode.next = currentNode.next
		} else {
			previousNode = currentNode
		}
		currentNode = currentNode.next
	}

}

// Contains returns either if the linked list contains the data parameter or not.
// Time Complexity: O(n).
// Space Complexity: O(1).
func (ll *LinkedList[T]) Contains(data interface{}) bool {
	currentNode := ll.head
	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// toArray transform the linked list to an array of Nodes.
// Time Complexity: O(n).
// Space Complexity: O(n).
func (ll *LinkedList[T]) toArray() []Node[T] {

	currentNode := ll.head
	var array []Node[T]

	for currentNode != nil {
		array = append(array, *currentNode)
		currentNode = currentNode.next
	}

	return array
}

// Size returns the size of the linked list.
// Time Complexity: O(n).
// Space Complexity: O(n).
func (ll *LinkedList[T]) Size() int {
	return len(ll.toArray())
}

// GetLast returns the last node of the linked list.
// Time Complexity: O(n).
// Space Complexity: O(1).
func GetLast[T comparable](node *Node[T]) *Node[T] {
	if node.next == nil {
		return node
	}
	nextNode := GetLast(node.next)
	return nextNode
}
