package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

type Company struct {
	Name      string
	Employees []Employee
}

func Walk(x any, fn func(string)) {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			Walk(value.Field(i).Interface(), fn)

		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			Walk(value.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			Walk(value.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			val, ok := value.Recv()
			if !ok {
				break
			}
			Walk(val.Interface(), fn)
		}
	case reflect.Func:
		if value.Type().NumIn() == 0 {
			results := value.Call(nil)
			for _, result := range results {
				Walk(result.Interface(), fn)
			}
		}

	}

}

func main() {
	company := Company{
		Name: "PlanetScale",
		Employees: []Employee{
			{
				Name: "Bob",
				Age:  25,
				Address: Address{
					City:    "Buffalo",
					Country: "USA",
				},
			},
			{
				Name: "Pratik",
				Age:  23,
				Address: Address{
					City:    "Pune",
					Country: "India",
				},
			},
		},
	}
	animals := map[string]string{
		"Cow":   "Moo",
		"Sheep": "Baa",
	}
	Walk(company, func(s string) {
		fmt.Println(s)
	})
	Walk(animals, func(s string) {
		fmt.Println(s)
	})

}
