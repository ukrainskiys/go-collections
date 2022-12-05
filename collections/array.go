package collections

import (
	"fmt"
	"strings"
)

type List[T comparable] interface {
	Get(index int) T
	SafeGet(index int) (T, bool)
	IndexOf(element T) int
	Slice() *[]T
	Equal(array List[T]) bool
	Collection[T]
}

type ArrayList[T comparable] struct {
	data []T
}

func NewList[T comparable](elements ...T) *ArrayList[T] {
	array := &ArrayList[T]{data: make([]T, len(elements))}
	for i, el := range elements {
		array.data[i] = el
	}
	return array
}

func NewListOf[T comparable](collection Collection[T]) *ArrayList[T] {
	array := &ArrayList[T]{data: []T{}}
	for el := range collection.Iterator() {
		array.data = append(array.data, el)
	}
	return array
}

func (a *ArrayList[T]) Get(index int) T {
	return a.data[index]
}

func (a *ArrayList[T]) SafeGet(index int) (T, bool) {
	if len(a.data) > index {
		return a.data[index], true
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
	return &a.data
}

func (a *ArrayList[T]) Equal(array List[T]) bool {
	if array == nil {
		return false
	}
	if len(a.data) != array.Size() {
		return false
	}

	for idx, val := range a.data {
		if array.Get(idx) != val {
			return false
		}
	}
	return true
}

func (a *ArrayList[T]) Add(element T) {
	a.data = append(a.data, element)
}

func (a *ArrayList[T]) AddAll(elements Collection[T]) {
	for el := range elements.Iterator() {
		a.Add(el)
	}
}

func (a *ArrayList[T]) AddAllSlice(elements []T) {
	for _, el := range elements {
		a.Add(el)
	}
}

func (a *ArrayList[T]) Contains(element T) bool {
	for _, el := range a.data {
		if el == element {
			return true
		}
	}
	return false
}

func (a *ArrayList[T]) ContainsAll(elements Collection[T]) bool {
	for el := range elements.Iterator() {
		if !a.Contains(el) {
			return false
		}
	}
	return true
}

func (a *ArrayList[T]) ContainsAllSlice(elements []T) bool {
	for _, el := range elements {
		if !a.Contains(el) {
			return false
		}
	}
	return true
}

func (a *ArrayList[T]) Remove(element T) bool {
	index := a.IndexOf(element)
	if index == -1 {
		return false
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return true
}

func (a *ArrayList[T]) RemoveAll(elements Collection[T]) bool {
	modified := false
	for el := range elements.Iterator() {
		if a.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (a *ArrayList[T]) RemoveAllSlice(elements []T) bool {
	modified := false
	for _, el := range elements {
		if a.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (a *ArrayList[T]) RemoveIf(predicate func(T) bool) bool {
	modified := false
	var arr []T
	for el := range a.Iterator() {
		if predicate(el) {
			arr = append(arr, el)
		}
	}

	for _, val := range arr {
		if a.Remove(val) {
			modified = true
		}
	}
	return modified
}

func (a *ArrayList[T]) Size() int {
	return len(a.data)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayList[T]) Clear() {
	a.data = nil
}

func (a *ArrayList[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		defer close(pool)
		for _, val := range a.data {
			pool <- val
		}
	}()

	return pool
}

func (a *ArrayList[T]) ForEach(do func(T)) {
	for _, val := range a.data {
		do(val)
	}
}

func (a *ArrayList[T]) String() string {
	var data []string
	for _, el := range a.data {
		data = append(data, fmt.Sprint(el))
	}
	return "Array=[" + strings.Join(data, ", ") + "]"
}
