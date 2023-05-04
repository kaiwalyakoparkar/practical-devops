package main
import "fmt"

func main() {
	//Compulsary to provide array size
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println(fruitList)//Prints the array
	fmt.Println(len(ist))//Prints the length of array strangly gives 4 here

	var vegList = [4]string{"potato", "mushroom", "brinjal"}
	fmt.Println(vegList)
}