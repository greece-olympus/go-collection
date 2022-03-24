package go_collection

func IsEmpty[T any](items []T) bool {
	return len(items) == 0
}

func (c *CollectionOpt[T]) IsEmpty() bool {
	return IsEmpty(c.items)
}
