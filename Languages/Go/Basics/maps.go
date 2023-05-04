package main
import "fmt"

func main() {
	languages:= make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	// for _, value := range languages {
	for key, value := range languages {
		fmt.Println(key, value)
	}

}