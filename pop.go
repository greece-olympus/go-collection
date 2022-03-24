package go_collection

func Pop[T any](items []T) (out T) {
	if len(items) <= 0 {
		return
	}
	out = items[0]
	items = items[1:]
	return
}

func (c *CollectionOpt[T]) Pop() T {
	return Pop(c.items)
}
