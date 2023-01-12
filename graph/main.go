package graph

import (
	"ds/queue"
	"ds/stack"
	"fmt"
)

type IGraph[T comparable] interface {
	nodeExists(node T) bool
	Add(node T) error
	AddDirectedEdge(src T, dst T) error
	AddUndirectedEdge(src T, dst T) error
	HasPathBFS(src T, dst T) bool
	HasPathDFS(src T, dst T) bool
	ComponentCount() int
	LargestComponentSize() int
	ShortestPath(src T, dst T) int
}

type Graph[T comparable] struct {
	nodes map[T][]T
}

func New[T comparable]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[T][]T),
	}
}

func (g *Graph[T]) nodeExists(node T) bool {
	_, exist := g.nodes[node]
	return exist
}

func (g *Graph[T]) Add(node T) error {
	// Check if node exist

	if g.nodeExists(node) {
		return fmt.Errorf("node already exist")
	}

	g.nodes[node] = make([]T, 0)
	return nil
}

func (g *Graph[T]) AddDirectedEdge(src T, dst T) error {

	if !g.nodeExists(src) || !g.nodeExists(dst) {
		return fmt.Errorf("at least one of the two nodes doesnt exist")
	}

	g.nodes[src] = append(g.nodes[src], dst)
	return nil
}

func (g *Graph[T]) AddUndirectedEdge(src T, dst T) error {

	err := g.AddDirectedEdge(src, dst)
	if err != nil {
		return err
	}

	g.nodes[dst] = append(g.nodes[dst], src)

	return err
}

func (g *Graph[T]) HasPathBFS(src T, dst T) bool {
	// BFS Way
	if !g.nodeExists(src) || !g.nodeExists(dst) {
		return false
	} else if src == dst {
		return true
	}

	visitedNodes := make(map[T]bool)
	var queue queue.IQueue[T] = queue.New[T](len(g.nodes))
	queue.Enqueue(src)

	for !queue.IsEmpty() {
		currentNode, _ := queue.Dequeue()
		visitedNodes[*currentNode] = true

		if *currentNode == dst {
			return true
		}

		for _, neighbor := range g.nodes[*currentNode] {
			// Enqueue all paths
			_, exist := visitedNodes[neighbor]
			if !exist {
				queue.Enqueue(neighbor)
			}
		}
	}

	return false
}

func (g *Graph[T]) HasPathDFS(src T, dst T) bool {
	// DFS Way
	if !g.nodeExists(src) || !g.nodeExists(dst) {
		return false
	} else if src == dst {
		return true
	}

	visitedNodes := make(map[T]bool)
	var stack stack.IStack[T] = stack.New[T](len(g.nodes))
	stack.Push(src)

	for !stack.IsEmpty() {
		currentNode, _ := stack.Pop()
		visitedNodes[currentNode] = true

		if currentNode == dst {
			return true
		}

		for _, neighbor := range g.nodes[currentNode] {
			// Enqueue all paths
			_, exist := visitedNodes[neighbor]
			if !exist {
				stack.Push(neighbor)
			}
		}
	}

	return false
}

func (g *Graph[T]) ComponentCount() int {
	visitedNodes := make(map[T]bool)
	count := 0

	for node := range g.nodes {
		_, exist := visitedNodes[node]
		if !exist {
			visitedNodes[node] = true

			var stack stack.IStack[T] = stack.New[T](len(g.nodes))
			stack.Push(node)

			for !stack.IsEmpty() {
				currentNode, _ := stack.Pop()
				visitedNodes[currentNode] = true

				for _, neighbor := range g.nodes[currentNode] {
					_, exist := visitedNodes[neighbor]
					if !exist {
						stack.Push(neighbor)
					}
				}
			}
			count++
		}
	}

	return count
}

func (g *Graph[T]) LargestComponentSize() int {
	visitedNodes := make(map[T]bool)
	maxSize := 0

	for node := range g.nodes {
		_, exist := visitedNodes[node]
		if !exist {
			visitedNodes[node] = true

			var stack stack.IStack[T] = stack.New[T](len(g.nodes))
			stack.Push(node)

			currentSize := 0
			for !stack.IsEmpty() {
				currentNode, _ := stack.Pop()
				currentSize++
				visitedNodes[currentNode] = true

				for _, neighbor := range g.nodes[currentNode] {
					_, exist := visitedNodes[neighbor]
					if !exist {
						stack.Push(neighbor)
					}
				}
			}

			if currentSize > maxSize {
				maxSize = currentSize
			}
		}
	}

	return maxSize
}

type Element[T any] struct {
	node     T
	distance int
}

func (g *Graph[T]) ShortestPath(src T, dst T) int {
	visitedNodes := make(map[T]bool)
	queue := queue.New[Element[T]](len(g.nodes))
	queue.Enqueue(Element[T]{
		node:     src,
		distance: 0,
	})

	for !queue.IsEmpty() {
		currentNode, _ := queue.Dequeue()

		if currentNode.node == dst {
			return currentNode.distance
		}

		visitedNodes[currentNode.node] = true

		for _, neighboor := range g.nodes[currentNode.node] {
			_, exist := visitedNodes[neighboor]
			if !exist {
				queue.Enqueue(Element[T]{
					node:     neighboor,
					distance: currentNode.distance + 1,
				})
			}
		}
	}

	return -1
}
