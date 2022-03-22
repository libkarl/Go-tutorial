package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome in Gofiles tutorial.")

	content := "This need to go in a file"
	file, err := os.Create("./mylcogofile.txt") //vytvoří txt soubor

	checkNilErr(err) // volání funkce pro test chyb po vytvoření souboru
	
	length, err := io.WriteString(file, content) // zapíše do file souboru obsah content
	checkNilErr(err) // volání funkce pro test chyb po vytvoření souboru 
	fmt.Println("length is: ", length)
	defer file.Close() // zavře soubor
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // přečte data ve formátu bytů a zapíše do proměnné databyte
	if err != nil { //pokud nastane problém přeruší program
		panic(err)
	}

	fmt.Println("Text data inside th file is \n", string(databyte))// převede obsah databyte na string a vypíše

}

func checkNilErr(err error) { // funkce pro testování chyb
	if err != nil { //pokud nastane problém přeruší program
		panic(err)
	} 
}


