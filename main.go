package main

import (
	"ds/graph"
	"fmt"
)

func main() {
	fmt.Println("Welcome to ds package !")

	var graph graph.IGraph[int] = graph.New[int]()

	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.Add(6)
	graph.Add(7)

	graph.AddUndirectedEdge(1, 2)
	graph.AddUndirectedEdge(2, 3)
	graph.AddUndirectedEdge(3, 4)
	graph.AddUndirectedEdge(4, 5)

	graph.AddUndirectedEdge(6, 7)

	fmt.Println(graph)
	fmt.Println(graph.ComponentCount())
	fmt.Println(graph.LargestComponentSize())
}
