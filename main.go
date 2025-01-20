package main

import (
	"fmt"

	"myapp/database"
	"myapp/routes"
	"myapp/workers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Start the worker pool with 4 workers
	go workers.StartWorkerPool(4)
	database.Connect()
	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	fmt.Println("Server running on port 8080")
	router.Run(":8080")

	// Close the task queue and wait for workers to finish
	close(workers.TaskQueue)
	workers.WaitForWorkers()
}
