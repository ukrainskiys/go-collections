package collect

import "fmt"

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
	ForEach(do func(T))
	String() string
}

type collectionWithSlice[T comparable] struct {
	data *[]T
}

func (c *collectionWithSlice[T]) Add(element T) {
	*c.data = append(*c.data, element)
}

func (c *collectionWithSlice[T]) AddAll(elements Collection[T]) {
	for el := range elements.Iterator() {
		c.Add(el)
	}
}

func (c *collectionWithSlice[T]) AddAllSlice(elements []T) {
	for _, el := range elements {
		c.Add(el)
	}
}

func (c *collectionWithSlice[T]) Contains(element T) bool {
	for _, el := range *c.data {
		if el == element {
			return true
		}
	}
	return false
}

func (c *collectionWithSlice[T]) ContainsAll(elements Collection[T]) bool {
	for el := range elements.Iterator() {
		if !c.Contains(el) {
			return false
		}
	}
	return true
}

func (c *collectionWithSlice[T]) ContainsAllSlice(elements []T) bool {
	for _, el := range elements {
		if !c.Contains(el) {
			return false
		}
	}
	return true
}

func (c *collectionWithSlice[T]) Remove(element T) bool {
	for idx, el := range *c.data {
		if el == element {
			*c.data = append((*c.data)[:idx], (*c.data)[idx+1:]...)
			return true
		}
	}
	return false
}

func (c *collectionWithSlice[T]) RemoveAll(elements Collection[T]) bool {
	modified := false
	for el := range elements.Iterator() {
		if c.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) RemoveAllSlice(elements []T) bool {
	modified := false
	for _, el := range elements {
		if c.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) RemoveIf(predicate func(T) bool) bool {
	modified := false
	for el := range c.Iterator() {
		if predicate(el) && c.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) Size() int {
	return len(*c.data)
}

func (c *collectionWithSlice[T]) IsEmpty() bool {
	return c.Size() == 0
}

func (c *collectionWithSlice[T]) Clear() {
	*c.data = nil
}

func (c *collectionWithSlice[T]) Iterator() <-chan T {
	pool := make(chan T, len(*c.data))
	defer close(pool)

	for _, val := range *c.data {
		pool <- val
	}

	return pool
}

func (c *collectionWithSlice[T]) ForEach(do func(T)) {
	for _, val := range *c.data {
		do(val)
	}
}

func (c *collectionWithSlice[T]) String() string {
	return fmt.Sprint(*c.data)
}
