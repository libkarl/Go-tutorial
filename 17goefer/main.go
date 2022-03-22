package main

import "fmt"


func main() {

	fmt.Println("Welcome in go defer.")
	defer fmt.Println("\nWorld") 
	fmt.Println("Welcome.")
	myDefer()
}

// defer před řádkem způsobí, že se ten řádek provede ve funkci jako poslední

func myDefer() { // defer v cyklu řadí požadavky do fronty a jakmile smyčka skončí provede je od konce
	for i:= 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
