package main

import "fmt"

func main() {

	fmt.Println("Go if else")

	loginCount := 5
	var result string

	if loginCount < 10 {
		result = "Mensi nez 10"
		
	} else if loginCount > 10 {
		result = "Vetsi nez deset"
	} else if loginCount == 10 {
		result = "loginCount odpovida podmince"
	}

	fmt.Println(result)


	 
}