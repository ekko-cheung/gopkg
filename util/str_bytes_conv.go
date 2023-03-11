package util

import (
	"reflect"
	"unsafe"
)

func StringToSliceByte(s string) []byte {
	l := len(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  l,
		Cap:  l,
	}))
}

func SliceByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
