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

import (
	"sync"
)

type BitMap struct {
	m sync.Map
}

func NewBitMap() *BitMap {
	return &BitMap{
		m: sync.Map{},
	}
}

func (b *BitMap) SetBit(key string, offset, value int) int {
	v, ok := b.m.Load(key)
	if ok {
		return v.(*BitArray).SetBit(offset, value)
	} else {
		bitArray := NewBitArray()
		bit := bitArray.SetBit(offset, value)
		b.m.Store(key, bitArray)

		return bit
	}
}

func (b *BitMap) GetBit(key string, offset int) int {
	v, ok := b.m.Load(key)
	if ok {
		return v.(*BitArray).GetBit(offset)
	}

	return 0
}

func (b *BitMap) BitCount(key string) int64 {
	v, ok := b.m.Load(key)
	if ok {
		return v.(*BitArray).BitCount()
	}

	return 0
}
