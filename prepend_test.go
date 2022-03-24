package go_collection

import "testing"

func TestPrepend(t *testing.T) {
	if NewCollection([]int{1, 2, 3}).Prepend(4).All()[0] != 4 {
		t.Fail()
	}
}
