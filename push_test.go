package go_collection

import "testing"

func TestPush(t *testing.T) {
	col := NewCollection([]int{1, 2, 3})
	col.Push(5)
	if col.All()[3] != 5 {
		t.Fail()
	}
}
