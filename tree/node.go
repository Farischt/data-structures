package tree

type Node struct {
	data   int
	left   *Node
	right  *Node
	parent *Node
}

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
)

func NewNode(data int) *Node {
	return &Node{data: data}
}

// HasChild returns true if the node has at least one child.
// This is useful to know if the node is a leaf or not.
// A leaf is a node that has no children.
func (n *Node) HasChild() bool {
	return n.left != nil || n.right != nil
}

// HasTwoChild returns true if the node has two children.
func (n *Node) HasTwoChild() bool {
	return n.left != nil && n.right != nil
}

// IsFromRight returns true if the node is the right child of the parent.
func (n *Node) IsFromRight() bool {

	if n.parent == nil { 
		return false
	}

	return n == n.parent.right
}

// RemoveNodeWithTwoChild removes a node with two children.
func (currentNode *Node) RemoveNodeWithTwoChild() {
	// In this case we need to find the inorder successor of the currentNode
	var temp *Node = currentNode.right
	var tempParent *Node = nil

	for temp.left != nil {
		tempParent = temp
		temp = temp.left
	}

	// Once we found it we need to place it as the right successor of the parent
	if tempParent != nil {
		tempParent.left = temp.right
	} else {
		currentNode.right = temp.right
	}
	// Here i found the minimum value
	currentNode.data = temp.data
}

// RemoveNodeWithOneChild removes a node with one child.
func (currentNode *Node) RemoveNodeWithOneChild(t *BinarySearchTree, direction Direction) {
	var newNode *Node = nil
	switch direction {
	case Left:
		newNode = currentNode.left
	case Right:
		newNode = currentNode.right
	}

	if currentNode == t.root {
		t.root = newNode
	} else if currentNode.IsFromRight() {
		currentNode.parent.right = newNode
	} else {
		currentNode.parent.left = newNode
	}
}

// RemoveNodeWithNoChild removes a node with no child. This is the case when the node is a leaf. In this case we just need to remove the reference to the node from the parent.
func (currentNode *Node) RemoveLeafNode(t *BinarySearchTree) {
	if currentNode == t.root {
		t.root = nil
	} else if currentNode.IsFromRight() {
		currentNode.parent.right = nil
	} else {
		currentNode.parent.left = nil
	}
}
