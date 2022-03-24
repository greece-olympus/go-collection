package go_collection

func First[T any](items []T) (out T) {
	if len(items) <= 0 {
		return
	}
	return items[0]
}

func (c *CollectionOpt[T]) First() T {
	return First(c.items)
}
