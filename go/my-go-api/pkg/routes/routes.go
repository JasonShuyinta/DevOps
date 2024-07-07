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
	//User
	r.GET("/user", handlers.GetAllUsersHandler)
	r.POST("/user", handlers.SaveUserHandler)
	r.GET("/user/:id", handlers.GetUserByIDHandler)
	r.DELETE("/user/:id", handlers.DeleteUserByIDHandler)
	r.PUT("/user", handlers.UpdateUserHandler)
	r.POST("/user/:n", handlers.CreateMultipleUsersHandler)
	r.POST("/user/go/:n", handlers.CreateMultipleUsersGoRoutineHandler)
	r.GET("/user/count", handlers.CountUsersHandler)
	r.DELETE("/user/all", handlers.DeleteAllUsersHandler)

	return r
}
