package main
 
import "fmt"
 
var area int   // Variable declare outside the main function
 
func main(){
    var l,b int //Declaration of multiple Variables 
    fmt.Print("Enter Length of Rectangle : ")
    fmt.Scan(&l)
    fmt.Print("Enter Breadth of Rectangle : ")
    fmt.Scan(&b)
    area = l * b
    fmt.Println("Area of Rectangle : ",area)  //move to new line
    
    fmt.Print("Enter Length of Square : ")
    fmt.Scan(&l)
    area = l*l
    fmt.Print("Area of Square : ",area)
}