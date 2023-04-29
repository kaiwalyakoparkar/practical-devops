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

	fmt.Print(b)
	fmt.Print("Welcome ", str, " your marks are ", i, "\n");
	fmt.Println(f)
	fmt.Printf("Your grades are %d \n", i)
	fmt.Print(name)
}