/*
 * Copyright 2025 ekko.cheung
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

package async

import "sync/atomic"

type completefuture[T any] struct {
	ch      chan T
	value   T
	errFunc func(any)
	isclose int32
}

func (c *completefuture[T]) RunAsync(f func() T) {
	go func() {
		defer func() {
			if err := recover(); err != nil && c.errFunc != nil {
				c.errFunc(err)
			}
		}()
		c.ch <- f()
	}()
}

func (c *completefuture[T]) RunAsyncWithErrHandle(f func() T, errHandle func(any)) {
	c.errFunc = errHandle
	c.RunAsync(f)
}

func (c *completefuture[T]) Get() T {
	if atomic.CompareAndSwapInt32(&c.isclose, 0, 1) {
		c.value = <-c.ch
		close(c.ch)
	}

	return c.value
}

func RunAsync[T any](f func() T) *completefuture[T] {
	c := &completefuture[T]{
		ch: make(chan T),
	}
	c.RunAsync(f)

	return c
}

func RunAsyncWithErrHandle[T any](f func() T, errHandle func(any)) *completefuture[T] {
	c := &completefuture[T]{
		ch: make(chan T),
	}
	c.RunAsyncWithErrHandle(f, errHandle)

	return c
}
