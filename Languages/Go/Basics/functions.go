package main
import "fmt"

func main() {
	greater(3)
	result := adder(3, 5)
	fmt.Println(result)
	fmt.Println(proAdder(2,3,4))
	fmt.Println(proAdder(2,3,))
	fmt.Println(proAdder(3,4,5,6,7,8,9))
}

//No return function
func greater(a int) {
	fmt.Printf("Greater %d \n", a)
}

//func name (par type) return-signature {}
func adder(one int, two int) int {
	return one + two
}

//Any number of arguments but with int type
func proAdder(values ...int) int {
	total := 0
	for _, val :=range values {
		total += val
	}
	return total
}