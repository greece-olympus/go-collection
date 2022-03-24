package go_collection

import (
	"math"
)

func Chunk[T any](items []T, num int) (out [][]T) {
	count := len(items)
	if count == 0 {
		return
	}
	length := int(math.Ceil(float64(count) / float64(num)))
	out = make([][]T, length)
	for i := 0; i < length; i++ {
		start := num * i
		if start > count {
			return
		}
		end := num * (i + 1)
		if end > count {
			end = count
		}
		out[i] = items[start:end]
		//	fmt.Println(c.items[start:end])
	}
	return
}

func (c *CollectionOpt[T]) Chunk(num int) [][]T {
	return Chunk(c.items, num)
}
