package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Go server for front-end tutorial.")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get" // URL na kterém běží můj front end

	response, err := http.Get(myurl) //pomocí funkce http.Get() do které URL pošlu dostanu response nebo error
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content lenth is: ", response.ContentLength)

	var responseString strings.Builder // metoda, která zapisuje data v bytech
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount) // vypíše to stejné jako response.ContentLength
	fmt.Println(responseString.String()) // překlopí byty uložené v responseString na string

	//fmt.Println(string(content)) přeloží byte formát content do stringu 
	// 
}