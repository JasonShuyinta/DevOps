package main

import (
	"my-go-api/pkg/db"
	"my-go-api/pkg/routes"
)

func main() {

	// Connect to mongodb
	client := db.ConnectDB()
	defer db.DisconnectDB()

	// Set up the Gin router
	r := routes.SetupRouter(client)

	// Start http server
	r.Run(":8081")
}
