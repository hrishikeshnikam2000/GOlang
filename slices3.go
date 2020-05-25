package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("my slice is :", mySlice)

	// 	newSlice := mySlice[2:3]
	// 	fmt.Println("new slice is :", newSlice)

	// 	newSlice[0] = 100
	// 	fmt.Println("my slice is :", mySlice)
	// 	fmt.Println("new slice is :", newSlice)
	//

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println("My Slice Is : ", mySlice)

}
