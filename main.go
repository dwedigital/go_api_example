package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 8.99},
	{ID: "2", Title: "The Wall", Artist: "Pink Floyd", Price: 8.99},
	{ID: "3", Title: "The Division Bell", Artist: "Pink Floyd", Price: 8.99},
}

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(http.StatusOK, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	albumID := c.Param("id")
	for _, album := range albums {
		if album.ID == albumID {
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
}
