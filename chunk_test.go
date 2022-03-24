package go_collection

import (
	"testing"
)

func FuzzFunChunk(f *testing.F) {
	f.Add(2, 5)
	f.Add(3, 4)
	f.Add(6, 2)
	f.Add(7, 2)
	f.Add(10, 1)
	f.Fuzz(func(t *testing.T, c int, length int) {
		if len(Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, c)) != length {
			t.Fail()
		}
	})
}

func FuzzChunk(f *testing.F) {
	f.Add(2, 5)
	f.Add(3, 4)
	f.Add(6, 2)
	f.Add(7, 2)
	f.Add(10, 1)
	f.Fuzz(func(t *testing.T, c int, length int) {
		list := NewCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		if len(list.Chunk(c)) != length {
			t.Fail()
		}
	})
}
