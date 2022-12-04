package collections

import (
	"fmt"
	"strings"
)

type Array[T comparable] struct {
	data []T
}

func NewArray[T comparable](elements ...T) *Array[T] {
	array := &Array[T]{data: make([]T, len(elements))}
	for i, el := range elements {
		array.data[i] = el
	}
	return array
}

func NewArrayOf[T comparable](collection Collection[T]) *Array[T] {
	array := &Array[T]{data: []T{}}
	for el := range collection.Iterator() {
		array.data = append(array.data, el)
	}
	return array
}

func (a *Array[T]) Get(index int) (T, bool) {
	if len(a.data) > index {
		return a.data[index], true
	} else {
		var t T
		return t, false
	}
}

func (a *Array[T]) Add(element T) {
	a.data = append(a.data, element)
}

func (a *Array[T]) AddAll(elements Collection[T]) {
	for el := range elements.Iterator() {
		a.Add(el)
	}
}

func (a *Array[T]) AddAllSlice(elements []T) {
	for _, el := range elements {
		a.Add(el)
	}
}

func (a *Array[T]) Contains(element T) bool {
	for _, el := range a.data {
		if el == element {
			return true
		}
	}
	return false
}

func (a *Array[T]) ContainsAll(elements Collection[T]) bool {
	for el := range elements.Iterator() {
		if !a.Contains(el) {
			return false
		}
	}
	return true
}

func (a *Array[T]) ContainsAllSlice(elements []T) bool {
	for _, el := range elements {
		if !a.Contains(el) {
			return false
		}
	}
	return true
}

func (a *Array[T]) IndexOf(element T) int {
	i := 0
	for el := range a.Iterator() {
		if el == element {
			return i
		}
		i++
	}
	return -1
}

func (a *Array[T]) Remove(element T) bool {
	index := a.IndexOf(element)
	if index == -1 {
		return false
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return true
}

func (a *Array[T]) RemoveAll(elements Collection[T]) bool {
	modified := false
	for el := range elements.Iterator() {
		if a.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (a *Array[T]) RemoveAllSlice(elements []T) bool {
	modified := false
	for _, el := range elements {
		if a.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (a *Array[T]) RemoveIf(predicate func(T) bool) bool {
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

func (a *Array[T]) Size() int {
	return len(a.data)
}

func (a *Array[T]) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Array[T]) Clear() {
	a.data = nil
}

func (a *Array[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		defer close(pool)
		for _, val := range a.data {
			pool <- val
		}
	}()

	return pool
}

func (a *Array[T]) Equal(elements Collection[T]) bool {
	if elements == nil {
		return false
	}
	if len(a.data) != elements.Size() {
		return false
	}

	for _, val := range a.data {
		if !elements.Contains(val) {
			return false
		}
	}
	return true
}

func (a *Array[T]) ForEach(do func(T)) {
	for _, val := range a.data {
		do(val)
	}
}

func (a *Array[T]) Slice() *[]T {
	return &a.data
}

func (a *Array[T]) String() string {
	var data []string
	for _, el := range a.data {
		data = append(data, fmt.Sprint(el))
	}
	return "Array=[" + strings.Join(data, ", ") + "]"
}
