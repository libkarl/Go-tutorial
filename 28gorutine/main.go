package main

import (
	"fmt"
	"net/http"
	"sync"
)

// získávání status kodu ze stránek urychluje pomocí go rutiny 
// aby nemusel zpomalovat main funkci pro získání odpovědi použil
// package sync up 
// waitgroup funguje tak, že neskončí funkci ve které je dokud nejsou všechny rutiny ukončeny
// (dokud nejsou odpovědi z ostatních core doručeny)

var wg sync.WaitGroup // má být implementováno jako pointer 

var signals = []string{"test"} // slice 

var mut sync.Mutex // opět by měl byt mutex zapsaný jako pointeer  

func main() {
	//go greeter("Hello")
	//greeter("World")

	webSiteList := []string{ //slice
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://fb.com",

	}

	for _, web := range webSiteList {
		
		go getStatusCode(web)
		wg.Add(1) // říkám waitgroup, že spouštím jednu go rutinu
		// pokaždé když for proběhne, přidá do wait group jednu rutinu na kterou se bude čekat
	}

	wg.Wait() // nedovolí main funkci skončit práci dokud  neskončí poslední go rutina neskončí
	fmt.Println(signals)
}


func getStatusCode (url string)  {
	defer wg.Done() // spustí se díky defer na konci celé funkce a řekne wg, že konkrétní go rutina skončila

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Jsi v prdeli")

	} else {
		mut.Lock() // zamkne pamět pro zapisování
		signals = append(signals, url)
		mut.Unlock()
		fmt.Printf("Status code: %d pro web %s\n", res.StatusCode, url)
	}
	
}
// %s říká zobraz odpovídající string z proměnné
// %d říká zobraz odpovídající argument, který musí být integer 
// mutex poskytuje zámek nad pamětí, existuje i zámek pro čtení nebo zápis
// říkám, že zamknu tuhle pamět dokud tato konkretní go rutina běží čímž nedovolím nikomu jinému tuto pamět používat dokud tato rutina nedoběhne 

// read mutex -> dovolí nějakou pamet čist, ale nedovoli do ni zapisovat 