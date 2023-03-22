//go:build go1.18 || 1.19 || 1.20

package stream

type Filter[T any] func(T) bool

type Compare[T any] func(T, T) int

type ForEach[T any] func(T)

type Stream[T any] interface {
	Filter(Filter[T]) Stream[T]
	ForEach(ForEach[T])
	Reverse() Stream[T]
	Count() int
	ToArray() []T
	Limit(int) Stream[T]
	Min(Compare[T]) T
	Max(Compare[T]) T
	Skip(int) Stream[T]
	Sorted(Compare[T]) Stream[T]
}

type streamImpl[T any] struct {
	val     []T
	source  []T
	isFirst bool
	f       func()
}

func New[T any](arr []T) Stream[T] {
	s := streamImpl[T]{
		source:  arr,
		f:       func() {},
		isFirst: true,
	}

	return &s
}

func (s *streamImpl[T]) Filter(f Filter[T]) Stream[T] {
	arr := s.arr()
	if len(arr) == 0 {
		return s
	}
	fun := s.f
	s.f = func() {
		fun()
		arr := s.arr()
		if len(arr) == 0 {
			return
		}
		temp := make([]T, 0, len(arr))
		for _, item := range arr {
			if f(item) {
				temp = append(temp, item)
			}
		}
		s.val = temp
		s.isFirst = false
	}

	return s
}

func (s *streamImpl[T]) ForEach(f ForEach[T]) {
	arr := s.arr()
	if len(arr) == 0 {
		return
	}
	s.f()
	arr = s.arr()
	for i := range arr {
		f(arr[i])
	}
}

func (s *streamImpl[T]) Reverse() Stream[T] {
	arr := s.arr()
	if len(arr) == 0 || len(arr) == 1 {
		return s
	}
	fun := s.f
	s.f = func() {
		fun()
		if s.val == nil {
			val := make([]T, len(s.source))
			copy(val, s.source)
			s.val = val
			s.isFirst = false
		}
		valLen := len(s.val)
		if valLen == 0 || valLen == 1 {
			return
		}
		head := 0
		tail := valLen - 1
		for head <= tail {
			temp := s.val[head]
			s.val[head] = s.val[tail]
			s.val[tail] = temp
			head++
			tail--
		}
	}
	return s
}

func (s *streamImpl[T]) Count() int {
	s.f()
	count := len(s.arr())
	s.free()

	return count
}

func (s *streamImpl[T]) ToArray() []T {
	s.f()
	val := s.arr()
	s.free()

	return val
}

func (s *streamImpl[T]) free() {
	s.f = nil
	s.val = nil
	s.source = nil
}

func (s *streamImpl[T]) Limit(i int) Stream[T] {
	valLen := len(s.arr())
	if valLen == 0 || i >= valLen {
		return s
	}
	fun := s.f
	s.f = func() {
		fun()
		arr := s.arr()
		valLen := len(arr)
		if valLen == 0 || i+1 > valLen {
			return
		}
		s.val = arr[:i]
	}

	return s
}

func (s *streamImpl[T]) Skip(i int) Stream[T] {
	valLen := len(s.arr())
	if valLen == 0 || i >= valLen {
		return s
	}
	fun := s.f
	s.f = func() {
		fun()
		arr := s.arr()
		valLen := len(arr)
		if valLen == 0 || i+1 > valLen {
			return
		}
		s.val = arr[i:]
	}

	return s
}

func (s *streamImpl[T]) Max(f Compare[T]) T {
	s.f()
	arr := s.arr()
	if len(arr) == 0 {
		t := new(T)
		return *t
	}
	val := arr[0]
	for i := 1; i < len(arr); i++ {
		if f(val, arr[i]) < 0 {
			val = arr[i]
		}
	}
	s.free()

	return val
}

func (s *streamImpl[T]) Min(f Compare[T]) T {
	s.f()
	arr := s.arr()
	if len(arr) == 0 {
		t := new(T)
		return *t
	}
	val := arr[0]
	for i := 1; i < len(arr); i++ {
		if f(val, arr[i]) >= 0 {
			val = arr[i]
		}
	}
	s.free()

	return val
}

func (s *streamImpl[T]) Sorted(f Compare[T]) Stream[T] {
	valLen := len(s.arr())
	if valLen == 0 || valLen == 1 {
		return s
	}
	fun := s.f
	s.f = func() {
		fun()
		if s.val == nil {
			val := make([]T, len(s.source))
			copy(val, s.source)
			s.val = val
			s.isFirst = false
		}
		valLen := len(s.val)
		if valLen == 0 || valLen == 1 {
			return
		}

		for i := valLen - 1; i > 0; i-- {
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

	return s
}

func (s *streamImpl[T]) arr() []T {
	if s.isFirst {
		return s.source
	}

	return s.val
}

func Map[T, R any](source []T, f func(T) R) []R {
	arr := make([]R, len(source))
	for i := range source {
		arr[i] = f(source[i])
	}

	return arr
}
