package go_collection

import "testing"

func TestAppend(t *testing.T) {
	if NewCollection([]int{1, 2, 3}).Append(4).All()[3] != 4 {
		t.Fail()
	}
}
