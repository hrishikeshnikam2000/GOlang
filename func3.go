package main

import "fmt"

func main() {
	var num1 int32
	var num2 int32

	fmt.Println("enter num1")
	fmt.Scanln(&num1)
	fmt.Println("Num 1 is", num1)

	fmt.Println("enter num2")
	fmt.Scanln(&num2)
	fmt.Println("Num 2 is", num2)

	sum := adding(num1, num2)
	fmt.Println("sum is ", sum)

	sum1, subt1 := addAndSub(num1, num2)
	fmt.Println("Sum is ", sum1, "Subt is :", subt1)
}

func adding(num1, num2 int32) float64 {
	fmt.Println("Inside adding function", num1, num2)
	return float64(num1 + num2)
}

func addAndSub(num1, num2 int32) (float64, float64) {
	return float64(num1 + num2), float64(num1 - num2)
}
