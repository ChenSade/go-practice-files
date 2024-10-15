package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// AddAlbum adds an album from JSON received in the request body.
func AddAlbum(c *gin.Context) {
	var newAlbum Album

	// bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func UpdateAlbum(c *gin.Context) {
	var updatedAlbum Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	for i, album := range albums {
		if album.ID == updatedAlbum.ID {
			pAlbum := &albums[i]
			pAlbum.Title = updatedAlbum.Title
			pAlbum.Artist = updatedAlbum.Artist
			pAlbum.Price = updatedAlbum.Price
			c.IndentedJSON(http.StatusOK, album.ID)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
