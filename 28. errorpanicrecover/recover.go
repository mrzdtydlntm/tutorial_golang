package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

/*
func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		// fmt.Println("end")
	}
}
*/

//using iife
/*
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()
	panic("some error happen")
}
*/
func main() {
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {

		func() {

			// recover untuk IIFE dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
		}()

	}
}
