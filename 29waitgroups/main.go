package main

import (
	"fmt"
	"sync"
)




func main() {
	fmt.Println("Go wait group")
	
	mut := &sync.Mutex{} // mutex reference 
	var wg = &sync.WaitGroup{} // reference na WaitGroup v balíčku sync

    var score = []int{0} // vytvořil slice integerů, do kterého zapsal počáteční hodnotu 0
	
	wg.Add(3) // přidávám do WaitGroup 3 rutiny
	go func(wg *sync.WaitGroup, m *sync.Mutex){ // do argumentu funkce se posílá to co je ve Wait group a v Mutex z balíčku balíčku sync
		// do WaitGroup jsem vložil, že tam budou 3 rutiny, než jsem to začal předávat těm funkcím s rutinami
		fmt.Println("First Go rutine")
		mut.Lock()
		score = append(score, 1) // přidá do score 1
		mut.Unlock()
		wg.Done() // na konci každé rutiny musím říct, že tady skončila
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Second Go rutine")
		mut.Lock() // v každé rutině si zamikám pamět pro tu danou operaci a pak zase odemykám, aby nedošlo ke střetu 
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Third Go rutine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait() // zajistí, že počká s dokončením main funkce než se vrátí výsledek ze všech go rutin
	fmt.Println(score)
}

// kontrola pro střet go rutin při využívání paměti, je spuštění souboru s race flag
// go run --race main.go, pokud nevypíše nic je využití paměti v pořádku
// jemožné i využití RWMutex ( read / write ), používá se k nastavení používání zdroje tak
// že dovoluje více vláknům ve stejnou chvíli číst ze zdroje, ale jakmile přijde vlákno, které chce 
// zapisovat odstřihne všechny čtecí rutiny a nechá zapsat vlákno co jde zapisovat,
// poté opět umožní čtení všem vláknům
// to jsou situace, kdy povolíme 
// dvoum odličným threads ve stejný okamžik číst z jednoho zdroje, ale nikdy
// nesmí ve stejný okamžik zapisovat