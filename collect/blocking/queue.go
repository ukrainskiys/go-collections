package blocking

import "sync"

type PrimaryQueue[T comparable] struct {
	collectionWithSlice[T]
}

func NewQueue[T comparable](elements ...T) *PrimaryQueue[T] {
	return &PrimaryQueue[T]{
		collectionWithSlice: collectionWithSlice[T]{
			mx:   &sync.RWMutex{},
			data: &elements,
		},
	}
}

func (p *PrimaryQueue[T]) Offer(element T) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.Add(element)
}

func (p *PrimaryQueue[T]) Pool() T {
	p.mx.Lock()
	defer p.mx.Unlock()
	result := (*p.data)[0]
	*p.data = (*p.data)[1:]
	return result
}

func (p *PrimaryQueue[T]) Peek() T {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return (*p.data)[0]
}

func (p *PrimaryQueue[T]) Equal(elements *PrimaryQueue[T]) bool {
	p.mx.RLock()
	defer p.mx.RUnlock()
	if elements == nil {
		return false
	}
	if len(*p.data) != elements.Size() {
		return false
	}

	i := 0
	for el := range elements.Iterator() {
		if el != (*p.data)[i] {
			return false
		}
		i++
	}
	return true
}
