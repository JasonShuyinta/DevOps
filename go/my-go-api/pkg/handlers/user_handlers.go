package handlers

import (
	"my-go-api/pkg/models"
	"my-go-api/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsersHandler(c *gin.Context) {
	client := c.MustGet("dbClient").(*mongo.Client)
	users, err := services.GetAllUsers(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func SaveUserHandler(c *gin.Context) {
	client := c.MustGet("dbClient").(*mongo.Client)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := services.SaveUser(client, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully!"})
}

func GetUserByIDHandler(c *gin.Context) {
	client := c.MustGet("dbClient").(*mongo.Client)
	id := c.Param("id")
	user, err := services.GetUserByID(client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user: " + err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUserByIDHandler(c *gin.Context) {
	client := c.MustGet("dbClient").(*mongo.Client)
	id := c.Param("id")
	if err := services.DeleteUserByID(client, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
