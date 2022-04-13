package go_collection

func Unique[T BaseType](list []T) (out []T) {
	Each(list, func(i int, v T) {
		if IndexOf(out, v, func(c T, p T) bool {
			return c == p
		}) == -1 {
			out = append(out, v)
		}
	})
	return
}
