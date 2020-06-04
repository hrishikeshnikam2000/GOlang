package main

import "fmt"

func main() {
	func() {
		fmt.Println("inside anonymous function")

	}() //immediately invoked function experssion. IIFE
	anonyFunc()
}

func anonyFunc() {
	var number = func(s string) int {
		fmt.Println("inside non-main anonymous function.")
		return 100
	}("hello rajesh,")
	fmt.Println("return value is:", number)
}
