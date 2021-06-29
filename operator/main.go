package main

import "fmt"

func main() {
	/*
		Ini menggunakan operator aritmatika
	*/
	value := (((2+6)%3)*4 - 2) / 3
	fmt.Printf("%d\n", value)
	var isEqual = (value == 2)
	fmt.Printf("nilai %d samadengan 2 (%t)\n", value, isEqual)
	/*
		Ini menggunakan operator aritmatika
	*/
	var left = false
	var right = true
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
