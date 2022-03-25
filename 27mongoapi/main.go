package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hiteshchoudhary/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000 ...")
}


// interface funguje tak, že očekává nějakou funkci s přesným názvem v něm definovanou
// 