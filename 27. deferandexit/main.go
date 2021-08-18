package main

import (
	"fmt"
	"os"
)

/*
func main() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")
}
*/

//defer pada return
/*
func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
*/

//defer pada iife
//halo 3 pada akhir fungsi
/*
func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		defer fmt.Println("halo 3")
	}

	fmt.Println("halo 2")
}
*/

//os.exit
func main() {
	defer fmt.Println("halo")
	os.Exit(10)
	fmt.Println("selamat datang")
}
