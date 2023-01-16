package tree

import (
	"testing"
)

func TestNode_New(t *testing.T) {
	node := NewNode(1)
	if node.data != 1 {
		t.Errorf("data is not 1")
	}
}

func TestNode_HasChild(t *testing.T) {
	node := NewNode(1)
	if node.HasChild() {
		t.Errorf("node has no child")
	}
	node.left = NewNode(2)
	if !node.HasChild() {
		t.Errorf("node has left child")
	}
	node.left = nil
}

func TestNode_HasTwoChild(t *testing.T) {
	node := NewNode(1)
	if node.HasTwoChild() {
		t.Errorf("node has no child")
	}
	node.left = NewNode(2)
	if node.HasTwoChild() {
		t.Errorf("node has only left child")
	}
	node.right = NewNode(3)
	if !node.HasTwoChild() {
		t.Errorf("node has two child")
	}
	node.left = nil
}

func TestNode_IsFromRight(t *testing.T) {
	node := NewNode(1)
	childNode := NewNode(2)
	
	if childNode.IsFromRight() {
		t.Errorf("node is not from right")
	}
	
	node.right = childNode
	childNode.parent = node
	if !childNode.IsFromRight() {
		t.Errorf("node is from right")
	}
}

func Test_RemoveRootNodeWithTwoChild(t *testing.T) {
	node := NewNode(1)
	node.left = NewNode(2)
	node.right = NewNode(3)
	node.RemoveNodeWithTwoChild()
	if node.data != 3 {
		t.Errorf("node is not 3")
	}
	if node.left.data != 2 {
		t.Errorf("left child is not 2")
	}
}

func Test_RemoveRootNodeWithLeftChild(t *testing.T) {
	tree := New(nil)
	node := NewNode(10)
	node.left = NewNode(5)
	tree.root = node

	node.RemoveNodeWithOneChild(tree, Left)
	if tree.root.data != 5 {
		t.Errorf("Root node is not 5. Delete root node with left child")
	}
}

func Test_RemoveRootNodeWithRightChild(t *testing.T) {
	tree := New(nil)
	node := NewNode(10)
	node.right = NewNode(15)
	tree.root = node

	node.RemoveNodeWithOneChild(tree, Right)
	if tree.root.data != 15 {
		t.Errorf("Root node is not 15. Delete root node with right child")
	}
}

func Test_RemoveRootNodeWithNoChild(t *testing.T) {
	tree := New(nil)
	node := NewNode(10)
	tree.root = node

	node.RemoveLeafNode(tree)
	if tree.root != nil {
		t.Errorf("Root node is not nil. Delete root node with no child")
	}
}

func Test_RemoveMiddleNodeWithLeftChild(t *testing.T) {
	tree := New(nil)
	parentNode := NewNode(10)
	currentNode := NewNode(5)
	currentNode.parent = parentNode
	parentNode.left = currentNode
	currentNodeLeftChild := NewNode(3)
	currentNode.left = currentNodeLeftChild
	tree.root = parentNode

	currentNode.RemoveNodeWithOneChild(tree, Left)
	if parentNode.left.data != 3 {
		t.Errorf("Parent child node is not 3. Delete middle node with left child")
	}
}

func Test_RemoveMiddleNodeWithRightChild(t *testing.T) {
	tree := New(nil)
	parentNode := NewNode(10)
	currentNode := NewNode(15)
	currentNode.parent = parentNode
	parentNode.right = currentNode
	currentNodeRightChild := NewNode(25)
	currentNode.left = currentNodeRightChild
	tree.root = parentNode

	currentNode.RemoveNodeWithOneChild(tree, Left)
	if parentNode.right.data != 25 {
		t.Errorf("Parent child node is not 3. Delete middle node with right child")
	}
}

func Test_RemoveLeafNode(t *testing.T) {
	tree := New(nil)
	parentNode := NewNode(10)
	currentNode := NewNode(5)
	parentNode.left = currentNode
	currentNode.parent = parentNode
	tree.root = parentNode

	currentNode.RemoveLeafNode(tree)
	if parentNode.left != nil {
		t.Errorf("Parent child node is not nil. Delete leaf node")
	}
}
