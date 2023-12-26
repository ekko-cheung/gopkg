package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenCharStr(t *testing.T) {
	s := LenCharStr(10)
	assert.Equal(t, 10, len(s))
}

func TestLenNumberStr(t *testing.T) {
	s := LenNumberStr(10)
	assert.Equal(t, 10, len(s))
}

func TestLenNumber(t *testing.T) {
	i := LenNumber(10)
	assert.NotEqual(t, i, 0)
}

func TestCamelCaseToUnderScoreCase(t *testing.T) {
	s1 := "helloWorld"
	s2 := CamelCaseToUnderScoreCase(s1)
	assert.Equal(t, s2, "hello_world")
}

func TestUnderScoreCaseToCamelCase(t *testing.T) {
	s1 := "hello_world"
	s2 := UnderScoreCaseToCamelCase(s1)
	assert.Equal(t, s2, "helloWorld")
}