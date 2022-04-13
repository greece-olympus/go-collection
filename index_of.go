package go_collection

func IndexOf[T any](list []T, v T, cmp func(T, T) bool) int {
	for i := 0; i < len(list); i++ {
		if cmp(list[i], v) {
			return i
		}
	}
	return -1
}

func (c *CollectionOpt[T]) IndexOf(v T, cmp func(T, T) bool) int {
	return IndexOf(c.items, v, cmp)
}
