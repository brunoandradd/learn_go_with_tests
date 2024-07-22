package arraysslices

import "fmt"

var p = fmt.Println

func AppendBytes(slice []byte, data ...byte) []byte {
	newSlice := append(slice, data...)
	p(newSlice)
	return newSlice
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
