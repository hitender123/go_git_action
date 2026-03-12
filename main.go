package main

import (
	"log"
)

func main() {
	// Setup Gin router
	r := setupRouter()

	// Start server on port 8080
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
