package main
import "fmt"

func main ()  {
	var i int = 100
	var str string = "Kaiwalya"
	var b bool = false
	var f float64 = 77.90

	//Assign value afterwards
	var name string
	name = "Koparkar"

	//Shorthand
	//1
	var v,p int = 3, 4
	fmt.Println(v,p);

	//2
	var (
		st string = "foo"
		val int = 54
	)
	fmt.Println(st, val);


	//3
	s := "Kaiwalya Koparkar"
	fmt.Println(s);

	fmt.Println(b)
	fmt.Printf("Welcome %v, your marks are %v \n", str, f);
	fmt.Printf("Your grades are %d \n", i)
	fmt.Println(name)
}