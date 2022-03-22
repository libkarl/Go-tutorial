package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.geeksforgeeks.org/interfaces-in-golang/"

func main() {

	fmt.Println("Welcome in golang handling URL tutorial.")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	

	qparams := result.Query() // Query ten výsledek překlopí do lepšího formátu 

	fmt.Printf("The type of query params are: %T\n", qparams)
	// výsledkem je, že jsou tam url.values (v podstatě klasické key hodnoty)
	// ty se dají pomocí volání klíče také získat
	//  fmt.Println(qparams["course name"]) a měl tam uloženo [reactjs] 
	
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

}
