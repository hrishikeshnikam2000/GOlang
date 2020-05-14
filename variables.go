package main

import (
	"fmt"
	"runtime"
)

var name string

func main() {
	var rollNo int
	rollNo = 20
	rollNo = 30
	// rollNo =  "20.5"
	// var firstName = "abc"
	var (
		firstName, lastName, flag = "bunny", "bunz", false
	)
	name = "AbC"
	//shorthand variable declararion
	test := 20
	fmt.Println(rollNo)
	fmt.Println(firstName, flag, lastName)
	fmt.Println(name)
	fmt.Println(test)
	fmt.Printf("%t", test)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
