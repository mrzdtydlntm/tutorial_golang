package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name":"john wick", "Age":27}`
	var jsonData = []byte(jsonString)

	//decode to map[string]interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("User :", data1["Name"])
	fmt.Println("Age :", data1["Age"])

	//decode to interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age :", decodedData["Age"])
}
