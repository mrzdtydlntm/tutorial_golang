package main

import (
	"fmt"
	"time"
)

func main() {
	//afterfunc
	/*
		var ch = make(chan bool)

		time.AfterFunc(4*time.Second, func() {
			fmt.Println("expired")
			ch <- true
		})

		fmt.Println("start")
		<-ch
		fmt.Println("finish")
	*/
	//after
	<-time.After(4 * time.Second)
	fmt.Println("expired")
}
