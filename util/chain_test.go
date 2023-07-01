package util

import (
	"context"
	"log"
	"testing"
)

type handle1 struct{}

func (h handle1) Handle(ctx context.Context, value string, c *Chain[string]) {
	log.Println("handle1: ", value)
	c.Next(ctx, value)
}

type handle2 struct{}

func (h handle2) Handle(ctx context.Context, value string, c *Chain[string]) {
	log.Println("handle2: ", value)
	c.Next(ctx, value)
}

type handle3 struct{}

func (h handle3) Handle(ctx context.Context, value string, c *Chain[string]) {
	log.Println("handle3: ", value)
}

type handle4 struct{}

func (h handle4) Handle(ctx context.Context, value string, c *Chain[string]) {
	log.Println("handle4: ", value)
	c.Next(ctx, value)
}

func TestChain(t *testing.T) {
	ch := NewSimpleChain[string]()
	ch.Add(handle1{})
	ch.Add(handle2{})
	ch.Add(handle3{})
	ch.Add(handle4{})
	ch.Execute(context.Background(), "chain execute value")
}
