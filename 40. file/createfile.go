package main

import (
	"fmt"
	"os"
)

// cara aing
/*
var pwd, _ = exec.Command("pwd").Output()
var pet = []string{string(pwd[:len(pwd)-1]), "test.txt"}
var path = strings.Join(pet, "/")

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	//deteksi file apakah sudah ada
	var _, err = os.Stat(path)

	//buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> file berhasil dibuat", path)
}

func main() {
	createFile()
}
*/

//cara tutorial

var path = "/home/CodeMirza/Golang/tutorial/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func main() {
	createFile()
}
