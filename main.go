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

	graph.AddUndirectedEdge(1, 2)
	graph.AddUndirectedEdge(2, 3)
	graph.AddUndirectedEdge(3, 5)
	graph.AddUndirectedEdge(1, 4)
	graph.AddUndirectedEdge(4, 5)

	fmt.Println(graph)
	fmt.Println(graph.ComponentCount())
	fmt.Println(graph.HasPathBFS(3, 1))
	fmt.Println(graph.LargestComponentSize())
	fmt.Println(graph.ShortestPath(3, 1))
}
