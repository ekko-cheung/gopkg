package list

import (
	"fmt"
	"strings"
)

type Deque[T any] struct {
	element    []*node[T]
	head, tail int
	size       int
}

type node[T any] struct {
	value T
}

func NewDeque[T any](num int) *Deque[T] {
	if num < 0 {
		num = 16
	}
	d := Deque[T]{
		element: make([]*node[T], num),
	}

	return &d
}

func (d *Deque[T]) AddFirst(value T) {
	d.head = dec(d.head, len(d.element))
	d.element[d.head] = &node[T]{value}
	d.size++
	if d.head == d.tail {
		d.grow()
	}
}

func (d *Deque[T]) AddLast(value T) {
	d.element[d.tail] = &node[T]{value}
	d.size++
	d.tail = inc(d.tail, len(d.element))
	if d.head == d.tail {
		d.grow()
	}
}

func (d *Deque[T]) GetFirst() T {
	n := d.element[d.head]
	if n != nil {
		return n.value
	}

	return *new(T)
}

func (d *Deque[T]) GetLast() T {
	n := d.element[dec(d.tail, len(d.element))]
	if n != nil {
		return n.value
	}
	return *new(T)
}

func (d *Deque[T]) RemoveFirst() T {
	n := d.element[d.head]
	if n != nil {
		d.element[d.head] = nil
		d.head = inc(d.head, len(d.element))
		d.size--

		return n.value
	}

	return *new(T)
}

func (d *Deque[T]) RemoveLast() T {
	t := dec(d.tail, len(d.element))
	n := d.element[t]
	if n != nil {
		d.element[t] = nil
		d.tail = t
		d.size--

		return n.value
	}

	return *new(T)
}

func dec(i, modules int) int {
	i = i - 1
	if i < 0 {
		i = modules - 1
	}

	return i
}

func inc(i, modules int) int {
	i = i + 1
	if i >= modules {
		i = 0
	}

	return i
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) String() string {
	if d.size == 0 {
		return "[]"
	}
	b := strings.Builder{}
	b.WriteByte('[')
	size := d.size
	i := d.head
	for size > 1 {
		n := d.element[i]
		i = inc(i, len(d.element))
		b.WriteString(fmt.Sprintf("%v", n.value))
		b.WriteByte(',')
		size--
	}
	n := d.element[i]
	b.WriteString(fmt.Sprintf("%v", n.value))
	b.WriteByte(']')

	return b.String()
}

func (d *Deque[T]) grow() {
	newSize := len(d.element) << 1
	newEle := make([]*node[T], newSize)
	newIndex := 0
	for i := d.head; i < len(d.element); i++ {
		newEle[newIndex] = d.element[i]
		newIndex++
	}
	for i := d.head; i > 0; i-- {
		newEle[newIndex] = d.element[i]
		newIndex++
	}
	d.element = newEle
	d.head = 0
	d.tail = newIndex
}
