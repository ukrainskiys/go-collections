package collections

import (
	"fmt"
	"strings"
)

type Queue[T comparable] interface {
	Offer(element T)
	Pool() T
	Peek() T
	Equal(elements Queue[T]) bool
	Collection[T]
}

type PrimaryQueue[T comparable] struct {
	data []T
	size int
}

func NewQueue[T comparable](elements ...T) *PrimaryQueue[T] {
	s := len(elements)
	queue := &PrimaryQueue[T]{data: make([]T, s), size: s}
	for idx, val := range elements {
		queue.data[idx] = val
	}
	return queue
}

func (p *PrimaryQueue[T]) Offer(element T) {
	p.data = append(p.data, element)
	p.size++
}

func (p *PrimaryQueue[T]) Pool() T {
	result := p.data[0]
	p.data = p.data[1:]
	p.size--
	return result
}

func (p *PrimaryQueue[T]) Peek() T {
	return p.data[0]
}

func (p *PrimaryQueue[T]) Equal(elements Queue[T]) bool {
	if elements == nil {
		return false
	}
	if len(p.data) != elements.Size() {
		return false
	}

	i := 0
	for el := range elements.Iterator() {
		if el != p.data[i] {
			return false
		}
		i++
	}
	return true
}

func (p *PrimaryQueue[T]) Add(element T) {
	p.Offer(element)
}

func (p *PrimaryQueue[T]) AddAll(elements Collection[T]) {
	for el := range elements.Iterator() {
		p.Offer(el)
	}
}

func (p *PrimaryQueue[T]) AddAllSlice(elements []T) {
	for _, el := range elements {
		p.Offer(el)
	}
}

func (p *PrimaryQueue[T]) Contains(element T) bool {
	for _, el := range p.data {
		if el == element {
			return true
		}
	}
	return false
}

func (p *PrimaryQueue[T]) ContainsAll(elements Collection[T]) bool {
	for el := range elements.Iterator() {
		if !p.Contains(el) {
			return false
		}
	}
	return true
}

func (p *PrimaryQueue[T]) ContainsAllSlice(elements []T) bool {
	for _, el := range elements {
		if !p.Contains(el) {
			return false
		}
	}
	return true
}

func (p *PrimaryQueue[T]) Remove(element T) bool {
	for idx, el := range p.data {
		if el == element {
			p.data = append(p.data[:idx], p.data[idx+1:]...)
			p.size--
			return true
		}
	}
	return false
}

func (p *PrimaryQueue[T]) RemoveAll(elements Collection[T]) bool {
	modified := false
	for el := range elements.Iterator() {
		if p.Remove(el) {
			p.size--
			modified = true
		}
	}
	return modified
}

func (p *PrimaryQueue[T]) RemoveAllSlice(elements []T) bool {
	modified := false
	for _, el := range elements {
		if p.Remove(el) {
			p.size--
			modified = true
		}
	}
	return modified
}

func (p *PrimaryQueue[T]) RemoveIf(predicate func(T) bool) bool {
	modified := false
	for _, el := range p.data {
		if predicate(el) && p.Remove(el) {
			p.size--
			modified = true
		}
	}
	return modified
}

func (p *PrimaryQueue[T]) Size() int {
	return p.size
}

func (p *PrimaryQueue[T]) IsEmpty() bool {
	return p.size == 0
}

func (p *PrimaryQueue[T]) Clear() {
	p.data = nil
	p.size = 0
}

func (p *PrimaryQueue[T]) Iterator() <-chan T {
	pool := make(chan T)

	go func() {
		defer close(pool)
		for _, val := range p.data {
			pool <- val
		}
	}()

	return pool
}

func (p *PrimaryQueue[T]) ForEach(do func(T)) {
	for _, el := range p.data {
		do(el)
	}
}

func (p *PrimaryQueue[T]) String() string {
	fmt.Println(p.data)
	var data []string
	for _, el := range p.data {
		data = append(data, fmt.Sprint(el))
	}
	return "Queue=[" + strings.Join(data, ", ") + "]"
}
