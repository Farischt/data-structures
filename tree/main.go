package tree

import (
	"ds/queue"
	"ds/stack"
	"fmt"
)

type ITree interface {
	Insert(data int)
	Remove(data int)
	Search(data int) *Node
	InOrderTraversal()
	DepthFirstValues(int) []int
	BreathFirstValues(int) []int
	FindMinimumValue() *Node
}

type BinarySearchTree struct {
	root *Node
}

func New(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: node,
	}
}

// Insert a new node in the tree, if the tree is empty the new node will be the root.
func (t *BinarySearchTree) Insert(data int) {
	newNode := NewNode(data)

	if t.root == nil {
		t.root = newNode
		return
	}

	var currentNode *Node = t.root

	for currentNode != nil {
		if data <= currentNode.data {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				return
			}
			currentNode = currentNode.right
		}
	}
}

// Remove a node from the tree and rearrange the tree based on the data.
func (t *BinarySearchTree) Remove(data int) {
	if t.root == nil {
		return
	}

	var currentNode *Node = t.root
	var parentNode *Node = nil

	// Find the node to delete
	for currentNode != nil && currentNode.data != data {
		parentNode = currentNode
		if currentNode.data > data {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	if currentNode == nil {
		// In this case the node to delete is not present in the tree
		return
	} else if currentNode.HasTwoChild() {
		// In this case the currentNode to delete has two child
		currentNode.RemoveNodeWithTwoChild()
		return
	} else if currentNode.left != nil {
		// In this case the currentNode to delete only has a left child
		currentNode.RemoveNodeWithOneChild(parentNode, t, Left)
		return
	} else if currentNode.right != nil {
		// In this case the currentNode to delete only has a right child
		currentNode.RemoveNodeWithOneChild(parentNode, t, Right)
		return
	} else {
		// In this case the currentNode to delete has no child
		// We need to cut the relation
		currentNode.RemoveNodeWithNoChild(parentNode, t)
		return
	}

}

// Search a node in the tree and return it.
func (t *BinarySearchTree) Search(data int) *Node {
	if t.root == nil || t.root.data == data {
		return t.root
	}

	currentNode := t.root

	if data <= currentNode.data {
		bst := New(currentNode.left)
		return bst.Search(data)
	}

	bst := New(currentNode.right)
	return bst.Search(data)
}

// InOrderTraversal traverse the tree in order and print the data.
func (t *BinarySearchTree) InOrderTraversal() {
	if t.root == nil {
		return
	}

	bst := New(t.root.left)
	bst.InOrderTraversal()

	fmt.Println(t.root.data)

	bst = New(t.root.right)
	bst.InOrderTraversal()
}

// DepthFirstValues traverse the tree in depth first order and return the values in a slice.
func (t *BinarySearchTree) DepthFirstValues(capacity int) []int {
	if t.root == nil {
		return []int{}
	}

	var tempStack stack.IStack[*Node]

	currentNode := t.root
	result := []int{}
	tempStack = stack.New[*Node](capacity)
	tempStack.Push(currentNode)

	for !tempStack.IsEmpty() {
		currentNode, _ = tempStack.Pop()
		result = append(result, currentNode.data)

		if currentNode.HasChild() {
			if currentNode.right != nil {
				tempStack.Push(currentNode.right)
			}

			if currentNode.left != nil {
				tempStack.Push(currentNode.left)
			}
		}
	}
	return result
}

// BreathFirstValues traverse the tree in breath first order and return the values in a slice.
func (t *BinarySearchTree) BreathFirstValues(capacity int) []int {
	if t.root == nil {
		return []int{}
	}

	var tempQueue queue.IQueue[*Node]

	currentNode := t.root
	result := []int{}
	tempQueue = queue.New[*Node](capacity)
	tempQueue.Enqueue(currentNode)

	for !tempQueue.IsEmpty() {
		dequeuedNode, _ := tempQueue.Dequeue()
		currentNode = *dequeuedNode
		result = append(result, currentNode.data)

		if currentNode.HasChild() {
			if currentNode.left != nil {
				tempQueue.Enqueue(currentNode.left)
			}

			if currentNode.right != nil {
				tempQueue.Enqueue(currentNode.right)
			}
		}
	}

	return result

}

// FindMinimumValue find the minimum value in the tree. Usefull for the remove method and inorder successor.
func (t *BinarySearchTree) FindMinimumValue() *Node {
	if t.root == nil {
		return nil
	}

	currentNode := t.root
	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode
}
