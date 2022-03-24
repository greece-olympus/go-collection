package go_collection

import (
	"testing"
)

func FuzzWhere(f *testing.F) {
	f.Add(2, 2)
	f.Add(1, 5)
	f.Add(1, 1)
	f.Add(2, 3)
	f.Fuzz(func(t *testing.T, result int, comparable int) {
		list := NewCollection[int]([]int{1, 2, 3, 3, 2, 5})
		items := list.
			Where(func(v int) bool {
				return v == comparable
			})
		if len(items) != result {
			t.Error("where error")
		}
	})
}

func FuzzWhereFirst(f *testing.F) {
	f.Add(2, 2)
	f.Add(5, 5)

	f.Fuzz(func(t *testing.T, result int, comparable int) {
		list := NewCollection[int]([]int{1, 2, 3, 3, 2, 5})
		item := list.FirstWhere(func(v int) bool {
			return v == comparable
		})
		if item != result {
			t.Error("where error")
		}
	})
}
