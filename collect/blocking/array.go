package blocking

import (
	"github.com/ukrainskiys/go-collections/collect"
	"sync"
)

type ArrayList[T comparable] struct {
	collectionWithSlice[T]
}

func NewList[T comparable](elements ...T) *ArrayList[T] {
	return &ArrayList[T]{
		collectionWithSlice: collectionWithSlice[T]{
			mx:   &sync.RWMutex{},
			data: &elements,
		},
	}
}

func NewListOf[T comparable](collection collect.Collection[T]) *ArrayList[T] {
	data := make([]T, collection.Size())
	i := 0
	for el := range collection.Iterator() {
		data[i] = el
		i++
	}
	return NewList[T](data...)
}

func (a *ArrayList[T]) Get(index int) T {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return (*a.data)[index]
}

func (a *ArrayList[T]) SafeGet(index int) (T, bool) {
	a.mx.RLock()
	defer a.mx.RUnlock()
	if len(*a.data) > index {
		return (*a.data)[index], true
	} else {
		var t T
		return t, false
	}
}

func (a *ArrayList[T]) IndexOf(element T) int {
	a.mx.RLock()
	defer a.mx.RUnlock()
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
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.data
}

func (a *ArrayList[T]) Equal(array *ArrayList[T]) bool {
	a.mx.RLock()
	defer a.mx.RUnlock()
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
