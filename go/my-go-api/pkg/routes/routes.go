package routes

import (
	"database/sql"
	"my-go-api/pkg/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(mongoClient *mongo.Client, postgresClient *sql.DB) *gin.Engine {
	r := gin.Default()

	// // Middleware to inject the MongoDB client into the context
	// r.Use(func(c *gin.Context) {
	// 	c.Set("dbClient", client)
	// 	c.Next()
	// })

	// Define routes
	//User
	r.GET("/user", handlers.GetAllUsersHandler(mongoClient))
	r.POST("/user", handlers.SaveUserHandler(mongoClient))
	r.GET("/user/:id", handlers.GetUserByIDHandler(mongoClient))
	r.DELETE("/user/:id", handlers.DeleteUserByIDHandler(mongoClient))
	r.PUT("/user", handlers.UpdateUserHandler(mongoClient))
	r.POST("/user/:n", handlers.CreateMultipleUsersHandler(mongoClient))
	r.POST("/user/go/:n", handlers.CreateMultipleUsersGoRoutineHandler(mongoClient))
	r.GET("/user/count", handlers.CountUsersHandler(mongoClient))
	r.DELETE("/user/all", handlers.DeleteAllUsersHandler(mongoClient))

	//Album
	r.GET("/album", handlers.GetAllAlbumsHandler(postgresClient))
	r.POST("/album", handlers.SaveAlbumHandler(postgresClient))
	r.GET("/album/:id", handlers.GetAlbumByIDHandler(postgresClient))
	r.DELETE("/album/:id", handlers.DeleteAlbumByIDHandler(postgresClient))
	r.POST("/album/:n", handlers.CreateMultipleAlbumsHandler(postgresClient))
	r.POST("/album/go/:n", handlers.CreateMultipleAlbumsEfficientHandler(postgresClient))

	return r
}
