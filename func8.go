package main

import (
	"fmt"
	"reflect"
)

func main() {
	test1 := first
	test2 := first()
	reflect.TypeOf(test1)
	reflect.TypeOf(test2)
	fmt.Println("test1", test1)
	fmt.Println("test2",test2)
}

func first(){

}
