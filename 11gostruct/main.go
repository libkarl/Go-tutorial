package main

import "fmt"

func main(){

	fmt.Println("Go struct tutorial")

	hitesh := User{"Hitesg", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details: %+v\n", hitesh) // při vypsání +v hodnoty nezobrazí jen hodnoty, ale i její struktury
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}