package ex00

import (
	"unsafe"
	"errors"
)

func GetElement(arr []int, idx int) (int, error) {
	if idx < 0 {
		return 0, errors.New("Index cannot be negative")
	}
	if idx >= len(arr) {
		return 0, errors.New("Index out of bounds")
	}
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&(arr[0]))) + uintptr(idx)*unsafe.Sizeof(int(0)))), nil
}
