package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome in JSON consuming")
	DecodeJsonMap()
}

func EncodeJson() {
	lcoCourses := []course{ // bude tam nějaký slice z courses
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bbc123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// cíl je package ty data jako JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t") // vytáhnu z knihovny json metodu Marshal do, které pošlu ta data
	if err != nil {                                               // prefix zlepšuje čitelnost dat protože dělá indent podle typu
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	
	`)
	
	var lcoCourse course // říkám, že proměnná lcoCourse bude mít struktutu course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // nechci, aby to pracovalo s kopii takže &
		fmt.Printf("%#v\n", lcoCourse)
		// %#v říkám, že chci vypsat hondotu uvnitř proměnné lcoCourse
		// i se strukturou, pokud bych dal jen  %v vypíše se bez struktury
	} else {
		fmt.Println("JSON WAS NOW VALID")
	}
	
}


// Go's terminology calls marshal the process of generating
// a JSON string from a data structure, and unmarshal the act
// of parsing JSON to a data structure.

func DecodeJsonMap () {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	
	`)

	var myOnlineData map[string]interface{}
	// vytvořená proměnná, která bude map s klíčem ve formátu string
	// nevím jaké budou hodnoty přicházející z webu, takže
	// nelze říct string nebo int, musím dát interface{} což
	// je pro více možností

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// & používám, abych zajistil, že vezme přímo ta data, ne jejich kopii
	

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
