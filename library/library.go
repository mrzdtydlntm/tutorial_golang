package library

import "fmt"

// func SayHello(name string) { //public
// fmt.Println("hello")
// introduce(name)
// }

// func introduce(name string) { //private
// fmt.Println("nama saya", name)
// }

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
