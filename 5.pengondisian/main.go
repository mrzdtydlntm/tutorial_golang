package main

import "fmt"

func main() {
	/*
		Menggunakan if else
	*/
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
	/*
		Menggunakan temporary variables
	*/
	var points = 8840.0

	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
	/*
		Menggunakan switch case
	*/
	var pointt = 6

	switch pointt {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	/*
		Menggunakan fallthrough
	*/
	var pointts = 6

	switch {
	case pointts == 8:
		fmt.Println("perfect")
	case (pointts < 8) && (pointts > 3):
		fmt.Println("awesome")
		fallthrough
	case pointts < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
