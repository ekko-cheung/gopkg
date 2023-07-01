package util

import "context"

type Observer[T any] interface {
	Update(ctx context.Context, value T)
}

type SimpleSubject[T any] struct {
	observers []Observer[T]
}

func NewSimpleSubject[T any]() *SimpleSubject[T] {
	return &SimpleSubject[T]{
		observers: make([]Observer[T], 0),
	}
}

func (s *SimpleSubject[T]) Register(o Observer[T]) {
	s.observers = append(s.observers, o)
}

func (s *SimpleSubject[T]) Notify(ctx context.Context, value T) {
	for i := range s.observers {
		s.observers[i].Update(ctx, value)
	}
}

func (s *SimpleSubject[T]) Remove(o Observer[T]) {
	obs := make([]Observer[T], 0, len(s.observers))
	for i := range s.observers {
		ob := s.observers[i]
		if ob == o {
			continue
		}
		obs = append(obs, ob)
	}
	s.observers = obs
}

func (s *SimpleSubject[T]) AsyncNotify(ctx context.Context, value T) {
	go func() {
		s.Notify(ctx, value)
	}()
}
