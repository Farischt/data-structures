package heap

import "testing"

func Test_NewItem(t *testing.T) {
	item := NewItem(1, nil)
	if item.Value != 1 {
		t.Error("should create new item")
	}
}
