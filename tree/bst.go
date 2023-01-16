package tree

import (
	"fmt"

	"github.com/farischt/ds/queue"
	"github.com/farischt/ds/stack"
)

type IBinarySearchTree interface {
	// Insert a new node in the tree, if the tree is empty the new node will be the root.
	// Time complexity: O(log(n)).
	// Space complexity: O(1).
	Insert(data int) *Node
	// Remove a node from the tree.
	// Time complexity: O(log(n)).
	// Space complexity: O(1).
	Remove(data int)
	// Search a node in the tree and return it.
	// Time complexity: O(log(n)).
	// Space complexity: O(1).
	Search(data int) *Node
	// InOrderTraversal returns the values of the tree in order.
	// Time complexity: O(n).
	// Space complexity: O(n).
	InOrderTraversal()
	// DepthFirstValues returns the values of the tree in depth first order.
	// Time complexity: O(n).
	// Space complexity: O(n).
	DepthFirstValues(int) []int
	// BreathFirstValues returns the values of the tree in breath first order.
	// Time complexity: O(n).
	// Space complexity: O(n).
	BreathFirstValues(int) []int
	// FindMinimumValue returns the minimum value of the tree.
	// Time complexity: O(n).
	// Space complexity: O(1).
	FindMinimumValue() *Node
	// BreathFirstSearch returns the node with the given data.
	// Time complexity: O(n).
	// Space complexity: O(n).
	BreathFirstSearch(data int, capacity int) *Node
	// FindInorderSuccessor returns the inorder successor of the given node.
	// The inorder successor is the node that comes after the given node in an in-order traversal.
	// Time complexity: O(n).
	// Space complexity: O(1).
	FindInorderSuccessor(data int) *Node
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
func (t *BinarySearchTree) Insert(data int) *Node {
	newNode := NewNode(data)

	if t.root == nil {
		newNode.parent = nil
		t.root = newNode
		return newNode
	}

	var currentNode *Node = t.root

	for currentNode != nil {
		if data <= currentNode.data {
			if currentNode.left == nil {
				newNode.parent = currentNode
				currentNode.left = newNode
				return newNode
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				newNode.parent = currentNode
				currentNode.right = newNode
				return newNode
			}
			currentNode = currentNode.right
		}
	}

	return nil
}

// Remove a node from the tree and rearrange the tree based on the data.
func (t *BinarySearchTree) Remove(data int) {
	if t.root == nil {
		return
	}

	var currentNode *Node = t.root
	
	// Find the node to delete
	for currentNode != nil && currentNode.data != data {
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
		currentNode.RemoveNodeWithOneChild(t, Left)
		return
	} else if currentNode.right != nil {
		// In this case the currentNode to delete only has a right child
		currentNode.RemoveNodeWithOneChild(t, Right)
		return
	} else {
		// In this case the currentNode to delete has no child
		// We need to cut the relation
		currentNode.RemoveLeafNode(t)
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
		node, _ := tempStack.Pop()
		currentNode = *node
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

func (t *BinarySearchTree) BreathFirstSearch(data int, capacity int) *Node {
	if t.root == nil {
		return nil
	}

	var tempQueue queue.IQueue[*Node]
	var currentNode *Node

	tempQueue = queue.New[*Node](capacity)
	currentNode = t.root
	tempQueue.Enqueue(currentNode)

	for !tempQueue.IsEmpty() {
		dequeuedNode, _ := tempQueue.Dequeue()
		currentNode = *dequeuedNode

		if currentNode.data == data {
			return currentNode
		}

		if currentNode.left != nil {
			tempQueue.Enqueue(currentNode.left)
		}

		if currentNode.right != nil {
			tempQueue.Enqueue(currentNode.right)
		}
	}
	return nil
}

func (t *BinarySearchTree) FindInorderSuccessor(data int) *Node {

	if t.root == nil {
		return nil
	}
	// Get to the node we are looking for
	currentNode := t.Search(data)

	if currentNode == nil {
		return nil
	}
	// At this point we found the Node we were looking for
	// First thing to check is if i have right child. Meaning i will be going for the right child

	if currentNode.right != nil {
		// In the case we need to go one step right and all the way left until it is nill
		currentNode = currentNode.right

		if currentNode.left != nil {
			for currentNode.left != nil {
				currentNode = currentNode.left
			}
		}

		return currentNode
	} else {

		for currentNode.parent != nil && currentNode.IsFromRight() {
			if currentNode.parent.parent == nil {
				return nil
			}
			currentNode = currentNode.parent
		}

		return currentNode.parent
	}
}
