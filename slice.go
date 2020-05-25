package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println(mySlice)

	newSlice := mySlice
	newSlice[0] = 100
	fmt.Println(mySlice)

}
