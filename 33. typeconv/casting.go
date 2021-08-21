package main

import "fmt"

func main() {
	//konversi data dengan casting
	var a float64 = float64(24)
	fmt.Println(a)
	fmt.Printf("%T\n\n", a)

	var b int32 = int32(24)
	fmt.Println(b)
	fmt.Printf("%T\n\n", b)

	//casting string to byte
	var text1 = "halo"
	var c = []byte(text1)

	fmt.Printf("%d %d %d %d \n\n", c[0], c[1], c[2], c[3])

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
	// halo

	var d int64 = int64('h')
	fmt.Println(d) // 104

	// var e string = string(104)
	r := rune('h')
	fmt.Println(r, string(r))
	// fmt.Println(e) // h
}
