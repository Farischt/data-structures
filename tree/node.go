package tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) HasChild() bool {
	return n.left != nil || n.right != nil
}
