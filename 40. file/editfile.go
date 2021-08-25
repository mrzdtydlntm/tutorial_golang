package main

import (
	"fmt"
	"os"
)

var path = "/home/CodeMirza/Golang/tutorial/test.txt"

func writeFile() {
	//buka file dengan akses Read Write
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Mari belajar golang\n")
	if isError(err) {
		return
	}

	//simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil diisi")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func main() {
	writeFile()
}
