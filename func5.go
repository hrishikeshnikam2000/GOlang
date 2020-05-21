package main

import (
	"fmt"
	"maths"
)


func main() {
	var rectangle,square,triangle,circle,side1,side2,side,edge1,edge2,edge3,rad float64
	fmt.Println("choose a shape from rectangle,square,triangle or circle")
	fmt.Scanln(&rectangle,&square,&rectangle,&circle)
	choise := "rectangle" or "triangle"
	fmt.Scanln(&choise)
	if choise = "rectangle" {
		fmt.Println("enter values of side1 side2 & ")
		fmt.Scanln(&side1,side2)
		rectangle()
	}else if choise == "square" {
		fmt.Println("enter value of side")
		fmt.Scanln(&side)
		square()
	}else if choise = "circle"{
		fmt.Println("enter value of radius n rad")
		fmt.Scanln(&rad)
		circle()
	}else{
		fmt.Println("enter value of edge1,edge2,edg3")
		fmt.Scanln(&edge1,&edge2,edge3)
		triangle()
	}
}

func rectangle() {
	return float64( 2*(side1 + side2) )
}

func square() {
	return float64(4*side)

}

func triangle() {
	return float64(edge1+edge2+edg3)
}

func circle() {
	return float64(2*3.14*rad)
}