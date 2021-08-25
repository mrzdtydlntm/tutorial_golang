package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//EncodeToString() dan DecodeString()
	/*
		var data = "john wick"
		var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
		fmt.Println("encodedString:", encodedString)

		var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
		fmt.Println("decodedByte:", decodedByte)
		var decodedString = string(decodedByte)
		fmt.Println("decodedString:", decodedString)
	*/

	//Encode() dan Decode()
	/*
		var data = "john wick"
		var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(encoded, []byte(data))
		var encodedString = string(encoded)
		fmt.Println(encodedString)

		var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encodedString)))
		var _, err = base64.StdEncoding.Decode(decoded, encoded)
		if err != nil {
			fmt.Println(err.Error())
		}
		var decodedString = string(decoded)
		fmt.Println(decodedString)
	*/

	//URLEncoding
	var data = "https://kalipare.com/"
	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	fmt.Println(decodedByte)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
