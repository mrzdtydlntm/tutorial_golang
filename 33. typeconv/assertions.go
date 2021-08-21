package main

import "fmt"

func main() {
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	/*
		fmt.Println(data["nama"].(string))
		fmt.Println(data["grade"].(int))
		fmt.Println(data["height"].(float64))
		fmt.Println(data["isMale"].(bool))
		fmt.Println(data["hobbies"].([]string))
	*/
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
