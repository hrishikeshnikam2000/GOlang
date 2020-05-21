package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b string = "hello"
	fmt.Println("value of a :", a)
	fmt.Println("memory address of a : ")
	// p := &a
	p := &b
	fmt.Println("value of p :", p)
	fmt.Println("deref of p :", *p) //value of a i.e derefencing of a pointer
	fmt.Println("type of p :", reflect.TypeOf(p))
	fmt.Println("memory address of p :", &p)
}
