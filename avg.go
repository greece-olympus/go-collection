package go_collection

func Avg[T any](items []T, fn func(T) int) int {
	count := 0
	Each(items, func(i int, v T) {
		count += fn(v)
	})
	if count == 0 {
		return 0
	}
	return count / len(items)
}

func (c *CollectionOpt[T]) Avg(fn func(T) int) int {
	return Avg(c.items, fn)
}
