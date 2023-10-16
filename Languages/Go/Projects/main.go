package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	routers "github.com/kaiwalyakoparkar/practical-devops/tree/main/Languages/Go/Projects/router"
)

func main() {
	fmt.Println("✅ API Started")
	r := routers.Router()
	
	env_error := godotenv.Load("../.env")
	if env_error != nil {
		log.Fatal("No env file found")
	}
	fmt.Println("🔒 Env loaded")

	fmt.Println("🚚 Starting server")

	log.Fatal(r.Run(":8080"))
}