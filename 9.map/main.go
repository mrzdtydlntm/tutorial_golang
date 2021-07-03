package main

import "fmt"

func main() {
	//Mapping
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0
	fmt.Println()

	//Inisialisasi nil
	var data map[string]int
	// data["one"] = 1 akan muncul error!
	data = map[string]int{}
	data["one"] = 1 // tidak ada error

	// cara memasukan data kedalam map
	var chicken1 = map[string]int{"januari": 50, "februari": 40} //horizontal
	var chicken2 = map[string]int{                               // vertical
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1["januari"])
	fmt.Println(chicken2["februari"])
	fmt.Println()

	//Variabelisasi map dengan cara lain
	var _ = map[string]int{}     //tanpa nilai awal
	var _ = make(map[string]int) //menggunakan make
	var _ = *new(map[string]int) //menggunakan pointer new

	//Iterasi item map dengan for range
	var chickened = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chickened {
		fmt.Println(key, "  \t:", val)
	}
	fmt.Println()

	//Menghapus item di map
	var chickens = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chickens)) // 2
	fmt.Println(chickens)

	delete(chickens, "januari")

	fmt.Println(len(chickens)) // 1
	fmt.Println(chickens)
	fmt.Println()

	//Deteksi keberadaan key dgn item tertentu
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
	fmt.Println()

	//slice and map
	var chicks = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chick := range chicks {
		fmt.Println(chick["gender"], chick["name"])
	}
	fmt.Println()
}
