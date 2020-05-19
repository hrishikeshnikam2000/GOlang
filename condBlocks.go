package main

import "fmt"

var a = "a"
var b = "b"
var c = "c"
var d = "d"

func main() {
	a = "a"
	b = "b"
	if true {
		c = "c"
		if true {
			d = "d"
			fmt.Println("a")
			fmt.Println("b")
			fmt.Println("c")
			fmt.Println("d")
		}
		fmt.Println("a")
		fmt.Println("b")
		fmt.Println("c")
	}
	fmt.Println("a")
	fmt.Println("b")
}
