package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings.Contains()
	/*
		var isExists = strings.Contains("John Wick", "Wick")
		fmt.Println(isExists) //true
	*/

	//strings.HasPrefix()
	/*
		var isPrefix1 = strings.HasPrefix("John Wick", "Jo")
		fmt.Println(isPrefix1)

		var isPrefix2 = strings.HasPrefix("John Wick", "Wi")
		fmt.Println(isPrefix2)
	*/

	//strings.HasSuffix()
	/*
		var isSuffix1 = strings.HasSuffix("John Wick", "ic")
		fmt.Println(isSuffix1)

		var isSuffix2 = strings.HasSuffix("John Wick", "ck")
		fmt.Println(isSuffix2)
	*/

	//strings.Count()
	/*
		var howMany = strings.Count("Ethan Hunt", "t")
		fmt.Println(howMany) //2
	*/

	//strings.Index()
	/*
		var index1 = strings.Index("Ethan Hunt", "ha")
		fmt.Println(index1)

		var index2 = strings.Index("Ethan Hunt", "n")
		fmt.Println(index2)
	*/

	//strings.Replace()
	/*
		var text = "banana"
		var find = "a"
		var replaceWith = "o"

		var newText1 = strings.Replace(text, find, replaceWith, 1)
		fmt.Println(newText1)

		var newText2 = strings.Replace(text, find, replaceWith, 2)
		fmt.Println(newText2)

		var newText3 = strings.Replace(text, find, replaceWith, -1)
		fmt.Println(newText3)
	*/

	//strings.Repeat()
	/*
		var str = strings.Repeat("na", 4)
		fmt.Println(str)
	*/

	//strings.Split()
	/*
		var string1 = strings.Split("The Dark Knight", " ")
		fmt.Println(string1)

		var string2 = strings.Split("Batman", "")
		fmt.Println(string2)
	*/

	//strings.Join()
	/*
		var data = []string{"banana", "papaya", "tomato"}
		var str = strings.Join(data, "-")
		fmt.Println(str)
	*/

	//strings.ToLower() and strings.ToUpper()
	var str_low = strings.ToLower("aLaY")
	fmt.Println(str_low)

	var str_up = strings.ToUpper("eat!")
	fmt.Println(str_up)
}
