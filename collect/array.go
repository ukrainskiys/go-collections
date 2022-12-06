package collect

type List[T comparable] interface {
	Get(index int) T
	SafeGet(index int) (T, bool)
	IndexOf(element T) int
	Slice() *[]T
}

type ArrayList[T comparable] struct {
	collectionWithSlice[T]
}

func NewList[T comparable](elements ...T) *ArrayList[T] {
	return &ArrayList[T]{collectionWithSlice[T]{data: &elements}}
}

func NewListOf[T comparable](collection Collection[T]) *ArrayList[T] {
	data := make([]T, collection.Size())
	array := &ArrayList[T]{collectionWithSlice[T]{data: &data}}

	i := 0
	for el := range collection.Iterator() {
		(*array.data)[i] = el
		i++
	}
	return array
}

func (a *ArrayList[T]) Get(index int) T {
	return (*a.data)[index]
}

func (a *ArrayList[T]) SafeGet(index int) (T, bool) {
	if len(*a.data) > index {
		return (*a.data)[index], true
	} else {
		var t T
		return t, false
	}
}

func (a *ArrayList[T]) IndexOf(element T) int {
	i := 0
	for el := range a.Iterator() {
		if el == element {
			return i
		}
		i++
	}
	return -1
}

func (a *ArrayList[T]) Slice() *[]T {
	return a.data
}

func (a *ArrayList[T]) Equal(array *ArrayList[T]) bool {
	if array.data == nil {
		return false
	}
	if len(*a.data) != array.Size() {
		return false
	}

	for idx, val := range *a.data {
		if array.Get(idx) != val {
			return false
		}
	}
	return true
}
