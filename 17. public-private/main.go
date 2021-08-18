package main

import (
	"fmt"
	"tutorial/library"
)

func main() {
	// library.SayHello("ethan")
	// library.introduce("ethan") //akan error karena private func

	// var s1 = library.Student{"ethan", 21}
	// fmt.Println("name ", s1.Name)
	// fmt.Println("grade", s1.Grade)

	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)

}
