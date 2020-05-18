package main

import "fmt"

func main() {

	var rad float64
	var l, b,area int
	 
	const PI float64 = 3.14 
	var areaCircle,ci float64
	
	fmt.Print("Enter radius of Circle : ")
	fmt.Scanln(&rad)
	areaCircle = PI * rad * rad
	fmt.Println("Area of Circle is : ", areaCircle)
	ci = 2 * PI * rad
	fmt.Println("Circumference of Circle is : ", ci)

	fmt.Print("Enter Length of Rectangle : ")
	fmt.Scan(&l)
	fmt.Print("Enter Breadth of Rectangle : ")
	fmt.Scan(&b)
	area = l * b
	fmt.Println("Area of Rectangle : ", area) 

	fmt.Print("Enter Length of Square : ")
	fmt.Scan(&l)
	area = l * l
	fmt.Print("Area of Square : ", area)
}
