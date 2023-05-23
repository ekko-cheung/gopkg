package list

type node[T any] struct {
	value T
}

func (n *node[T]) Val() T {
	if n == nil {
		return *new(T)
	}

	return n.value
}
