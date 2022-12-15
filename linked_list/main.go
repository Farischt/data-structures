package ll

import (
	"fmt"
)

type ILinkedList interface {
	GetLast(*Node) *Node
	isEmpty() bool
	InsertAtStart(interface{}) *Node
	InsertAtEnd(interface{}) *Node
	DeleteNode(interface{})
	Contains(interface{})
	toArray() []Node
	Size() int
}

type LinkedList struct {
	head *Node
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (ll *LinkedList) String() string {
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

// GetLast returns the last node of the linked list.
func GetLast(node *Node) *Node {
	if node.next == nil {
		return node
	}
	nextNode := GetLast(node.next)
	return nextNode
}

// IsEmpty returns either if the linked list is empty or not.
func (ll *LinkedList) isEmpty() bool {
	return ll.head == nil
}

// InsertAtStart insert a new node based on the data parameter, at the begining of the linked list.
// It returns a pointer to the inserted Node
func (ll *LinkedList) InsertAtStart(data interface{}) *Node {
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
// It returns a pointer to the inserted Node
func (ll *LinkedList) InsertAtEnd(data interface{}) *Node {
	newNode := NewNode(data)

	if ll.isEmpty() {
		ll.head = newNode
		return newNode
	}

	lastNode := GetLast(ll.head)
	lastNode.next = newNode
	return newNode
}

// In order to delete a node we need to :
// First find that node
// If we don't find that node we return an error
// Else we remove the node
// To remove the node we need to :
// take the previous node and set it next node to the current node next node
func (ll *LinkedList) DeleteNode(data interface{}) {
	if ll.isEmpty() {
		return
	}

	var previousNode *Node = nil
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

func (ll *LinkedList) Contains(data interface{}) bool {
	currentNode := ll.head
	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// ToArray transform the linked list to an array of Nodes.
func (ll *LinkedList) toArray() []Node {

	currentNode := ll.head
	var array []Node

	for currentNode != nil {
		array = append(array, *currentNode)
		currentNode = currentNode.next
	}

	return array
}

// Size returns the size of the linked list
func (ll *LinkedList) Size() int {
	return len(ll.toArray())
}
