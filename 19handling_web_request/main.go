package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// tutorial na zpracování webového požadavku
// code posílá posílá request na URL a odpověd zobrazí jako string

const url = "https://gcc.gnu.org/onlinedocs/cpp/Header-Files.html"

func main() {
	fmt.Println("Welcome in web handle")

	response, err := http.Get(url)

	if err != nil {
		panic (err)
	}

	fmt.Printf("Responese is of type: %T\n", response) // vypíše typ odpovědi 

	defer response.Body.Close() // uzavření těla odpovědi na konci programu pomocí defer

	databytes, err := ioutil.ReadAll(response.Body) // přečte tělo odpovědi a uloží v bytech do proměnné databytes

	if err != nil { // pokud se tělo nepodaří přečíst změní hodnotu err z nil na error a shodí program
		panic(err)
	}

	content := string(databytes) // překládá tělo odpovědi v bytech na string, který uloží do proměnné content

	fmt.Println(content) // zobrazí výslednou odpověd ve formatu stringu 
}