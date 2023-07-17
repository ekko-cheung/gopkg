package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap := NewBitMap()
	assertions := assert.New(t)
	assertions.EqualValues(bitMap.SetBit("a", 0, 1), 0)
	assertions.EqualValues(bitMap.SetBit("a", 1, 1), 0)
	assertions.EqualValues(bitMap.GetBit("a", 0), 1)
	assertions.EqualValues(bitMap.GetBit("a", 1), 1)
	assertions.EqualValues(bitMap.BitCount("a"), 2)
}
