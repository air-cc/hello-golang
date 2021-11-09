package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iaircc.com/go/demo/hello-gin/model"
	"iaircc.com/go/demo/hello-gin/services/users"
)

func addRouterGroupUsers(rg *gin.RouterGroup) {
	userRG := rg.Group("/users")

	// 路由参数
	userRG.GET("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "illegal params")
			return
		}

		user := users.Get(id)
		if user == nil {
			c.String(http.StatusNotFound, "no user")
			return
		}

		c.JSON(http.StatusOK, *user)
	})

	userRG.POST("/", func(c *gin.Context) {
		userBody := model.User{}

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
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "illegal params")
			return
		}

		userBody := model.User{}
		bindErr := c.ShouldBind(&userBody)
		if bindErr != nil {
			c.String(http.StatusBadRequest, "illegal data")
			return
		}

		userBody.ID = id

		ok := users.Update(userBody)

		c.JSON(http.StatusOK, gin.H{
			"ok": ok,
		})
	})

	userRG.DELETE("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "illegal params")
			return
		}

		ok := users.Delete(id)

		c.JSON(http.StatusOK, gin.H{
			"ok": ok,
		})
	})
}
