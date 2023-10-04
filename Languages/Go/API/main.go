package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type album struct {
	ID string `json: "id"`
	Title string `json: "title"`
	Artist string `json: "artist"`
	Price float64 `json: "price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main()  {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/health", getHealth)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, albums)
}


func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Your stuff works"})
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id you are trying to find is not available"})
}