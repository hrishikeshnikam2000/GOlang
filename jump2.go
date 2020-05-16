package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if 1 == 4 {
			return
		}

		fmt.Println("iterating", i)
		if 1 == 6 {
			continue
		}
		if i == 8 {
			break
		}
	}

}
