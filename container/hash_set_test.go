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
