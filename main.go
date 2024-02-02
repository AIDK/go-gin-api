package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album is a struct that represents a music album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums is a slice of album
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Register a new route for GET requests to /albums
	// and calls getAlbums function
	router.GET("/albums", getAlbums)
	// Register a new route for POST requests to /albums
	// and calls postAlbums function
	router.POST("/albums", postAlbums)
	// Register a new route for GET requests to /albums/:id
	// and calls getAlbumByID function
	router.GET("/albums/:id", getAlbumByID)

	// Run the server on port 8080
	log.Fatal(router.Run("localhost:3000"))
}

func getAlbums(c *gin.Context) {
	// gin.Context is a struct that holds the request and response objects
	// IndentedJSON serializes the given struct as JSON into the response body.
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {

	// Create a new newAlbum from the request's JSON
	newAlbum := album{}

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	// Return a JSON response with the new album
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {

	// Get the id from the URL
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	// Return the album as JSON
	for _, a := range albums {
		if a.ID == id {
			// Return the album as JSON
			c.JSON(http.StatusOK, a)
			return
		}
	}

	// If no album is found, return an error with a 404 status code
	// gin.H is a shortcut for map[string]interface{}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
