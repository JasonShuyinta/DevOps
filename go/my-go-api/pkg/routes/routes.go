package routes

import (
	"my-go-api/pkg/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Middleware to inject the MongoDB client into the context
	r.Use(func(c *gin.Context) {
		c.Set("dbClient", client)
		c.Next()
	})

	// Define routes
	r.GET("/get-all-users", handlers.GetAllUsersHandler)
	r.POST("/save-user", handlers.SaveUserHandler)
	r.GET("/get-user/:id", handlers.GetUserByIDHandler)
	r.DELETE("/delete-user/:id", handlers.DeleteUserByIDHandler)

	return r
}
