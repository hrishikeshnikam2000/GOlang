package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Close resource")
		details := recover()
		fmt.Println("details:", details)
	}()
	fmt.Println("Hello")
	panic("panicking: i don't know what to do!")
	
}
