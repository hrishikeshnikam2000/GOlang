package main

import "fmt"

func main() {
	result := mathOperations(divide)
	fmt.Println(result)
	fmt.Println(mathOperations(multiply))
}

func divide(num1, num2 float32) float64 {
	return float64(num1 / num2)
}

func multiply(num1, num2 float32) float64 {
	return float64(num1 * num2)
}

func mathOperations(funcName func(num1, num2 float32) float64) float64 {
	result := funcName(20, 10)
	return float64(result)
}
   