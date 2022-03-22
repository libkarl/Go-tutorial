package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome in slices")

	var fruitList = []string{"Apple", "Bannan", "Jahoda"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Bako")

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)   

	fruitList = append(fruitList[0:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 321) // přidá nakonec highScores tato čísla
	sort.Ints(highScores) // setřídí hodnoty v highScores od nejmenší po největší
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	// how to remove a value frome slices based on index
	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) 
	fmt.Println(courses) // vynechá hodnotu na indexu 2
}