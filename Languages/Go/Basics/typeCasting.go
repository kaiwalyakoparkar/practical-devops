package main
import ("fmt"
"strconv")

func main() {
	//int to float
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	//float to int
	var f float64 = 90.00
	var i int = int(f)
	fmt.Printf("%d\n", i)

	//int to string - strconv package
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)

	//string to int
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T", err, err)
}