package main

import "fmt"

func main() {
  fmt.Println("Welcome to loop in golang")

  //days := []string{"Sunday", "Thuesday", "Wednesday", "Friday", "Saturday"}

  //for d :=0; d < len(days); d++ {
	//fmt.Println(days[d]) // vypisuje obsah arry days dle indexu pozice, který dostává
	//fmt.Println(len(days)) // vypisuje délku pole stringů
  //}

  //for i := range days { //range projde všechny indexy v days a postupně je při každém opakování for zapisuje do i
//	  fmt.Println(days[i])
//	  fmt.Println(i)
 // }
 //for _, day := range days {
//	 fmt.Printf("index is %v\n, and value is  ", day)
 //}

 	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco // když napočítá do dvou skočí na řádek s lco:
		}
	
		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue ++

	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")
 
}