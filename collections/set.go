package collections

type Set[T any] struct {
	data map[any]interface{}
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{make(map[any]interface{})}
}

func NewSetOf[T any](elements Collection[T]) *Set[T] {
	set := NewSet[T]()
	for el := range elements.Iterator() {
		set.data[el] = nil
	}
	return set
}

func NewSetOfSlice[T any](elements []T) *Set[T] {
	set := NewSet[T]()
	for _, el := range elements {
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
	for el := range elements.Iterator() {
		if !s.Remove(el) {
			return false
		}
	}
	return true
}

func (s *Set[T]) RemoveAllSlice(elements []T) bool {
	for _, e := range elements {
		if !s.Remove(e) {
			return false
		}
	}
	return true
}

func (s *Set[T]) RemoveIf(predicate func(T) bool) bool {
	for key := range s.data {
		if predicate(key.(T)) {
			if !s.Remove(key.(T)) {
				return false
			}
		}
	}
	return true
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
