package go_collection

import "testing"

func TestAvg(t *testing.T) {
	if a := NewCollection([]int{1, 1, 2, 4}).Avg(func(i int) int {
		return i
	}); a != 2 {
		t.Fail()
	}
}
