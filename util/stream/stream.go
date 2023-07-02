//go:build go1.18 || 1.19 || 1.20

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

type Filter[T any] func(T) bool

// Compare a and b
//
// if a > b return number > 0
//
// if a == b return 0
//
// if a < b return number < 0
type Compare[T any] func(a T, b T) int

type ForEach[T any] func(T)

type Stream[T any] interface {
	Filter(Filter[T]) Stream[T]
	ForEach(ForEach[T])
	Reverse() Stream[T]
	Count() int
	ToSlice() []T
	Limit(int) Stream[T]
	Min(Compare[T]) T
	Max(Compare[T]) T
	Skip(int) Stream[T]
	Sorted(Compare[T]) Stream[T]
}

type streamImpl[T any] struct {
	val []T
	fs  []func()
}

func New[T any](arr []T) Stream[T] {
	val := make([]T, len(arr))
	copy(val, arr)
	s := streamImpl[T]{
		fs:  make([]func(), 0),
		val: val,
	}

	return &s
}

func (s *streamImpl[T]) notEmpty() bool {
	return len(s.val) != 0
}

func (s *streamImpl[T]) Filter(f Filter[T]) Stream[T] {
	s.fs = append(s.fs, func() {
		if s.notEmpty() {
			val := make([]T, 0, len(s.val))
			for i := range s.val {
				v := s.val[i]
				if f(v) {
					val = append(val, v)
				}
			}
			s.val = val
		}
	})

	return s
}

func (s *streamImpl[T]) ForEach(f ForEach[T]) {
	s.execute()
	if s.notEmpty() {
		for i := range s.val {
			f(s.val[i])
		}
	}
	s.free()
}

func (s *streamImpl[T]) Reverse() Stream[T] {
	s.fs = append(s.fs, func() {
		l := len(s.val)
		if s.notEmpty() && l != 1 {
			head := 0
			tail := l - 1
			for head < tail {
				tem := s.val[head]
				s.val[head] = s.val[tail]
				s.val[tail] = tem
				head++
				tail--
			}
		}
	})

	return s
}

func (s *streamImpl[T]) Count() int {
	s.execute()
	count := len(s.val)
	s.free()

	return count
}

func (s *streamImpl[T]) ToSlice() []T {
	s.execute()
	val := s.val
	s.free()

	return val
}

func (s *streamImpl[T]) free() {
	s.val = nil
	s.fs = nil
}

func (s *streamImpl[T]) Limit(i int) Stream[T] {
	s.fs = append(s.fs, func() {
		if s.notEmpty() {
			l := len(s.val)
			if i < l {
				s.val = s.val[:i]
			}
		}
	})

	return s
}

func (s *streamImpl[T]) Skip(i int) Stream[T] {
	s.fs = append(s.fs, func() {
		if s.notEmpty() {
			l := len(s.val)
			if i < l {
				s.val = s.val[i:]
			}
		}
	})

	return s
}

func (s *streamImpl[T]) Max(f Compare[T]) T {
	s.execute()

	if s.notEmpty() {
		max := s.val[0]
		for i := 1; i < len(s.val); i++ {
			if f(s.val[i], max) > 0 {
				max = s.val[i]
			}
		}

		return max
	}
	s.free()

	return *new(T)
}

func (s *streamImpl[T]) Min(f Compare[T]) T {
	s.execute()
	if s.notEmpty() {
		min := s.val[0]
		for i := 1; i < len(s.val); i++ {
			if f(s.val[i], min) < 0 {
				min = s.val[i]
			}
		}
	}

	s.free()

	return *new(T)
}

func (s *streamImpl[T]) Sorted(f Compare[T]) Stream[T] {
	s.fs = append(s.fs, func() {
		l := len(s.val)
		if l > 1 {
			for i := l - 1; i > 0; i-- {
				flag := 0
				for j := 0; j < i; j++ {
					if f(s.val[j], s.val[j+1]) > 0 {
						temp := s.val[j]
						s.val[j] = s.val[j+1]
						s.val[j+1] = temp
						flag = 1
					}
				}
				if flag == 0 {
					break
				}
			}
		}
	})

	return s
}

func (s *streamImpl[T]) execute() {
	for i := range s.fs {
		s.fs[i]()
	}
}

func Map[T, R any](source []T, f func(T) R) []R {
	arr := make([]R, len(source))
	for i := range source {
		arr[i] = f(source[i])
	}

	return arr
}
