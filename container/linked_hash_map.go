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
		keys: make([]K, 0),
	}
}

func (l *LinkedHashMap[K, V]) Insert(key K, value V) {
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
	delete(l.inner, key)
	l.findKeyInKeysDel(key)
}

func (l *LinkedHashMap[K, V]) ForEach(f func(key K, value V)) {
	for _, key := range l.keys {
		value := l.inner[key]
		f(key, value)
	}
}

func (l *LinkedHashMap[K, V]) findKeyInKeysDel(key K) {
	index := 0
	for _, v := range l.keys {
		if v != key {
			l.keys[index] = v
			index++
		}
	}

	l.keys = l.keys[:index]
}
