package container

type HashSet[T comparable] struct {
	defaultSize int
	data        map[T]struct{}
}

func NewHashSet[T comparable](i ...int) *HashSet[T] {
	if len(i) == 0 {
		return &HashSet[T]{}
	}

	return &HashSet[T]{
		defaultSize: i[0],
	}
}

func (h *HashSet[T]) Add(t T) {
	h.ifNilCreate()
	h.data[t] = struct{}{}
}

func (h *HashSet[T]) isNil() bool {
	return h.data == nil
}

func (h *HashSet[T]) ifNilCreate() {
	if h.isNil() {
		if h.defaultSize > 0 {
			h.data = make(map[T]struct{}, h.defaultSize)
		} else {
			h.data = make(map[T]struct{})
		}
	}
}

func (h *HashSet[T]) Contains(t T) bool {
	if h.isNil() {
		return false
	}

	_, ok := h.data[t]

	return ok
}

func (h *HashSet[T]) Remove(t T) {
	if !h.isNil() {
		delete(h.data, t)
	}
}

func (h *HashSet[T]) Clear() {
	h.data = nil
}

func (h *HashSet[T]) Len() int {
	return len(h.data)
}

func (h *HashSet[T]) IsEmpty() bool {
	return h.data == nil || h.Len() == 0
}

func (h *HashSet[T]) ToSlice() []T {
	if h.isNil() {
		return []T{}
	}

	arr := make([]T, len(h.data))
	i := 0
	for t := range h.data {
		arr[i] = t
		i++
	}

	return arr
}

func (h *HashSet[T]) ForEach(f func(T)) {
	if h.isNil() {
		return
	}

	for t := range h.data {
		f(t)
	}
}
