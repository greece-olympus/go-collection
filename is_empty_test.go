package go_collection

import "testing"

func TestFuncIsEmpty(t *testing.T) {
	if !IsEmpty([]int{}) {
		t.Fail()
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewCollection[int]([]int{})
	if !list.IsEmpty() {
		t.Fail()
	}
}
