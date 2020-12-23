package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iaircc.com/go/demo/hello-gin/services/users"
)

func addRouterGroupUsers(rg *gin.RouterGroup) {
	userRG := rg.Group("/users")

	// 路由参数
	userRG.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		user, ok := users.Get(id)
		if ok == false {
			c.String(http.StatusNotFound, "no user")
			return
		}

		c.JSON(http.StatusOK, user)
	})

	userRG.POST("/", func(c *gin.Context) {
		userBody := users.User{}

		bindErr := c.ShouldBind(&userBody)
		if bindErr != nil {
			c.String(http.StatusBadRequest, "illegal data")
			return
		}

		id := users.Save(userBody)

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	userRG.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		userBody := users.User{}

		bindErr := c.ShouldBind(&userBody)
		if bindErr != nil {
			c.String(http.StatusBadRequest, "illegal data")
			return
		}

		ok := users.Update(id, userBody)

		c.JSON(http.StatusOK, gin.H{
			"ok": ok,
		})
	})

	userRG.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")

		ok := users.Delete(id)

		c.JSON(http.StatusOK, gin.H{
			"ok": ok,
		})
	})
}
