package ll

type Node struct {
	data interface{}
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}
