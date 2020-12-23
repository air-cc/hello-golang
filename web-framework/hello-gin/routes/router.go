package routes

import (
	"github.com/gin-gonic/gin"
)

// RouterHandler returns all routes
func RouterHandler(r *gin.Engine) {
	rg := r.Group("/")

	// ping
	addRouterGroupPing(rg)

	// users
	addRouterGroupUsers(rg)
}
