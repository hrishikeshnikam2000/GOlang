package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	mySlice()
	fmt.Println("enter 5 numbers in array: ")
	fmt.Scanln(&mySlice[0])
	fmt.Scanln(&mySlice[1])
	fmt.Scanln(&mySlice[2])
	fmt.Scanln(&mySlice[3])
	fmt.Scanln(&mySlice[4])

	for i := 0; i < size; i++ {
		var scanNum int
		fmt.Scanln(&scanNum)
		mySlice[i] = scanNum
	}
	fmt.Println("my slice is", mySlice)
	delSlice()
	addSlice()
	addMySLice()
}

func mySclice(sklice []int) []int {
	mySlice := make([]int, 5, 5)
	return mySlice
}

func delSlice(sklice []int) []int {
	mySlice()
	a := 0
	fmt.Println("the index of the element you want to delete:")
	fmt.Scanln(&a)
	fmt.Println("my slice is", mySlice)

	mySlice = append(mySlice[:a], mySlice[(a+1):]...)
	fmt.Println("My Slice Is : ", mySlice)
}

func addSlice(sklice []int) []int {
	mySlice()
	b := 0
	fmt.Println("the index you want to add element in")
	fmt.Scanln(&b)
	fmt.Println("you want to add the number at", b, "index")
	fmt.Println("My Slice Is : ", mySlice)
}

func addMySLice(sklice []int) []int {
	mySlice()
	b := 0
	c := 0
	fmt.Println("the number you want to enter on ", b, "index :")
	fmt.Scanln(&c)
	fmt.Println("My Slice Is : ", mySlice)

	mySlice[b] = c

	fmt.Println("the final output is:", mySlice)
}
