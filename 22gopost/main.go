package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	fmt.Println("Welcome in gopost Tutorial.")
	PerformPostJsonRequest(arg)
}

type Payload struct {
	CourseName string `json:"coursename"`
}

func PerformPostJsonRequest(arg string) {
	const myurl = "http://localhost:8000/post"

	payload := Payload{
		CourseName: arg,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewBuffer(data)

	println(string(data))

	// zde je nadefinované request body ve formátu json
	//když takhle použiji NewReader z package strings a zadám do závorek sintax `` mohu tam psát jakákoli data

	response, err := http.Post(myurl, "application/json", reader) // application/json je celé typ dat, která do post posílám
	// první argument post je url kam to chci poslat, další je format dat, která posílám
	// poslední je název proměnné kam jsem uložil tělo v JSON formatu (requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Počet bytů v v těle odpovědi je: ", byteCount)
	fmt.Println(responseString.String())

}
