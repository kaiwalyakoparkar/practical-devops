package main
import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Pointer is", ptr)
	myNumber := 23
	var ptr = &myNumber
	// *ptr = *ptr + 2 //works too, outputs 25
	fmt.Println("Pointer points to ", *ptr)
}