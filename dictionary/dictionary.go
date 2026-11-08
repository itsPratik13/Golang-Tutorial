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
func (d Dictionary) Add(key, value string) (string, error) {
	// _, err := d.Search(key)
	// if err==nil {
    //     return "", errors.New("key already exists")
	// }
	_,exists:=d[key]
	if exists{
		return "", errors.New("key already exists")
	}
	d[key] = value
	return "word is entered in dictionary", nil
}
func (d Dictionary) Update(key,updated_value string) (string, error){
	_,exists:=d[key]
	if !exists{
		return "",errors.New("key not found")
	}
	d[key]=updated_value
	return "key updated",nil
}
func (d Dictionary) Delete(key string) (string, error){
	_,exists:=d[key]
	if !exists{
		return "key not found",errors.New("key not found")
	}
	delete(d,key)
	return "key deleted",nil
}

func main() {
	d := Dictionary{"test": "This is test"}
	d.Search("test")
}
