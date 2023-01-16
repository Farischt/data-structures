package ll

type Node[T comparable] struct {
	data T
	next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
	}
}
