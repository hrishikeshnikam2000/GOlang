package main

import (
	"fmt"
)

func main() {
	outer := a()
	fmt.Println("break")
	outer()
}

func a() func() {
	fmt.Println("A")
	var b = func() {
		fmt.Println("B")
	}
	return b
}
