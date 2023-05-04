//Used like a  member function for a class/struct
package main
import "fmt"

func main() {
	kaiwalya := User{"Kaiwalya", "kai@mail.com", true, 20}
	fmt.Println(kaiwalya)
	fmt.Println(kaiwalya.Email)
	kaiwalya.Score()
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}

func (u User) Score() {
	fmt.Println("Kaiwalya has 100  marks")
}