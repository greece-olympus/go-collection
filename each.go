package go_collection

func Each[T any](items []T, fn func(int, T)) {
	for k, v := range items {
		fn(k, v)
	}
}

func (c CollectionOpt[T]) Each(fn func(int, T)) {
	Each(c.items, fn)
}
