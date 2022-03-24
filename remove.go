package go_collection

func Remove[T any](items []T, index int) (out []T) {
	if index >= len(items) {
		return
	}
	out = append(items[:index], items[index+1:]...)
	return
}

func (c *CollectionOpt[T]) Remove(index int) {
	c.items = Remove(c.items, index)
}
