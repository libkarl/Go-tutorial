package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

langulages := make(map[string]string) // klíč v map uloženém v languages bude strring a value bude string

langulages["JS"] = "Java Script" // položka s klíčem JS a hodnotou Javascript
langulages["SV"] = "Swift"
langulages["RB"] = "Ruby"
langulages["PT"] = "Python"

fmt.Println(langulages)

delete(langulages, "SV")

fmt.Print(langulages)

// loops in go lang

for key, value := range langulages {
	fmt.Printf("For key %v, value is %v\n", key, value)
}

}