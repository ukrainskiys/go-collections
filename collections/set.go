package collections

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	data map[any]interface{}
}

func NewSet[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{make(map[any]interface{})}
	for _, el := range elements {
		set.data[el] = nil
	}
	return set
}

func NewSetOf[T comparable](elements Collection[T]) *Set[T] {
	set := &Set[T]{make(map[any]interface{})}
	for el := range elements.Iterator() {
		set.data[el] = nil
	}
	return set
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set[T]) Contains(element T) bool {
	_, ok := s.data[element]
	return ok
}

func (s *Set[T]) ContainsAll(elements Collection[T]) bool {
	for el := range elements.Iterator() {
		if !s.Contains(el) {
			return false
		}
	}
	return true
}

func (s *Set[T]) ContainsAllSlice(elements []T) bool {
	for _, e := range elements {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Remove(element T) bool {
	delete(s.data, element)
	_, ok := s.data[element]
	return !ok
}

func (s *Set[T]) RemoveAll(elements Collection[T]) bool {
	modified := false
	for el := range elements.Iterator() {
		if s.Remove(el) {
			modified = true
		}
	}
	return modified
}

func (s *Set[T]) RemoveAllSlice(elements []T) bool {
	modified := false
	for _, e := range elements {
		if s.Remove(e) {
			modified = true
		}
	}
	return modified
}

func (s *Set[T]) RemoveIf(predicate func(T) bool) bool {
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

func (s *Set[T]) Add(element T) {
	s.data[element] = nil
}

func (s *Set[T]) AddAll(elements Collection[T]) {
	for el := range elements.Iterator() {
		s.data[el] = nil
	}
}

func (s *Set[T]) AddAllSlice(elements []T) {
	for _, el := range elements {
		s.data[el] = nil
	}
}

func (s *Set[T]) Clear() {
	for k := range s.data {
		delete(s.data, k)
	}
}

func (s *Set[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		defer close(pool)
		for key := range s.data {
			pool <- key.(T)
		}
	}()

	return pool
}

func (s *Set[T]) Equal(elements Collection[T]) bool {
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

func (s *Set[T]) ForEach(do func(T)) {
	for key := range s.data {
		do(key.(T))
	}
}

func (s *Set[T]) String() string {
	var data []string
	for el := range s.data {
		data = append(data, fmt.Sprint(el))
	}
	return "Set=[" + strings.Join(data, ", ") + "]"
}
