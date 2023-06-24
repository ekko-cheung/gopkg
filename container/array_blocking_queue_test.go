package container

import "testing"

func TestArrayBlockDeque(t *testing.T) {
	d := NewArrayBlockDeque[int](4)

	t.Run("Put", func(t *testing.T) {
		d.Put(1)
		d.Put(2)
		d.Put(3)
		d.Put(4)
	})

	t.Run("Take", func(t *testing.T) {
		t.Log(d.Take())
		t.Log(d.Take())
		t.Log(d.Take())
		t.Log(d.Take())
	})
}
