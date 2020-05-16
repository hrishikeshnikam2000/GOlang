package main

import "fmt"

func main() {
	userNumber := addOne(9)
	fmt.Println(userNumber)

}

func addOne(number int) int {
	number = number + 1 //	number is 10
	return number
}
