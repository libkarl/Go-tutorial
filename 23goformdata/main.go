package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func main() {
	fmt.Println("Welcome in Go form data-tutorial")
	PerformPostFormRequest()
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform" //tímhle řeknu, že to posílám na app.POST na frontendu

	//form data 
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("nation", "Czech")

	response, err := http.PostForm(myurl, data) // u post form musím definovat jen co a kam to posílám, nemusím definovat datový typ jako u Post()
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Response is: ", string(content))



}

