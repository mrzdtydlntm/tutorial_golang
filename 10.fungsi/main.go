package main

import (
	"fmt"
	"math"
)

//Inisialisasi fungsi
/*
import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"John", "Wick"}
	printMessage("halo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
*/

//fungsi dengan return value
/*
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
*/

//fungsi dengan return untuk menghentikan proses (tanpa nilai atau value di return)
// import "fmt"
//
// func main() {
// divideNumber(10, 2)
// divideNumber(4, 0)
// divideNumber(8, -4)
// }
//
// func divideNumber(m, n int) {
// if n == 0 {
// fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// return
// }
//
// var res = m / n
// fmt.Printf("%d / %d = %d\n", m, n, res)
// }

//Fungsi multiple return
func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

func main() {
	var diameter float64 = 15
	area, circumference := calculate(diameter)
	fmt.Printf("luas lingkaran \t \t: %.2f\n", area)
	fmt.Printf("keliling lingkaran \t: %.2f\n", circumference)
}
