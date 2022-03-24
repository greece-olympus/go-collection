package go_collection

import "testing"

type Pro struct {
	Name string
}

func TestFilter(t *testing.T) {
	col := NewCollection([]int{1, 3, 2, 3, 4, 6, 3, 3})
	if col.Filter(func(i, v int) bool {
		return v == 3
	}).Count() != 4 {
		t.Fail()
	}
}
