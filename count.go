package go_collection

func Count[T any](items []T) int {
	return len(items)
}

func (c CollectionOpt[T]) Count() int {
	return len(c.items)
}
