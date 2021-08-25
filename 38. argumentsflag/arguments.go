package main

import (
	"fmt"
	"os"
)

func main() {
	//penggunaan arguments

	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)

}

//penulisan di terminal
// $ go run 38.\ argumentsflag/arguments.go <arguments>
// contoh: $ go run 38.\ argumentsflag/arguments.go wow keren banget anjay
