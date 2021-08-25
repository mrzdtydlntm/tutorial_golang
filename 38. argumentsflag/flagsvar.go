package main

import (
	"flag"
	"fmt"
)

func main() {
	// cara ke-1
	var data1 = flag.String("name", "anonymous", "type your name")
	// cara ke-2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	flag.Parse()

	fmt.Println(*data1)
	fmt.Println(data2)
}

// $ go build flagsvar.go
// $ ./flagsvar --help
