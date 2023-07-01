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
