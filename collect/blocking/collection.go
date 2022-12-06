package blocking

import (
	"fmt"
	"github.com/ukrainskiys/go-collections/collect"
	"sync"
)

type collectionWithSlice[T comparable] struct {
	mx   *sync.RWMutex
	data *[]T
}

func (c *collectionWithSlice[T]) Add(element T) {
	c.mx.Lock()
	defer c.mx.Unlock()
	*c.data = append(*c.data, element)
}

func (c *collectionWithSlice[T]) AddAll(elements collect.Collection[T]) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for el := range elements.Iterator() {
		c.Add(el)
	}
}

func (c *collectionWithSlice[T]) AddAllSlice(elements []T) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for _, el := range elements {
		c.Add(el)
	}
}

func (c *collectionWithSlice[T]) Contains(element T) bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, el := range *c.data {
		if el == element {
			return true
		}
	}
	return false
}

func (c *collectionWithSlice[T]) ContainsAll(elements collect.Collection[T]) bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for el := range elements.Iterator() {
		if !c.Contains(el) {
			return false
		}
	}
	return true
}

func (c *collectionWithSlice[T]) ContainsAllSlice(elements []T) bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, el := range elements {
		if !c.Contains(el) {
			return false
		}
	}
	return true
}

func (c *collectionWithSlice[T]) Remove(element T) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	for idx, el := range *c.data {
		if el == element {
			*c.data = append((*c.data)[:idx], (*c.data)[idx+1:]...)
			return true
		}
	}
	return false
}

func (c *collectionWithSlice[T]) RemoveAll(elements collect.Collection[T]) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	modified := false
	for el := range elements.Iterator() {
		if c.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) RemoveAllSlice(elements []T) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	modified := false
	for _, el := range elements {
		if c.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) RemoveIf(predicate func(T) bool) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	modified := false
	var arr []T
	for el := range c.Iterator() {
		if predicate(el) {
			arr = append(arr, el)
		}
	}

	for _, val := range arr {
		if c.Remove(val) {
			modified = true
		}
	}
	return modified
}

func (c *collectionWithSlice[T]) Size() int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return len(*c.data)
}

func (c *collectionWithSlice[T]) IsEmpty() bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.Size() == 0
}

func (c *collectionWithSlice[T]) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()
	*c.data = nil
}

func (c *collectionWithSlice[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		defer close(pool)
		c.mx.RLock()
		defer c.mx.RUnlock()
		for _, val := range *c.data {
			pool <- val
		}
	}()

	return pool
}

func (c *collectionWithSlice[T]) ForEach(do func(T)) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, val := range *c.data {
		do(val)
	}
}

func (c *collectionWithSlice[T]) String() string {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return fmt.Sprint(*c.data)
}
