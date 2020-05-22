package main

import "fmt"

func main() {
	number := 20
	case1(number)
	fmt.Println("Adress of number in main:", &number)
	fmt.Println("after case1, In main: ", number)
	case2(&number)
	fmt.Println("after case2, In main: ", number)
}

// pass by value

func case1(number int) {
	number = +2
	fmt.Println("In Case1, value is:", number)
	fmt.Println("In Case1, address is :", &number)
}

// pass by refrence

func case2(number *int) {
	*number = *number + 2
	fmt.Println("In Case1, value is:", number)
	fmt.Println("In Case1, address is :", &number)
}
