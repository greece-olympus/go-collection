package go_collection

func (c CollectionOpt[T]) All() []T {
	return c.items
}

func (c CollectionOpt[T]) Get() []T {
	return c.items
}
