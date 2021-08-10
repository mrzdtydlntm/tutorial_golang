package main

/*
type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	var s4 = student{name: "wayne", grade: 2}
	var s5 = student{grade: 2, name: "bruce"}
	fmt.Println("student 4 :", s4.name)
	fmt.Println("student 5 :", s5.name)

	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)
}
*/
type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	// var s1 = student{}
	// s1.name = "wick"
	// s1.age = 21
	// s1.grade = 2

	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	// Substruct
	// var p1 = person{name: "wick", age: 21}
	// var s1 = student{person: p1, grade: 2}

	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("grade :", s1.grade)

	//Anonymous struct
	// var s1 = struct {
	// person
	// grade int
	// }{}
	// s1.person = person{"wick", 21}
	// s1.grade = 2

	// fmt.Println("name  :", s1.person.name)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	//slice and struct
	// type person struct {
	// name string
	// age  int
	// }

	// var allStudents = []person{
	// {name: "Wick", age: 23},
	// {name: "Ethan", age: 23},
	// {name: "Bourne", age: 22},
	// }

	// for _, student := range allStudents {
	// fmt.Println(student.name, "age is", student.age)
	// }

	//Anonymous slice and struct
	// var student struct {
	// person
	// grade int
	// }

	// student.person = person{"wick", 21}
	// student.grade = 2

	//Nested struct
	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
}
