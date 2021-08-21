package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration = 5
	duration_sec := n * time.Second
	fmt.Println("time:", duration_sec)

	duration_min := n * time.Minute
	fmt.Println("time:", duration_min)

	duration_milsec := n * time.Millisecond
	fmt.Println("time:", duration_milsec)
}
