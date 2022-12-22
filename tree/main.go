package tree

import (
	"fmt"
)

type ITree interface {
	Insert(data int)
	Search(data int) *Node
	InOrderTraversal()
	DepthFirstValues() []int
}

type BinarySearchTree struct {
	root *Node
}

func New(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: node,
	}
}

func (t *BinarySearchTree) Insert(data int) {
	newNode := NewNode(data)

	if t.root == nil {
		t.root = newNode
		return
	}

	var currentNode *Node = t.root

	for currentNode != nil {
		if data < currentNode.data {
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

func (t *BinarySearchTree) Search(data int) *Node {
	if t.root == nil || t.root.data == data {
		return t.root
	}

	currentNode := t.root

	if data < currentNode.data {
		bst := New(currentNode.left)
		return bst.Search(data)
	}

	bst := New(currentNode.right)
	return bst.Search(data)
}

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

func (t *BinarySearchTree) DepthFirstValues() []int {
	if t.root == nil {
		return []int{}
	}

	var tempStack stacker

	currentNode := t.root
	result := []int{}
	tempStack = NewStack(6)
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
