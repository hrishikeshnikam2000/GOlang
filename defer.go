package main

import "fmt"

func main() {
	var number = 50
	f1()
	defer f2(number)
	number = 100
	fmt.Println("end of main, Number is:", number)
}

func f1() {
	defer fmt.Println("Hello From F1")
	fmt.Println("End of f1")
}

func f2(num int) {
	fmt.Println("end of f2,number is", num)
}
