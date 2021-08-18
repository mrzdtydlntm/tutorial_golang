package main

import "fmt"

func main() {
	/*
		var secret interface{}

		secret = "ethan hunt"
		fmt.Println(secret)

		secret = []string{"apple", "manggo", "banana"}
		fmt.Println(secret)

		secret = 12.4
		fmt.Println(secret)
	*/

	//casting variabel intervace kosong ke objek pointer
	/*
		type person struct {
			name string
			age  int
		}

		var secret interface{} = &person{name: "wick", age: 27}
		var name = secret.(*person).name
		fmt.Println(name)
	*/

	//kombinasi slice, map dan interface{}
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
