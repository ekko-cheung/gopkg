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

func TestLinkedHashMapInsert(t *testing.T) {
	m := NewLinkedHashMap[string, int]()
	m.Insert("a", 1)
	m.Insert("b", 2)
}

func TestLinkedHashMapGet(t *testing.T) {
	m := NewLinkedHashMap[string, int]()
	m.Insert("a", 1)
	t.Log(m.Get("a"))
}

func TestLinkedHashMapDel(t *testing.T) {
	m := NewLinkedHashMap[string, int]()
	m.Insert("a", 1)
	m.Delete("a")
}

func TestLinkedHashMapForEach(t *testing.T) {
	m := NewLinkedHashMap[string, int]()
	m.Insert("a", 1)
	m.Insert("b", 2)
	m.Insert("c", 3)
	m.Delete("b")
	m.ForEach(func(key string, value int) {
		t.Logf("%s=%d\n", key, value)
	})
}