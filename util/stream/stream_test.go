//go:build go1.18 || 1.19 || 1.20

package stream

import (
	"strconv"
	"testing"
)

var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestStreamImpl_ToArray(t *testing.T) {
	s := New(a)
	t.Log(s.ToArray())
}

func TestStreamImpl_ForEach(t *testing.T) {
	s := New(a)
	s.ForEach(func(i int) {
		t.Log(i)
	})
}

func TestStreamImpl_Count(t *testing.T) {
	s := New(a)
	t.Log(s.Count())
}

func TestStreamImpl_Reverse(t *testing.T) {
	s := New(a)
	t.Log(s.Reverse().ToArray())
}

func TestStreamImpl_Filter(t *testing.T) {
	s := New(a)
	array := s.Filter(func(i int) bool {
		return i > 10
	}).Reverse().ToArray()
	t.Log(array)
}

func TestStreamImpl_Limit(t *testing.T) {
	s := New(a)
	t.Log(s.Limit(5).ToArray())
}

func TestStreamImpl_Skip(t *testing.T) {
	s := New(a)
	t.Log(s.Skip(5).ToArray())
}

func TestStreamImpl_Max(t *testing.T) {
	s := New(a)
	max := s.Max(func(a, b int) int {
		return a - b
	})
	t.Log(max)
}

func TestStreamImpl_Min(t *testing.T) {
	s := New(a)
	min := s.Min(func(a, b int) int {
		return a - b
	})
	t.Log(min)
}

func TestStreamImpl_Sorted(t *testing.T) {
	s := New([]int{3, 1, 4, 6, 2, 6, 10, 7})
	array := s.Sorted(func(a, b int) int {
		return a - b
	}).ToArray()
	t.Log(array)
}

func TestMap(t *testing.T) {
	arr := Map(a, func(i int) string {
		return strconv.Itoa(i)
	})
	t.Log(arr)
}
