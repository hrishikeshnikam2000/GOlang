package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"

)

func main() {
	fmt.Println("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade,err := strconv.ParseFloat(input,64)
	if err != nil{
		log.Fatal(err)
	}
	var status string
	if grade >= 60{
		status = "failing"
	}else {
		status = "passing"
	}
	fmt.Println("A grade of", grade, "is", status)
	

}