package go_collection

func Prepend[T any](items []T, item T) (out []T) {
	out = append(out, item)
	out = append(out, items...)
	return
}

func (c *CollectionOpt[T]) Prepend(item T) Collection[T] {
	c.items = Prepend(c.items, item)
	return c
}
