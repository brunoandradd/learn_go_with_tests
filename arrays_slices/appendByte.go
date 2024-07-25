package arraysslices

import "fmt"

var p = fmt.Println

func AppendBytes(slice []byte, data ...byte) []byte {
	return append(slice, data...)
}

func AppendBytesWithCopy(slice []byte, data ...byte) []byte {
	length := len(slice)
	newCapacity := length + len(data)
	if newCapacity > cap(slice) {
		newSlice := make([]byte, newCapacity)
		copy(newSlice, slice)
		slice = newSlice
	}
	copy(slice[length:], data)
	return slice
}
