package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ambil port dari environment variable, default 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := gin.Default()

	// Serve static files dari folder ./static
	r.Static("/", "./static")

	log.Printf("listening on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
