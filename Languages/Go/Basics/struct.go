//No inheritance in golang, no super or parent
package main
import "fmt"

func main() {
	kaiwalya := User{"Kaiwalya", "kai@mail.com", true, 20}
	fmt.Println(kaiwalya)
	fmt.Println(kaiwalya.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}