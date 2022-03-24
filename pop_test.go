package go_collection

import "testing"

func TestPop(t *testing.T) {
	if NewCollection([]int{1, 2, 3}).Pop() != 1 {
		t.Fail()
	}

	if NewCollection([]int{2}).Pop() != 2 {
		t.Fail()
	}
}
