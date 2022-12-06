package blocking

import (
	"fmt"
	"github.com/ukrainskiys/go-collections/collect"
	"strings"
	"sync"
)

type Set[T comparable] interface {
	Equal(elements Set[T]) bool
	collect.Collection[T]
}

type HashSet[T comparable] struct {
	data map[any]interface{}
	mx   *sync.RWMutex
}

func NewSet[T comparable](elements ...T) *HashSet[T] {
	data := make(map[any]interface{})
	for _, el := range elements {
		data[el] = nil
	}
	return &HashSet[T]{
		data: data,
		mx:   &sync.RWMutex{},
	}
}

func NewSetOf[T comparable](elements collect.Collection[T]) *HashSet[T] {
	data := make(map[any]interface{})
	for el := range elements.Iterator() {
		data[el] = nil
	}
	return &HashSet[T]{
		data: data,
		mx:   &sync.RWMutex{},
	}
}

func (s *HashSet[T]) Equal(elements Set[T]) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	if elements == nil {
		return false
	}
	if len(s.data) != elements.Size() {
		return false
	}

	for key := range s.data {
		if !elements.Contains(key.(T)) {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) Size() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.data)
}

func (s *HashSet[T]) IsEmpty() bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.data) == 0
}

func (s *HashSet[T]) Contains(element T) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	_, ok := s.data[element]
	return ok
}

func (s *HashSet[T]) ContainsAll(elements collect.Collection[T]) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for el := range elements.Iterator() {
		if !s.Contains(el) {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) ContainsAllSlice(elements []T) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for _, e := range elements {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) Remove(element T) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.data, element)
	_, ok := s.data[element]
	return !ok
}

func (s *HashSet[T]) RemoveAll(elements collect.Collection[T]) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	modified := false
	for el := range elements.Iterator() {
		if s.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (s *HashSet[T]) RemoveAllSlice(elements []T) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	modified := false
	for _, e := range elements {
		if s.Remove(e) {
			modified = true
		}
	}
	return modified
}

func (s *HashSet[T]) RemoveIf(predicate func(T) bool) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	modified := false
	for key := range s.data {
		if predicate(key.(T)) {
			if s.Remove(key.(T)) {
				modified = true
			}
		}
	}
	return modified
}

func (s *HashSet[T]) Add(element T) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	s.data[element] = nil
}

func (s *HashSet[T]) AddAll(elements collect.Collection[T]) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for el := range elements.Iterator() {
		s.data[el] = nil
	}
}

func (s *HashSet[T]) AddAllSlice(elements []T) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for _, el := range elements {
		s.data[el] = nil
	}
}

func (s *HashSet[T]) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	for k := range s.data {
		delete(s.data, k)
	}
}

func (s *HashSet[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		s.mx.RLock()
		defer s.mx.RUnlock()
		defer close(pool)
		for key := range s.data {
			pool <- key.(T)
		}
	}()

	return pool
}

func (s *HashSet[T]) ForEach(do func(T)) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for key := range s.data {
		do(key.(T))
	}
}

func (s *HashSet[T]) String() string {
	s.mx.RLock()
	defer s.mx.RUnlock()
	var data []string
	for el := range s.data {
		data = append(data, fmt.Sprint(el))
	}
	return "[" + strings.Join(data, " ") + "]"
}
