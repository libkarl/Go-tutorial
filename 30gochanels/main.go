package main

import (
	"fmt"
	"sync"
)

// chanels jsou způsob jak mezi sebou jednotlive go rutiny
// mohou komunikovat
func main() {
	fmt.Println("Go channels")

	myCh := make(chan int, 2) // vytvoření kanálu přes make, musím definovat co budu kanálem posílat, zde int a kolik hodnot jím bude procházet, zde 1
	wg := &sync.WaitGroup{}
	//myCh <- 5 // říkám, že vkládám do kanálu 5 hodnot
	// kanál si lze představit jako box a <- tento zápis jako něco co do něj vkládá hodnoty 
	//fmt.Println(<-myCh)
	wg.Add(2) // říkám, že main func čeká na dvě rutiny
	go func (ch chan int, wg *sync.WaitGroup){ // prvni rutina je rozpovědná za obdržení hodnoty
		val, isChanelOpen := <-myCh 

		 fmt.Println(isChanelOpen) // zobrazuje stav kanálu, pokud ho uzavřu bude vracet false jinak true pro otevřený kanál
		 fmt.Println(val) // zobrazuje obdržené hodnoty
	
	
		wg.Done() // říkám, že tato rutina skončila
	}(myCh, wg) // (myCh, wg) říká, že předávám funkci stav těchto proměnných takový jaký byl před jejím spuštěním
	// pokud by to tam neměl a pod go rutinou bych obsah proměnné změnil, pracovala by rutina s tím novým stavem po ní
	go func (ch chan int, wg *sync.WaitGroup){ // druhá bude zodpovědná za předání hodnoty kanálu

		  // rutina předala kanálu číslo 5
		myCh <- 6
		myCh <- 7
		close(myCh) // uzavře kanál
		wg.Done() // říkám, že tato rutina skončila 
	}(myCh, wg)

	wg.Wait()
}
