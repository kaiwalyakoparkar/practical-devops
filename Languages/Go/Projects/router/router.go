package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "github.com/kaiwalyakoparkar/practical-devops/tree/main/Languages/Go/Projects/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()

	fmt.Println("🎉 Server started")

	//Router perfoming on all records
	router.GET("/api/movies", controllers.GetAllMovies)
	router.DELETE("/api/movies", controllers.DeleteAllMovies)

	//Router perfoming on specific records
	router.POST("/api/movie", controllers.CreateMovie)
	router.PUT("/api/movie/:id", controllers.MarkedAsWatched)
	router.DELETE("/api/movie/:id", controllers.DeleteMovie)

	return router
}