package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id" validate:"required,min=1"`
	Title  string  `json:"title" validate:"required,min=1"`
	Artist string  `json:"artist" validate:"required,min=1"`
	Price  float64 `json:"price" validate:"required,gt=0"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// validator instance
var validate *validator.Validate

func main() {
	validate = validator.New()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Validate the struct using validator tags
	if err := validate.Struct(newAlbum); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				validationErrors = append(validationErrors, err.Field()+" is required")
			case "min":
				validationErrors = append(validationErrors, err.Field()+" must not be empty")
			case "gt":
				validationErrors = append(validationErrors, err.Field()+" must be greater than 0")
			default:
				validationErrors = append(validationErrors, err.Field()+" is invalid")
			}
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	// Check if album with same ID already exists
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Album with this ID already exists"})
			return
		}
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
