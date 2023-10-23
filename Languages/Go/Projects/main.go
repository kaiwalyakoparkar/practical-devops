package main

import (
	"fmt"
	"log"
	routers "github.com/kaiwalyakoparkar/practical-devops/tree/main/Languages/Go/Projects/router"
)

func main() {
	fmt.Println("✅ API Started")
	r := routers.Router()
	
	fmt.Println("🔒 Env loaded")

	fmt.Println("🚚 Starting server")

	log.Fatal(r.Run("localhost:8080"))
}