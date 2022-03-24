package go_collection

func Filter[T any](items []T, fn func(int, T) bool) (list []T) {
	Each(items, func(i int, v T) {
		if result := fn(i, v); result {
			list = append(list, v)
		}
	})
	return
}

func (c *CollectionOpt[T]) Filter(fn func(int, T) bool) Collection[T] {
	c.items = Filter(c.items, fn)
	return c
}
