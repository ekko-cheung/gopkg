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
