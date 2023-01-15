package graph

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	graph := New[int]()
	fmt.Println(graph.nodes)
	if graph.nodes == nil {
		t.Errorf("root is not nil, test new graph.")
	}
}

func Test_Add(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	if graph.nodes[1] == nil {
		t.Errorf("root is not 1, test add node.")
	}
}

func Test_AddDirectedEdge(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.AddDirectedEdge(1, 2)
	if graph.nodes[1][0] != 2 {
		t.Errorf("edge is not 2, test add directed edge.")
	}
}

func Test_AddUndirectedEdge(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.AddUndirectedEdge(1, 2)
	if graph.nodes[1][0] != 2 {
		t.Errorf("edge is not 2, test add undirected edge.")
	}
	if graph.nodes[2][0] != 1 {
		t.Errorf("edge is not 1, test add undirected edge.")
	}
}

func Test_HasPath(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.AddDirectedEdge(1, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, 5)
	if !graph.HasPath(1, 5) {
		t.Errorf("path is not 1->2->3->4->5, test has path.")
	}
}

func Test_HasPathWithInvalidSrc(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.AddDirectedEdge(1, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, 5)
	if graph.HasPath(0, 5) {
		t.Errorf("path is not 1->2->3->4->5, test has path.")
	}
}

func Test_HasPathWithInvalidDst(t *testing.T) {
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.Add(3)
	graph.Add(4)
	graph.Add(5)
	graph.AddDirectedEdge(1, 2)
	graph.AddDirectedEdge(2, 3)
	graph.AddDirectedEdge(3, 4)
	graph.AddDirectedEdge(4, 5)
	if graph.HasPath(1, 0) {
		t.Errorf("path is not 1->2->3->4->5, test has path.")
	}
}

func Test_HasPathWithSameSrcDst(t *testing.T) {
	graph := New[int]()
	graph.Add(1)

	if !graph.HasPath(1, 1) {
		t.Errorf("path is not 1->1, test has path.")
	}
}

func Test_ComponentCount(t *testing.T) {
	expectedCount := 1
	graph := New[int]()
	graph.Add(1)
	graph.Add(2)
	graph.AddUndirectedEdge(1, 2)

	if graph.ComponentCount() != expectedCount {
		t.Errorf("component count is not 1, test component count.")
	}
}

func Test_ComponentCountEmpty(t *testing.T) {
	expectedCount := 0
	graph := New[int]()

	if graph.ComponentCount() != expectedCount {
		t.Errorf("component count is not 0, test component count.")
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
		t.Errorf("component count is not 2, test component count.")
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
		t.Errorf("component count is not 1, test component count.")
	}

	if graph.LargestComponentSize() != expectedSize {
		t.Errorf("largest component size is not 5, test largest component size.")
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
		t.Errorf("component count is not 2, test component count.")
	}

	if graph.LargestComponentSize() != expectedSize {
		t.Errorf("largest component size is not 5, test largest component size.")
	}
}
