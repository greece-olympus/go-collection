package go_collection

func Append[T any](items []T, item T) []T {
	items = append(items, item)
	return items
}

func (c *CollectionOpt[T]) Append(item T) Collection[T] {
	c.items = Append(c.items, item)
	return c
}
