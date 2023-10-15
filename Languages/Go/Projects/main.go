package main

//Importing essential packages
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

//This is the main function
func main() {
	fmt.Println("This is an mongo-api")
	
	//Defining router
	router := gin.Default()

	//All the routes here
	router.GET("/ping", welcome)

	//Runs on 8080 port of localhost
	router.Run("localhost:8080")
}


//=======Goes in contollers=========
//Function definition
func welcome(c *gin.Context) {
	//Responding with simple json info
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong!"})
}

//===== Goes in Models =========