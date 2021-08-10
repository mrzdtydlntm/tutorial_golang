package main

import "fmt"

func main() {
	// Inisialisasi array cara 1
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Printf("%T\n", names)
	fmt.Println()

	// Inisialisasi array cara 2
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println()

	//Inisialisasi array tanpa mengetahui banyak data
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))
	fmt.Println()

	//Array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
	fmt.Println()

	//Perulangan array dengan for
	var fruitss = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruitss); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruitss[i])
	}
	fmt.Println()

	//Perulangan dengan for range
	for i, fruit := range fruitss {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	fmt.Println()

	//For loop dengan variabel i tidak akan digunakan
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
	fmt.Println()

	//Array menggunakan make
	var fruitses = make([]string, 2)
	fruitses[0] = "apple"
	fruitses[1] = "manggo"
	fmt.Println(fruitses) // [apple manggo]
}
