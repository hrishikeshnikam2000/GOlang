package main

import "fmt"

func main() {
	i := 5
	for {
		fmt.Println("Hello", i)
		i += 5
		if i == 25 {
			break
		}

	}

	for j := 10; j > 0; j-=2 {

		fmt.Println("hello", j)
	}

	sum := 10

	for ; sum > 0; sum-- {
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
