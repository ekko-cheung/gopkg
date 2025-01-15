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
