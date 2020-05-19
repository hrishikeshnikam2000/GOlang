package main

import "fmt"

func main() {
	var originalCount = 10
	fmt.Println("I started with", originalCount, "apples")
	var eatenCount = 4
	fmt.Println("some jerk ate", eatenCount, "apples")
	fmt.Println("there are ", originalCount-eatenCount, "apples left.")

}
