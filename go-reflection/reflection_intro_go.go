package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	city string
	age  int
}
type Username string

func main2() {
	p := Person{name: "Pratik", city: "Pune", age: 23}
	reflect_p := reflect.ValueOf(p)
	fmt.Println(reflect_p)            // has copied exact values
	fmt.Println(reflect.TypeOf(p))    //type of value here its type is Person
	fmt.Println(reflect_p.Kind())     // whats the type its struct here
	fmt.Println(reflect_p.NumField()) // how many fields exist
	for i := 0; i < reflect_p.NumField(); i++ {
		field := reflect_p.Field(i)
		fmt.Println(field)
	}
	person := Person{
		name: "Chris",
		city: "London",
		age:  33,
	}
	value := reflect.ValueOf(&person)
	value = value.Elem() // gets value of the pointer
	fmt.Println(value)
	fmt.Println(value.Kind())

	//type and kind are different
	// var name Username="Peter"
	// fmt.Println(reflect.TypeOf(name))
	// fmt.Println(reflect.ValueOf(name))

	// usage of interface
	name := "Alice"                // type string
	x := reflect.ValueOf(name)     //type reflect.value
	original_name := x.Interface() // type =any but underlying dynamic value=string

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(original_name))

	animals := map[string]string{
		"Cow":   "Moo",
		"Sheep": "Baa",
	}
	y:=reflect.ValueOf(animals)

	fmt.Println(y.Type())

// Elem() → Dereference a pointer.

// Recv() → Receive from a channel.

// Call() → Execute a function.
}
