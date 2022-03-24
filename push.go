package go_collection

func Push[T any](items []T, item T) []T {
	items = append(items, item)
	return items
}

func (c *CollectionOpt[T]) Push(item T) {
	c.items = Push(c.items, item)
}
