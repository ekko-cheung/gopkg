//go:build go1.18

/*
 * Copyright 2025 veerdone
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

type ringQueue[T any] struct {
	array             []*node[T]
	rear, front, size int
}

func NewRingQueue[T any](size int) *ringQueue[T] {
	return &ringQueue[T]{
		array: make([]*node[T], size),
	}
}

func (r *ringQueue[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *ringQueue[T]) IsFull() bool {
	return r.size == len(r.array)
}

func (r *ringQueue[T]) Push(value T) bool {
	if r.IsFull() {
		return false
	}

	r.array[r.rear] = &node[T]{value: value}
	r.rear = (r.rear + 1) % len(r.array)
	r.size++

	return true
}

func (r *ringQueue[T]) Pop() T {
	if r.IsEmpty() {
		return *new(T)
	}

	node := r.array[r.front]
	r.array[r.front] = nil
	r.size--
	r.front = (r.front + 1) % len(r.array)

	return node.value
}

func (r *ringQueue[T]) Front() T {
	if r.IsEmpty() {
		return *new(T)
	}

	return r.array[r.front].value
}

func (r *ringQueue[T]) Rear() T {
	if r.IsEmpty() {
		return *new(T)
	}

	return r.array[r.rear].value
}
