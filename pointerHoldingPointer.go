package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	fmt.Println("value of a :", a)
	fmt.Println("memory address of a : ", &a)
	var p *int = &a
	fmt.Println("value of p :", p)
	fmt.Println("deref of p :", *p) //value of a i.e derefencing of a pointer
	fmt.Println("type of p :", reflect.TypeOf(p))
	fmt.Println("memory address of p :", &p)
	var pointerToPointer **int = &p
	fmt.Println("value of p2p :", pointerToPointer)
	fmt.Println("deref of p2p :", *pointerToPointer) //value of a i.e derefencing of a pointer
	fmt.Println("type of p2p :", reflect.TypeOf(pointerToPointer))
	fmt.Println("memory address of p2p :", &pointerToPointer)
}
