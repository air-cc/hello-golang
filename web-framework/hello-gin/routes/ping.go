package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRouterGroupPing(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("", func(c *gin.Context) {
		c.String(200, "pong")
	})

	ping.GET("/echo", func(c *gin.Context) {
		c.DataFromReader(
			http.StatusOK,
			c.Request.ContentLength,
			c.ContentType(),
			c.Request.Body,
			map[string]string{},
		)
	})
}
