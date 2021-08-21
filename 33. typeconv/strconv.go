package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv string to int
	/*
		var str = "124"
		var num, err = strconv.Atoi(str)

		if err == nil {
			fmt.Println(num)
			fmt.Printf("%T\n", num)
		}
	*/

	//strconv int to string
	/*
		var num = 124
		var str = strconv.Itoa(num)

		fmt.Println(str) // "124"
		fmt.Printf("%T\n", str)
	*/

	//strconv string to int
	/*
		var str = "124"
		var num, err = strconv.ParseInt(str, 10, 64)
		if err == nil {
			fmt.Println(num)
		}

		var str_byte = "1010"
		var num_byte, err_byte = strconv.ParseInt(str_byte, 2, 8)
		if err_byte == nil {
			fmt.Println(num_byte)
		}
	*/

	//strconv int to string
	/*
		var num = int64(24)
		var str = strconv.FormatInt(num, 8)

		fmt.Println(str) // 30
	*/

	//strconv string to float
	/*
		var str = "24.12"
		var num, err = strconv.ParseFloat(str, 32)

		if err == nil {
			fmt.Println(num) // 24.1200008392334
		}
	*/

	//strconv float to string
	/*
		var num = float64(24.12)
		var str = strconv.FormatFloat(num, 'f', 6, 64)

		fmt.Println(str) // 24.120000
	*/

	//strconv string to bool
	/*
		var str = "true"
		var bul, err = strconv.ParseBool(str)

		if err == nil {
			fmt.Println(bul) // true
		}
	*/

	//strconv bool to string
	var bul = true
	var str = strconv.FormatBool(bul)
	fmt.Println(str) //true
}
