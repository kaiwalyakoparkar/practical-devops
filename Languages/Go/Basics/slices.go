//More functionality then arrays [Has more power than array]
package main
import "fmt"
import "sort"

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("%T \n", fruitList)

	//Add to the list
	fruitList = append(fruitList, "Banana")
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	//Slice your slice ;)
	//[Start:End]
	fruitList = append(fruitList[1:])//Will omit Apple
	fmt.Println(fruitList)

	//Different syntax
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 433
	highScores[2] = 656
	highScores[3] = 343
	// highScores[4] = 343//Will throw error

	//append methods reallocates the memory elements
	highScores = append(highScores, 4333, 5555, 3333)

	//Sorts the slice
	sort.Ints(highScores)

	//returns boolean if the slice is sorted
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)
}