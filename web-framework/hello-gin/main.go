package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"iaircc.com/go/demo/hello-gin/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// middleware
	r.Use(func(c *gin.Context) {
		fmt.Println(c.Request.Method, c.Request.URL)
		traceID := uuid.NewV4().String()
		c.Header("trace-id", traceID)
		c.Next()
	})

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
