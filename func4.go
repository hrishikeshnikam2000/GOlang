package main

import "fmt"

func main() {

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}
}
