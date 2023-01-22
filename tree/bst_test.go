package tree

import "testing"

func TestNew(t *testing.T) {
	tree := New(nil)
	if tree.root != nil {
		t.Errorf("root is not nil, test new tree.")
	}
}

func TestInsert(t *testing.T) {
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

func TestRemoveNonExistingNode(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Remove(2)
	if tree.root.data != 1 {
		t.Errorf("root is not 1, test remove non existing node.")
	}
}

func TestRemoveNodeWithTwoChildren(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(0)
	tree.Remove(1)

	if tree.root.data != 2 {
		t.Errorf("root is not 0, test remove node with two children.")
	}
}

func TestRemoveNodeWithLeftChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(2)
	tree.Insert(1)
	tree.Remove(2)

	if tree.root.data != 1 {
		t.Errorf("root is not 0, test remove node with left child.")
	}
}

func TestRemoveNodeWithRightChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Insert(2)
	tree.Remove(1)

	if tree.root.data != 2 {
		t.Errorf("root is not 0, test remove node with right child.")
	}
}

func TestRemoveNodeWithNoChild(t *testing.T) {
	tree := New(nil)
	tree.Insert(1)
	tree.Remove(1)

	if tree.root != nil {
		t.Errorf("root is not nil, test remove node with no child.")
	}
}

func TestSearch(t *testing.T) {
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

func TestDepthFirstValues(t *testing.T) {
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

func TestEmptyDepthFirstValues(t *testing.T) {
	tree := New(nil)
	values := tree.DepthFirstValues(4)
	if len(values) != 0 {
		t.Errorf("values are not empty, test empty depth first values.")
	}
}

func TestBreathFirstValues(t *testing.T) {
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

func TestEmptyBreathFirstValues(t *testing.T) {
	tree := New(nil)
	values := tree.BreathFirstValues(3)
	if len(values) != 0 {
		t.Errorf("values are not empty, test empty breath first values.")
	}
}

func TestFindMinimumValue(t *testing.T) {
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

func TestBreathFirstSearch(t *testing.T) {
	tree := New(nil)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(9)
	tree.Insert(8)
	node := tree.BreathFirstSearch(8, 4)
	secondNode := tree.BreathFirstSearch(50, 4)
	if node.data != 8 {
		t.Errorf("invalid node, test breath first search.")
	}
	if secondNode != nil {
		t.Errorf("node is not nil, test breath first search.")
	}
}

func TestEmptyBreathFirstSearch(t *testing.T) {
	tree := New(nil)
	node := tree.BreathFirstSearch(8, 4)
	if node != nil {
		t.Errorf("node is not nil, test empty breath first search.")
	}
}

func TestBreathFirstSearchRoot(t *testing.T) {
	tree := New(nil)
	tree.Insert(10)
	node := tree.BreathFirstSearch(10, 4)
	if node.data != 10 {
		t.Errorf("invalid node, test breath first search root.")
	}
}

func TestFindInorderSuccessor(t *testing.T) {
	tree := New(nil)
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(9)
	tree.Insert(5)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(14)

	node := tree.FindInorderSuccessor(9)
	if node.data != 11 {
		t.Errorf("invalid node, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(14)
	if node.data != 20 {
		t.Errorf("invalid node, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(25)
	if node != nil {
		t.Errorf("node is not nil, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(50)
	if node != nil {
		t.Errorf("node is not nil, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(5)
	if node.data != 9 {
		t.Errorf("invalid node, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(12)
	if node.data != 14 {
		t.Errorf("invalid node, test find inorder successor.")
	}

	node = tree.FindInorderSuccessor(20)
	if node.data != 25 {
		t.Errorf("invalid node, test find inorder successor.")
	}
}
