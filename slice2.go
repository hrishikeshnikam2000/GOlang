package main

import "fmt"

func main() {
	// var mySlice []int
	// fmt.Println(mySlice == nil)
	mySlice := make([]int, 5, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("My Slice Is:", mySlice)
	fmt.Println("My Slice Length  Is:", len(mySlice))
	fmt.Println("My Slice Capacity  Is:", cap(mySlice))
	fmt.Println("Address of the first element  Is:", &(mySlice[0]))

	mySlice = append(mySlice, 60, 70)
	fmt.Println("My Slice Is:", mySlice)
	fmt.Println("My Slice Length  Is:", len(mySlice))
	fmt.Println("My Slice Capacity  Is:", cap(mySlice))
	fmt.Println("Address of the first element  Is:", &(mySlice[0]))

	mySlice = append(mySlice, 80, 90)
	fmt.Println("My Slice Is:", mySlice)
	fmt.Println("My Slice Length  Is:", len(mySlice))
	fmt.Println("My Slice Capacity  Is:", cap(mySlice))
	fmt.Println("Address of the first element  Is:", &(mySlice[0]))

	newSlice := append(mySlice, 100, 110)
	fmt.Println("My new Slice Length  Is:", len(newSlice))
	fmt.Println("My new Slice Capacity  Is:", cap(newSlice))
	fmt.Println("My new Slice Is:", newSlice)
	fmt.Println("Address of the first element  Is:", &(newSlice[0]))

	fmt.Println("My Slice Is:", mySlice)
	fmt.Println("My Slice Length  Is:", len(mySlice))
	fmt.Println("My Slice Capacity  Is:", cap(mySlice))
	fmt.Println("Address of the first element  Is:", &(mySlice[0]))
}
