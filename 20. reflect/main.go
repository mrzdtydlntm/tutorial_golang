package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

//akses informasi property objek
/*
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s) //mengambil value dari student

	if reflectValue.Kind() == reflect.Ptr { //jika true, berarti reflectValue berupa pointer
		reflectValue = reflectValue.Elem() //mengambil objek reflect aslinya dengan .Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ { //perulangan sesuai banyaknya data
		fmt.Println("nama      :", reflectType.Field(i).Name)         //mengembalikan nama property
		fmt.Println("tipe data :", reflectType.Field(i).Type)         //mengembalikan tipe data property
		fmt.Println("nilai     :", reflectValue.Field(i).Interface()) //mengembalikan nilai property dalam bentuk interface()
		fmt.Println("")
	}
}
*/
func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	//mencari tipe data dan value dengan reflect
	/*
		var number = 23
		var reflectValue = reflect.ValueOf(number)

		fmt.Println("tipe  variabel :", reflectValue.Type())

		if reflectValue.Kind() == reflect.Int {
			fmt.Println("nilai variabel :", reflectValue.Int())
		}
	*/

	//akses informasi property objek
	/*
		s1 := &student{Name: "wick", Grade: 2}
		s1.getPropertyInfo()
	*/

	//akses informasi method variabel objek
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)
}
