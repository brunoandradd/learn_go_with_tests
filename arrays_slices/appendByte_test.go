package arraysslices

import (
	"reflect"
	"testing"
)

func TestAppendByte(t *testing.T) {
	slice := []byte{1, 2, 3}
	got := AppendBytes(slice, 4, 5, 6)
	want := []byte{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
