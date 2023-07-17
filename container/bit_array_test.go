package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitArray(t *testing.T) {
	b := NewBitArray()
	assertions := assert.New(t)
	assertions.EqualValues(b.SetBit(1, 1), 0)
	assertions.EqualValues(b.GetBit(1), 1)
	assertions.EqualValues(b.SetBit(1, 0), 1)
	assertions.EqualValues(b.GetBit(1), 0)
}
