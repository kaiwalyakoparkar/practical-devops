package main
import "fmt"

func main() {
	var name string
	fmt.Print("Enter your namve: ")
	fmt.Scanf("%s", &name)//Takes 1 input from user
	fmt.Println("Hey there, ", name)
}