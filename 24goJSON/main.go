package main

import (
	"encoding/json"
	"fmt"
)

// řeší konvertování jakýchkoli dat na validní JSON

type course struct {
	Name string `json:"coursename"` // říkám, že Name ve struktuře nechci volat jako name, ale jako course name
	Price int
	Platform string `json:"website"`
	Password string `json:"-"` // pomlčka řekne, že nechci aby bylo heslo viditelné
	Tags []string `json:"tags,omitempty"`// omitempty říká, že pokud je tato položka null vůbec se nebude tag vypisovat
}


func main() {
	fmt.Println("Welcome in GoJson-tutorial")
	EncodeJson()
}

func EncodeJson(){
	lcoCourses := []course  {// bude tam nějaký slice z courses
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bbc123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	} 

	// cíl je package ty data jako JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t") // vytáhnu z knihovny json metodu Marshal do, které pošlu ta data
	if err != nil{ // prefix zlepšuje čitelnost dat protože dělá indent podle typu 
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}