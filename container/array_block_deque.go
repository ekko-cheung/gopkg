package container

import "sync"

type ArrayBlockDeque[T any] struct {
	element           []*node[T]
	notFull, notEmpty *sync.Cond
	index, size       int
}

func NewArrayBlockDeque[T any](num int) *ArrayBlockDeque[T] {
	if num <= 0 {
		num = 16
	}

	return &ArrayBlockDeque[T]{
		element:  make([]*node[T], num),
		notFull:  sync.NewCond(&sync.Mutex{}),
		notEmpty: sync.NewCond(&sync.Mutex{}),
	}
}

func (a *ArrayBlockDeque[T]) Put(value T) {
	a.notFull.L.Lock()
	defer a.notFull.L.Unlock()
	if a.size == len(a.element) {
		a.notFull.Wait()
	}
	a.element[a.index] = &node[T]{value}
	a.index++
	if a.index == len(a.element) {
		a.index = 0
	}
	a.size++
	a.notEmpty.Signal()
}

func (a *ArrayBlockDeque[T]) Take() T {
	a.notEmpty.L.Lock()
	defer a.notEmpty.L.Unlock()
	if a.size == 0 {
		a.notEmpty.Wait()
	}
	n := a.element[a.index]
	a.element[a.index] = nil
	a.index++
	if a.index == len(a.element) {
		a.index = 0
	}
	a.size--
	a.notFull.Signal()

	return n.Val()
}
