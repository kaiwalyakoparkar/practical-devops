package main
import "fmt"

func main() {
	var name string
	var is_muggle bool
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)//Takes 1 input from user
	fmt.Println("Hey there, ", name)

	//Multiple inputs
	fmt.Print("Enter your name and muggle status: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println("Hello ",name," your muggle status is ", is_muggle)

	//Shorthand style
	var a string
	var b int
	fmt.Print("Enter a string and a number: ")
	//count and err are scanf return values
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("Count: ", count)
	fmt.Println("Error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}