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

package stream

import (
	"strconv"
	"testing"
)

var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestStreamImpl_ToArray(t *testing.T) {
	s := New(a)
	t.Log(s.ToSlice())
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
	t.Log(s.Reverse().ToSlice())
}

func TestStreamImpl_Filter(t *testing.T) {
	s := New(a)
	array := s.Filter(func(i int) bool {
		return i > 8
	}).Reverse().ToSlice()
	t.Log(array)
}

func TestStreamImpl_Limit(t *testing.T) {
	s := New(a)
	t.Log(s.Limit(5).ToSlice())
}

func TestStreamImpl_Skip(t *testing.T) {
	s := New(a)
	t.Log(s.Skip(5).ToSlice())
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
	}).ToSlice()
	t.Log(array)
}

func TestMap(t *testing.T) {
	arr := Map(a, func(i int) string {
		return strconv.Itoa(i)
	})
	t.Log(arr)
}
