package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//random seed
	/*
		rand.Seed(10)
		fmt.Println("random ke-1:", rand.Int()) // 5221277731205826435
		fmt.Println("random ke-2:", rand.Int()) // 3852159813000522384
		fmt.Println("random ke-3:", rand.Int()) // 8532807521486154107
	*/

	//unique seed
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
}
