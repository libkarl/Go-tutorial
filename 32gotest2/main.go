package main

import "fmt"

type Developer struct {
	Name string
	Age string
}

// funkce má za úkol ze slice stringů vyfiltrovat unikátní hodnoty  (odstraní co je tam víckrát)
func FilterUnique (developers []Developer) []string {
	var unique []string
	check := make(map[string]int)
	for _, developer := range developers { // tady se zbavím hodnot, které jsou tam 2x
		check[developer.Name] = 1
      // pro každé jméno developera zapíše do check jméno a k němu 1
	}
	fmt.Println(check)
	for name := range check {
		unique = append(unique, name)
	}

	return unique
}

func main () {
	fmt.Println("Gettin Started with testify")

	input := []Developer { 
		{Name: "Elliot"},
		{Name: "Elliot"}, 
		{Name: "David"}, 
		{Name: "Alexander"},
		{Name: "David"}, 
		{Name: "Eva"}, 
		{Name: "Alan"}, 
	}

	result := FilterUnique(input)
	fmt.Println(result)
}

// sum := 0
//for _, value := range array {
//    sum += value
//}
// If you only need the second item in the range (the value), use
// the blank identifier, an underscore, to discard the first:

