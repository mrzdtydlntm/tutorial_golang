package main

import "fmt"

func main() {
	var firstName string = "John"
	var lastName string
	lastName = "Udin"
	fmt.Printf("%s %s\n", firstName, lastName)

	firstName2 := "Maman"
	lastName2 := "Abdurrahman"
	fmt.Printf("%s %s\n", firstName2, lastName2)
	fmt.Printf("%T\n", firstName2) //untuk mengecek tipe data
}
