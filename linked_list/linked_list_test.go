package ll

import "testing"

// Test the linked list
func TestGetLast(t *testing.T) {
	lastData := 300
	ll := New()
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(lastData)

	last := GetLast(ll.head)
	if last.data != lastData {
		t.Errorf("Expected %d, got %v", lastData, last.data)
	}
}

func TestIsEmpty(t *testing.T) {
	ll := New()
	if !ll.isEmpty() {
		t.Errorf("Expected true, got false")
	}
}

func TestInsertAtStart(t *testing.T) {
	firstData := 200
	ll := New()
	ll.InsertAtStart(firstData)
	if ll.head.data != firstData {
		t.Errorf("Expected %d, got %v", firstData, ll.head.data)
	}
}

func TestInsertAtEnd(t *testing.T) {
	lastData := 300
	ll := New()
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(lastData)
	last := GetLast(ll.head)
	if last.data != lastData {
		t.Errorf("Expected %d, got %v", lastData, last.data)
	}
}

func TestDeleteFirstNode(t *testing.T) {
	ll := New()
	firstNodeData := 200
	secondNodeData := 300
	ll.InsertAtEnd(firstNodeData)
	ll.InsertAtEnd(secondNodeData)
	ll.InsertAtEnd(400)
	ll.InsertAtEnd(500)

	ll.DeleteNode(firstNodeData)
	if ll.head.data != secondNodeData {
		t.Errorf("Expected %d, got %v", secondNodeData, ll.head.data)
	}
}

func TestDeleteLastNode(t *testing.T) {
	ll := New()
	previousNodeData := 400
	lastNodeData := 500
	ll.InsertAtEnd(100)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(300)
	ll.InsertAtEnd(previousNodeData)
	ll.InsertAtEnd(lastNodeData)

	ll.DeleteNode(lastNodeData)
	last := GetLast(ll.head)
	if last.data != previousNodeData {
		t.Errorf("Expected %d, got %v", previousNodeData, last.data)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	ll := New()
	middleNodeData := 400

	ll.InsertAtEnd(100)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(300)
	ll.InsertAtEnd(middleNodeData)
	ll.InsertAtEnd(500)
	ll.InsertAtEnd(600)

	ll.DeleteNode(middleNodeData)

	if ll.Contains(middleNodeData) {
		t.Errorf("Expected false, got true")
	}
}

func TestDeleteNode(t *testing.T) {
	ll := New()
	dataToDelete := 200
	ll.InsertAtEnd(dataToDelete)
	ll.InsertAtEnd(300)
	ll.InsertAtEnd(dataToDelete)
	ll.InsertAtEnd(dataToDelete)
	ll.InsertAtEnd(dataToDelete)

	ll.DeleteNode(dataToDelete)

	if ll.Contains(dataToDelete) {
		t.Errorf("Expected false, got true")
	}
}

func TestContains(t *testing.T) {
	ll := New()
	firstNodeData := 200
	secondNodeData := 300
	ll.InsertAtEnd(firstNodeData)
	ll.InsertAtEnd(secondNodeData)

	if !ll.Contains(firstNodeData) {
		t.Errorf("Expected true, got false")
	}
}

func TestToArray(t *testing.T) {
	ll := New()
	nodesData := []int{200, 300}
	ll.InsertAtEnd(nodesData[0])
	ll.InsertAtEnd(nodesData[1])

	array := ll.toArray()

	for i, node := range array {
		if node.data != nodesData[i] {
			t.Errorf("Expected %d, got %v", nodesData[i], node.data)
		}
	}
}

func TestSize(t *testing.T) {
	ll := New()
	expectedSize := 4
	for i := 0; i < expectedSize; i++ {
		ll.InsertAtEnd(i * 100)
	}

	if ll.Size() != expectedSize {
		t.Errorf("Expected %d, got %d", expectedSize, ll.Size())
	}
}
