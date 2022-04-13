package go_collection

import "testing"

func TestUnique(t *testing.T) {
	u := Unique([]int{1, 2, 2, 4, 6})
	if len(u) != 4 {
		t.Fail()
	}
	if u[0] != 1 {
		t.Fail()
	}
	if u[1] != 2 {
		t.Fail()
	}
	if u[2] != 4 {
		t.Fail()
	}
}
