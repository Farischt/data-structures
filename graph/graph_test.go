package graph

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	graph := New[int]()
	fmt.Println(graph.nodes)
	if graph.nodes == nil {
		t.Errorf("nodes should be initialized, test new graph.")
	}
}

func Test_Add(t *testing.T) {
	data := 1
	graph := New[int]()
	graph.Add(data)
	if graph.nodes[data] == nil {
		t.Errorf("graph should contain a node with data %d, test add node.", data)
	}
}

func Test_AddDirectedEdge(t *testing.T) {
	src := 1
	dst := 2
	graph := New[int]()
	graph.Add(src)
	graph.Add(dst)
	graph.AddDirectedEdge(src, dst)
	if graph.nodes[src][0] != dst {
		t.Errorf("graph should have a direct edge %d -> %d , test add directed edge.", src, dst)
	}
}

func Test_AddUndirectedEdge(t *testing.T) {
	src := 1
	dst := 2
	graph := New[int]()
	graph.Add(src)
	graph.Add(dst)
	graph.AddUndirectedEdge(src, dst)
	if graph.nodes[src][0] != dst {
		t.Errorf("graph should have a undirect edge %d -- %d, test add undirected edge.", src, dst)
	}
	if graph.nodes[dst][0] != src {
		t.Errorf("graph should have a undirect edge %d -- %d, test add undirected edge.", dst, src)
	}
}

func Test_HasPath(t *testing.T) {
	src := 1
	dst := 5
	graph := New[int]()
	graph.Add(src)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(dst)
	graph.AddDirectedEdge(src, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, dst)
	if !graph.HasPath(src, dst) {
		t.Errorf("graph should have a path from %d to %d, test has path.", src, dst)
	}
}

func Test_HasPathWithInvalidSrc(t *testing.T) {
	src := 0
	dst := 5
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(dst)
	graph.AddDirectedEdge(1, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, dst)
	if graph.HasPath(src, dst) {
		t.Errorf("graph shouldn't have a path from %d to %d since %d doesn't exist, test has path with invalid src.", src, dst, src)
	}
}

func Test_HasPathWithInvalidDst(t *testing.T) {
	src := 1
	dst := 0
	graph := New[int]()
	graph.Add(src)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.AddDirectedEdge(src, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, 5)
	if graph.HasPath(src, dst) {
		t.Errorf("graph shouldn't have a path from %d to %d since %d doesn't exist, test has path with invalid dst.", src, dst, dst)
	}
}

func Test_HasPathWithSameSrcDst(t *testing.T) {
	src := 1
	graph := New[int]()
	graph.Add(src)

	if !graph.HasPath(src, src) {
		t.Errorf("graph should have a path from %d to %d, test has path with same src dst.", src, src)
	}
}

func Test_ComponentCount(t *testing.T) {
	expectedCount := 1
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.AddUndirectedEdge(1, 2)

	if graph.ComponentCount() != expectedCount {
		t.Errorf("component count should be %d, test component count.", expectedCount)
	}
}

func Test_ComponentCountEmpty(t *testing.T) {
	expectedCount := 0
	graph := New[int]()

	if graph.ComponentCount() != expectedCount {
		t.Errorf("component count should be %d, test component count empty.", expectedCount)
	}
}

func Test_ComponentCountWithTwoComponents(t *testing.T) {
	expectedCount := 2
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)

	graph.AddUndirectedEdge(1, 2)
	graph.AddUndirectedEdge(3, 4)

	if graph.ComponentCount() != expectedCount {
		t.Errorf("component count should be %d, test component count with two components.", expectedCount)
	}
}

func Test_LargestComponentSizeWithOneComponent(t *testing.T) {
	expectedSize := 5
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.AddUndirectedEdge(1, 2)
	graph.AddUndirectedEdge(2, 3)
	graph.AddUndirectedEdge(3, 4)
	graph.AddUndirectedEdge(4, 5)

	if graph.ComponentCount() != 1 {
		t.Errorf("component count should be 1, test largest component size with one component.")
	}

	if graph.LargestComponentSize() != expectedSize {
		t.Errorf("largest component size should be %d, test largest component size with one component.", expectedSize)
	}
}

func Test_LargestComponentSizeWithTwoComponents(t *testing.T) {
	expectedSize := 5
	graph := New[int]()
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

	if graph.ComponentCount() != 2 {
		t.Errorf("component count should be 2, test largest component size with two components.")
	}

	if graph.LargestComponentSize() != expectedSize {
		t.Errorf("largest component size should be %d, test largest component size with two components.", expectedSize)
	}
}
