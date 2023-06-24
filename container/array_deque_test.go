package container

import "testing"

func TestDeque(t *testing.T) {
	d := NewArrayDeque[int](6)

	t.Run("AddFirst", func(t *testing.T) {
		d.AddFirst(3)
		d.AddFirst(2)
		d.AddFirst(1)
	})

	t.Run("AddLast", func(t *testing.T) {
		d.AddLast(4)
		d.AddLast(5)
		d.AddLast(6)
	})

	t.Run("Size", func(t *testing.T) {
		t.Log(d.Size())
	})

	t.Run("GetFirst", func(t *testing.T) {
		t.Log(d.GetFirst())
	})

	t.Run("GetLast", func(t *testing.T) {
		t.Log(d.GetLast())
	})

	t.Run("RemoveFirst", func(t *testing.T) {
		t.Log(d.RemoveFirst())
	})

	t.Run("RemoveLast", func(t *testing.T) {
		t.Log(d.RemoveLast())
	})

	t.Run("String", func(t *testing.T) {
		t.Log(d.String())
	})
}
