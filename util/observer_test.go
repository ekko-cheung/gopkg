package util

import (
	"context"
	"log"
	"testing"
	"time"
)

type subject1 struct{}

func (s subject1) Update(ctx context.Context, value string) {
	log.Println("subject1: ", value)
}

type subject2 struct{}

func (s subject2) Update(ctx context.Context, value string) {
	log.Println("subject2:", value)
}

func TestObserver(t *testing.T) {
	s := NewSimpleSubject[string]()
	s.Register(subject1{})
	s.Register(subject2{})

	t.Run("test_Notify", func(t *testing.T) {
		s.Notify(context.Background(), "notify content")
	})

	t.Run("test_AsyncNotify", func(t *testing.T) {
		s.AsyncNotify(context.Background(), "async notify content")
		time.Sleep(time.Second)
	})
}
