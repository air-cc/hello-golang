package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"iaircc.com/go/demo/hello-gin/middlewares"
	"iaircc.com/go/demo/hello-gin/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// middleware
	r.Use(middlewares.RequestTrace)

	routes.RouterHandler(r)

	return r
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + port)
}
