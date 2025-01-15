//go:build go1.18

/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package container

import (
	"fmt"
	"strings"
)

type ArrayDeque[T any] struct {
	element    []*node[T]
	head, tail int
	size       int
}

func NewArrayDeque[T any](num int) *ArrayDeque[T] {
	if num < 0 {
		num = 16
	}
	d := ArrayDeque[T]{
		element: make([]*node[T], num),
	}

	return &d
}

func (d *ArrayDeque[T]) AddFirst(value T) {
	d.head = dec(d.head, len(d.element))
	d.element[d.head] = &node[T]{value}
	d.size++
	if d.head == d.tail {
		d.grow()
	}
}

func (d *ArrayDeque[T]) AddLast(value T) {
	d.element[d.tail] = &node[T]{value}
	d.size++
	d.tail = inc(d.tail, len(d.element))
	if d.head == d.tail {
		d.grow()
	}
}

func (d *ArrayDeque[T]) GetFirst() T {
	n := d.element[d.head]

	return n.Val()
}

func (d *ArrayDeque[T]) GetLast() T {
	n := d.element[dec(d.tail, len(d.element))]

	return n.Val()
}

func (d *ArrayDeque[T]) RemoveFirst() T {
	n := d.element[d.head]
	if n != nil {
		d.element[d.head] = nil
		d.head = inc(d.head, len(d.element))
		d.size--
	}

	return n.Val()
}

func (d *ArrayDeque[T]) RemoveLast() T {
	t := dec(d.tail, len(d.element))
	n := d.element[t]
	if n != nil {
		d.element[t] = nil
		d.tail = t
		d.size--
	}

	return n.Val()
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

func (d *ArrayDeque[T]) Size() int {
	return d.size
}

func (d *ArrayDeque[T]) String() string {
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

func (d *ArrayDeque[T]) grow() {
	newSize := len(d.element) << 1
	newEle := make([]*node[T], newSize)
	newIndex := 0
	for i := d.head; i < len(d.element); i++ {
		newEle[newIndex] = d.element[i]
		newIndex++
	}
	for i := 0; i < d.head; i++ {
		newEle[newIndex] = d.element[i]
		newIndex++
	}
	d.element = newEle
	d.head = 0
	d.tail = newIndex
}
