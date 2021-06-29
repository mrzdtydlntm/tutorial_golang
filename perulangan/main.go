package main

import "fmt"

func main() {
	/*
		Menggunakan for
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("Hello World ke-", i)
	}
	/*
		Menggunakan for dengan less argumen (hanya kondisi)
	*/
	var ii = 0

	for ii < 5 {
		fmt.Println("Angka", ii)
		ii++
	}
	/*
		Menggunakan for tanpa argumen apapun
	*/
	var iii = 0

	for {
		fmt.Println("Angka", iii)

		iii++
		if iii == 5 {
			break
		}
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
