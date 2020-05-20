package main

import (
	"fmt"
)

var num int
var num1 int

func main() {
	num = 40
	display()
	fmt.Println("in main",num1)

}

func display() {
	var num1 int 
	num = 30
	fmt.Println("the value in main" ,num1)
	fmt.Println("in display" ,num)
}