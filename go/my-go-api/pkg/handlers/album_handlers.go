package handlers

import (
	"database/sql"
	"my-go-api/pkg/models"
	"my-go-api/pkg/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllAlbumsHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		albums, err := services.GetAllAlbums(client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching albums: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, albums)
	}
	return gin.HandlerFunc(fn)
}

func SaveAlbumHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var album models.Albums

		if err := c.ShouldBindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		album.CreationTimestamp = time.Now()
		album.UpdateTimestamp = time.Now()

		if err := services.SaveAlbum(client, &album); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving album: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, album)
	}
	return gin.HandlerFunc(fn)
}

func GetAlbumByIDHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		album, err := services.GetAlbumByID(client, id)
		if err != nil {
			if err.Error() == "album with id "+idStr+" not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetchin albumByID: " + err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, album)
	}
	return gin.HandlerFunc(fn)
}

func DeleteAlbumByIDHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		err = services.DeleteAlbumByID(client, id)
		if err != nil {
			if err.Error() == "album with id "+idStr+" not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting albumByID: " + err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, "Album with id "+idStr+" was deleted")
	}
	return gin.HandlerFunc(fn)
}

func CreateMultipleAlbumsHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nStr := c.Param("n")
		n, err := strconv.Atoi(nStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid n format"})
			return
		}

		err = services.CreateMultipleAlbums(client, n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating " + nStr + " albums"})
			return
		}
		c.JSON(http.StatusOK, "Created "+nStr+" albums")
	}
	return gin.HandlerFunc(fn)
}

func CreateMultipleAlbumsEfficientHandler(client *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nStr := c.Param("n")
		n, err := strconv.Atoi(nStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid n format"})
			return
		}

		err = services.CreateMultipleAlbums(client, n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating " + nStr + " albums"})
			return
		}
		c.JSON(http.StatusOK, "Created "+nStr+" albums")
	}
	return gin.HandlerFunc(fn)
}
