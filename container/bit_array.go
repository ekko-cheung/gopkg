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

var bitsinbyte = []byte{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8}

type BitArray struct {
	bytes []byte
}

func NewBitArray() *BitArray {
	return &BitArray{
		bytes: make([]byte, 0),
	}
}

func (b *BitArray) SetBit(offset, val int) int {
	if val|1 != 1 {
		panic("val must be 1 on 0")
	}
	index := b.checkCapacityAndGetIndex(offset)
	bitval := b.bytes[index]
	bit := 7 - (offset & 0x7)
	bitval = bitval & (1 << bit)

	if val == 0 {
		b.bytes[index] = b.bytes[index] ^ (1 << bit)
	} else {
		b.bytes[index] = b.bytes[index] | (1 << bit)
	}

	if bitval > 0 {
		return 1
	}

	return 0
}

func (b *BitArray) GetBit(offset int) int {
	l := offset >> 3
	if len(b.bytes) < l+1 {
		return 0
	}

	bitval := b.bytes[l]
	bit := 7 - (offset & 0x7)
	bitval = bitval & (1 << bit)

	if bitval > 0 {
		return 1
	}

	return 0
}

func (b *BitArray) BitCount() int64 {
	var count int64
	for i := range b.bytes {
		bitval := b.bytes[i]
		count += int64(bitsinbyte[bitval])
	}

	return count
}

func (b *BitArray) checkCapacityAndGetIndex(offset int) int {
	l := (offset >> 3) + 1
	if len(b.bytes) < l {
		newBytes := make([]byte, l)
		copy(newBytes, b.bytes)
		b.bytes = newBytes
	}

	return l - 1
}
