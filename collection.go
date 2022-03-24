package go_collection

type (
	Collection[T any] interface {
		Count() int
		All() []T
		Get() []T
		Where(func(T) bool) []T
		FirstWhere(func(T) bool) T
		Filter(func(int, T) bool) Collection[T]
		First() T
		Each(func(int, T))
		Chunk(int) [][]T
		IsEmpty() bool
		Map(func(T) any) []any
		Avg(func(T) int) int
		Prepend(T) Collection[T]
		Append(T) Collection[T]
		Pop() T
		Push(T)
		Remove(int)
	}

	CollectionOpt[T any] struct {
		items []T
	}
)

func NewCollection[T any](items []T) Collection[T] {
	return &CollectionOpt[T]{
		items: items,
	}
}
