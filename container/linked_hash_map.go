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

type LinkedHashMap[K comparable, V any] struct {
	inner map[K]V
	keys []K
}

func NewLinkedHashMap[K comparable, V any]() LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		inner: make(map[K]V),
		keys: make([]K, 0, 10),
	}
}

func (l *LinkedHashMap[K, V]) Put(key K, value V) {
	_, ok := l.inner[key]
	if !ok {
		l.keys = append(l.keys, key)
	}
	l.inner[key] = value
}

func (l *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	v, ok := l.inner[key]

	return v, ok
}

func (l *LinkedHashMap[K, V]) Delete(key K) {
	if _, ok := l.inner[key]; ok {
		delete(l.inner, key)
		l.findKeyInKeysDel(key)
	}
}

func (l *LinkedHashMap[K, V]) ForEach(f func(key K, value V)) {
	for _, key := range l.keys {
		value := l.inner[key]
		f(key, value)
	}
}

func (l *LinkedHashMap[K, V]) Contains(key K) bool {
	_, ok := l.inner[key]

	return ok
}

func (l *LinkedHashMap[K, V]) findKeyInKeysDel(key K) {
	for i, k := range l.keys {
		if k == key {
			newKeys := make([]K, len(l.keys), cap(l.keys))
			copy(newKeys, l.keys[:i])
			copy(newKeys, l.keys[i+1:])
			l.keys = newKeys
			break
		}
	}
}

func (l *LinkedHashMap[K, V]) Size() int {
	return len(l.keys)
}

func (l *LinkedHashMap[K, V]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LinkedHashMap[K, V]) Keys() []K {
	copyKeys := make([]K, len(l.keys))
	copy(copyKeys, l.keys)

	return copyKeys
}

func (l *LinkedHashMap[K, V]) Clear() {
	l.inner = nil
	l.inner = make(map[K]V)

	l.keys = nil
	l.keys = make([]K, 0)
}