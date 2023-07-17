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
