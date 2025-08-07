package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// The C parameter is from *gin.Context
// `IndentedJSON` means the response will be formatted as JSON with indentation for readability.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")        // c.Param retrieves the value of the URL parameter "id"
	for _, a := range albums { // for loop to find the correct position
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a) //if found, http200
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album // create a variable to hold entered data

	// c.BindJSON means read JSON from the request Body
	// Bind it to the newAlbum variable
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	albums = append(albums, newAlbum)            // add the new album to the albums slice
	c.IndentedJSON(http.StatusCreated, newAlbum) // http201
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)

	router.Run(":8080")
}
