package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type album struct {
	// json dictates what the fields are called
	ID     string  
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

var albums = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 8.99},
	{ID: "2", Title: "The Wall", Artist: "Pink Floyd", Price: 8.99},
	{ID: "3", Title: "The Division Bell", Artist: "Pink Floyd", Price: 8.99},
}



func router() *gin.Engine{
	r := gin.Default()
	// routes
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
	return r
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
	// auto-increment ID by converting len of albums to int + 1
	newAlbum.ID = strconv.FormatInt(int64(len(albums) + 1),10)
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

func main(){
	r := router()
	r.Run(":8080")

}