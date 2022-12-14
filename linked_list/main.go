package ll

import (
	"ds/node"
	"fmt"
)

type ILinkedList interface {
	GetLast(*node.Node) *node.Node
	isEmpty() bool
	InsertAtStart(interface{}) *node.Node
	InsertAtEnd(interface{}) *node.Node
	DeleteNode(interface{})
	Contains(interface{})
	toArray() []node.Node
	Size() int
}

type LinkedList struct {
	Head *node.Node
}

func New() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (ll *LinkedList) String() string {
	if ll.isEmpty() {
		return "nil"
	}
	var str string
	currentNode := ll.Head
	for currentNode != nil {
		str += fmt.Sprintf("%v -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	str += "nil"
	return str
}

// GetLast returns the last node of the linked list.
func GetLast(node *node.Node) *node.Node {
	if node.Next == nil {
		return node
	}
	nextNode := GetLast(node.Next)
	return nextNode
}

// IsEmpty returns either if the linked list is empty or not.
func (ll *LinkedList) isEmpty() bool {
	return ll.Head == nil
}

// InsertAtStart insert a new node based on the data parameter, at the begining of the linked list.
// It returns a pointer to the inserted Node
func (ll *LinkedList) InsertAtStart(data interface{}) *node.Node {
	newNode := node.NewNode(data)
	// Check if the list is empty
	if ll.isEmpty() {
		ll.Head = newNode
		return newNode
	}

	newNode.Next = ll.Head
	ll.Head = newNode
	return ll.Head
}

// InsertAtEnd insert a new node based on the data parameter, at the end of the linked list.
// It returns a pointer to the inserted Node
func (ll *LinkedList) InsertAtEnd(data interface{}) *node.Node {
	newNode := node.NewNode(data)

	if ll.isEmpty() {
		ll.Head = newNode
		return newNode
	}

	lastNode := GetLast(ll.Head)
	lastNode.Next = newNode
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

	var previousNode *node.Node = nil
	currentNode := ll.Head

	for currentNode != nil {
		if currentNode == ll.Head && currentNode.Data == data {
			ll.Head = currentNode.Next
		} else if currentNode.Data == data {
			previousNode.Next = currentNode.Next
		} else {
			previousNode = currentNode
		}
		currentNode = currentNode.Next
	}

}

func (ll *LinkedList) Contains(data interface{}) bool {
	currentNode := ll.Head
	for currentNode != nil {
		if currentNode.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

// ToArray transform the linked list to an array of Nodes.
func (ll *LinkedList) toArray() []node.Node {

	currentNode := ll.Head
	var array []node.Node

	for currentNode != nil {
		array = append(array, *currentNode)
		currentNode = currentNode.Next
	}

	return array
}

// Size returns the size of the linked list
func (ll *LinkedList) Size() int {
	return len(ll.toArray())
}
