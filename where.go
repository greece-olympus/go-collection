package go_collection

func Where[T any](items []T, fn func(T) bool) (list []T) {
	Each(items, func(k int, v T) {
		if result := fn(v); result {
			list = append(list, v)
		}
	})
	return
}

func (c *CollectionOpt[T]) Where(fn func(T) bool) []T {
	return Where(c.items, fn)
}

func (c *CollectionOpt[T]) FirstWhere(fn func(T) bool) (out T) {
	return First(Where(c.items, fn))
}
