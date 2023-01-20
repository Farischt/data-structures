package heap

import (
	"testing"
)

func Test_NewHeap(t *testing.T) {
	heap := New[int](MinHeap)
	if heap == nil {
		t.Error("should create new heap")
	}
}

func Test_HeapSize(t *testing.T) {
	heap := New[int](MinHeap)
	if heap.size() != 0 {
		t.Error("should return correct size")
	}
}

func Test_IsEmpty(t *testing.T) {
	heap := New[int](MinHeap)
	if !heap.isEmpty() {
		t.Error("should be empty")
	}
}

func Test_HeapLess(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.less(0, 1) {
		t.Error("should be less")
	}
}

func Test_HeapGreater(t *testing.T) {
	heap := New[int](MaxHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if heap.greater(0, 1) {
		t.Error("should be greater")
	}
}

func Test_HeapCompare(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.compare(0, 1) {
		t.Error("should compare correctly for min heap")
	}

	maxHeap := New[int](MaxHeap)
	maxHeap.data = append(heap.data, *NewItem(0, nil))
	maxHeap.data = append(heap.data, *NewItem(1, nil))

	if maxHeap.compare(0, 1) {
		t.Error("should compare correctly for max heap")
	}
}

func Test_HeapSwap(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	heap.swap(0, 1)

	if heap.data[0] != *secondItem {
		t.Error("should swap correctly")
	} else if heap.data[1] != *firstItem {
		t.Error("should swap correctly")
	}
}

func Test_HeapParent(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	parentIndex := heap.parent(1)

	if parentIndex != 0 {
		t.Error("should get parent index correctly")
	}
}

func Test_HeapLeft(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	leftIndex := heap.left(0)

	if leftIndex != 1 {
		t.Error("should get left index correctly")
	}
}

func Test_HeapRight(t *testing.T) {
	heap := New[int](MinHeap)

	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	rightIndex := heap.right(0)

	if rightIndex != 2 {
		t.Error("should get right index correctly")
	}
}

func Test_IsInBound(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.isInBound(1) {
		t.Error("should return true for valid child")
	}
}

func Test_HeapBubbleUp(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	heap.up(1)

	// if heap.data[0] != *secondItem {
	// 	t.Error("should bubble up correctly")
	// } else if heap.data[1] != *firstItem {
	// 	t.Error("should bubble up correctly")
	// }
}

func Test_HeapPush(t *testing.T) {
	heap := New[int](MinHeap)
	item := NewItem(1, nil)
	heap.Push(item.Value, item.Information)
	if heap.data[0] != *item {
		t.Error("should correctly insert a first element into heap")
	}
}

func Test_HeapPush2(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(1, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(2, nil)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 2 elements into heap, and maintain heap property. Top element should be %d", top)
	}
}

func Test_HeapPush3(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 3 elements into heap, and maintain heap property. Top element should be %d", top)
	}
}

func Test_HeapPush3MaxHeap(t *testing.T) {
	heap := New[int](MaxHeap)
	top := NewItem(10, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 3 elements into heap, and maintain heap property. Top element should be %d", top)
	}
}

func Test_HeapPushMany(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(12, nil)
	heap.Push(3, nil)
	heap.Push(4, nil)
	heap.Push(20, nil)
	heap.Push(5, nil)
	heap.Push(12, nil)
	heap.Push(6, nil)
	heap.Push(7, nil)
	heap.Push(8, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert many elements into heap, and maintain heap property. Top element should be %d", top)
	}
}

func Test_Pop(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %v", top)
	}
}

func Test_Pop2(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(12, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d", top)
	}
}

func Test_PopEmpty(t *testing.T) {
	heap := New[int](MinHeap)

	_, err := heap.Pop()

	if err == nil {
		t.Error("should return error when pop from empty heap")
	}
}

func Test_PopOneElement(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d", top)
	}
}

func Test_PopManyElement(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(15, nil)
	heap.Push(1, nil)
	heap.Push(7, nil)
	heap.Push(9, nil)
	heap.Push(11, nil)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d", top)
	}
}

func Test_PopManyMaxHeap(t *testing.T) {
	heap := New[int](MaxHeap)
	top := NewItem(15, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(-1, nil)
	heap.Push(1, nil)
	heap.Push(7, nil)
	heap.Push(9, nil)
	heap.Push(11, nil)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d", top)
	}
}
