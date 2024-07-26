package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

type Dollar int

var ErrNotFound = errors.New("word not found")

func (d Dictionary) Search(key string) (string, error) {
	definition, keyFound := d[key]
	if keyFound {
		return definition, nil
	}
	fmt.Printf("definition %v", definition)

	return definition, ErrNotFound
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dollar) ToSring() string {
	return fmt.Sprintf("$ %d", d)
}
