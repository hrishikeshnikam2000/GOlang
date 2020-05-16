package main

import "fmt"

func main() {
	var userNumber int
test:
	fmt.Println("Enter a number:")
	fmt.Scanln(&userNumber)
	fmt.Println("The number is:", userNumber)

if userNumber  >200{
	fmt.Println("exiting$")
	return
}
	if userNumber > 100 {
		fmt.Println("number very high")
		goto test
	}
	fmt.Println("good number!")

}
