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

import "sync"

type CopyOnWriteArray[T any] struct {
	lock  sync.Locker
	inner []T
}

func NewCopyOnWriteArray[T any]() CopyOnWriteArray[T] {
	return CopyOnWriteArray[T]{
		inner: make([]T, 0),
		lock:  &sync.Mutex{},
	}
}

func (c *CopyOnWriteArray[T]) Add(item T) {
	c.lock.Lock()
	defer c.lock.Unlock()

	oldInner := c.inner
	l := len(oldInner)
	newInner := make([]T, l+1)
	copy(newInner, oldInner)
	newInner[l] = item
	c.inner = newInner
}

func (c *CopyOnWriteArray[T]) Get(index int) T {
	if index >= len(c.inner) {
		return *new(T)
	}

	return c.inner[index]
}

func (c *CopyOnWriteArray[T]) Remove(index int) {
	if index >= len(c.inner) {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	oldInner := c.inner
	newInner := make([]T, len(oldInner))
	copy(newInner, oldInner)

	i := 0
	for j, item := range newInner {
		if j != index {
			newInner[i] = item
			i++
		}
	}
	newInner = newInner[:i]
	c.inner = newInner
}

func (c *CopyOnWriteArray[T]) Set(index int, item T) {
	if index >= len(c.inner) {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	oldInner := c.inner
	newInner := make([]T, len(oldInner))
	copy(newInner, oldInner)
	newInner[index] = item
	c.inner = newInner
}
