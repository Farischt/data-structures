package tree

import "testing"

func Test_New(t *testing.T) {
	tree := New(nil)
	if tree.root != nil {
		t.Errorf("root is not nil, test new tree.")
	}
}

func Test_Insert(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	if tree.root.data != 1 {
		t.Errorf("root is not 1, test insert root node.")
	}
	tree.Insert(2)
	if tree.root.right.data != 2 {
		t.Errorf("right child is not 2, test insert right node.")
	}
	tree.Insert(0)
	if tree.root.left.data != 0 {
		t.Errorf("left child is not 0, test insert left node.")
	}
}

func Test_RemoveNonExistingNode(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Remove(2)
	if tree.root.data != 1 {
		t.Errorf("root is not 1, test remove non existing node.")
	}
}

func Test_RemoveNodeWithTwoChildren(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(0)
	tree.Remove(1)

	if tree.root.data != 2 {
		t.Errorf("root is not 0, test remove node with two children.")
	}
}

func Test_RemoveNodeWithLeftChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(2)
	tree.Insert(1)
	tree.Remove(2)

	if tree.root.data != 1 {
		t.Errorf("root is not 0, test remove node with left child.")
	}
}

func Test_RemoveNodeWithRightChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Insert(2)
	tree.Remove(1)

	if tree.root.data != 2 {
		t.Errorf("root is not 0, test remove node with right child.")
	}
}

func Test_RemoveNodeWithNoChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Remove(1)

	if tree.root != nil {
		t.Errorf("root is not nil, test remove node with no child.")
	}
}

func Test_Search(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(0)
	if tree.Search(1) == nil {
		t.Errorf("node is not nil, test search node.")
	}
	if tree.Search(0) == nil {
		t.Errorf("node is not nil, test search node.")
	}
	if tree.Search(3) != nil {
		t.Errorf("node is nil, test search non existing node.")
	}
}

func Test_DepthFirstValues(t *testing.T) {
	tree := New(nil)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(9)
	tree.Insert(8)

	values := tree.DepthFirstValues(4)
	expectedResult := []int{10, 9, 8, 15}

	for i := 0; i < len(values); i++ {
		if values[i] != expectedResult[i] {
			t.Errorf("values are not equal, test depth first values.")
		}
	}
}

func Test_EmptyDepthFirstValues(t *testing.T) {
	tree := New(nil)
	values := tree.DepthFirstValues(4)
	if len(values) != 0 {
		t.Errorf("values are not empty, test empty depth first values.")
	}
}

func Test_BreathFirstValues(t *testing.T) {
	tree := New(nil)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(9)
	tree.Insert(8)
	values := tree.BreathFirstValues(3)
	expectedResult := []int{10, 9, 15, 8}

	for i := 0; i < len(values); i++ {
		if values[i] != expectedResult[i] {
			t.Errorf("values are not equal, test breath first values.")
		}
	}
}

func Test_EmptyBreathFirstValues(t *testing.T) {
	tree := New(nil)
	values := tree.BreathFirstValues(3)
	if len(values) != 0 {
		t.Errorf("values are not empty, test empty breath first values.")
	}
}

func Test_FindMinimumValue(t *testing.T) {
	tree := New(nil)
	node := NewNode(8)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(9)
	tree.Insert(node.data)
	min := tree.FindMinimumValue()
	if min.data != node.data {
		t.Errorf("invalid min value, test find minimum value.")
	}
}
