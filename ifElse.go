package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Hour()
	if now <= 12 {
		fmt.Println("Good morning!")
	} else if now <= 16 {
		fmt.Println("Good afternon")
	} else if now <= 22 {
		fmt.Println("good Evening!")

	} else {
		fmt.Println("slep now!")
	}

	fmt.Println("have a nice day!!!!")
}
