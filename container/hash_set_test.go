package container

import "testing"

func TestHashSet(t *testing.T) {
	set := NewHashSet[string]()

	t.Run("test_Add", func(t *testing.T) {
		set.Add("a")
		set.Add("b")
		set.Add("c")
	})

	t.Run("test_Contains", func(t *testing.T) {
		t.Log(set.Contains("a"))
		t.Log(set.Contains("c"))
		t.Log(set.Contains("b"))
	})

	t.Run("test_Len", func(t *testing.T) {
		t.Log(set.Len())
	})

	t.Run("test_Remove", func(t *testing.T) {
		set.Remove("a")
	})

	t.Run("test_ForEach", func(t *testing.T) {
		set.ForEach(func(s string) {
			t.Log(s)
		})
	})

	t.Run("test_ToSlice", func(t *testing.T) {
		t.Log(set.ToSlice())
	})

	t.Run("test_Clear", func(t *testing.T) {
		set.Clear()
	})

	t.Run("test_IsEmpty", func(t *testing.T) {
		t.Log(set.IsEmpty())
	})
}
