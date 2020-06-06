package util

import (
	"reflect"
	"unsafe"
)

func Int2Bool(v int) bool {
	return v != 0
}

func Bool2Int(v bool) int {
	if v {
		return 1
	}
	return 0
}

func SetSliceDataLen(pSlice, pData unsafe.Pointer, len int) {
	header := (*reflect.SliceHeader)(pSlice)
	header.Cap = len
	header.Len = len
	header.Data = uintptr(pData)
}

func GetZeroTermArrayLen(p unsafe.Pointer, size uintptr) int {
	if p == nil {
		return 0
	}
	var count uintptr
	//println("size is", size)
	switch size {
	case 8:
		for {
			value := *(*int64)(unsafe.Pointer(uintptr(p) + (count * size)))
			if value == 0 {
				break
			}
			count++
		}

	case 4:
		for {
			value := *(*int32)(unsafe.Pointer(uintptr(p) + (count * size)))
			if value == 0 {
				break
			}
			count++
		}

	case 2:
		for {
			value := *(*int16)(unsafe.Pointer(uintptr(p) + (count * size)))
			if value == 0 {
				break
			}
			count++
		}

	case 1:
		for {
			value := *(*int8)(unsafe.Pointer(uintptr(p) + (count * size)))
			if value == 0 {
				break
			}
			count++
		}

	case 0:
		panic("assert failed size != 0")

	default:
		for {
			value := (*[999]byte)(unsafe.Pointer(uintptr(p) + (count * size)))
			if allZero(value, size) {
				break
			}
			count++
		}
	}
	return int(count)
}

func allZero(value *[999]byte, length uintptr) bool {
	for i := uintptr(0); i < length; i++ {
		if value[i] != byte(0) {
			return false
		}
	}
	return true
}
