package maps

import (
	"fmt"
)

type Dictionary map[string]string

type Dollar int

const (
	ErrNotFound         = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word exists")
	ErrWordDoesNotExist = DictionaryErr("word not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, keyFound := d[key]
	if keyFound {
		return definition, nil
	}

	return definition, ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	if d[key] != "" {
		return ErrWordExists
	}
	d[key] = value

	return nil
}

func (d Dictionary) Update(key, value string) error {
	if d[key] == "" {
		return ErrWordDoesNotExist
	}

	d[key] = value
	return nil
}

func (d Dollar) ToSring() string {
	return fmt.Sprintf("$ %d", d)
}
