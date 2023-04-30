package main
import "fmt"

func main() {
	i:=100

	switch i {
		//fallthrough keyword will force it to execute next case as well (Opposite to break) -- here we don't need to mention break :)
		case 10:
			fmt.Println("i is 10")
		case 20:
			fmt.Println("i is 10")
		default:
			fmt.Println("i is neither 10 nor 20")
	}
}