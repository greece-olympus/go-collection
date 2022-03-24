package go_collection

func Map[T, B any](items []T, fn func(T) B) (out []B) {
	Each(items, func(i int, v T) {
		out = append(out, fn(v))
	})
	return
}

func (c *CollectionOpt[T]) Map(fn func(T) any) (out []any) {
	return Map(c.items, fn)
}
