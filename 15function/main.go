package main 

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	
	result := adder(3, 5)

	fmt.Println("Result is: ", result)
}

func adder(a int, b int) int {
	return a + b

}