package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	key   = "abcdefgh"
	value = "helloworld"
)

func TestDesEncoding(t *testing.T) {
	s, err := DesEncoding(value, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestDesDecoding(t *testing.T) {
	s := "a//xM5D6Q3XwT14eXPsa/A=="
	desStr, err := DesDecoding(s, key)
	if err != nil {
		t.Fatal(err)
	}

	assert.New(t).Equal(value, desStr)
}
