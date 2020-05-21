 package main

import "fmt"

func main() {
	fmt.Println("addition is", add(30, 20))
	fmt.Println("addition is", subt(30, 10))
	fmt.Println("multiplication is", multi(20, 30))
	fmt.Println("division is", divi(30, 2))
	fmt.Println("square is:-", square(30))
	fmt.Println("square and addition is", complexOperations(30, 25))
	fmt.Println(newFunc(30, 20))
}
func add(num1, num2 int) float64 {
	return float64(num1 + num2)
}
func subt(num1, num2 int) float64 {
	return float64(num1 - num2)
}
func multi(num1, num2 int) float64 {
	return float64(num1 * num2)
}
func divi(num1, num2 int) float64 {
	return float64(num1 / num2)
}
func square(num1 int) int {
	return num1 * num1
}
func complexOperations(num1, num2 int) float64 {
	return float64(square(int(add(num1, num2))))
}
func newFunc(num1, num2 int) (float64, float64, float64, float64, int) {
	return float64(add(num1, num2)), float64(subt(num1, num2)), float64(multi(num1, num2)), float64(divi(num1, num2)), int(square(num1))
}
