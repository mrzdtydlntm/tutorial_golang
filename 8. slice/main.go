package main

import "fmt"

func main() {
	//slicing
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits) // "apple"
	fmt.Printf("%T\n", fruits)
	fmt.Println()

	//slice to array
	var newFruits = fruits[2:]
	fmt.Println(newFruits)
	fmt.Println()

	//slice is a reference type
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruits)      // [apple grape banana melon]
	fmt.Println(aFruits)     // [apple grape banana]
	fmt.Println(bFruits)     // [grape banana melon]
	fmt.Println(aaFruits)    // [grape]
	fmt.Println(baFruits)    // [grape]
	baFruits[0] = "pinnaple" // Buah "grape" diubah menjadi "pinnaple"
	fmt.Println(fruits)      // [apple pinnaple banana melon]
	fmt.Println(aFruits)     // [apple pinnaple banana]
	fmt.Println(bFruits)     // [pinnaple banana melon]
	fmt.Println(aaFruits)    // [pinnaple]
	fmt.Println(baFruits)    // [pinnaple]
	fmt.Println()

	//Fungsi cap dan len
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var cFruits = fruits[0:3]
	fmt.Println(len(cFruits)) // len: 3
	fmt.Println(cap(cFruits)) // cap: 4

	var dFruits = fruits[1:4]
	fmt.Println(len(dFruits)) // len: 3
	fmt.Println(cap(dFruits)) // cap: 3
	fmt.Println()

	//append slice
	var eFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(eFruits) // ["apple", "grape", "banana", "papaya"]
	fmt.Println()

	//copy slice
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3
	fmt.Println()

	//Akses slice dengan 3 indeks
	var fruitss = []string{"apple", "grape", "banana"}
	var fFruits = fruitss[0:2]
	var gFruits = fruitss[0:2:2]

	fmt.Println(fruitss)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruitss)) // len: 3
	fmt.Println(cap(fruitss)) // cap: 3

	fmt.Println(fFruits)      // ["apple", "grape"]
	fmt.Println(len(fFruits)) // len: 2
	fmt.Println(cap(fFruits)) // cap: 3

	fmt.Println(gFruits)      // ["apple", "grape"]
	fmt.Println(len(gFruits)) // len: 2
	fmt.Println(cap(gFruits)) // cap: 2
}
