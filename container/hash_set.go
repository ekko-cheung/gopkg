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

type HashSet[T comparable] struct {
	defaultSize int
	data        map[T]struct{}
}

func NewHashSet[T comparable](i ...int) *HashSet[T] {
	if len(i) == 0 {
		return &HashSet[T]{}
	}

	return &HashSet[T]{
		defaultSize: i[0],
	}
}

func NewHashSetFromSlice[T comparable](s []T) *HashSet[T] {
	if len(s) == 0 {
		return &HashSet[T]{}
	}

	h := NewHashSet[T](len(s))
	for i := range s {
		h.Add(s[i])
	}

	return h
}

func (h *HashSet[T]) Add(t T) {
	h.ifNilCreate()
	h.data[t] = struct{}{}
}

func (h *HashSet[T]) isNil() bool {
	return h.data == nil
}

func (h *HashSet[T]) ifNilCreate() {
	if h.isNil() {
		if h.defaultSize > 0 {
			h.data = make(map[T]struct{}, h.defaultSize)
		} else {
			h.data = make(map[T]struct{})
		}
	}
}

func (h *HashSet[T]) Contains(t T) bool {
	if h.isNil() {
		return false
	}

	_, ok := h.data[t]

	return ok
}

func (h *HashSet[T]) Remove(t T) {
	if !h.isNil() {
		delete(h.data, t)
	}
}

func (h *HashSet[T]) Clear() {
	h.data = nil
}

func (h *HashSet[T]) Len() int {
	return len(h.data)
}

func (h *HashSet[T]) IsEmpty() bool {
	return h.data == nil || h.Len() == 0
}

func (h *HashSet[T]) ToSlice() []T {
	if h.isNil() {
		return []T{}
	}

	arr := make([]T, len(h.data))
	i := 0
	for t := range h.data {
		arr[i] = t
		i++
	}

	return arr
}

func (h *HashSet[T]) ForEach(f func(T)) {
	if h.isNil() {
		return
	}

	for t := range h.data {
		f(t)
	}
}

func (h *HashSet[T]) Intersection(h2 *HashSet[T]) []T {
	res := make([]T, 0, 4)
	h.ForEach(func(item T) {
		if h2.Contains(item) {
			res = append(res, item)
		}
	})

	return res
}

func (h *HashSet[T]) Union(h2 *HashSet[T]) []T {
	h2.ForEach(func(item T) {
		if !h.Contains(item) {
			h.Add(item)
		}
	})

	return h.ToSlice()
}
