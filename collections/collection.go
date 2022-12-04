package collections

type Collection[T comparable] interface {
	Add(element T)
	AddAll(elements Collection[T])
	AddAllSlice(elements []T)

	Contains(element T) bool
	ContainsAll(elements Collection[T]) bool
	ContainsAllSlice(elements []T) bool

	Remove(element T) bool
	RemoveAll(elements Collection[T]) bool
	RemoveAllSlice(elements []T) bool
	RemoveIf(predicate func(T) bool) bool

	Size() int
	IsEmpty() bool
	Clear()

	Iterator() <-chan T
	Equal(elements Collection[T]) bool
	ForEach(do func(T))
	String() string
}
