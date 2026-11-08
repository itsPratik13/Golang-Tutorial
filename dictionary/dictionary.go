package main

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, exists := d[key]
	if !exists {
		return "", errors.New("word not found")
	}

	return definition, nil
}
func (d Dictionary) Add(key, value string) string {
	d[key] = value
	return "word is entered in dictionary"
}

func main() {
	d := Dictionary{"test": "This is test"}
	d.Search("test")
}
