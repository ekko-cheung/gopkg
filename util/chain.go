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

package util

import "context"

type Handler[T any] interface {
	Handle(ctx context.Context, value T, c *Chain[T])
}

type SimpleChain[T any] struct {
	c *Chain[T]
}

type Chain[T any] struct {
	index    int
	handlers []Handler[T]
}

func (c *Chain[T]) Next(ctx context.Context, value T) {
	c.index++
	if c.index == len(c.handlers) {
		return
	}
	c.handlers[c.index].Handle(ctx, value, c)
}

func NewSimpleChain[T any]() *SimpleChain[T] {
	return &SimpleChain[T]{
		c: &Chain[T]{
			handlers: make([]Handler[T], 0),
		},
	}
}

func (s *SimpleChain[T]) Add(h Handler[T]) {
	s.c.handlers = append(s.c.handlers, h)
}

func (s *SimpleChain[T]) Execute(ctx context.Context, value T) {
	h := s.c.handlers[0]
	h.Handle(ctx, value, s.c)
	s.c.index = 0
}
