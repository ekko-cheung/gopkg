package util

import (
	"bytes"
	"math/rand"
	"strconv"
	"sync"
)

var (
	charByte []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	charByteLen   int    = len(charByte)
	numberByte    []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	numberByteLen int    = len(numberByte)

	pool sync.Pool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
)

func LenCharStr(l int) string {
	return buildStr(l, charByteLen, charByte)
}

func LenNumberStr(l int) string {
	return buildStr(l, numberByteLen, numberByte)
}

func LenNumber(l int) int64 {
	if l >= 19 {
		l = 18
	}
	s := LenNumberStr(l)
	i, _ := strconv.ParseInt(s, 10, 64)

	return i
}

func buildStr(l, arrLen int, b []byte) string {
	buf := pool.Get().(*bytes.Buffer)
	buf.Grow(l)
	for i := 0; i < l; i++ {
		buf.WriteByte(b[rand.Intn(arrLen)])
	}
	s := buf.String()
	buf.Reset()
	pool.Put(buf)

	return s
}

func CamelCaseToUnderScoreCase(s string) string {
	buf := pool.Get().(*bytes.Buffer)
	buf.Grow(len(s))

	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >=64 && b <= 90 {			
			buf.WriteByte('_')
			b = b + 32
		}
		
		buf.WriteByte(b)
	}
	res := buf.String()
	buf.Reset()
	pool.Put(buf)

	return res
}

func UnderScoreCaseToCamelCase(s string) string {
	upperCase := false
	buf := pool.Get().(*bytes.Buffer)
	buf.Grow(len(s))

	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '_' {
			upperCase = true
			continue
		}
		if upperCase {
			b = b - 32
			upperCase = false
		}
		buf.WriteByte(b)
	}
	res := buf.String()
	buf.Reset()
	pool.Put(buf)

	return res
}