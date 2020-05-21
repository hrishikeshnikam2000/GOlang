package main

import "fmt"

func main() {
	test1 := empty //calling a function .test1 is a function here
	fmt.Println("test1", test1) //this will print address yad rakh.test2 is a variable 
	empty() //this is invoking  function

}

func add(num1, num2 int) int { // 2 variables of int typ returning int type
	print("addition ") // when func is called this will execute
	return num1 + num2 // when a function is invoked this will execute
}

func empty() { // empty function 
	fmt.Println("print nothing:") // will print if function is invoked
}
  