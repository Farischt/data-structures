package node

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}
