package go_collection

import (
	"testing"
)

func TestRemove(t *testing.T) {
	col := NewCollection([]int{1, 2, 3})
	col.Remove(1)
	if col.All()[0] != 1 {
		t.Fail()
	}
	if col.All()[1] != 3 {
		t.Fail()
	}

	col1 := NewCollection([]int{1})
	col1.Remove(0)
	if col1.Count() != 0 {
		t.Fail()
	}

	NewCollection([]int{1}).Remove(10)
}
