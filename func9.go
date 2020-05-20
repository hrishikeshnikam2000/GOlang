package main

import (
	"fmt"
	"reflect"
)

var num int = 100

func main() {
	test1 := third
	test2, test3, test4 := third(200)
	// func()

	// t2 := reflect.TypeOf(test2)
	fmt.Println("break")
	fmt.Println("Type is :", reflect.TypeOf(test1))
	fmt.Println("type2 is", reflect.TypeOf(test2))
	fmt.Println(test1(20))
	fmt.Println("test2", test2)
	fmt.Println("test3", test3)
	fmt.Println("test4", test4)

	//  func(num int) int
	// func(num3 int)(string,int,float64)
}

func first() {
	fmt.Println("inside first.")
}

func second(num int) int {
	fmt.Println("inside second", num)
	return 64
}

func third(num1 int) (string, int, float64) {
	var num5 int = 21
	fmt.Println("inside third.", num5)
	return "hello", num1, 0.64

}
