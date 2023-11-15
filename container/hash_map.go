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

type HashMap[K comparable, V any] struct {
	inner map[K]V
	capatiy int
}

func NewHashMap[K comparable, V any](i ...int) HashMap[K, V] {
	h := HashMap[K,V]{}
	if len(i) == 1 {
		h.capatiy = i[0]
		h.inner = make(map[K]V, i[0])
	}

	return h
}

func (h *HashMap[K, V]) Put(key K, value V) {
	h.inner[key] = value
}

func (h *HashMap[K, V]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *HashMap[K, V]) Size() int {
	return len(h.inner)
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	v, ok := h.inner[key]

	return v, ok
}

func (h *HashMap[K, V]) Delete(key K) {
	delete(h.inner, key)
}

func (h *HashMap[K, V]) Clear() {
	h.inner = nil
	h.inner = make(map[K]V)
}

func (h *HashMap[K, V]) ClearSaveCapatiy() {
	h.inner = nil
	h.inner = make(map[K]V, h.capatiy)
}

func (h *HashMap[K, V]) Keys() []K {
	if h.IsEmpty() {
		return []K{}
	}

	keys := make([]K, 0,h.Size())
	for key := range h.inner {
		keys = append(keys, key)
	}

	return keys
}

func (h *HashMap[K, V]) Contains(key K) bool {
	_, ok := h.Get(key)

	return ok
}

func (h *HashMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range h.inner {
		f(k, v)
	}
}