package main
import (
	"fmt"
	"reflect"
)

func main() {
	one1 := three
	one2 := three(21) 
	fmt.Println("the type one1 is:", reflect.TypeOf(one1))
	fmt.Println("the type one2 is:", reflect.TypeOf(one2))
	fmt.Println(one1(2))
	fmt.Println("one1", one1)
	fmt.Println("one2", one2)
}

func one() {
	fmt.Println("in one")
}

func two(num int)int{
	num = 10
	fmt.Println("two is", num)
	return 20
}

func three(num int)float64{
	fmt.Println("three is",num)
	return 6.4
}