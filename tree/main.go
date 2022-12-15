package tree

type ITree interface {
	Insert(data int)
	Search(data int) bool
	Remove(data int)
}

type BinarySearchTree struct {
	root *Node 
}

func New() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
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


func (t *BinarySearchTree) Search(data int) bool {
	return false
}

func (t *BinarySearchTree) Remove(data interface{}) {
}


	

