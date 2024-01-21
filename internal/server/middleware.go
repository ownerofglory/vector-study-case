package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerofglory/vector-study-case/internal/logging"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logging.Info("Incoming request")
		logging.Debugf("Request %v %v", ctx.Request.Method, ctx.Request.RequestURI)

		ctx.Next()

		logging.Debugf("Response status %v", ctx.Writer.Status())
	}
}
