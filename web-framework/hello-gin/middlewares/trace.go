package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestTrace 请求追踪中间件
func RequestTrace(ctx *gin.Context) {
	fmt.Println(ctx.Request.Method, ctx.Request.URL)

	traceID := uuid.NewV4().String()
	ctx.Header("trace-id", traceID)

	ctx.Next()
}
