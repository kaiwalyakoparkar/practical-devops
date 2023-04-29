//Variable Scope
package main
import "fmt"

//Global variabl
func main () {
	//Local Variable of function main 
	//Outer block
	city:="London"
	{
		country:="Uk"
	}
	fmt.Println(count)
	fmt.Println(city)
	// fmt.Println(country)//Error: can't access sten
}