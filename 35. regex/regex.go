package main

import (
	"fmt"
	"regexp"
)

func main() {
	//regex.FindAllString()
	/*
		var text = "banana burger soup"
		var regex, err = regexp.Compile(`[a-z]+`)

		if err != nil {
			fmt.Println(err.Error())
		}

		var res1 = regex.FindAllString(text, 2)
		fmt.Printf("%#v \n", res1)
		// []string{"banana", "burger"}

		var res2 = regex.FindAllString(text, -1)
		fmt.Printf("%#v \n", res2)
		// []string{"banana", "burger", "soup"}
	*/

	//regex.MatchString()
	/*
		var text = "banana burger soup"
		var regex, _ = regexp.Compile(`[a-z]+`)

		var isMatch = regex.MatchString(text)
		fmt.Println(isMatch)
		//true
	*/

	//regex.FindStringIndex()
	/*
		var text = "banana burger soup"
		var regex, _ = regexp.Compile(`[a-z]+`)

		var idx = regex.FindStringIndex(text)
		fmt.Println(idx)
		//[0, 6]

		var str = text[0:6]
		fmt.Println(str)
	*/

	//regex.FindAllString()
	/*
		var text = "banana burger soup"
		var regex, _ = regexp.Compile(`[a-z]+`)

		var str1 = regex.FindAllString(text, -1)
		fmt.Println(str1)
		// ["banana", "burger", "soup"]

		var str2 = regex.FindAllString(text, 1)
		fmt.Println(str2)
		// ["banana"]
	*/

	//regex.ReplaceAllString()
	/*
		var text = "banana burger soup"
		var regex, _ = regexp.Compile(`[a-z]+`)

		var str = regex.ReplaceAllString(text, "potato")
		fmt.Println(str)
	*/

	//regex.ReplaceAllStringFunc()
	/*
		var text = "banana burger soup"
		var regex, _ = regexp.Compile(`[a-z]+`)

		var str = regex.ReplaceAllStringFunc(text, func(each string) string {
			if each == "burger" {
				return "potato"
			}
			return each
		})
		fmt.Println(str)
		// "banana potato soup"
	*/

	//regex.Split()
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-b]+`)

	var str = regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
	// []string{"", "n", "n", " ", "urger soup"}
}
