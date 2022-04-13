package go_collection

import "testing"

func TestIndexOf(t *testing.T) {
	if IndexOf([]int{1, 2, 4, 5}, 4, func(c int, p int) bool {
		return c == p
	}) != 2 {
		t.Fail()
	}
	if IndexOf([]int{1, 2, 4, 5}, 9, func(c int, p int) bool {
		return c == p
	}) != -1 {
		t.Fail()
	}
}
