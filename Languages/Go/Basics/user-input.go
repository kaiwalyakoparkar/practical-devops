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
}